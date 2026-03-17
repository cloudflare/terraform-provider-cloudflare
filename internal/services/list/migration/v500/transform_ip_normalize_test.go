package v500

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNormalizeIPAddress(t *testing.T) {
	tests := []struct {
		name string
		in   types.String
		want types.String
	}{
		{
			name: "preserve network CIDR",
			in:   types.StringValue("10.0.0.0/8"),
			want: types.StringValue("10.0.0.0/8"),
		},
		{
			name: "strip ipv4 host CIDR",
			in:   types.StringValue("192.0.2.1/32"),
			want: types.StringValue("192.0.2.1"),
		},
		{
			name: "strip ipv6 host CIDR",
			in:   types.StringValue("2001:db8::1/128"),
			want: types.StringValue("2001:db8::1"),
		},
		{
			name: "preserve plain ip",
			in:   types.StringValue("1.1.1.1"),
			want: types.StringValue("1.1.1.1"),
		},
		{
			name: "preserve null",
			in:   types.StringNull(),
			want: types.StringNull(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeIPAddress(tt.in)
			if got.String() != tt.want.String() {
				t.Fatalf("normalizeIPAddress(%s) = %s, want %s", tt.in.String(), got.String(), tt.want.String())
			}
		})
	}
}
