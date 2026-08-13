// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_secret

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecretsStoreSecretResultDataSourceEnvelope struct {
	Result SecretsStoreSecretDataSourceModel `json:"result,computed"`
}

type SecretsStoreSecretDataSourceModel struct {
	ID        types.String                                `tfsdk:"id" path:"secret_id,computed"`
	SecretID  types.String                                `tfsdk:"secret_id" path:"secret_id,optional"`
	AccountID types.String                                `tfsdk:"account_id" path:"account_id,required"`
	StoreID   types.String                                `tfsdk:"store_id" path:"store_id,required"`
	Comment   types.String                                `tfsdk:"comment" json:"comment,computed"`
	Created   timetypes.RFC3339                           `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                           `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                                `tfsdk:"name" json:"name,computed"`
	Status    types.String                                `tfsdk:"status" json:"status,computed"`
	Scopes    customfield.List[types.String]              `tfsdk:"scopes" json:"scopes,computed"`
	Filter    *SecretsStoreSecretFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SecretsStoreSecretDataSourceModel) toReadParams(_ context.Context) (params secrets_store.StoreSecretGetParams, diags diag.Diagnostics) {
	params = secrets_store.StoreSecretGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *SecretsStoreSecretDataSourceModel) toListParams(_ context.Context) (params secrets_store.StoreSecretListParams, diags diag.Diagnostics) {
	mFilterScopes := [][]string{}
	if m.Filter.Scopes != nil {
		for _, item := range *m.Filter.Scopes {
			mFilterItem := []string{}
			if item != nil {
				for _, item := range *item {
					mFilterItem = append(mFilterItem, item.ValueString())
				}
			}
			mFilterScopes = append(mFilterScopes, mFilterItem)
		}
	}

	params = secrets_store.StoreSecretListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Scopes:    cloudflare.F(mFilterScopes),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(secrets_store.StoreSecretListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(secrets_store.StoreSecretListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type SecretsStoreSecretFindOneByDataSourceModel struct {
	Direction types.String       `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String       `tfsdk:"order" query:"order,computed_optional"`
	Scopes    *[]*[]types.String `tfsdk:"scopes" query:"scopes,optional"`
	Search    types.String       `tfsdk:"search" query:"search,optional"`
}
