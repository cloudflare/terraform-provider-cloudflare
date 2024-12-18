package api_shield_operation

func setOperationValues(data *APIShieldOperationModel, env APIShieldOperationResultEnvelope) {
	data.Endpoint = (*env.Result)[0].Endpoint
	data.Host = (*env.Result)[0].Host
	data.Method = (*env.Result)[0].Method
	// TODO:: after schema changes land
	//data.LastUpdated = (*env.Result)[0].LastUpdated
	//data.OperationID = (*env.Result)[0].OperationID
}
