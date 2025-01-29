// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PagesDomainsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Name of the project.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PagesDomainsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"certificate_authority": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("google", "lets_encrypt"),
							},
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"domain_id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"initializing",
									"pending",
									"active",
									"deactivated",
									"blocked",
									"error",
								),
							},
						},
						"validation_data": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[PagesDomainsValidationDataDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"error_message": schema.StringAttribute{
									Computed: true,
								},
								"method": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("http", "txt"),
									},
								},
								"status": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"initializing",
											"pending",
											"active",
											"deactivated",
											"error",
										),
									},
								},
								"txt_name": schema.StringAttribute{
									Computed: true,
								},
								"txt_value": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"verification_data": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[PagesDomainsVerificationDataDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"error_message": schema.StringAttribute{
									Computed: true,
								},
								"status": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"pending",
											"active",
											"deactivated",
											"blocked",
											"error",
										),
									},
								},
							},
						},
						"zone_tag": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *PagesDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PagesDomainsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
