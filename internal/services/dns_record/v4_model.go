package dns_record

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// V4CloudflareRecordModel represents the v4 cloudflare_record state structure.
// This is used by MoveState to parse the source state from v4 provider.
type V4CloudflareRecordModel struct {
	ID             types.String   `tfsdk:"id"`
	ZoneID         types.String   `tfsdk:"zone_id"`
	Name           types.String   `tfsdk:"name"`
	Type           types.String   `tfsdk:"type"`
	Value          types.String   `tfsdk:"value"`
	Content        types.String   `tfsdk:"content"`
	TTL            types.Int64    `tfsdk:"ttl"`
	Priority       types.Int64    `tfsdk:"priority"`
	Proxied        types.Bool     `tfsdk:"proxied"`
	Comment        types.String   `tfsdk:"comment"`
	Tags           types.Set      `tfsdk:"tags"`
	AllowOverwrite types.Bool     `tfsdk:"allow_overwrite"`
	Hostname       types.String   `tfsdk:"hostname"`
	Proxiable      types.Bool     `tfsdk:"proxiable"`
	CreatedOn      types.String   `tfsdk:"created_on"`
	ModifiedOn     types.String   `tfsdk:"modified_on"`
	Metadata       types.Map      `tfsdk:"metadata"`
	Data           []V4DataModel  `tfsdk:"data"`
}

// V4DataModel represents the v4 data block structure.
// In v4, this is a list with MaxItems: 1.
type V4DataModel struct {
	// CAA fields
	Flags   types.String `tfsdk:"flags"`
	Tag     types.String `tfsdk:"tag"`
	Content types.String `tfsdk:"content"` // CAA uses content, renamed to value in v5

	// SRV fields
	Service  types.String `tfsdk:"service"`
	Proto    types.String `tfsdk:"proto"`
	Name     types.String `tfsdk:"name"`
	Priority types.Int64  `tfsdk:"priority"`
	Weight   types.Int64  `tfsdk:"weight"`
	Port     types.Int64  `tfsdk:"port"`
	Target   types.String `tfsdk:"target"`

	// DNSKEY/DS/CERT fields
	Algorithm   types.Int64  `tfsdk:"algorithm"`
	KeyTag      types.Int64  `tfsdk:"key_tag"`
	Type        types.Int64  `tfsdk:"type"`
	Protocol    types.Int64  `tfsdk:"protocol"`
	PublicKey   types.String `tfsdk:"public_key"`
	Digest      types.String `tfsdk:"digest"`
	DigestType  types.Int64  `tfsdk:"digest_type"`
	Certificate types.String `tfsdk:"certificate"`

	// TLSA fields
	Usage        types.Int64 `tfsdk:"usage"`
	Selector     types.Int64 `tfsdk:"selector"`
	MatchingType types.Int64 `tfsdk:"matching_type"`

	// LOC fields
	Altitude      types.Float64 `tfsdk:"altitude"`
	LatDegrees    types.Int64   `tfsdk:"lat_degrees"`
	LatDirection  types.String  `tfsdk:"lat_direction"`
	LatMinutes    types.Int64   `tfsdk:"lat_minutes"`
	LatSeconds    types.Float64 `tfsdk:"lat_seconds"`
	LongDegrees   types.Int64   `tfsdk:"long_degrees"`
	LongDirection types.String  `tfsdk:"long_direction"`
	LongMinutes   types.Int64   `tfsdk:"long_minutes"`
	LongSeconds   types.Float64 `tfsdk:"long_seconds"`
	PrecisionHorz types.Float64 `tfsdk:"precision_horz"`
	PrecisionVert types.Float64 `tfsdk:"precision_vert"`
	Size          types.Float64 `tfsdk:"size"`

	// NAPTR fields
	Order       types.Int64  `tfsdk:"order"`
	Preference  types.Int64  `tfsdk:"preference"`
	Regex       types.String `tfsdk:"regex"`
	Replacement types.String `tfsdk:"replacement"`

	// SSHFP fields
	Fingerprint types.String `tfsdk:"fingerprint"`

	// URI fields
	Value types.String `tfsdk:"value"`
}
