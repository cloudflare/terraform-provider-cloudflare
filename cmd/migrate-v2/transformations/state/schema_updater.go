package state

// SchemaVersionUpdater creates a transformer that updates the schema version
func SchemaVersionUpdater(newVersion int) StateTransformer {
	return func(state map[string]interface{}) error {
		state["schema_version"] = newVersion
		return nil
	}
}

// DefaultValueSetter creates a transformer that sets default values for missing fields
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