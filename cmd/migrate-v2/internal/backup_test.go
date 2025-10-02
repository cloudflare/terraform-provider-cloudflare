package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBackupManager_CreateBackup(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test files
	testFile1 := filepath.Join(tempDir, "main.tf")
	testFile2 := filepath.Join(tempDir, "variables.tfvars")
	require.NoError(t, os.WriteFile(testFile1, []byte("resource \"test\" \"example\" {}"), 0644))
	require.NoError(t, os.WriteFile(testFile2, []byte("var = \"value\""), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)
	require.NotNil(t, backup)

	// Verify backup properties
	assert.NotEmpty(t, backup.ID)
	assert.Equal(t, "v4", backup.FromVersion)
	assert.Equal(t, "v5", backup.ToVersion)
	assert.Len(t, backup.Files, 2)
	assert.Contains(t, backup.Files, "main.tf")
	assert.Contains(t, backup.Files, "variables.tfvars")
	assert.Equal(t, []byte("resource \"test\" \"example\" {}"), backup.Files["main.tf"])
	assert.NotEmpty(t, backup.Checksums["main.tf"])

	// Verify backup was saved to disk
	backupFile := filepath.Join(tempDir, ".terraform-migrate-backups", backup.ID+".backup.json")
	assert.FileExists(t, backupFile)
}

func TestBackupManager_CreateBackupWithState(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test file and state
	testFile := filepath.Join(tempDir, "main.tf")
	stateFile := filepath.Join(tempDir, "terraform.tfstate")
	require.NoError(t, os.WriteFile(testFile, []byte("resource \"test\" \"example\" {}"), 0644))

	stateData := map[string]interface{}{
		"version": 4,
		"resources": []interface{}{
			map[string]interface{}{
				"type": "test",
				"name": "example",
			},
		},
	}
	stateJSON, _ := json.Marshal(stateData)
	require.NoError(t, os.WriteFile(stateFile, stateJSON, 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		StateFile:     stateFile,
		Diagnostics:   []Diagnostic{},
	}

	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)
	require.NotNil(t, backup)

	// Verify state was backed up
	assert.NotNil(t, backup.State)
	assert.Equal(t, float64(4), backup.State["version"])
	assert.Equal(t, stateFile, backup.Metadata["state_file"])
	assert.NotEmpty(t, backup.Checksums["__state__"])
}

func TestBackupManager_Rollback(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create original files
	originalContent := "original content"
	testFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(testFile, []byte(originalContent), 0644))

	// Create backup
	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}
	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Modify file
	modifiedContent := "modified content"
	require.NoError(t, os.WriteFile(testFile, []byte(modifiedContent), 0644))

	// Verify file was modified
	content, _ := os.ReadFile(testFile)
	assert.Equal(t, modifiedContent, string(content))

	// Rollback
	err = bm.Rollback(backup.ID)
	require.NoError(t, err)

	// Verify file was restored
	content, _ = os.ReadFile(testFile)
	assert.Equal(t, originalContent, string(content))
}

func TestBackupManager_RollbackLatest(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test file
	testFile := filepath.Join(tempDir, "main.tf")
	originalContent := "original"
	require.NoError(t, os.WriteFile(testFile, []byte(originalContent), 0644))

	// Create multiple backups
	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	// First backup
	_, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Modify and create second backup
	require.NoError(t, os.WriteFile(testFile, []byte("modified1"), 0644))
	time.Sleep(10 * time.Millisecond) // Ensure different timestamp
	_, err = bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Modify again
	require.NoError(t, os.WriteFile(testFile, []byte("modified2"), 0644))

	// Rollback to latest (should be backup2)
	err = bm.Rollback("latest")
	require.NoError(t, err)

	// Verify rollback to backup2 content
	content, _ := os.ReadFile(testFile)
	assert.Equal(t, "modified1", string(content))
}

