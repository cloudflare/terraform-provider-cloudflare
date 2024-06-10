// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r PagesDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Name of the project.",
				Required:    true,
			},
			"domain_name": schema.StringAttribute{
				Description: "Name of the domain.",
				Optional:    true,
			},
		},
	}
}
