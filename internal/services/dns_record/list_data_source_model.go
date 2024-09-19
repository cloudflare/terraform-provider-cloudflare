// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/dns"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSRecordsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSRecordsResultDataSourceModel] `json:"result,computed"`
}

type DNSRecordsDataSourceModel struct {
	ZoneID    types.String                                                  `tfsdk:"zone_id" path:"zone_id,required"`
	Content   types.String                                                  `tfsdk:"content" query:"content,optional"`
	Name      types.String                                                  `tfsdk:"name" query:"name,optional"`
	Search    types.String                                                  `tfsdk:"search" query:"search,optional"`
	Type      types.String                                                  `tfsdk:"type" query:"type,optional"`
	Comment   *DNSRecordsCommentDataSourceModel                             `tfsdk:"comment" query:"comment,optional"`
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
	if !m.Content.IsNull() {
		params.Content = cloudflare.F(m.Content.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(shared.SortDirection(m.Direction.ValueString()))
	}
	if !m.Match.IsNull() {
		params.Match = cloudflare.F(dns.RecordListParamsMatch(m.Match.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
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
	Absent     types.String `tfsdk:"absent" json:"absent,computed_optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,computed_optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,computed_optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,computed_optional"`
	Present    types.String `tfsdk:"present" json:"present,computed_optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,computed_optional"`
}

type DNSRecordsTagDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent,computed_optional"`
	Contains   types.String `tfsdk:"contains" json:"contains,computed_optional"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith,computed_optional"`
	Exact      types.String `tfsdk:"exact" json:"exact,computed_optional"`
	Present    types.String `tfsdk:"present" json:"present,computed_optional"`
	Startswith types.String `tfsdk:"startswith" json:"startswith,computed_optional"`
}

type DNSRecordsResultDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	Comment           types.String                   `tfsdk:"comment" json:"comment,computed"`
	CreatedOn         timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Meta              jsontypes.Normalized           `tfsdk:"meta" json:"meta,computed"`
	ModifiedOn        timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name              types.String                   `tfsdk:"name" json:"name,computed"`
	Proxiable         types.Bool                     `tfsdk:"proxiable" json:"proxiable,computed"`
	Proxied           types.Bool                     `tfsdk:"proxied" json:"proxied,computed"`
	Settings          jsontypes.Normalized           `tfsdk:"settings" json:"settings,computed"`
	Tags              customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	TTL               types.Float64                  `tfsdk:"ttl" json:"ttl,computed"`
	CommentModifiedOn timetypes.RFC3339              `tfsdk:"comment_modified_on" json:"comment_modified_on,computed" format:"date-time"`
	TagsModifiedOn    timetypes.RFC3339              `tfsdk:"tags_modified_on" json:"tags_modified_on,computed" format:"date-time"`
}
