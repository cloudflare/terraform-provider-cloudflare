package cloudflare_record

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/generator"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v5_models"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// DNSRecordGenerator implements ModelGenerator for DNS records
// This contains all DNS-specific HCL generation logic
type DNSRecordGenerator struct {
	*generator.GenericHCLGenerator
}

// NewDNSRecordGenerator creates a new DNS record generator
func NewDNSRecordGenerator() *DNSRecordGenerator {
	return &DNSRecordGenerator{
		GenericHCLGenerator: generator.NewGenericHCLGenerator(),
	}
}

// Generate creates an HCL block from a v5 DNS record model
// This method knows how to generate DNS-specific HCL structure
func (g *DNSRecordGenerator) Generate(v5Model interface{}, labels []string) *hclwrite.Block {
	model, ok := v5Model.(*v5_models.DNSRecordV5Model)
	if !ok {
		// Return empty block if type assertion fails
		return hclwrite.NewBlock("resource", labels)
	}

	// Create the resource block with updated resource type
	// v4: cloudflare_record -> v5: cloudflare_dns_record
	resourceLabels := []string{"cloudflare_dns_record"}
	if len(labels) > 1 {
		resourceLabels = append(resourceLabels, labels[1])
	}

	block := g.CreateBlock("resource", resourceLabels)
	body := block.Body()

	// Add DNS record specific attributes in logical order
	g.AddStringAttribute(body, "zone_id", model.ZoneID)
	g.AddStringAttribute(body, "name", model.Name)
	g.AddStringAttribute(body, "type", model.Type)

	// Add content (not value - that's v4)
	g.AddStringAttribute(body, "content", model.Content)

	// TTL is required in v5
	if !model.TTL.IsNull() {
		body.SetAttributeValue("ttl", cty.NumberFloatVal(model.TTL.ValueFloat64()))
	}

	// Add optional attributes
	g.AddBoolAttribute(body, "proxied", model.Proxied)
	g.AddFloat64Attribute(body, "priority", model.Priority)
	g.AddStringAttribute(body, "comment", model.Comment)

	// Add tags if present
	if len(model.Tags) > 0 {
		g.AddStringArrayAttribute(body, "tags", model.Tags)
	}

	// Add data block if present (DNS-specific structure)
	if model.Data != nil {
		g.generateDataBlock(body, model.Data, model.Type.ValueString())
	}

	return block
}

// generateDataBlock generates a DNS record data block
func (g *DNSRecordGenerator) generateDataBlock(body *hclwrite.Body, data *v5_models.DNSRecordDataV5Model, recordType string) {
	// Check if data block has any content
	if !g.hasDataContent(data) {
		return
	}

	dataBlock := g.AddBlock(body, "data", nil)
	dataBody := dataBlock.Body()

	// Add fields based on DNS record type
	switch recordType {
	case "CAA":
		g.generateCAAData(dataBody, data)
	case "SRV":
		g.generateSRVData(dataBody, data)
	case "LOC":
		g.generateLOCData(dataBody, data)
	case "URI":
		g.generateURIData(dataBody, data)
	default:
		g.generateGenericData(dataBody, data)
	}
}

// generateCAAData generates CAA-specific data fields
func (g *DNSRecordGenerator) generateCAAData(body *hclwrite.Body, data *v5_models.DNSRecordDataV5Model) {
	// Flags as number in v5
	if !data.Flags.IsNull() {
		body.SetAttributeValue("flags", cty.NumberFloatVal(data.Flags.ValueFloat64()))
	}

	g.AddStringAttribute(body, "tag", data.Tag)
	g.AddStringAttribute(body, "value", data.Value) // Note: v5 uses 'value', not 'content'
}

// generateSRVData generates SRV-specific data fields
func (g *DNSRecordGenerator) generateSRVData(body *hclwrite.Body, data *v5_models.DNSRecordDataV5Model) {
	g.AddFloat64Attribute(body, "priority", data.Priority)
	g.AddFloat64Attribute(body, "weight", data.Weight)
	g.AddFloat64Attribute(body, "port", data.Port)
	g.AddStringAttribute(body, "target", data.Target)
	g.AddStringAttribute(body, "service", data.Service)
}

// generateLOCData generates LOC-specific data fields
func (g *DNSRecordGenerator) generateLOCData(body *hclwrite.Body, data *v5_models.DNSRecordDataV5Model) {
	g.AddFloat64Attribute(body, "altitude", data.Altitude)
	g.AddFloat64Attribute(body, "lat_degrees", data.LatDegrees)
	g.AddStringAttribute(body, "lat_direction", data.LatDirection)
	g.AddFloat64Attribute(body, "lat_minutes", data.LatMinutes)
	g.AddFloat64Attribute(body, "lat_seconds", data.LatSeconds)
	g.AddFloat64Attribute(body, "long_degrees", data.LongDegrees)
	g.AddStringAttribute(body, "long_direction", data.LongDirection)
	g.AddFloat64Attribute(body, "long_minutes", data.LongMinutes)
	g.AddFloat64Attribute(body, "long_seconds", data.LongSeconds)
	g.AddFloat64Attribute(body, "precision_horz", data.PrecisionHorz)
	g.AddFloat64Attribute(body, "precision_vert", data.PrecisionVert)
	g.AddFloat64Attribute(body, "size", data.Size)
}

// generateURIData generates URI-specific data fields
func (g *DNSRecordGenerator) generateURIData(body *hclwrite.Body, data *v5_models.DNSRecordDataV5Model) {
	g.AddFloat64Attribute(body, "priority", data.Priority)
	g.AddFloat64Attribute(body, "weight", data.Weight)
	g.AddStringAttribute(body, "target", data.Target)
}

// generateGenericData generates generic data fields for other record types
func (g *DNSRecordGenerator) generateGenericData(body *hclwrite.Body, data *v5_models.DNSRecordDataV5Model) {
	g.AddStringAttribute(body, "target", data.Target)
	g.AddFloat64Attribute(body, "priority", data.Priority)
	g.AddFloat64Attribute(body, "weight", data.Weight)
	g.AddFloat64Attribute(body, "port", data.Port)
}

// hasDataContent checks if the data model has any non-null values
func (g *DNSRecordGenerator) hasDataContent(data *v5_models.DNSRecordDataV5Model) bool {
	if data == nil {
		return false
	}

	// Check if any field has a value
	return !data.Flags.IsNull() ||
		!data.Tag.IsNull() ||
		!data.Value.IsNull() ||
		!data.Priority.IsNull() ||
		!data.Weight.IsNull() ||
		!data.Port.IsNull() ||
		!data.Target.IsNull() ||
		!data.Service.IsNull() ||
		!data.Altitude.IsNull() ||
		!data.LatDegrees.IsNull() ||
		!data.LongDegrees.IsNull()
}