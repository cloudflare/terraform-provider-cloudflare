// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.ResourceWithConfigValidators = (*UserResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"country": schema.StringAttribute{
				Description: "The country in which the user lives.",
				Optional:    true,
			},
			"first_name": schema.StringAttribute{
				Description: "User's first name",
				Optional:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "User's last name",
				Optional:    true,
			},
			"telephone": schema.StringAttribute{
				Description: "User's telephone number",
				Optional:    true,
			},
			"zipcode": schema.StringAttribute{
				Description: "The zipcode or postal code where the user lives.",
				Optional:    true,
			},
		},
	}
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
