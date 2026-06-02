package v500

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UpgradeFromV0 converts schema_version=0 → 500: id changes Int64 → String
// (decimal representation of the prior int).
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var prior SourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	next := TargetModel{
		AccountID:               prior.AccountID,
		Email:                   prior.Email,
		IsEmailRegex:            prior.IsEmailRegex,
		Name:                    prior.Name,
		Comments:                prior.Comments,
		DirectoryID:             prior.DirectoryID,
		DirectoryNodeID:         prior.DirectoryNodeID,
		ExternalDirectoryNodeID: prior.ExternalDirectoryNodeID,
		Provenance:              prior.Provenance,
		CreatedAt:               prior.CreatedAt,
		LastModified:            prior.LastModified,
	}
	if !prior.ID.IsNull() && !prior.ID.IsUnknown() {
		next.ID = types.StringValue(strconv.FormatInt(prior.ID.ValueInt64(), 10))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &next)...)
}
