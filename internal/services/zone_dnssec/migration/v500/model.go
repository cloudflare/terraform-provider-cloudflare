// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareZoneDNSSECModel represents the v4 (SDKv2) state structure
// This model is used to read the old state format before migration
type SourceCloudflareZoneDNSSECModel struct {
	ID              types.String `tfsdk:"id"`
	ZoneID          types.String `tfsdk:"zone_id"`
	Status          types.String `tfsdk:"status"`
	Flags           types.Int64  `tfsdk:"flags"`            // Int64 in v4
	Algorithm       types.String `tfsdk:"algorithm"`
	KeyType         types.String `tfsdk:"key_type"`
	DigestType      types.String `tfsdk:"digest_type"`
	DigestAlgorithm types.String `tfsdk:"digest_algorithm"`
	Digest          types.String `tfsdk:"digest"`
	DS              types.String `tfsdk:"ds"`
	KeyTag          types.Int64  `tfsdk:"key_tag"`    // Int64 in v4
	PublicKey       types.String `tfsdk:"public_key"`
	ModifiedOn      types.String `tfsdk:"modified_on"` // String in RFC1123Z format in v4
}

// TargetZoneDNSSECModel represents the v5 (Plugin Framework) state structure
// This model is used to write the new state format after migration
type TargetZoneDNSSECModel struct {
	ID                types.String  `tfsdk:"id"`
	ZoneID            types.String  `tfsdk:"zone_id"`
	DNSSECMultiSigner types.Bool    `tfsdk:"dnssec_multi_signer"` // New in v5
	DNSSECPresigned   types.Bool    `tfsdk:"dnssec_presigned"`    // New in v5
	DNSSECUseNsec3    types.Bool    `tfsdk:"dnssec_use_nsec3"`    // New in v5
	Status            types.String  `tfsdk:"status"`
	Algorithm         types.String  `tfsdk:"algorithm"`
	Digest            types.String  `tfsdk:"digest"`
	DigestAlgorithm   types.String  `tfsdk:"digest_algorithm"`
	DigestType        types.String  `tfsdk:"digest_type"`
	DS                types.String  `tfsdk:"ds"`
	Flags             types.Float64 `tfsdk:"flags"`     // Float64 in v5
	KeyTag            types.Float64 `tfsdk:"key_tag"`   // Float64 in v5
	KeyType           types.String  `tfsdk:"key_type"`
	ModifiedOn        types.String  `tfsdk:"modified_on"` // String in RFC3339 format in v5
	PublicKey         types.String  `tfsdk:"public_key"`
}
