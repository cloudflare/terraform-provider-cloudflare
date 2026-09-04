package r2_custom_domain

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestR2CustomDomainDataSourceToReadParamsJurisdiction(t *testing.T) {
	t.Parallel()

	params, diags := (&R2CustomDomainDataSourceModel{Jurisdiction: types.StringValue("us")}).toReadParams(context.Background())
	if diags.HasError() {
		t.Fatal(diags)
	}
	if !params.Jurisdiction.Present || string(params.Jurisdiction.Value) != "us" {
		t.Fatalf("expected us jurisdiction, got %#v", params.Jurisdiction)
	}

	params, diags = (&R2CustomDomainDataSourceModel{Jurisdiction: types.StringNull()}).toReadParams(context.Background())
	if diags.HasError() {
		t.Fatal(diags)
	}
	if params.Jurisdiction.Present {
		t.Fatalf("expected omitted jurisdiction, got %#v", params.Jurisdiction)
	}
}
