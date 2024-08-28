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

type ByoIPPrefixResultDataSourceEnvelope struct {
	Result ByoIPPrefixDataSourceModel `json:"result,computed"`
}

type ByoIPPrefixResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ByoIPPrefixDataSourceModel] `json:"result,computed"`
}

type ByoIPPrefixDataSourceModel struct {
	PrefixID             types.String                         `tfsdk:"prefix_id" path:"prefix_id"`
	AccountID            types.String                         `tfsdk:"account_id" path:"account_id,computed_optional"`
	AdvertisedModifiedAt timetypes.RFC3339                    `tfsdk:"advertised_modified_at" json:"advertised_modified_at,computed" format:"date-time"`
	CreatedAt            timetypes.RFC3339                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt           timetypes.RFC3339                    `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Advertised           types.Bool                           `tfsdk:"advertised" json:"advertised,computed_optional"`
	Approved             types.String                         `tfsdk:"approved" json:"approved,computed_optional"`
	ASN                  types.Int64                          `tfsdk:"asn" json:"asn,computed_optional"`
	CIDR                 types.String                         `tfsdk:"cidr" json:"cidr,computed_optional"`
	Description          types.String                         `tfsdk:"description" json:"description,computed_optional"`
	ID                   types.String                         `tfsdk:"id" json:"id,computed_optional"`
	LOADocumentID        types.String                         `tfsdk:"loa_document_id" json:"loa_document_id,computed_optional"`
	OnDemandEnabled      types.Bool                           `tfsdk:"on_demand_enabled" json:"on_demand_enabled,computed_optional"`
	OnDemandLocked       types.Bool                           `tfsdk:"on_demand_locked" json:"on_demand_locked,computed_optional"`
	Filter               *ByoIPPrefixFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ByoIPPrefixDataSourceModel) toReadParams() (params addressing.PrefixGetParams, diags diag.Diagnostics) {
	params = addressing.PrefixGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ByoIPPrefixDataSourceModel) toListParams() (params addressing.PrefixListParams, diags diag.Diagnostics) {
	params = addressing.PrefixListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ByoIPPrefixFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}