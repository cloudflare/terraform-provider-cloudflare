// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the settings have been created.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Email Routing settings identifier.",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the settings have been modified.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Domain of your zone.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Show the state of your account, and the type or configuration error.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ready",
						"unconfigured",
						"misconfigured",
						"misconfigured/locked",
						"unlocked",
					),
				},
			},
			"tag": schema.StringAttribute{
				Description: "Email Routing settings tag. (Deprecated, replaced by Email Routing settings identifier)",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "State of the zone settings for Email Routing.",
				Computed:    true,
				Optional:    true,
			},
			"skip_wizard": schema.BoolAttribute{
				Description: "Flag to check if the user skipped the configuration wizard.",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (d *EmailRoutingSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailRoutingSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
