// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersScriptDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the script was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"etag": schema.StringAttribute{
				Description: "Hashed script content, can be used in a If-None-Match header when updating.",
				Computed:    true,
			},
			"has_assets": schema.BoolAttribute{
				Description: "Whether a Worker contains assets.",
				Computed:    true,
			},
			"has_modules": schema.BoolAttribute{
				Description: "Whether a Worker contains modules.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "The id of the script in the Workers system. Usually the script name.",
				Computed:    true,
			},
			"logpush": schema.BoolAttribute{
				Description: "Whether Logpush is turned on for the Worker.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the script was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"placement_mode": schema.StringAttribute{
				Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("smart"),
				},
			},
			"placement_status": schema.StringAttribute{
				Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"SUCCESS",
						"NO_VALID_HOSTS",
						"NO_VALID_BINDINGS",
						"UNSUPPORTED_APPLICATION",
						"INSUFFICIENT_INVOCATIONS",
						"INSUFFICIENT_SUBREQUESTS",
					),
				},
			},
			"usage_model": schema.StringAttribute{
				Description: "Usage model for the Worker invocations.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("standard"),
				},
			},
			"placement": schema.SingleNestedAttribute{
				Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WorkersScriptPlacementDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("smart"),
						},
					},
					"status": schema.StringAttribute{
						Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"SUCCESS",
								"NO_VALID_HOSTS",
								"NO_VALID_BINDINGS",
								"UNSUPPORTED_APPLICATION",
								"INSUFFICIENT_INVOCATIONS",
								"INSUFFICIENT_SUBREQUESTS",
							),
						},
					},
				},
			},
			"tail_consumers": schema.ListNestedAttribute{
				Description: "List of Workers that will consume logs from the attached Worker.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkersScriptTailConsumersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service": schema.StringAttribute{
							Description: "Name of Worker that is to be the consumer.",
							Computed:    true,
						},
						"environment": schema.StringAttribute{
							Description: "Optional environment if the Worker utilizes one.",
							Computed:    true,
						},
						"namespace": schema.StringAttribute{
							Description: "Optional dispatch namespace the script belongs to.",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *WorkersScriptDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersScriptDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("script_name")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("script_name")),
	}
}
