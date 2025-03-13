// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordResultEnvelope struct {
	Result DNSRecordModel `json:"result"`
}

type DNSRecordModel struct {
	ID                types.String                                     `tfsdk:"id" json:"id,computed" path:"dns_record_id,optional"`
	ZoneID            types.String                                     `tfsdk:"zone_id" path:"zone_id,required"`
	Content           types.String                                     `tfsdk:"content" json:"content,computed_optional"`
	Priority          types.Float64                                    `tfsdk:"priority" json:"priority,optional"`
	Type              types.String                                     `tfsdk:"type" json:"type,required"`
	Data              customfield.NestedObject[DNSRecordDataModel]     `tfsdk:"data" json:"data,computed_optional"`
	Settings          customfield.NestedObject[DNSRecordSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
	Comment           types.String                                     `tfsdk:"comment" json:"comment,computed_optional"`
	CommentModifiedOn timetypes.RFC3339                                `tfsdk:"comment_modified_on" json:"comment_modified_on,computed" format:"date-time"`
	CreatedOn         timetypes.RFC3339                                `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn        timetypes.RFC3339                                `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name              types.String                                     `tfsdk:"name" json:"name,required"`
	Proxiable         types.Bool                                       `tfsdk:"proxiable" json:"proxiable,computed"`
	Proxied           types.Bool                                       `tfsdk:"proxied" json:"proxied,computed_optional"`
	TagsModifiedOn    timetypes.RFC3339                                `tfsdk:"tags_modified_on" json:"tags_modified_on,computed" format:"date-time"`
	TTL               types.Float64                                    `tfsdk:"ttl" json:"ttl,required"`
	Tags              customfield.List[types.String]                   `tfsdk:"tags" json:"tags,computed_optional"`
	Meta              jsontypes.Normalized                             `tfsdk:"meta" json:"meta,computed"`
}

func (m DNSRecordModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSRecordModel) MarshalJSONForUpdate(state DNSRecordModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type DNSRecordDataModel struct {
	Flags         types.Float64 `tfsdk:"flags" json:"flags,optional"`
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
	IPV4Only     types.Bool `tfsdk:"ipv4_only" json:"ipv4_only,computed_optional"`
	IPV6Only     types.Bool `tfsdk:"ipv6_only" json:"ipv6_only,computed_optional"`
	FlattenCNAME types.Bool `tfsdk:"flatten_cname" json:"flatten_cname,computed_optional"`
}
