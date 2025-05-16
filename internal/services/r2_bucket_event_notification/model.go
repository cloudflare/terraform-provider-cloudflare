// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketEventNotificationResultEnvelope struct {
	Result R2BucketEventNotificationModel `json:"result"`
}

type R2BucketEventNotificationModel struct {
	AccountID                       types.String                                                                            `tfsdk:"account_id" path:"account_id,required"`
	BucketName                      types.String                                                                            `tfsdk:"bucket_name" path:"bucket_name,required"`
	QueueID                         types.String                                                                            `tfsdk:"queue_id" path:"queue_id,required"`
	Rules                           *[]*R2BucketEventNotificationRulesModel                                                 `tfsdk:"rules" json:"rules,optional,no_refresh"`
	Enabled                         types.Bool                                                                              `tfsdk:"enabled" json:"enabled,computed"`
	ID                              types.String                                                                            `tfsdk:"id" json:"id,computed"`
	AbortMultipartUploadsTransition customfield.NestedObject[R2BucketEventNotificationAbortMultipartUploadsTransitionModel] `tfsdk:"abort_multipart_uploads_transition" json:"abortMultipartUploadsTransition,computed"`
	Conditions                      customfield.NestedObject[R2BucketEventNotificationConditionsModel]                      `tfsdk:"conditions" json:"conditions,computed"`
	DeleteObjectsTransition         customfield.NestedObject[R2BucketEventNotificationDeleteObjectsTransitionModel]         `tfsdk:"delete_objects_transition" json:"deleteObjectsTransition,computed"`
	StorageClassTransitions         customfield.NestedObjectList[R2BucketEventNotificationStorageClassTransitionsModel]     `tfsdk:"storage_class_transitions" json:"storageClassTransitions,computed"`
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

type R2BucketEventNotificationAbortMultipartUploadsTransitionModel struct {
	Condition customfield.NestedObject[R2BucketEventNotificationAbortMultipartUploadsTransitionConditionModel] `tfsdk:"condition" json:"condition,computed"`
}

type R2BucketEventNotificationAbortMultipartUploadsTransitionConditionModel struct {
	MaxAge types.Int64  `tfsdk:"max_age" json:"maxAge,computed"`
	Type   types.String `tfsdk:"type" json:"type,computed"`
}

type R2BucketEventNotificationConditionsModel struct {
	Prefix types.String `tfsdk:"prefix" json:"prefix,computed"`
}

type R2BucketEventNotificationDeleteObjectsTransitionModel struct {
	Condition customfield.NestedObject[R2BucketEventNotificationDeleteObjectsTransitionConditionModel] `tfsdk:"condition" json:"condition,computed"`
}

type R2BucketEventNotificationDeleteObjectsTransitionConditionModel struct {
	MaxAge types.Int64       `tfsdk:"max_age" json:"maxAge,computed"`
	Type   types.String      `tfsdk:"type" json:"type,computed"`
	Date   timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}

type R2BucketEventNotificationStorageClassTransitionsModel struct {
	Condition    customfield.NestedObject[R2BucketEventNotificationStorageClassTransitionsConditionModel] `tfsdk:"condition" json:"condition,computed"`
	StorageClass types.String                                                                             `tfsdk:"storage_class" json:"storageClass,computed"`
}

type R2BucketEventNotificationStorageClassTransitionsConditionModel struct {
	MaxAge types.Int64       `tfsdk:"max_age" json:"maxAge,computed"`
	Type   types.String      `tfsdk:"type" json:"type,computed"`
	Date   timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}
