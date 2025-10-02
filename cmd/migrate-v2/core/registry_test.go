package core

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockMigration implements ResourceMigration for testing
type mockMigration struct {
	resourceType  string
	sourceVersion string
	targetVersion string
	migrateError  error
}

func (m *mockMigration) ResourceType() string  { return m.resourceType }
func (m *mockMigration) SourceVersion() string { return m.sourceVersion }
func (m *mockMigration) TargetVersion() string { return m.targetVersion }
func (m *mockMigration) MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error {
	return m.migrateError
}
func (m *mockMigration) MigrateState(state map[string]interface{}, ctx *MigrationContext) error {
	return m.migrateError
}
func (m *mockMigration) Validate(block *hclwrite.Block) error { return nil }

func TestDefaultRegistry_Register(t *testing.T) {
	tests := []struct {
		name        string
		migrations  []ResourceMigration
		expectError bool
	}{
		{
			name: "register single migration",
			migrations: []ResourceMigration{
				&mockMigration{
					resourceType:  "test_resource",
					sourceVersion: "v4",
					targetVersion: "v5",
				},
			},
			expectError: false,
		},
		{
			name: "register multiple migrations for different resources",
			migrations: []ResourceMigration{
				&mockMigration{
					resourceType:  "resource_a",
					sourceVersion: "v4",
					targetVersion: "v5",
				},
				&mockMigration{
					resourceType:  "resource_b",
					sourceVersion: "v4",
					targetVersion: "v5",
				},
			},
			expectError: false,
		},
		{
			name: "register multiple versions for same resource",
			migrations: []ResourceMigration{
				&mockMigration{
					resourceType:  "test_resource",
					sourceVersion: "v4",
					targetVersion: "v5",
				},
				&mockMigration{
					resourceType:  "test_resource",
					sourceVersion: "v5",
					targetVersion: "v6",
				},
			},
			expectError: false,
		},
		{
			name: "register duplicate migration fails",
			migrations: []ResourceMigration{
				&mockMigration{
					resourceType:  "test_resource",
					sourceVersion: "v4",
					targetVersion: "v5",
				},
				&mockMigration{
					resourceType:  "test_resource",
					sourceVersion: "v4",
					targetVersion: "v5",
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := NewDefaultRegistry()
			var err error
			
			for i, migration := range tt.migrations {
				err = registry.Register(migration)
				if tt.expectError && i == len(tt.migrations)-1 {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}

func TestDefaultRegistry_Get(t *testing.T) {
	registry := NewDefaultRegistry()
	
	// Register test migrations
	migrations := []ResourceMigration{
		&mockMigration{
			resourceType:  "resource_a",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		&mockMigration{
			resourceType:  "resource_b",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		&mockMigration{
			resourceType:  "resource_a",
			sourceVersion: "v5",
			targetVersion: "v6",
		},
	}
	
	for _, m := range migrations {
		require.NoError(t, registry.Register(m))
	}

	tests := []struct {
		name          string
		resourceType  string
		sourceVersion string
		targetVersion string
		expectError   bool
	}{
		{
			name:          "get existing migration",
			resourceType:  "resource_a",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectError:   false,
		},
		{
			name:          "get different existing migration",
			resourceType:  "resource_b",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectError:   false,
		},
		{
			name:          "get v5 to v6 migration",
			resourceType:  "resource_a",
			sourceVersion: "v5",
			targetVersion: "v6",
			expectError:   false,
		},
		{
			name:          "get non-existent resource",
			resourceType:  "non_existent",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectError:   true,
		},
		{
			name:          "get non-existent version",
			resourceType:  "resource_a",
			sourceVersion: "v3",
			targetVersion: "v4",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migration, err := registry.Get(tt.resourceType, tt.sourceVersion, tt.targetVersion)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, migration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, migration)
				assert.Equal(t, tt.resourceType, migration.ResourceType())
				assert.Equal(t, tt.sourceVersion, migration.SourceVersion())
				assert.Equal(t, tt.targetVersion, migration.TargetVersion())
			}
		})
	}
}

func TestDefaultRegistry_ListAvailable(t *testing.T) {
	registry := NewDefaultRegistry()
	
	// Start with empty registry
	infos := registry.ListAvailable()
	assert.Empty(t, infos)
	
	// Register migrations
	migrations := []ResourceMigration{
		&mockMigration{
			resourceType:  "resource_a",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		&mockMigration{
			resourceType:  "resource_b",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		&mockMigration{
			resourceType:  "resource_a",
			sourceVersion: "v5",
			targetVersion: "v6",
		},
	}
	
	for _, m := range migrations {
		require.NoError(t, registry.Register(m))
	}
	
	// List available migrations
	infos = registry.ListAvailable()
	assert.Len(t, infos, 3)
	
	// Verify all migrations are listed
	found := make(map[string]bool)
	for _, info := range infos {
		key := info.ResourceType + ":" + info.SourceVersion + "->" + info.TargetVersion
		found[key] = true
	}
	
	assert.True(t, found["resource_a:v4->v5"])
	assert.True(t, found["resource_b:v4->v5"])
	assert.True(t, found["resource_a:v5->v6"])
}

func TestDefaultRegistry_GetPath(t *testing.T) {
	registry := NewDefaultRegistry()
	
	// Register migrations for path finding
	migrations := []ResourceMigration{
		&mockMigration{
			resourceType:  "resource_a",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		&mockMigration{
			resourceType:  "resource_a",
			sourceVersion: "v5",
			targetVersion: "v6",
		},
	}
	
	for _, m := range migrations {
		require.NoError(t, registry.Register(m))
	}

	tests := []struct {
		name          string
		resourceType  string
		sourceVersion string
		targetVersion string
		expectPath    int
		expectError   bool
	}{
		{
			name:          "direct migration path",
			resourceType:  "resource_a",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectPath:    1,
			expectError:   false,
		},
		{
			name:          "no migration path available",
			resourceType:  "resource_b",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectError:   true,
		},
		// Note: Multi-hop path finding is not fully implemented yet
		// This test documents current behavior
		{
			name:          "multi-hop path not directly available",
			resourceType:  "resource_a",
			sourceVersion: "v4",
			targetVersion: "v6",
			expectError:   true, // Current implementation doesn't find multi-hop paths
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := registry.GetPath(tt.resourceType, tt.sourceVersion, tt.targetVersion)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, path)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, path)
				assert.Len(t, path, tt.expectPath)
			}
		})
	}
}

func TestPathFinder_FindPath(t *testing.T) {
	registry := NewDefaultRegistry()
	pathFinder := NewPathFinder(registry)
	
	// Register migrations for path finding
	migrations := []ResourceMigration{
		&mockMigration{
			resourceType:  "test_resource",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		&mockMigration{
			resourceType:  "test_resource",
			sourceVersion: "v5",
			targetVersion: "v6",
		},
	}
	
	for _, m := range migrations {
		require.NoError(t, registry.Register(m))
	}

	tests := []struct {
		name          string
		resourceType  string
		sourceVersion string
		targetVersion string
		expectPath    int
		expectError   bool
	}{
		{
			name:          "direct path v4 to v5",
			resourceType:  "test_resource",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectPath:    1,
			expectError:   false,
		},
		{
			name:          "direct path v5 to v6",
			resourceType:  "test_resource",
			sourceVersion: "v5",
			targetVersion: "v6",
			expectPath:    1,
			expectError:   false,
		},
		{
			name:          "multi-hop path v4 to v6",
			resourceType:  "test_resource",
			sourceVersion: "v4",
			targetVersion: "v6",
			expectPath:    2,
			expectError:   false,
		},
		{
			name:          "no path available",
			resourceType:  "test_resource",
			sourceVersion: "v3",
			targetVersion: "v4",
			expectError:   true,
		},
		{
			name:          "non-existent resource",
			resourceType:  "non_existent",
			sourceVersion: "v4",
			targetVersion: "v5",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := pathFinder.FindPath(tt.resourceType, tt.sourceVersion, tt.targetVersion)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, path)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, path)
				assert.Len(t, path, tt.expectPath)
				
				// Verify path is correct
				if len(path) > 0 {
					assert.Equal(t, tt.sourceVersion, path[0].SourceVersion())
					assert.Equal(t, tt.targetVersion, path[len(path)-1].TargetVersion())
				}
			}
		})
	}
}

func TestDefaultRegistry_ConcurrentAccess(t *testing.T) {
	registry := NewDefaultRegistry()
	
	// Test concurrent registration and retrieval
	done := make(chan bool, 20)
	
	// Register 10 migrations concurrently
	for i := 0; i < 10; i++ {
		go func(idx int) {
			migration := &mockMigration{
				resourceType:  "resource",
				sourceVersion: "v1",
				targetVersion: "v" + string(rune('2'+idx)),
			}
			_ = registry.Register(migration)
			done <- true
		}(i)
	}
	
	// Retrieve migrations concurrently
	for i := 0; i < 10; i++ {
		go func(idx int) {
			_, _ = registry.Get("resource", "v1", "v"+string(rune('2'+idx)))
			done <- true
		}(i)
	}
	
	// Wait for all goroutines
	for i := 0; i < 20; i++ {
		<-done
	}
	
	// Verify registry is still consistent
	infos := registry.ListAvailable()
	assert.NotEmpty(t, infos)
}