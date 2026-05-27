// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/secrets_store"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecretsStoreResultDataSourceEnvelope struct {
	Result SecretsStoreDataSourceModel `json:"result,computed"`
}

type SecretsStoreDataSourceModel struct {
	ID        types.String                          `tfsdk:"id" path:"store_id,computed"`
	StoreID   types.String                          `tfsdk:"store_id" path:"store_id,optional"`
	AccountID types.String                          `tfsdk:"account_id" path:"account_id,required"`
	Created   timetypes.RFC3339                     `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                     `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                          `tfsdk:"name" json:"name,computed"`
	Filter    *SecretsStoreFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SecretsStoreDataSourceModel) toReadParams(_ context.Context) (params secrets_store.StoreGetParams, diags diag.Diagnostics) {
	params = secrets_store.StoreGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *SecretsStoreDataSourceModel) toListParams(_ context.Context) (params secrets_store.StoreListParams, diags diag.Diagnostics) {
	params = secrets_store.StoreListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(secrets_store.StoreListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(secrets_store.StoreListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type SecretsStoreFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String `tfsdk:"order" query:"order,computed_optional"`
}
