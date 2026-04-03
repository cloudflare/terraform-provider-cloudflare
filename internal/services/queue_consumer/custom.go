package queue_consumer

func FixInconsistentCRUDResponses(data *QueueConsumerModel) {
	// The API returns "script" in the response, but we expose "script_name" in the schema.
	// Copy the script value to script_name to ensure state consistency.
	if !data.Script.IsNull() && !data.Script.IsUnknown() && data.Script.ValueString() != "" {
		data.ScriptName = data.Script
	}
}
