// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname_ipfs_universal_path_content_list_entry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r Web3HostnameIPFSUniversalPathContentListEntryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"content": schema.StringAttribute{
				Description: "CID or content path of content to block.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Type of content list entry to block.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("cid", "content_path"),
				},
			},
			"description": schema.StringAttribute{
				Description: "An optional description of the content list entry.",
				Optional:    true,
			},
		},
	}
}
