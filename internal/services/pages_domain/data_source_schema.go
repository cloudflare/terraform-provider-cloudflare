// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &PagesDomainDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"domain_name": schema.StringAttribute{
				Description: "Name of the domain.",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Name of the project.",
				Optional:    true,
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
			"id": schema.StringAttribute{
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
			"zone_tag": schema.StringAttribute{
				Computed: true,
			},
			"validation_data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesDomainValidationDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_message": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"method": schema.StringAttribute{
						Computed: true,
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("http", "txt"),
						},
					},
					"status": schema.StringAttribute{
						Computed: true,
						Optional: true,
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
						Optional: true,
					},
					"txt_value": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"verification_data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesDomainVerificationDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_message": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"status": schema.StringAttribute{
						Computed: true,
						Optional: true,
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
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"project_name": schema.StringAttribute{
						Description: "Name of the project.",
						Required:    true,
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
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(
			path.MatchRoot("account_id"),
			path.MatchRoot("domain_name"),
			path.MatchRoot("project_name"),
		),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("domain_name")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("project_name")),
	}
}
