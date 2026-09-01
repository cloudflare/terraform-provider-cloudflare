package r2_bucket_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestR2BucketFedRAMPHighJurisdictionValidation(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	jurisdiction := r2_bucket.ResourceSchema(ctx).Attributes["jurisdiction"].(resourceschema.StringAttribute)
	request := validator.StringRequest{ConfigValue: types.StringValue("fedramp-high")}
	response := &validator.StringResponse{}

	for _, jurisdictionValidator := range jurisdiction.Validators {
		jurisdictionValidator.ValidateString(ctx, request, response)
	}

	if response.Diagnostics.HasError() {
		t.Fatalf("expected fedramp-high to be a valid jurisdiction: %v", response.Diagnostics)
	}
}
