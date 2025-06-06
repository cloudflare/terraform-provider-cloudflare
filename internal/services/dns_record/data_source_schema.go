// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSRecordDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dns_record_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"comment": schema.StringAttribute{
				Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
				Computed:    true,
			},
			"content": schema.StringAttribute{
				Description: "A valid IPv4 address.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "DNS record name (or @ for the zone apex) in Punycode.",
				Computed:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "Required for MX, SRV and URI records; unused by other record types. Records with lower priorities are preferred.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 65535),
				},
			},
			"proxied": schema.BoolAttribute{
				Description: "Whether the record is receiving the performance and security benefits of Cloudflare.",
				Computed:    true,
			},
			"ttl": schema.Float64Attribute{
				Description: "Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "Record type.\nAvailable values: \"A\", \"AAAA\", \"CNAME\", \"MX\", \"NS\", \"OPENPGPKEY\", \"PTR\", \"TXT\", \"CAA\", \"CERT\", \"DNSKEY\", \"DS\", \"HTTPS\", \"LOC\", \"NAPTR\", \"SMIMEA\", \"SRV\", \"SSHFP\", \"SVCB\", \"TLSA\", \"URI\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"A",
						"AAAA",
						"CNAME",
						"MX",
						"NS",
						"OPENPGPKEY",
						"PTR",
						"TXT",
						"CAA",
						"CERT",
						"DNSKEY",
						"DS",
						"HTTPS",
						"LOC",
						"NAPTR",
						"SMIMEA",
						"SRV",
						"SSHFP",
						"SVCB",
						"TLSA",
						"URI",
					),
				},
			},
			"tags": schema.ListAttribute{
				Description: "Custom tags for the DNS record. This field has no effect on DNS responses.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Components of a CAA record.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[DNSRecordDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"flags": schema.DynamicAttribute{
						Description: "Flags for the CAA record.",
						Computed:    true,
						Validators: []validator.Dynamic{
							customvalidator.AllowedSubtypes(basetypes.Float64Type{}, basetypes.StringType{}),
						},
					},
					"tag": schema.StringAttribute{
						Description: "Name of the property controlled by this record (e.g.: issue, issuewild, iodef).",
						Computed:    true,
					},
					"value": schema.StringAttribute{
						Description: "Value of the record. This field's semantics depend on the chosen tag.",
						Computed:    true,
					},
					"algorithm": schema.Float64Attribute{
						Description: "Algorithm.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"certificate": schema.StringAttribute{
						Description: "Certificate.",
						Computed:    true,
					},
					"key_tag": schema.Float64Attribute{
						Description: "Key Tag.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"type": schema.Float64Attribute{
						Description: "Type.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.AtLeast(0),
						},
					},
					"protocol": schema.Float64Attribute{
						Description: "Protocol.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"public_key": schema.StringAttribute{
						Description: "Public Key.",
						Computed:    true,
					},
					"digest": schema.StringAttribute{
						Description: "Digest.",
						Computed:    true,
					},
					"digest_type": schema.Float64Attribute{
						Description: "Digest Type.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"priority": schema.Float64Attribute{
						Description: "priority.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"target": schema.StringAttribute{
						Description: "target.",
						Computed:    true,
					},
					"altitude": schema.Float64Attribute{
						Description: "Altitude of location in meters.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(-100000, 42849672.95),
						},
					},
					"lat_degrees": schema.Float64Attribute{
						Description: "Degrees of latitude.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90),
						},
					},
					"lat_direction": schema.StringAttribute{
						Description: "Latitude direction.\nAvailable values: \"N\", \"S\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("N", "S"),
						},
					},
					"lat_minutes": schema.Float64Attribute{
						Description: "Minutes of latitude.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59),
						},
					},
					"lat_seconds": schema.Float64Attribute{
						Description: "Seconds of latitude.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59.999),
						},
					},
					"long_degrees": schema.Float64Attribute{
						Description: "Degrees of longitude.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 180),
						},
					},
					"long_direction": schema.StringAttribute{
						Description: "Longitude direction.\nAvailable values: \"E\", \"W\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("E", "W"),
						},
					},
					"long_minutes": schema.Float64Attribute{
						Description: "Minutes of longitude.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59),
						},
					},
					"long_seconds": schema.Float64Attribute{
						Description: "Seconds of longitude.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59.999),
						},
					},
					"precision_horz": schema.Float64Attribute{
						Description: "Horizontal precision of location.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
					},
					"precision_vert": schema.Float64Attribute{
						Description: "Vertical precision of location.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
					},
					"size": schema.Float64Attribute{
						Description: "Size of location in meters.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
					},
					"order": schema.Float64Attribute{
						Description: "Order.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"preference": schema.Float64Attribute{
						Description: "Preference.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"regex": schema.StringAttribute{
						Description: "Regex.",
						Computed:    true,
					},
					"replacement": schema.StringAttribute{
						Description: "Replacement.",
						Computed:    true,
					},
					"service": schema.StringAttribute{
						Description: "Service.",
						Computed:    true,
					},
					"matching_type": schema.Float64Attribute{
						Description: "Matching Type.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"selector": schema.Float64Attribute{
						Description: "Selector.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"usage": schema.Float64Attribute{
						Description: "Usage.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"port": schema.Float64Attribute{
						Description: "The port of the service.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"weight": schema.Float64Attribute{
						Description: "The record weight.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"fingerprint": schema.StringAttribute{
						Description: "fingerprint.",
						Computed:    true,
					},
				},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Settings for the DNS record.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[DNSRecordSettingsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ipv4_only": schema.BoolAttribute{
						Description: "When enabled, only A records will be generated, and AAAA records will not be created. This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.",
						Computed:    true,
					},
					"ipv6_only": schema.BoolAttribute{
						Description: "When enabled, only AAAA records will be generated, and A records will not be created. This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.",
						Computed:    true,
					},
					"flatten_cname": schema.BoolAttribute{
						Description: "If enabled, causes the CNAME record to be resolved externally and the resulting address records (e.g., A and AAAA) to be returned instead of the CNAME record itself. This setting is unavailable for proxied records, since they are always flattened.",
						Computed:    true,
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
	return []datasource.ConfigValidator{}
}
