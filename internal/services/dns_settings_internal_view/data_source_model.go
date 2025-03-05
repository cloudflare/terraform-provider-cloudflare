// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings_internal_view

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSSettingsInternalViewResultDataSourceEnvelope struct {
	Result DNSSettingsInternalViewDataSourceModel `json:"result,computed"`
}

type DNSSettingsInternalViewDataSourceModel struct {
	ID           types.String                                     `tfsdk:"id" json:"-,computed"`
	ViewID       types.String                                     `tfsdk:"view_id" path:"view_id,optional"`
	AccountID    types.String                                     `tfsdk:"account_id" path:"account_id,required"`
	CreatedTime  timetypes.RFC3339                                `tfsdk:"created_time" json:"created_time,computed" format:"date-time"`
	ModifiedTime timetypes.RFC3339                                `tfsdk:"modified_time" json:"modified_time,computed" format:"date-time"`
	Name         types.String                                     `tfsdk:"name" json:"name,computed"`
	Zones        customfield.List[types.String]                   `tfsdk:"zones" json:"zones,computed"`
	Filter       *DNSSettingsInternalViewFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *DNSSettingsInternalViewDataSourceModel) toReadParams(_ context.Context) (params dns.SettingViewGetParams, diags diag.Diagnostics) {
	params = dns.SettingViewGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *DNSSettingsInternalViewDataSourceModel) toListParams(_ context.Context) (params dns.SettingViewListParams, diags diag.Diagnostics) {
	params = dns.SettingViewListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(dns.SettingViewListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Match.IsNull() {
		params.Match = cloudflare.F(dns.SettingViewListParamsMatch(m.Filter.Match.ValueString()))
	}
	if m.Filter.Name != nil {
		paramsName := dns.SettingViewListParamsName{}
		if !m.Filter.Name.Contains.IsNull() {
			paramsName.Contains = cloudflare.F(m.Filter.Name.Contains.ValueString())
		}
		if !m.Filter.Name.Endswith.IsNull() {
			paramsName.Endswith = cloudflare.F(m.Filter.Name.Endswith.ValueString())
		}
		if !m.Filter.Name.Exact.IsNull() {
			paramsName.Exact = cloudflare.F(m.Filter.Name.Exact.ValueString())
		}
		if !m.Filter.Name.Startswith.IsNull() {
			paramsName.Startswith = cloudflare.F(m.Filter.Name.Startswith.ValueString())
		}
		params.Name = cloudflare.F(paramsName)
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(dns.SettingViewListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.ZoneID.IsNull() {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}
	if !m.Filter.ZoneName.IsNull() {
		params.ZoneName = cloudflare.F(m.Filter.ZoneName.ValueString())
	}

	return
}

type DNSSettingsInternalViewFindOneByDataSourceModel struct {
	Direction types.String                                 `tfsdk:"direction" query:"direction,computed_optional"`
	Match     types.String                                 `tfsdk:"match" query:"match,computed_optional"`
	Name      *DNSSettingsInternalViewsNameDataSourceModel `tfsdk:"name" query:"name,optional"`
	Order     types.String                                 `tfsdk:"order" query:"order,optional"`
	ZoneID    types.String                                 `tfsdk:"zone_id" query:"zone_id,optional"`
	ZoneName  types.String                                 `tfsdk:"zone_name" query:"zone_name,optional"`
}
