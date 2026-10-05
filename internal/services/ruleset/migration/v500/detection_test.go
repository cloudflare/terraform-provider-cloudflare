package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/stretchr/testify/require"
)

func TestDetectV4RulesetState(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		json string
		v4   bool
		err  bool
	}{
		{"empty rules", `{"rules":[]}`, false, false},
		{"null rules", `{"rules":null}`, false, false},
		{"absent rules", `{}`, false, false},
		{"v4 empty block", `{"rules":[{"action_parameters":[]}]}`, true, false},
		{"v4 later rule", `{"rules":[{}, {"action_parameters":[{}]}]}`, true, false},
		{"v4 logging", `{"rules":[{"action_parameters":null,"logging":[]}]}`, true, false},
		{"v4 ratelimit", `{"rules":[{"ratelimit":[{}]}]}`, true, false},
		{"v4 credential check", `{"rules":[{"exposed_credential_check":[]}]}`, true, false},
		{"v4 metadata", `{"rules":[{"version":null,"last_updated":null}]}`, true, false},
		{"v5 shared lists", `{"rules":[{"action_parameters":{"algorithms":[{"name":"gzip"}],"overrides":{"rules":[],"categories":[]}}}]}`, false, false},
		{"v5 later rule", `{"rules":[{}, {"logging":{"enabled":false}}]}`, false, false},
		{"mixed shapes", `{"rules":[{"action_parameters":[],"logging":{}}]}`, false, true},
		{"invalid block", `{"rules":[{"action_parameters":false}]}`, false, true},
		{"invalid rule", `{"rules":[false]}`, false, true},
		{"invalid rules", `{"rules":{}}`, false, true},
		{"invalid JSON", `{`, false, true},
		{"null root", `null`, false, true},
		{"empty JSON", ``, false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := detectV4RulesetState(resource.UpgradeStateRequest{RawState: &tfprotov6.RawState{JSON: []byte(tc.json)}})
			if tc.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.v4, got)
		})
	}
}

func TestUpgradeFromAmbiguousMissingRawState(t *testing.T) {
	t.Parallel()
	resp := &resource.UpgradeStateResponse{}
	UpgradeFromAmbiguous(context.Background(), resource.UpgradeStateRequest{}, resp)
	require.True(t, resp.Diagnostics.HasError())
	require.Contains(t, resp.Diagnostics[0].Detail(), "missing or empty")
}
