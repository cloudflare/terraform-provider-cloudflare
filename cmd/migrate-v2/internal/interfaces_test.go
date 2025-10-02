package internal

import (
	"errors"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock implementation of ResourceMigration for testing interfaces
type mockInterfaceMigration struct {
	resourceType   string
	sourceVersion  string
	targetVersion  string
	migrateError   error
	validateError  error
	stateModified  bool
	configModified bool
}

func (m *mockInterfaceMigration) ResourceType() string {
	return m.resourceType
}

func (m *mockInterfaceMigration) SourceVersion() string {
	return m.sourceVersion
}

func (m *mockInterfaceMigration) TargetVersion() string {
	return m.targetVersion
}

func (m *mockInterfaceMigration) MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error {
	m.configModified = true
	if m.migrateError != nil {
		return m.migrateError
	}
	// Simulate adding an attribute
	block.Body().SetAttributeRaw("migrated", hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("true")},
	})
	return nil
}

func (m *mockInterfaceMigration) MigrateState(state map[string]interface{}, ctx *MigrationContext) error {
	m.stateModified = true
	if m.migrateError != nil {
		return m.migrateError
	}
	state["migrated"] = true
	state["schema_version"] = 1
	return nil
}

func (m *mockInterfaceMigration) Validate(block *hclwrite.Block) error {
	return m.validateError
}

// Mock implementation of MigrationRegistry
type mockRegistry struct {
	migrations map[string]ResourceMigration
	errors     map[string]error
}

func newMockRegistry() *mockRegistry {
	return &mockRegistry{
		migrations: make(map[string]ResourceMigration),
		errors:     make(map[string]error),
	}
}

func (r *mockRegistry) makeKey(resourceType, sourceVersion, targetVersion string) string {
	return resourceType + ":" + sourceVersion + ":" + targetVersion
}

func (r *mockRegistry) Register(migration ResourceMigration) error {
	if migration == nil {
		return errors.New("migration cannot be nil")
	}
	key := r.makeKey(migration.ResourceType(), migration.SourceVersion(), migration.TargetVersion())
	if err, exists := r.errors[key]; exists {
		return err
	}
	r.migrations[key] = migration
	return nil
}

func (r *mockRegistry) Get(resourceType, sourceVersion, targetVersion string) (ResourceMigration, error) {
	key := r.makeKey(resourceType, sourceVersion, targetVersion)
	if migration, exists := r.migrations[key]; exists {
		return migration, nil
	}
	return nil, errors.New("migration not found")
}

func (r *mockRegistry) GetPath(resourceType, sourceVersion, targetVersion string) ([]ResourceMigration, error) {
	// Simple implementation: try to find direct migration
	direct, err := r.Get(resourceType, sourceVersion, targetVersion)
	if err == nil {
		return []ResourceMigration{direct}, nil
	}

	// Try to find a multi-step path (v4 -> v5 -> v6)
	if sourceVersion == "v4" && targetVersion == "v6" {
		step1, err1 := r.Get(resourceType, "v4", "v5")
		step2, err2 := r.Get(resourceType, "v5", "v6")
		if err1 == nil && err2 == nil {
			return []ResourceMigration{step1, step2}, nil
		}
	}

	return nil, errors.New("no migration path found")
}

func (r *mockRegistry) ListAvailable() []MigrationInfo {
	var infos []MigrationInfo
	for _, migration := range r.migrations {
		infos = append(infos, MigrationInfo{
			ResourceType:  migration.ResourceType(),
			SourceVersion: migration.SourceVersion(),
			TargetVersion: migration.TargetVersion(),
			Description:   "Mock migration",
		})
	}
	return infos
}

