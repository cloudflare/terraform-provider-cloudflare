// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*EmailSecurityImpersonationRegistryResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloud Email Security: Read",
				"Cloud Email Security: Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Impersonation registry entry identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"email": schema.StringAttribute{
				Description: "Email address (or pattern) of the protected identity.",
				Required:    true,
			},
			"is_email_regex": schema.BoolAttribute{
				Description: "Whether `email` is a regular expression instead of a literal address.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Display name of the protected identity.",
				Required:    true,
			},
			"comments": schema.StringAttribute{
				Description: "Optional note describing the entry.",
				Optional:    true,
			},
			"directory_id": schema.Int64Attribute{
				Description: "Identifier of the directory the entry was synced from, when directory-synced.",
				Optional:    true,
			},
			"directory_node_id": schema.Int64Attribute{
				Description: "Identifier of the directory node the entry was synced from, when directory-synced.",
				Optional:    true,
			},
			"external_directory_node_id": schema.StringAttribute{
				Description:        "Deprecated. External identifier of the directory node.",
				Optional:           true,
				DeprecationMessage: "This field is deprecated.",
			},
			"provenance": schema.StringAttribute{
				Description: "Source the entry was created from.\nAvailable values: \"A1S_INTERNAL\", \"SNOOPY-CASB_OFFICE_365\", \"SNOOPY-OFFICE_365\", \"SNOOPY-GOOGLE_DIRECTORY\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"A1S_INTERNAL",
						"SNOOPY-CASB_OFFICE_365",
						"SNOOPY-OFFICE_365",
						"SNOOPY-GOOGLE_DIRECTORY",
					),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"last_modified": schema.StringAttribute{
				Description:        "Deprecated, use `modified_at` instead. End of life: November 1, 2026.",
				Computed:           true,
				DeprecationMessage: "Use `modified_at` instead.",
				CustomType:         timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *EmailSecurityImpersonationRegistryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailSecurityImpersonationRegistryResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
