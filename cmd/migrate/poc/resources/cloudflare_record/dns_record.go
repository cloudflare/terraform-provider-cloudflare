package cloudflare_record

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DNSRecord struct{}

func NewDNSRecord() interfaces.ResourceTransformer {
	return &DNSRecord{}
}

func (d *DNSRecord) CanHandle(resourceType string) bool {
	return resourceType == "cloudflare_record" || resourceType == "cloudflare_dns_record"
}

func (d *DNSRecord) GetResourceType() string {
	return "cloudflare_record"
}

// Preprocess handles string-level transformations before HCL parsing
func (d *DNSRecord) Preprocess(content string) string {
	if !strings.Contains(content, "cloudflare_record") {
		return content
	}
	
	content = strings.ReplaceAll(content, `resource "cloudflare_record"`, `resource "cloudflare_dns_record"`)
	content = strings.ReplaceAll(content, `data "cloudflare_record"`, `data "cloudflare_dns_record"`)

	return content
}

func (d *DNSRecord) handleDataBlockTransformations(block *hclwrite.Block, recordType string) {
	body := block.Body()
	hasDataBlock := false
	for _, b := range body.Blocks() {
		if b.Type() == "data" {
			hasDataBlock = true
			break
		}
	}

	if !hasDataBlock {
		return
	}

	if recordType == "SRV" || recordType == "MX" || recordType == "URI" {
		for _, dataBlock := range body.Blocks() {
			if dataBlock.Type() != "data" {
				continue
			}

			if priorityAttr := dataBlock.Body().GetAttribute("priority"); priorityAttr != nil {
				if body.GetAttribute("priority") == nil {
					priorityTokens := priorityAttr.Expr().BuildTokens(nil)
					body.SetAttributeRaw("priority", priorityTokens)
				}
			}
		}
	}

	var dataBlocksToRemove []*hclwrite.Block
	for _, dataBlock := range body.Blocks() {
		if dataBlock.Type() != "data" {
			continue
		}

		objTokens := d.buildDataObjectTokens(dataBlock, recordType)
		body.SetAttributeRaw("data", objTokens)
		dataBlocksToRemove = append(dataBlocksToRemove, dataBlock)
	}

	for _, dataBlock := range dataBlocksToRemove {
		body.RemoveBlock(dataBlock)
	}
}

func (d *DNSRecord) handleDataAttributeTransformations(block *hclwrite.Block, recordType string) {
	body := block.Body()
	dataAttr := body.GetAttribute("data")

	if dataAttr != nil && recordType == "CAA" {
		expr := dataAttr.Expr()
		tokens := expr.BuildTokens(nil)

		newTokens := make(hclwrite.Tokens, 0, len(tokens))
		for i := 0; i < len(tokens); i++ {
			token := tokens[i]
			if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "flags" {
				if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual {
					valueIdx := i + 2
					for valueIdx < len(tokens) && tokens[valueIdx].Type == hclsyntax.TokenNewline {
						valueIdx++
					}

					if valueIdx < len(tokens) && tokens[valueIdx].Type == hclsyntax.TokenQuotedLit {
						quotedValue := string(tokens[valueIdx].Bytes)
						unquoted := strings.Trim(quotedValue, `"`)
						if _, err := strconv.Atoi(unquoted); err == nil {
							newTokens = append(newTokens, token)
							newTokens = append(newTokens, tokens[i+1])
							numberToken := &hclwrite.Token{
								Type:  hclsyntax.TokenNumberLit,
								Bytes: []byte(unquoted),
							}
							newTokens = append(newTokens, numberToken)
							i = valueIdx
							continue
						}
					}
				}
			}

			if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "content" {
				if i+1 < len(tokens) && (tokens[i+1].Type == hclsyntax.TokenEqual ||
					(string(tokens[i+1].Bytes) == " " && i+2 < len(tokens) && tokens[i+2].Type == hclsyntax.TokenEqual)) {
					valueToken := &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("value"),
					}
					newTokens = append(newTokens, valueToken)
				} else {
					newTokens = append(newTokens, token)
				}
			} else {
				newTokens = append(newTokens, token)
			}
		}

		body.SetAttributeRaw("data", newTokens)
	}
}

