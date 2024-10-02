package list

import (
	"context"
	"fmt"
	"regexp"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *ListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides Lists (IPs, Redirects, Hostname, ASNs) to be used in Edge Rules Engine across all zones within the same account.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				Description: consts.AccountIDSchemaDescription,
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the list.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-z_]+$"), "List name must only contain lowercase letters, numbers, and underscores"),
				},
			},
			"description": schema.StringAttribute{
				Description: "An optional description of the list.",
				Optional:    true,
			},
			"kind": schema.StringAttribute{
				Description: fmt.Sprintf("The type of items the list will contain. %s", utils.RenderMustProviderOnlyOneOfDocumentationValuesStringSlice([]string{"ip", "redirect", "hostname", "asn"})),
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("ip", "redirect", "hostname", "asn"),
				},
			},
			"item": schema.SetNestedAttribute{
				Description: "The items in the list.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.ListNestedAttribute{
							Required: true,
							Validators: []validator.List{
								listvalidator.SizeBetween(1, 1),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("redirect"),
												path.MatchRelative().AtParent().AtName("asn"),
												path.MatchRelative().AtParent().AtName("hostname"),
											),
										},
									},
									"redirect": schema.ListNestedAttribute{
										Optional: true,
										Validators: []validator.List{
											listvalidator.SizeBetween(1, 1),
										},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"source_url": schema.StringAttribute{
													Description: "The source url of the redirect.",
													Required:    true,
												},
												"target_url": schema.StringAttribute{
													Description: "The target url of the redirect.",
													Required:    true,
												},
												"include_subdomains": schema.StringAttribute{
													Description: fmt.Sprintf("Whether the redirect also matches subdomains of the source url. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.OneOf("disabled", "enabled"),
													},
												},
												"subpath_matching": schema.StringAttribute{
													Description: fmt.Sprintf("Whether the redirect also matches subpaths of the source url. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.OneOf("disabled", "enabled"),
													},
												},
												"status_code": schema.Int64Attribute{
													Description: "The status code to be used when redirecting a request.",
													Optional:    true,
												},
												"preserve_query_string": schema.StringAttribute{
													Description: fmt.Sprintf("Whether the redirect target url should keep the query string of the request's url. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.OneOf("disabled", "enabled"),
													},
												},
												"preserve_path_suffix": schema.StringAttribute{
													Description: fmt.Sprintf("Whether to preserve the path suffix when doing subpath matching. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.OneOf("disabled", "enabled"),
													},
												},
											},
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("hostname"),
													path.MatchRelative().AtParent().AtName("ip"),
													path.MatchRelative().AtParent().AtName("asn"),
												),
											},
										},
									},
									"asn": schema.Int64Attribute{
										Optional: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("redirect"),
												path.MatchRelative().AtParent().AtName("ip"),
												path.MatchRelative().AtParent().AtName("hostname"),
											),
										},
									},
									"hostname": schema.ListNestedAttribute{
										Optional: true,
										Validators: []validator.List{
											listvalidator.SizeBetween(1, 1),
										},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"url_hostname": schema.StringAttribute{
													Description: "The FQDN to match on. Wildcard sub-domain matching is allowed. Eg. *.abc.com",
													Required:    true,
												},
											},
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("redirect"),
													path.MatchRelative().AtParent().AtName("ip"),
													path.MatchRelative().AtParent().AtName("asn"),
												),
											},
										},
									},
								},
							},
						},
						"comment": schema.StringAttribute{
							Description: "An optional comment for the item.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}
