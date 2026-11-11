package load_balancer

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestIsZeroOrEmptyTime(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want bool
	}{
		{name: "empty", in: "", want: true},
		{name: "go zero time", in: "0001-01-01T00:00:00Z", want: true},
		{name: "real timestamp", in: "2026-11-11T09:07:16.596986Z", want: false},
		{name: "unparseable", in: "not-a-time", want: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isZeroOrEmptyTime(tc.in); got != tc.want {
				t.Fatalf("isZeroOrEmptyTime(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestPreserveCreatedOnFromState(t *testing.T) {
	prior := types.StringValue("2026-11-11T09:07:16Z")
	for _, tc := range []struct {
		name     string
		response types.String
		state    types.String
		want     types.String
	}{
		{
			name:     "zero time restored from state",
			response: types.StringValue("0001-01-01T00:00:00Z"), state: prior, want: prior,
		},
		{
			name:     "empty restored from state",
			response: types.StringValue(""), state: prior, want: prior,
		},
		{
			name:     "null restored from state",
			response: types.StringNull(), state: prior, want: prior,
		},
		{
			name:     "real response kept",
			response: types.StringValue("2026-12-01T00:00:00Z"), state: prior,
			want: types.StringValue("2026-12-01T00:00:00Z"),
		},
		{
			name:     "no prior leaves zero response untouched",
			response: types.StringValue("0001-01-01T00:00:00Z"), state: types.StringNull(),
			want: types.StringValue("0001-01-01T00:00:00Z"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data := &LoadBalancerModel{CreatedOn: tc.response}
			state := &LoadBalancerModel{CreatedOn: tc.state}
			preserveCreatedOnFromState(data, state)
			if !data.CreatedOn.Equal(tc.want) {
				t.Fatalf("created_on = %s, want %s", data.CreatedOn, tc.want)
			}
		})
	}
}

func TestPreserveTTLFromState(t *testing.T) {
	prior := types.Float64Value(30)
	for _, tc := range []struct {
		name     string
		response types.Float64
		state    types.Float64
		want     types.Float64
	}{
		{name: "null restored from state", response: types.Float64Null(), state: prior, want: prior},
		{name: "unknown restored from state", response: types.Float64Unknown(), state: prior, want: prior},
		{
			name: "real response kept", response: types.Float64Value(120), state: prior,
			want: types.Float64Value(120),
		},
		{
			name:     "no prior leaves null response untouched",
			response: types.Float64Null(), state: types.Float64Null(), want: types.Float64Null(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data := &LoadBalancerModel{TTL: tc.response}
			state := &LoadBalancerModel{TTL: tc.state}
			preserveTTLFromState(data, state)
			if !data.TTL.Equal(tc.want) {
				t.Fatalf("ttl = %s, want %s", data.TTL, tc.want)
			}
		})
	}
}
