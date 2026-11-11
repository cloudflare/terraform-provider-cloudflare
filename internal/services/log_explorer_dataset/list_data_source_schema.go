package log_explorer_dataset

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*LogExplorerDatasetsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Logs Read",
				"Logs Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"include_zones": schema.BoolAttribute{
				Description: "Set to true to include zone-scoped datasets belonging to this account. Only supported with account_id.",
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
				CustomType:  customfield.NewNestedObjectListType[LogExplorerDatasetsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique dataset ID.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "RFC3339 timestamp recording when the API created this dataset.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dataset": schema.StringAttribute{
							Description: "Dataset type name (e.g. `http_requests`).",
							Computed:    true,
						},
						"dataset_id": schema.StringAttribute{
							Description: "Unique dataset ID.",
							Computed:    true,
						},
						"deletion_protection": schema.BoolAttribute{
							Description: "Whether deletion is blocked. Set to `false` before deleting the dataset.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether log ingest is currently active for this dataset.",
							Computed:    true,
						},
						"object_id": schema.StringAttribute{
							Description: "Public ID of the account or zone that owns this dataset.",
							Computed:    true,
						},
						"object_type": schema.StringAttribute{
							Description: "Whether this dataset belongs to an account or a zone.\nAvailable values: \"account\", \"zone\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("account", "zone"),
							},
						},
						"updated_at": schema.StringAttribute{
							Description: "RFC3339 timestamp recording when the API last updated this dataset.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *LogExplorerDatasetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *LogExplorerDatasetsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.All(
			datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
			datasourcevalidator.Conflicting(path.MatchRoot("zone_id"), path.MatchRoot("include_zones")),
		),
	}
}
