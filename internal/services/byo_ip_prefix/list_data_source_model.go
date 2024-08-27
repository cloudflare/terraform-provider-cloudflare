// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ByoIPPrefixesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ByoIPPrefixesResultDataSourceModel] `json:"result,computed"`
}

type ByoIPPrefixesDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ByoIPPrefixesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ByoIPPrefixesDataSourceModel) toListParams() (params addressing.PrefixListParams, diags diag.Diagnostics) {
	params = addressing.PrefixListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ByoIPPrefixesResultDataSourceModel struct {
	ID                   types.String      `tfsdk:"id" json:"id,computed_optional"`
	AccountID            types.String      `tfsdk:"account_id" json:"account_id,computed_optional"`
	Advertised           types.Bool        `tfsdk:"advertised" json:"advertised,computed_optional"`
	AdvertisedModifiedAt timetypes.RFC3339 `tfsdk:"advertised_modified_at" json:"advertised_modified_at,computed" format:"date-time"`
	Approved             types.String      `tfsdk:"approved" json:"approved,computed_optional"`
	ASN                  types.Int64       `tfsdk:"asn" json:"asn,computed_optional"`
	CIDR                 types.String      `tfsdk:"cidr" json:"cidr,computed_optional"`
	CreatedAt            timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description          types.String      `tfsdk:"description" json:"description,computed_optional"`
	LOADocumentID        types.String      `tfsdk:"loa_document_id" json:"loa_document_id,computed_optional"`
	ModifiedAt           timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	OnDemandEnabled      types.Bool        `tfsdk:"on_demand_enabled" json:"on_demand_enabled,computed_optional"`
	OnDemandLocked       types.Bool        `tfsdk:"on_demand_locked" json:"on_demand_locked,computed_optional"`
}
