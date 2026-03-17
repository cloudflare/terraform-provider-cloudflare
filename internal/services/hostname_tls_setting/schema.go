// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*HostnameTLSSettingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The TLS Setting name. The value type depends on the setting:\n- `ciphers`: value is an array of cipher suite strings (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`)\n- `min_tls_version`: value is a TLS version string (`\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"`)\n- `http2`: value is `\"on\"` or `\"off\"`\nAvailable values: \"ciphers\", \"min_tls_version\", \"http2\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ciphers",
						"min_tls_version",
						"http2",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"setting_id": schema.StringAttribute{
				Description: "The TLS Setting name. The value type depends on the setting:\n- `ciphers`: value is an array of cipher suite strings (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`)\n- `min_tls_version`: value is a TLS version string (`\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"`)\n- `http2`: value is `\"on\"` or `\"off\"`\nAvailable values: \"ciphers\", \"min_tls_version\", \"http2\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ciphers",
						"min_tls_version",
						"http2",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description:   "The hostname for which the tls settings are set.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.DynamicAttribute{
				Description: "The TLS setting value. The type depends on the `setting_id` used in the request path:\n- `ciphers`: an array of allowed cipher suite strings in BoringSSL format (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`)\n- `min_tls_version`: a string indicating the minimum TLS version — one of `\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"` (e.g., `\"1.2\"`)\n- `http2`: a string indicating whether HTTP/2 is enabled — `\"on\"` or `\"off\"` (e.g., `\"on\"`)\nAvailable values: \"1.0\", \"1.1\", \"1.2\", \"1.3\", \"on\", \"off\".",
				Required:    true,
				CustomType:  customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{
					customfield.NormalizeDynamicPlanModifier(),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "This is the time the tls setting was originally created for this hostname.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Deployment status for the given tls setting.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "This is the time the tls setting was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *HostnameTLSSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *HostnameTLSSettingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
