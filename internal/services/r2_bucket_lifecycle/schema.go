// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lifecycle

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketLifecycleResource)(nil)

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
			"rules": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[R2BucketLifecycleRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for this rule",
							Required:    true,
						},
						"conditions": schema.SingleNestedAttribute{
							Description: "Conditions that apply to all transitions of this rule",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"prefix": schema.StringAttribute{
									Description: "Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads",
									Required:    true,
								},
							},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether or not this rule is in effect",
							Required:    true,
						},
						"abort_multipart_uploads_transition": schema.SingleNestedAttribute{
							Description: "Transition to abort ongoing multipart uploads",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesAbortMultipartUploadsTransitionModel](ctx),
							Attributes: map[string]schema.Attribute{
								"condition": schema.SingleNestedAttribute{
									Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesAbortMultipartUploadsTransitionConditionModel](ctx),
									Attributes: map[string]schema.Attribute{
										"max_age": schema.Int64Attribute{
											Required: true,
										},
										"type": schema.StringAttribute{
											Required: true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("Age"),
											},
										},
									},
								},
							},
						},
						"delete_objects_transition": schema.SingleNestedAttribute{
							Description: "Transition to delete objects",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesDeleteObjectsTransitionModel](ctx),
							Attributes: map[string]schema.Attribute{
								"condition": schema.SingleNestedAttribute{
									Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesDeleteObjectsTransitionConditionModel](ctx),
									Attributes: map[string]schema.Attribute{
										"max_age": schema.Int64Attribute{
											Optional: true,
										},
										"type": schema.StringAttribute{
											Required: true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("Age", "Date"),
											},
										},
										"date": schema.StringAttribute{
											Optional:   true,
											CustomType: timetypes.RFC3339Type{},
										},
									},
								},
							},
						},
						"storage_class_transitions": schema.ListNestedAttribute{
							Description: "Transitions to change the storage class of objects",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectListType[R2BucketLifecycleRulesStorageClassTransitionsModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"condition": schema.SingleNestedAttribute{
										Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds",
										Required:    true,
										Attributes: map[string]schema.Attribute{
											"max_age": schema.Int64Attribute{
												Optional: true,
											},
											"type": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("Age", "Date"),
												},
											},
											"date": schema.StringAttribute{
												Optional:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
									"storage_class": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("InfrequentAccess"),
										},
									},
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *R2BucketLifecycleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *R2BucketLifecycleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
