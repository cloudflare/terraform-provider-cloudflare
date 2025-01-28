// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordResultDataSourceEnvelope struct {
	Result DNSRecordDataSourceModel `json:"result,computed"`
}

type DNSRecordResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSRecordDataSourceModel] `json:"result,computed"`
}

type DNSRecordDataSourceModel struct {
	ID                types.String                                               `tfsdk:"id" json:"-,computed"`
	DNSRecordID       types.String                                               `tfsdk:"dns_record_id" path:"dns_record_id,optional"`
	ZoneID            types.String                                               `tfsdk:"zone_id" path:"zone_id,required"`
	Comment           types.String                                               `tfsdk:"comment" json:"comment,computed"`
	CommentModifiedOn timetypes.RFC3339                                          `tfsdk:"comment_modified_on" json:"comment_modified_on,computed" format:"date-time"`
	Content           types.String                                               `tfsdk:"content" json:"content,computed"`
	CreatedOn         timetypes.RFC3339                                          `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn        timetypes.RFC3339                                          `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name              types.String                                               `tfsdk:"name" json:"name,computed"`
	Priority          types.Float64                                              `tfsdk:"priority" json:"priority,computed"`
	Proxiable         types.Bool                                                 `tfsdk:"proxiable" json:"proxiable,computed"`
	Proxied           types.Bool                                                 `tfsdk:"proxied" json:"proxied,computed"`
	TagsModifiedOn    timetypes.RFC3339                                          `tfsdk:"tags_modified_on" json:"tags_modified_on,computed" format:"date-time"`
	TTL               types.Float64                                              `tfsdk:"ttl" json:"ttl,computed"`
	Type              types.String                                               `tfsdk:"type" json:"type,computed"`
	Tags              customfield.List[types.String]                             `tfsdk:"tags" json:"tags,computed"`
	Data              customfield.NestedObject[DNSRecordDataDataSourceModel]     `tfsdk:"data" json:"data,computed"`
	Settings          customfield.NestedObject[DNSRecordSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
	Meta              jsontypes.Normalized                                       `tfsdk:"meta" json:"meta,computed"`
	Filter            *DNSRecordFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *DNSRecordDataSourceModel) toReadParams(_ context.Context) (params dns.RecordGetParams, diags diag.Diagnostics) {
	params = dns.RecordGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *DNSRecordDataSourceModel) toListParams(_ context.Context) (params dns.RecordListParams, diags diag.Diagnostics) {
	params = dns.RecordListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if m.Filter.Comment != nil {
		paramsComment := dns.RecordListParamsComment{}
		if !m.Filter.Comment.Absent.IsNull() {
			paramsComment.Absent = cloudflare.F(m.Filter.Comment.Absent.ValueString())
		}
		if !m.Filter.Comment.Contains.IsNull() {
			paramsComment.Contains = cloudflare.F(m.Filter.Comment.Contains.ValueString())
		}
		if !m.Filter.Comment.Endswith.IsNull() {
			paramsComment.Endswith = cloudflare.F(m.Filter.Comment.Endswith.ValueString())
		}
		if !m.Filter.Comment.Exact.IsNull() {
			paramsComment.Exact = cloudflare.F(m.Filter.Comment.Exact.ValueString())
		}
		if !m.Filter.Comment.Present.IsNull() {
			paramsComment.Present = cloudflare.F(m.Filter.Comment.Present.ValueString())
		}
		if !m.Filter.Comment.Startswith.IsNull() {
			paramsComment.Startswith = cloudflare.F(m.Filter.Comment.Startswith.ValueString())
		}
		params.Comment = cloudflare.F(paramsComment)
	}
	if m.Filter.Content != nil {
		paramsContent := dns.RecordListParamsContent{}
		if !m.Filter.Content.Contains.IsNull() {
			paramsContent.Contains = cloudflare.F(m.Filter.Content.Contains.ValueString())
		}
		if !m.Filter.Content.Endswith.IsNull() {
			paramsContent.Endswith = cloudflare.F(m.Filter.Content.Endswith.ValueString())
		}
		if !m.Filter.Content.Exact.IsNull() {
			paramsContent.Exact = cloudflare.F(m.Filter.Content.Exact.ValueString())
		}
		if !m.Filter.Content.Startswith.IsNull() {
			paramsContent.Startswith = cloudflare.F(m.Filter.Content.Startswith.ValueString())
		}
		params.Content = cloudflare.F(paramsContent)
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(shared.SortDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Match.IsNull() {
		params.Match = cloudflare.F(dns.RecordListParamsMatch(m.Filter.Match.ValueString()))
	}
	if m.Filter.Name != nil {
		paramsName := dns.RecordListParamsName{}
		if !m.Filter.Name.Contains.IsNull() {
			paramsName.Contains = cloudflare.F(m.Filter.Name.Contains.ValueString())
		}
		if !m.Filter.Name.Endswith.IsNull() {
			paramsName.Endswith = cloudflare.F(m.Filter.Name.Endswith.ValueString())
		}
		if !m.Filter.Name.Exact.IsNull() {
			paramsName.Exact = cloudflare.F(m.Filter.Name.Exact.ValueString())
		}
		if !m.Filter.Name.Startswith.IsNull() {
			paramsName.Startswith = cloudflare.F(m.Filter.Name.Startswith.ValueString())
		}
		params.Name = cloudflare.F(paramsName)
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(dns.RecordListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Proxied.IsNull() {
		params.Proxied = cloudflare.F(m.Filter.Proxied.ValueBool())
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}
	if m.Filter.Tag != nil {
		paramsTag := dns.RecordListParamsTag{}
		if !m.Filter.Tag.Absent.IsNull() {
			paramsTag.Absent = cloudflare.F(m.Filter.Tag.Absent.ValueString())
		}
		if !m.Filter.Tag.Contains.IsNull() {
			paramsTag.Contains = cloudflare.F(m.Filter.Tag.Contains.ValueString())
		}
		if !m.Filter.Tag.Endswith.IsNull() {
			paramsTag.Endswith = cloudflare.F(m.Filter.Tag.Endswith.ValueString())
		}
		if !m.Filter.Tag.Exact.IsNull() {
			paramsTag.Exact = cloudflare.F(m.Filter.Tag.Exact.ValueString())
		}
		if !m.Filter.Tag.Present.IsNull() {
			paramsTag.Present = cloudflare.F(m.Filter.Tag.Present.ValueString())
		}
		if !m.Filter.Tag.Startswith.IsNull() {
			paramsTag.Startswith = cloudflare.F(m.Filter.Tag.Startswith.ValueString())
		}
		params.Tag = cloudflare.F(paramsTag)
	}
	if !m.Filter.TagMatch.IsNull() {
		params.TagMatch = cloudflare.F(dns.RecordListParamsTagMatch(m.Filter.TagMatch.ValueString()))
	}
	if !m.Filter.Type.IsNull() {
		params.Type = cloudflare.F(dns.RecordListParamsType(m.Filter.Type.ValueString()))
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

type DNSRecordFindOneByDataSourceModel struct {
	Comment   *DNSRecordCommentDataSourceModel `tfsdk:"comment" query:"comment,optional"`
	Content   *DNSRecordContentDataSourceModel `tfsdk:"content" query:"content,optional"`
	Direction types.String                     `tfsdk:"direction" query:"direction,computed_optional"`
	Match     types.String                     `tfsdk:"match" query:"match,computed_optional"`
	Name      *DNSRecordNameDataSourceModel    `tfsdk:"name" query:"name,optional"`
	Order     types.String                     `tfsdk:"order" query:"order,computed_optional"`
	Proxied   types.Bool                       `tfsdk:"proxied" query:"proxied,computed_optional"`
	Search    types.String                     `tfsdk:"search" query:"search,optional"`
	Tag       *DNSRecordTagDataSourceModel     `tfsdk:"tag" query:"tag,optional"`
	TagMatch  types.String                     `tfsdk:"tag_match" query:"tag_match,computed_optional"`
	Type      types.String                     `tfsdk:"type" query:"type,optional"`
}

type DNSRecordCommentDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent,optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Present    types.String `tfsdk:"present" json:"present,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordContentDataSourceModel struct {
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordNameDataSourceModel struct {
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordTagDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent,optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Present    types.String `tfsdk:"present" json:"present,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}
