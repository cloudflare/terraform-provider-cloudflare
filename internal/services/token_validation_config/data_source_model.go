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

type TokenValidationConfigResultDataSourceEnvelope struct {
	Result TokenValidationConfigDataSourceModel `json:"result,computed"`
}

type TokenValidationConfigDataSourceModel struct {
	ID           types.String                                                              `tfsdk:"id" path:"config_id,computed"`
	ConfigID     types.String                                                              `tfsdk:"config_id" path:"config_id,optional"`
	ZoneID       types.String                                                              `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedAt    timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description  types.String                                                              `tfsdk:"description" json:"description,computed"`
	LastUpdated  timetypes.RFC3339                                                         `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Title        types.String                                                              `tfsdk:"title" json:"title,computed"`
	TokenType    types.String                                                              `tfsdk:"token_type" json:"token_type,computed"`
	TokenSources customfield.List[types.String]                                            `tfsdk:"token_sources" json:"token_sources,computed"`
	Credentials  customfield.NestedObject[TokenValidationConfigCredentialsDataSourceModel] `tfsdk:"credentials" json:"credentials,computed"`
}

func (m *TokenValidationConfigDataSourceModel) toReadParams(_ context.Context) (params token_validation.ConfigurationGetParams, diags diag.Diagnostics) {
	params = token_validation.ConfigurationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type TokenValidationConfigCredentialsDataSourceModel struct {
	Keys customfield.NestedObjectList[TokenValidationConfigCredentialsKeysDataSourceModel] `tfsdk:"keys" json:"keys,computed"`
}

type TokenValidationConfigCredentialsKeysDataSourceModel struct {
	Alg types.String `tfsdk:"alg" json:"alg,computed"`
	Kid types.String `tfsdk:"kid" json:"kid,computed"`
	Kty types.String `tfsdk:"kty" json:"kty,computed"`
	X   types.String `tfsdk:"x" json:"x,computed"`
	Y   types.String `tfsdk:"y" json:"y,computed"`
	E   types.String `tfsdk:"e" json:"e,computed"`
	N   types.String `tfsdk:"n" json:"n,computed"`
}
