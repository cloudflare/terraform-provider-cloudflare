// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"etag": schema.StringAttribute{
				Description: "Hashed script content, can be used in a If-None-Match header when updating.",
				Optional:    true,
			},
			"has_assets": schema.BoolAttribute{
				Description: "Whether a Worker contains assets.",
				Optional:    true,
			},
			"has_modules": schema.BoolAttribute{
				Description: "Whether a Worker contains modules.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The id of the script in the Workers system. Usually the script name.",
				Optional:    true,
			},
			"logpush": schema.BoolAttribute{
				Description: "Whether Logpush is turned on for the Worker.",
				Optional:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the script was last modified.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"placement_mode": schema.StringAttribute{
				Description: "Specifies the placement mode for the Worker (e.g. 'smart').",
				Optional:    true,
			},
			"usage_model": schema.StringAttribute{
				Description: "Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').",
				Optional:    true,
			},
			"tail_consumers": schema.ListNestedAttribute{
				Description: "List of Workers that will consume logs from the attached Worker.",
				Optional:    true,
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
