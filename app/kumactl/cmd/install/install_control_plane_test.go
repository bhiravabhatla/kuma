package install_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/kumahq/kuma/app/kumactl/cmd"
	"github.com/kumahq/kuma/app/kumactl/cmd/install"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/app/kumactl/pkg/install/data"
	"github.com/kumahq/kuma/pkg/tls"
	kuma_version "github.com/kumahq/kuma/pkg/version"
)

var _ = Describe("kumactl install control-plane", func() {

	var backupNewSelfSignedCert func(string, tls.CertType, ...string) (tls.KeyPair, error)
	BeforeEach(func() {
		backupNewSelfSignedCert = install.NewSelfSignedCert
	})
	AfterEach(func() {
		install.NewSelfSignedCert = backupNewSelfSignedCert
	})

	BeforeEach(func() {
		install.NewSelfSignedCert = func(string, tls.CertType, ...string) (tls.KeyPair, error) {
			return tls.KeyPair{
				CertPEM: []byte("CERT"),
				KeyPEM:  []byte("KEY"),
			}, nil
		}
	})

	var backupBuildInfo kuma_version.BuildInfo
	BeforeEach(func() {
		backupBuildInfo = kuma_version.Build
	})
	AfterEach(func() {
		kuma_version.Build = backupBuildInfo
	})

	BeforeEach(func() {
		kuma_version.Build = kuma_version.BuildInfo{
			Version:   "0.0.1",
			GitTag:    "v0.0.1",
			GitCommit: "91ce236824a9d875601679aa80c63783fb0e8725",
			BuildDate: "2019-08-07T11:26:06Z",
		}
	})

	var stdout *bytes.Buffer
	var stderr *bytes.Buffer

	BeforeEach(func() {
		stdout = &bytes.Buffer{}
		stderr = &bytes.Buffer{}
	})

	type testCase struct {
		extraArgs  []string
		goldenFile string
	}

	DescribeTable("should generate Kubernetes resources",
		func(given testCase) {
			// given
			rootCmd := cmd.DefaultRootCmd()
			rootCmd.SetArgs(append([]string{"install", "control-plane"}, given.extraArgs...))
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// when
			err := rootCmd.Execute()
			// then
			Expect(err).ToNot(HaveOccurred())
			// and
			Expect(stderr.Bytes()).To(BeNil())

			// when
			expected, err := ioutil.ReadFile(filepath.Join("testdata", given.goldenFile))
			// then
			Expect(err).ToNot(HaveOccurred())
			// and
			expectedManifests := data.SplitYAML(data.File{Data: expected})

			// when
			actual := stdout.Bytes()
			// then
			Expect(actual).To(MatchYAML(expected))
			// and
			actualManifests := data.SplitYAML(data.File{Data: actual})

			// and
			Expect(len(actualManifests)).To(Equal(len(expectedManifests)))
			// and
			for i := range expectedManifests {
				Expect(actualManifests[i]).To(MatchYAML(expectedManifests[i]))
			}
		},
		Entry("should generate Kubernetes resources with default settings", testCase{
			extraArgs:  nil,
			goldenFile: "install-control-plane.defaults.golden.yaml",
		}),
		Entry("should generate Kubernetes resources with custom settings", testCase{
			extraArgs: []string{
				"--namespace", "kuma",
				"--image-pull-policy", "Never",
				"--control-plane-version", "greatest",
				"--control-plane-image", "kuma-ci/kuma-cp",
				"--control-plane-service-name", "kuma-ctrl-plane",
				"--admission-server-tls-cert", "AdmissionCert",
				"--admission-server-tls-key", "AdmissionKey",
				"--injector-failure-policy", "Crash",
				"--dataplane-image", "kuma-ci/kuma-dp",
				"--dataplane-init-image", "kuma-ci/kuma-init",
				"--sds-tls-cert", "SdsCert",
				"--sds-tls-key", "SdsKey",
				"--kds-tls-cert", "KdsCert",
				"--kds-tls-key", "KdsKey",
				"--mode", "remote",
				"--kds-global-address", "grpcs://192.168.0.1:5685",
				"--zone", "zone-1",
				"--use-node-port",
			},
			goldenFile: "install-control-plane.overrides.golden.yaml",
		}),
		Entry("should generate Kubernetes resources with CNI plugin", testCase{
			extraArgs: []string{
				"--cni-enabled",
			},
			goldenFile: "install-control-plane.cni-enabled.golden.yaml",
		}),
		Entry("should generate Kubernetes resources for Global", testCase{
			extraArgs: []string{
				"--mode", "global",
			},
			goldenFile: "install-control-plane.global.golden.yaml",
		}),
		Entry("should generate Kubernetes resources for Remote", testCase{
			extraArgs: []string{
				"--mode", "remote",
				"--zone", "zone-1",
				"--kds-global-address", "grpcs://192.168.0.1:5685",
			},
			goldenFile: "install-control-plane.remote.golden.yaml",
		}),
	)
	It("should fail to install control plane when `kumactl install control-plane run with unknown mode`", func() {
		// given
		rootCmd := cmd.DefaultRootCmd()
		rootCmd.SetArgs([]string{"install", "control-plane", "--mode", "test"})
		//when
		err := rootCmd.Execute()
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("invalid mode. Available modes: standalone, remote, global"))
	})
})
