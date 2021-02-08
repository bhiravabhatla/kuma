package install

import (
	"context"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	install_context "github.com/kumahq/kuma/app/kumactl/cmd/install/context"
	"github.com/kumahq/kuma/app/kumactl/pkg/install/k8s"

	k8s_apixv1beta1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"

	"github.com/kumahq/kuma/app/kumactl/pkg/install/data"
)

func newInstallCrdsCmd(ctx *install_context.InstallCrdsContext) *cobra.Command {
	args := ctx.Args

	cmd := &cobra.Command{
		Use:   "crds",
		Short: "Install Kuma Custom Resource Definitions on Kubernetes",
		RunE: func(cmd *cobra.Command, _ []string) error {
			wantCrdFiles, err := ctx.InstallCrdTemplateFiles(args)
			if err != nil {
				return errors.Wrap(err, "Failed to read CRD files")
			}

			if !args.OnlyMissing {
				singleFile := data.JoinYAML(wantCrdFiles)

				if _, err := cmd.OutOrStdout().Write(singleFile.Data); err != nil {
					return errors.Wrap(err, "Failed to output rendered resources")
				}

				return nil
			}

			crdsToInstallMap, err := mapCrdNamesToFiles(wantCrdFiles)
			if err != nil {
				return errors.Wrap(err, "Failed mapping CRD files with CRD names")
			}

			kubeClientConfig, err := k8s.DefaultClientConfig()
			if err != nil {
				return errors.Wrap(err, "Could not detect Kubernetes configuration")
			}

			k8sClient, err := k8s_apixv1beta1client.NewForConfig(kubeClientConfig)
			if err != nil {
				return errors.Wrap(err, "Failed obtaining Kubernetes client")
			}

			crds, err := k8sClient.CustomResourceDefinitions().List(context.Background(), v1.ListOptions{})
			if err != nil {
				return errors.Wrap(err, "Failed obtaining CRDs from Kubernetes cluster")
			}

			installedCrds := ctx.FilterCrdNamesToInstall(getCrdNamesFromList(crds))
			for _, installedCrdName := range installedCrds {
				delete(crdsToInstallMap, installedCrdName)
			}

			var crdsToInstall []data.File
			for _, crd := range crdsToInstallMap {
				crdsToInstall = append(crdsToInstall, crd)
			}

			if len(crdsToInstallMap) > 0 {
				singleFile := data.JoinYAML(crdsToInstall)

				if _, err := cmd.OutOrStdout().Write(singleFile.Data); err != nil {
					return errors.Wrap(err, "Failed to output rendered resources")
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&args.OnlyMissing, "only-missing", false, "install only resources which are not already present in a cluster")

	return cmd
}

func getCrdNamesFromList(crds *v1beta1.CustomResourceDefinitionList) []string {
	var names []string

	for _, crd := range crds.Items {
		names = append(names, crd.Name)
	}

	return names
}

func mapCrdNamesToFiles(files []data.File) (map[string]data.File, error) {
	result := map[string]data.File{}

	for _, file := range files {
		var crd v1beta1.CustomResourceDefinition

		if err := yaml.Unmarshal(file.Data, &crd); err != nil {
			return nil, errors.Wrap(err, "Failed parsing file as CRD")
		}

		result[crd.ObjectMeta.Name] = file
	}

	return result, nil
}