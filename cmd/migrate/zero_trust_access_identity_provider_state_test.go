package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestTransformZeroTrustAccessIdentityProviderStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "Transform config from array to object",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "GitHub OAuth",
									"type":       "github",
									"config": []interface{}{
										map[string]interface{}{
											"client_id":     "github-client-id",
											"client_secret": "github-secret",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "GitHub OAuth",
									"type":       "github",
									"config": map[string]interface{}{
										"client_id":     "github-client-id",
										"client_secret": "github-secret",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Transform scim_config from array to object",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "Azure AD",
									"type":       "azureAD",
									"config": []interface{}{
										map[string]interface{}{
											"client_id": "azure-client-id",
										},
									},
									"scim_config": []interface{}{
										map[string]interface{}{
											"enabled":                   true,
											"secret":                    "scim-secret",
											"user_deprovision":          true,
											"seat_deprovision":          false,
											"group_member_deprovision":  true,
											"identity_update_behavior":  "automatic",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "Azure AD",
									"type":       "azureAD",
									"config": map[string]interface{}{
										"client_id": "azure-client-id",
									},
									"scim_config": map[string]interface{}{
										"enabled":                  true,
										"secret":                   "scim-secret",
										"user_deprovision":         true,
										"seat_deprovision":         false,
										// group_member_deprovision should be removed
										"identity_update_behavior": "automatic",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Transform idp_public_cert to idp_public_certs",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "SAML Provider",
									"type":       "saml",
									"config": []interface{}{
										map[string]interface{}{
											"issuer_url":      "https://saml.example.com",
											"sso_target_url":  "https://saml.example.com/sso",
											"idp_public_cert": "MIIDpDCCAoygAwIBAgIGAV...",
											"attributes":      []interface{}{"email", "name"},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "SAML Provider",
									"type":       "saml",
									"config": map[string]interface{}{
										"issuer_url":       "https://saml.example.com",
										"sso_target_url":   "https://saml.example.com/sso",
										"idp_public_certs": []interface{}{"MIIDpDCCAoygAwIBAgIGAV..."},
										"attributes":       []interface{}{"email", "name"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Remove deprecated api_token field",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "Custom Provider",
									"type":       "custom",
									"config": []interface{}{
										map[string]interface{}{
											"client_id": "custom-client-id",
											"api_token": "deprecated-token",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "Custom Provider",
									"type":       "custom",
									"config": map[string]interface{}{
										"client_id": "custom-client-id",
										// api_token should be removed
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Handle empty config array to empty object",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "OneTimePin",
									"type":       "onetimepin",
									"config":     []interface{}{},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "OneTimePin",
									"type":       "onetimepin",
									"config":     map[string]interface{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Remove empty/false fields from config",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "OAuth Provider",
									"type":       "oauth",
									"config": []interface{}{
										map[string]interface{}{
											"client_id":                 "oauth-client-id",
											"client_secret":             "oauth-secret",
											"sign_request":              false,
											"conditional_access_enabled": false,
											"support_groups":            false,
											"pkce_enabled":              false,
											"apps_domain":               "",
											"auth_url":                  "",
											"claims":                    []interface{}{},
											"scopes":                    []interface{}{},
											"prompt":                    nil,
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "OAuth Provider",
									"type":       "oauth",
									"config": map[string]interface{}{
										"client_id":     "oauth-client-id",
										"client_secret": "oauth-secret",
										// All false booleans, empty strings, empty arrays, and null values should be removed
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Handle resource type rename from access_identity_provider",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "GitHub",
									"type":       "github",
									"config": []interface{}{
										map[string]interface{}{
											"client_id": "github-id",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_zero_trust_access_identity_provider",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"id":         "test-id",
									"account_id": "test-account",
									"name":       "GitHub",
									"type":       "github",
									"config": map[string]interface{}{
										"client_id": "github-id",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Convert input to JSON
			inputJSON, err := json.Marshal(test.input)
			require.NoError(t, err)

			// Transform the state
			outputJSON, err := transformStateJSON(inputJSON)
			require.NoError(t, err)

			// Parse the output
			var output map[string]interface{}
			err = json.Unmarshal(outputJSON, &output)
			require.NoError(t, err)

			// Get the specific resource instance for comparison
			outputResource := gjson.GetBytes(outputJSON, "resources.0")
			expectedJSON, _ := json.Marshal(test.expected)
			expectedResource := gjson.GetBytes(expectedJSON, "resources.0")

			// Compare the transformed resource
			assert.JSONEq(t, expectedResource.String(), outputResource.String(),
				"State transformation did not produce expected result")
		})
	}
}

func TestTransformZeroTrustAccessIdentityProviderStateJSON_ComplexScenarios(t *testing.T) {
	t.Run("Full Azure AD configuration with SCIM", func(t *testing.T) {
		input := map[string]interface{}{
			"resources": []interface{}{
				map[string]interface{}{
					"type": "cloudflare_access_identity_provider",
					"instances": []interface{}{
						map[string]interface{}{
							"attributes": map[string]interface{}{
								"id":         "azure-id",
								"account_id": "test-account",
								"zone_id":    "test-zone",
								"name":       "Azure AD with SCIM",
								"type":       "azureAD",
								"config": []interface{}{
									map[string]interface{}{
										"client_id":                 "azure-client-id",
										"client_secret":             "azure-secret",
										"directory_id":              "azure-directory-id",
										"conditional_access_enabled": true,
										"support_groups":            true,
										"api_token":                 "deprecated-api-token",
										"claims":                    []interface{}{"email", "name", "groups"},
										"scopes":                    []interface{}{"openid", "profile", "email"},
										"email_claim_name":          "email",
									},
								},
								"scim_config": []interface{}{
									map[string]interface{}{
										"enabled":                   true,
										"secret":                    "scim-secret-value",
										"user_deprovision":          true,
										"seat_deprovision":          true,
										"group_member_deprovision":  true,
										"identity_update_behavior":  "reauth",
									},
								},
							},
						},
					},
				},
			},
		}

		inputJSON, err := json.Marshal(input)
		require.NoError(t, err)

		outputJSON, err := transformStateJSON(inputJSON)
		require.NoError(t, err)

		// Verify the transformation
		output := gjson.ParseBytes(outputJSON)
		
		// Check resource type was renamed
		assert.Equal(t, "cloudflare_zero_trust_access_identity_provider", 
			output.Get("resources.0.type").String())
		
		// Check config is now an object, not an array
		assert.True(t, output.Get("resources.0.instances.0.attributes.config").IsObject())
		
		// Check scim_config is now an object, not an array
		assert.True(t, output.Get("resources.0.instances.0.attributes.scim_config").IsObject())
		
		// Check api_token was removed
		assert.False(t, output.Get("resources.0.instances.0.attributes.config.api_token").Exists())
		
		// Check group_member_deprovision was removed
		assert.False(t, output.Get("resources.0.instances.0.attributes.scim_config.group_member_deprovision").Exists())
		
		// Check other fields are preserved
		assert.Equal(t, "azure-client-id", 
			output.Get("resources.0.instances.0.attributes.config.client_id").String())
		assert.Equal(t, true,
			output.Get("resources.0.instances.0.attributes.config.conditional_access_enabled").Bool())
		assert.Equal(t, "reauth",
			output.Get("resources.0.instances.0.attributes.scim_config.identity_update_behavior").String())
	})

	t.Run("SAML provider with multiple certificates", func(t *testing.T) {
		input := map[string]interface{}{
			"resources": []interface{}{
				map[string]interface{}{
					"type": "cloudflare_zero_trust_access_identity_provider",
					"instances": []interface{}{
						map[string]interface{}{
							"attributes": map[string]interface{}{
								"id":         "saml-id",
								"account_id": "test-account",
								"name":       "SAML SSO",
								"type":       "saml",
								"config": []interface{}{
									map[string]interface{}{
										"issuer_url":      "https://saml.example.com",
										"sso_target_url":  "https://saml.example.com/sso",
										"idp_public_cert": "CERT_CONTENT_HERE",
										"sign_request":    true,
										"attributes":      []interface{}{"email", "name", "groups"},
										"email_attribute_name": "email",
									},
								},
							},
						},
					},
				},
			},
		}

		inputJSON, err := json.Marshal(input)
		require.NoError(t, err)

		outputJSON, err := transformStateJSON(inputJSON)
		require.NoError(t, err)

		output := gjson.ParseBytes(outputJSON)
		
		// Check idp_public_cert was renamed to idp_public_certs and converted to array
		assert.False(t, output.Get("resources.0.instances.0.attributes.config.idp_public_cert").Exists())
		assert.True(t, output.Get("resources.0.instances.0.attributes.config.idp_public_certs").IsArray())
		assert.Equal(t, "CERT_CONTENT_HERE",
			output.Get("resources.0.instances.0.attributes.config.idp_public_certs.0").String())
	})
}