package origin_cloud_regions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*OriginCloudRegionsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"mappings": schema.SetNestedAttribute{
				Description: "Set of origin IP to public cloud vendor and region mappings. Up to 3,500 entries per zone.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"origin_ip": schema.StringAttribute{
							Description: "Origin server IP address (IPv4 or IPv6).",
							Required:    true,
						},
						"vendor": schema.StringAttribute{
							Description: "Cloud vendor.\nAvailable values: \"aws\", \"azure\", \"gcp\", \"oci\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("aws", "azure", "gcp", "oci"),
							},
						},
						"region": schema.StringAttribute{
							Description: "Cloud region identifier (e.g. \"us-east-1\" for AWS).",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (r *OriginCloudRegionsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OriginCloudRegionsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
