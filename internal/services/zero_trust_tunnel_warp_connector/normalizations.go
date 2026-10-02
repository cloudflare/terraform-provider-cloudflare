package zero_trust_tunnel_warp_connector

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// warpConnectorMetadataHA is a minimal struct used to extract the ha field from
// the opaque metadata JSON object returned by the Cloudflare API.
type warpConnectorMetadataHA struct {
	HA *bool `json:"ha"`
}

// normalizeWARPConnectorHA populates data.Ha from data.Metadata.
//
// The Cloudflare WARP Connector API accepts ha on Create but never returns it
// as a top-level field in Create/Get/Update responses. Instead, ha is embedded
// in the opaque metadata object (e.g. {"ha": true}). The model tag
// `no_refresh` prevents the JSON decoder from trying to read a nonexistent
// top-level ha, so after Import data.Ha is null even though the connector
// may be HA.
//
// This function reads metadata.ha and writes it to data.Ha:
//   - If data.Ha is already a concrete value, it is left unchanged.
//   - If metadata is absent, null, unknown, or unparseable, the function
//     returns early without modifying data.Ha.
//   - If metadata is valid JSON but contains no "ha" key, data.Ha is set to
//     false (the default for non-HA connectors).
//   - Otherwise data.Ha is set to the value of metadata.ha.
func normalizeWARPConnectorHA(data *ZeroTrustTunnelWARPConnectorModel) {
	if data == nil {
		return
	}

	// If ha is already set to a concrete value, respect it.
	if !data.Ha.IsNull() && !data.Ha.IsUnknown() {
		return
	}

	// Metadata absent or not yet known — cannot determine HA status.
	if data.Metadata.IsNull() || data.Metadata.IsUnknown() {
		return
	}

	var meta warpConnectorMetadataHA
	if err := json.Unmarshal([]byte(data.Metadata.ValueString()), &meta); err != nil {
		// Unparseable metadata — leave ha unchanged.
		return
	}

	// Metadata present but ha key absent → non-HA connector (default false).
	ha := false
	if meta.HA != nil {
		ha = *meta.HA
	}
	data.Ha = types.BoolValue(ha)
}
