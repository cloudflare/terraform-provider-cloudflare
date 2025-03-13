// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketEventNotificationResultDataSourceEnvelope struct {
	Result R2BucketEventNotificationDataSourceModel `json:"result,computed"`
}

type R2BucketEventNotificationDataSourceModel struct {
	AccountID  types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	BucketName types.String                                                                 `tfsdk:"bucket_name" path:"bucket_name,required"`
	BucketName types.String                                                                 `tfsdk:"bucket_name" json:"bucketName,computed"`
	Queues     customfield.NestedObjectList[R2BucketEventNotificationQueuesDataSourceModel] `tfsdk:"queues" json:"queues,computed"`
}

func (m *R2BucketEventNotificationDataSourceModel) toReadParams(_ context.Context) (params r2.BucketEventNotificationGetParams, diags diag.Diagnostics) {
	params = r2.BucketEventNotificationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2BucketEventNotificationQueuesDataSourceModel struct {
	QueueID   types.String                                                                      `tfsdk:"queue_id" json:"queueId,computed"`
	QueueName types.String                                                                      `tfsdk:"queue_name" json:"queueName,computed"`
	Rules     customfield.NestedObjectList[R2BucketEventNotificationQueuesRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

type R2BucketEventNotificationQueuesRulesDataSourceModel struct {
	Actions     customfield.List[types.String] `tfsdk:"actions" json:"actions,computed"`
	CreatedAt   types.String                   `tfsdk:"created_at" json:"createdAt,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	Prefix      types.String                   `tfsdk:"prefix" json:"prefix,computed"`
	RuleID      types.String                   `tfsdk:"rule_id" json:"ruleId,computed"`
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
}
