package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransform_DefaultsIncludeShadowMetadata(t *testing.T) {
	t.Parallel()
	for _, recordType := range []string{"A", "TXT", "MX"} {
		t.Run(recordType, func(t *testing.T) {
			target, diags := Transform(context.Background(), SourceCloudflareRecordModel{
				ID:      types.StringValue("record-id"),
				ZoneID:  types.StringValue("zone-id"),
				Name:    types.StringValue("example.com"),
				Type:    types.StringValue(recordType),
				Content: types.StringValue("existing-content"),
				TTL:     types.Int64Value(300),
				Proxied: types.BoolValue(false),
			})
			if diags.HasError() {
				t.Fatalf("transform: %v", diags)
			}
			if !target.IncludeShadowMetadata.Equal(types.BoolValue(false)) {
				t.Errorf("migrated shadow metadata = %s, want false", target.IncludeShadowMetadata)
			}
		})
	}
}
