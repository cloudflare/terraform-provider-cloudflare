// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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
	DNSRecordID       types.String                       `tfsdk:"dns_record_id" path:"dns_record_id,optional"`
	ZoneID            types.String                       `tfsdk:"zone_id" path:"zone_id,optional"`
	Comment           types.String                       `tfsdk:"comment" json:"comment,computed"`
	CommentModifiedOn timetypes.RFC3339                  `tfsdk:"comment_modified_on" json:"comment_modified_on,computed" format:"date-time"`
	CreatedOn         timetypes.RFC3339                  `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ID                types.String                       `tfsdk:"id" json:"id,computed"`
	ModifiedOn        timetypes.RFC3339                  `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name              types.String                       `tfsdk:"name" json:"name,computed"`
	Proxiable         types.Bool                         `tfsdk:"proxiable" json:"proxiable,computed"`
	Proxied           types.Bool                         `tfsdk:"proxied" json:"proxied,computed"`
	TagsModifiedOn    timetypes.RFC3339                  `tfsdk:"tags_modified_on" json:"tags_modified_on,computed" format:"date-time"`
	TTL               types.Float64                      `tfsdk:"ttl" json:"ttl,computed"`
	Tags              customfield.List[types.String]     `tfsdk:"tags" json:"tags,computed"`
	Meta              jsontypes.Normalized               `tfsdk:"meta" json:"meta,computed"`
	Settings          jsontypes.Normalized               `tfsdk:"settings" json:"settings,computed"`
	Filter            *DNSRecordFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *DNSRecordDataSourceModel) toReadParams(_ context.Context) (params dns.RecordGetParams, diags diag.Diagnostics) {
	params = dns.RecordGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *DNSRecordDataSourceModel) toListParams(_ context.Context) (params dns.RecordListParams, diags diag.Diagnostics) {
	params = dns.RecordListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
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
	if !m.Filter.Content.IsNull() {
		params.Content = cloudflare.F(m.Filter.Content.ValueString())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(shared.SortDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Match.IsNull() {
		params.Match = cloudflare.F(dns.RecordListParamsMatch(m.Filter.Match.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
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

type DNSRecordFindOneByDataSourceModel struct {
	ZoneID    types.String                     `tfsdk:"zone_id" path:"zone_id,required"`
	Comment   *DNSRecordCommentDataSourceModel `tfsdk:"comment" query:"comment,optional"`
	Content   types.String                     `tfsdk:"content" query:"content,optional"`
	Direction types.String                     `tfsdk:"direction" query:"direction,computed_optional"`
	Match     types.String                     `tfsdk:"match" query:"match,computed_optional"`
	Name      types.String                     `tfsdk:"name" query:"name,optional"`
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

type DNSRecordTagDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent,optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,optional"`
	Present    types.String `tfsdk:"present" json:"present,optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}
