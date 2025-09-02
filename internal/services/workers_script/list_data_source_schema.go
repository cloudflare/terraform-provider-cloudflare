// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersScriptsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"tags": schema.StringAttribute{
				Description: "Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where allowed is 'yes' or 'no'.",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[WorkersScriptsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the script in the Workers system. Usually the script name.",
							Computed:    true,
						},
						"compatibility_date": schema.StringAttribute{
							Description: "Date indicating targeted support in the Workers runtime. Backwards incompatible fixes to the runtime following this date will not affect this Worker.",
							Computed:    true,
						},
						"compatibility_flags": schema.SetAttribute{
							Description: "Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
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
						"handlers": schema.ListAttribute{
							Description: "The names of handlers exported as part of the default export.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"has_assets": schema.BoolAttribute{
							Description: "Whether a Worker contains assets.",
							Computed:    true,
						},
						"has_modules": schema.BoolAttribute{
							Description: "Whether a Worker contains modules.",
							Computed:    true,
						},
						"last_deployed_from": schema.StringAttribute{
							Description: "The client most recently used to deploy this Worker.",
							Computed:    true,
						},
						"logpush": schema.BoolAttribute{
							Description: "Whether Logpush is turned on for the Worker.",
							Computed:    true,
						},
						"migration_tag": schema.StringAttribute{
							Description: "The tag of the Durable Object migration that was most recently applied for this Worker.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the script was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"named_handlers": schema.ListNestedAttribute{
							Description: "Named exports, such as Durable Object class implementations and named entrypoints.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[WorkersScriptsNamedHandlersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"handlers": schema.ListAttribute{
										Description: "The names of handlers exported as part of the named export.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"name": schema.StringAttribute{
										Description: "The name of the export.",
										Computed:    true,
									},
								},
							},
						},
						"placement": schema.SingleNestedAttribute{
							Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersScriptsPlacementDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"last_analyzed_at": schema.StringAttribute{
									Description: "The last time the script was analyzed for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"mode": schema.StringAttribute{
									Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"smart\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("smart"),
									},
								},
								"status": schema.StringAttribute{
									Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"SUCCESS\", \"UNSUPPORTED_APPLICATION\", \"INSUFFICIENT_INVOCATIONS\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"SUCCESS",
											"UNSUPPORTED_APPLICATION",
											"INSUFFICIENT_INVOCATIONS",
										),
									},
								},
							},
						},
						"placement_mode": schema.StringAttribute{
							Description:        "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"smart\".",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("smart"),
							},
						},
						"placement_status": schema.StringAttribute{
							Description:        "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"SUCCESS\", \"UNSUPPORTED_APPLICATION\", \"INSUFFICIENT_INVOCATIONS\".",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"SUCCESS",
									"UNSUPPORTED_APPLICATION",
									"INSUFFICIENT_INVOCATIONS",
								),
							},
						},
						"tail_consumers": schema.SetNestedAttribute{
							Description: "List of Workers that will consume logs from the attached Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectSetType[WorkersScriptsTailConsumersDataSourceModel](ctx),
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
						"usage_model": schema.StringAttribute{
							Description: "Usage model for the Worker invocations.\nAvailable values: \"standard\", \"bundled\", \"unbound\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"standard",
									"bundled",
									"unbound",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *WorkersScriptsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersScriptsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
