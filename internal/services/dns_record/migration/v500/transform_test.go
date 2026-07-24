package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformLegacyURIUsesRootPriorityAndPreservesZeroWeight(t *testing.T) {
	t.Parallel()

	target, diags := Transform(context.Background(), SourceCloudflareRecordModel{
		ZoneID:   types.StringValue("zone"),
		Name:     types.StringValue("_https.example.com"),
		Type:     types.StringValue("URI"),
		Priority: types.Int64Value(20),
		Data: []SourceDataModel{{
			Priority: types.Int64Value(0), // SDKv2 default for an unset data field.
			Weight:   types.Int64Value(0),
			Target:   types.StringValue("https://example.com/"),
		}},
	})
	if diags.HasError() {
		t.Fatalf("transform diagnostics: %v", diags)
	}
	if target.Data == nil {
		t.Fatal("data was not migrated")
	}
	if got := target.Data.Priority.ValueFloat64(); got != 20 {
		t.Errorf("data.priority = %v, want 20", got)
	}
	if target.Data.Weight.IsNull() || target.Data.Weight.ValueFloat64() != 0 {
		t.Errorf("data.weight = %s, want explicit zero", target.Data.Weight)
	}
}

func TestTransformLegacyCAAKeepsUnrelatedSyntheticZerosNull(t *testing.T) {
	t.Parallel()

	target, diags := Transform(context.Background(), SourceCloudflareRecordModel{
		ZoneID: types.StringValue("zone"),
		Name:   types.StringValue("example.com"),
		Type:   types.StringValue("CAA"),
		Data: []SourceDataModel{{
			Priority: types.Int64Value(0),
			Weight:   types.Int64Value(0),
			Port:     types.Int64Value(0),
		}},
	})
	if diags.HasError() {
		t.Fatalf("transform diagnostics: %v", diags)
	}
	if target.Data == nil {
		t.Fatal("data was not migrated")
	}
	if !target.Data.Priority.IsNull() || !target.Data.Weight.IsNull() || !target.Data.Port.IsNull() {
		t.Errorf("synthetic zero fields were retained: priority=%s weight=%s port=%s", target.Data.Priority, target.Data.Weight, target.Data.Port)
	}
}
