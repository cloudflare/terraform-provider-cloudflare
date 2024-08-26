// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapResultEnvelope struct {
	Result AddressMapModel `json:"result"`
}

type AddressMapModel struct {
	ID           types.String                                          `tfsdk:"id" json:"id,computed"`
	AccountID    types.String                                          `tfsdk:"account_id" path:"account_id"`
	IPs          *[]types.String                                       `tfsdk:"ips" json:"ips"`
	Memberships  *[]*AddressMapMembershipsModel                        `tfsdk:"memberships" json:"memberships"`
	DefaultSNI   types.String                                          `tfsdk:"default_sni" json:"default_sni"`
	Description  types.String                                          `tfsdk:"description" json:"description"`
	Enabled      types.Bool                                            `tfsdk:"enabled" json:"enabled"`
	CanDelete    types.Bool                                            `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool                                            `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339                                     `tfsdk:"created_at" json:"created_at,computed"`
	ModifiedAt   timetypes.RFC3339                                     `tfsdk:"modified_at" json:"modified_at,computed"`
	Success      types.Bool                                            `tfsdk:"success" json:"success,computed"`
	Errors       customfield.NestedObjectList[AddressMapErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages     customfield.NestedObjectList[AddressMapMessagesModel] `tfsdk:"messages" json:"messages,computed"`
	ResultInfo   customfield.NestedObject[AddressMapResultInfoModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

type AddressMapMembershipsModel struct {
	CanDelete  types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Identifier types.String      `tfsdk:"identifier" json:"identifier"`
	Kind       types.String      `tfsdk:"kind" json:"kind"`
}

type AddressMapErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type AddressMapMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type AddressMapResultInfoModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count"`
	Page       types.Float64 `tfsdk:"page" json:"page"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count"`
}
