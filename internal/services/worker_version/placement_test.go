package worker_version_test

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPlacementFromSettings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input worker_version.ScriptPlacementGetResponse
		want  *worker_version.WorkerVersionPlacementModel
	}{
		{
			name:  "empty returns nil",
			input: worker_version.ScriptPlacementGetResponse{},
			want:  nil,
		},
		{
			name: "region is populated",
			input: worker_version.ScriptPlacementGetResponse{
				Region: "WEUR",
			},
			want: &worker_version.WorkerVersionPlacementModel{
				Region:   types.StringValue("WEUR"),
				Mode:     types.StringNull(),
				Hostname: types.StringNull(),
				Host:     types.StringNull(),
			},
		},
		{
			name: "mode is populated",
			input: worker_version.ScriptPlacementGetResponse{
				Mode: worker_version.ScriptPlacementModeMode("smart"),
			},
			want: &worker_version.WorkerVersionPlacementModel{
				Region:   types.StringNull(),
				Mode:     types.StringValue("smart"),
				Hostname: types.StringNull(),
				Host:     types.StringNull(),
			},
		},
		{
			name: "hostname is populated",
			input: worker_version.ScriptPlacementGetResponse{
				Hostname: "example.com",
			},
			want: &worker_version.WorkerVersionPlacementModel{
				Region:   types.StringNull(),
				Mode:     types.StringNull(),
				Hostname: types.StringValue("example.com"),
				Host:     types.StringNull(),
			},
		},
		{
			name: "host is populated",
			input: worker_version.ScriptPlacementGetResponse{
				Host: "10.0.0.1:8080",
			},
			want: &worker_version.WorkerVersionPlacementModel{
				Region:   types.StringNull(),
				Mode:     types.StringNull(),
				Hostname: types.StringNull(),
				Host:     types.StringValue("10.0.0.1:8080"),
			},
		},
		{
			name: "all non-empty fields in response are populated in model",
			input: worker_version.ScriptPlacementGetResponse{
				Region:   "WEUR",
				Hostname: "example.com",
			},
			want: &worker_version.WorkerVersionPlacementModel{
				Region:   types.StringValue("WEUR"),
				Mode:     types.StringNull(),
				Hostname: types.StringValue("example.com"),
				Host:     types.StringNull(),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := worker_version.PlacementFromSettingsForTest(tc.input)

			if tc.want == nil {
				if got != nil {
					t.Errorf("expected nil, got %+v", got)
				}
				return
			}
			if got == nil {
				t.Fatalf("expected non-nil result, got nil")
			}
			if got.Region != tc.want.Region {
				t.Errorf("Region: got %v, want %v", got.Region, tc.want.Region)
			}
			if got.Mode != tc.want.Mode {
				t.Errorf("Mode: got %v, want %v", got.Mode, tc.want.Mode)
			}
			if got.Hostname != tc.want.Hostname {
				t.Errorf("Hostname: got %v, want %v", got.Hostname, tc.want.Hostname)
			}
			if got.Host != tc.want.Host {
				t.Errorf("Host: got %v, want %v", got.Host, tc.want.Host)
			}
		})
	}
}
