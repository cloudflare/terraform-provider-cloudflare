// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationProfileResultEnvelope struct {
	Result OrganizationProfileModel `json:"result"`
}

type OrganizationProfileModel struct {
	OrganizationID   types.String `tfsdk:"organization_id" path:"organization_id,required"`
	BusinessAddress  types.String `tfsdk:"business_address" json:"business_address,optional"`
	BusinessEmail    types.String `tfsdk:"business_email" json:"business_email,optional"`
	BusinessName     types.String `tfsdk:"business_name" json:"business_name,optional"`
	BusinessPhone    types.String `tfsdk:"business_phone" json:"business_phone,optional"`
	ExternalMetadata types.String `tfsdk:"external_metadata" json:"external_metadata,optional"`
}

func (m OrganizationProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OrganizationProfileModel) MarshalJSONForUpdate(state OrganizationProfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
