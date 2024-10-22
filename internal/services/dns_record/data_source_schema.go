// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSRecordDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dns_record_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"comment": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"absent": schema.StringAttribute{
								Description: "If this parameter is present, only records *without* a comment are returned.\n",
								Optional:    true,
							},
							"contains": schema.StringAttribute{
								Description: "Substring of the DNS record comment. Comment filters are case-insensitive.\n",
								Optional:    true,
							},
							"endswith": schema.StringAttribute{
								Description: "Suffix of the DNS record comment. Comment filters are case-insensitive.\n",
								Optional:    true,
							},
							"exact": schema.StringAttribute{
								Description: "Exact value of the DNS record comment. Comment filters are case-insensitive.\n",
								Optional:    true,
							},
							"present": schema.StringAttribute{
								Description: "If this parameter is present, only records *with* a comment are returned.\n",
								Optional:    true,
							},
							"startswith": schema.StringAttribute{
								Description: "Prefix of the DNS record comment. Comment filters are case-insensitive.\n",
								Optional:    true,
							},
						},
					},
					"content": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"contains": schema.StringAttribute{
								Description: "Substring of the DNS record content. Content filters are case-insensitive.\n",
								Optional:    true,
							},
							"endswith": schema.StringAttribute{
								Description: "Suffix of the DNS record content. Content filters are case-insensitive.\n",
								Optional:    true,
							},
							"exact": schema.StringAttribute{
								Description: "Exact value of the DNS record content. Content filters are case-insensitive.\n",
								Optional:    true,
							},
							"startswith": schema.StringAttribute{
								Description: "Prefix of the DNS record content. Content filters are case-insensitive.\n",
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
					"name": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"contains": schema.StringAttribute{
								Description: "Substring of the DNS record name. Name filters are case-insensitive.\n",
								Optional:    true,
							},
							"endswith": schema.StringAttribute{
								Description: "Suffix of the DNS record name. Name filters are case-insensitive.\n",
								Optional:    true,
							},
							"exact": schema.StringAttribute{
								Description: "Exact value of the DNS record name. Name filters are case-insensitive.\n",
								Optional:    true,
							},
							"startswith": schema.StringAttribute{
								Description: "Prefix of the DNS record name. Name filters are case-insensitive.\n",
								Optional:    true,
							},
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
					"search": schema.StringAttribute{
						Description: "Allows searching in multiple properties of a DNS record simultaneously. This parameter is intended for human users, not automation. Its exact behavior is intentionally left unspecified and is subject to change in the future. This parameter works independently of the `match` setting. For automated searches, please use the other available parameters.\n",
						Optional:    true,
					},
					"tag": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"absent": schema.StringAttribute{
								Description: "Name of a tag which must *not* be present on the DNS record. Tag filters are case-insensitive.\n",
								Optional:    true,
							},
							"contains": schema.StringAttribute{
								Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value contains `<tag-value>`. Tag filters are case-insensitive.\n",
								Optional:    true,
							},
							"endswith": schema.StringAttribute{
								Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value ends with `<tag-value>`. Tag filters are case-insensitive.\n",
								Optional:    true,
							},
							"exact": schema.StringAttribute{
								Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag filters are case-insensitive.\n",
								Optional:    true,
							},
							"present": schema.StringAttribute{
								Description: "Name of a tag which must be present on the DNS record. Tag filters are case-insensitive.\n",
								Optional:    true,
							},
							"startswith": schema.StringAttribute{
								Description: "A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return DNS records that have a tag named `<tag-name>` whose value starts with `<tag-value>`. Tag filters are case-insensitive.\n",
								Optional:    true,
							},
						},
					},
					"tag_match": schema.StringAttribute{
						Description: "Whether to match all tag search requirements or at least one (any). If set to `all`, acts like a logical AND between tag filters. If set to `any`, acts like a logical OR instead. Note that the regular `match` parameter is still used to combine the resulting condition with other filters that aren't related to tags.\n",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("any", "all"),
						},
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
				},
			},
		},
	}
}

func (d *DNSRecordDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DNSRecordDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("dns_record_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("dns_record_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}
