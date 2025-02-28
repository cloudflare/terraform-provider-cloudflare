// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketEventNotificationResultEnvelope struct {
	Result R2BucketEventNotificationModel `json:"result"`
}

type R2BucketEventNotificationModel struct {
	AccountID    types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String                                                       `tfsdk:"bucket_name" path:"bucket_name,required"`
	Jurisdiction types.String                                                       `tfsdk:"jurisdiction" json:"-,computed_optional"`
	QueueID      types.String                                                       `tfsdk:"queue_id" path:"queue_id,optional"`
	Rules        customfield.NestedObjectList[R2BucketEventNotificationRulesModel]  `tfsdk:"rules" json:"rules,computed_optional"`
	Queues       customfield.NestedObjectList[R2BucketEventNotificationQueuesModel] `tfsdk:"queues" json:"queues,computed"`
}

func (m R2BucketEventNotificationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketEventNotificationModel) MarshalJSONForUpdate(state R2BucketEventNotificationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketEventNotificationRulesModel struct {
	Actions     *[]types.String `tfsdk:"actions" json:"actions,required"`
	Description types.String    `tfsdk:"description" json:"description,optional"`
	Prefix      types.String    `tfsdk:"prefix" json:"prefix,optional"`
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,optional"`
}

type R2BucketEventNotificationQueuesModel struct {
	QueueID   types.String                                                            `tfsdk:"queue_id" json:"queueId,computed"`
	QueueName types.String                                                            `tfsdk:"queue_name" json:"queueName,computed"`
	Rules     customfield.NestedObjectList[R2BucketEventNotificationQueuesRulesModel] `tfsdk:"rules" json:"rules,computed"`
}

type R2BucketEventNotificationQueuesRulesModel struct {
	Actions     customfield.List[types.String] `tfsdk:"actions" json:"actions,computed"`
	CreatedAt   types.String                   `tfsdk:"created_at" json:"createdAt,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	Prefix      types.String                   `tfsdk:"prefix" json:"prefix,computed"`
	RuleID      types.String                   `tfsdk:"rule_id" json:"ruleId,computed"`
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
}
