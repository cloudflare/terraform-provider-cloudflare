// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_records

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordsResultEnvelope struct {
	Result DNSRecordsModel `json:"result,computed"`
}

type DNSRecordsModel struct {
	ZoneID      types.String         `tfsdk:"zone_id" path:"zone_id"`
	DNSRecordID types.String         `tfsdk:"dns_record_id" path:"dns_record_id"`
	Content     types.String         `tfsdk:"content" json:"content"`
	Name        types.String         `tfsdk:"name" json:"name"`
	Type        types.String         `tfsdk:"type" json:"type"`
	Comment     types.String         `tfsdk:"comment" json:"comment"`
	Proxied     types.Bool           `tfsdk:"proxied" json:"proxied"`
	Tags        []types.String       `tfsdk:"tags" json:"tags"`
	TTL         types.Float64        `tfsdk:"ttl" json:"ttl"`
	Data        *DNSRecordsDataModel `tfsdk:"data" json:"data"`
	Priority    types.Float64        `tfsdk:"priority" json:"priority"`
	ID          types.String         `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String         `tfsdk:"created_on" json:"created_on,computed"`
	Locked      types.Bool           `tfsdk:"locked" json:"locked,computed"`
	ModifiedOn  types.String         `tfsdk:"modified_on" json:"modified_on,computed"`
	Proxiable   types.Bool           `tfsdk:"proxiable" json:"proxiable,computed"`
	ZoneName    types.String         `tfsdk:"zone_name" json:"zone_name,computed"`
}

type DNSRecordsDataModel struct {
	Flags types.Float64 `tfsdk:"flags" json:"flags"`
	Tag   types.String  `tfsdk:"tag" json:"tag"`
	Value types.String  `tfsdk:"value" json:"value"`
}
