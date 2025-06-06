// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordResultDataSourceEnvelope struct {
	Result DNSRecordDataSourceModel `json:"result,computed"`
}

type DNSRecordDataSourceModel struct {
	DNSRecordID types.String                                               `tfsdk:"dns_record_id" path:"dns_record_id,required"`
	ZoneID      types.String                                               `tfsdk:"zone_id" path:"zone_id,required"`
	Comment     types.String                                               `tfsdk:"comment" json:"comment,computed"`
	Content     types.String                                               `tfsdk:"content" json:"content,computed"`
	Name        types.String                                               `tfsdk:"name" json:"name,computed"`
	Priority    types.Float64                                              `tfsdk:"priority" json:"priority,computed"`
	Proxied     types.Bool                                                 `tfsdk:"proxied" json:"proxied,computed"`
	TTL         types.Float64                                              `tfsdk:"ttl" json:"ttl,computed"`
	Type        types.String                                               `tfsdk:"type" json:"type,computed"`
	Tags        customfield.List[types.String]                             `tfsdk:"tags" json:"tags,computed"`
	Data        customfield.NestedObject[DNSRecordDataDataSourceModel]     `tfsdk:"data" json:"data,computed"`
	Settings    customfield.NestedObject[DNSRecordSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
}

func (m *DNSRecordDataSourceModel) toReadParams(_ context.Context) (params dns.RecordGetParams, diags diag.Diagnostics) {
	params = dns.RecordGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type DNSRecordDataDataSourceModel struct {
	Flags         types.Dynamic `tfsdk:"flags" json:"flags,computed"`
	Tag           types.String  `tfsdk:"tag" json:"tag,computed"`
	Value         types.String  `tfsdk:"value" json:"value,computed"`
	Algorithm     types.Float64 `tfsdk:"algorithm" json:"algorithm,computed"`
	Certificate   types.String  `tfsdk:"certificate" json:"certificate,computed"`
	KeyTag        types.Float64 `tfsdk:"key_tag" json:"key_tag,computed"`
	Type          types.Float64 `tfsdk:"type" json:"type,computed"`
	Protocol      types.Float64 `tfsdk:"protocol" json:"protocol,computed"`
	PublicKey     types.String  `tfsdk:"public_key" json:"public_key,computed"`
	Digest        types.String  `tfsdk:"digest" json:"digest,computed"`
	DigestType    types.Float64 `tfsdk:"digest_type" json:"digest_type,computed"`
	Priority      types.Float64 `tfsdk:"priority" json:"priority,computed"`
	Target        types.String  `tfsdk:"target" json:"target,computed"`
	Altitude      types.Float64 `tfsdk:"altitude" json:"altitude,computed"`
	LatDegrees    types.Float64 `tfsdk:"lat_degrees" json:"lat_degrees,computed"`
	LatDirection  types.String  `tfsdk:"lat_direction" json:"lat_direction,computed"`
	LatMinutes    types.Float64 `tfsdk:"lat_minutes" json:"lat_minutes,computed"`
	LatSeconds    types.Float64 `tfsdk:"lat_seconds" json:"lat_seconds,computed"`
	LongDegrees   types.Float64 `tfsdk:"long_degrees" json:"long_degrees,computed"`
	LongDirection types.String  `tfsdk:"long_direction" json:"long_direction,computed"`
	LongMinutes   types.Float64 `tfsdk:"long_minutes" json:"long_minutes,computed"`
	LongSeconds   types.Float64 `tfsdk:"long_seconds" json:"long_seconds,computed"`
	PrecisionHorz types.Float64 `tfsdk:"precision_horz" json:"precision_horz,computed"`
	PrecisionVert types.Float64 `tfsdk:"precision_vert" json:"precision_vert,computed"`
	Size          types.Float64 `tfsdk:"size" json:"size,computed"`
	Order         types.Float64 `tfsdk:"order" json:"order,computed"`
	Preference    types.Float64 `tfsdk:"preference" json:"preference,computed"`
	Regex         types.String  `tfsdk:"regex" json:"regex,computed"`
	Replacement   types.String  `tfsdk:"replacement" json:"replacement,computed"`
	Service       types.String  `tfsdk:"service" json:"service,computed"`
	MatchingType  types.Float64 `tfsdk:"matching_type" json:"matching_type,computed"`
	Selector      types.Float64 `tfsdk:"selector" json:"selector,computed"`
	Usage         types.Float64 `tfsdk:"usage" json:"usage,computed"`
	Port          types.Float64 `tfsdk:"port" json:"port,computed"`
	Weight        types.Float64 `tfsdk:"weight" json:"weight,computed"`
	Fingerprint   types.String  `tfsdk:"fingerprint" json:"fingerprint,computed"`
}

type DNSRecordSettingsDataSourceModel struct {
	IPV4Only     types.Bool `tfsdk:"ipv4_only" json:"ipv4_only,computed"`
	IPV6Only     types.Bool `tfsdk:"ipv6_only" json:"ipv6_only,computed"`
	FlattenCNAME types.Bool `tfsdk:"flatten_cname" json:"flatten_cname,computed"`
}
