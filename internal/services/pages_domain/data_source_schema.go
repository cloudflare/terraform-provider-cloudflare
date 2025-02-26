// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PagesDomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Name of the domain.",
				Computed:    true,
			},
			"domain_name": schema.StringAttribute{
				Description: "Name of the domain.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Name of the project.",
				Required:    true,
			},
			"certificate_authority": schema.StringAttribute{
				Description: "available values: \"google\", \"lets_encrypt\"",
				Computed:    true,
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
				Description: "available values: \"initializing\", \"pending\", \"active\", \"deactivated\", \"blocked\", \"error\"",
				Computed:    true,
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
			"zone_tag": schema.StringAttribute{
				Computed: true,
			},
			"validation_data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesDomainValidationDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_message": schema.StringAttribute{
						Computed: true,
					},
					"method": schema.StringAttribute{
						Description: "available values: \"http\", \"txt\"",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("http", "txt"),
						},
					},
					"status": schema.StringAttribute{
						Description: "available values: \"initializing\", \"pending\", \"active\", \"deactivated\", \"error\"",
						Computed:    true,
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
				CustomType: customfield.NewNestedObjectType[PagesDomainVerificationDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_message": schema.StringAttribute{
						Computed: true,
					},
					"status": schema.StringAttribute{
						Description: "available values: \"pending\", \"active\", \"deactivated\", \"blocked\", \"error\"",
						Computed:    true,
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
		},
	}
}

func (d *PagesDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PagesDomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