func (d *DNSRecord) buildDataObjectTokens(dataBlock *hclwrite.Block, recordType string) hclwrite.Tokens {
	var objTokens hclwrite.Tokens
	objTokens = append(objTokens, &hclwrite.Token{
		Type:  hclsyntax.TokenOBrace,
		Bytes: []byte("{"),
	})
	objTokens = append(objTokens, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte("\n"),
	})

	attrs := dataBlock.Body().Attributes()
	processedAttrs := make(map[string]bool)
	var attrOrder []string
	switch recordType {
	case "CAA":
		attrOrder = []string{"flags", "tag", "value"}
	case "SRV":
		attrOrder = []string{"priority", "weight", "port", "target", "service"}
	case "URI":
		attrOrder = []string{"weight", "target"}
	default:
		attrOrder = []string{}
	}

	for _, attrName := range attrOrder {
		var attr *hclwrite.Attribute
		var finalName string
		var origName string

		if attrName == "value" {
			if contentAttr, exists := attrs["content"]; exists && recordType == "CAA" {
				attr = contentAttr
				finalName = "value"
				origName = "content"
			} else if valueAttr, exists := attrs["value"]; exists {
				attr = valueAttr
				finalName = "value"
				origName = "value"
			}
		} else {
			if a, exists := attrs[attrName]; exists {
				attr = a
				finalName = attrName
				origName = attrName
			}
		}

		if attr != nil {
			processedAttrs[origName] = true
			objTokens = d.addAttributeToTokens(objTokens, finalName, attr)
		}
	}

	for name, attr := range attrs {
		if !processedAttrs[name] {
			objTokens = d.addAttributeToTokens(objTokens, name, attr)
		}
	}
	objTokens = append(objTokens, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("  "),
	})
	objTokens = append(objTokens, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrace,
		Bytes: []byte("}"),
	})

	return objTokens
}

func (d *DNSRecord) addAttributeToTokens(tokens hclwrite.Tokens, name string, attr *hclwrite.Attribute) hclwrite.Tokens {
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("    "),
	})

	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte(name),
	})

	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenEqual,
		Bytes: []byte(" = "),
	})

	tokens = append(tokens, attr.Expr().BuildTokens(nil)...)

	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte("\n"),
	})

	return tokens
}

func (d *DNSRecord) TransformState(json gjson.Result, resourcePath string) (string, error) {
	result := json.String()
	resourceType := gjson.Get(result, resourcePath+".type").String()
	if resourceType == "cloudflare_record" {
		result, _ = sjson.Set(result, resourcePath+".type", "cloudflare_dns_record")
	}

	instances := gjson.Get(result, resourcePath+".instances")
	if instances.Exists() && instances.IsArray() {
		for i, instance := range instances.Array() {
			instancePath := resourcePath + ".instances." + strconv.Itoa(i)
			result = d.transformDNSRecordStateInstance(result, instancePath, instance)
		}
	}

	return result, nil
}

func (d *DNSRecord) transformDNSRecordStateInstance(result string, path string, instance gjson.Result) string {
	if !instance.Exists() || !instance.Get("attributes").Exists() {
		return result
	}

	attrs := instance.Get("attributes")
	if !attrs.Get("name").Exists() || !attrs.Get("type").Exists() || !attrs.Get("zone_id").Exists() {
		return result
	}

	meta := instance.Get("attributes.meta")
	if meta.Exists() {
		if meta.String() == "{}" || (meta.IsObject() && len(meta.Map()) == 0) {
			result, _ = sjson.Delete(result, path+".attributes.meta")
		}
	}

	settings := instance.Get("attributes.settings")
	if settings.Exists() {
		flattenCname := settings.Get("flatten_cname")
		ipv4Only := settings.Get("ipv4_only")
		ipv6Only := settings.Get("ipv6_only")

		allNull := (!flattenCname.Exists() || flattenCname.Type == gjson.Null || flattenCname.Value() == nil) &&
			(!ipv4Only.Exists() || ipv4Only.Type == gjson.Null || ipv4Only.Value() == nil) &&
			(!ipv6Only.Exists() || ipv6Only.Type == gjson.Null || ipv6Only.Value() == nil)

		if allNull {
			result, _ = sjson.Delete(result, path+".attributes.settings")
		}
	}

	createdOn := instance.Get("attributes.created_on")
	if !createdOn.Exists() {
		result, _ = sjson.Set(result, path+".attributes.created_on", "2024-01-01T00:00:00Z")
	}

	modifiedOn := instance.Get("attributes.modified_on")
	if !modifiedOn.Exists() {
		if createdOn.Exists() {
			result, _ = sjson.Set(result, path+".attributes.modified_on", createdOn.String())
		} else {
			result, _ = sjson.Set(result, path+".attributes.modified_on", "2024-01-01T00:00:00Z")
		}
	}

	value := instance.Get("attributes.value")
	content := instance.Get("attributes.content")
	if value.Exists() && !content.Exists() {
		result, _ = sjson.Set(result, path+".attributes.content", value.String())
		result, _ = sjson.Delete(result, path+".attributes.value")
	} else if value.Exists() && content.Exists() {
		result, _ = sjson.Delete(result, path+".attributes.value")
	}

	ttl := instance.Get("attributes.ttl")
	if !ttl.Exists() {
		result, _ = sjson.Set(result, path+".attributes.ttl", 1.0)
	}

	metadata := instance.Get("attributes.metadata")
	if metadata.Exists() {
		if metadata.IsObject() {
			if len(metadata.Map()) == 0 {
				result, _ = sjson.Delete(result, path+".attributes.metadata")
			} else {
				metaJSON, _ := json.Marshal(metadata.Value())
				result, _ = sjson.Set(result, path+".attributes.meta", string(metaJSON))
				result, _ = sjson.Delete(result, path+".attributes.metadata")
			}
		} else if metadata.String() == "" || metadata.String() == "{}" {
			result, _ = sjson.Delete(result, path+".attributes.metadata")
		} else {
			result, _ = sjson.Set(result, path+".attributes.meta", metadata.String())
			result, _ = sjson.Delete(result, path+".attributes.metadata")
		}
	}

	if instance.Get("attributes.hostname").Exists() {
		result, _ = sjson.Delete(result, path+".attributes.hostname")
	}
	if instance.Get("attributes.allow_overwrite").Exists() {
		result, _ = sjson.Delete(result, path+".attributes.allow_overwrite")
	}
	if instance.Get("attributes.timeouts").Exists() {
		result, _ = sjson.Delete(result, path+".attributes.timeouts")
	}

	recordType := instance.Get("attributes.type").String()
	result = d.transformDataField(result, path, instance, recordType)

	return result
}

