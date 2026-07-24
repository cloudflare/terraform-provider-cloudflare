package dns_record

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var priorityDataRecordTypes = map[string]struct{}{
	"MX":  {},
	"SRV": {},
	"URI": {},
}

// DNSRecordConfigValidator enforces the canonical structured representation
// for record types whose priority is part of their RDATA.
type DNSRecordConfigValidator struct{}

func (DNSRecordConfigValidator) Description(context.Context) string {
	return "validates type-specific DNS record data"
}

func (DNSRecordConfigValidator) MarkdownDescription(context.Context) string {
	return "validates type-specific DNS record data"
}

func (DNSRecordConfigValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var recordType types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("type"), &recordType)...)
	if resp.Diagnostics.HasError() || recordType.IsNull() || recordType.IsUnknown() {
		return
	}

	recordTypeValue := strings.ToUpper(recordType.ValueString())
	if _, ok := priorityDataRecordTypes[recordTypeValue]; !ok {
		return
	}

	var priority types.Float64
	priorityPath := path.Root("data").AtName("priority")
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, priorityPath, &priority)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if priority.IsNull() {
		resp.Diagnostics.AddAttributeError(
			priorityPath,
			"Missing DNS record priority",
			"MX, SRV, and URI records must configure priority within the data object.",
		)
	}

	if recordTypeValue != "MX" {
		return
	}

	var target types.String
	targetPath := path.Root("data").AtName("target")
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, targetPath, &target)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if target.IsNull() {
		resp.Diagnostics.AddAttributeError(
			targetPath,
			"Missing MX record target",
			"MX records must configure the mail server hostname in data.target instead of content.",
		)
	}

	var content types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("content"), &content)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !content.IsNull() && !content.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("content"),
			"MX record content is read-only",
			"Configure the mail server hostname in data.target. The formatted content is returned by the API.",
		)
	}
}

// marshalDNSRecordForCreate converts the canonical Terraform representation
// to the legacy API request representation used during the transition.
func marshalDNSRecordForCreate(data *DNSRecordModel) ([]byte, error) {
	encoded, err := data.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return normalizeDNSRecordRequestJSON(encoded, data.Type.ValueString())
}

// marshalDNSRecordForUpdate converts a canonical Terraform representation to
// the legacy full-update request representation.
func marshalDNSRecordForUpdate(data, state *DNSRecordModel) ([]byte, error) {
	encoded, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		return nil, err
	}
	return normalizeDNSRecordRequestJSON(encoded, data.Type.ValueString())
}

func normalizeDNSRecordRequestJSON(encoded []byte, recordType string) ([]byte, error) {
	var record map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &record); err != nil {
		return nil, err
	}

	switch strings.ToUpper(recordType) {
	case "MX":
		data, err := rawJSONObject(record["data"])
		if err != nil {
			return nil, err
		}
		if priority, ok := presentRawJSON(data["priority"]); ok {
			record["priority"] = priority
		}
		if target, ok := presentRawJSON(data["target"]); ok {
			record["content"] = target
		}
		delete(record, "data")
	case "URI":
		data, err := rawJSONObject(record["data"])
		if err != nil {
			return nil, err
		}
		if priority, ok := presentRawJSON(data["priority"]); ok {
			record["priority"] = priority
			delete(data, "priority")
		}
		if len(data) == 0 {
			delete(record, "data")
		} else {
			record["data"], err = json.Marshal(data)
			if err != nil {
				return nil, err
			}
		}
		delete(record, "content")
	case "SRV":
		// SRV already uses data.priority in the API. The top-level value is
		// response-only compatibility data and must not be sent on writes.
		delete(record, "priority")
		delete(record, "content")
	}

	return json.Marshal(record)
}

// normalizeDNSRecordResponseJSON materializes canonical nested fields from
// either the current legacy response or a future dual-field response. Existing
// nested values win because the API transition contract gives data precedence.
func normalizeDNSRecordResponseJSON(encoded []byte) ([]byte, error) {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &envelope); err != nil {
		return nil, err
	}

	result, ok := envelope["result"]
	if !ok {
		return encoded, nil
	}

	trimmed := bytes.TrimSpace(result)
	if len(trimmed) == 0 {
		return encoded, nil
	}

	if trimmed[0] == '[' {
		var records []json.RawMessage
		if err := json.Unmarshal(result, &records); err != nil {
			return nil, err
		}
		for i := range records {
			normalized, err := normalizeDNSRecordResponseObject(records[i])
			if err != nil {
				return nil, err
			}
			records[i] = normalized
		}
		envelope["result"], _ = json.Marshal(records)
	} else {
		normalized, err := normalizeDNSRecordResponseObject(result)
		if err != nil {
			return nil, err
		}
		envelope["result"] = normalized
	}

	return json.Marshal(envelope)
}

func normalizeDNSRecordResponseObject(encoded json.RawMessage) (json.RawMessage, error) {
	var record map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &record); err != nil {
		return nil, err
	}

	var recordType string
	if err := json.Unmarshal(record["type"], &recordType); err != nil {
		return encoded, nil
	}
	if _, ok := priorityDataRecordTypes[strings.ToUpper(recordType)]; !ok {
		return encoded, nil
	}

	data, err := rawJSONObject(record["data"])
	if err != nil {
		return nil, err
	}
	if _, ok := presentRawJSON(data["priority"]); !ok {
		if priority, ok := presentRawJSON(record["priority"]); ok {
			data["priority"] = priority
		}
	}
	if strings.EqualFold(recordType, "MX") {
		if _, ok := presentRawJSON(data["target"]); !ok {
			if content, ok := presentRawJSON(record["content"]); ok {
				data["target"] = content
			}
		}
	}

	if len(data) > 0 {
		record["data"], err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(record)
}

func rawJSONObject(encoded json.RawMessage) (map[string]json.RawMessage, error) {
	result := map[string]json.RawMessage{}
	if _, ok := presentRawJSON(encoded); !ok {
		return result, nil
	}
	if err := json.Unmarshal(encoded, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func presentRawJSON(value json.RawMessage) (json.RawMessage, bool) {
	trimmed := bytes.TrimSpace(value)
	return value, len(trimmed) > 0 && !bytes.Equal(trimmed, []byte("null"))
}