func TestResourceMigration_Interface(t *testing.T) {
	migration := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}

	// Test interface methods
	assert.Equal(t, "cloudflare_test_resource", migration.ResourceType())
	assert.Equal(t, "v4", migration.SourceVersion())
	assert.Equal(t, "v5", migration.TargetVersion())

	// Test MigrateConfig
	block := hclwrite.NewBlock("resource", []string{"cloudflare_test_resource", "test"})
	ctx := NewMigrationContext()

	err := migration.MigrateConfig(block, ctx)
	require.NoError(t, err)
	assert.True(t, migration.configModified)

	// Verify attribute was added
	attr := block.Body().GetAttribute("migrated")
	assert.NotNil(t, attr)

	// Test MigrateState
	state := make(map[string]interface{})
	err = migration.MigrateState(state, ctx)
	require.NoError(t, err)
	assert.True(t, migration.stateModified)
	assert.Equal(t, true, state["migrated"])
	assert.Equal(t, 1, state["schema_version"])
}

func TestResourceMigration_Errors(t *testing.T) {
	expectedError := errors.New("migration failed")
	migration := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
		migrateError:  expectedError,
		validateError: errors.New("validation failed"),
	}

	// Test MigrateConfig with error
	block := hclwrite.NewBlock("resource", []string{"cloudflare_test_resource", "test"})
	ctx := NewMigrationContext()

	err := migration.MigrateConfig(block, ctx)
	assert.Equal(t, expectedError, err)

	// Test MigrateState with error
	state := make(map[string]interface{})
	err = migration.MigrateState(state, ctx)
	assert.Equal(t, expectedError, err)

	// Test Validate with error
	err = migration.Validate(block)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

func TestDiagnosticSeverity(t *testing.T) {
	// Test severity constants
	assert.Equal(t, DiagnosticSeverity(0), DiagnosticSeverityError)
	assert.Equal(t, DiagnosticSeverity(1), DiagnosticSeverityWarning)
	assert.Equal(t, DiagnosticSeverity(2), DiagnosticSeverityInfo)

	// Test diagnostic creation
	diag := Diagnostic{
		Severity: DiagnosticSeverityError,
		Summary:  "Test error",
		Detail:   "This is a test error",
		Resource: "cloudflare_test_resource",
		Line:     10,
		Column:   5,
	}

	assert.Equal(t, DiagnosticSeverityError, diag.Severity)
	assert.Equal(t, "Test error", diag.Summary)
	assert.Equal(t, "This is a test error", diag.Detail)
	assert.Equal(t, "cloudflare_test_resource", diag.Resource)
	assert.Equal(t, 10, diag.Line)
	assert.Equal(t, 5, diag.Column)
}

func TestMigrationMetrics(t *testing.T) {
	metrics := &MigrationMetrics{
		TotalResources:       10,
		MigratedResources:    7,
		FailedResources:      2,
		WarningCount:         5,
		ManualMigrationCount: 1,
	}

	assert.Equal(t, 10, metrics.TotalResources)
	assert.Equal(t, 7, metrics.MigratedResources)
	assert.Equal(t, 2, metrics.FailedResources)
	assert.Equal(t, 5, metrics.WarningCount)
	assert.Equal(t, 1, metrics.ManualMigrationCount)

	// Test that we can calculate success rate
	successRate := float64(metrics.MigratedResources) / float64(metrics.TotalResources) * 100
	assert.Equal(t, 70.0, successRate)
}

func TestMigrationOptions(t *testing.T) {
	opts := MigrationOptions{
		DryRun:       true,
		Verbose:      true,
		Parallel:     true,
		Workers:      4,
		Preview:      true,
		Backup:       true,
		AutoRollback: false,
		WorkingDir:   "/test/dir",
	}

	assert.True(t, opts.DryRun)
	assert.True(t, opts.Verbose)
	assert.True(t, opts.Parallel)
	assert.Equal(t, 4, opts.Workers)
	assert.True(t, opts.Preview)
	assert.True(t, opts.Backup)
	assert.False(t, opts.AutoRollback)
	assert.Equal(t, "/test/dir", opts.WorkingDir)
}

func TestMigrationRegistry_Register(t *testing.T) {
	registry := newMockRegistry()

	migration1 := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}

	// Test successful registration
	err := registry.Register(migration1)
	require.NoError(t, err)

	// Test registration of nil migration
	err = registry.Register(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be nil")

	// Test duplicate registration (should overwrite)
	migration2 := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}
	err = registry.Register(migration2)
	require.NoError(t, err)
}

