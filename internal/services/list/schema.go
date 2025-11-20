// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"
	"net/url"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ListResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique ID of the list.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID for this resource.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"kind": schema.StringAttribute{
				Description: "The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).\nAvailable values: \"ip\", \"redirect\", \"hostname\", \"asn\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ip",
						"redirect",
						"hostname",
						"asn",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "An informative name for the list. Use this name in filter and rule expressions.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the list.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the list was created.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the list was last modified.",
				Computed:    true,
			},
			"num_items": schema.Float64Attribute{
				Description: "The number of items in the list.",
				Computed:    true,
			},
			"num_referencing_filters": schema.Float64Attribute{
				Description: "The number of [filters](/api/resources/filters/) referencing the list.",
				Computed:    true,
			},
			"items": schema.SetNestedAttribute{
				Description:   "The items in the list. If set, this overwrites all items in the list. Do not use with `cloudflare_list_item`.",
				Optional:      true,
				CustomType:    customfield.NewNestedObjectSetType[ListItemModel](ctx),
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"asn": schema.Int64Attribute{
							Description: "A non-negative 32 bit integer",
							Optional:    true,
							Validators: []validator.Int64{
								int64validator.ConflictsWith(path.MatchRelative().AtParent().AtName("ip")),
								int64validator.ConflictsWith(path.MatchRelative().AtParent().AtName("hostname")),
								int64validator.ConflictsWith(path.MatchRelative().AtParent().AtName("redirect")),
							},
						},
						"comment": schema.StringAttribute{
							Description: "An informative summary of the list item.",
							Optional:    true,
						},
						"hostname": schema.SingleNestedAttribute{
							Description: "Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).",
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[ListItemHostnameModel](ctx),
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ip")),
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("asn")),
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("redirect")),
							},
							Attributes: map[string]schema.Attribute{
								"url_hostname": schema.StringAttribute{
									Required: true,
								},
								"exclude_exact_hostname": schema.BoolAttribute{
									Description: "Only applies to wildcard hostnames (e.g., *.example.com). When true (default), only subdomains are blocked. When false, both the root domain and subdomains are blocked.",
									Optional:    true,
								},
							},
						},
						"ip": schema.StringAttribute{
							Description: "An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("asn")),
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("hostname")),
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("redirect")),
							},
						},
						"redirect": schema.SingleNestedAttribute{
							Description: "The definition of the redirect.",
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[ListItemRedirectModel](ctx),
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ip")),
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("asn")),
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("hostname")),
							},
							Attributes: map[string]schema.Attribute{
								"source_url": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										SourceUrlValidator{},
									},
								},
								"target_url": schema.StringAttribute{
									Required: true,
								},
								"include_subdomains": schema.BoolAttribute{
									Optional: true,
								},
								"preserve_path_suffix": schema.BoolAttribute{
									Optional: true,
								},
								"preserve_query_string": schema.BoolAttribute{
									Optional: true,
								},
								"status_code": schema.Int64Attribute{
									Description: "Available values: 301, 302, 307, 308.",
									Optional:    true,
									Validators: []validator.Int64{
										int64validator.OneOf(
											301,
											302,
											307,
											308,
										),
									},
								},
								"subpath_matching": schema.BoolAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ListResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		ListValidator{},
	}
}

type SourceUrlValidator struct{}

func (v SourceUrlValidator) Description(ctx context.Context) string {
	return "Validates that the URL path is not empty."
}

func (v SourceUrlValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that the URL path is not empty."
}

func (v SourceUrlValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	rawUrl := req.ConfigValue.ValueString()
	if !strings.HasPrefix(rawUrl, "http://") && !strings.HasPrefix(rawUrl, "https://") {
		rawUrl = "https://" + rawUrl
	}
	u, err := url.Parse(rawUrl)
	if err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid URL", err.Error())
		return
	}
	if u.Path == "" {
		resp.Diagnostics.AddAttributeError(req.Path, "source_url path is empty", "The source_url path must not be empty, append a '/' at the end of the URL.")
		return
	}
}

type ListValidator struct {
}

func (v ListValidator) Description(context.Context) string {
	return "validates a cloudflare_list_item"
}

func (v ListValidator) MarkdownDescription(context.Context) string {
	return "validates a cloudflare_list_item"
}
func (v ListValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data *ListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	kind := data.Kind.ValueString()

	if !data.Items.IsNull() {
		items, diags := data.Items.AsStructSliceT(ctx)
		resp.Diagnostics.Append(diags...)

		values, diags := data.Items.Value(ctx)
		resp.Diagnostics.Append(diags...)

		for i, item := range items {
			switch kind {
			case "ip":
				if item.IP.IsNull() {
					resp.Diagnostics.AddAttributeError(path.Root("items").AtSetValue(values[i]).AtName("ip"), "`ip` is not set on list item", "Each list item must be the same type as the list kind.")
				}
			case "hostname":
				if item.Hostname.IsNull() {
					resp.Diagnostics.AddAttributeError(path.Root("items").AtSetValue(values[i]).AtName("hostname"), "`hostname` is not set on list item", "Each list item must be the same type as the list kind.")
				}
			case "asn":
				if item.ASN.IsNull() {
					resp.Diagnostics.AddAttributeError(path.Root("items").AtSetValue(values[i]).AtName("asn"), "`asn` is not set on list item", "Each list item must be the same type as the list kind.")
				}
			case "redirect":
				if item.Redirect.IsNull() {
					resp.Diagnostics.AddAttributeError(path.Root("items").AtSetValue(values[i]).AtName("redirect"), "`redirect` is not set on list item", "Each list item must be the same type as the list kind.")
				}
			}

			if !item.Hostname.IsNull() {
				hostname, diag := item.Hostname.Value(context.Background())
				resp.Diagnostics.Append(diag...)
				if strings.HasPrefix(hostname.URLHostname.ValueString(), "*") && hostname.ExcludeExactHostname.IsNull() {
					resp.Diagnostics.AddAttributeError(path.Root("hostname").AtName("exclude_exact_hostname"), "exclude_exact_hostname is required for wildcard hostnames", "exclude_exact_hostname is required for wildcard hostnames, set it to true or false.")
				}

				if !strings.HasPrefix(hostname.URLHostname.ValueString(), "*") && !hostname.ExcludeExactHostname.IsNull() {
					resp.Diagnostics.AddAttributeError(path.Root("hostname").AtName("exclude_exact_hostname"), "exclude_exact_hostname is only allowed for wildcard hostnames", "exclude_exact_hostname is only allowed for wildcard hostnames, remove it from the resource.")
				}
			}
		}
	}
}
