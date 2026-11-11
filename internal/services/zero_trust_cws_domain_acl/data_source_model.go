// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_domain_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSDomainACLResultDataSourceEnvelope struct {
	Result ZeroTrustCWSDomainACLDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSDomainACLDataSourceModel struct {
	ID            types.String                                                             `tfsdk:"id" path:"domain_acl_id,computed"`
	DomainACLID   types.String                                                             `tfsdk:"domain_acl_id" path:"domain_acl_id,optional"`
	AccountID     types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt     timetypes.RFC3339                                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DefaultAction types.String                                                             `tfsdk:"default_action" json:"default_action,computed"`
	Name          types.String                                                             `tfsdk:"name" json:"name,computed"`
	UpdatedAt     timetypes.RFC3339                                                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Entries       customfield.NestedObjectSet[ZeroTrustCWSDomainACLEntriesDataSourceModel] `tfsdk:"entries" json:"entries,computed"`
	Filter        *ZeroTrustCWSDomainACLFindOneByDataSourceModel                           `tfsdk:"filter"`
}

func (m *ZeroTrustCWSDomainACLDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSDomainACLGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSDomainACLGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSDomainACLDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSDomainACLListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSDomainACLListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.DeploymentID.IsNull() {
		params.DeploymentID = cloudflare.F(m.Filter.DeploymentID.ValueString())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSDomainACLListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSDomainACLListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSDomainACLEntriesDataSourceModel struct {
	Action types.String `tfsdk:"action" json:"action,computed"`
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustCWSDomainACLFindOneByDataSourceModel struct {
	DeploymentID types.String `tfsdk:"deployment_id" query:"deployment_id,optional"`
	Direction    types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name         types.String `tfsdk:"name" query:"name,optional"`
	Order        types.String `tfsdk:"order" query:"order,computed_optional"`
}
