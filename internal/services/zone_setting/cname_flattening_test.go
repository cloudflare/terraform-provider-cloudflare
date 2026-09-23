package zone_setting

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestIsCNAMEFlatteningSetting(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name      string
		settingID types.String
		want      bool
	}{
		{name: "exact", settingID: types.StringValue("cname_flattening"), want: true},
		{name: "case insensitive", settingID: types.StringValue("CNAME_FLATTENING"), want: true},
		{name: "other setting", settingID: types.StringValue("http3"), want: false},
		{name: "null", settingID: types.StringNull(), want: false},
		{name: "unknown", settingID: types.StringUnknown(), want: false},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isCNAMEFlatteningSetting(test.settingID); got != test.want {
				t.Errorf("isCNAMEFlatteningSetting(%q) = %t, want %t", test.settingID, got, test.want)
			}
		})
	}
}

func TestImportCNAMEFlatteningIsRejectedBeforeAPIRequest(t *testing.T) {
	t.Parallel()

	r := &ZoneSettingResource{}
	resp := &resource.ImportStateResponse{}
	r.ImportState(context.Background(), resource.ImportStateRequest{ID: "zone-id/cname_flattening"}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected an error diagnostic")
	}
}
