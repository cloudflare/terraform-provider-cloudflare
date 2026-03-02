// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*R2DataCatalogDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Specifies the R2 bucket name.",
				Computed:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Specifies the R2 bucket name.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Use this to identify the account.",
				Required:    true,
			},
			"bucket": schema.StringAttribute{
				Description: "Specifies the associated R2 bucket name.",
				Computed:    true,
			},
			"credential_status": schema.StringAttribute{
				Description: "Shows the credential configuration status.\nAvailable values: \"present\", \"absent\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("present", "absent"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Specifies the catalog name (generated from account and bucket name).",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Indicates the status of the catalog.\nAvailable values: \"active\", \"inactive\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "inactive"),
				},
			},
			"maintenance_config": schema.SingleNestedAttribute{
				Description: "Configures maintenance for the catalog.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[R2DataCatalogMaintenanceConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"compaction": schema.SingleNestedAttribute{
						Description: "Configures compaction for catalog maintenance.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[R2DataCatalogMaintenanceConfigCompactionDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"state": schema.StringAttribute{
								Description: "Specifies the state of maintenance operations.\nAvailable values: \"enabled\", \"disabled\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
								},
							},
							"target_size_mb": schema.StringAttribute{
								Description: "Sets the target file size for compaction in megabytes. Defaults to \"128\".\nAvailable values: \"64\", \"128\", \"256\", \"512\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"64",
										"128",
										"256",
										"512",
									),
								},
							},
						},
					},
					"snapshot_expiration": schema.SingleNestedAttribute{
						Description: "Configures snapshot expiration settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[R2DataCatalogMaintenanceConfigSnapshotExpirationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"max_snapshot_age": schema.StringAttribute{
								Description: "Specifies the maximum age for snapshots. The system deletes snapshots older than this age.\nFormat: <number><unit> where unit is d (days), h (hours), m (minutes), or s (seconds).\nExamples: \"7d\" (7 days), \"48h\" (48 hours), \"2880m\" (2,880 minutes).\nDefaults to \"7d\".",
								Computed:    true,
							},
							"min_snapshots_to_keep": schema.Int64Attribute{
								Description: "Specifies the minimum number of snapshots to retain. Defaults to 100.",
								Computed:    true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
							"state": schema.StringAttribute{
								Description: "Specifies the state of maintenance operations.\nAvailable values: \"enabled\", \"disabled\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *R2DataCatalogDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2DataCatalogDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
