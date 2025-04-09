// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ByoIPPrefixResultEnvelope struct {
	Result ByoIPPrefixModel `json:"result"`
}

type ByoIPPrefixModel struct {
	ID                   types.String      `tfsdk:"id" json:"id,computed"`
	AccountID            types.String      `tfsdk:"account_id" path:"account_id,required"`
	ASN                  types.Int64       `tfsdk:"asn" json:"asn,required"`
	CIDR                 types.String      `tfsdk:"cidr" json:"cidr,required"`
	LOADocumentID        types.String      `tfsdk:"loa_document_id" json:"loa_document_id,required"`
	Description          types.String      `tfsdk:"description" json:"description,optional"`
	Advertised           types.Bool        `tfsdk:"advertised" json:"advertised,computed"`
	AdvertisedModifiedAt timetypes.RFC3339 `tfsdk:"advertised_modified_at" json:"advertised_modified_at,computed" format:"date-time"`
	Approved             types.String      `tfsdk:"approved" json:"approved,computed"`
	CreatedAt            timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt           timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	OnDemandEnabled      types.Bool        `tfsdk:"on_demand_enabled" json:"on_demand_enabled,computed"`
	OnDemandLocked       types.Bool        `tfsdk:"on_demand_locked" json:"on_demand_locked,computed"`
}

func (m ByoIPPrefixModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ByoIPPrefixModel) MarshalJSONForUpdate(state ByoIPPrefixModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
