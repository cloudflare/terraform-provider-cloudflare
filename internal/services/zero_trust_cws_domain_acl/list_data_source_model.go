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

type ZeroTrustCWSDomainACLsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSDomainACLsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSDomainACLsDataSourceModel struct {
	AccountID    types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	DeploymentID types.String                                                              `tfsdk:"deployment_id" query:"deployment_id,optional"`
	Name         types.String                                                              `tfsdk:"name" query:"name,optional"`
	Direction    types.String                                                              `tfsdk:"direction" query:"direction,computed_optional"`
	Order        types.String                                                              `tfsdk:"order" query:"order,computed_optional"`
	MaxItems     types.Int64                                                               `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[ZeroTrustCWSDomainACLsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSDomainACLsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSDomainACLListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSDomainACLListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.DeploymentID.IsNull() {
		params.DeploymentID = cloudflare.F(m.DeploymentID.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSDomainACLListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSDomainACLListParamsOrder(m.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSDomainACLsResultDataSourceModel struct {
	ID            types.String                                                              `tfsdk:"id" json:"id,computed"`
	CreatedAt     timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DefaultAction types.String                                                              `tfsdk:"default_action" json:"default_action,computed"`
	Entries       customfield.NestedObjectSet[ZeroTrustCWSDomainACLsEntriesDataSourceModel] `tfsdk:"entries" json:"entries,computed"`
	Name          types.String                                                              `tfsdk:"name" json:"name,computed"`
	UpdatedAt     timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustCWSDomainACLsEntriesDataSourceModel struct {
	Action types.String `tfsdk:"action" json:"action,computed"`
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}
