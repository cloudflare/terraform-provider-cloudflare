// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lifecycle

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketLifecycleResultEnvelope struct {
	Result R2BucketLifecycleModel `json:"result"`
}

type R2BucketLifecycleModel struct {
	AccountID    types.String                    `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String                    `tfsdk:"bucket_name" path:"bucket_name,required"`
	Jurisdiction types.String                    `tfsdk:"jurisdiction" json:"-,computed_optional,no_refresh"`
	Rules        *[]*R2BucketLifecycleRulesModel `tfsdk:"rules" json:"rules,optional"`
}

func (m R2BucketLifecycleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketLifecycleModel) MarshalJSONForUpdate(state R2BucketLifecycleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketLifecycleRulesModel struct {
	ID                              types.String                                                `tfsdk:"id" json:"id,required"`
	Conditions                      *R2BucketLifecycleRulesConditionsModel                      `tfsdk:"conditions" json:"conditions,required"`
	Enabled                         types.Bool                                                  `tfsdk:"enabled" json:"enabled,required"`
	AbortMultipartUploadsTransition *R2BucketLifecycleRulesAbortMultipartUploadsTransitionModel `tfsdk:"abort_multipart_uploads_transition" json:"abortMultipartUploadsTransition,optional"`
	DeleteObjectsTransition         *R2BucketLifecycleRulesDeleteObjectsTransitionModel         `tfsdk:"delete_objects_transition" json:"deleteObjectsTransition,optional"`
	StorageClassTransitions         *[]*R2BucketLifecycleRulesStorageClassTransitionsModel      `tfsdk:"storage_class_transitions" json:"storageClassTransitions,optional"`
}

type R2BucketLifecycleRulesConditionsModel struct {
	Prefix types.String `tfsdk:"prefix" json:"prefix,required"`
}

type R2BucketLifecycleRulesAbortMultipartUploadsTransitionModel struct {
	Condition *R2BucketLifecycleRulesAbortMultipartUploadsTransitionConditionModel `tfsdk:"condition" json:"condition,optional"`
}

type R2BucketLifecycleRulesAbortMultipartUploadsTransitionConditionModel struct {
	MaxAge types.Int64  `tfsdk:"max_age" json:"maxAge,required"`
	Type   types.String `tfsdk:"type" json:"type,required"`
}

type R2BucketLifecycleRulesDeleteObjectsTransitionModel struct {
	Condition *R2BucketLifecycleRulesDeleteObjectsTransitionConditionModel `tfsdk:"condition" json:"condition,optional"`
}

type R2BucketLifecycleRulesDeleteObjectsTransitionConditionModel struct {
	MaxAge types.Int64       `tfsdk:"max_age" json:"maxAge,optional"`
	Type   types.String      `tfsdk:"type" json:"type,required"`
	Date   timetypes.RFC3339 `tfsdk:"date" json:"date,optional" format:"date"`
}

type R2BucketLifecycleRulesStorageClassTransitionsModel struct {
	Condition    *R2BucketLifecycleRulesStorageClassTransitionsConditionModel `tfsdk:"condition" json:"condition,required"`
	StorageClass types.String                                                 `tfsdk:"storage_class" json:"storageClass,required"`
}

type R2BucketLifecycleRulesStorageClassTransitionsConditionModel struct {
	MaxAge types.Int64       `tfsdk:"max_age" json:"maxAge,optional"`
	Type   types.String      `tfsdk:"type" json:"type,required"`
	Date   timetypes.RFC3339 `tfsdk:"date" json:"date,optional" format:"date"`
}
