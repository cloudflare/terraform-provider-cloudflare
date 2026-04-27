// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityTrustedDomainsListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloud Email Security: Read",
				"Cloud Email Security: Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "The sorting direction.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"is_recent": schema.BoolAttribute{
				Description: "Filter to show only recently registered domains that are trusted to prevent triggering Suspicious or Malicious dispositions.",
				Optional:    true,
			},
			"is_similarity": schema.BoolAttribute{
				Description: "Filter to show only proximity domains (partner or approved domains with similar spelling to connected domains) that prevent Spoof dispositions.",
				Optional:    true,
			},
			"order": schema.StringAttribute{
				Description: "Field to sort by.\nAvailable values: \"pattern\", \"created_at\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("pattern", "created_at"),
				},
			},
			"pattern": schema.StringAttribute{
				Optional: true,
			},
			"search": schema.StringAttribute{
				Description: "Search term for filtering records. Behavior may change.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EmailSecurityTrustedDomainsListResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Trusted domain identifier",
							Computed:    true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"is_recent": schema.BoolAttribute{
							Description: "Select to prevent recently registered domains from triggering a Suspicious or Malicious disposition.",
							Computed:    true,
						},
						"is_regex": schema.BoolAttribute{
							Computed: true,
						},
						"is_similarity": schema.BoolAttribute{
							Description: "Select for partner or other approved domains that have similar spelling to your connected domains. Prevents listed domains from triggering a Spoof disposition.",
							Computed:    true,
						},
						"last_modified": schema.StringAttribute{
							Description:        "Deprecated, use `modified_at` instead. End of life: November 1, 2026.",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							CustomType:         timetypes.RFC3339Type{},
						},
						"modified_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"pattern": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *EmailSecurityTrustedDomainsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EmailSecurityTrustedDomainsListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
