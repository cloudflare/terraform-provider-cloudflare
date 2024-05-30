// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsNamespaceResultEnvelope struct {
	Result WorkersForPlatformsNamespaceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespaceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	Name        types.String `tfsdk:"name" json:"name"`
}
