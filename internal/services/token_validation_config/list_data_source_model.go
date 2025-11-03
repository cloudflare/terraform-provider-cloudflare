// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_config

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/token_validation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TokenValidationConfigsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[TokenValidationConfigsResultDataSourceModel] `json:"result,computed"`
}

type TokenValidationConfigsDataSourceModel struct {
	ZoneID   types.String                                                              `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                               `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[TokenValidationConfigsResultDataSourceModel] `tfsdk:"result"`
}

func (m *TokenValidationConfigsDataSourceModel) toListParams(_ context.Context) (params token_validation.ConfigurationListParams, diags diag.Diagnostics) {
	params = token_validation.ConfigurationListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type TokenValidationConfigsResultDataSourceModel struct {
	ID           types.String                                                               `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Credentials  customfield.NestedObject[TokenValidationConfigsCredentialsDataSourceModel] `tfsdk:"credentials" json:"credentials,computed"`
	Description  types.String                                                               `tfsdk:"description" json:"description,computed"`
	LastUpdated  timetypes.RFC3339                                                          `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Title        types.String                                                               `tfsdk:"title" json:"title,computed"`
	TokenSources customfield.List[types.String]                                             `tfsdk:"token_sources" json:"token_sources,computed"`
	TokenType    types.String                                                               `tfsdk:"token_type" json:"token_type,computed"`
}

type TokenValidationConfigsCredentialsDataSourceModel struct {
	Keys customfield.NestedObjectList[TokenValidationConfigsCredentialsKeysDataSourceModel] `tfsdk:"keys" json:"keys,computed"`
}

type TokenValidationConfigsCredentialsKeysDataSourceModel struct {
	Alg types.String `tfsdk:"alg" json:"alg,computed"`
	E   types.String `tfsdk:"e" json:"e,computed"`
	Kid types.String `tfsdk:"kid" json:"kid,computed"`
	Kty types.String `tfsdk:"kty" json:"kty,computed"`
	N   types.String `tfsdk:"n" json:"n,computed"`
	Crv types.String `tfsdk:"crv" json:"crv,computed"`
	X   types.String `tfsdk:"x" json:"x,computed"`
	Y   types.String `tfsdk:"y" json:"y,computed"`
}
