// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketEventNotificationResultDataSourceEnvelope struct {
	Result R2BucketEventNotificationDataSourceModel `json:"result,computed"`
}

type R2BucketEventNotificationDataSourceModel struct {
	AccountID                       types.String                                                                                      `tfsdk:"account_id" path:"account_id,required"`
	BucketName                      types.String                                                                                      `tfsdk:"bucket_name" path:"bucket_name,required"`
	QueueID                         types.String                                                                                      `tfsdk:"queue_id" path:"queue_id,required"`
	Enabled                         types.Bool                                                                                        `tfsdk:"enabled" json:"enabled,computed"`
	ID                              types.String                                                                                      `tfsdk:"id" json:"id,computed"`
	AbortMultipartUploadsTransition customfield.NestedObject[R2BucketEventNotificationAbortMultipartUploadsTransitionDataSourceModel] `tfsdk:"abort_multipart_uploads_transition" json:"abortMultipartUploadsTransition,computed"`
	Conditions                      customfield.NestedObject[R2BucketEventNotificationConditionsDataSourceModel]                      `tfsdk:"conditions" json:"conditions,computed"`
	DeleteObjectsTransition         customfield.NestedObject[R2BucketEventNotificationDeleteObjectsTransitionDataSourceModel]         `tfsdk:"delete_objects_transition" json:"deleteObjectsTransition,computed"`
	StorageClassTransitions         customfield.NestedObjectList[R2BucketEventNotificationStorageClassTransitionsDataSourceModel]     `tfsdk:"storage_class_transitions" json:"storageClassTransitions,computed"`
}

func (m *R2BucketEventNotificationDataSourceModel) toReadParams(_ context.Context) (params r2.BucketEventNotificationGetParams, diags diag.Diagnostics) {
	params = r2.BucketEventNotificationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2BucketEventNotificationAbortMultipartUploadsTransitionDataSourceModel struct {
	Condition customfield.NestedObject[R2BucketEventNotificationAbortMultipartUploadsTransitionConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
}

type R2BucketEventNotificationAbortMultipartUploadsTransitionConditionDataSourceModel struct {
	MaxAge types.Int64  `tfsdk:"max_age" json:"maxAge,computed"`
	Type   types.String `tfsdk:"type" json:"type,computed"`
}

type R2BucketEventNotificationConditionsDataSourceModel struct {
	Prefix types.String `tfsdk:"prefix" json:"prefix,computed"`
}

type R2BucketEventNotificationDeleteObjectsTransitionDataSourceModel struct {
	Condition customfield.NestedObject[R2BucketEventNotificationDeleteObjectsTransitionConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
}

type R2BucketEventNotificationDeleteObjectsTransitionConditionDataSourceModel struct {
	MaxAge types.Int64       `tfsdk:"max_age" json:"maxAge,computed"`
	Type   types.String      `tfsdk:"type" json:"type,computed"`
	Date   timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}

type R2BucketEventNotificationStorageClassTransitionsDataSourceModel struct {
	Condition    customfield.NestedObject[R2BucketEventNotificationStorageClassTransitionsConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
	StorageClass types.String                                                                                       `tfsdk:"storage_class" json:"storageClass,computed"`
}

type R2BucketEventNotificationStorageClassTransitionsConditionDataSourceModel struct {
	MaxAge types.Int64       `tfsdk:"max_age" json:"maxAge,computed"`
	Type   types.String      `tfsdk:"type" json:"type,computed"`
	Date   timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}
