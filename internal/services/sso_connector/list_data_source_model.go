// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SSOConnectorsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SSOConnectorsResultDataSourceModel] `json:"result,computed"`
}

type SSOConnectorsDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[SSOConnectorsResultDataSourceModel] `tfsdk:"result"`
}

type SSOConnectorsResultDataSourceModel struct {
	ID                 types.String                                                       `tfsdk:"id" json:"id,computed"`
	CreatedOn          timetypes.RFC3339                                                  `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	EmailDomain        types.String                                                       `tfsdk:"email_domain" json:"email_domain,computed"`
	Enabled            types.Bool                                                         `tfsdk:"enabled" json:"enabled,computed"`
	UpdatedOn          timetypes.RFC3339                                                  `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
	UseFedrampLanguage types.Bool                                                         `tfsdk:"use_fedramp_language" json:"use_fedramp_language,computed"`
	Verification       customfield.NestedObject[SSOConnectorsVerificationDataSourceModel] `tfsdk:"verification" json:"verification,computed"`
}

type SSOConnectorsVerificationDataSourceModel struct {
	Code   types.String `tfsdk:"code" json:"code,computed"`
	Status types.String `tfsdk:"status" json:"status,computed"`
}
