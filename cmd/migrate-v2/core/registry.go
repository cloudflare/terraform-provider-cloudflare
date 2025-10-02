package core

import (
	"fmt"
	"sync"
)

// DefaultRegistry is the default implementation of MigrationRegistry
type DefaultRegistry struct {
	migrations map[string]ResourceMigration
	mu         sync.RWMutex
}

// NewDefaultRegistry creates a new default registry
func NewDefaultRegistry() *DefaultRegistry {
	return &DefaultRegistry{
		migrations: make(map[string]ResourceMigration),
	}
}

// Register adds a migration to the registry
func (r *DefaultRegistry) Register(migration ResourceMigration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := r.makeKey(migration.ResourceType(), migration.SourceVersion(), migration.TargetVersion())
	
	if _, exists := r.migrations[key]; exists {
		return fmt.Errorf("migration already registered for %s", key)
	}

	r.migrations[key] = migration
	return nil
}

// Get retrieves a migration for a specific resource type and version transition
func (r *DefaultRegistry) Get(resourceType, sourceVersion, targetVersion string) (ResourceMigration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := r.makeKey(resourceType, sourceVersion, targetVersion)
	
	migration, exists := r.migrations[key]
	if !exists {
		return nil, fmt.Errorf("no migration found for %s from %s to %s", resourceType, sourceVersion, targetVersion)
	}

	return migration, nil
}

// GetPath finds a migration path from source to target version
func (r *DefaultRegistry) GetPath(resourceType, sourceVersion, targetVersion string) ([]ResourceMigration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// For now, only support direct migrations
	// Future enhancement: implement path finding for multi-hop migrations
	migration, err := r.Get(resourceType, sourceVersion, targetVersion)
	if err != nil {
		return nil, err
	}

	return []ResourceMigration{migration}, nil
}

// ListAvailable lists all registered migrations
func (r *DefaultRegistry) ListAvailable() []MigrationInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var infos []MigrationInfo
	
	for _, migration := range r.migrations {
		infos = append(infos, MigrationInfo{
			ResourceType:  migration.ResourceType(),
			SourceVersion: migration.SourceVersion(),
			TargetVersion: migration.TargetVersion(),
		})
	}

	return infos
}

// makeKey creates a unique key for a migration
func (r *DefaultRegistry) makeKey(resourceType, sourceVersion, targetVersion string) string {
	return fmt.Sprintf("%s:%s->%s", resourceType, sourceVersion, targetVersion)
}

// PathFinder implements migration path discovery
type PathFinder struct {
	registry *DefaultRegistry
}

// NewPathFinder creates a new path finder
func NewPathFinder(registry *DefaultRegistry) *PathFinder {
	return &PathFinder{
		registry: registry,
	}
}

// FindPath finds the shortest migration path between versions
func (pf *PathFinder) FindPath(resourceType, sourceVersion, targetVersion string) ([]ResourceMigration, error) {
	// This is a simplified implementation
	// For multi-hop migrations, we would implement a graph search algorithm
	
	// Try direct migration first
	if migration, err := pf.registry.Get(resourceType, sourceVersion, targetVersion); err == nil {
		return []ResourceMigration{migration}, nil
	}

	// For v4 -> v6, try v4 -> v5 -> v6
	if sourceVersion == "v4" && targetVersion == "v6" {
		path := []ResourceMigration{}
		
		// Get v4 -> v5
		if m1, err := pf.registry.Get(resourceType, "v4", "v5"); err == nil {
			path = append(path, m1)
			
			// Get v5 -> v6
			if m2, err := pf.registry.Get(resourceType, "v5", "v6"); err == nil {
				path = append(path, m2)
				return path, nil
			}
		}
	}

	return nil, fmt.Errorf("no migration path found from %s to %s for %s", sourceVersion, targetVersion, resourceType)
}