func TestBackupManager_ListBackups(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test file
	testFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(testFile, []byte("content"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	// Create multiple backups
	var createdBackups []string
	for i := 0; i < 3; i++ {
		backup, err := bm.CreateBackup(ctx)
		require.NoError(t, err)
		createdBackups = append(createdBackups, backup.ID)
		time.Sleep(10 * time.Millisecond) // Ensure different timestamps
	}

	// List backups
	backups, err := bm.ListBackups()
	require.NoError(t, err)
	assert.Len(t, backups, 3)

	// Verify order (newest first)
	for i := 0; i < len(backups)-1; i++ {
		assert.True(t, backups[i].Timestamp.After(backups[i+1].Timestamp))
	}

	// Verify backup details
	for _, backup := range backups {
		assert.NotEmpty(t, backup.ID)
		assert.Equal(t, "v4", backup.FromVersion)
		assert.Equal(t, "v5", backup.ToVersion)
		assert.Equal(t, 1, backup.FileCount)
		assert.Greater(t, backup.TotalSize, int64(0))
	}
}

func TestBackupManager_GetLatestBackup(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// No backups initially
	latest, err := bm.GetLatestBackup()
	require.NoError(t, err)
	assert.Nil(t, latest)

	// Create test file and backup
	testFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(testFile, []byte("content"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	// Create backups with delays
	_, err = bm.CreateBackup(ctx)
	require.NoError(t, err)
	time.Sleep(10 * time.Millisecond)

	backup2, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Get latest
	latest, err = bm.GetLatestBackup()
	require.NoError(t, err)
	require.NotNil(t, latest)
	assert.Equal(t, backup2.ID, latest.ID)
}

func TestBackupManager_CleanOldBackups(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test file
	testFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(testFile, []byte("content"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	// Create old backup
	oldBackup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Manually set old timestamp (30 days ago)
	oldBackup.Timestamp = time.Now().Add(-30 * 24 * time.Hour)
	require.NoError(t, bm.saveBackup(oldBackup))

	// Create recent backup
	time.Sleep(10 * time.Millisecond)
	_, err = bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Clean backups older than 7 days
	removed, err := bm.CleanOldBackups(7 * 24 * time.Hour)
	require.NoError(t, err)
	assert.Equal(t, 1, removed)

	// Verify only recent backup remains
	backups, err := bm.ListBackups()
	require.NoError(t, err)
	assert.Len(t, backups, 1)
}

func TestBackupManager_VerifyBackupIntegrity(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test file
	testFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(testFile, []byte("content"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Verify integrity passes for valid backup
	err = bm.verifyBackup(backup)
	require.NoError(t, err)

	// Corrupt backup data
	backup.Files["main.tf"] = []byte("corrupted")

	// Verify integrity fails for corrupted backup
	err = bm.verifyBackup(backup)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "checksum mismatch")
}

func TestBackupManager_SkipTerraformDirectory(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create files in main directory
	mainFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(mainFile, []byte("main"), 0644))

	// Create files in .terraform directory (should be skipped)
	terraformDir := filepath.Join(tempDir, ".terraform")
	require.NoError(t, os.MkdirAll(terraformDir, 0755))
	skipFile := filepath.Join(terraformDir, "modules.tf")
	require.NoError(t, os.WriteFile(skipFile, []byte("skip"), 0644))

	// Create files in backup directory (should be skipped)
	backupDir := filepath.Join(tempDir, ".terraform-migrate-backups")
	require.NoError(t, os.MkdirAll(backupDir, 0755))
	backupFile := filepath.Join(backupDir, "old.tf")
	require.NoError(t, os.WriteFile(backupFile, []byte("backup"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Verify only main.tf was backed up
	assert.Len(t, backup.Files, 1)
	assert.Contains(t, backup.Files, "main.tf")
	assert.NotContains(t, backup.Files, skipFile)
	assert.NotContains(t, backup.Files, backupFile)
}

func TestBackupManager_RollbackWithMissingDirectory(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create nested directory structure
	nestedDir := filepath.Join(tempDir, "modules", "vpc")
	require.NoError(t, os.MkdirAll(nestedDir, 0755))
	nestedFile := filepath.Join(nestedDir, "main.tf")
	require.NoError(t, os.WriteFile(nestedFile, []byte("nested content"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Remove the directory
	require.NoError(t, os.RemoveAll(nestedDir))

	// Rollback should recreate the directory
	err = bm.Rollback(backup.ID)
	require.NoError(t, err)

	// Verify file was restored with directory
	content, err := os.ReadFile(nestedFile)
	require.NoError(t, err)
	assert.Equal(t, "nested content", string(content))
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		hasError bool
	}{
		{"24h", 24 * time.Hour, false},
		{"30m", 30 * time.Minute, false},
		{"7d", 7 * 24 * time.Hour, false},
		{"30d", 30 * 24 * time.Hour, false},
		{"1d", 24 * time.Hour, false},
		{"invalid", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			duration, err := ParseDuration(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, duration)
			}
		})
	}
}

func TestCalculateChecksum(t *testing.T) {
	data1 := []byte("test data")
	data2 := []byte("test data")
	data3 := []byte("different data")

	checksum1 := calculateChecksum(data1)
	checksum2 := calculateChecksum(data2)
	checksum3 := calculateChecksum(data3)

	// Same data should produce same checksum
	assert.Equal(t, checksum1, checksum2)

	// Different data should produce different checksum
	assert.NotEqual(t, checksum1, checksum3)

	// Checksum should be hex string
	assert.Regexp(t, "^[a-f0-9]{64}$", checksum1)
}

func TestGenerateShortID(t *testing.T) {
	id1 := generateShortID()
	time.Sleep(10 * time.Millisecond)
	id2 := generateShortID()

	// IDs should be different
	assert.NotEqual(t, id1, id2)

	// IDs should be 8 characters hex
	assert.Len(t, id1, 8)
	assert.Regexp(t, "^[a-f0-9]{8}$", id1)
}

func TestBackupManager_EmergencyRestore(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create backup data
	backup := &Backup{
		ID:        "test-backup",
		Timestamp: time.Now(),
		Files: map[string][]byte{
			"main.tf":      []byte("main content"),
			"variables.tf": []byte("var content"),
		},
	}

	// Perform emergency restore
	err := bm.emergencyRestore(backup)
	require.NoError(t, err)

	// Verify files were created
	mainContent, err := os.ReadFile(filepath.Join(tempDir, "main.tf"))
	require.NoError(t, err)
	assert.Equal(t, "main content", string(mainContent))

	varContent, err := os.ReadFile(filepath.Join(tempDir, "variables.tf"))
	require.NoError(t, err)
	assert.Equal(t, "var content", string(varContent))
}

func TestBackupManager_BackupIDFormat(t *testing.T) {
	tempDir := t.TempDir()
	bm := NewBackupManager(tempDir)

	// Create test file
	testFile := filepath.Join(tempDir, "main.tf")
	require.NoError(t, os.WriteFile(testFile, []byte("content"), 0644))

	ctx := &MigrationContext{
		SourceVersion: "v4",
		TargetVersion: "v5",
		Diagnostics:   []Diagnostic{},
	}

	backup, err := bm.CreateBackup(ctx)
	require.NoError(t, err)

	// Verify ID format: timestamp-shortid
	parts := strings.Split(backup.ID, "-")
	assert.Len(t, parts, 3) // YYYYMMDD-HHMMSS-shortid

	// Verify timestamp format
	timestampPart := strings.Join(parts[:2], "-")
	_, err = time.Parse("20060102-150405", timestampPart)
	assert.NoError(t, err)
}
