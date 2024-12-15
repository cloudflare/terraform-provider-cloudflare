// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordResultEnvelope struct {
	Result DNSRecordModel `json:"result"`
}

type DNSRecordModel struct {
	ZoneID      types.String                                     `tfsdk:"zone_id" path:"zone_id,required"`
	DNSRecordID types.String                                     `tfsdk:"dns_record_id" path:"dns_record_id,optional"`
	Comment     types.String                                     `tfsdk:"comment" json:"comment,optional"`
	Content     types.String                                     `tfsdk:"content" json:"content,optional"`
	Name        types.String                                     `tfsdk:"name" json:"name,optional"`
	Priority    types.Float64                                    `tfsdk:"priority" json:"priority,optional"`
	Type        types.String                                     `tfsdk:"type" json:"type,optional"`
	Proxied     types.Bool                                       `tfsdk:"proxied" json:"proxied,computed_optional"`
	TTL         types.Float64                                    `tfsdk:"ttl" json:"ttl,computed_optional"`
	Tags        customfield.List[types.String]                   `tfsdk:"tags" json:"tags,computed_optional"`
	Data        customfield.NestedObject[DNSRecordDataModel]     `tfsdk:"data" json:"data,computed_optional"`
	Settings    customfield.NestedObject[DNSRecordSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
}

func (m DNSRecordModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSRecordModel) MarshalJSONForUpdate(state DNSRecordModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type DNSRecordDataModel struct {
	Flags         types.Dynamic `tfsdk:"flags" json:"flags,optional"`
	Tag           types.String  `tfsdk:"tag" json:"tag,optional"`
	Value         types.String  `tfsdk:"value" json:"value,optional"`
	Algorithm     types.Float64 `tfsdk:"algorithm" json:"algorithm,optional"`
	Certificate   types.String  `tfsdk:"certificate" json:"certificate,optional"`
	KeyTag        types.Float64 `tfsdk:"key_tag" json:"key_tag,optional"`
	Type          types.Float64 `tfsdk:"type" json:"type,optional"`
	Protocol      types.Float64 `tfsdk:"protocol" json:"protocol,optional"`
	PublicKey     types.String  `tfsdk:"public_key" json:"public_key,optional"`
	Digest        types.String  `tfsdk:"digest" json:"digest,optional"`
	DigestType    types.Float64 `tfsdk:"digest_type" json:"digest_type,optional"`
	Priority      types.Float64 `tfsdk:"priority" json:"priority,optional"`
	Target        types.String  `tfsdk:"target" json:"target,optional"`
	Altitude      types.Float64 `tfsdk:"altitude" json:"altitude,optional"`
	LatDegrees    types.Float64 `tfsdk:"lat_degrees" json:"lat_degrees,optional"`
	LatDirection  types.String  `tfsdk:"lat_direction" json:"lat_direction,optional"`
	LatMinutes    types.Float64 `tfsdk:"lat_minutes" json:"lat_minutes,computed_optional"`
	LatSeconds    types.Float64 `tfsdk:"lat_seconds" json:"lat_seconds,computed_optional"`
	LongDegrees   types.Float64 `tfsdk:"long_degrees" json:"long_degrees,optional"`
	LongDirection types.String  `tfsdk:"long_direction" json:"long_direction,optional"`
	LongMinutes   types.Float64 `tfsdk:"long_minutes" json:"long_minutes,computed_optional"`
	LongSeconds   types.Float64 `tfsdk:"long_seconds" json:"long_seconds,computed_optional"`
	PrecisionHorz types.Float64 `tfsdk:"precision_horz" json:"precision_horz,computed_optional"`
	PrecisionVert types.Float64 `tfsdk:"precision_vert" json:"precision_vert,computed_optional"`
	Size          types.Float64 `tfsdk:"size" json:"size,computed_optional"`
	Order         types.Float64 `tfsdk:"order" json:"order,optional"`
	Preference    types.Float64 `tfsdk:"preference" json:"preference,optional"`
	Regex         types.String  `tfsdk:"regex" json:"regex,optional"`
	Replacement   types.String  `tfsdk:"replacement" json:"replacement,optional"`
	Service       types.String  `tfsdk:"service" json:"service,optional"`
	MatchingType  types.Float64 `tfsdk:"matching_type" json:"matching_type,optional"`
	Selector      types.Float64 `tfsdk:"selector" json:"selector,optional"`
	Usage         types.Float64 `tfsdk:"usage" json:"usage,optional"`
	Port          types.Float64 `tfsdk:"port" json:"port,optional"`
	Weight        types.Float64 `tfsdk:"weight" json:"weight,optional"`
	Fingerprint   types.String  `tfsdk:"fingerprint" json:"fingerprint,optional"`
}

type DNSRecordSettingsModel struct {
	FlattenCNAME types.Bool `tfsdk:"flatten_cname" json:"flatten_cname,computed_optional"`
}
