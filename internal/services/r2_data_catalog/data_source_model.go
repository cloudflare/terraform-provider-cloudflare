// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/r2_data_catalog"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2DataCatalogResultDataSourceEnvelope struct {
	Result R2DataCatalogDataSourceModel `json:"result,computed"`
}

type R2DataCatalogDataSourceModel struct {
	ID                types.String                                                            `tfsdk:"id" path:"bucket_name,computed"`
	BucketName        types.String                                                            `tfsdk:"bucket_name" path:"bucket_name,required"`
	AccountID         types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	Bucket            types.String                                                            `tfsdk:"bucket" json:"bucket,computed"`
	CredentialStatus  types.String                                                            `tfsdk:"credential_status" json:"credential_status,computed"`
	Name              types.String                                                            `tfsdk:"name" json:"name,computed"`
	Status            types.String                                                            `tfsdk:"status" json:"status,computed"`
	MaintenanceConfig customfield.NestedObject[R2DataCatalogMaintenanceConfigDataSourceModel] `tfsdk:"maintenance_config" json:"maintenance_config,computed"`
}

func (m *R2DataCatalogDataSourceModel) toReadParams(_ context.Context) (params r2_data_catalog.R2DataCatalogGetParams, diags diag.Diagnostics) {
	params = r2_data_catalog.R2DataCatalogGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2DataCatalogMaintenanceConfigDataSourceModel struct {
	Compaction         customfield.NestedObject[R2DataCatalogMaintenanceConfigCompactionDataSourceModel]         `tfsdk:"compaction" json:"compaction,computed"`
	SnapshotExpiration customfield.NestedObject[R2DataCatalogMaintenanceConfigSnapshotExpirationDataSourceModel] `tfsdk:"snapshot_expiration" json:"snapshot_expiration,computed"`
}

type R2DataCatalogMaintenanceConfigCompactionDataSourceModel struct {
	State        types.String `tfsdk:"state" json:"state,computed"`
	TargetSizeMB types.String `tfsdk:"target_size_mb" json:"target_size_mb,computed"`
}

type R2DataCatalogMaintenanceConfigSnapshotExpirationDataSourceModel struct {
	MaxSnapshotAge     types.String `tfsdk:"max_snapshot_age" json:"max_snapshot_age,computed"`
	MinSnapshotsToKeep types.Int64  `tfsdk:"min_snapshots_to_keep" json:"min_snapshots_to_keep,computed"`
	State              types.String `tfsdk:"state" json:"state,computed"`
}
