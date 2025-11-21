package queue_consumer

func FixInconsistentCRUDResponses(data *QueueConsumerModel) {
	if data.ScriptName.IsNull() && !data.Script.IsNull() {
		data.ScriptName = data.Script
	}
	if data.Script.IsNull() && !data.ScriptName.IsNull() {
		data.Script = data.ScriptName
	}
}
