package cmd

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	kuma_dp "github.com/kumahq/kuma/pkg/config/app/kuma-dp"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/model/rest"
	"github.com/kumahq/kuma/pkg/util/template"
)

func readDataplaneResource(cmd *cobra.Command, cfg *kuma_dp.Config) (*rest.Resource, error) {
	var b []byte
	var err error
	// load from file first
	if cfg.DataplaneRuntime.ResourcePath == "-" {
		if b, err = ioutil.ReadAll(cmd.InOrStdin()); err != nil {
			return nil, err
		}
	} else if cfg.DataplaneRuntime.ResourcePath != "" {
		if b, err = ioutil.ReadFile(cfg.DataplaneRuntime.ResourcePath); err != nil {
			return nil, errors.Wrap(err, "error while reading provided file")
		}
	}
	// override with inline resource
	if cfg.DataplaneRuntime.Resource != "" {
		b = []byte(cfg.DataplaneRuntime.Resource)
	}

	if len(b) == 0 {
		return nil, nil
	}

	b = template.Render(string(b), cfg.DataplaneRuntime.ResourceVars)
	runLog.Info("rendered dataplane", "dataplane", string(b))

	res, err := rest.Unmarshall(b)
	if err != nil {
		return nil, err
	}
	if res.Meta.Type != string(core_mesh.DataplaneType) {
		return nil, errors.Errorf("invalid resource of type: %s. Expected: Dataplane", res.Meta.Type)
	}
	if err := core_mesh.ValidateMeta(res.Meta.GetName(), res.Meta.GetMesh(), model.ScopeMesh); err.HasViolations() {
		return nil, &err
	}

	return res, nil
}
