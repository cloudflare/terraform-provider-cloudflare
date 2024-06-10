// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_smart_routing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r ArgoSmartRoutingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "Enables Argo Smart Routing.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("on", "off"),
				},
			},
		},
	}
}
