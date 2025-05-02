// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*Web3HostnameResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Specify the identifier of the hostname.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Specify the identifier of the hostname.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Specify the hostname that points to the target gateway via CNAME.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"target": schema.StringAttribute{
				Description: "Specify the target gateway of the hostname.\nAvailable values: \"ethereum\", \"ipfs\", \"ipfs_universal_path\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ethereum",
						"ipfs",
						"ipfs_universal_path",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "Specify an optional description of the hostname.",
				Optional:    true,
			},
			"dnslink": schema.StringAttribute{
				Description: "Specify the DNSLink value used if the target is ipfs.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Specifies the status of the hostname's activation.\nAvailable values: \"active\", \"pending\", \"deleting\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"pending",
						"deleting",
						"error",
					),
				},
			},
		},
	}
}

func (r *Web3HostnameResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *Web3HostnameResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
