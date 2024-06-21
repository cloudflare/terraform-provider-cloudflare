// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r ListItemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Required:    true,
			},
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"item_id": schema.StringAttribute{
				Description: "The unique ID of the item in the List.",
				Optional:    true,
			},
			"asn": schema.Int64Attribute{
				Description: "A non-negative 32 bit integer",
				Optional:    true,
			},
			"comment": schema.StringAttribute{
				Description: "An informative summary of the list item.",
				Optional:    true,
			},
			"hostname": schema.SingleNestedAttribute{
				Description: "Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"url_hostname": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"ip": schema.StringAttribute{
				Description: "An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.",
				Optional:    true,
			},
			"redirect": schema.SingleNestedAttribute{
				Description: "The definition of the redirect.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"source_url": schema.StringAttribute{
						Required: true,
					},
					"target_url": schema.StringAttribute{
						Required: true,
					},
					"include_subdomains": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"preserve_path_suffix": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"preserve_query_string": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"status_code": schema.Int64Attribute{
						Computed: true,
						Optional: true,
						Validators: []validator.Int64{
							int64validator.OneOf(301, 302, 307, 308),
						},
						Default: int64default.StaticInt64(301),
					},
					"subpath_matching": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
				},
			},
		},
	}
}
