package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

// SourceCloudflareRecordModel represents the source cloudflare_record state structure.
// This corresponds to schema_version=3 from the legacy (SDKv2) cloudflare provider.
// Used by both MoveState (Terraform 1.8+) and UpgradeFromLegacyV3 (Terraform < 1.8) to parse legacy state.
type SourceCloudflareRecordModel struct {
	ID             types.String      `tfsdk:"id"`
	ZoneID         types.String      `tfsdk:"zone_id"`
	Name           types.String      `tfsdk:"name"`
	Type           types.String      `tfsdk:"type"`
	Value          types.String      `tfsdk:"value"`
	Content        types.String      `tfsdk:"content"`
	TTL            types.Int64       `tfsdk:"ttl"`
	Priority       types.Int64       `tfsdk:"priority"`
	Proxied        types.Bool        `tfsdk:"proxied"`
	Comment        types.String      `tfsdk:"comment"`
	Tags           types.Set         `tfsdk:"tags"`
	AllowOverwrite types.Bool        `tfsdk:"allow_overwrite"`
	Hostname       types.String      `tfsdk:"hostname"`
	Proxiable      types.Bool        `tfsdk:"proxiable"`
	CreatedOn      types.String      `tfsdk:"created_on"`
	ModifiedOn     types.String      `tfsdk:"modified_on"`
	Metadata       types.Map         `tfsdk:"metadata"`
	Data           []SourceDataModel `tfsdk:"data"`
}

// SourceDataModel represents the source data block structure.
// In the legacy provider, this is a list with MaxItems: 1.
type SourceDataModel struct {
	// CAA fields
	Flags   types.String `tfsdk:"flags"`
	Tag     types.String `tfsdk:"tag"`
	Content types.String `tfsdk:"content"` // CAA uses content, renamed to value in v500

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

// TargetDNSRecordModel represents the target cloudflare_dns_record state structure (v500).
type TargetDNSRecordModel struct {
	ID                types.String                  `tfsdk:"id"`
	ZoneID            types.String                  `tfsdk:"zone_id"`
	Name              types.String                  `tfsdk:"name"`
	Type              types.String                  `tfsdk:"type"`
	Content           types.String                  `tfsdk:"content"`
	TTL               types.Float64                 `tfsdk:"ttl"`
	Priority          types.Float64                 `tfsdk:"priority"`
	Proxied           types.Bool                    `tfsdk:"proxied"`
	Comment           types.String                  `tfsdk:"comment"`
	Tags              customfield.Set[types.String] `tfsdk:"tags"`
	Data              *TargetDNSRecordDataModel     `tfsdk:"data"`
	CreatedOn         timetypes.RFC3339             `tfsdk:"created_on"`
	ModifiedOn        timetypes.RFC3339             `tfsdk:"modified_on"`
	CommentModifiedOn timetypes.RFC3339             `tfsdk:"comment_modified_on"`
	TagsModifiedOn    timetypes.RFC3339             `tfsdk:"tags_modified_on"`
	// Computed fields (not migrated, will be refreshed from API)
	Proxiable types.Bool                                             `tfsdk:"proxiable"`
	Meta      jsontypes.Normalized                                   `tfsdk:"meta"`
	Settings  customfield.NestedObject[TargetDNSRecordSettingsModel] `tfsdk:"settings"`
}

// TargetDNSRecordDataModel represents the target data nested object (v500).
type TargetDNSRecordDataModel struct {
	Flags         customfield.NormalizedDynamicValue `tfsdk:"flags"`
	Tag           types.String                       `tfsdk:"tag"`
	Value         types.String                       `tfsdk:"value"`
	Algorithm     types.Float64                      `tfsdk:"algorithm"`
	Altitude      types.Float64                      `tfsdk:"altitude"`
	Certificate   types.String                       `tfsdk:"certificate"`
	Digest        types.String                       `tfsdk:"digest"`
	DigestType    types.Float64                      `tfsdk:"digest_type"`
	Fingerprint   types.String                       `tfsdk:"fingerprint"`
	KeyTag        types.Float64                      `tfsdk:"key_tag"`
	LatDegrees    types.Float64                      `tfsdk:"lat_degrees"`
	LatDirection  types.String                       `tfsdk:"lat_direction"`
	LatMinutes    types.Float64                      `tfsdk:"lat_minutes"`
	LatSeconds    types.Float64                      `tfsdk:"lat_seconds"`
	LongDegrees   types.Float64                      `tfsdk:"long_degrees"`
	LongDirection types.String                       `tfsdk:"long_direction"`
	LongMinutes   types.Float64                      `tfsdk:"long_minutes"`
	LongSeconds   types.Float64                      `tfsdk:"long_seconds"`
	MatchingType  types.Float64                      `tfsdk:"matching_type"`
	Order         types.Float64                      `tfsdk:"order"`
	Port          types.Float64                      `tfsdk:"port"`
	PrecisionHorz types.Float64                      `tfsdk:"precision_horz"`
	PrecisionVert types.Float64                      `tfsdk:"precision_vert"`
	Preference    types.Float64                      `tfsdk:"preference"`
	Priority      types.Float64                      `tfsdk:"priority"`
	Protocol      types.Float64                      `tfsdk:"protocol"`
	PublicKey     types.String                       `tfsdk:"public_key"`
	Regex         types.String                       `tfsdk:"regex"`
	Replacement   types.String                       `tfsdk:"replacement"`
	Selector      types.Float64                      `tfsdk:"selector"`
	Service       types.String                       `tfsdk:"service"`
	Size          types.Float64                      `tfsdk:"size"`
	Target        types.String                       `tfsdk:"target"`
	Type          types.Float64                      `tfsdk:"type"`
	Usage         types.Float64                      `tfsdk:"usage"`
	Weight        types.Float64                      `tfsdk:"weight"`
}

// TargetDNSRecordSettingsModel represents the target settings nested object (v500).
// These are computed/optional fields that control DNS record behavior.
// Must match dns_record.DNSRecordSettingsModel structure exactly.
type TargetDNSRecordSettingsModel struct {
	IPV4Only     types.Bool `tfsdk:"ipv4_only" json:"ipv4_only,computed_optional"`
	IPV6Only     types.Bool `tfsdk:"ipv6_only" json:"ipv6_only,computed_optional"`
	FlattenCNAME types.Bool `tfsdk:"flatten_cname" json:"flatten_cname,computed_optional"`
}
