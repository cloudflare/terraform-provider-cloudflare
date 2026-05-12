// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2DataCatalogResultEnvelope struct {
	Result R2DataCatalogModel `json:"result"`
}

type R2DataCatalogModel struct {
	ID                types.String                                                  `tfsdk:"id" json:"id,computed"`
	AccountID         types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	BucketName        types.String                                                  `tfsdk:"bucket_name" path:"bucket_name,required"`
	Bucket            types.String                                                  `tfsdk:"bucket" json:"bucket,computed"`
	CredentialStatus  types.String                                                  `tfsdk:"credential_status" json:"credential_status,computed"`
	Name              types.String                                                  `tfsdk:"name" json:"name,computed"`
	Status            types.String                                                  `tfsdk:"status" json:"status,computed"`
	MaintenanceConfig customfield.NestedObject[R2DataCatalogMaintenanceConfigModel] `tfsdk:"maintenance_config" json:"maintenance_config,computed"`
}

func (m R2DataCatalogModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2DataCatalogModel) MarshalJSONForUpdate(state R2DataCatalogModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2DataCatalogMaintenanceConfigModel struct {
	Compaction         customfield.NestedObject[R2DataCatalogMaintenanceConfigCompactionModel]         `tfsdk:"compaction" json:"compaction,computed"`
	SnapshotExpiration customfield.NestedObject[R2DataCatalogMaintenanceConfigSnapshotExpirationModel] `tfsdk:"snapshot_expiration" json:"snapshot_expiration,computed"`
}

type R2DataCatalogMaintenanceConfigCompactionModel struct {
	State        types.String `tfsdk:"state" json:"state,computed"`
	TargetSizeMB types.String `tfsdk:"target_size_mb" json:"target_size_mb,computed"`
}

type R2DataCatalogMaintenanceConfigSnapshotExpirationModel struct {
	MaxSnapshotAge     types.String `tfsdk:"max_snapshot_age" json:"max_snapshot_age,computed"`
	MinSnapshotsToKeep types.Int64  `tfsdk:"min_snapshots_to_keep" json:"min_snapshots_to_keep,computed"`
	State              types.String `tfsdk:"state" json:"state,computed"`
}
