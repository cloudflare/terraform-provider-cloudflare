// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lifecycle

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*R2BucketLifecycleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID.",
				Required:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket.",
				Required:    true,
			},
			"rules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[R2BucketLifecycleRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for this rule.",
							Computed:    true,
						},
						"conditions": schema.SingleNestedAttribute{
							Description: "Conditions that apply to all transitions of this rule.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesConditionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"prefix": schema.StringAttribute{
									Description: "Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads.",
									Computed:    true,
								},
							},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether or not this rule is in effect.",
							Computed:    true,
						},
						"abort_multipart_uploads_transition": schema.SingleNestedAttribute{
							Description: "Transition to abort ongoing multipart uploads.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesAbortMultipartUploadsTransitionDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"condition": schema.SingleNestedAttribute{
									Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesAbortMultipartUploadsTransitionConditionDataSourceModel](ctx),
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
						"delete_objects_transition": schema.SingleNestedAttribute{
							Description: "Transition to delete objects.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesDeleteObjectsTransitionDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"condition": schema.SingleNestedAttribute{
									Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesDeleteObjectsTransitionConditionDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[R2BucketLifecycleRulesStorageClassTransitionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"condition": schema.SingleNestedAttribute{
										Description: "Condition for lifecycle transitions to apply after an object reaches an age in seconds.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[R2BucketLifecycleRulesStorageClassTransitionsConditionDataSourceModel](ctx),
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
				},
			},
		},
	}
}

func (d *R2BucketLifecycleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2BucketLifecycleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
