package cloudflare_record

import (
	"fmt"
	"strconv"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v4_models"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v5_models"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DNSRecordTransformer handles transformation from v4 to v5 DNS record models
type DNSRecordTransformer struct {
	// Configuration options
	AddDefaultTTL    bool
	RemoveDeprecated bool
	PreserveComments bool
}

// NewDNSRecordTransformer creates a new transformer with default settings
func NewDNSRecordTransformer() *DNSRecordTransformer {
	return &DNSRecordTransformer{
		AddDefaultTTL:    true,
		RemoveDeprecated: true,
		PreserveComments: false,
	}
}

// Transform implements the ModelTransformer interface
func (t *DNSRecordTransformer) Transform(v4Model interface{}) (interface{}, error) {
	v4, ok := v4Model.(*v4_models.DNSRecordV4Model)
	if !ok {
		return nil, fmt.Errorf("expected *v4_models.DNSRecordV4Model, got %T", v4Model)
	}
	return t.TransformDNSRecord(*v4)
}

// TransformDNSRecord converts a v4 DNS record to v5
func (t *DNSRecordTransformer) TransformDNSRecord(v4 v4_models.DNSRecordV4Model) (*v5_models.DNSRecordV5Model, error) {
	v5 := &v5_models.DNSRecordV5Model{
		ID:      v4.ID,
		ZoneID:  v4.ZoneID,
		Name:    v4.Name,
		Type:    v4.Type,
		Proxied: v4.Proxied,
	}

	// Handle primary value field transformation
	// v4: could be 'value' or 'content' -> v5: always 'content'
	primaryValue := v4.GetPrimaryValue()
	if !primaryValue.IsNull() && !primaryValue.IsUnknown() {
		v5.Content = primaryValue
	}

	// Handle TTL
	if !v4.TTL.IsNull() && !v4.TTL.IsUnknown() {
		v5.TTL = v4.TTL
	} else if t.AddDefaultTTL {
		v5.TTL = types.Float64Value(1) // Automatic TTL
	}

	// Handle priority (for MX, SRV, URI records)
	if !v4.Priority.IsNull() && !v4.Priority.IsUnknown() {
		v5.Priority = v4.Priority
	}

	// Transform data block if present
	if v4.HasData() {
		v5Data, err := t.transformDataBlock(v4.Data, v4.Type.ValueString())
		if err != nil {
			return nil, err
		}
		v5.Data = v5Data

		// For SRV records, priority might be in the data block
		if v5Data != nil && !v5Data.Priority.IsNull() && v5.Priority.IsNull() {
			v5.Priority = v5Data.Priority
		}
	}

	// Apply defaults only if we're adding defaults
	if t.AddDefaultTTL {
		v5.SetDefaults()
	} else {
		// Only set proxied default if it's null (not TTL)
		if v5.Proxied.IsNull() {
			v5.Proxied = types.BoolValue(false)
		}
	}

	// Note: Deprecated fields (allow_overwrite, hostname) are intentionally not copied

	return v5, nil
}

// transformDataBlock converts v4 data array to v5 data object
func (t *DNSRecordTransformer) transformDataBlock(v4Data []v4_models.DNSRecordDataV4Model, recordType string) (*v5_models.DNSRecordDataV5Model, error) {
	if len(v4Data) == 0 {
		return nil, nil
	}

	// Take the first element (v4 used array, v5 uses single object)
	first := v4Data[0]
	v5Data := &v5_models.DNSRecordDataV5Model{}

	// Transform based on record type
	switch recordType {
	case "CAA":
		t.transformCAAData(&first, v5Data)
	case "SRV":
		t.transformSRVData(&first, v5Data)
	case "LOC":
		t.transformLOCData(&first, v5Data)
	case "URI":
		t.transformURIData(&first, v5Data)
	default:
		// Generic transformation for other types
		t.transformGenericData(&first, v5Data)
	}

	return v5Data, nil
}

// transformCAAData handles CAA-specific field transformations
func (t *DNSRecordTransformer) transformCAAData(v4 *v4_models.DNSRecordDataV4Model, v5 *v5_models.DNSRecordDataV5Model) {
	// Transform flags from string to number
	if !v4.Flags.IsNull() && !v4.Flags.IsUnknown() {
		flagStr := v4.Flags.ValueString()
		if flagNum, err := strconv.ParseFloat(flagStr, 64); err == nil {
			v5.Flags = types.Float64Value(flagNum)
		} else {
			// If not a number, default to 0
			v5.Flags = types.Float64Value(0)
		}
	}

	// Copy tag
	v5.Tag = v4.Tag

	// Transform content -> value
	if !v4.Content.IsNull() && !v4.Content.IsUnknown() {
		v5.Value = v4.Content
	}
}

// transformSRVData handles SRV-specific field transformations
func (t *DNSRecordTransformer) transformSRVData(v4 *v4_models.DNSRecordDataV4Model, v5 *v5_models.DNSRecordDataV5Model) {
	v5.Priority = v4.Priority
	v5.Weight = v4.Weight
	v5.Port = v4.Port
	v5.Target = v4.Target
	v5.Service = v4.Service

	// Note: 'proto' and 'name' fields are removed in v5
}

// transformLOCData handles LOC-specific field transformations
func (t *DNSRecordTransformer) transformLOCData(v4 *v4_models.DNSRecordDataV4Model, v5 *v5_models.DNSRecordDataV5Model) {
	// LOC fields are mostly unchanged
	v5.Altitude = v4.Altitude
	v5.LatDegrees = v4.LatDegrees
	v5.LatDirection = v4.LatDirection
	v5.LatMinutes = v4.LatMinutes
	v5.LatSeconds = v4.LatSeconds
	v5.LongDegrees = v4.LongDegrees
	v5.LongDirection = v4.LongDirection
	v5.LongMinutes = v4.LongMinutes
	v5.LongSeconds = v4.LongSeconds
	v5.PrecisionHorz = v4.PrecisionHorz
	v5.PrecisionVert = v4.PrecisionVert
	v5.Size = v4.Size
}

// transformURIData handles URI-specific field transformations
func (t *DNSRecordTransformer) transformURIData(v4 *v4_models.DNSRecordDataV4Model, v5 *v5_models.DNSRecordDataV5Model) {
	v5.Priority = v4.Priority
	v5.Weight = v4.Weight
	v5.Target = v4.Target
}

// transformGenericData handles generic field transformations
func (t *DNSRecordTransformer) transformGenericData(v4 *v4_models.DNSRecordDataV4Model, v5 *v5_models.DNSRecordDataV5Model) {
	// Copy common fields that might exist
	if !v4.Target.IsNull() {
		v5.Target = v4.Target
	}
	if !v4.Priority.IsNull() {
		v5.Priority = v4.Priority
	}
	if !v4.Weight.IsNull() {
		v5.Weight = v4.Weight
	}
	if !v4.Port.IsNull() {
		v5.Port = v4.Port
	}
}