func (d *DNSRecord) transformDataField(result string, path string, instance gjson.Result, recordType string) string {
	simpleTypes := map[string]bool{
		"A": true, "AAAA": true, "CNAME": true, "MX": true,
		"NS": true, "PTR": true, "TXT": true, "OPENPGPKEY": true,
	}

	if simpleTypes[recordType] {
		data := instance.Get("attributes.data")
		if data.Exists() {
			result, _ = sjson.Delete(result, path+".attributes.data")
		}
		return result
	}

	data := instance.Get("attributes.data")
	if !data.Exists() {
		dataObj := make(map[string]interface{})
		if recordType == "CAA" {
			dataObj["flags"] = nil
		}
		result, _ = sjson.Set(result, path+".attributes.data", dataObj)
		return result
	}

	dataObj := make(map[string]interface{})

	if data.IsArray() {
		dataArray := data.Array()
		if len(dataArray) == 0 {
			result, _ = sjson.Delete(result, path+".attributes.data")
			return result
		}

		if len(dataArray) > 0 {
			firstElem := dataArray[0]
			firstElem.ForEach(func(key, value gjson.Result) bool {
				k := key.String()

				if k == "name" || k == "proto" {
					return true
				}

				if recordType == "CAA" && k == "content" {
					dataObj["value"] = value.String()
					return true
				}

				if k == "flags" {
					dataObj[k] = d.wrapFlagsValue(value)
				} else if recordType == "CAA" && k == "value" {
					dataObj[k] = value.String()
				} else {
					switch value.Type {
					case gjson.Number:
						dataObj[k] = value.Float()
					case gjson.String:
						dataObj[k] = value.String()
					case gjson.True:
						dataObj[k] = true
					case gjson.False:
						dataObj[k] = false
					case gjson.Null:
						dataObj[k] = nil
					default:
						dataObj[k] = value.Value()
					}
				}
				return true
			})
		}

		if recordType == "CAA" {
			if _, hasFlags := dataObj["flags"]; !hasFlags {
				dataObj["flags"] = nil
			}
		}

		if len(dataObj) == 0 {
			result, _ = sjson.Delete(result, path+".attributes.data")
			return result
		}
		result, _ = sjson.Set(result, path+".attributes.data", dataObj)

	} else if data.IsObject() {
		data.ForEach(func(key, value gjson.Result) bool {
			k := key.String()

			if k == "name" || k == "proto" {
				return true
			}

			if k == "flags" {
				if value.IsObject() && value.Get("value").Exists() {
					dataObj[k] = value.Value()
				} else {
					dataObj[k] = d.wrapFlagsValue(value)
				}
			} else {
				switch value.Type {
				case gjson.Number:
					dataObj[k] = value.Float()
				case gjson.String:
					dataObj[k] = value.String()
				case gjson.True:
					dataObj[k] = true
				case gjson.False:
					dataObj[k] = false
				case gjson.Null:
					dataObj[k] = nil
				default:
					dataObj[k] = value.Value()
				}
			}
			return true
		})

		if recordType == "CAA" {
			if content := data.Get("content"); content.Exists() {
				dataObj["value"] = content.String()
				delete(dataObj, "content")
			}

			if _, hasFlags := dataObj["flags"]; !hasFlags {
				dataObj["flags"] = nil
			}
		}

		if len(dataObj) == 0 {
			result, _ = sjson.Delete(result, path+".attributes.data")
		} else {
			result, _ = sjson.Set(result, path+".attributes.data", dataObj)
		}
	}

	if recordType == "SRV" || recordType == "URI" {
		if priority, ok := dataObj["priority"]; ok && priority != nil {
			result, _ = sjson.Set(result, path+".attributes.priority", priority)
		}
	}

	return result
}

