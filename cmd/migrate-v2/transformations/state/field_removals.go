package state

// FieldRemover creates a transformer that removes fields from state
//
// Example usage in state migration:
//   StateTransformer: FieldRemover(
//     "deprecated_field",
//     "obsolete_setting",
//     "removed_feature",
//   )
//
// Transforms state:
//   {
//     "attributes": {
//       "id": "abc123",
//       "name": "test",
//       "deprecated_field": "old_value",
//       "obsolete_setting": true,
//       "removed_feature": {"data": "value"},
//       "keep_field": "stays"
//     }
//   }
//
// Into:
//   {
//     "attributes": {
//       "id": "abc123",
//       "name": "test",
//       "keep_field": "stays"
//     }
//   }
func FieldRemover(fields ...string) StateTransformer {
	return func(state map[string]interface{}) error {
		for _, field := range fields {
			delete(state, field)
		}
		return nil
	}
}