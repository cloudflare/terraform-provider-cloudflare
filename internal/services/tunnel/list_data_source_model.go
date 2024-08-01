// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelsResultListDataSourceEnvelope struct {
	Result *[]*TunnelsResultDataSourceModel `json:"result,computed"`
}

type TunnelsDataSourceModel struct {
	AccountID     types.String                     `tfsdk:"account_id" path:"account_id"`
	ExcludePrefix types.String                     `tfsdk:"exclude_prefix" query:"exclude_prefix"`
	ExistedAt     timetypes.RFC3339                `tfsdk:"existed_at" query:"existed_at"`
	IncludePrefix types.String                     `tfsdk:"include_prefix" query:"include_prefix"`
	IsDeleted     types.Bool                       `tfsdk:"is_deleted" query:"is_deleted"`
	Name          types.String                     `tfsdk:"name" query:"name"`
	PerPage       types.Float64                    `tfsdk:"per_page" query:"per_page"`
	Status        types.String                     `tfsdk:"status" query:"status"`
	TunTypes      types.String                     `tfsdk:"tun_types" query:"tun_types"`
	UUID          types.String                     `tfsdk:"uuid" query:"uuid"`
	WasActiveAt   timetypes.RFC3339                `tfsdk:"was_active_at" query:"was_active_at"`
	WasInactiveAt timetypes.RFC3339                `tfsdk:"was_inactive_at" query:"was_inactive_at"`
	Page          types.Float64                    `tfsdk:"page" query:"page"`
	MaxItems      types.Int64                      `tfsdk:"max_items"`
	Result        *[]*TunnelsResultDataSourceModel `tfsdk:"result"`
}

type TunnelsResultDataSourceModel struct {
}
