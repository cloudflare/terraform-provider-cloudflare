// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.ResourceWithConfigValidators = (*DNSRecordResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"comment": schema.StringAttribute{
				Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "A valid IPv4 address.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "DNS record name (or @ for the zone apex) in Punycode.",
				Optional:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "Required for MX, SRV and URI records; unused by other record types. Records with lower priorities are preferred.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 65535),
				},
			},
			"type": schema.StringAttribute{
				Description: "Record type.\nAvailable values: \"A\".",
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
			"proxied": schema.BoolAttribute{
				Description: "Whether the record is receiving the performance and security benefits of Cloudflare.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"ttl": schema.Float64Attribute{
				Description: "Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.",
				Computed:    true,
				Optional:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Custom tags for the DNS record. This field has no effect on DNS responses.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Components of a CAA record.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[DNSRecordDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"flags": schema.DynamicAttribute{
						Description: "Flags for the CAA record.",
						Optional:    true,
						Validators: []validator.Dynamic{
							customvalidator.AllowedSubtypes(basetypes.Float64Type{}, basetypes.StringType{}),
						},
					},
					"tag": schema.StringAttribute{
						Description: "Name of the property controlled by this record (e.g.: issue, issuewild, iodef).",
						Optional:    true,
					},
					"value": schema.StringAttribute{
						Description: "Value of the record. This field's semantics depend on the chosen tag.",
						Optional:    true,
					},
					"algorithm": schema.Float64Attribute{
						Description: "Algorithm.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"certificate": schema.StringAttribute{
						Description: "Certificate.",
						Optional:    true,
					},
					"key_tag": schema.Float64Attribute{
						Description: "Key Tag.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"type": schema.Float64Attribute{
						Description: "Type.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.AtLeast(0),
						},
					},
					"protocol": schema.Float64Attribute{
						Description: "Protocol.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"public_key": schema.StringAttribute{
						Description: "Public Key.",
						Optional:    true,
					},
					"digest": schema.StringAttribute{
						Description: "Digest.",
						Optional:    true,
					},
					"digest_type": schema.Float64Attribute{
						Description: "Digest Type.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"priority": schema.Float64Attribute{
						Description: "priority.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"target": schema.StringAttribute{
						Description: "target.",
						Optional:    true,
					},
					"altitude": schema.Float64Attribute{
						Description: "Altitude of location in meters.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(-100000, 42849672.95),
						},
					},
					"lat_degrees": schema.Float64Attribute{
						Description: "Degrees of latitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90),
						},
					},
					"lat_direction": schema.StringAttribute{
						Description: "Latitude direction.\nAvailable values: \"N\", \"S\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("N", "S"),
						},
					},
					"lat_minutes": schema.Float64Attribute{
						Description: "Minutes of latitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59),
						},
					},
					"lat_seconds": schema.Float64Attribute{
						Description: "Seconds of latitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59.999),
						},
					},
					"long_degrees": schema.Float64Attribute{
						Description: "Degrees of longitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 180),
						},
					},
					"long_direction": schema.StringAttribute{
						Description: "Longitude direction.\nAvailable values: \"E\", \"W\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("E", "W"),
						},
					},
					"long_minutes": schema.Float64Attribute{
						Description: "Minutes of longitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59),
						},
					},
					"long_seconds": schema.Float64Attribute{
						Description: "Seconds of longitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59.999),
						},
					},
					"precision_horz": schema.Float64Attribute{
						Description: "Horizontal precision of location.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
					},
					"precision_vert": schema.Float64Attribute{
						Description: "Vertical precision of location.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
					},
					"size": schema.Float64Attribute{
						Description: "Size of location in meters.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
					},
					"order": schema.Float64Attribute{
						Description: "Order.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"preference": schema.Float64Attribute{
						Description: "Preference.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"regex": schema.StringAttribute{
						Description: "Regex.",
						Optional:    true,
					},
					"replacement": schema.StringAttribute{
						Description: "Replacement.",
						Optional:    true,
					},
					"service": schema.StringAttribute{
						Description: "Service.",
						Optional:    true,
					},
					"matching_type": schema.Float64Attribute{
						Description: "Matching Type.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"selector": schema.Float64Attribute{
						Description: "Selector.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"usage": schema.Float64Attribute{
						Description: "Usage.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 255),
						},
					},
					"port": schema.Float64Attribute{
						Description: "The port of the service.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"weight": schema.Float64Attribute{
						Description: "The record weight.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"fingerprint": schema.StringAttribute{
						Description: "fingerprint.",
						Optional:    true,
					},
				},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Settings for the DNS record.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[DNSRecordSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ipv4_only": schema.BoolAttribute{
						Description: "When enabled, only A records will be generated, and AAAA records will not be created. This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"ipv6_only": schema.BoolAttribute{
						Description: "When enabled, only AAAA records will be generated, and A records will not be created. This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"flatten_cname": schema.BoolAttribute{
						Description: "If enabled, causes the CNAME record to be resolved externally and the resulting address records (e.g., A and AAAA) to be returned instead of the CNAME record itself. This setting is unavailable for proxied records, since they are always flattened.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
				},
			},
			"comment_modified_on": schema.StringAttribute{
				Description: "When the record comment was last modified. Omitted if there is no comment.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"created_on": schema.StringAttribute{
				Description: "When the record was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the record was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"proxiable": schema.BoolAttribute{
				Description: "Whether the record can be proxied by Cloudflare or not.",
				Computed:    true,
			},
			"tags_modified_on": schema.StringAttribute{
				Description: "When the record tags were last modified. Omitted if there are no tags.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"meta": schema.StringAttribute{
				Description: "Extra Cloudflare-specific information about the record.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (r *DNSRecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *DNSRecordResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
