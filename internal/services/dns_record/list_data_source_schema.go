// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSRecordsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"content": schema.StringAttribute{
				Description: "DNS record content.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "DNS record name (or @ for the zone apex) in Punycode.",
				Optional:    true,
			},
			"search": schema.StringAttribute{
				Description: "Allows searching in multiple properties of a DNS record simultaneously. This parameter is intended for human users, not automation. Its exact behavior is intentionally left unspecified and is subject to change in the future. This parameter works independently of the `match` setting. For automated searches, please use the other available parameters.\n",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "Record type.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"A",
						"AAAA",
						"CAA",
						"CERT",
						"CNAME",
						"DNSKEY",
						"DS",
						"HTTPS",
						"LOC",
						"MX",
						"NAPTR",
						"NS",
						"OPENPGPKEY",
						"PTR",
						"SMIMEA",
						"SRV",
						"SSHFP",
						"SVCB",
						"TLSA",
						"TXT",
						"URI",
					),
				},
			},
			"comment": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"absent": schema.StringAttribute{
						Description: "If this parameter is present, only records *without* a comment are returned.\n",
						Computed:    true,
						Optional:    true,
					},
					"contains": schema.StringAttribute{
						Description: "Substring of the DNS record comment. Comment filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "Suffix of the DNS record comment. Comment filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "Exact value of the DNS record comment. Comment filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"present": schema.StringAttribute{
						Description: "If this parameter is present, only records *with* a comment are returned.\n",
						Computed:    true,
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "Prefix of the DNS record comment. Comment filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"tag": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"absent": schema.StringAttribute{
						Description: "Name of a tag which must *not* be present on the DNS record. Tag filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"contains": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value contains `<tag-value>`. Tag filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value ends with `<tag-value>`. Tag filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"present": schema.StringAttribute{
						Description: "Name of a tag which must be present on the DNS record. Tag filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value starts with `<tag-value>`. Tag filters are case-insensitive.\n",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order DNS records in.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"match": schema.StringAttribute{
				Description: "Whether to match all search requirements or at least one (any). If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead. Note that the interaction between tag filters is controlled by the `tag-match` parameter instead.\n",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("any", "all"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Field to order DNS records by.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"type",
						"name",
						"content",
						"ttl",
						"proxied",
					),
				},
			},
			"proxied": schema.BoolAttribute{
				Description: "Whether the record is receiving the performance and security benefits of Cloudflare.",
				Computed:    true,
				Optional:    true,
			},
			"tag_match": schema.StringAttribute{
				Description: "Whether to match all tag search requirements or at least one (any). If set to `all`, acts like a logical AND between tag filters. If set to `any`, acts like a logical OR instead. Note that the regular `match` parameter is still used to combine the resulting condition with other filters that aren't related to tags.\n",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("any", "all"),
				},
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
				CustomType:  customfield.NewNestedObjectListType[DNSRecordsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"comment": schema.StringAttribute{
							Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the record was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"meta": schema.StringAttribute{
							Description: "Extra Cloudflare-specific information about the record.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"modified_on": schema.StringAttribute{
							Description: "When the record was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "DNS record name (or @ for the zone apex) in Punycode.",
							Computed:    true,
						},
						"proxiable": schema.BoolAttribute{
							Description: "Whether the record can be proxied by Cloudflare or not.",
							Computed:    true,
						},
						"proxied": schema.BoolAttribute{
							Description: "Whether the record is receiving the performance and security benefits of Cloudflare.",
							Computed:    true,
						},
						"settings": schema.StringAttribute{
							Description: "Settings for the DNS record.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"tags": schema.ListAttribute{
							Description: "Custom tags for the DNS record. This field has no effect on DNS responses.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"ttl": schema.Float64Attribute{
							Description: "Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(30, 86400),
							},
						},
						"comment_modified_on": schema.StringAttribute{
							Description: "When the record comment was last modified. Omitted if there is no comment.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"tags_modified_on": schema.StringAttribute{
							Description: "When the record tags were last modified. Omitted if there are no tags.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *DNSRecordsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DNSRecordsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}