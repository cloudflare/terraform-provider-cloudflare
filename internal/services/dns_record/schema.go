// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.ResourceWithConfigValidators = &DNSRecordResource{}

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "DNS record name (or @ for the zone apex) in Punycode.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Record type.",
				Required:    true,
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
			"comment": schema.StringAttribute{
				Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "A valid IPv4 address.",
				Optional:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "Required for MX, SRV and URI records; unused by other record types. Records with lower priorities are preferred.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 65535),
				},
			},
			"data": schema.SingleNestedAttribute{
				Description: "Components of a CAA record.",
				Optional:    true,
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
							float64validator.Between(0, 65535),
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
						Description: "Latitude direction.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("N", "S"),
						},
					},
					"lat_minutes": schema.Float64Attribute{
						Description: "Minutes of latitude.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59),
						},
						Default: float64default.StaticFloat64(0),
					},
					"lat_seconds": schema.Float64Attribute{
						Description: "Seconds of latitude.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59.999),
						},
						Default: float64default.StaticFloat64(0),
					},
					"long_degrees": schema.Float64Attribute{
						Description: "Degrees of longitude.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 180),
						},
					},
					"long_direction": schema.StringAttribute{
						Description: "Longitude direction.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("E", "W"),
						},
					},
					"long_minutes": schema.Float64Attribute{
						Description: "Minutes of longitude.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59),
						},
						Default: float64default.StaticFloat64(0),
					},
					"long_seconds": schema.Float64Attribute{
						Description: "Seconds of longitude.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 59.999),
						},
						Default: float64default.StaticFloat64(0),
					},
					"precision_horz": schema.Float64Attribute{
						Description: "Horizontal precision of location.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
						Default: float64default.StaticFloat64(0),
					},
					"precision_vert": schema.Float64Attribute{
						Description: "Vertical precision of location.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
						Default: float64default.StaticFloat64(0),
					},
					"size": schema.Float64Attribute{
						Description: "Size of location in meters.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 90000000),
						},
						Default: float64default.StaticFloat64(0),
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
					"name": schema.StringAttribute{
						Description: "A valid hostname. Deprecated in favor of the regular 'name' outside the data map. This data map field represents the remainder of the full 'name' after the service and protocol.",
						Optional:    true,
					},
					"port": schema.Float64Attribute{
						Description: "The port of the service.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"proto": schema.StringAttribute{
						Description: "A valid protocol, prefixed with an underscore. Deprecated in favor of the regular 'name' outside the data map. This data map field normally represents the second label of that 'name'.",
						Optional:    true,
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
				Validators: []validator.Float64{
					float64validator.Between(30, 86400),
				},
			},
			"tags": schema.ListAttribute{
				Description: "Custom tags for the DNS record. This field has no effect on DNS responses.",
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
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
