// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_finding_type

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbFindingTypesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCasbFindingTypesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCasbFindingTypesDataSourceModel struct {
	AccountID    types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	Direction    types.String                                                                 `tfsdk:"direction" query:"direction,optional"`
	MaxUpdatedAt timetypes.RFC3339                                                            `tfsdk:"max_updated_at" query:"max_updated_at,optional" format:"date-time"`
	MinUpdatedAt timetypes.RFC3339                                                            `tfsdk:"min_updated_at" query:"min_updated_at,optional" format:"date-time"`
	Order        types.String                                                                 `tfsdk:"order" query:"order,optional"`
	Quarantined  types.Bool                                                                   `tfsdk:"quarantined" query:"quarantined,optional"`
	Search       types.String                                                                 `tfsdk:"search" query:"search,optional"`
	Origins      *[]types.String                                                              `tfsdk:"origins" query:"origins,optional"`
	Severities   *[]types.String                                                              `tfsdk:"severities" query:"severities,optional"`
	Types        *[]types.String                                                              `tfsdk:"types" query:"types,optional"`
	Vendors      *[]types.String                                                              `tfsdk:"vendors" query:"vendors,optional"`
	MaxItems     types.Int64                                                                  `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[ZeroTrustCasbFindingTypesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCasbFindingTypesDataSourceModel) toListParams(_ context.Context) (params zero_trust.CasbPostureFindingTypeListParams, diags diag.Diagnostics) {
	mOrigins := []zero_trust.CasbPostureFindingTypeListParamsOrigin{}
	if m.Origins != nil {
		for _, item := range *m.Origins {
			mOrigins = append(mOrigins, zero_trust.CasbPostureFindingTypeListParamsOrigin(item.ValueString()))
		}
	}
	mSeverities := []zero_trust.CasbPostureFindingTypeListParamsSeverity{}
	if m.Severities != nil {
		for _, item := range *m.Severities {
			mSeverities = append(mSeverities, zero_trust.CasbPostureFindingTypeListParamsSeverity(item.ValueString()))
		}
	}
	mTypes := []zero_trust.CasbPostureFindingTypeListParamsType{}
	if m.Types != nil {
		for _, item := range *m.Types {
			mTypes = append(mTypes, zero_trust.CasbPostureFindingTypeListParamsType(item.ValueString()))
		}
	}
	mVendors := []zero_trust.CasbPostureFindingTypeListParamsVendor{}
	if m.Vendors != nil {
		for _, item := range *m.Vendors {
			mVendors = append(mVendors, zero_trust.CasbPostureFindingTypeListParamsVendor(item.ValueString()))
		}
	}
	mMaxUpdatedAt, errs := m.MaxUpdatedAt.ValueRFC3339Time()
	diags.Append(errs...)
	mMinUpdatedAt, errs := m.MinUpdatedAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.CasbPostureFindingTypeListParams{
		AccountID:  cloudflare.F(m.AccountID.ValueString()),
		Origins:    cloudflare.F(mOrigins),
		Severities: cloudflare.F(mSeverities),
		Types:      cloudflare.F(mTypes),
		Vendors:    cloudflare.F(mVendors),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CasbPostureFindingTypeListParamsDirection(m.Direction.ValueString()))
	}
	if !m.MaxUpdatedAt.IsNull() {
		params.MaxUpdatedAt = cloudflare.F(mMaxUpdatedAt)
	}
	if !m.MinUpdatedAt.IsNull() {
		params.MinUpdatedAt = cloudflare.F(mMinUpdatedAt)
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CasbPostureFindingTypeListParamsOrder(m.Order.ValueString()))
	}
	if !m.Quarantined.IsNull() {
		params.Quarantined = cloudflare.F(m.Quarantined.ValueBool())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type ZeroTrustCasbFindingTypesResultDataSourceModel struct {
	ID                       types.String                                                               `tfsdk:"id" json:"id,computed"`
	Category                 customfield.NestedObject[ZeroTrustCasbFindingTypesCategoryDataSourceModel] `tfsdk:"category" json:"category,computed"`
	Name                     types.String                                                               `tfsdk:"name" json:"name,computed"`
	Severity                 types.String                                                               `tfsdk:"severity" json:"severity,computed"`
	Vendor                   types.String                                                               `tfsdk:"vendor" json:"vendor,computed"`
	AppliesToAllIntegrations types.Bool                                                                 `tfsdk:"applies_to_all_integrations" json:"applies_to_all_integrations,computed"`
	AssetClass               types.String                                                               `tfsdk:"asset_class" json:"asset_class,computed"`
	AssetVendor              types.String                                                               `tfsdk:"asset_vendor" json:"asset_vendor,computed"`
	CreatedAt                timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description              types.String                                                               `tfsdk:"description" json:"description,computed"`
	Expression               types.String                                                               `tfsdk:"expression" json:"expression,computed"`
	IntegrationIDs           customfield.List[types.String]                                             `tfsdk:"integration_ids" json:"integration_ids,computed"`
	MatchCount               types.Int64                                                                `tfsdk:"match_count" json:"match_count,computed"`
	Origin                   types.String                                                               `tfsdk:"origin" json:"origin,computed"`
	Quarantined              types.Bool                                                                 `tfsdk:"quarantined" json:"quarantined,computed"`
	UpdatedAt                timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version                  types.Int64                                                                `tfsdk:"version" json:"version,computed"`
}

type ZeroTrustCasbFindingTypesCategoryDataSourceModel struct {
	Observation types.String `tfsdk:"observation" json:"observation,computed"`
	Product     types.String `tfsdk:"product" json:"product,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
}
