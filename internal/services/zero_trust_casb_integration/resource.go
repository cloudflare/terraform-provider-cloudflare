package zero_trust_casb_integration

import (
	"context"
	"crypto/rand"
	"crypto/sha3"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"strings"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Auth method slugs, matching the setup flow IDs the API accepts.
const (
	secretsDigestVersion  = "v1"
	secretsDigestSaltSize = 32
	secretsDigestHashSize = 32

	authMethodAnthropicAdminAPIKey                = "anthropic_admin_api_key"
	authMethodAnthropicWorkspaceAPIKey            = "anthropic_workspace_api_key"
	authMethodAnthropicComplianceAPIKey           = "anthropic_compliance_api_key"
	authMethodAWSIAMRole                          = "aws_iam_role"
	authMethodBoxServerAuthentication             = "box_server_authentication"
	authMethodGoogleCloudPlatformServiceAccount   = "google_cloud_platform_service_account"
	authMethodGoogleWorkspaceDomainWideDelegation = "google_domain_wide_delegation_service_account"
	authMethodOpenAIStandardAPIKey                = "chatgpt_standard_api_key"
	authMethodOpenAIComplianceAPIKey              = "chatgpt_compliance_api_key"
)

// supportedAuthMethods maps each supported application to its auth methods.
//
// Deliberately excluded:
//   - OAuth based auth methods (oauth2_standard, oauth2_tenant_id,
//     oauth2_domain, servicenow_oauth, salesforce_fedramp) and therefore the
//     vendors that offer nothing else: Bitbucket, Confluence, Dropbox, GitLab,
//     Jira, Microsoft, Salesforce, ServiceNow, Slack and Zoom. They need an
//     interactive browser redirect that Terraform cannot perform.
//   - github_app_installation: no API credential payload exists, since the
//     credential is an access token obtained through app installation.
//   - google_workspace_workload_identity_federation: alpha-gated.
//   - Google Chat, Okta and 1Password: not live, and absent from the public
//     application enum.
var supportedAuthMethods = map[string]map[string]struct{}{
	string(zero_trust.CasbIntegrationNewParamsApplicationAnthropic): {
		authMethodAnthropicAdminAPIKey:      {},
		authMethodAnthropicWorkspaceAPIKey:  {},
		authMethodAnthropicComplianceAPIKey: {},
	},
	string(zero_trust.CasbIntegrationNewParamsApplicationAws): {
		authMethodAWSIAMRole: {},
	},
	string(zero_trust.CasbIntegrationNewParamsApplicationBox): {
		authMethodBoxServerAuthentication: {},
	},
	string(zero_trust.CasbIntegrationNewParamsApplicationGoogleCloudPlatform): {
		authMethodGoogleCloudPlatformServiceAccount: {},
	},
	string(zero_trust.CasbIntegrationNewParamsApplicationGoogleWorkspace): {
		authMethodGoogleWorkspaceDomainWideDelegation: {},
	},
	string(zero_trust.CasbIntegrationNewParamsApplicationOpenAI): {
		authMethodOpenAIStandardAPIKey:   {},
		authMethodOpenAIComplianceAPIKey: {},
	},
}

var _ resource.ResourceWithConfigure = (*ZeroTrustCasbIntegrationResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustCasbIntegrationResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustCasbIntegrationResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustCasbIntegrationResource{}
}

// ZeroTrustCasbIntegrationResource is manually maintained because the Terraform
// resource exposes only the vendors and auth methods that can be created
// non-interactively through the public API.
type ZeroTrustCasbIntegrationResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustCasbIntegrationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_casb_integration"
}

func (r *ZeroTrustCasbIntegrationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

// ModifyPlan recomputes the secrets digest from configuration. Confidential
// arguments are write-only, so they are null in both plan and state; without
// recomputing the digest here a secret rotation would produce no diff and
// Terraform would never call Update.
func (r *ZeroTrustCasbIntegrationResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return // Destroy.
	}

	var config ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resolved, err := config.resolveAuth()
	if err != nil {
		// The config validators report this with better context.
		return
	}
	if resolved.builder.hasUnknown() {
		// Secrets derive from values not known until apply.
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("secrets_digest"), types.StringUnknown())...)
		return
	}

	var state ZeroTrustCasbIntegrationModel
	if !req.State.Raw.IsNull() {
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	salt, err := secretsDigestSaltOrNew(state.SecretsDigest.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to generate secrets digest salt", err.Error())
		return
	}

	digest, err := config.secretsDigest(salt)
	if err != nil {
		// Malformed secrets are reported during apply with full context.
		return
	}
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("secrets_digest"), types.StringValue(digest))...)
}

func (r *ZeroTrustCasbIntegrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Confidential arguments are write-only, so they are null in the plan and
	// must be read from configuration.
	var config ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.newParams(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	salt, err := secretsDigestSaltOrNew(data.SecretsDigest.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to generate secrets digest salt", err.Error())
		return
	}
	digest, err := config.secretsDigest(salt)
	if err != nil {
		resp.Diagnostics.AddError("invalid integration credentials", err.Error())
		return
	}
	data.SecretsDigest = types.StringValue(digest)

	requestedPaused := data.Paused.ValueBool()
	// Credential-bearing requests intentionally omit request logging.
	result, err := r.client.ZeroTrust.Casb.Integrations.New(
		ctx,
		params,
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	setIntegrationState(&data, result.ID, result.Name, result.IsPaused)
	resolveComputedDLPProfiles(ctx, &data, result.DLPProfiles)
	if result.IsPaused != requestedPaused {
		if err := r.setPaused(ctx, &data, requestedPaused); err != nil {
			resp.Diagnostics.AddError("failed to set integration pause state", err.Error())
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustCasbIntegrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Confidential arguments are write-only and absent from plan and state.
	var config ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = state.ID
	salt, ok := secretsDigestSalt(data.SecretsDigest.ValueString())
	if !ok {
		var saltErr error
		salt, saltErr = secretsDigestSaltOrNew(state.SecretsDigest.ValueString())
		if saltErr != nil {
			resp.Diagnostics.AddError("failed to generate secrets digest salt", saltErr.Error())
			return
		}
	}
	digest, err := config.secretsDigest(salt)
	if err != nil {
		resp.Diagnostics.AddError("invalid integration credentials", err.Error())
		return
	}

	desiredPaused := data.Paused.ValueBool()
	currentPaused := state.Paused.ValueBool()
	data.SecretsDigest = types.StringValue(digest)

	nameChanged := data.Name.ValueString() != state.Name.ValueString()
	credentialsChanged := credentialsChanged(&config, &state, digest)
	collectionsChanged := !data.DLPProfiles.Equal(state.DLPProfiles) ||
		!data.Permissions.Equal(state.Permissions) ||
		!data.UseCases.Equal(state.UseCases)

	if nameChanged || credentialsChanged || collectionsChanged {
		params := zero_trust.CasbIntegrationUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		}
		if nameChanged {
			params.Name = cloudflare.F(data.Name.ValueString())
		}
		if credentialsChanged {
			credentials, err := config.credentials()
			if err != nil {
				resp.Diagnostics.AddError("invalid integration credentials", err.Error())
				return
			}
			params.Credentials = cloudflare.F(credentials)
		}
		if collectionsChanged {
			resp.Diagnostics.Append(data.applyUpdateCollections(ctx, &params)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}

		// The update may contain replacement secrets.
		result, err := r.client.ZeroTrust.Casb.Integrations.Update(
			ctx,
			data.ID.ValueString(),
			params,
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		setIntegrationState(&data, result.ID, result.Name, result.IsPaused)
		resolveComputedDLPProfiles(ctx, &data, result.DLPProfiles)
		currentPaused = result.IsPaused
	} else {
		data.Paused = types.BoolValue(currentPaused)
	}

	if currentPaused != desiredPaused {
		if err := r.setPaused(ctx, &data, desiredPaused); err != nil {
			resp.Diagnostics.AddError("failed to set integration pause state", err.Error())
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustCasbIntegrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var httpResponse *http.Response
	result, err := r.client.ZeroTrust.Casb.Integrations.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.CasbIntegrationGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseInto(&httpResponse),
	)
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	setIntegrationState(&data, result.ID, result.Name, result.IsPaused)

	// Refresh detects drift for the collections the API reports back.
	// Permissions are not returned, so drift in them cannot be detected.
	data.DLPProfiles = stringsToSet(ctx, result.DLPProfiles)
	data.UseCases = stringsToSet(ctx, enabledUseCases(result.UseCases))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustCasbIntegrationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ZeroTrustCasbIntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var httpResponse *http.Response
	err := r.client.ZeroTrust.Casb.Integrations.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.CasbIntegrationDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseInto(&httpResponse),
	)
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
	}
}

func (r *ZeroTrustCasbIntegrationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var accountID string
	var integrationID string
	resp.Diagnostics.Append(importpath.ParseImportID(
		req.ID,
		"<account_id>/<id>",
		&accountID,
		&integrationID,
	)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.ZeroTrust.Casb.Integrations.Get(
		ctx,
		integrationID,
		zero_trust.CasbIntegrationGetParams{
			AccountID: cloudflare.F(accountID),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if err := validateImportableIntegration(result); err != nil {
		resp.Diagnostics.AddError("unsupported CASB integration", err.Error())
		return
	}

	// Credentials are never returned by the API, so the vendor block cannot be
	// populated here; it must be supplied in configuration after import. The
	// null secrets digest makes the first apply push the configured secrets.
	data := ZeroTrustCasbIntegrationModel{
		AccountID:   types.StringValue(accountID),
		Permissions: customfield.NullSet[types.String](ctx),
	}
	setIntegrationState(&data, result.ID, result.Name, result.IsPaused)
	data.DLPProfiles = stringsToSet(ctx, result.DLPProfiles)
	data.UseCases = stringsToSet(ctx, enabledUseCases(result.UseCases))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustCasbIntegrationResource) setPaused(ctx context.Context, data *ZeroTrustCasbIntegrationModel, paused bool) error {
	if paused {
		result, err := r.client.ZeroTrust.Casb.Integrations.Pause(
			ctx,
			data.ID.ValueString(),
			zero_trust.CasbIntegrationPauseParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
			},
		)
		if err != nil {
			return err
		}
		setIntegrationState(data, result.ID, result.Name, result.IsPaused)
		resolveComputedDLPProfiles(ctx, data, result.DLPProfiles)
		return nil
	}

	result, err := r.client.ZeroTrust.Casb.Integrations.Resume(
		ctx,
		data.ID.ValueString(),
		zero_trust.CasbIntegrationResumeParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		return err
	}
	setIntegrationState(data, result.ID, result.Name, result.IsPaused)
	resolveComputedDLPProfiles(ctx, data, result.DLPProfiles)
	return nil
}

// ── Auth method resolution ──────────────────────────────────────────────────

// credentialBuilder is implemented by every auth method model. Arguments are
// split by confidentiality: confidentialArgs never reach state and are covered
// by the digest, publicArgs are stored in state and diffed normally.
type credentialBuilder interface {
	confidentialArgs() (map[string]any, error)
	publicArgs() map[string]any
	hasUnknown() bool
}

// buildCredentials assembles the full payload the API expects.
func buildCredentials(builder credentialBuilder) (map[string]any, error) {
	confidential, err := builder.confidentialArgs()
	if err != nil {
		return nil, err
	}
	public := builder.publicArgs()

	credentials := make(map[string]any, len(confidential)+len(public))
	for key, value := range public {
		credentials[key] = value
	}
	for key, value := range confidential {
		credentials[key] = value
	}
	return credentials, nil
}

func anyUnknown(values ...attr.Value) bool {
	for _, value := range values {
		if value.IsUnknown() {
			return true
		}
	}
	return false
}

// setOptional adds a value only when the practitioner supplied one, so omitted
// optional arguments are absent from the payload rather than sent as "".
func setOptional(args map[string]any, key string, value types.String) {
	if !value.IsNull() && !value.IsUnknown() {
		args[key] = value.ValueString()
	}
}

type resolvedAuth struct {
	application zero_trust.CasbIntegrationNewParamsApplication
	authMethod  string
	builder     credentialBuilder
}

// resolveAuth finds the single configured vendor and auth method. The schema
// validators already enforce exactly one of each, so the error paths here are
// defensive and also cover models built directly in tests.
func (m ZeroTrustCasbIntegrationModel) resolveAuth() (resolvedAuth, error) {
	var matches []resolvedAuth

	if v := m.Anthropic; v != nil {
		app := zero_trust.CasbIntegrationNewParamsApplicationAnthropic
		if a := v.AdminAPIKey; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodAnthropicAdminAPIKey, a})
		}
		if a := v.WorkspaceAPIKey; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodAnthropicWorkspaceAPIKey, a})
		}
		if a := v.ComplianceAPIKey; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodAnthropicComplianceAPIKey, a})
		}
	}
	if v := m.AWS; v != nil {
		app := zero_trust.CasbIntegrationNewParamsApplicationAws
		if a := v.IAMRole; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodAWSIAMRole, a})
		}
	}
	if v := m.Box; v != nil {
		app := zero_trust.CasbIntegrationNewParamsApplicationBox
		if a := v.ServerAuthentication; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodBoxServerAuthentication, a})
		}
	}
	if v := m.GoogleCloudPlatform; v != nil {
		app := zero_trust.CasbIntegrationNewParamsApplicationGoogleCloudPlatform
		if a := v.ServiceAccount; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodGoogleCloudPlatformServiceAccount, a})
		}
	}
	if v := m.GoogleWorkspace; v != nil {
		app := zero_trust.CasbIntegrationNewParamsApplicationGoogleWorkspace
		if a := v.DomainWideDelegationServiceAccount; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodGoogleWorkspaceDomainWideDelegation, a})
		}
	}
	if v := m.OpenAI; v != nil {
		app := zero_trust.CasbIntegrationNewParamsApplicationOpenAI
		if a := v.StandardAPIKey; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodOpenAIStandardAPIKey, a})
		}
		if a := v.ComplianceAPIKey; a != nil {
			matches = append(matches, resolvedAuth{app, authMethodOpenAIComplianceAPIKey, a})
		}
	}

	switch len(matches) {
	case 1:
		return matches[0], nil
	case 0:
		return resolvedAuth{}, fmt.Errorf("exactly one vendor and one authentication method must be configured, got none")
	default:
		configured := make([]string, 0, len(matches))
		for _, match := range matches {
			configured = append(configured, match.authMethod)
		}
		return resolvedAuth{}, fmt.Errorf(
			"exactly one vendor and one authentication method must be configured, got %d: %s",
			len(matches), strings.Join(configured, ", "),
		)
	}
}

