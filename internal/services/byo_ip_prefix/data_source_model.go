// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/addressing"
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
	ID                   types.String      `tfsdk:"id" json:"-,computed"`
	PrefixID             types.String      `tfsdk:"prefix_id" path:"prefix_id,optional"`
	AccountID            types.String      `tfsdk:"account_id" path:"account_id,computed"`
	Advertised           types.Bool        `tfsdk:"advertised" json:"advertised,computed"`
	AdvertisedModifiedAt timetypes.RFC3339 `tfsdk:"advertised_modified_at" json:"advertised_modified_at,computed" format:"date-time"`
	Approved             types.String      `tfsdk:"approved" json:"approved,computed"`
	ASN                  types.Int64       `tfsdk:"asn" json:"asn,computed"`
	CIDR                 types.String      `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt            timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description          types.String      `tfsdk:"description" json:"description,computed"`
	LOADocumentID        types.String      `tfsdk:"loa_document_id" json:"loa_document_id,computed"`
	ModifiedAt           timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	OnDemandEnabled      types.Bool        `tfsdk:"on_demand_enabled" json:"on_demand_enabled,computed"`
	OnDemandLocked       types.Bool        `tfsdk:"on_demand_locked" json:"on_demand_locked,computed"`
}

func (m *ByoIPPrefixDataSourceModel) toReadParams(_ context.Context) (params addressing.PrefixGetParams, diags diag.Diagnostics) {
	params = addressing.PrefixGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ByoIPPrefixDataSourceModel) toListParams(_ context.Context) (params addressing.PrefixListParams, diags diag.Diagnostics) {
	params = addressing.PrefixListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
