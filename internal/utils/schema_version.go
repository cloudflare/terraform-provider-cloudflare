package utils

import "os"

// GetSchemaVersion returns the appropriate schema version based on TF_MIG_TEST environment variable.
//
// This function allows controlled rollout of StateUpgrader migrations:
//   - During development/testing: Set TF_MIG_TEST=1 to enable migrations (returns postMigration version)
//   - In production: StateUpgraders remain dormant (returns preMigration version)
//   - For coordinated release: Remove this wrapper and set Version directly to enable all migrations at once
//
// Parameters:
//   - preMigration: The version to use when migrations are disabled (typically 0)
//   - postMigration: The version to use when migrations are enabled (typically 500)
//
// Example usage:
//
//	Version: GetSchemaVersion(0, 500)  // Returns 0 normally, 500 when TF_MIG_TEST=1
//
// TODO:: Remove once the migration work is complete
func GetSchemaVersion(preMigration, postMigration int64) int64 {
	if os.Getenv("TF_MIG_TEST") == "" {
		return preMigration
	}
	return postMigration
}
