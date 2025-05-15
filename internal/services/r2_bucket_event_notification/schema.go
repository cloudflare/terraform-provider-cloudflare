// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketEventNotificationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Account ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bucket_name": schema.StringAttribute{
				Description:   "Name of the bucket.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue_id": schema.StringAttribute{
				Description:   "Queue ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rules": schema.ListNestedAttribute{
				Description: "Array of rules to drive notifications.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"actions": schema.ListAttribute{
							Description: "Array of R2 object actions that will trigger notifications.",
							Required:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"PutObject",
										"CopyObject",
										"DeleteObject",
										"CompleteMultipartUpload",
										"LifecycleDeletion",
									),
								),
							},
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Description: "A description that can be used to identify the event notification rule after creation.",
							Optional:    true,
						},
						"prefix": schema.StringAttribute{
							Description: "Notifications will be sent only for objects with this prefix.",
							Optional:    true,
						},
						"suffix": schema.StringAttribute{
							Description: "Notifications will be sent only for objects with this suffix.",
							Optional:    true,
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether or not this rule is in effect.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Unique identifier for this rule.",
				Computed:    true,
			},
			"abort_multipart_uploads_transition": schema.SingleNestedAttribute{
				Description: "Transition to abort ongoing multipart uploads.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketEventNotificationAbortMultipartUploadsTransitionModel](ctx),
				Attributes: map[string]schema.Attribute{
					"condition": schema.SingleNestedAttribute{
						Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[R2BucketEventNotificationAbortMultipartUploadsTransitionConditionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"max_age": schema.Int64Attribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Description: `Available values: "Age".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("Age"),
								},
							},
						},
					},
				},
			},
			"conditions": schema.SingleNestedAttribute{
				Description: "Conditions that apply to all transitions of this rule.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketEventNotificationConditionsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"prefix": schema.StringAttribute{
						Description: "Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads.",
						Computed:    true,
					},
				},
			},
			"delete_objects_transition": schema.SingleNestedAttribute{
				Description: "Transition to delete objects.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketEventNotificationDeleteObjectsTransitionModel](ctx),
				Attributes: map[string]schema.Attribute{
					"condition": schema.SingleNestedAttribute{
						Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[R2BucketEventNotificationDeleteObjectsTransitionConditionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"max_age": schema.Int64Attribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Description: `Available values: "Age", "Date".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("Age", "Date"),
								},
							},
							"date": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
						},
					},
				},
			},
			"storage_class_transitions": schema.ListNestedAttribute{
				Description: "Transitions to change the storage class of objects.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[R2BucketEventNotificationStorageClassTransitionsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"condition": schema.SingleNestedAttribute{
							Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketEventNotificationStorageClassTransitionsConditionModel](ctx),
							Attributes: map[string]schema.Attribute{
								"max_age": schema.Int64Attribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Description: `Available values: "Age", "Date".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("Age", "Date"),
									},
								},
								"date": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
							},
						},
						"storage_class": schema.StringAttribute{
							Description: `Available values: "InfrequentAccess".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("InfrequentAccess"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *R2BucketEventNotificationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *R2BucketEventNotificationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
