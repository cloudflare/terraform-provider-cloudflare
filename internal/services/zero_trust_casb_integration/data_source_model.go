// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_integration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbIntegrationResultDataSourceEnvelope struct {
	Result ZeroTrustCasbIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustCasbIntegrationDataSourceModel struct {
	ID                types.String                                                                       `tfsdk:"id" path:"id,computed_optional"`
	AccountID         types.String                                                                       `tfsdk:"account_id" path:"account_id,required"`
	Created           timetypes.RFC3339                                                                  `tfsdk:"created" json:"created,computed" format:"date-time"`
	CredentialsExpiry timetypes.RFC3339                                                                  `tfsdk:"credentials_expiry" json:"credentials_expiry,computed" format:"date-time"`
	IsPaused          types.Bool                                                                         `tfsdk:"is_paused" json:"is_paused,computed"`
	LastHydrated      timetypes.RFC3339                                                                  `tfsdk:"last_hydrated" json:"last_hydrated,computed" format:"date-time"`
	Name              types.String                                                                       `tfsdk:"name" json:"name,computed"`
	Status            types.String                                                                       `tfsdk:"status" json:"status,computed"`
	Updated           timetypes.RFC3339                                                                  `tfsdk:"updated" json:"updated,computed" format:"date-time"`
	Application       customfield.Map[types.String]                                                      `tfsdk:"application" json:"application,computed"`
	AuthMethod        customfield.Map[types.String]                                                      `tfsdk:"auth_method" json:"auth_method,computed"`
	DLPProfiles       customfield.List[types.String]                                                     `tfsdk:"dlp_profiles" json:"dlp_profiles,computed"`
	HealthDetails     customfield.List[customfield.Map[jsontypes.Normalized]]                            `tfsdk:"health_details" json:"health_details,computed"`
	UseCases          customfield.List[customfield.Map[jsontypes.Normalized]]                            `tfsdk:"use_cases" json:"use_cases,computed"`
	AuthorizationLink customfield.NestedObject[ZeroTrustCasbIntegrationAuthorizationLinkDataSourceModel] `tfsdk:"authorization_link" json:"authorization_link,computed"`
	Filter            *ZeroTrustCasbIntegrationFindOneByDataSourceModel                                  `tfsdk:"filter"`
}

func (m *ZeroTrustCasbIntegrationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CasbIntegrationGetParams, diags diag.Diagnostics) {
	params = zero_trust.CasbIntegrationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCasbIntegrationDataSourceModel) toListParams(_ context.Context) (params zero_trust.CasbIntegrationListParams, diags diag.Diagnostics) {
	params = zero_trust.CasbIntegrationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Application.IsNull() {
		params.Application = cloudflare.F(m.Filter.Application.ValueString())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CasbIntegrationListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.DLPEnabled.IsNull() {
		params.DLPEnabled = cloudflare.F(m.Filter.DLPEnabled.ValueBool())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CasbIntegrationListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Page.IsNull() {
		params.Page = cloudflare.F(m.Filter.Page.ValueInt64())
	}
	if !m.Filter.PageSize.IsNull() {
		params.PageSize = cloudflare.F(m.Filter.PageSize.ValueInt64())
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(zero_trust.CasbIntegrationListParamsStatus(m.Filter.Status.ValueString()))
	}
	if !m.Filter.UseCases.IsNull() {
		params.UseCases = cloudflare.F(m.Filter.UseCases.ValueString())
	}

	return
}

type ZeroTrustCasbIntegrationAuthorizationLinkDataSourceModel struct {
	Components customfield.Map[jsontypes.Normalized] `tfsdk:"components" json:"components,computed"`
	Link       types.String                          `tfsdk:"link" json:"link,computed"`
}

type ZeroTrustCasbIntegrationFindOneByDataSourceModel struct {
	Application types.String `tfsdk:"application" query:"application,optional"`
	Direction   types.String `tfsdk:"direction" query:"direction,optional"`
	DLPEnabled  types.Bool   `tfsdk:"dlp_enabled" query:"dlp_enabled,optional"`
	Order       types.String `tfsdk:"order" query:"order,optional"`
	Page        types.Int64  `tfsdk:"page" query:"page,optional"`
	PageSize    types.Int64  `tfsdk:"page_size" query:"page_size,optional"`
	Search      types.String `tfsdk:"search" query:"search,optional"`
	Status      types.String `tfsdk:"status" query:"status,optional"`
	UseCases    types.String `tfsdk:"use_cases" query:"use_cases,optional"`
}
