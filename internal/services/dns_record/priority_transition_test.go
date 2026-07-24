package dns_record

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMarshalDNSRecordForCreateUsesCanonicalData(t *testing.T) {
	t.Parallel()

	record := &DNSRecordModel{
		ZoneID:   types.StringValue("zone"),
		Name:     types.StringValue("example.com"),
		Type:     types.StringValue("MX"),
		TTL:      types.Float64Value(300),
		Priority: types.Float64Value(99),
		Data: &DNSRecordDataModel{
			Priority: types.Float64Value(0),
			Target:   types.StringValue("mail.example.com"),
		},
	}

	actual, err := marshalDNSRecordForCreate(record)
	if err != nil {
		t.Fatalf("marshal create request: %v", err)
	}
	assertJSONValueEqual(t, actual, []byte(`{
		"name":"example.com",
		"type":"MX",
		"ttl":300,
		"priority":0,
		"content":"mail.example.com"
	}`))
}

func TestMarshalDNSRecordForUpdateUsesCanonicalData(t *testing.T) {
	t.Parallel()

	state := &DNSRecordModel{
		Name:    types.StringValue("_service.example.com"),
		Type:    types.StringValue("URI"),
		TTL:     types.Float64Value(300),
		Content: types.StringValue("10 5 https://example.com/old"),
		Data: &DNSRecordDataModel{
			Priority: types.Float64Value(10),
			Weight:   types.Float64Value(5),
			Target:   types.StringValue("https://example.com/old"),
		},
	}
	plan := *state
	plan.Data = &DNSRecordDataModel{
		Priority: types.Float64Value(20),
		Weight:   types.Float64Value(0),
		Target:   types.StringValue("https://example.com/new"),
	}

	actual, err := marshalDNSRecordForUpdate(&plan, state)
	if err != nil {
		t.Fatalf("marshal update request: %v", err)
	}
	var value map[string]any
	if err := json.Unmarshal(actual, &value); err != nil {
		t.Fatalf("decode update request: %v", err)
	}
	if value["priority"] != float64(20) {
		t.Errorf("root priority = %#v, want 20", value["priority"])
	}
	if _, ok := value["content"]; ok {
		t.Error("read-only URI content was sent in the update request")
	}
	data := value["data"].(map[string]any)
	if _, ok := data["priority"]; ok {
		t.Error("data.priority was sent to the legacy URI API")
	}
	if data["weight"] != float64(0) || data["target"] != "https://example.com/new" {
		t.Errorf("unexpected URI data: %#v", data)
	}
}

func TestNormalizeDNSRecordRequestJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		recordType string
		input      string
		expected   string
	}{
		{
			name:       "MX maps nested priority and target to root fields and removes data",
			recordType: "MX",
			input:      `{"type":"MX","name":"example.com","priority":99,"content":"old.example.com","data":{"priority":10,"target":"mail.example.com"}}`,
			expected:   `{"type":"MX","name":"example.com","priority":10,"content":"mail.example.com"}`,
		},
		{
			name:       "URI maps priority to root and retains target and weight in data",
			recordType: "URI",
			input:      `{"type":"URI","name":"example.com","priority":99,"content":"20 5 https://example.com","data":{"priority":20,"target":"https://example.com","weight":5}}`,
			expected:   `{"type":"URI","name":"example.com","priority":20,"data":{"target":"https://example.com","weight":5}}`,
		},
		{
			name:       "SRV keeps nested priority and removes root priority",
			recordType: "SRV",
			input:      `{"type":"SRV","name":"_sip._tcp.example.com","priority":99,"content":"10 5060 sip.example.com","data":{"priority":30,"target":"sip.example.com","port":5060,"weight":10}}`,
			expected:   `{"type":"SRV","name":"_sip._tcp.example.com","data":{"priority":30,"target":"sip.example.com","port":5060,"weight":10}}`,
		},
		{
			name:       "unrelated record type is unchanged",
			recordType: "CNAME",
			input:      `{"type":"CNAME","name":"www.example.com","content":"example.com","priority":7,"data":{"priority":8,"target":"untouched.example.com"}}`,
			expected:   `{"type":"CNAME","name":"www.example.com","content":"example.com","priority":7,"data":{"priority":8,"target":"untouched.example.com"}}`,
		},
		{
			name:       "zero priority is preserved",
			recordType: "MX",
			input:      `{"type":"MX","name":"example.com","data":{"priority":0,"target":"mail.example.com"}}`,
			expected:   `{"type":"MX","name":"example.com","priority":0,"content":"mail.example.com"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := normalizeDNSRecordRequestJSON([]byte(test.input), test.recordType)
			if err != nil {
				t.Fatalf("normalize request JSON: %v", err)
			}
			assertJSONValueEqual(t, actual, []byte(test.expected))
		})
	}
}

func TestNormalizeDNSRecordResponseJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "legacy MX root fields synthesize data",
			input:    `{"success":true,"result":{"id":"mx","type":"MX","priority":10,"content":"mail.example.com"}}`,
			expected: `{"success":true,"result":{"id":"mx","type":"MX","priority":10,"content":"mail.example.com","data":{"priority":10,"target":"mail.example.com"}}}`,
		},
		{
			name:     "legacy URI root priority merges existing data and preserves zero",
			input:    `{"success":true,"result":{"id":"uri","type":"URI","priority":0,"data":{"target":"https://example.com","weight":5}}}`,
			expected: `{"success":true,"result":{"id":"uri","type":"URI","priority":0,"data":{"priority":0,"target":"https://example.com","weight":5}}}`,
		},
		{
			name:     "SRV nested priority wins over conflicting root priority",
			input:    `{"success":true,"result":{"id":"srv","type":"SRV","priority":99,"data":{"priority":30,"target":"sip.example.com","port":5060,"weight":10}}}`,
			expected: `{"success":true,"result":{"id":"srv","type":"SRV","priority":99,"data":{"priority":30,"target":"sip.example.com","port":5060,"weight":10}}}`,
		},
		{
			name:     "future dual MX nested fields win over conflicting root fields",
			input:    `{"success":true,"result":{"id":"mx","type":"MX","priority":99,"content":"old.example.com","data":{"priority":10,"target":"mail.example.com"}}}`,
			expected: `{"success":true,"result":{"id":"mx","type":"MX","priority":99,"content":"old.example.com","data":{"priority":10,"target":"mail.example.com"}}}`,
		},
		{
			name:     "future nested-only MX response is unchanged",
			input:    `{"success":true,"result":{"id":"mx","type":"MX","content":"10 mail.example.com","data":{"priority":10,"target":"mail.example.com"}}}`,
			expected: `{"success":true,"result":{"id":"mx","type":"MX","content":"10 mail.example.com","data":{"priority":10,"target":"mail.example.com"}}}`,
		},
		{
			name: "list envelope normalizes multiple relevant records and leaves irrelevant records unchanged",
			input: `{
				"success": true,
				"result": [
					{"id":"mx","type":"MX","priority":0,"content":"mail.example.com"},
					{"id":"uri","type":"URI","priority":20,"data":{"target":"https://example.com","weight":5}},
					{"id":"a","type":"A","content":"192.0.2.1","priority":77,"data":{"custom":"unchanged"}}
				],
				"result_info": {"count":3}
			}`,
			expected: `{
				"success": true,
				"result": [
					{"id":"mx","type":"MX","priority":0,"content":"mail.example.com","data":{"priority":0,"target":"mail.example.com"}},
					{"id":"uri","type":"URI","priority":20,"data":{"priority":20,"target":"https://example.com","weight":5}},
					{"id":"a","type":"A","content":"192.0.2.1","priority":77,"data":{"custom":"unchanged"}}
				],
				"result_info": {"count":3}
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := normalizeDNSRecordResponseJSON([]byte(test.input))
			if err != nil {
				t.Fatalf("normalize response JSON: %v", err)
			}
			assertJSONValueEqual(t, actual, []byte(test.expected))
		})
	}
}

func assertJSONValueEqual(t *testing.T, actual, expected []byte) {
	t.Helper()

	var actualValue any
	if err := json.Unmarshal(actual, &actualValue); err != nil {
		t.Fatalf("decode actual JSON %q: %v", actual, err)
	}

	var expectedValue any
	if err := json.Unmarshal(expected, &expectedValue); err != nil {
		t.Fatalf("decode expected JSON %q: %v", expected, err)
	}

	if !reflect.DeepEqual(actualValue, expectedValue) {
		t.Errorf("JSON values differ\nactual:   %#v\nexpected: %#v", actualValue, expectedValue)
	}
}
