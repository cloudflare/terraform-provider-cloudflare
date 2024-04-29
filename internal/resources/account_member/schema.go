// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r AccountMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Membership identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"email": schema.StringAttribute{
				Description: "The contact email address of the user.",
				Required:    true,
			},
			"roles": schema.ListAttribute{
				Description: "Array of roles associated with this member.",
				Required:    true,
				ElementType: types.StringType,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
				},
				Default: stringdefault.StaticString("pending"),
			},
		},
	}
}
