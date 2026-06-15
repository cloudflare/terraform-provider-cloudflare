// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_config

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TokenValidationConfigResultEnvelope struct {
	Result TokenValidationConfigModel `json:"result"`
}

type TokenValidationConfigModel struct {
	ID           types.String                           `tfsdk:"id" json:"id,computed"`
	ZoneID       types.String                           `tfsdk:"zone_id" path:"zone_id,required"`
	TokenType    types.String                           `tfsdk:"token_type" json:"token_type,required"`
	Credentials  *TokenValidationConfigCredentialsModel `tfsdk:"credentials" json:"credentials,required"`
	Description  types.String                           `tfsdk:"description" json:"description,required"`
	Title        types.String                           `tfsdk:"title" json:"title,required"`
	TokenSources *[]types.String                        `tfsdk:"token_sources" json:"token_sources,required"`
	CreatedAt    timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastUpdated  timetypes.RFC3339                      `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
}

func (m TokenValidationConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TokenValidationConfigModel) MarshalJSONForUpdate(state TokenValidationConfigModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type TokenValidationConfigCredentialsModel struct {
	Keys *[]*TokenValidationConfigCredentialsKeysModel `tfsdk:"keys" json:"keys,required"`
}

type TokenValidationConfigCredentialsKeysModel struct {
	Alg types.String `tfsdk:"alg" json:"alg,required"`
	E   types.String `tfsdk:"e" json:"e,optional"`
	Kid types.String `tfsdk:"kid" json:"kid,required"`
	Kty types.String `tfsdk:"kty" json:"kty,required"`
	N   types.String `tfsdk:"n" json:"n,optional"`
	Crv types.String `tfsdk:"crv" json:"crv,optional"`
	X   types.String `tfsdk:"x" json:"x,optional"`
	Y   types.String `tfsdk:"y" json:"y,optional"`
}
