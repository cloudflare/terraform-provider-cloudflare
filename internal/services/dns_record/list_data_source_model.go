// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSRecordsResultDataSourceModel] `json:"result,computed"`
}

type DNSRecordsDataSourceModel struct {
	ZoneID    types.String                                                  `tfsdk:"zone_id" path:"zone_id,required"`
	Search    types.String                                                  `tfsdk:"search" query:"search,optional"`
	Type      types.String                                                  `tfsdk:"type" query:"type,optional"`
	Comment   *DNSRecordsCommentDataSourceModel                             `tfsdk:"comment" query:"comment,optional"`
	Content   *DNSRecordsContentDataSourceModel                             `tfsdk:"content" query:"content,optional"`
	Name      *DNSRecordsNameDataSourceModel                                `tfsdk:"name" query:"name,optional"`
	Tag       *DNSRecordsTagDataSourceModel                                 `tfsdk:"tag" query:"tag,optional"`
	Direction types.String                                                  `tfsdk:"direction" query:"direction,computed_optional"`
	Match     types.String                                                  `tfsdk:"match" query:"match,computed_optional"`
	Order     types.String                                                  `tfsdk:"order" query:"order,computed_optional"`
	Proxied   types.Bool                                                    `tfsdk:"proxied" query:"proxied,computed_optional"`
	TagMatch  types.String                                                  `tfsdk:"tag_match" query:"tag_match,computed_optional"`
	MaxItems  types.Int64                                                   `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DNSRecordsResultDataSourceModel] `tfsdk:"result"`
}

func (m *DNSRecordsDataSourceModel) toListParams(_ context.Context) (params dns.RecordListParams, diags diag.Diagnostics) {
	params = dns.RecordListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if m.Comment != nil {
		paramsComment := dns.RecordListParamsComment{}
		if !m.Comment.Absent.IsNull() {
			paramsComment.Absent = cloudflare.F(m.Comment.Absent.ValueString())
		}
		if !m.Comment.Contains.IsNull() {
			paramsComment.Contains = cloudflare.F(m.Comment.Contains.ValueString())
		}
		if !m.Comment.Endswith.IsNull() {
			paramsComment.Endswith = cloudflare.F(m.Comment.Endswith.ValueString())
		}
		if !m.Comment.Exact.IsNull() {
			paramsComment.Exact = cloudflare.F(m.Comment.Exact.ValueString())
		}
		if !m.Comment.Present.IsNull() {
			paramsComment.Present = cloudflare.F(m.Comment.Present.ValueString())
		}
		if !m.Comment.Startswith.IsNull() {
			paramsComment.Startswith = cloudflare.F(m.Comment.Startswith.ValueString())
		}
		params.Comment = cloudflare.F(paramsComment)
	}
	// if m.Content != nil {
	// 	paramsContent := dns.RecordListParamsContent{}
	// 	if !m.Content.Contains.IsNull() {
	// 		paramsContent.Contains = cloudflare.F(m.Content.Contains.ValueString())
	// 	}
	// 	if !m.Content.Endswith.IsNull() {
	// 		paramsContent.Endswith = cloudflare.F(m.Content.Endswith.ValueString())
	// 	}
	// 	if !m.Content.Exact.IsNull() {
	// 		paramsContent.Exact = cloudflare.F(m.Content.Exact.ValueString())
	// 	}
	// 	if !m.Content.Startswith.IsNull() {
	// 		paramsContent.Startswith = cloudflare.F(m.Content.Startswith.ValueString())
	// 	}
	// 	params.Content = cloudflare.F(paramsContent)
	// }
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(shared.SortDirection(m.Direction.ValueString()))
	}
	if !m.Match.IsNull() {
		params.Match = cloudflare.F(dns.RecordListParamsMatch(m.Match.ValueString()))
	}
	// if m.Name != nil {
	// 	paramsName := dns.RecordListParamsName{}
	// 	if !m.Name.Contains.IsNull() {
	// 		paramsName.Contains = cloudflare.F(m.Name.Contains.ValueString())
	// 	}
	// 	if !m.Name.Endswith.IsNull() {
	// 		paramsName.Endswith = cloudflare.F(m.Name.Endswith.ValueString())
	// 	}
	// 	if !m.Name.Exact.IsNull() {
	// 		paramsName.Exact = cloudflare.F(m.Name.Exact.ValueString())
	// 	}
	// 	if !m.Name.Startswith.IsNull() {
	// 		paramsName.Startswith = cloudflare.F(m.Name.Startswith.ValueString())
	// 	}
	// 	params.Name = cloudflare.F(paramsName)
	// }
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(dns.RecordListParamsOrder(m.Order.ValueString()))
	}
	if !m.Proxied.IsNull() {
		params.Proxied = cloudflare.F(m.Proxied.ValueBool())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}
	if m.Tag != nil {
		paramsTag := dns.RecordListParamsTag{}
		if !m.Tag.Absent.IsNull() {
			paramsTag.Absent = cloudflare.F(m.Tag.Absent.ValueString())
		}
		if !m.Tag.Contains.IsNull() {
			paramsTag.Contains = cloudflare.F(m.Tag.Contains.ValueString())
		}
		if !m.Tag.Endswith.IsNull() {
			paramsTag.Endswith = cloudflare.F(m.Tag.Endswith.ValueString())
		}
		if !m.Tag.Exact.IsNull() {
			paramsTag.Exact = cloudflare.F(m.Tag.Exact.ValueString())
		}
		if !m.Tag.Present.IsNull() {
			paramsTag.Present = cloudflare.F(m.Tag.Present.ValueString())
		}
		if !m.Tag.Startswith.IsNull() {
			paramsTag.Startswith = cloudflare.F(m.Tag.Startswith.ValueString())
		}
		params.Tag = cloudflare.F(paramsTag)
	}
	if !m.TagMatch.IsNull() {
		params.TagMatch = cloudflare.F(dns.RecordListParamsTagMatch(m.TagMatch.ValueString()))
	}
	if !m.Type.IsNull() {
		params.Type = cloudflare.F(dns.RecordListParamsType(m.Type.ValueString()))
	}

	return
}

type DNSRecordsCommentDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent,optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Present    types.String `tfsdk:"present" json:"present,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordsContentDataSourceModel struct {
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordsNameDataSourceModel struct {
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordsTagDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent,optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Present    types.String `tfsdk:"present" json:"present,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type DNSRecordsResultDataSourceModel struct {
	Comment  types.String                                                `tfsdk:"comment" json:"comment,computed"`
	Content  types.String                                                `tfsdk:"content" json:"content,computed"`
	Name     types.String                                                `tfsdk:"name" json:"name,computed"`
	Proxied  types.Bool                                                  `tfsdk:"proxied" json:"proxied,computed"`
	Tags     customfield.List[types.String]                              `tfsdk:"tags" json:"tags,computed"`
	TTL      types.Float64                                               `tfsdk:"ttl" json:"ttl,computed"`
	Type     types.String                                                `tfsdk:"type" json:"type,computed"`
	Data     customfield.NestedObject[DNSRecordsDataDataSourceModel]     `tfsdk:"data" json:"data,computed"`
	Settings customfield.NestedObject[DNSRecordsSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
	Priority types.Float64                                               `tfsdk:"priority" json:"priority,computed"`
}

type DNSRecordsDataDataSourceModel struct {
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

type DNSRecordsSettingsDataSourceModel struct {
	FlattenCNAME types.Bool `tfsdk:"flatten_cname" json:"flatten_cname,computed"`
}
