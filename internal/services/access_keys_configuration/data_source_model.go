// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_keys_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessKeysConfigurationResultDataSourceEnvelope struct {
	Result AccessKeysConfigurationDataSourceModel `json:"result,computed"`
}

type AccessKeysConfigurationDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
