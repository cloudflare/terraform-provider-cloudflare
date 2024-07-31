// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fallback_domain

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FallbackDomainResultEnvelope struct {
	Result FallbackDomainModel `json:"result,computed"`
}

type FallbackDomainModel struct {
	ID          types.String            `tfsdk:"id" json:"-,computed"`
	AccountID   types.String            `tfsdk:"account_id" path:"account_id"`
	PolicyID    types.String            `tfsdk:"policy_id" path:"policy_id"`
	Suffix      types.String            `tfsdk:"suffix" json:"suffix"`
	Description types.String            `tfsdk:"description" json:"description"`
	DNSServer   *[]jsontypes.Normalized `tfsdk:"dns_server" json:"dns_server"`
}
