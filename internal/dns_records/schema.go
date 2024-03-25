// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_records

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r DNSRecordsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"dns_record_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "A valid IPv4 address.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "DNS record name (or @ for the zone apex) in Punycode.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Record type.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("A"),
				},
			},
			"comment": schema.StringAttribute{
				Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
				Optional:    true,
			},
			"proxied": schema.BoolAttribute{
				Description: "Whether the record is receiving the performance and security benefits of Cloudflare.",
				Optional:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Custom tags for the DNS record. This field has no effect on DNS responses.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ttl": schema.Float64Attribute{
				Description: "Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.",
				Optional:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Components of a CAA record.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"flags": schema.Float64Attribute{
						Description: "Flags for the CAA record.",
						Optional:    true,
					},
					"tag": schema.StringAttribute{
						Description: "Name of the property controlled by this record (e.g.: issue, issuewild, iodef).",
						Optional:    true,
					},
					"value": schema.StringAttribute{
						Description: "Value of the record. This field's semantics depend on the chosen tag.",
						Optional:    true,
					},
				},
			},
			"priority": schema.Float64Attribute{
				Description: "Required for MX, SRV and URI records; unused by other record types. Records with lower priorities are preferred.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the record was created.",
				Computed:    true,
			},
			"locked": schema.BoolAttribute{
				Description: "Whether this record can be modified/deleted (true means it's managed by Cloudflare).",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the record was last modified.",
				Computed:    true,
			},
			"proxiable": schema.BoolAttribute{
				Description: "Whether the record can be proxied by Cloudflare or not.",
				Computed:    true,
			},
			"zone_name": schema.StringAttribute{
				Description: "The domain of the record.",
				Computed:    true,
			},
		},
	}
}
