// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lock

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*R2BucketLockDataSource)(nil)

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
			"rules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[R2BucketLockRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for this rule",
							Computed:    true,
						},
						"condition": schema.SingleNestedAttribute{
							Description: "Condition to apply a lock rule to an object for how long in seconds",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketLockRulesConditionDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"max_age_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Description: "available values: \"Age\"",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"Age",
											"Date",
											"Indefinite",
										),
									},
								},
								"date": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether or not this rule is in effect",
							Computed:    true,
						},
						"prefix": schema.StringAttribute{
							Description: "Rule will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *R2BucketLockDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2BucketLockDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
