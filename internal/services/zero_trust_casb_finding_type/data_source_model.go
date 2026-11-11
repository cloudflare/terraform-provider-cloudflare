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

type ZeroTrustCasbFindingTypeResultDataSourceEnvelope struct {
	Result ZeroTrustCasbFindingTypeDataSourceModel `json:"result,computed"`
}

type ZeroTrustCasbFindingTypeDataSourceModel struct {
	ID                       types.String                                                              `tfsdk:"id" path:"finding_type_id,computed"`
	FindingTypeID            types.String                                                              `tfsdk:"finding_type_id" path:"finding_type_id,optional"`
	AccountID                types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	AppliesToAllIntegrations types.Bool                                                                `tfsdk:"applies_to_all_integrations" json:"applies_to_all_integrations,computed"`
	AssetClass               types.String                                                              `tfsdk:"asset_class" json:"asset_class,computed"`
	AssetVendor              types.String                                                              `tfsdk:"asset_vendor" json:"asset_vendor,computed"`
	CreatedAt                timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description              types.String                                                              `tfsdk:"description" json:"description,computed"`
	Expression               types.String                                                              `tfsdk:"expression" json:"expression,computed"`
	MatchCount               types.Int64                                                               `tfsdk:"match_count" json:"match_count,computed"`
	Name                     types.String                                                              `tfsdk:"name" json:"name,computed"`
	Origin                   types.String                                                              `tfsdk:"origin" json:"origin,computed"`
	Quarantined              types.Bool                                                                `tfsdk:"quarantined" json:"quarantined,computed"`
	Severity                 types.String                                                              `tfsdk:"severity" json:"severity,computed"`
	UpdatedAt                timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Vendor                   types.String                                                              `tfsdk:"vendor" json:"vendor,computed"`
	Version                  types.Int64                                                               `tfsdk:"version" json:"version,computed"`
	IntegrationIDs           customfield.List[types.String]                                            `tfsdk:"integration_ids" json:"integration_ids,computed"`
	Category                 customfield.NestedObject[ZeroTrustCasbFindingTypeCategoryDataSourceModel] `tfsdk:"category" json:"category,computed"`
	Filter                   *ZeroTrustCasbFindingTypeFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *ZeroTrustCasbFindingTypeDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CasbPostureFindingTypeGetParams, diags diag.Diagnostics) {
	params = zero_trust.CasbPostureFindingTypeGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCasbFindingTypeDataSourceModel) toListParams(_ context.Context) (params zero_trust.CasbPostureFindingTypeListParams, diags diag.Diagnostics) {
	mFilterOrigins := []zero_trust.CasbPostureFindingTypeListParamsOrigin{}
	if m.Filter.Origins != nil {
		for _, item := range *m.Filter.Origins {
			mFilterOrigins = append(mFilterOrigins, zero_trust.CasbPostureFindingTypeListParamsOrigin(item.ValueString()))
		}
	}
	mFilterSeverities := []zero_trust.CasbPostureFindingTypeListParamsSeverity{}
	if m.Filter.Severities != nil {
		for _, item := range *m.Filter.Severities {
			mFilterSeverities = append(mFilterSeverities, zero_trust.CasbPostureFindingTypeListParamsSeverity(item.ValueString()))
		}
	}
	mFilterTypes := []zero_trust.CasbPostureFindingTypeListParamsType{}
	if m.Filter.Types != nil {
		for _, item := range *m.Filter.Types {
			mFilterTypes = append(mFilterTypes, zero_trust.CasbPostureFindingTypeListParamsType(item.ValueString()))
		}
	}
	mFilterVendors := []zero_trust.CasbPostureFindingTypeListParamsVendor{}
	if m.Filter.Vendors != nil {
		for _, item := range *m.Filter.Vendors {
			mFilterVendors = append(mFilterVendors, zero_trust.CasbPostureFindingTypeListParamsVendor(item.ValueString()))
		}
	}
	mFilterMaxUpdatedAt, errs := m.Filter.MaxUpdatedAt.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterMinUpdatedAt, errs := m.Filter.MinUpdatedAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.CasbPostureFindingTypeListParams{
		AccountID:  cloudflare.F(m.AccountID.ValueString()),
		Origins:    cloudflare.F(mFilterOrigins),
		Severities: cloudflare.F(mFilterSeverities),
		Types:      cloudflare.F(mFilterTypes),
		Vendors:    cloudflare.F(mFilterVendors),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CasbPostureFindingTypeListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.MaxUpdatedAt.IsNull() {
		params.MaxUpdatedAt = cloudflare.F(mFilterMaxUpdatedAt)
	}
	if !m.Filter.MinUpdatedAt.IsNull() {
		params.MinUpdatedAt = cloudflare.F(mFilterMinUpdatedAt)
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CasbPostureFindingTypeListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Quarantined.IsNull() {
		params.Quarantined = cloudflare.F(m.Filter.Quarantined.ValueBool())
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type ZeroTrustCasbFindingTypeCategoryDataSourceModel struct {
	Observation types.String `tfsdk:"observation" json:"observation,computed"`
	Product     types.String `tfsdk:"product" json:"product,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
}

type ZeroTrustCasbFindingTypeFindOneByDataSourceModel struct {
	Direction    types.String      `tfsdk:"direction" query:"direction,optional"`
	MaxUpdatedAt timetypes.RFC3339 `tfsdk:"max_updated_at" query:"max_updated_at,optional" format:"date-time"`
	MinUpdatedAt timetypes.RFC3339 `tfsdk:"min_updated_at" query:"min_updated_at,optional" format:"date-time"`
	Order        types.String      `tfsdk:"order" query:"order,optional"`
	Origins      *[]types.String   `tfsdk:"origins" query:"origins,optional"`
	Quarantined  types.Bool        `tfsdk:"quarantined" query:"quarantined,optional"`
	Search       types.String      `tfsdk:"search" query:"search,optional"`
	Severities   *[]types.String   `tfsdk:"severities" query:"severities,optional"`
	Types        *[]types.String   `tfsdk:"types" query:"types,optional"`
	Vendors      *[]types.String   `tfsdk:"vendors" query:"vendors,optional"`
}
