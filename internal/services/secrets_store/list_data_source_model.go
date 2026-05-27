// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecretsStoresResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecretsStoresResultDataSourceModel] `json:"result,computed"`
}

type SecretsStoresDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String                                                     `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String                                                     `tfsdk:"order" query:"order,computed_optional"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[SecretsStoresResultDataSourceModel] `tfsdk:"result"`
}

func (m *SecretsStoresDataSourceModel) toListParams(_ context.Context) (params secrets_store.StoreListParams, diags diag.Diagnostics) {
	params = secrets_store.StoreListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(secrets_store.StoreListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(secrets_store.StoreListParamsOrder(m.Order.ValueString()))
	}

	return
}

type SecretsStoresResultDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	AccountID types.String      `tfsdk:"account_id" json:"account_id,computed"`
}
