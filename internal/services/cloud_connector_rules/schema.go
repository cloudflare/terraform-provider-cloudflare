// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*CloudConnectorRulesResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.ListNestedAttribute{
				Description: "List of Cloud Connector rules",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional: true,
						},
						"description": schema.StringAttribute{
							Optional: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
						},
						"expression": schema.StringAttribute{
							Optional: true,
						},
						"parameters": schema.SingleNestedAttribute{
							Description: "Parameters of Cloud Connector Rule",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Description: "Host to perform Cloud Connection to",
									Optional:    true,
								},
							},
						},
						"provider": schema.StringAttribute{
							Description: "Cloud Provider type",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"aws_s3",
									"r2",
									"gcp_storage",
									"azure_storage",
								),
							},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *CloudConnectorRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CloudConnectorRulesResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