func (m ZeroTrustCasbIntegrationModel) credentials() (map[string]any, error) {
	resolved, err := m.resolveAuth()
	if err != nil {
		return nil, err
	}
	return buildCredentials(resolved.builder)
}

// secretsDigest returns a randomly salted SHA3-256 digest over the
// confidential arguments only. The payload is hashed after decoding, and
// encoding/json sorts object keys, so cosmetic JSON changes are not drift.
func (m ZeroTrustCasbIntegrationModel) secretsDigest(salt []byte) (string, error) {
	if len(salt) != secretsDigestSaltSize {
		return "", fmt.Errorf("secrets digest salt must be %d bytes", secretsDigestSaltSize)
	}
	resolved, err := m.resolveAuth()
	if err != nil {
		return "", err
	}
	confidential, err := resolved.builder.confidentialArgs()
	if err != nil {
		return "", err
	}
	canonical, err := json.Marshal(confidential)
	if err != nil {
		return "", fmt.Errorf("failed to encode secrets for hashing: %w", err)
	}

	digest := sha3.New256()
	digest.Write(salt)
	digest.Write(canonical)
	return fmt.Sprintf("%s:%s:%s", secretsDigestVersion, hex.EncodeToString(salt), hex.EncodeToString(digest.Sum(nil))), nil
}

