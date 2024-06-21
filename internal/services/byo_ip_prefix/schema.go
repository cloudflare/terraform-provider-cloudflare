// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r ByoIPPrefixResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"asn": schema.Int64Attribute{
				Description: "Autonomous System Number (ASN) the prefix will be advertised under.",
				Required:    true,
			},
			"cidr": schema.StringAttribute{
				Description: "IP Prefix in Classless Inter-Domain Routing format.",
				Required:    true,
			},
			"loa_document_id": schema.StringAttribute{
				Description: "Identifier for the uploaded LOA document.",
				Required:    true,
			},
		},
	}
}
