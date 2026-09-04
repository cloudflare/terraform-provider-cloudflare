package r2_bucket_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_cors"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_event_notification"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_lifecycle"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_lock"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_sippy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_custom_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_managed_domain"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestR2USJurisdictionValidation(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	resourceSchemas := []struct {
		name   string
		schema func(context.Context) resourceschema.Schema
	}{
		{"bucket", r2_bucket.ResourceSchema},
		{"bucket CORS", r2_bucket_cors.ResourceSchema},
		{"bucket event notification", r2_bucket_event_notification.ResourceSchema},
		{"bucket lifecycle", r2_bucket_lifecycle.ResourceSchema},
		{"bucket lock", r2_bucket_lock.ResourceSchema},
		{"bucket Sippy", r2_bucket_sippy.ResourceSchema},
		{"custom domain", r2_custom_domain.ResourceSchema},
		{"managed domain", r2_managed_domain.ResourceSchema},
	}
	for _, test := range resourceSchemas {
		t.Run("resource "+test.name, func(t *testing.T) {
			attribute := test.schema(ctx).Attributes["jurisdiction"].(resourceschema.StringAttribute)
			assertUSJurisdictionValidation(t, ctx, attribute.Validators)
		})
	}

	dataSourceSchemas := []struct {
		name   string
		schema func(context.Context) datasourceschema.Schema
	}{
		{"bucket", r2_bucket.DataSourceSchema},
		{"bucket CORS", r2_bucket_cors.DataSourceSchema},
		{"bucket event notification", r2_bucket_event_notification.DataSourceSchema},
		{"bucket lifecycle", r2_bucket_lifecycle.DataSourceSchema},
		{"bucket lock", r2_bucket_lock.DataSourceSchema},
		{"bucket Sippy", r2_bucket_sippy.DataSourceSchema},
		{"custom domain", r2_custom_domain.DataSourceSchema},
	}
	for _, test := range dataSourceSchemas {
		t.Run("data source "+test.name, func(t *testing.T) {
			attribute := test.schema(ctx).Attributes["jurisdiction"].(datasourceschema.StringAttribute)
			if !attribute.Optional {
				t.Fatal("expected jurisdiction to be optional")
			}
			assertUSJurisdictionValidation(t, ctx, attribute.Validators)
		})
	}
}

func assertUSJurisdictionValidation(t *testing.T, ctx context.Context, validators []validator.String) {
	t.Helper()
	if len(validators) == 0 {
		t.Fatal("expected jurisdiction validator")
	}

	for _, test := range []struct {
		value     string
		wantError bool
	}{
		{value: "us"},
		{value: "invalid", wantError: true},
	} {
		response := &validator.StringResponse{}
		request := validator.StringRequest{ConfigValue: types.StringValue(test.value)}
		for _, jurisdictionValidator := range validators {
			jurisdictionValidator.ValidateString(ctx, request, response)
		}
		if got := response.Diagnostics.HasError(); got != test.wantError {
			t.Fatalf("validation error for %q: got %t, want %t", test.value, got, test.wantError)
		}
	}
}
