package list

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNormalizeListIP(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		in      string
		want    string
		parseOK bool
	}{
		{"ipv4 /32 stripped", "54.81.199.220/32", "54.81.199.220", true},
		{"ipv6 /128 stripped", "2001:db8::1/128", "2001:db8::1", true},
		{"ipv4 cidr host bits masked", "192.168.1.5/24", "192.168.1.0/24", true},
		{"ipv4 cidr already normalized", "10.0.0.0/24", "10.0.0.0/24", true},
		{"ipv4 bare already normalized", "1.2.3.4", "1.2.3.4", true},
		{"ipv6 canonicalized", "2001:0db8:0000:0000:0000:0000:0000:0001", "2001:db8::1", true},
		{"ipv6 uppercase lowercased", "2001:DB8::1", "2001:db8::1", true},
		{"ipv6 cidr host bits masked", "2001:db8::1/32", "2001:db8::/32", true},
		{"unparseable", "banana", "", false},
		{"empty string", "", "", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, ok := normalizeListIP(tc.in)
			if ok != tc.parseOK {
				t.Fatalf("parse ok = %v, want %v", ok, tc.parseOK)
			}
			if got != tc.want {
				t.Errorf("normalizeListIP(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestListIPNormalizer_PlanModifyString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		config types.String
		want   types.String
	}{
		{"null passes through", types.StringNull(), types.StringNull()},
		{"unknown passes through", types.StringUnknown(), types.StringUnknown()},
		{"/32 rewritten", types.StringValue("1.2.3.4/32"), types.StringValue("1.2.3.4")},
		{"/128 rewritten", types.StringValue("::1/128"), types.StringValue("::1")},
		{"already canonical unchanged", types.StringValue("10.0.0.0/24"), types.StringValue("10.0.0.0/24")},
		{"unparseable left alone", types.StringValue("not-an-ip"), types.StringValue("not-an-ip")},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := planmodifier.StringRequest{
				Path:        path.Root("ip"),
				ConfigValue: tc.config,
				PlanValue:   tc.config,
			}
			resp := &planmodifier.StringResponse{PlanValue: tc.config}
			listIPNormalizer{}.PlanModifyString(context.Background(), req, resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Errorf("PlanValue = %v, want %v", resp.PlanValue, tc.want)
			}
			if resp.Diagnostics.HasError() {
				t.Errorf("unexpected diagnostics: %v", resp.Diagnostics)
			}
		})
	}
}
