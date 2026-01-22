package hyperdrive_config

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPreserveWriteOnlyFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		source   *HyperdriveConfigModel
		dest     *HyperdriveConfigModel
		expected *HyperdriveConfigModel
	}{
		{
			name:   "nil source does nothing",
			source: nil,
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
		},
		{
			name: "nil dest does nothing",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringValue("secret"),
				},
			},
			dest:     nil,
			expected: nil,
		},
		{
			name: "preserves password from source to dest",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("my-secret-password"),
					AccessClientSecret: types.StringNull(),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringNull(),
					AccessClientSecret: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("my-secret-password"),
					AccessClientSecret: types.StringNull(),
				},
			},
		},
		{
			name: "preserves access_client_secret from source to dest",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("password"),
					AccessClientSecret: types.StringValue("my-access-secret"),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringNull(),
					AccessClientSecret: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("password"),
					AccessClientSecret: types.StringValue("my-access-secret"),
				},
			},
		},
		{
			name: "preserves both password and access_client_secret",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Database:           types.StringValue("mydb"),
					Host:               types.StringValue("localhost"),
					Password:           types.StringValue("password123"),
					AccessClientID:     types.StringValue("client-id"),
					AccessClientSecret: types.StringValue("client-secret"),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Database:           types.StringValue("mydb"),
					Host:               types.StringValue("localhost"),
					Password:           types.StringNull(), // API doesn't return this
					AccessClientID:     types.StringValue("client-id"),
					AccessClientSecret: types.StringNull(), // API doesn't return this
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Database:           types.StringValue("mydb"),
					Host:               types.StringValue("localhost"),
					Password:           types.StringValue("password123"),
					AccessClientID:     types.StringValue("client-id"),
					AccessClientSecret: types.StringValue("client-secret"),
				},
			},
		},
		{
			name: "does not overwrite when source has unknown value",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringUnknown(),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
		},
		{
			name: "handles nil origin in source",
			source: &HyperdriveConfigModel{
				Origin: nil,
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
		},
		{
			name: "handles nil origin in dest",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringValue("password"),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: nil,
			},
			expected: &HyperdriveConfigModel{
				Origin: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			preserveWriteOnlyFields(tt.source, tt.dest)

			// For nil dest, just verify it stayed nil
			if tt.expected == nil {
				if tt.dest != nil {
					t.Errorf("expected dest to be nil, got %v", tt.dest)
				}
				return
			}

			// For nil expected origin, verify dest origin is nil
			if tt.expected.Origin == nil {
				if tt.dest.Origin != nil {
					t.Errorf("expected dest.Origin to be nil, got %v", tt.dest.Origin)
				}
				return
			}

			// Verify password
			if !tt.dest.Origin.Password.Equal(tt.expected.Origin.Password) {
				t.Errorf("password mismatch: got %v, want %v",
					tt.dest.Origin.Password, tt.expected.Origin.Password)
			}

			// Verify access_client_secret
			if !tt.dest.Origin.AccessClientSecret.Equal(tt.expected.Origin.AccessClientSecret) {
				t.Errorf("access_client_secret mismatch: got %v, want %v",
					tt.dest.Origin.AccessClientSecret, tt.expected.Origin.AccessClientSecret)
			}
		})
	}
}

