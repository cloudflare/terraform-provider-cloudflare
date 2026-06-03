// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_client

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OAuthClientsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[OAuthClientsResultDataSourceModel] `json:"result,computed"`
}

type OAuthClientsDataSourceModel struct {
	AccountID types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                     `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[OAuthClientsResultDataSourceModel] `tfsdk:"result"`
}

func (m *OAuthClientsDataSourceModel) toListParams(_ context.Context) (params iam.OAuthClientListParams, diags diag.Diagnostics) {
	params = iam.OAuthClientListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type OAuthClientsResultDataSourceModel struct {
	ClientID                types.String                                                               `tfsdk:"client_id" json:"client_id,computed"`
	Visibility              types.String                                                               `tfsdk:"visibility" json:"visibility,computed"`
	AllowedCORSOrigins      customfield.List[types.String]                                             `tfsdk:"allowed_cors_origins" json:"allowed_cors_origins,computed"`
	ClientName              types.String                                                               `tfsdk:"client_name" json:"client_name,computed"`
	ClientURI               types.String                                                               `tfsdk:"client_uri" json:"client_uri,computed"`
	ClientURIVerification   customfield.NestedObject[OAuthClientsClientURIVerificationDataSourceModel] `tfsdk:"client_uri_verification" json:"client_uri_verification,computed"`
	CreatedAt               timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	GrantTypes              customfield.List[types.String]                                             `tfsdk:"grant_types" json:"grant_types,computed"`
	HasRotatedSecret        types.Bool                                                                 `tfsdk:"has_rotated_secret" json:"has_rotated_secret,computed"`
	LogoURI                 types.String                                                               `tfsdk:"logo_uri" json:"logo_uri,computed"`
	PolicyURI               types.String                                                               `tfsdk:"policy_uri" json:"policy_uri,computed"`
	PostLogoutRedirectURIs  customfield.List[types.String]                                             `tfsdk:"post_logout_redirect_uris" json:"post_logout_redirect_uris,computed"`
	PromotedAt              timetypes.RFC3339                                                          `tfsdk:"promoted_at" json:"promoted_at,computed" format:"date-time"`
	RedirectURIs            customfield.List[types.String]                                             `tfsdk:"redirect_uris" json:"redirect_uris,computed"`
	ResponseTypes           customfield.List[types.String]                                             `tfsdk:"response_types" json:"response_types,computed"`
	Scopes                  customfield.List[types.String]                                             `tfsdk:"scopes" json:"scopes,computed"`
	TokenEndpointAuthMethod types.String                                                               `tfsdk:"token_endpoint_auth_method" json:"token_endpoint_auth_method,computed"`
	TosURI                  types.String                                                               `tfsdk:"tos_uri" json:"tos_uri,computed"`
	UpdatedAt               timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type OAuthClientsClientURIVerificationDataSourceModel struct {
	Status types.String `tfsdk:"status" json:"status,computed"`
	Text   types.String `tfsdk:"text" json:"text,computed"`
}
