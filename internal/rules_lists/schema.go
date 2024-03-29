// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules_lists

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r RulesListsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Optional:    true,
			},
			"kind": schema.StringAttribute{
				Description: "The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ip", "redirect", "hostname", "asn"),
				},
			},
			"name": schema.StringAttribute{
				Description: "An informative name for the list. Use this name in filter and rule expressions.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the list.",
				Optional:    true,
			},
		},
	}
}
