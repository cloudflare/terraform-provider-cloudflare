package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts v4 cloudflare_access_identity_provider state to v5 cloudflare_zero_trust_access_identity_provider.
func Transform(ctx context.Context, source SourceAccessIdentityProviderModel) (*TargetAccessIdentityProviderModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError("Missing required field", "name is required for access identity provider migration")
		return nil, diags
	}
	if source.Type.IsNull() || source.Type.IsUnknown() {
		diags.AddError("Missing required field", "type is required for access identity provider migration")
		return nil, diags
	}

	target := &TargetAccessIdentityProviderModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		ZoneID:    source.ZoneID,
		Name:      source.Name,
		Type:      source.Type,
	}

	if len(source.Config) > 0 {
		targetConfig, configDiags := transformConfig(ctx, source.Config[0])
		diags.Append(configDiags...)
		if !diags.HasError() {
			target.Config = targetConfig
		}
	} else {
		target.Config = &TargetConfigModel{}
	}

	if len(source.ScimConfig) > 0 {
		targetScim := transformScimConfig(source.ScimConfig[0])
		scimObj, scimDiags := customfield.NewObject(ctx, &targetScim)
		diags.Append(scimDiags...)
		if !diags.HasError() {
			target.SCIMConfig = scimObj
		}
	} else {
		target.SCIMConfig = customfield.NullObject[TargetScimConfigModel](ctx)
	}

	return target, diags
}

func transformConfig(ctx context.Context, source SourceConfigModel) (*TargetConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetConfigModel{
		ClientID:              migrations.FalseyStringToNull(source.ClientID),
		ClientSecret:          migrations.FalseyStringToNull(source.ClientSecret),
		DirectoryID:           migrations.FalseyStringToNull(source.DirectoryID),
		EmailClaimName:        migrations.FalseyStringToNull(source.EmailClaimName),
		Prompt:                migrations.FalseyStringToNull(source.Prompt),
		CentrifyAccount:       migrations.FalseyStringToNull(source.CentrifyAccount),
		CentrifyAppID:         migrations.FalseyStringToNull(source.CentrifyAppID),
		AppsDomain:            migrations.FalseyStringToNull(source.AppsDomain),
		AuthURL:               migrations.FalseyStringToNull(source.AuthURL),
		CERTsURL:              migrations.FalseyStringToNull(source.CERTsURL),
		TokenURL:              migrations.FalseyStringToNull(source.TokenURL),
		AuthorizationServerID: migrations.FalseyStringToNull(source.AuthorizationServerID),
		OktaAccount:           migrations.FalseyStringToNull(source.OktaAccount),
		OneloginAccount:       migrations.FalseyStringToNull(source.OneloginAccount),
		PingEnvID:             migrations.FalseyStringToNull(source.PingEnvID),
		EmailAttributeName:    migrations.FalseyStringToNull(source.EmailAttributeName),
		IssuerURL:             migrations.FalseyStringToNull(source.IssuerURL),
		SSOTargetURL:          migrations.FalseyStringToNull(source.SSOTargetURL),
		RedirectURL:           source.RedirectURL,

		ConditionalAccessEnabled: migrations.FalseyBoolToNull(source.ConditionalAccessEnabled),
		SupportGroups:            migrations.FalseyBoolToNull(source.SupportGroups),
		PKCEEnabled:              migrations.FalseyBoolToNull(source.PKCEEnabled),
		SignRequest:              migrations.FalseyBoolToNull(source.SignRequest),
	}

	// api_token: deprecated, not copied

	if !source.Claims.IsNull() && !source.Claims.IsUnknown() {
		var v []types.String
		diags.Append(source.Claims.ElementsAs(ctx, &v, false)...)
		if !diags.HasError() && len(v) > 0 {
			target.Claims = &v
		}
	}

	if !source.Scopes.IsNull() && !source.Scopes.IsUnknown() {
		var v []types.String
		diags.Append(source.Scopes.ElementsAs(ctx, &v, false)...)
		if !diags.HasError() && len(v) > 0 {
			target.Scopes = &v
		}
	}

	if !source.Attributes.IsNull() && !source.Attributes.IsUnknown() {
		var v []types.String
		diags.Append(source.Attributes.ElementsAs(ctx, &v, false)...)
		if !diags.HasError() && len(v) > 0 {
			target.Attributes = &v
		}
	}

	if !source.HeaderAttributes.IsNull() && !source.HeaderAttributes.IsUnknown() {
		var src []SourceHeaderAttributesModel
		diags.Append(source.HeaderAttributes.ElementsAs(ctx, &src, false)...)
		if !diags.HasError() && len(src) > 0 {
			dst := make([]*TargetHeaderAttributesModel, len(src))
			for i, h := range src {
				dst[i] = &TargetHeaderAttributesModel{
					AttributeName: h.AttributeName,
					HeaderName:    h.HeaderName,
				}
			}
			target.HeaderAttributes = &dst
		}
	}

	// idp_public_cert (string) → idp_public_certs (list)
	if !source.IdpPublicCert.IsNull() && !source.IdpPublicCert.IsUnknown() && source.IdpPublicCert.ValueString() != "" {
		certs := []types.String{source.IdpPublicCert}
		target.IdPPublicCERTs = &certs
	}

	return target, diags
}

func transformScimConfig(source SourceScimConfigModel) TargetScimConfigModel {
	return TargetScimConfigModel{
		Enabled:                source.Enabled,
		IdentityUpdateBehavior: source.IdentityUpdateBehavior,
		SCIMBaseURL:            source.SCIMBaseURL,
		SeatDeprovision:        source.SeatDeprovision,
		Secret:                 source.Secret,
		UserDeprovision:        source.UserDeprovision,
		// group_member_deprovision: deprecated, not copied
	}
}

type SourceHeaderAttributesModel struct {
	AttributeName types.String `tfsdk:"attribute_name"`
	HeaderName    types.String `tfsdk:"header_name"`
}
