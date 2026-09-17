package zero_trust_casb_integration

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCasbIntegrationResource)(nil)

// authMethodPathsByVendor is the single source of truth for the exposed
// vendor/auth method matrix. It drives the "exactly one" config validators.
func authMethodPathsByVendor() map[string][]string {
	return map[string][]string{
		"anthropic": {
			"anthropic_admin_api_key",
			"anthropic_workspace_api_key",
			"anthropic_compliance_api_key",
		},
		"aws":                   {"aws_iam_role"},
		"box":                   {"box_server_authentication"},
		"google_cloud_platform": {"google_cloud_platform_service_account"},
		"google_workspace":      {"google_domain_wide_delegation_service_account"},
		"openai": {
			"chatgpt_standard_api_key",
			"chatgpt_compliance_api_key",
		},
	}
}

func vendorPaths() []path.Expression {
	paths := []path.Expression{}
	for vendor := range authMethodPathsByVendor() {
		paths = append(paths, path.MatchRoot(vendor))
	}
	return paths
}

func authMethodPaths() []path.Expression {
	paths := []path.Expression{}
	for vendor, authMethods := range authMethodPathsByVendor() {
		for _, authMethod := range authMethods {
			paths = append(paths, path.MatchRoot(vendor).AtName(authMethod))
		}
	}
	return paths
}

// Every auth method belongs to exactly one vendor, so requiring exactly one
// auth method also rejects two vendors and a vendor left without an auth method.
func (r *ZeroTrustCasbIntegrationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(vendorPaths()...),
		resourcevalidator.ExactlyOneOf(authMethodPaths()...),
	}
}

// requiresReplaceOnAuthMethodChange forces replacement when an auth method
// block is added or removed. Neither the vendor nor the auth method can be
// changed in place: CasbIntegrationUpdateParams carries no application or
// auth_method field. Editing arguments inside a block updates in place.
func requiresReplaceOnAuthMethodChange() planmodifier.Object {
	const description = "Changing the vendor or authentication method requires the integration to be recreated."
	return objectplanmodifier.RequiresReplaceIf(
		func(ctx context.Context, req planmodifier.ObjectRequest, resp *objectplanmodifier.RequiresReplaceIfFuncResponse) {
			if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
				return
			}
			if req.StateValue.IsNull() && !req.PlanValue.IsNull() {
				var state ZeroTrustCasbIntegrationModel
				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
				if resp.Diagnostics.HasError() || state.SecretsDigest.IsNull() {
					return
				}
			}
			resp.RequiresReplace = req.StateValue.IsNull() != req.PlanValue.IsNull()
		},
		description,
		description,
	)
}

func authMethodAttribute(description string, attributes map[string]schema.Attribute) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description:   description,
		Optional:      true,
		Attributes:    attributes,
		PlanModifiers: []planmodifier.Object{requiresReplaceOnAuthMethodChange()},
	}
}

// secretAttribute is a confidential argument: write-only, so it never reaches
// state, and covered by secrets_digest instead.
func secretAttribute(description string, required bool) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description + " This value is write-only and is never persisted to Terraform state.",
		Required:    required,
		Optional:    !required,
		Sensitive:   true,
		WriteOnly:   true,
	}
}

// plainAttribute is a non-confidential argument, stored in state so that
// ordinary Terraform diffing applies.
func plainAttribute(description string, required bool) schema.StringAttribute {
	return schema.StringAttribute{
		Description: description,
		Required:    required,
		Optional:    !required,
	}
}

func serviceAccountKeyJSONAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Description: "Contents of a Google service account JSON key file. This value is write-only and is never persisted to Terraform state.",
		Required:    true,
		Sensitive:   true,
		WriteOnly:   true,
		CustomType:  jsontypes.NormalizedType{},
	}
}

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Creates and manages a CASB integration. Exactly one vendor and exactly one authentication method must be configured.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Integration ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Cloudflare account identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the integration.",
				Required:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "Whether the integration is paused.",
				Required:    true,
			},

			"dlp_profiles": schema.SetAttribute{
				Description:   "DLP profile IDs to associate with the integration.",
				CustomType:    customfield.NewSetType[types.String](ctx),
				ElementType:   types.StringType,
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			// The API does not return permissions, so this is tracked from
			// configuration only and cannot detect out-of-band drift.
			"permissions": schema.SetAttribute{
				Description: "Permission scopes granted to the integration.",
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
				Optional:    true,
			},
			"use_cases": schema.SetAttribute{
				Description: "Use cases to enroll the integration in.",
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Default: setdefault.StaticValue(
					types.SetValueMust(types.StringType, []attr.Value{types.StringValue("casb")}),
				),
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("casb", "ces", "auto_remediation"),
					),
				},
			},

			// Confidential arguments are write-only, so nothing about them
			// reaches state except this digest. ModifyPlan recomputes it from
			// configuration on every plan, which is what makes a secret
			// rotation visible to Terraform at all.
			"secrets_digest": schema.StringAttribute{
				Description: "Secure digest of the confidential arguments. Managed by the provider.",
				Computed:    true,
			},

			"anthropic": schema.SingleNestedAttribute{
				Description: "Anthropic integration configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"anthropic_admin_api_key": authMethodAttribute(
						"Authenticate with an Anthropic Admin API key.",
						map[string]schema.Attribute{
							"api_key":   secretAttribute("Anthropic Admin API key.", true),
							"tenant_id": plainAttribute("Organization ID. Auto-extracted from the key if not provided.", false),
						},
					),
					"anthropic_workspace_api_key": authMethodAttribute(
						"Authenticate with an Anthropic Workspace API key.",
						map[string]schema.Attribute{
							"api_key":   secretAttribute("Anthropic Workspace API key.", true),
							"tenant_id": plainAttribute("Workspace ID, found in the Anthropic Console URL after /workspaces/.", true),
						},
					),
					"anthropic_compliance_api_key": authMethodAttribute(
						"Authenticate with an Anthropic Compliance API key.",
						map[string]schema.Attribute{
							"compliance_api_key": secretAttribute("Anthropic Compliance API key.", true),
							"tenant_id":          plainAttribute("Organization ID. Auto-extracted from the key if not provided.", false),
						},
					),
				},
			},

			"aws": schema.SingleNestedAttribute{
				Description: "AWS integration configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"aws_iam_role": authMethodAttribute(
						"Authenticate by delegating to a cross-account IAM role.",
						map[string]schema.Attribute{
							"role_arn":    plainAttribute("ARN of the cross-account IAM role Cloudflare will assume.", true),
							"external_id": plainAttribute("External ID required when assuming the IAM role.", true),
						},
					),
				},
			},

			"box": schema.SingleNestedAttribute{
				Description: "Box integration configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"box_server_authentication": authMethodAttribute(
						"Authenticate with Box server authentication. Before creating the integration, add the Cloudflare CASB application in Box Admin Console > Integrations > Platform Apps Manager > Server Authentication Apps using client ID `puaghckpy0578r8p6f3g0rf860unup4r`.",
						map[string]schema.Attribute{
							"enterprise_id": plainAttribute("Box Enterprise ID from Admin Console > Accounts & Billing.", true),
						},
					),
				},
			},

			"google_cloud_platform": schema.SingleNestedAttribute{
				Description: "Google Cloud Platform integration configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"google_cloud_platform_service_account": authMethodAttribute(
						"Authenticate with a service account key.",
						map[string]schema.Attribute{
							"service_account_key_json": serviceAccountKeyJSONAttribute(),
						},
					),
				},
			},

			"google_workspace": schema.SingleNestedAttribute{
				Description: "Google Workspace integration configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"google_domain_wide_delegation_service_account": authMethodAttribute(
						"Authenticate with a service account granted domain-wide delegation.",
						map[string]schema.Attribute{
							"service_account_key_json": serviceAccountKeyJSONAttribute(),
							"administrator_email":      plainAttribute("A Google Workspace super administrator email address.", true),
						},
					),
				},
			},

			"openai": schema.SingleNestedAttribute{
				Description: "OpenAI integration configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"chatgpt_standard_api_key": authMethodAttribute(
						"Authenticate with an OpenAI Admin API key.",
						map[string]schema.Attribute{
							"admin_api_key":   secretAttribute("OpenAI Admin API key with api.management.read access.", true),
							"organization_id": plainAttribute("OpenAI Organization ID.", true),
							"project_api_key": secretAttribute("OpenAI Project API key, used for DLP.", false),
							"project_id":      plainAttribute("OpenAI Project ID, used for DLP.", false),
						},
					),
					"chatgpt_compliance_api_key": authMethodAttribute(
						"Authenticate with an OpenAI Compliance API key. Requires an Enterprise plan.",
						map[string]schema.Attribute{
							"admin_api_key":      secretAttribute("OpenAI Admin API key with api.management.read access.", true),
							"organization_id":    plainAttribute("OpenAI Organization ID.", true),
							"compliance_api_key": secretAttribute("OpenAI Compliance API key for audit logs.", true),
							"workspace_id":       plainAttribute("OpenAI Workspace ID for compliance data.", true),
							"project_api_key":    secretAttribute("OpenAI Project API key, used for DLP.", false),
							"project_id":         plainAttribute("OpenAI Project ID, used for DLP.", false),
						},
					),
				},
			},
		},
	}
}

func (r *ZeroTrustCasbIntegrationResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}
