// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r CloudforceOneRequestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"content": schema.StringAttribute{
				Description: "Request content",
				Optional:    true,
			},
			"priority": schema.StringAttribute{
				Description: "Priority for analyzing the request",
				Optional:    true,
			},
			"request_type": schema.StringAttribute{
				Description: "Requested information from request",
				Optional:    true,
			},
			"summary": schema.StringAttribute{
				Description: "Brief description of the request",
				Optional:    true,
			},
			"tlp": schema.StringAttribute{
				Description: "The CISA defined Traffic Light Protocol (TLP)",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("clear", "amber", "amber-strict", "green", "red"),
				},
			},
		},
	}
}
