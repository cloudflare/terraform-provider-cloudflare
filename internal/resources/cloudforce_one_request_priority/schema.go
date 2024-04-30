// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_priority

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r CloudforceOneRequestPriorityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"labels": schema.ListAttribute{
				Description: "List of labels",
				Required:    true,
				ElementType: types.StringType,
			},
			"priority": schema.Int64Attribute{
				Description: "Priority",
				Required:    true,
			},
			"requirement": schema.StringAttribute{
				Description: "Requirement",
				Required:    true,
			},
			"tlp": schema.StringAttribute{
				Description: "The CISA defined Traffic Light Protocol (TLP)",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("clear", "amber", "amber-strict", "green", "red"),
				},
			},
		},
	}
}
