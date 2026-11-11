// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_domain_acl

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSDomainACLResultEnvelope struct {
	Result ZeroTrustCWSDomainACLModel `json:"result"`
}

type ZeroTrustCWSDomainACLModel struct {
	ID            types.String                          `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                          `tfsdk:"account_id" path:"account_id,required"`
	DefaultAction types.String                          `tfsdk:"default_action" json:"default_action,required"`
	Name          types.String                          `tfsdk:"name" json:"name,required"`
	Entries       *[]*ZeroTrustCWSDomainACLEntriesModel `tfsdk:"entries" json:"entries,required"`
	CreatedAt     timetypes.RFC3339                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt     timetypes.RFC3339                     `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSDomainACLModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSDomainACLModel) MarshalJSONForUpdate(state ZeroTrustCWSDomainACLModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustCWSDomainACLEntriesModel struct {
	Action types.String `tfsdk:"action" json:"action,required"`
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}
