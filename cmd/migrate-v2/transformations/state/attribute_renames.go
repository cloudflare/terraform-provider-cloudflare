package state

// AttributeRenamer creates a transformer that renames attributes in state
func AttributeRenamer(mappings map[string]string) StateTransformer {
	return func(state map[string]interface{}) error {
		for oldName, newName := range mappings {
			if val, exists := state[oldName]; exists {
				state[newName] = val
				delete(state, oldName)
			}
		}
		return nil
	}
}