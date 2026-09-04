package r2_bucket_lock

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestR2BucketLockDataSourceToReadParamsJurisdiction(t *testing.T) {
	t.Parallel()

	params, diags := (&R2BucketLockDataSourceModel{Jurisdiction: types.StringValue("us")}).toReadParams(context.Background())
	if diags.HasError() {
		t.Fatal(diags)
	}
	if !params.Jurisdiction.Present || string(params.Jurisdiction.Value) != "us" {
		t.Fatalf("expected us jurisdiction, got %#v", params.Jurisdiction)
	}

	params, diags = (&R2BucketLockDataSourceModel{Jurisdiction: types.StringNull()}).toReadParams(context.Background())
	if diags.HasError() {
		t.Fatal(diags)
	}
	if params.Jurisdiction.Present {
		t.Fatalf("expected omitted jurisdiction, got %#v", params.Jurisdiction)
	}
}
