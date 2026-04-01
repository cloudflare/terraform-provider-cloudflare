package origin_cloud_regions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// OriginCloudRegionsResultEnvelope wraps the standard Cloudflare v4 API envelope.
type OriginCloudRegionsResultEnvelope struct {
	Result OriginCloudRegionsResult `json:"result"`
}

// OriginCloudRegionsResult is the top-level result object returned by
// GET /zones/:zone_id/cache/origin_public_cloud_region.
type OriginCloudRegionsResult struct {
	ID         string                   `json:"id"`
	Value      []OriginCloudRegionEntry `json:"value"`
	ModifiedOn string                   `json:"modified_on"`
	Editable   bool                     `json:"editable"`
}

// OriginCloudRegionEntry is one IP→vendor:region mapping in the value array.
type OriginCloudRegionEntry struct {
	OriginIP   string `json:"origin-ip"`
	Vendor     string `json:"vendor"`
	Region     string `json:"region"`
	ModifiedOn string `json:"modified_on,omitempty"`
}

// OriginCloudRegionMappingModel is the Terraform model for a single mapping block.
type OriginCloudRegionMappingModel struct {
	OriginIP types.String `tfsdk:"origin_ip"`
	Vendor   types.String `tfsdk:"vendor"`
	Region   types.String `tfsdk:"region"`
}

// OriginCloudRegionsModel is the Terraform resource model.
type OriginCloudRegionsModel struct {
	ID       types.String                    `tfsdk:"id"`
	ZoneID   types.String                    `tfsdk:"zone_id"`
	Mappings []OriginCloudRegionMappingModel `tfsdk:"mappings"`
}

// OriginCloudRegionUpsertRequest is the POST/PATCH request body.
type OriginCloudRegionUpsertRequest struct {
	IP     string `json:"ip"`
	Vendor string `json:"vendor"`
	Region string `json:"region"`
}
