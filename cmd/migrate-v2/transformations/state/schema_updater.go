package state

// SchemaVersionUpdater creates a transformer that updates the schema version
//
// Example usage in state migration:
//   StateTransformer: SchemaVersionUpdater(2)
//
// Transforms state:
//   {
//     "attributes": {
//       "id": "abc123",
//       "name": "test"
//     },
//     "schema_version": 1
//   }
//
// Into:
//   {
//     "attributes": {
//       "id": "abc123",
//       "name": "test"
//     },
//     "schema_version": 2
//   }
func SchemaVersionUpdater(newVersion int) StateTransformer {
	return func(state map[string]interface{}) error {
		state["schema_version"] = newVersion
		return nil
	}
}

// DefaultValueSetter creates a transformer that sets default values for missing fields
//
// Example usage in state migration:
//   StateTransformer: DefaultValueSetter(map[string]interface{}{
//     "enabled": true,
//     "timeout": 30,
//     "retries": 3,
//   })
//
// Transforms state:
//   {
//     "attributes": {
//       "id": "abc123",
//       "name": "test",
//       "timeout": 60  // Existing value preserved
//     }
//   }
//
// Into:
//   {
//     "attributes": {
//       "id": "abc123",
//       "name": "test",
//       "timeout": 60,    // Existing value preserved
//       "enabled": true,  // Default added
//       "retries": 3      // Default added
//     }
//   }
func DefaultValueSetter(defaults map[string]interface{}) StateTransformer {
	return func(state map[string]interface{}) error {
		for key, defaultValue := range defaults {
			if _, exists := state[key]; !exists {
				state[key] = defaultValue
			}
		}
		return nil
	}
}