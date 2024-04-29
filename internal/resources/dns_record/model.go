// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordResultEnvelope struct {
	Result DNSRecordModel `json:"result,computed"`
}

type DNSRecordModel struct {
	ID       types.String        `tfsdk:"id" json:"id,computed"`
	ZoneID   types.String        `tfsdk:"zone_id" path:"zone_id"`
	Content  types.String        `tfsdk:"content" json:"content"`
	Name     types.String        `tfsdk:"name" json:"name"`
	Type     types.String        `tfsdk:"type" json:"type"`
	Comment  types.String        `tfsdk:"comment" json:"comment"`
	Proxied  types.Bool          `tfsdk:"proxied" json:"proxied"`
	Tags     types.String        `tfsdk:"tags" json:"tags"`
	TTL      types.Float64       `tfsdk:"ttl" json:"ttl"`
	Data     *DNSRecordDataModel `tfsdk:"data" json:"data"`
	Priority types.Float64       `tfsdk:"priority" json:"priority"`
}

type DNSRecordDataModel struct {
	Flags         types.String  `tfsdk:"flags" json:"flags"`
	Tag           types.String  `tfsdk:"tag" json:"tag"`
	Value         types.String  `tfsdk:"value" json:"value"`
	Algorithm     types.Float64 `tfsdk:"algorithm" json:"algorithm"`
	Certificate   types.String  `tfsdk:"certificate" json:"certificate"`
	KeyTag        types.Float64 `tfsdk:"key_tag" json:"key_tag"`
	Type          types.Float64 `tfsdk:"type" json:"type"`
	Protocol      types.Float64 `tfsdk:"protocol" json:"protocol"`
	PublicKey     types.String  `tfsdk:"public_key" json:"public_key"`
	Digest        types.String  `tfsdk:"digest" json:"digest"`
	DigestType    types.Float64 `tfsdk:"digest_type" json:"digest_type"`
	Priority      types.Float64 `tfsdk:"priority" json:"priority"`
	Target        types.String  `tfsdk:"target" json:"target"`
	Altitude      types.Float64 `tfsdk:"altitude" json:"altitude"`
	LatDegrees    types.Float64 `tfsdk:"lat_degrees" json:"lat_degrees"`
	LatDirection  types.String  `tfsdk:"lat_direction" json:"lat_direction"`
	LatMinutes    types.Float64 `tfsdk:"lat_minutes" json:"lat_minutes"`
	LatSeconds    types.Float64 `tfsdk:"lat_seconds" json:"lat_seconds"`
	LongDegrees   types.Float64 `tfsdk:"long_degrees" json:"long_degrees"`
	LongDirection types.String  `tfsdk:"long_direction" json:"long_direction"`
	LongMinutes   types.Float64 `tfsdk:"long_minutes" json:"long_minutes"`
	LongSeconds   types.Float64 `tfsdk:"long_seconds" json:"long_seconds"`
	PrecisionHorz types.Float64 `tfsdk:"precision_horz" json:"precision_horz"`
	PrecisionVert types.Float64 `tfsdk:"precision_vert" json:"precision_vert"`
	Size          types.Float64 `tfsdk:"size" json:"size"`
	Order         types.Float64 `tfsdk:"order" json:"order"`
	Preference    types.Float64 `tfsdk:"preference" json:"preference"`
	Regex         types.String  `tfsdk:"regex" json:"regex"`
	Replacement   types.String  `tfsdk:"replacement" json:"replacement"`
	Service       types.String  `tfsdk:"service" json:"service"`
	MatchingType  types.Float64 `tfsdk:"matching_type" json:"matching_type"`
	Selector      types.Float64 `tfsdk:"selector" json:"selector"`
	Usage         types.Float64 `tfsdk:"usage" json:"usage"`
	Name          types.String  `tfsdk:"name" json:"name"`
	Port          types.Float64 `tfsdk:"port" json:"port"`
	Proto         types.String  `tfsdk:"proto" json:"proto"`
	Weight        types.Float64 `tfsdk:"weight" json:"weight"`
	Fingerprint   types.String  `tfsdk:"fingerprint" json:"fingerprint"`
	Content       types.String  `tfsdk:"content" json:"content"`
}
