// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*R2BucketEventNotificationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID",
				Required:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket",
				Required:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket.",
				Computed:    true,
			},
			"queues": schema.ListNestedAttribute{
				Description: "List of queues associated with the bucket.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[R2BucketEventNotificationQueuesDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectListType[R2BucketEventNotificationQueuesRulesDataSourceModel](ctx),
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

func (d *R2BucketEventNotificationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2BucketEventNotificationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
