// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordResultEnvelope struct {
	Result DNSRecordModel `json:"result"`
}

type DNSRecordModel struct {
	ID                types.String                                 `tfsdk:"id" json:"id,computed"`
	ZoneID            types.String                                 `tfsdk:"zone_id" path:"zone_id"`
	Content           types.String                                 `tfsdk:"content" json:"content"`
	Priority          types.Float64                                `tfsdk:"priority" json:"priority"`
	Type              types.String                                 `tfsdk:"type" json:"type"`
	Data              *DNSRecordDataModel                          `tfsdk:"data" json:"data"`
	Comment           types.String                                 `tfsdk:"comment" json:"comment,computed"`
	CommentModifiedOn timetypes.RFC3339                            `tfsdk:"comment_modified_on" json:"comment_modified_on,computed" format:"date-time"`
	CreatedOn         timetypes.RFC3339                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn        timetypes.RFC3339                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name              types.String                                 `tfsdk:"name" json:"name,computed"`
	Proxiable         types.Bool                                   `tfsdk:"proxiable" json:"proxiable,computed"`
	Proxied           types.Bool                                   `tfsdk:"proxied" json:"proxied,computed"`
	TagsModifiedOn    timetypes.RFC3339                            `tfsdk:"tags_modified_on" json:"tags_modified_on,computed" format:"date-time"`
	TTL               types.Float64                                `tfsdk:"ttl" json:"ttl,computed"`
	Tags              types.List                                   `tfsdk:"tags" json:"tags,computed"`
	Meta              customfield.NestedObject[DNSRecordMetaModel] `tfsdk:"meta" json:"meta,computed"`
}

type DNSRecordDataModel struct {
	Flags         types.Dynamic `tfsdk:"flags" json:"flags"`
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
	LatMinutes    types.Float64 `tfsdk:"lat_minutes" json:"lat_minutes,computed_optional"`
	LatSeconds    types.Float64 `tfsdk:"lat_seconds" json:"lat_seconds,computed_optional"`
	LongDegrees   types.Float64 `tfsdk:"long_degrees" json:"long_degrees"`
	LongDirection types.String  `tfsdk:"long_direction" json:"long_direction"`
	LongMinutes   types.Float64 `tfsdk:"long_minutes" json:"long_minutes,computed_optional"`
	LongSeconds   types.Float64 `tfsdk:"long_seconds" json:"long_seconds,computed_optional"`
	PrecisionHorz types.Float64 `tfsdk:"precision_horz" json:"precision_horz,computed_optional"`
	PrecisionVert types.Float64 `tfsdk:"precision_vert" json:"precision_vert,computed_optional"`
	Size          types.Float64 `tfsdk:"size" json:"size,computed_optional"`
	Order         types.Float64 `tfsdk:"order" json:"order"`
	Preference    types.Float64 `tfsdk:"preference" json:"preference"`
	Regex         types.String  `tfsdk:"regex" json:"regex"`
	Replacement   types.String  `tfsdk:"replacement" json:"replacement"`
	Service       types.String  `tfsdk:"service" json:"service"`
	MatchingType  types.Float64 `tfsdk:"matching_type" json:"matching_type"`
	Selector      types.Float64 `tfsdk:"selector" json:"selector"`
	Usage         types.Float64 `tfsdk:"usage" json:"usage"`
	Port          types.Float64 `tfsdk:"port" json:"port"`
	Weight        types.Float64 `tfsdk:"weight" json:"weight"`
	Fingerprint   types.String  `tfsdk:"fingerprint" json:"fingerprint"`
}

type DNSRecordMetaModel struct {
	AutoAdded types.Bool `tfsdk:"auto_added" json:"auto_added,computed_optional"`
}