func (d *DNSRecord) wrapFlagsValue(value gjson.Result) interface{} {
	switch value.Type {
	case gjson.Number:
		return map[string]interface{}{
			"value": json.Number(value.Raw),
			"type":  "number",
		}
	case gjson.String:
		if _, err := strconv.ParseFloat(value.String(), 64); err == nil {
			return map[string]interface{}{
				"value": json.Number(value.String()),
				"type":  "number",
			}
		} else if value.String() == "" {
			return nil
		} else {
			return map[string]interface{}{
				"value": value.String(),
				"type":  "string",
			}
		}
	case gjson.Null:
		return nil
	default:
		return nil
	}
}

// Transform implements ResourceTransformer interface
func (d *DNSRecord) TransformConfig(block *hclwrite.Block) (*interfaces.TransformResult, error) {
	labels := block.Labels()
	if len(labels) < 2 {
		return &interfaces.TransformResult{
			Blocks:         []*hclwrite.Block{block},
			RemoveOriginal: false,
		}, nil
	}

	resourceType := labels[0]
	if resourceType != "cloudflare_dns_record" && resourceType != "cloudflare_record" {
		return &interfaces.TransformResult{
			Blocks:         []*hclwrite.Block{block},
			RemoveOriginal: false,
		}, nil
	}

	if resourceType == "cloudflare_record" {
		labels[0] = "cloudflare_dns_record"
		block.SetLabels(labels)
	}

	body := block.Body()

	ttlAttr := body.GetAttribute("ttl")
	if ttlAttr == nil {
		ttlToken := &hclwrite.Token{
			Type:  hclsyntax.TokenNumberLit,
			Bytes: []byte("1"),
		}
		body.SetAttributeRaw("ttl", hclwrite.Tokens{ttlToken})
	}

	typeAttr := body.GetAttribute("type")
	if typeAttr == nil {
		return &interfaces.TransformResult{
			Blocks:         []*hclwrite.Block{block},
			RemoveOriginal: false,
		}, nil
	}

	typeTokens := typeAttr.Expr().BuildTokens(nil)
	var recordType string
	for _, token := range typeTokens {
		if token.Type == hclsyntax.TokenQuotedLit || token.Type == hclsyntax.TokenIdent {
			recordType = strings.Trim(string(token.Bytes), "\"")
			break
		}
	}
	simpleTypes := map[string]bool{
		"A": true, "AAAA": true, "CNAME": true, "MX": true,
		"NS": true, "PTR": true, "TXT": true, "OPENPGPKEY": true,
	}
	if simpleTypes[recordType] {
		valueAttr := body.GetAttribute("value")
		if valueAttr != nil {
			valueTokens := valueAttr.Expr().BuildTokens(nil)
			body.SetAttributeRaw("content", valueTokens)
			body.RemoveAttribute("value")
		}
	}

	if allowOverwrite := body.GetAttribute("allow_overwrite"); allowOverwrite != nil {
		body.RemoveAttribute("allow_overwrite")
	}
	if hostname := body.GetAttribute("hostname"); hostname != nil {
		body.RemoveAttribute("hostname")
	}

	d.handleDataBlockTransformations(block, recordType)
	d.handleDataAttributeTransformations(block, recordType)

	return &interfaces.TransformResult{
		Blocks:         []*hclwrite.Block{block},
		RemoveOriginal: false,
	}, nil
}

// compile-time check to ensure we implement the interface
var _ interfaces.ResourceTransformer = (*DNSRecord)(nil)
