// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r EmailRoutingAddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Destination address identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_identifier": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"email": schema.StringAttribute{
				Description: "The contact email address of the user.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the destination address has been created.",
				Computed:    true,
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the destination address was last modified.",
				Computed:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Destination address tag. (Deprecated, replaced by destination address identifier)",
				Computed:    true,
			},
			"verified": schema.StringAttribute{
				Description: "The date and time the destination address has been verified. Null means not verified yet.",
				Computed:    true,
			},
		},
	}
}
