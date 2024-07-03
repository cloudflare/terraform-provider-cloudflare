// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainResultEnvelope struct {
	Result PagesDomainModel `json:"result,computed"`
}

type PagesDomainModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	ProjectName types.String `tfsdk:"project_name" path:"project_name"`
	DomainName  types.String `tfsdk:"domain_name" path:"domain_name"`
}
