package internal

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Backup represents a backup of files and state before migration
type Backup struct {
	ID          string                 `json:"id"`
	Timestamp   time.Time              `json:"timestamp"`
	FromVersion string                 `json:"from_version"`
	ToVersion   string                 `json:"to_version"`
	Files       map[string][]byte      `json:"files"`
	State       map[string]interface{} `json:"state,omitempty"`
	Checksums   map[string]string      `json:"checksums"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
}

// BackupInfo provides summary information about a backup
type BackupInfo struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	FromVersion string    `json:"from_version"`
	ToVersion   string    `json:"to_version"`
	FileCount   int       `json:"file_count"`
	TotalSize   int64     `json:"total_size"`
}

// BackupManager handles creating and managing backups
type BackupManager struct {
	backupDir string
	workDir   string
}

// NewBackupManager creates a new backup manager
func NewBackupManager(workDir string) *BackupManager {
	backupDir := filepath.Join(workDir, ".terraform-migrate-backups")
	return &BackupManager{
		backupDir: backupDir,
		workDir:   workDir,
	}
}

// CreateBackup creates a backup of current configuration and state
func (bm *BackupManager) CreateBackup(ctx *MigrationContext) (*Backup, error) {
	// Ensure backup directory exists
	if err := os.MkdirAll(bm.backupDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Generate backup ID
	timestamp := time.Now()
	id := fmt.Sprintf("%s-%s", timestamp.Format("20060102-150405"), generateShortID())

	backup := &Backup{
		ID:          id,
		Timestamp:   timestamp,
		FromVersion: ctx.SourceVersion,
		ToVersion:   ctx.TargetVersion,
		Files:       make(map[string][]byte),
		Checksums:   make(map[string]string),
		Metadata:    make(map[string]string),
	}

	// Find and backup Terraform configuration files
	tfFiles, err := bm.findTerraformFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to find Terraform files: %w", err)
	}

	for _, file := range tfFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", file, err)
		}

		relPath, err := filepath.Rel(bm.workDir, file)
		if err != nil {
			relPath = file
		}

		backup.Files[relPath] = content
		backup.Checksums[relPath] = calculateChecksum(content)
	}

	// Backup state file if it exists
	if ctx.StateFile != "" {
		if err := bm.backupStateFile(ctx.StateFile, backup); err != nil {
			// State file backup is optional, log warning but continue
			ctx.AddWarning("State backup failed", err.Error(), "backup")
		}
	}

	// Save backup metadata
	backup.Metadata["work_dir"] = bm.workDir
	backup.Metadata["file_count"] = fmt.Sprintf("%d", len(backup.Files))
	backup.Metadata["created_by"] = "terraform-migrate-v2"

	// Save backup to disk
	if err := bm.saveBackup(backup); err != nil {
		return nil, fmt.Errorf("failed to save backup: %w", err)
	}

	return backup, nil
}

// Rollback restores a previous backup
func (bm *BackupManager) Rollback(backupID string) error {
	// Handle "latest" keyword
	if backupID == "latest" {
		latest, err := bm.GetLatestBackup()
		if err != nil {
			return fmt.Errorf("failed to get latest backup: %w", err)
		}
		if latest == nil {
			return fmt.Errorf("no backups found")
		}
		backupID = latest.ID
	}

	// Load the backup
	backup, err := bm.loadBackup(backupID)
	if err != nil {
		return fmt.Errorf("failed to load backup %s: %w", backupID, err)
	}

	// Verify backup integrity
	if err := bm.verifyBackup(backup); err != nil {
		return fmt.Errorf("backup integrity check failed: %w", err)
	}

	// Create a safety backup of current state before rollback
	safetyBackup := &Backup{
		ID:        fmt.Sprintf("rollback-safety-%s", time.Now().Format("20060102-150405")),
		Timestamp: time.Now(),
		Files:     make(map[string][]byte),
		Checksums: make(map[string]string),
		Metadata: map[string]string{
			"type":        "rollback_safety",
			"original_id": backupID,
		},
	}

	// Backup current files before overwriting
	for path := range backup.Files {
		fullPath := filepath.Join(bm.workDir, path)
		if content, err := os.ReadFile(fullPath); err == nil {
			safetyBackup.Files[path] = content
			safetyBackup.Checksums[path] = calculateChecksum(content)
		}
	}

	if err := bm.saveBackup(safetyBackup); err != nil {
		// Log warning but continue with rollback
		fmt.Printf("Warning: Failed to create safety backup: %v\n", err)
	}

	// Restore files from backup
	var restoreErrors []error
	for path, content := range backup.Files {
		fullPath := filepath.Join(bm.workDir, path)

		// Ensure directory exists
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			restoreErrors = append(restoreErrors, fmt.Errorf("failed to create directory %s: %w", dir, err))
			continue
		}

		// Write file
		if err := os.WriteFile(fullPath, content, 0644); err != nil {
			restoreErrors = append(restoreErrors, fmt.Errorf("failed to restore %s: %w", path, err))
		}
	}

	// If there were errors, attempt to restore from safety backup
	if len(restoreErrors) > 0 {
		fmt.Println("Errors occurred during rollback, attempting to restore from safety backup...")
		if err := bm.emergencyRestore(safetyBackup); err != nil {
			return fmt.Errorf("rollback failed and emergency restore failed: %w", err)
		}
		return fmt.Errorf("rollback failed with %d errors", len(restoreErrors))
	}

	// Restore state if present
	if len(backup.State) > 0 && backup.Metadata["state_file"] != "" {
		statePath := backup.Metadata["state_file"]
		stateBytes, err := json.MarshalIndent(backup.State, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal state: %w", err)
		}
		if err := os.WriteFile(statePath, stateBytes, 0644); err != nil {
			return fmt.Errorf("failed to restore state: %w", err)
		}
	}

	return nil
}

// ListBackups returns information about all available backups
func (bm *BackupManager) ListBackups() ([]BackupInfo, error) {
	if _, err := os.Stat(bm.backupDir); os.IsNotExist(err) {
		return []BackupInfo{}, nil
	}

	entries, err := os.ReadDir(bm.backupDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}

	var backups []BackupInfo
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".backup.json") {
			continue
		}

		// backupPath := filepath.Join(bm.backupDir, entry.Name())
		backup, err := bm.loadBackup(strings.TrimSuffix(entry.Name(), ".backup.json"))
		if err != nil {
			continue // Skip invalid backups
		}

		info := BackupInfo{
			ID:          backup.ID,
			Timestamp:   backup.Timestamp,
			FromVersion: backup.FromVersion,
			ToVersion:   backup.ToVersion,
			FileCount:   len(backup.Files),
		}

		// Calculate total size
		for _, content := range backup.Files {
			info.TotalSize += int64(len(content))
		}

		backups = append(backups, info)
	}

	// Sort by timestamp, newest first
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Timestamp.After(backups[j].Timestamp)
	})

	return backups, nil
}

// GetLatestBackup returns the most recent backup
func (bm *BackupManager) GetLatestBackup() (*BackupInfo, error) {
	backups, err := bm.ListBackups()
	if err != nil {
		return nil, err
	}
	if len(backups) == 0 {
		return nil, nil
	}
	return &backups[0], nil
}

// CleanOldBackups removes backups older than the specified duration
func (bm *BackupManager) CleanOldBackups(olderThan time.Duration) (int, error) {
	backups, err := bm.ListBackups()
	if err != nil {
		return 0, err
	}

	cutoff := time.Now().Add(-olderThan)
	removed := 0

	for _, backup := range backups {
		if backup.Timestamp.Before(cutoff) {
			backupFile := filepath.Join(bm.backupDir, backup.ID+".backup.json")
			if err := os.Remove(backupFile); err != nil {
				// Log but continue
				fmt.Printf("Warning: Failed to remove backup %s: %v\n", backup.ID, err)
			} else {
				removed++
			}
		}
	}

	return removed, nil
}

// Private helper methods

func (bm *BackupManager) findTerraformFiles() ([]string, error) {
	var files []string

	err := filepath.Walk(bm.workDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip backup directory
		if strings.Contains(path, bm.backupDir) {
			return nil
		}

		// Skip .terraform directory
		if strings.Contains(path, ".terraform") {
			return nil
		}

		// Include .tf and .tfvars files
		if strings.HasSuffix(path, ".tf") || strings.HasSuffix(path, ".tfvars") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func (bm *BackupManager) backupStateFile(statePath string, backup *Backup) error {
	content, err := os.ReadFile(statePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // State file doesn't exist, that's okay
		}
		return err
	}

	var state map[string]interface{}
	if err := json.Unmarshal(content, &state); err != nil {
		return fmt.Errorf("failed to parse state file: %w", err)
	}

	backup.State = state
	backup.Metadata["state_file"] = statePath
	backup.Checksums["__state__"] = calculateChecksum(content)

	return nil
}

func (bm *BackupManager) saveBackup(backup *Backup) error {
	backupPath := filepath.Join(bm.backupDir, backup.ID+".backup.json")

	data, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal backup: %w", err)
	}

	return os.WriteFile(backupPath, data, 0644)
}

func (bm *BackupManager) loadBackup(backupID string) (*Backup, error) {
	backupPath := filepath.Join(bm.backupDir, backupID+".backup.json")

	data, err := os.ReadFile(backupPath)
	if err != nil {
		return nil, err
	}

	var backup Backup
	if err := json.Unmarshal(data, &backup); err != nil {
		return nil, fmt.Errorf("failed to unmarshal backup: %w", err)
	}

	return &backup, nil
}

func (bm *BackupManager) verifyBackup(backup *Backup) error {
	// Verify checksums
	for path, expectedChecksum := range backup.Checksums {
		if path == "__state__" {
			continue // State checksum handled separately
		}

		content, exists := backup.Files[path]
		if !exists {
			return fmt.Errorf("file %s missing from backup", path)
		}

		actualChecksum := calculateChecksum(content)
		if actualChecksum != expectedChecksum {
			return fmt.Errorf("checksum mismatch for %s", path)
		}
	}

	return nil
}

func (bm *BackupManager) emergencyRestore(backup *Backup) error {
	// Best-effort restore, ignore errors
	for path, content := range backup.Files {
		fullPath := filepath.Join(bm.workDir, path)
		_ = os.WriteFile(fullPath, content, 0644)
	}
	return nil
}

// Helper functions

func calculateChecksum(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func generateShortID() string {
	// Generate a short random ID
	h := sha256.New()
	h.Write([]byte(time.Now().String()))
	hash := h.Sum(nil)
	return fmt.Sprintf("%x", hash[:4])
}

// ParseDuration parses duration strings like "30d", "24h", etc.
func ParseDuration(s string) (time.Duration, error) {
	// Handle days specially
	if strings.HasSuffix(s, "d") {
		days := strings.TrimSuffix(s, "d")
		var d int
		if _, err := fmt.Sscanf(days, "%d", &d); err != nil {
			return 0, err
		}
		return time.Duration(d) * 24 * time.Hour, nil
	}

	// Otherwise use standard duration parsing
	return time.ParseDuration(s)
}
