// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_keys_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r AccessKeysConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"key_rotation_interval_days": schema.Float64Attribute{
				Description: "The number of days between key rotations.",
				Required:    true,
				Validators: []validator.Float64{
					float64validator.Between(21, 365),
				},
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
		},
	}
}
