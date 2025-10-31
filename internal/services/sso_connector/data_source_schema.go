// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*SSOConnectorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "SSO Connector identifier tag.",
				Computed:    true,
			},
			"sso_connector_id": schema.StringAttribute{
				Description: "SSO Connector identifier tag.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "Timestamp for the creation of the SSO connector",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"email_domain": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"updated_on": schema.StringAttribute{
				Description: "Timestamp for the last update of the SSO connector",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"use_fedramp_language": schema.BoolAttribute{
				Description: "Controls the display of FedRAMP language to the user during SSO login",
				Computed:    true,
			},
			"verification": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SSOConnectorVerificationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"code": schema.StringAttribute{
						Description: "DNS verification code. Add this entire string to the DNS TXT record of the email domain to validate ownership.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "The status of the verification code from the verification process.\nAvailable values: \"awaiting\", \"pending\", \"failed\", \"verified\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"awaiting",
								"pending",
								"failed",
								"verified",
							),
						},
					},
				},
			},
		},
	}
}

func (d *SSOConnectorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SSOConnectorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
