package state

// StateTransformer is a function that transforms state data
type StateTransformer func(state map[string]interface{}) error

// TransformContext holds context information for state transformations
type TransformContext struct {
	Diagnostics []string
}