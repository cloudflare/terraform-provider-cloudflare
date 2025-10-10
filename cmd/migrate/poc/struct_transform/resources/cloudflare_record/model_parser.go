package cloudflare_record

import (
	"fmt"
	"strconv"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v4_models"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HCLParser parses HCL blocks into v4 model structs
type HCLParser struct {
	// Options
	StrictMode bool // If true, fail on unknown attributes
}

// NewHCLParser creates a new parser
func NewHCLParser() *HCLParser {
	return &HCLParser{
		StrictMode: false,
	}
}

// Parse implements the ModelParser interface
func (p *HCLParser) Parse(block *hclwrite.Block) (interface{}, error) {
	return p.ParseDNSRecord(block)
}

// ParseDNSRecord parses an HCL block into a v4 DNS record model
func (p *HCLParser) ParseDNSRecord(block *hclwrite.Block) (*v4_models.DNSRecordV4Model, error) {
	model := &v4_models.DNSRecordV4Model{}
	body := block.Body()

	// Parse attributes
	for name, attr := range body.Attributes() {
		if err := p.parseAttribute(name, attr, model); err != nil {
			if p.StrictMode {
				return nil, err
			}
			// In non-strict mode, just log and continue
			fmt.Printf("Warning: %v\n", err)
		}
	}

	// Parse data blocks
	dataBlocks := []v4_models.DNSRecordDataV4Model{}
	for _, nestedBlock := range body.Blocks() {
		if nestedBlock.Type() == "data" {
			dataModel, err := p.parseDataBlock(nestedBlock)
			if err != nil {
				if p.StrictMode {
					return nil, err
				}
				fmt.Printf("Warning parsing data block: %v\n", err)
				continue
			}
			dataBlocks = append(dataBlocks, *dataModel)
		}
	}
	model.Data = dataBlocks

	return model, nil
}

// parseAttribute parses a single attribute into the model
func (p *HCLParser) parseAttribute(name string, attr *hclwrite.Attribute, model *v4_models.DNSRecordV4Model) error {
	tokens := attr.Expr().BuildTokens(nil)

	switch name {
	case "zone_id":
		model.ZoneID = p.parseStringAttribute(tokens)
	case "name":
		model.Name = p.parseStringAttribute(tokens)
	case "type":
		model.Type = p.parseStringAttribute(tokens)
	case "value":
		model.Value = p.parseStringAttribute(tokens)
	case "content":
		model.Content = p.parseStringAttribute(tokens)
	case "ttl":
		model.TTL = p.parseFloat64Attribute(tokens)
	case "proxied":
		model.Proxied = p.parseBoolAttribute(tokens)
	case "priority":
		model.Priority = p.parseFloat64Attribute(tokens)
	case "allow_overwrite":
		model.AllowOverwrite = p.parseBoolAttribute(tokens)
	case "hostname":
		model.Hostname = p.parseStringAttribute(tokens)
	default:
		if p.StrictMode {
			return fmt.Errorf("unknown attribute: %s", name)
		}
	}

	return nil
}

// parseDataBlock parses a data block into a v4 data model
func (p *HCLParser) parseDataBlock(block *hclwrite.Block) (*v4_models.DNSRecordDataV4Model, error) {
	model := &v4_models.DNSRecordDataV4Model{}
	body := block.Body()

	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)

		switch name {
		// CAA fields
		case "flags":
			model.Flags = p.parseStringAttribute(tokens) // v4 uses string for flags
		case "tag":
			model.Tag = p.parseStringAttribute(tokens)
		case "content":
			model.Content = p.parseStringAttribute(tokens)

		// SRV fields
		case "priority":
			model.Priority = p.parseFloat64Attribute(tokens)
		case "weight":
			model.Weight = p.parseFloat64Attribute(tokens)
		case "port":
			model.Port = p.parseFloat64Attribute(tokens)
		case "target":
			model.Target = p.parseStringAttribute(tokens)
		case "service":
			model.Service = p.parseStringAttribute(tokens)
		case "proto":
			model.Proto = p.parseStringAttribute(tokens)
		case "name":
			model.Name = p.parseStringAttribute(tokens)

		// LOC fields
		case "altitude":
			model.Altitude = p.parseFloat64Attribute(tokens)
		case "lat_degrees":
			model.LatDegrees = p.parseFloat64Attribute(tokens)
		case "lat_direction":
			model.LatDirection = p.parseStringAttribute(tokens)
		case "lat_minutes":
			model.LatMinutes = p.parseFloat64Attribute(tokens)
		case "lat_seconds":
			model.LatSeconds = p.parseFloat64Attribute(tokens)
		case "long_degrees":
			model.LongDegrees = p.parseFloat64Attribute(tokens)
		case "long_direction":
			model.LongDirection = p.parseStringAttribute(tokens)
		case "long_minutes":
			model.LongMinutes = p.parseFloat64Attribute(tokens)
		case "long_seconds":
			model.LongSeconds = p.parseFloat64Attribute(tokens)
		case "precision_horz":
			model.PrecisionHorz = p.parseFloat64Attribute(tokens)
		case "precision_vert":
			model.PrecisionVert = p.parseFloat64Attribute(tokens)
		case "size":
			model.Size = p.parseFloat64Attribute(tokens)

		default:
			if p.StrictMode {
				return nil, fmt.Errorf("unknown data attribute: %s", name)
			}
		}
	}

	return model, nil
}

// parseStringAttribute extracts a string value from tokens
func (p *HCLParser) parseStringAttribute(tokens hclwrite.Tokens) types.String {
	for _, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenQuotedLit:
			// TokenQuotedLit contains the string content without quotes
			return types.StringValue(string(token.Bytes))
		case hclsyntax.TokenIdent:
			// Variable reference or literal identifier
			return types.StringValue(string(token.Bytes))
		}
	}
	return types.StringNull()
}

// parseFloat64Attribute extracts a float64 value from tokens
func (p *HCLParser) parseFloat64Attribute(tokens hclwrite.Tokens) types.Float64 {
	for _, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenNumberLit:
			value, err := strconv.ParseFloat(string(token.Bytes), 64)
			if err == nil {
				return types.Float64Value(value)
			}
		case hclsyntax.TokenQuotedLit:
			// Try to parse quoted number (TokenQuotedLit already has quotes removed)
			strValue := string(token.Bytes)
			value, err := strconv.ParseFloat(strValue, 64)
			if err == nil {
				return types.Float64Value(value)
			}
		}
	}
	return types.Float64Null()
}

// parseBoolAttribute extracts a bool value from tokens
func (p *HCLParser) parseBoolAttribute(tokens hclwrite.Tokens) types.Bool {
	for _, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenIdent:
			value := string(token.Bytes)
			if value == "true" {
				return types.BoolValue(true)
			} else if value == "false" {
				return types.BoolValue(false)
			}
		}
	}
	return types.BoolNull()
}