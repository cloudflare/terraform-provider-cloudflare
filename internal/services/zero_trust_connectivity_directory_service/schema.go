// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_directory_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustConnectivityDirectoryServiceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Account identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"service_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Description: `Available values: "http".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("http"),
				},
			},
			"host": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"hostname": schema.StringAttribute{
						Optional: true,
					},
					"ipv4": schema.StringAttribute{
						Optional: true,
					},
					"ipv6": schema.StringAttribute{
						Optional: true,
					},
					"network": schema.StringAttribute{
						Optional:   true,
						CustomType: jsontypes.NormalizedType{},
					},
					"resolver_network": schema.StringAttribute{
						Optional:   true,
						CustomType: jsontypes.NormalizedType{},
					},
				},
			},
			"http_port": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"https_port": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
		},
	}
}

func (r *ZeroTrustConnectivityDirectoryServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustConnectivityDirectoryServiceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
