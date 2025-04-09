// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lifecycle

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/r2"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketLifecycleResultDataSourceEnvelope struct {
Result R2BucketLifecycleDataSourceModel `json:"result,computed"`
}

type R2BucketLifecycleDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
BucketName types.String `tfsdk:"bucket_name" path:"bucket_name,required"`
Rules customfield.NestedObjectList[R2BucketLifecycleRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *R2BucketLifecycleDataSourceModel) toReadParams(_ context.Context) (params r2.BucketLifecycleGetParams, diags diag.Diagnostics) {
  params = r2.BucketLifecycleGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type R2BucketLifecycleRulesDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Conditions customfield.NestedObject[R2BucketLifecycleRulesConditionsDataSourceModel] `tfsdk:"conditions" json:"conditions,computed"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
AbortMultipartUploadsTransition customfield.NestedObject[R2BucketLifecycleRulesAbortMultipartUploadsTransitionDataSourceModel] `tfsdk:"abort_multipart_uploads_transition" json:"abortMultipartUploadsTransition,computed"`
DeleteObjectsTransition customfield.NestedObject[R2BucketLifecycleRulesDeleteObjectsTransitionDataSourceModel] `tfsdk:"delete_objects_transition" json:"deleteObjectsTransition,computed"`
StorageClassTransitions customfield.NestedObjectList[R2BucketLifecycleRulesStorageClassTransitionsDataSourceModel] `tfsdk:"storage_class_transitions" json:"storageClassTransitions,computed"`
}

type R2BucketLifecycleRulesConditionsDataSourceModel struct {
Prefix types.String `tfsdk:"prefix" json:"prefix,computed"`
}

type R2BucketLifecycleRulesAbortMultipartUploadsTransitionDataSourceModel struct {
Condition customfield.NestedObject[R2BucketLifecycleRulesAbortMultipartUploadsTransitionConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
}

type R2BucketLifecycleRulesAbortMultipartUploadsTransitionConditionDataSourceModel struct {
MaxAge types.Int64 `tfsdk:"max_age" json:"maxAge,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type R2BucketLifecycleRulesDeleteObjectsTransitionDataSourceModel struct {
Condition customfield.NestedObject[R2BucketLifecycleRulesDeleteObjectsTransitionConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
}

type R2BucketLifecycleRulesDeleteObjectsTransitionConditionDataSourceModel struct {
MaxAge types.Int64 `tfsdk:"max_age" json:"maxAge,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
Date timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}

type R2BucketLifecycleRulesStorageClassTransitionsDataSourceModel struct {
Condition customfield.NestedObject[R2BucketLifecycleRulesStorageClassTransitionsConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
StorageClass types.String `tfsdk:"storage_class" json:"storageClass,computed"`
}

type R2BucketLifecycleRulesStorageClassTransitionsConditionDataSourceModel struct {
MaxAge types.Int64 `tfsdk:"max_age" json:"maxAge,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
Date timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}
