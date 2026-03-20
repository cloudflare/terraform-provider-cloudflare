package secrets_store

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecretsStoreResultEnvelope struct {
	Result SecretsStoreModel `json:"result"`
}

type SecretsStoreModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String      `tfsdk:"name" json:"name,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	ModifiedAt timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed"`
}

func (m SecretsStoreModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SecretsStoreModel) MarshalJSONForUpdate(state SecretsStoreModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
