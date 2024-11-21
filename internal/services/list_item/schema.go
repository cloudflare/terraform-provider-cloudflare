// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ListItemResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"list_id": schema.StringAttribute{
				Description:   "The unique ID of the list.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_identifier": schema.StringAttribute{
				Description:   "Identifier",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"item_id": schema.StringAttribute{
				Description:   "The unique ID of the item in the List.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
										int64validator.OneOf(
											301,
											302,
											307,
											308,
										),
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
				},
			},
			"include_subdomains": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"operation_id": schema.StringAttribute{
				Description: "The unique operation ID of the asynchronous action.",
				Computed:    true,
			},
			"preserve_path_suffix": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"preserve_query_string": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"source_url": schema.StringAttribute{
				Computed: true,
			},
			"status_code": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.OneOf(
						301,
						302,
						307,
						308,
					),
				},
				Default: int64default.StaticInt64(301),
			},
			"subpath_matching": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"target_url": schema.StringAttribute{
				Computed: true,
			},
			"url_hostname": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ListItemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ListItemResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
