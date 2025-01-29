// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_cookies

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldCookiesResultDataSourceEnvelope struct {
	Result PageShieldCookiesDataSourceModel `json:"result,computed"`
}

type PageShieldCookiesDataSourceModel struct {
	CookieID          types.String                   `tfsdk:"cookie_id" path:"cookie_id,required"`
	ZoneID            types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	DomainAttribute   types.String                   `tfsdk:"domain_attribute" json:"domain_attribute,computed"`
	ExpiresAttribute  timetypes.RFC3339              `tfsdk:"expires_attribute" json:"expires_attribute,computed" format:"date-time"`
	FirstSeenAt       timetypes.RFC3339              `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
	Host              types.String                   `tfsdk:"host" json:"host,computed"`
	HTTPOnlyAttribute types.Bool                     `tfsdk:"http_only_attribute" json:"http_only_attribute,computed"`
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	LastSeenAt        timetypes.RFC3339              `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	MaxAgeAttribute   types.Int64                    `tfsdk:"max_age_attribute" json:"max_age_attribute,computed"`
	Name              types.String                   `tfsdk:"name" json:"name,computed"`
	PathAttribute     types.String                   `tfsdk:"path_attribute" json:"path_attribute,computed"`
	SameSiteAttribute types.String                   `tfsdk:"same_site_attribute" json:"same_site_attribute,computed"`
	SecureAttribute   types.Bool                     `tfsdk:"secure_attribute" json:"secure_attribute,computed"`
	Type              types.String                   `tfsdk:"type" json:"type,computed"`
	PageURLs          customfield.List[types.String] `tfsdk:"page_urls" json:"page_urls,computed"`
}

func (m *PageShieldCookiesDataSourceModel) toReadParams(_ context.Context) (params page_shield.CookieGetParams, diags diag.Diagnostics) {
	params = page_shield.CookieGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
