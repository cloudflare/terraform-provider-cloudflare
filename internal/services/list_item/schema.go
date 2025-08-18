// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ListItemResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 2,
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "The Account ID for this resource.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"list_id": schema.StringAttribute{
				Description:   "The unique ID of the list.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Description:   "The unique ID of the item in the List.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"asn": schema.Int64Attribute{
				Description:   "A non-negative 32 bit integer",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"comment": schema.StringAttribute{
				Description:   "An informative summary of the list item.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"created_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the item was created.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the item was last modified.",
				Computed:    true,
			},
			"operation_id": schema.StringAttribute{
				Description: "The unique operation ID of the asynchronous action.",
				Computed:    true,
			},
			"hostname": schema.SingleNestedAttribute{
				Description: "Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).",
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ListItemHostnameModel](ctx),
				Attributes: map[string]schema.Attribute{
					"url_hostname": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"exclude_exact_hostname": schema.BoolAttribute{
						Description:   "Only applies to wildcard hostnames (e.g., *.example.com). When true (default), only subdomains are blocked. When false, both the root domain and subdomains are blocked.",
						Optional:      true,
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
					},
				},
			},
			"ip": schema.StringAttribute{
				Description:   "An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"redirect": schema.SingleNestedAttribute{
				Description: "The definition of the redirect.",
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ListItemRedirectModel](ctx),
				Attributes: map[string]schema.Attribute{
					"source_url": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"target_url": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"include_subdomains": schema.BoolAttribute{
						Computed:      true,
						Optional:      true,
						Default:       booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
					},
					"preserve_path_suffix": schema.BoolAttribute{
						Computed:      true,
						Optional:      true,
						Default:       booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
					},
					"preserve_query_string": schema.BoolAttribute{
						Computed:      true,
						Optional:      true,
						Default:       booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
					},
					"status_code": schema.Int64Attribute{
						Description: "Available values: 301, 302, 307, 308.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.OneOf(
								301,
								302,
								307,
								308,
							),
						},
						Default:       int64default.StaticInt64(301),
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
					},
					"subpath_matching": schema.BoolAttribute{
						Computed:      true,
						Optional:      true,
						Default:       booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
					},
				},
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
