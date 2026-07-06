// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSRecordsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"DNS Read",
				"DNS Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"search": schema.StringAttribute{
				Description: "Allows searching in multiple properties of a DNS record simultaneously. This parameter is intended for human users, not automation. Its exact behavior is intentionally left unspecified and is subject to change in the future. This parameter works independently of the `match` setting. For automated searches, please use the other available parameters.",
				Optional:    true,
			},
			"shadowed_by_name": schema.StringAttribute{
				Description: "Filters to records at or below the given NS delegation name, excluding the NS records that form the delegation itself. The value must be a subdomain of the zone; the zone apex is not accepted. Requires `include_shadow_metadata=true`. See [Shadowed records](https://developers.cloudflare.com/dns/manage-dns-records/reference/shadowed-records).",
				Optional:    true,
			},
			"shadowing_name": schema.StringAttribute{
				Description: "Returns NS records that shadow the given name, searching at the name itself and each of its ancestor names within the zone, excluding the zone apex. The value must be a subdomain of the zone; the zone apex is not accepted. See [Shadowed records](https://developers.cloudflare.com/dns/manage-dns-records/reference/shadowed-records).",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "Record type.\nAvailable values: \"A\", \"AAAA\", \"CAA\", \"CERT\", \"CNAME\", \"DNSKEY\", \"DS\", \"HTTPS\", \"LOC\", \"MX\", \"NAPTR\", \"NS\", \"OPENPGPKEY\", \"PTR\", \"SMIMEA\", \"SRV\", \"SSHFP\", \"SVCB\", \"TLSA\", \"TXT\", \"URI\".",
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
						Description: "If this parameter is present, only records *without* a comment are returned.",
						Optional:    true,
					},
					"contains": schema.StringAttribute{
						Description: "Substring of the DNS record comment. Comment filters are case-insensitive.",
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "Suffix of the DNS record comment. Comment filters are case-insensitive.",
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "Exact value of the DNS record comment. Comment filters are case-insensitive.",
						Optional:    true,
					},
					"present": schema.StringAttribute{
						Description: "If this parameter is present, only records *with* a comment are returned.",
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "Prefix of the DNS record comment. Comment filters are case-insensitive.",
						Optional:    true,
					},
				},
			},
			"content": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"contains": schema.StringAttribute{
						Description: "Substring of the DNS record content. Content filters are case-insensitive.",
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "Suffix of the DNS record content. Content filters are case-insensitive.",
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "Exact value of the DNS record content. Content filters are case-insensitive.",
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "Prefix of the DNS record content. Content filters are case-insensitive.",
						Optional:    true,
					},
				},
			},
			"name": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"contains": schema.StringAttribute{
						Description: "Substring of the DNS record name. Name filters are case-insensitive.",
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "Suffix of the DNS record name. Name filters are case-insensitive.",
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "Exact value of the DNS record name. Name filters are case-insensitive.",
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "Prefix of the DNS record name. Name filters are case-insensitive.",
						Optional:    true,
					},
				},
			},
			"tag": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"absent": schema.StringAttribute{
						Description: "Name of a tag which must *not* be present on the DNS record. Tag filters are case-insensitive.",
						Optional:    true,
					},
					"contains": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value contains `<tag-value>`. Tag filters are case-insensitive.",
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value ends with `<tag-value>`. Tag filters are case-insensitive.",
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag filters are case-insensitive.",
						Optional:    true,
					},
					"present": schema.StringAttribute{
						Description: "Name of a tag which must be present on the DNS record. Tag filters are case-insensitive.",
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value starts with `<tag-value>`. Tag filters are case-insensitive.",
						Optional:    true,
					},
				},
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order DNS records in.\nAvailable values: \"asc\", \"desc\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"include_shadow_metadata": schema.BoolAttribute{
				Description: "Whether to include shadow metadata in the `meta` field of each record in the response. See [Shadowed records](https://developers.cloudflare.com/dns/manage-dns-records/reference/shadowed-records).",
				Computed:    true,
				Optional:    true,
			},
			"match": schema.StringAttribute{
				Description: "Whether to match all search requirements or at least one (any). If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead. Note that the interaction between tag filters is controlled by the `tag-match` parameter instead.\nAvailable values: \"any\", \"all\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("any", "all"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Field to order DNS records by.\nAvailable values: \"type\", \"name\", \"content\", \"ttl\", \"proxied\".",
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
				Description: "Whether to match all tag search requirements or at least one (any). If set to `all`, acts like a logical AND between tag filters. If set to `any`, acts like a logical OR instead. Note that the regular `match` parameter is still used to combine the resulting condition with other filters that aren't related to tags.\nAvailable values: \"any\", \"all\".",
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
			"result": schema.DynamicAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
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
