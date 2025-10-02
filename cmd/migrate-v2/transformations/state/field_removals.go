package state

// FieldRemover creates a transformer that removes fields from state
func FieldRemover(fields ...string) StateTransformer {
	return func(state map[string]interface{}) error {
		for _, field := range fields {
			delete(state, field)
		}
		return nil
	}
}