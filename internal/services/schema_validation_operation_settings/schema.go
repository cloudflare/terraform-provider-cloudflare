// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_operation_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*SchemaValidationOperationSettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"operation_id": schema.StringAttribute{
				Description:   "UUID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"mitigation_action": schema.StringAttribute{
				Description: "When set, this applies a mitigation action to this operation\n\n  - `\"log\"` - log request when request does not conform to schema for this operation\n  - `\"block\"` - deny access to the site when request does not conform to schema for this operation\n  - `\"none\"` - will skip mitigation for this operation\n  - `null` - clears any mitigation action\nAvailable values: \"log\", \"block\", \"none\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"log",
						"block",
						"none",
					),
				},
			},
		},
	}
}

func (r *SchemaValidationOperationSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SchemaValidationOperationSettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
