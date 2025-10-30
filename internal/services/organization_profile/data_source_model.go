// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_profile

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationProfileResultDataSourceEnvelope struct {
	Result OrganizationProfileDataSourceModel `json:"result,computed"`
}

type OrganizationProfileDataSourceModel struct {
	OrganizationID   types.String `tfsdk:"organization_id" path:"organization_id,required"`
	BusinessAddress  types.String `tfsdk:"business_address" json:"business_address,computed"`
	BusinessEmail    types.String `tfsdk:"business_email" json:"business_email,computed"`
	BusinessName     types.String `tfsdk:"business_name" json:"business_name,computed"`
	BusinessPhone    types.String `tfsdk:"business_phone" json:"business_phone,computed"`
	ExternalMetadata types.String `tfsdk:"external_metadata" json:"external_metadata,computed"`
}