func TestMigrationRegistry_Get(t *testing.T) {
	registry := newMockRegistry()

	migration := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}

	// Register migration
	err := registry.Register(migration)
	require.NoError(t, err)

	// Test successful get
	retrieved, err := registry.Get("cloudflare_test_resource", "v4", "v5")
	require.NoError(t, err)
	assert.NotNil(t, retrieved)
	assert.Equal(t, "cloudflare_test_resource", retrieved.ResourceType())

	// Test get non-existent migration
	_, err = registry.Get("cloudflare_test_resource", "v3", "v4")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestMigrationRegistry_GetPath(t *testing.T) {
	registry := newMockRegistry()

	// Register v4->v5 and v5->v6 migrations
	migration1 := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}
	migration2 := &mockInterfaceMigration{
		resourceType:  "cloudflare_test_resource",
		sourceVersion: "v5",
		targetVersion: "v6",
	}

	require.NoError(t, registry.Register(migration1))
	require.NoError(t, registry.Register(migration2))

	// Test direct path (single migration)
	path, err := registry.GetPath("cloudflare_test_resource", "v4", "v5")
	require.NoError(t, err)
	assert.Len(t, path, 1)
	assert.Equal(t, "v5", path[0].TargetVersion())

	// Test multi-step path
	path, err = registry.GetPath("cloudflare_test_resource", "v4", "v6")
	require.NoError(t, err)
	assert.Len(t, path, 2)
	assert.Equal(t, "v5", path[0].TargetVersion())
	assert.Equal(t, "v6", path[1].TargetVersion())

	// Test no path found
	_, err = registry.GetPath("cloudflare_test_resource", "v2", "v3")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no migration path found")
}

func TestMigrationRegistry_ListAvailable(t *testing.T) {
	registry := newMockRegistry()

	// Register multiple migrations
	migrations := []*mockInterfaceMigration{
		{
			resourceType:  "cloudflare_test_resource",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		{
			resourceType:  "cloudflare_other_resource",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		{
			resourceType:  "cloudflare_test_resource",
			sourceVersion: "v5",
			targetVersion: "v6",
		},
	}

	for _, m := range migrations {
		require.NoError(t, registry.Register(m))
	}

	// List all available migrations
	available := registry.ListAvailable()
	assert.Len(t, available, 3)

	// Check that all migrations are listed
	resourceTypes := make(map[string]bool)
	for _, info := range available {
		resourceTypes[info.ResourceType] = true
		assert.NotEmpty(t, info.Description)
		assert.NotEmpty(t, info.SourceVersion)
		assert.NotEmpty(t, info.TargetVersion)
	}

	assert.True(t, resourceTypes["cloudflare_test_resource"])
	assert.True(t, resourceTypes["cloudflare_other_resource"])
}

func TestMigrationInfo(t *testing.T) {
	info := MigrationInfo{
		ResourceType:  "cloudflare_test_resource",
		SourceVersion: "v4",
		TargetVersion: "v5",
		Description:   "Migrate test resource from v4 to v5",
	}

	assert.Equal(t, "cloudflare_test_resource", info.ResourceType)
	assert.Equal(t, "v4", info.SourceVersion)
	assert.Equal(t, "v5", info.TargetVersion)
	assert.Equal(t, "Migrate test resource from v4 to v5", info.Description)
}

func TestResourceMigrationInterface(t *testing.T) {
	// Test that mockInterfaceMigration implements ResourceMigration
	var rm ResourceMigration = &mockInterfaceMigration{
		resourceType:  "cloudflare_test",
		sourceVersion: "v4",
		targetVersion: "v5",
	}

	assert.Equal(t, "cloudflare_test", rm.ResourceType())
	assert.Equal(t, "v4", rm.SourceVersion())
	assert.Equal(t, "v5", rm.TargetVersion())
}
