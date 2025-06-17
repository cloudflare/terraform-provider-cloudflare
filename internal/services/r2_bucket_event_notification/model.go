// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketEventNotificationResultEnvelope struct {
	Result R2BucketEventNotificationModel `json:"result"`
}

type R2BucketEventNotificationModel struct {
	AccountID    types.String                            `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String                            `tfsdk:"bucket_name" path:"bucket_name,required"`
	QueueID      types.String                            `tfsdk:"queue_id" path:"queue_id,required"`
	Rules        *[]*R2BucketEventNotificationRulesModel `tfsdk:"rules" json:"rules,optional"`
	QueueName    types.String                            `tfsdk:"queue_name" json:"queueName,computed"`
	Jurisdiction types.String                            `tfsdk:"jurisdiction" json:"-,computed_optional"`
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
