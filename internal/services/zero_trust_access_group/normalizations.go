package zero_trust_access_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type IsNull interface {
	IsNull() bool
}

func normalizeEmptyAndNullSlice[T any](data **[]T, stateData *[]T) {
	if (*data != nil && len(**data) != 0) || (stateData != nil && len(*stateData) != 0) {
		return
	}
	*data = stateData
}

// Normalizing function to ensure consistency between the state/plan and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustAccessGroupAPIData(ctx context.Context, data, sourceData *ZeroTrustAccessGroupModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	normalizeEmptyAndNullSlice(&data.Include, sourceData.Include)
	normalizeEmptyAndNullSlice(&data.Require, sourceData.Require)
	normalizeEmptyAndNullSlice(&data.Exclude, sourceData.Exclude)

	// Normalize IP addresses in include/exclude/require rules to handle /32 and /128 CIDR notation
	if data.Include != nil && sourceData.Include != nil {
		for i := range *data.Include {
			if i < len(*sourceData.Include) {
				if (*data.Include)[i].IP != nil && (*sourceData.Include)[i].IP != nil {
					utils.NormalizeIPStringWithCIDR(&(*data.Include)[i].IP.IP, (*sourceData.Include)[i].IP.IP)
				}
			}
		}
	}
	if data.Exclude != nil && sourceData.Exclude != nil {
		for i := range *data.Exclude {
			if i < len(*sourceData.Exclude) {
				if (*data.Exclude)[i].IP != nil && (*sourceData.Exclude)[i].IP != nil {
					utils.NormalizeIPStringWithCIDR(&(*data.Exclude)[i].IP.IP, (*sourceData.Exclude)[i].IP.IP)
				}
			}
		}
	}
	if data.Require != nil && sourceData.Require != nil {
		for i := range *data.Require {
			if i < len(*sourceData.Require) {
				if (*data.Require)[i].IP != nil && (*sourceData.Require)[i].IP != nil {
					utils.NormalizeIPStringWithCIDR(&(*data.Require)[i].IP.IP, (*sourceData.Require)[i].IP.IP)
				}
			}
		}
	}

	return diags
}

func normalizeImportZeroTrustAccessGroupAPIData(ctx context.Context, data *ZeroTrustAccessGroupModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	if data.Include != nil && len(*data.Include) == 0 {
		data.Include = nil
	}

	if data.Require != nil && len(*data.Require) == 0 {
		data.Require = nil
	}

	if data.Exclude != nil && len(*data.Exclude) == 0 {
		data.Exclude = nil
	}

	return diags
}
