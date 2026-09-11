// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingDNSDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zone Settings Read",
				"Zone Settings Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"subdomain": schema.StringAttribute{
				Description: "Deprecated. When supplied, the response shape differs from the documented default and is not modeled in generated SDKs. Do not rely on this parameter.",
				Optional:    true,
			},
			"dns": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSDNSDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"content": schema.StringAttribute{
							Description: "DNS record content.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "DNS record name (or @ for the zone apex).",
							Computed:    true,
						},
						"priority": schema.Float64Attribute{
							Description: "Required for MX, SRV and URI records. Unused by other record types. Records with lower priorities are preferred.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 65535),
							},
						},
						"ttl": schema.Float64Attribute{
							Description: "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "DNS record type.\nAvailable values: \"A\", \"AAAA\", \"CNAME\", \"HTTPS\", \"TXT\", \"SRV\", \"LOC\", \"MX\", \"NS\", \"CERT\", \"DNSKEY\", \"DS\", \"NAPTR\", \"SMIMEA\", \"SSHFP\", \"SVCB\", \"TLSA\", \"URI\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"A",
									"AAAA",
									"CNAME",
									"HTTPS",
									"TXT",
									"SRV",
									"LOC",
									"MX",
									"NS",
									"CERT",
									"DNSKEY",
									"DS",
									"NAPTR",
									"SMIMEA",
									"SSHFP",
									"SVCB",
									"TLSA",
									"URI",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *EmailRoutingDNSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailRoutingDNSDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
