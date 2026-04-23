package v500

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func stringValue(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

func boolValue(b bool) types.Bool {
	return types.BoolValue(b)
}

// rawStateJSON is a minimal struct for JSON-unmarshalling the raw state to detect
// whether it looks like a v5 state (has zone_id) or a v4 state (needs transformation).
type rawStateJSON struct {
	ID                     string           `json:"id"`
	ZoneID                 string           `json:"zone_id"`
	ManagedRequestHeaders  json.RawMessage  `json:"managed_request_headers"`
	ManagedResponseHeaders json.RawMessage  `json:"managed_response_headers"`
}

// rawHeaderEntry is a minimal struct for JSON-unmarshalling header entries from raw state.
type rawHeaderEntry struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
}

// parseRawState converts raw JSON state bytes into a SourceManagedHeadersModel.
// This is used when PriorSchema is nil and req.State is unavailable.
func parseRawState(rawJSON []byte) (*SourceManagedHeadersModel, error) {
	var raw rawStateJSON
	if err := json.Unmarshal(rawJSON, &raw); err != nil {
		return nil, err
	}

	state := &SourceManagedHeadersModel{
		ID:     stringValue(raw.ID),
		ZoneID: stringValue(raw.ZoneID),
	}

	var reqEntries []rawHeaderEntry
	if len(raw.ManagedRequestHeaders) > 0 {
		if err := json.Unmarshal(raw.ManagedRequestHeaders, &reqEntries); err != nil {
			return nil, err
		}
	}
	reqHeaders := make([]*SourceHeaderEntryModel, 0, len(reqEntries))
	for _, e := range reqEntries {
		reqHeaders = append(reqHeaders, &SourceHeaderEntryModel{
			ID:      stringValue(e.ID),
			Enabled: boolValue(e.Enabled),
		})
	}
	state.ManagedRequestHeaders = &reqHeaders

	var respEntries []rawHeaderEntry
	if len(raw.ManagedResponseHeaders) > 0 {
		if err := json.Unmarshal(raw.ManagedResponseHeaders, &respEntries); err != nil {
			return nil, err
		}
	}
	respHeaders := make([]*SourceHeaderEntryModel, 0, len(respEntries))
	for _, e := range respEntries {
		respHeaders = append(respHeaders, &SourceHeaderEntryModel{
			ID:      stringValue(e.ID),
			Enabled: boolValue(e.Enabled),
		})
	}
	state.ManagedResponseHeaders = &respHeaders

	return state, nil
}

// UpgradeFromV0 handles state upgrades from schema_version=0 to v5 (version=500).
//
// There are two sources of schema_version=0 state:
//
//  1. State from v4 cloudflare_managed_headers moved via `terraform state mv`.
//     This state has the v4 schema (optional set nested blocks).
//
//  2. State from early v5 cloudflare_managed_transforms (versions 5.0.0–5.x.y before
//     the version was bumped to 500). These already have the correct v5 structure.
//
// Because PriorSchema is nil for this upgrader (to accept both formats), req.State
// is nil. We must use req.RawState.JSON to parse the prior state ourselves.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading managed_transforms state from schema_version=0")

	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		resp.Diagnostics.AddError(
			"unable to upgrade managed_transforms state",
			"raw state is empty or nil",
		)
		return
	}

	sourceState, err := parseRawState(req.RawState.JSON)
	if err != nil {
		resp.Diagnostics.AddError(
			"unable to upgrade managed_transforms state",
			"failed to parse raw state JSON: "+err.Error(),
		)
		return
	}

	if !sourceState.ZoneID.IsNull() {
		tflog.Info(ctx, "Upgrading managed_transforms state: detected early-v5 structure, transforming")
	} else {
		tflog.Info(ctx, "Upgrading managed_transforms state: detected v4 managed_headers structure, transforming")
	}

	newV5State, diags := Transform(ctx, *sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newV5State)...)
	tflog.Info(ctx, "State upgrade from schema_version=0 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade. Version 1 is the current v5 schema version.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading managed_transforms state from version=1 to version=500 (no-op)")

	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
