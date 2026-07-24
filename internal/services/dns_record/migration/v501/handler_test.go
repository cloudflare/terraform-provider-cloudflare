package v501

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestCanonicalizePriorityData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		state            v500.TargetDNSRecordModel
		expectData       bool
		expectedPriority types.Float64
		expectedTarget   types.String
		expectedRoot     types.Float64
	}{
		{
			name: "MX moves zero priority and target into data",
			state: v500.TargetDNSRecordModel{
				Type:     types.StringValue("MX"),
				Priority: types.Float64Value(0),
				Content:  types.StringValue("mail.example.com"),
			},
			expectData:       true,
			expectedPriority: types.Float64Value(0),
			expectedTarget:   types.StringValue("mail.example.com"),
			expectedRoot:     types.Float64Value(0),
		},
		{
			name: "MX nested values take precedence",
			state: v500.TargetDNSRecordModel{
				Type:     types.StringValue("mx"),
				Priority: types.Float64Value(99),
				Content:  types.StringValue("legacy.example.com"),
				Data: &v500.TargetDNSRecordDataModel{
					Priority: types.Float64Value(10),
					Target:   types.StringValue("mail.example.com"),
				},
			},
			expectData:       true,
			expectedPriority: types.Float64Value(10),
			expectedTarget:   types.StringValue("mail.example.com"),
			expectedRoot:     types.Float64Value(99),
		},
		{
			name: "URI merges root priority into existing data",
			state: v500.TargetDNSRecordModel{
				Type:     types.StringValue("URI"),
				Priority: types.Float64Value(20),
				Data: &v500.TargetDNSRecordDataModel{
					Target: types.StringValue("https://example.com"),
					Weight: types.Float64Value(0),
				},
			},
			expectData:       true,
			expectedPriority: types.Float64Value(20),
			expectedTarget:   types.StringValue("https://example.com"),
			expectedRoot:     types.Float64Value(20),
		},
		{
			name: "SRV fills missing nested priority",
			state: v500.TargetDNSRecordModel{
				Type:     types.StringValue("SRV"),
				Priority: types.Float64Value(5),
				Data: &v500.TargetDNSRecordDataModel{
					Target: types.StringValue("sip.example.com"),
				},
			},
			expectData:       true,
			expectedPriority: types.Float64Value(5),
			expectedTarget:   types.StringValue("sip.example.com"),
			expectedRoot:     types.Float64Value(5),
		},
		{
			name: "unrelated record is unchanged",
			state: v500.TargetDNSRecordModel{
				Type:     types.StringValue("A"),
				Priority: types.Float64Value(7),
			},
			expectData:   false,
			expectedRoot: types.Float64Value(7),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			CanonicalizePriorityData(&test.state)

			if !test.state.Priority.Equal(test.expectedRoot) {
				t.Fatalf("root priority = %s, want %s", test.state.Priority, test.expectedRoot)
			}
			if !test.expectData {
				if test.state.Data != nil {
					t.Fatalf("data unexpectedly materialized: %#v", test.state.Data)
				}
				return
			}
			if test.state.Data == nil {
				t.Fatal("data was not materialized")
			}
			if !test.state.Data.Priority.Equal(test.expectedPriority) {
				t.Errorf("data.priority = %s, want %s", test.state.Data.Priority, test.expectedPriority)
			}
			if !test.state.Data.Target.Equal(test.expectedTarget) {
				t.Errorf("data.target = %s, want %s", test.state.Data.Target, test.expectedTarget)
			}
		})
	}
}
