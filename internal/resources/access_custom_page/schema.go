// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_custom_page

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r AccessCustomPageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID",
				Required:    true,
			},
			"custom_html": schema.StringAttribute{
				Description: "Custom page HTML.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Custom page name.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Custom page type.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("identity_denied", "forbidden"),
				},
			},
			"app_count": schema.Int64Attribute{
				Description: "Number of apps the custom page is assigned to.",
				Optional:    true,
			},
		},
	}
}
