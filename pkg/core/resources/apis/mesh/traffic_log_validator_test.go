package mesh_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/Kong/kuma/pkg/core/resources/apis/mesh"

	util_proto "github.com/Kong/kuma/pkg/util/proto"
	"github.com/ghodss/yaml"
)

var _ = Describe("TrafficLog", func() {
	Describe("Validate()", func() {
		DescribeTable("should pass validation",
			func(trafficLogYAML string) {
				// setup
				trafficLog := TrafficLogResource{}

				// when
				err := util_proto.FromYAML([]byte(trafficLogYAML), &trafficLog.Spec)
				// then
				Expect(err).ToNot(HaveOccurred())

				// when
				verr := trafficLog.Validate()

				// then
				Expect(verr).ToNot(HaveOccurred())
			},
			Entry("full example", `
                selectors:
                - match:
                    region: eu
                conf:
                  backend: file`,
			),
			Entry("empty backend", `
                selectors:
                - match:
                    region: eu
                conf:
                  backend: # backend can be empty, default backend from mesh is chosen`,
			),
			Entry("empty conf", `
                selectors:
                - match:
                    region: eu`,
			),
		)

		type testCase struct {
			trafficLog string
			expected   string
		}
		DescribeTable("should validate all fields and return as much individual errors as possible",
			func(given testCase) {
				// setup
				trafficLog := TrafficLogResource{}

				// when
				err := util_proto.FromYAML([]byte(given.trafficLog), &trafficLog.Spec)
				// then
				Expect(err).ToNot(HaveOccurred())

				// when
				verr := trafficLog.Validate()
				// and
				actual, err := yaml.Marshal(verr)

				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(actual).To(MatchYAML(given.expected))
			},
			Entry("empty spec", testCase{
				trafficLog: ``,
				expected: `
                violations:
                - field: selectors
                  message: must have at least one element
`,
			}),
			Entry("selectors without tags", testCase{
				trafficLog: `
                selectors:
                - match: {}
`,
				expected: `
                violations:
                - field: selectors[0].match
                  message: must have at least one tag
`,
			}),
			Entry("selectors with empty tags values", testCase{
				trafficLog: `
                selectors:
                - match:
                    service:
                    region:
`,
				expected: `
                violations:
                - field: selectors[0].match["region"]
                  message: tag value must be non-empty
                - field: selectors[0].match["service"]
                  message: tag value must be non-empty
`,
			}),
			Entry("multiple selectors", testCase{
				trafficLog: `
                selectors:
                - match:
                    service:
                    region:
                - match: {}
`,
				expected: `
                violations:
                - field: selectors[0].match["region"]
                  message: tag value must be non-empty
                - field: selectors[0].match["service"]
                  message: tag value must be non-empty
                - field: selectors[1].match
                  message: must have at least one tag
`,
			}),
		)
	})
})
