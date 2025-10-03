package state

// AttributeRenamer creates a transformer that renames attributes in state
//
// Example usage in state migration:
//   StateTransformer: AttributeRenamer(map[string]string{
//     "old_field": "new_field",
//     "deprecated_name": "current_name",
//   })
//
// Transforms state:
//   {
//     "attributes": {
//       "id": "abc123",
//       "old_field": "value1",
//       "deprecated_name": "value2",
//       "unchanged": "stays"
//     }
//   }
//
// Into:
//   {
//     "attributes": {
//       "id": "abc123", 
//       "new_field": "value1",
//       "current_name": "value2",
//       "unchanged": "stays"
//     }
//   }
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