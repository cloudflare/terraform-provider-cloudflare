package consts

const (
	// Schema key for the API token configuration.
	APITokenSchemaKey = "api_token"

	// Environment variable key for the API token configuration.
	APITokenEnvVarKey = "CLOUDFLARE_API_TOKEN"

	// Schema key for the API key configuration.
	APIKeySchemaKey = "api_key"

	// Environment variable key for the API key configuration.
	APIKeyEnvVarKey = "CLOUDFLARE_API_KEY"

	// Schema key for the email configuration.
	EmailSchemaKey = "email"

	// Environment variable key for the email configuration.
	EmailEnvVarKey = "CLOUDFLARE_EMAIL"

	// Schema key for the API user service key configuration.
	APIUserServiceKeySchemaKey = "api_user_service_key"

	// Environment variable key for the API user service key configuration.
	APIUserServiceKeyEnvVarKey = "CLOUDFLARE_API_USER_SERVICE_KEY"

	// Schema key for the User Agent operator suffix.
	UserAgentOperatorSuffixSchemaKey = "user_agent_operator_suffix"

	// Environment variable key for the User Agent operator suffix.
	UserAgentOperatorSuffixEnvVarKey = "CLOUDFLARE_USER_AGENT_OPERATOR_SUFFIX"

	// Environment variable key for the account ID configuration.
	//
	// Deprecated: Use resource specific account ID values instead.
	AccountIDEnvVarKey = "CLOUDFLARE_ACCOUNT_ID"

	// Schema key for the account ID configuration.
	AccountIDSchemaKey = "account_id"

	// Schema description for `account_id` field.
	AccountIDSchemaDescription = "The account identifier to target for the resource."

	// Schema key for the zone ID configuration.
	ZoneIDSchemaKey = "zone_id"

	// Schema description for `zone_id` field.
	ZoneIDSchemaDescription = "The zone identifier to target for the resource."

	// Schema key for IDs.
	IDSchemaKey = "id"

	// Schema description for all ID fields.
	IDSchemaDescription = "The identifier of this resource."

	// Schema key for the base URL field.
	BaseURLSchemaKey = "base_url"

	// Environment variable key for the client base URL.
	BaseURLEnvVarKey = "CLOUDFLARE_BASE_URL"
)
