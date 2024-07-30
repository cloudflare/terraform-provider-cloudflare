// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package record

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RecordsResultListDataSourceEnvelope struct {
	Result *[]*RecordsResultDataSourceModel `json:"result,computed"`
}

type RecordsDataSourceModel struct {
	ZoneID    types.String                     `tfsdk:"zone_id" path:"zone_id"`
	Comment   *RecordsCommentDataSourceModel   `tfsdk:"comment" query:"comment"`
	Content   types.String                     `tfsdk:"content" query:"content"`
	Direction types.String                     `tfsdk:"direction" query:"direction"`
	Match     types.String                     `tfsdk:"match" query:"match"`
	Name      types.String                     `tfsdk:"name" query:"name"`
	Order     types.String                     `tfsdk:"order" query:"order"`
	Page      types.Float64                    `tfsdk:"page" query:"page"`
	PerPage   types.Float64                    `tfsdk:"per_page" query:"per_page"`
	Proxied   types.Bool                       `tfsdk:"proxied" query:"proxied"`
	Search    types.String                     `tfsdk:"search" query:"search"`
	Tag       *RecordsTagDataSourceModel       `tfsdk:"tag" query:"tag"`
	TagMatch  types.String                     `tfsdk:"tag_match" query:"tag_match"`
	Type      types.String                     `tfsdk:"type" query:"type"`
	MaxItems  types.Int64                      `tfsdk:"max_items"`
	Result    *[]*RecordsResultDataSourceModel `tfsdk:"result"`
}

type RecordsCommentDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent"`
	Contains   types.String `tfsdk:"contains" json:"contains"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith"`
	Exact      types.String `tfsdk:"exact" json:"exact"`
	Present    types.String `tfsdk:"present" json:"present"`
	Startswith types.String `tfsdk:"startswith" json:"startswith"`
}

type RecordsTagDataSourceModel struct {
	Absent     types.String `tfsdk:"absent" json:"absent"`
	Contains   types.String `tfsdk:"contains" json:"contains"`
	Endswith   types.String `tfsdk:"endswith" json:"endswith"`
	Exact      types.String `tfsdk:"exact" json:"exact"`
	Present    types.String `tfsdk:"present" json:"present"`
	Startswith types.String `tfsdk:"startswith" json:"startswith"`
}

type RecordsResultDataSourceModel struct {
}
