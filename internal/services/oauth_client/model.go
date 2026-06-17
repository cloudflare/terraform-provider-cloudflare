// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_client

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OAuthClientResultEnvelope struct {
	Result OAuthClientModel `json:"result"`
}

type OAuthClientModel struct {
	AccountID               types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	OAuthClientID           types.String                                                    `tfsdk:"oauth_client_id" path:"oauth_client_id,optional"`
	ClientName              types.String                                                    `tfsdk:"client_name" json:"client_name,required"`
	TokenEndpointAuthMethod types.String                                                    `tfsdk:"token_endpoint_auth_method" json:"token_endpoint_auth_method,required"`
	GrantTypes              *[]types.String                                                 `tfsdk:"grant_types" json:"grant_types,required"`
	RedirectURIs            *[]types.String                                                 `tfsdk:"redirect_uris" json:"redirect_uris,required"`
	ResponseTypes           *[]types.String                                                 `tfsdk:"response_types" json:"response_types,required"`
	Scopes                  *[]types.String                                                 `tfsdk:"scopes" json:"scopes,required"`
	ClientURI               types.String                                                    `tfsdk:"client_uri" json:"client_uri,optional"`
	LogoURI                 types.String                                                    `tfsdk:"logo_uri" json:"logo_uri,optional"`
	PolicyURI               types.String                                                    `tfsdk:"policy_uri" json:"policy_uri,optional"`
	TosURI                  types.String                                                    `tfsdk:"tos_uri" json:"tos_uri,optional"`
	Visibility              types.String                                                    `tfsdk:"visibility" json:"visibility,optional"`
	AllowedCORSOrigins      *[]types.String                                                 `tfsdk:"allowed_cors_origins" json:"allowed_cors_origins,optional"`
	PostLogoutRedirectURIs  *[]types.String                                                 `tfsdk:"post_logout_redirect_uris" json:"post_logout_redirect_uris,optional"`
	ClientID                types.String                                                    `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret            types.String                                                    `tfsdk:"client_secret" json:"client_secret,computed,no_refresh"`
	CreatedAt               timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	HasRotatedSecret        types.Bool                                                      `tfsdk:"has_rotated_secret" json:"has_rotated_secret,computed"`
	PromotedAt              timetypes.RFC3339                                               `tfsdk:"promoted_at" json:"promoted_at,computed" format:"date-time"`
	UpdatedAt               timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ClientURIVerification   customfield.NestedObject[OAuthClientClientURIVerificationModel] `tfsdk:"client_uri_verification" json:"client_uri_verification,computed"`
}

func (m OAuthClientModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OAuthClientModel) MarshalJSONForUpdate(state OAuthClientModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type OAuthClientClientURIVerificationModel struct {
	Status types.String `tfsdk:"status" json:"status,computed"`
	Text   types.String `tfsdk:"text" json:"text,computed"`
}
