// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ContentScanningResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Defines an identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.StringAttribute{
				Description: "The status value for Content Scanning.\nAvailable values: \"enabled\", \"disabled\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
				},
			},
			"modified": schema.StringAttribute{
				Description: "Defines the last modification date (ISO 8601) of the Content Scanning status.",
				Computed:    true,
			},
		},
	}
}

func (r *ContentScanningResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ContentScanningResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