func newSecretsDigestSalt() ([]byte, error) {
	salt := make([]byte, secretsDigestSaltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func secretsDigestSaltOrNew(value string) ([]byte, error) {
	if salt, ok := secretsDigestSalt(value); ok {
		return salt, nil
	}
	return newSecretsDigestSalt()
}

func secretsDigestSalt(value string) ([]byte, bool) {
	parts := strings.Split(value, ":")
	if len(parts) != 3 || parts[0] != secretsDigestVersion {
		return nil, false
	}
	salt, err := hex.DecodeString(parts[1])
	if err != nil || len(salt) != secretsDigestSaltSize {
		return nil, false
	}
	digest, err := hex.DecodeString(parts[2])
	if err != nil || len(digest) != secretsDigestHashSize {
		return nil, false
	}
	return salt, true
}

// credentialsChanged reports whether the credential payload needs to be resent.
// Confidential arguments are compared through the digest, since they are absent
// from state; non-confidential arguments are compared directly.
func credentialsChanged(config, state *ZeroTrustCasbIntegrationModel, digest string) bool {
	if state.SecretsDigest.ValueString() != digest {
		return true
	}

	configAuth, configErr := config.resolveAuth()
	stateAuth, stateErr := state.resolveAuth()
	if configErr != nil || stateErr != nil {
		return (configErr == nil) != (stateErr == nil)
	}
	if configAuth.authMethod != stateAuth.authMethod {
		return true
	}
	return !reflect.DeepEqual(configAuth.builder.publicArgs(), stateAuth.builder.publicArgs())
}

// ── Credential payloads ─────────────────────────────────────────────────────

func serviceAccountCredentials(key jsontypes.Normalized) (map[string]any, error) {
	var parsed map[string]any
	if err := json.Unmarshal([]byte(key.ValueString()), &parsed); err != nil {
		return nil, fmt.Errorf("service_account_key_json must contain a valid JSON object: %w", err)
	}
	if parsed == nil {
		return nil, fmt.Errorf("service_account_key_json must contain a JSON object")
	}
	return parsed, nil
}

// Anthropic admin API key.

func (m ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel) confidentialArgs() (map[string]any, error) {
	return map[string]any{"api_key": m.APIKey.ValueString()}, nil
}

func (m ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel) publicArgs() map[string]any {
	args := map[string]any{}
	setOptional(args, "tenant_id", m.TenantID)
	return args
}

func (m ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel) hasUnknown() bool {
	return anyUnknown(m.APIKey, m.TenantID)
}

// Anthropic workspace API key.

func (m ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel) confidentialArgs() (map[string]any, error) {
	return map[string]any{"api_key": m.APIKey.ValueString()}, nil
}

func (m ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel) publicArgs() map[string]any {
	return map[string]any{"tenant_id": m.TenantID.ValueString()}
}

func (m ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel) hasUnknown() bool {
	return anyUnknown(m.APIKey, m.TenantID)
}

// Anthropic compliance API key.

func (m ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel) confidentialArgs() (map[string]any, error) {
	return map[string]any{"compliance_api_key": m.ComplianceAPIKey.ValueString()}, nil
}

func (m ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel) publicArgs() map[string]any {
	args := map[string]any{}
	setOptional(args, "tenant_id", m.TenantID)
	return args
}

func (m ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel) hasUnknown() bool {
	return anyUnknown(m.ComplianceAPIKey, m.TenantID)
}

// AWS IAM role. No confidential arguments.

func (m ZeroTrustCasbIntegrationAWSIAMRoleModel) confidentialArgs() (map[string]any, error) {
	return map[string]any{}, nil
}

func (m ZeroTrustCasbIntegrationAWSIAMRoleModel) publicArgs() map[string]any {
	return map[string]any{
		"role_arn":    m.RoleARN.ValueString(),
		"external_id": m.ExternalID.ValueString(),
	}
}

func (m ZeroTrustCasbIntegrationAWSIAMRoleModel) hasUnknown() bool {
	return anyUnknown(m.RoleARN, m.ExternalID)
}

// Box server authentication. No confidential arguments.

func (m ZeroTrustCasbIntegrationBoxServerAuthenticationModel) confidentialArgs() (map[string]any, error) {
	return map[string]any{}, nil
}

func (m ZeroTrustCasbIntegrationBoxServerAuthenticationModel) publicArgs() map[string]any {
	return map[string]any{"enterprise_id": m.EnterpriseID.ValueString()}
}

func (m ZeroTrustCasbIntegrationBoxServerAuthenticationModel) hasUnknown() bool {
	return anyUnknown(m.EnterpriseID)
}

// Google Cloud Platform service account.

func (m ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel) confidentialArgs() (map[string]any, error) {
	parsed, err := serviceAccountCredentials(m.ServiceAccountKeyJSON)
	if err != nil {
		return nil, err
	}
	return map[string]any{"service_account_credentials": parsed}, nil
}

func (m ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel) publicArgs() map[string]any {
	return map[string]any{}
}

func (m ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel) hasUnknown() bool {
	return anyUnknown(m.ServiceAccountKeyJSON)
}

// Google Workspace domain-wide delegation.

func (m ZeroTrustCasbIntegrationGoogleWorkspaceDomainWideDelegationServiceAccountModel) confidentialArgs() (map[string]any, error) {
	parsed, err := serviceAccountCredentials(m.ServiceAccountKeyJSON)
	if err != nil {
		return nil, err
	}
	return map[string]any{"service_account_credentials": parsed}, nil
}

func (m ZeroTrustCasbIntegrationGoogleWorkspaceDomainWideDelegationServiceAccountModel) publicArgs() map[string]any {
	return map[string]any{"administrator_email": m.AdministratorEmail.ValueString()}
}

func (m ZeroTrustCasbIntegrationGoogleWorkspaceDomainWideDelegationServiceAccountModel) hasUnknown() bool {
	return anyUnknown(m.ServiceAccountKeyJSON, m.AdministratorEmail)
}

// OpenAI standard API key.

func (m ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel) confidentialArgs() (map[string]any, error) {
	args := map[string]any{"admin_api_key": m.AdminAPIKey.ValueString()}
	setOptional(args, "project_api_key", m.ProjectAPIKey)
	return args, nil
}

func (m ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel) publicArgs() map[string]any {
	args := map[string]any{"organization_id": m.OrganizationID.ValueString()}
	setOptional(args, "project_id", m.ProjectID)
	return args
}

func (m ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel) hasUnknown() bool {
	return anyUnknown(m.AdminAPIKey, m.OrganizationID, m.ProjectAPIKey, m.ProjectID)
}

// OpenAI compliance API key.

func (m ZeroTrustCasbIntegrationOpenAIComplianceAPIKeyModel) confidentialArgs() (map[string]any, error) {
	args := map[string]any{
		"admin_api_key":      m.AdminAPIKey.ValueString(),
		"compliance_api_key": m.ComplianceAPIKey.ValueString(),
	}
	setOptional(args, "project_api_key", m.ProjectAPIKey)
	return args, nil
}

func (m ZeroTrustCasbIntegrationOpenAIComplianceAPIKeyModel) publicArgs() map[string]any {
	args := map[string]any{
		"organization_id": m.OrganizationID.ValueString(),
		"workspace_id":    m.WorkspaceID.ValueString(),
	}
	setOptional(args, "project_id", m.ProjectID)
	return args
}

func (m ZeroTrustCasbIntegrationOpenAIComplianceAPIKeyModel) hasUnknown() bool {
	return anyUnknown(m.AdminAPIKey, m.OrganizationID, m.ComplianceAPIKey, m.WorkspaceID, m.ProjectAPIKey, m.ProjectID)
}

// ── Request parameters ──────────────────────────────────────────────────────

// newParams builds the create request. Non-credential fields come from the
// plan; credentials come from configuration because they are write-only.
func (m ZeroTrustCasbIntegrationModel) newParams(ctx context.Context, config *ZeroTrustCasbIntegrationModel) (zero_trust.CasbIntegrationNewParams, diag.Diagnostics) {
	var diags diag.Diagnostics

	resolved, err := config.resolveAuth()
	if err != nil {
		diags.AddError("invalid integration credentials", err.Error())
		return zero_trust.CasbIntegrationNewParams{}, diags
	}
	credentials, err := buildCredentials(resolved.builder)
	if err != nil {
		diags.AddError("invalid integration credentials", err.Error())
		return zero_trust.CasbIntegrationNewParams{}, diags
	}

	params := zero_trust.CasbIntegrationNewParams{
		AccountID:   cloudflare.F(m.AccountID.ValueString()),
		Application: cloudflare.F(resolved.application),
		Credentials: cloudflare.F(credentials),
		Name:        cloudflare.F(m.Name.ValueString()),
		AuthMethod:  cloudflare.F(resolved.authMethod),
	}

	dlpProfiles, d := setToStrings(ctx, m.DLPProfiles)
	diags.Append(d...)
	if dlpProfiles != nil {
		params.DLPProfiles = cloudflare.F(dlpProfiles)
	}

	permissions, d := setToStrings(ctx, m.Permissions)
	diags.Append(d...)
	if permissions != nil {
		params.Permissions = cloudflare.F(permissions)
	}

	useCases, d := setToStrings(ctx, m.UseCases)
	diags.Append(d...)
	if useCases != nil {
		converted := make([]zero_trust.CasbIntegrationNewParamsUseCase, 0, len(useCases))
		for _, useCase := range useCases {
			converted = append(converted, zero_trust.CasbIntegrationNewParamsUseCase(useCase))
		}
		params.UseCases = cloudflare.F(converted)
	}

	return params, diags
}

func (m ZeroTrustCasbIntegrationModel) applyUpdateCollections(ctx context.Context, params *zero_trust.CasbIntegrationUpdateParams) diag.Diagnostics {
	var diags diag.Diagnostics

	dlpProfiles, d := setToStrings(ctx, m.DLPProfiles)
	diags.Append(d...)
	if dlpProfiles != nil {
		params.DLPProfiles = cloudflare.F(dlpProfiles)
	}

	permissions, d := setToStrings(ctx, m.Permissions)
	diags.Append(d...)
	if permissions != nil {
		params.Permissions = cloudflare.F(permissions)
	}

	useCases, d := setToStrings(ctx, m.UseCases)
	diags.Append(d...)
	if useCases != nil {
		converted := make([]zero_trust.CasbIntegrationUpdateParamsUseCase, 0, len(useCases))
		for _, useCase := range useCases {
			converted = append(converted, zero_trust.CasbIntegrationUpdateParamsUseCase(useCase))
		}
		params.UseCases = cloudflare.F(converted)
	}

	return diags
}

// ── Collection helpers ──────────────────────────────────────────────────────

// setToStrings returns nil for a null or unknown set so that the caller omits
// the field from the request rather than sending an empty list, which the API
// would read as "remove everything".
func setToStrings(ctx context.Context, set customfield.Set[types.String]) ([]string, diag.Diagnostics) {
	if set.IsNullOrUnknown() {
		return nil, nil
	}
	values := []string{}
	diags := set.ElementsAs(ctx, &values, false)
	return values, diags
}

func stringsToSet(ctx context.Context, values []string) customfield.Set[types.String] {
	elements := make([]attr.Value, 0, len(values))
	for _, value := range values {
		elements = append(elements, types.StringValue(value))
	}
	return customfield.NewSetMust[types.String](ctx, elements)
}

// enabledUseCases extracts the enabled use case slugs from an API response.
func enabledUseCases(useCases []map[string]any) []string {
	enabled := []string{}
	for _, useCase := range useCases {
		id, idOK := useCase["id"].(string)
		isEnabled, enabledOK := useCase["is_enabled"].(bool)
		if idOK && enabledOK && isEnabled {
			enabled = append(enabled, id)
		}
	}
	return enabled
}

// resolveComputedDLPProfiles fills dlp_profiles from the API only when the
// practitioner left it unset, so a configured value is never contradicted.
func resolveComputedDLPProfiles(ctx context.Context, data *ZeroTrustCasbIntegrationModel, dlpProfiles []string) {
	if data.DLPProfiles.IsUnknown() {
		data.DLPProfiles = stringsToSet(ctx, dlpProfiles)
	}
}

// ── Import ──────────────────────────────────────────────────────────────────

func validateImportableIntegration(result *zero_trust.CasbIntegrationGetResponse) error {
	applicationID := result.Application["id"]
	authMethods, ok := supportedAuthMethods[applicationID]
	if !ok {
		application := applicationID
		if application == "" {
			application = result.Application["display_name"]
		}
		return fmt.Errorf("unsupported application %q", application)
	}

	authMethod := result.AuthMethod["id"]
	if _, ok := authMethods[authMethod]; !ok {
		return fmt.Errorf("unsupported auth method %q for application %q", authMethod, applicationID)
	}

	if !slices.Contains(enabledUseCases(result.UseCases), "casb") {
		return fmt.Errorf("expected the casb use case to be enabled")
	}
	return nil
}

func setIntegrationState(data *ZeroTrustCasbIntegrationModel, id, name string, paused bool) {
	data.ID = types.StringValue(id)
	data.Name = types.StringValue(name)
	data.Paused = types.BoolValue(paused)
}
