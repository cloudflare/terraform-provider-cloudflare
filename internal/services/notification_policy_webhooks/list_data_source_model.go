// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksListResultListDataSourceEnvelope struct {
	Result *[]*NotificationPolicyWebhooksListItemsDataSourceModel `json:"result,computed"`
}

type NotificationPolicyWebhooksListDataSourceModel struct {
	AccountID types.String                                           `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                            `tfsdk:"max_items"`
	Items     *[]*NotificationPolicyWebhooksListItemsDataSourceModel `tfsdk:"items"`
}

type NotificationPolicyWebhooksListItemsDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String `tfsdk:"created_at" json:"created_at,computed"`
	LastFailure types.String `tfsdk:"last_failure" json:"last_failure,computed"`
	LastSuccess types.String `tfsdk:"last_success" json:"last_success,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Secret      types.String `tfsdk:"secret" json:"secret,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
	URL         types.String `tfsdk:"url" json:"url,computed"`
}
