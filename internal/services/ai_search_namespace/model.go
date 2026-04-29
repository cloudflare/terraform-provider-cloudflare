// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchNamespaceResultEnvelope struct {
	Result AISearchNamespaceModel `json:"result"`
}

type AISearchNamespaceModel struct {
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	Description types.String      `tfsdk:"description" json:"description,optional"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
}

func (m AISearchNamespaceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AISearchNamespaceModel) MarshalJSONForUpdate(state AISearchNamespaceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
