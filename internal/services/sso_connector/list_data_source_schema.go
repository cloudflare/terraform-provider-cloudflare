// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*SSOConnectorsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
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
				CustomType:  customfield.NewNestedObjectListType[SSOConnectorsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "SSO Connector identifier tag.",
							Computed:    true,
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
							CustomType: customfield.NewNestedObjectType[SSOConnectorsVerificationDataSourceModel](ctx),
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
				},
			},
		},
	}
}

func (d *SSOConnectorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *SSOConnectorsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
