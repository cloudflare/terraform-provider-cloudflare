package v500

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// V5TargetSchema is set by parent package init() to provide the target schema.
// This avoids circular imports between the migration package and parent package.
var V5TargetSchema func(context.Context) schema.Schema

// UpgradeFromVersion0 handles state upgrades from schema_version=0 to version=500.
//
// IMPORTANT: Both v4 SDKv2 provider AND v5.16.0 (dormant) have schema_version=0.
// PriorSchema is nil because v4 and v5 have incompatible schemas:
// - v4 state: cors_headers, saas_app are ARRAYS (ListNestedBlock)
// - v5.16.0 state: these fields are OBJECTS (SingleNestedAttribute)
//
// Detection strategy: Parse raw JSON and check if cors_headers is array (v4) or object (v5).
func UpgradeFromVersion0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from version=0 (detecting v4 vs v5.16.0 format)")

	// Detect v4 vs v5 format by inspecting raw JSON
	isV4, err := detectV4State(req)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect state format",
			fmt.Sprintf("Could not determine v4 vs v5.16.0 state format: %s", err))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 state format, performing transformation")
		upgradeFromV4Internal(ctx, req, resp)
	} else {
		tflog.Info(ctx, "Detected v5.16.0+ state format, no-op upgrade")
		unmarshalV5StateToResponse(ctx, req.RawState, resp)
	}
}

// detectV4State checks if the state is v4 format.
// v4: has cors_headers or saas_app as array [], OR has v4-only fields like domain_type
// v5: has cors_headers or saas_app as object {}, OR has v5-only fields like allow_iframe
func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			// First, check for v4-only fields that don't exist in v5 schema.
			// If any of these exist, it's definitely v4 format.
			v4OnlyFields := []string{
				"domain_type", // Removed in v5
			}
			for _, field := range v4OnlyFields {
				if _, ok := rawJSON[field]; ok {
					return true, nil // v4 format - has v4-only field
				}
			}

			// Check for v5-only fields that don't exist in v4 schema.
			// If any of these exist, it's definitely v5 format.
			v5OnlyFields := []string{
				"allow_iframe",               // Added in v5.14
				"read_service_tokens",        // v5-only
				"app_launcher_customization", // v5-only structure
				"policies",                   // In v5, this became a list of objects, not strings
			}
			for _, field := range v5OnlyFields {
				if val, ok := rawJSON[field]; ok && val != nil {
					// For policies, check if it's a list of objects (v5) vs list of strings (v4)
					if field == "policies" {
						if arr, isArray := val.([]interface{}); isArray && len(arr) > 0 {
							// If first element is a map, it's v5 format
							if _, isMap := arr[0].(map[string]interface{}); isMap {
								return false, nil // v5 format - policies is list of objects
							}
							// If first element is a string, it's v4 format
							if _, isString := arr[0].(string); isString {
								return true, nil // v4 format - policies is list of strings
							}
						}
						continue
					}
					return false, nil // v5 format - has v5-only field
				}
			}

			// Check if cors_headers is an array (v4) or object (v5)
			if corsHeaders, ok := rawJSON["cors_headers"]; ok && corsHeaders != nil {
				if _, isArray := corsHeaders.([]interface{}); isArray {
					return true, nil // v4 format - cors_headers is array
				}
				if _, isMap := corsHeaders.(map[string]interface{}); isMap {
					return false, nil // v5 format - cors_headers is object
				}
			}
			// Check saas_app as secondary indicator
			if saasApp, ok := rawJSON["saas_app"]; ok && saasApp != nil {
				if _, isArray := saasApp.([]interface{}); isArray {
					return true, nil // v4 format - saas_app is array
				}
				if _, isMap := saasApp.(map[string]interface{}); isMap {
					return false, nil // v5 format - saas_app is object
				}
			}
			// Check scim_config as another indicator
			if scimConfig, ok := rawJSON["scim_config"]; ok && scimConfig != nil {
				if _, isArray := scimConfig.([]interface{}); isArray {
					return true, nil // v4 format - scim_config is array
				}
				if _, isMap := scimConfig.(map[string]interface{}); isMap {
					return false, nil // v5 format - scim_config is object
				}
			}
			// Check landing_page_design as another indicator
			if landingPageDesign, ok := rawJSON["landing_page_design"]; ok && landingPageDesign != nil {
				if _, isArray := landingPageDesign.([]interface{}); isArray {
					return true, nil // v4 format - landing_page_design is array
				}
				if _, isMap := landingPageDesign.(map[string]interface{}); isMap {
					return false, nil // v5 format - landing_page_design is object
				}
			}
			// Default to v5 if we can't determine - v5.16.0 state with minimal config
			// is more likely than v4 state with minimal config in practice
			return false, nil
		}
	}
	return true, nil // Default to v4 if no raw state (fallback)
}

// upgradeFromV4Internal performs the actual v4 → v5 transformation.
func upgradeFromV4Internal(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 source model
	var v4State SourceAccessApplicationModel
	diags := unmarshalV4State(ctx, req.RawState, &v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, transformDiags := Transform(ctx, v4State)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// unmarshalV4State parses raw state using v4 schema.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceAccessApplicationModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceAccessApplicationSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError("Failed to unmarshal v4 state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err))
		return diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	diags.Append(state.Get(ctx, target)...)
	return diags
}

// unmarshalV5StateToResponse unmarshals v5 raw state and sets it on the response.
func unmarshalV5StateToResponse(ctx context.Context, rawState *tfprotov6.RawState, resp *resource.UpgradeStateResponse) {
	if V5TargetSchema == nil {
		resp.Diagnostics.AddError("Migration configuration error",
			"V5TargetSchema not set. Ensure parent package init() sets v500.V5TargetSchema.")
		return
	}

	targetSchema := V5TargetSchema(ctx)
	targetType := targetSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(targetType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v5 state",
			fmt.Sprintf("Could not parse raw state as v5 format: %s", err))
		return
	}

	resp.State.Raw = rawValue
}

// UpgradeFromV0 is an alias for UpgradeFromVersion0 for backward compatibility.
// Deprecated: Use UpgradeFromVersion0 instead.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	UpgradeFromVersion0(ctx, req, resp)
}

// UpgradeFromV1 handles state upgrades from schema_version=1 to current version.
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from schema_version=1 (no-op)")

	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State upgrade from schema_version=1 completed")
}

// MoveFromAccessApplication handles moves from cloudflare_access_application (v4) to cloudflare_zero_trust_access_application (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_application.example
//	    to   = cloudflare_zero_trust_access_application.example
//	}
func MoveFromAccessApplication(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Verify source is cloudflare_access_application from cloudflare provider
	if req.SourceTypeName != "cloudflare_access_application" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_access_application to cloudflare_zero_trust_access_application",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the v4 state using the v4 schema
	var v4State SourceAccessApplicationModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 state to v5 state
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the v5 state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5State)...)

	tflog.Info(ctx, "State move from cloudflare_access_application to cloudflare_zero_trust_access_application completed successfully")
}

// isCloudflareProvider checks if the provider address is the Cloudflare provider.
func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
