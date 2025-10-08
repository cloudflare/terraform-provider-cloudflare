// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomHostnameFallbackOriginDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "This is the time the fallback origin was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"origin": schema.StringAttribute{
				Description: "Your origin hostname that requests to your custom hostnames will be sent to.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the fallback origin's activation.\nAvailable values: \"initializing\", \"pending_deployment\", \"pending_deletion\", \"active\", \"deployment_timed_out\", \"deletion_timed_out\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"initializing",
						"pending_deployment",
						"pending_deletion",
						"active",
						"deployment_timed_out",
						"deletion_timed_out",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "This is the time the fallback origin was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"errors": schema.ListAttribute{
				Description: "These are errors that were encountered while trying to activate a fallback origin.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *CustomHostnameFallbackOriginDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomHostnameFallbackOriginDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
