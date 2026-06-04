// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_secret

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecretsStoreSecretResultEnvelope struct {
	Result SecretsStoreSecretModel `json:"result"`
}

type SecretsStoreSecretModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	StoreID   types.String      `tfsdk:"store_id" path:"store_id,required"`
	Comment   types.String      `tfsdk:"comment" json:"comment,optional"`
	Value     types.String      `tfsdk:"value" json:"value,optional,no_refresh"`
	Scopes    *[]types.String   `tfsdk:"scopes" json:"scopes,optional"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

func (m SecretsStoreSecretModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SecretsStoreSecretModel) MarshalJSONForUpdate(state SecretsStoreSecretModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
