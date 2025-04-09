// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
				Description:   "Account ID",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bucket_name": schema.StringAttribute{
				Description:   "Name of the bucket",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue_id": schema.StringAttribute{
				Description:   "Queue ID",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rules": schema.ListNestedAttribute{
				Description: "Array of rules to drive notifications",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[R2BucketEventNotificationRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"actions": schema.ListAttribute{
							Description: "Array of R2 object actions that will trigger notifications",
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
							Description: "A description that can be used to identify the event notification rule after creation",
							Optional:    true,
						},
						"prefix": schema.StringAttribute{
							Description: "Notifications will be sent only for objects with this prefix",
							Optional:    true,
						},
						"suffix": schema.StringAttribute{
							Description: "Notifications will be sent only for objects with this suffix",
							Optional:    true,
						},
					},
				},
			},
			"queues": schema.ListNestedAttribute{
				Description: "List of queues associated with the bucket.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[R2BucketEventNotificationQueuesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"queue_id": schema.StringAttribute{
							Description: "Queue ID",
							Computed:    true,
						},
						"queue_name": schema.StringAttribute{
							Description: "Name of the queue",
							Computed:    true,
						},
						"rules": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[R2BucketEventNotificationQueuesRulesModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"actions": schema.ListAttribute{
										Description: "Array of R2 object actions that will trigger notifications",
										Computed:    true,
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
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"created_at": schema.StringAttribute{
										Description: "Timestamp when the rule was created",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "A description that can be used to identify the event notification rule after creation",
										Computed:    true,
									},
									"prefix": schema.StringAttribute{
										Description: "Notifications will be sent only for objects with this prefix",
										Computed:    true,
									},
									"rule_id": schema.StringAttribute{
										Description: "Rule ID",
										Computed:    true,
									},
									"suffix": schema.StringAttribute{
										Description: "Notifications will be sent only for objects with this suffix",
										Computed:    true,
									},
								},
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
