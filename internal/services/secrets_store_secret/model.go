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
	Name      types.String      `tfsdk:"name" json:"name,required"`
	SecretText types.String     `tfsdk:"secret_text" json:"value,required"`
	Scopes    []types.String    `tfsdk:"scopes" json:"scopes,optional"`
	Comment   types.String      `tfsdk:"comment" json:"comment,optional"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	ModifiedAt timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed"`
}

func (m SecretsStoreSecretModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SecretsStoreSecretModel) MarshalJSONForUpdate(state SecretsStoreSecretModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
