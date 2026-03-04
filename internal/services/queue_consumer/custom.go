package queue_consumer

func FixInconsistentCRUDResponses(data *QueueConsumerModel) {
	// The Script field was removed from QueueConsumerModel; ScriptName stands alone.
	// Nothing to sync.
}
