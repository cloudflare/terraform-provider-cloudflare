// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ListItemResource)(nil)

func (r *ListItemResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// upgrade from version 1 to version 2 (v4 to v5)
		1: {
			// provider v4 cloudflare_list_item schema
			PriorSchema: &schema.Schema{
				Version: 1,
				Attributes: map[string]schema.Attribute{
					consts.AccountIDSchemaKey: schema.StringAttribute{
						MarkdownDescription: consts.AccountIDSchemaDescription,
						Required:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"list_id": schema.StringAttribute{
						MarkdownDescription: "The list identifier to target for the resource.",
						Required:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: "The list item identifier.",
						Computed:            true,
					},
					"ip": schema.StringAttribute{
						MarkdownDescription: fmt.Sprintf("IP address to include in the list. %s", utils.RenderMustProviderOnlyOneOfDocumentationValuesStringSlice([]string{"ip", "asn", "redirect", "hostname"})),
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("redirect"),
								path.MatchRelative().AtParent().AtName("asn"),
								path.MatchRelative().AtParent().AtName("hostname"),
							),
						},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"asn": schema.Int64Attribute{
						MarkdownDescription: fmt.Sprintf("Autonomous system number to include in the list. %s", utils.RenderMustProviderOnlyOneOfDocumentationValuesStringSlice([]string{"ip", "asn", "redirect", "hostname"})),
						Optional:            true,
						Validators: []validator.Int64{
							int64validator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("redirect"),
								path.MatchRelative().AtParent().AtName("ip"),
								path.MatchRelative().AtParent().AtName("hostname"),
							),
						},
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"comment": schema.StringAttribute{
						MarkdownDescription: "An optional comment for the item.",
						Optional:            true,
					},
				},
				Blocks: map[string]schema.Block{
					"hostname": schema.ListNestedBlock{
						MarkdownDescription: fmt.Sprintf("Hostname to store in the list. %s", utils.RenderMustProviderOnlyOneOfDocumentationValuesStringSlice([]string{"ip", "asn", "redirect", "hostname"})),
						Validators: []validator.List{
							listvalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("redirect"),
								path.MatchRelative().AtParent().AtName("asn"),
								path.MatchRelative().AtParent().AtName("ip"),
							),
							listvalidator.SizeAtMost(1),
						},
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplace(),
						},
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"url_hostname": schema.StringAttribute{
									MarkdownDescription: "The FQDN to match on.",
									Required:            true,
								},
							},
						},
					},
					"redirect": schema.ListNestedBlock{
						MarkdownDescription: fmt.Sprintf("Redirect configuration to store in the list. %s", utils.RenderMustProviderOnlyOneOfDocumentationValuesStringSlice([]string{"ip", "asn", "redirect", "hostname"})),
						Validators: []validator.List{
							listvalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("asn"),
								path.MatchRelative().AtParent().AtName("hostname"),
								path.MatchRelative().AtParent().AtName("ip"),
							),
							listvalidator.SizeAtMost(1),
						},
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplace(),
						},
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"source_url": schema.StringAttribute{
									MarkdownDescription: "The source url of the redirect.",
									Required:            true,
								},
								"target_url": schema.StringAttribute{
									MarkdownDescription: "The target url of the redirect.",
									Required:            true,
								},
								"include_subdomains": schema.BoolAttribute{
									MarkdownDescription: "Whether the redirect also matches subdomains of the source url.",
									Optional:            true,
								},
								"subpath_matching": schema.BoolAttribute{
									MarkdownDescription: "Whether the redirect also matches subpaths of the source url.",
									Optional:            true,
								},
								"status_code": schema.Int64Attribute{
									MarkdownDescription: "The status code to be used when redirecting a request.",
									Optional:            true,
								},
								"preserve_path_suffix": schema.BoolAttribute{
									MarkdownDescription: "Whether the redirect target url should keep the query string of the request's url.",
									Optional:            true,
								},
								"preserve_query_string": schema.BoolAttribute{
									MarkdownDescription: "Whether the redirect target url should keep the query string of the request's url.",
									Optional:            true,
								},
							},
						},
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData ListItemModelV1
				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgradedStateData := ListItemModel{
					AccountID: priorStateData.AccountID,
					ListID:    priorStateData.ListID,
					ID:        priorStateData.ID,
					IP:        priorStateData.IP,
					ASN:       priorStateData.ASN,
					//these were dynamic blocks, but limited to a max of 1
					//Hostname:  priorStateData.Hostname,
					//Redirect:  priorStateData.Redirect,
					Comment: priorStateData.Comment,
				}

				if len(priorStateData.Hostname) > 0 {
					first := priorStateData.Hostname[0]
					nested, diags := customfield.NewObject[ListItemHostnameModel](ctx, &ListItemHostnameModel{
						URLHostname: first.URLHostname,
					})
					resp.Diagnostics.Append(diags...)
					if resp.Diagnostics.HasError() {
						return
					}
					upgradedStateData.Hostname = nested
				}

				if len(priorStateData.Redirect) > 0 {
					first := priorStateData.Redirect[0]
					nested, diags := customfield.NewObject[ListItemRedirectModel](ctx, &ListItemRedirectModel{
						SourceURL:           first.SourceURL,
						TargetURL:           first.TargetURL,
						IncludeSubdomains:   first.IncludeSubdomains,
						SubpathMatching:     first.SubpathMatching,
						StatusCode:          first.StatusCode,
						PreservePathSuffix:  first.PreservePathSuffix,
						PreserveQueryString: first.PreserveQueryString,
					})
					resp.Diagnostics.Append(diags...)
					if resp.Diagnostics.HasError() {
						return
					}

					upgradedStateData.Redirect = nested
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, &upgradedStateData)...)
			},
		},
	}
}

type ListItemModelV1 struct {
	AccountID types.String               `tfsdk:"account_id"`
	ListID    types.String               `tfsdk:"list_id"`
	ID        types.String               `tfsdk:"id"`
	IP        types.String               `tfsdk:"ip"`
	ASN       types.Int64                `tfsdk:"asn"`
	Hostname  []*ListItemHostnameModelV1 `tfsdk:"hostname"`
	Redirect  []*ListItemRedirectModelV1 `tfsdk:"redirect"`
	Comment   types.String               `tfsdk:"comment"`
}

type ListItemHostnameModelV1 struct {
	URLHostname types.String `tfsdk:"url_hostname"`
}

type ListItemRedirectModelV1 struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string"`
}
