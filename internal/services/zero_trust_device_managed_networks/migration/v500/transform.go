package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts a source cloudflare_device_managed_networks state to target cloudflare_zero_trust_device_managed_networks state.
func Transform(ctx context.Context, source SourceCloudflareDeviceManagedNetworksModel) (*TargetZeroTrustDeviceManagedNetworksModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError("Missing required field", "account_id is required for device managed networks migration")
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError("Missing required field", "name is required for device managed networks migration")
		return nil, diags
	}
	if source.Type.IsNull() || source.Type.IsUnknown() {
		diags.AddError("Missing required field", "type is required for device managed networks migration")
		return nil, diags
	}

	target := &TargetZeroTrustDeviceManagedNetworksModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,
		Type:      source.Type,
		// NetworkID: v4 stored the network ID in the id field, v5 has both id and network_id
		// Copy id to network_id (they should have the same value)
		NetworkID: source.ID,
	}

	// Config: Transform from []SourceConfigModel (TypeList MaxItems:1 in v4) to *TargetConfigModel (SingleNestedAttribute in v5)
	// v4 stores config as a single-element array, v5 stores as an object
	if len(source.Config) > 0 {
		sourceConfig := source.Config[0]
		target.Config = &TargetConfigModel{
			TLSSockaddr: sourceConfig.TLSSockaddr,
			Sha256:      sourceConfig.Sha256,
		}
	} else {
		// If config is missing, this is an error since it's required in both v4 and v5
		diags.AddError("Missing required field", "config is required for device managed networks migration")
		return nil, diags
	}

	return target, diags
}
