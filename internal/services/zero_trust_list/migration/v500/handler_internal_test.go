package v500

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// TestDetectV4ZeroTrustListState is a regression guard for the schema_version=0
// state-upgrade crash (#7133 follow-up): v5 object-format `items` stamped at
// version 0 was routed through the v4 string PriorSchema and failed with
// "AttributeName(\"items\").ElementKeyInt(0): unsupported type json.Delim sent
// as tftypes.String". The detector must classify such state as v5 (false) so
// UpgradeFromV0Ambiguous takes the no-op re-decode path instead of the v4
// string transform.
func TestDetectV4ZeroTrustListState(t *testing.T) {
	cases := []struct {
		name string
		json string
		want bool
	}{
		{
			name: "v5 object items at version 0 (crash trigger)",
			json: `{"name":"x","type":"IP","items":[{"value":"1.1.1.1","description":null},{"value":"8.8.8.8","description":"dns"}]}`,
			want: false,
		},
		{
			name: "v4 string items with items_with_description key",
			json: `{"name":"x","type":"IP","items":["1.1.1.1","2.2.2.2"],"items_with_description":[]}`,
			want: true,
		},
		{
			name: "v4 string items, no items_with_description key (string fallback)",
			json: `{"name":"x","type":"IP","items":["1.1.1.1"]}`,
			want: true,
		},
		{
			name: "v4 items_with_description populated, items absent",
			json: `{"name":"x","type":"SERIAL","items_with_description":[{"value":"abc","description":"laptop"}]}`,
			want: true,
		},
		{
			name: "v5 empty/null items",
			json: `{"name":"x","type":"IP","items":null}`,
			want: false,
		},
		{
			name: "v5 no items key at all",
			json: `{"name":"x","type":"IP"}`,
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := resource.UpgradeStateRequest{
				RawState: &tfprotov6.RawState{JSON: []byte(tc.json)},
			}
			got, err := detectV4ZeroTrustListState(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("detectV4ZeroTrustListState(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestDetectV4ZeroTrustListState_NilRawState(t *testing.T) {
	got, err := detectV4ZeroTrustListState(resource.UpgradeStateRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got {
		t.Fatalf("nil RawState should not be classified as v4")
	}
}
