// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*CloudforceOneRequestAssetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "Asset ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"request_id": schema.StringAttribute{
				Description:   "UUID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"page": schema.Int64Attribute{
				Description:   "Page number of results.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"per_page": schema.Int64Attribute{
				Description:   "Number of results per page.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"source": schema.StringAttribute{
				Description: "Asset file to upload.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Defines the asset creation time.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "Asset description.",
				Computed:    true,
			},
			"file_type": schema.StringAttribute{
				Description: "Asset file type.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Asset name.",
				Computed:    true,
			},
		},
	}
}

func (r *CloudforceOneRequestAssetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CloudforceOneRequestAssetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
