package rulesets

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &RulesetResource{}
var _ resource.ResourceWithImportState = &RulesetResource{}

func NewResource() resource.Resource {
	return &RulesetResource{}
}

type RulesetResource struct {
	client *cloudflare.API
}

type RulesetResourceModel struct {
	AccountID                types.String  `tfsdk:"account_id"`
	Description              types.String  `tfsdk:"description"`
	ID                       types.String  `tfsdk:"id"`
	Kind                     types.String  `tfsdk:"kind"`
	Name                     types.String  `tfsdk:"name"`
	Phase                    types.String  `tfsdk:"phase"`
	Rules                    []*RulesModel `tfsdk:"rules"`
	ShareableEntitlementName types.String  `tfsdk:"shareable_entitlement_name"`
	ZoneID                   types.String  `tfsdk:"zone_id"`
}

type RulesModel struct {
	Action                 types.String                 `tfsdk:"action"`
	ActionParameters       *ActionParametersModel       `tfsdk:"action_parameters"`
	Description            types.String                 `tfsdk:"description"`
	Enabled                types.Bool                   `tfsdk:"enabled"`
	ExposedCredentialCheck *ExposedCredentialCheckModel `tfsdk:"exposed_credential_check"`
	Expression             types.String                 `tfsdk:"expression"`
	ID                     types.String                 `tfsdk:"id"`
	LastUpdated            types.String                 `tfsdk:"last_updated"`
	Logging                *LoggingModel                `tfsdk:"logging"`
	Ratelimit              *RatelimitModel              `tfsdk:"ratelimit"`
	Ref                    types.String                 `tfsdk:"ref"`
	Version                types.String                 `tfsdk:"version"`
}

type ActionParametersModel struct {
	AutomaticHTTPSRewrites  types.Bool   `tfsdk:"automatic_https_rewrites"`
	BIC                     types.Bool   `tfsdk:"bic"`
	Cache                   types.Bool   `tfsdk:"cache"`
	Content                 types.String `tfsdk:"content"`
	ContentType             types.String `tfsdk:"content_type"`
	CookieFields            types.Set    `tfsdk:"cookie_fields"`
	DisableApps             types.Bool   `tfsdk:"disable_apps"`
	DisableRailgun          types.Bool   `tfsdk:"disable_railgun"`
	DisableZaraz            types.Bool   `tfsdk:"disable_zaraz"`
	EmailObfuscation        types.Bool   `tfsdk:"email_obfuscation"`
	HostHeader              types.String `tfsdk:"host_header"`
	HotlinkProtection       types.Bool   `tfsdk:"hotlink_protection"`
	ID                      types.String `tfsdk:"id"`
	Increment               types.Int64  `tfsdk:"increment"`
	Mirage                  types.Bool   `tfsdk:"mirage"`
	OpportunisticEncryption types.Bool   `tfsdk:"opportunistic_encryption"`
	Phases                  types.Set    `tfsdk:"phases"`
	Polish                  types.String `tfsdk:"polish"`
	Products                types.Set    `tfsdk:"products"`
	RequestFields           types.Set    `tfsdk:"request_fields"`
	RespectStrongEtags      types.Bool   `tfsdk:"respect_strong_etags"`
	ResponseFields          types.Set    `tfsdk:"response_fields"`
	RocketLoader            types.String `tfsdk:"rocket_loader"`
	Rules                   types.Map    `tfsdk:"rules"`
	Ruleset                 types.String `tfsdk:"ruleset"`
	Rulesets                types.Set    `tfsdk:"rulesets"`
	SecurityLevel           types.String `tfsdk:"security_level"`
	ServerSideExcludes      types.Bool   `tfsdk:"server_side_excludes"`
	SSL                     types.String `tfsdk:"ssl"`
	StatusCode              types.Int64  `tfsdk:"status_code"`
	SXG                     types.Bool   `tfsdk:"sxg"`
	Version                 types.String `tfsdk:"version"`
}

type RatelimitModel struct {
	Characteristics         types.Set    `tfsdk:"characteristics"`
	CountingExpression      types.String `tfsdk:"counting_expression"`
	MitigationTimeout       types.Int64  `tfsdk:"mitigation_timeout"`
	Period                  types.Int64  `tfsdk:"period"`
	RequestsPerPeriod       types.Int64  `tfsdk:"requests_per_period"`
	RequestsToOrigin        types.Bool   `tfsdk:"requests_to_origin"`
	ScorePerPeriod          types.Int64  `tfsdk:"score_per_period"`
	ScoreResponseHeaderName types.String `tfsdk:"score_response_header_name"`
}

type ExposedCredentialCheckModel struct {
	PasswordExpression types.String `tfsdk:"password_expression"`
	UsernameExpression types.String `tfsdk:"username_expression"`
}

type LoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

func (r *RulesetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ruleset"
}

func (r *RulesetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Cloudflare Ruleset Engine](https://developers.cloudflare.com/firewall/cf-rulesets)
			allows you to create and deploy rules and rulesets.

			The engine syntax, inspired by the Wireshark Display Filter language, is the
			same syntax used in custom Firewall Rules. Cloudflare uses the Ruleset Engine
			in different products, allowing you to configure several products using the same
			basic syntax.
		`),

		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: consts.IDSchemaDescription,
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: "The account identifier to target for the resource.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.ZoneIDSchemaKey))),
					),
				},
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: "The zone identifier to target for the resource.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.AccountIDSchemaKey))),
					),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the ruleset.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Brief summary of the ruleset and its intended use.",
			},
			"shareable_entitlement_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of entitlement that is shareable between entities.",
			},
			"kind": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(cloudflare.RulesetKindValues()...),
				},
				MarkdownDescription: fmt.Sprintf("Type of Ruleset to create. %s", utils.RenderAvailableDocumentationValuesStringSlice(cloudflare.RulesetKindValues())),
			},
			"phase": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(cloudflare.RulesetPhaseValues()...),
				},
				MarkdownDescription: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s", utils.RenderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
			},
		},

		Blocks: map[string]schema.Block{
			"rules": schema.ListNestedBlock{
				MarkdownDescription: "List of rules to apply to the ruleset.",
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"action_parameters": schema.ListNestedBlock{
							MarkdownDescription: "List of parameters that configure the behavior of the ruleset rule action.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"automatic_https_rewrites": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off Cloudflare Automatic HTTPS rewrites.",
									},
									"bic": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Inspect the visitor's browser for headers commonly associated with spammers and certain bots.",
									},
									"cache": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Whether to cache if expression matches.",
									},
									"content": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Content of the custom error response.",
									},
									"content_type": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Content-Type of the custom error response.",
									},
									"cookie_fields": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of cookie values to include as part of custom fields logging.",
									},
									"disable_apps": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn off all active Cloudflare Apps.",
									},
									"disable_railgun": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn off railgun feature of the Cloudflare Speed app.",
									},
									"disable_zaraz": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn off zaraz feature.",
									},
									"email_obfuscation": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the Cloudflare Email Obfuscation feature of the Cloudflare Scrape Shield app.",
									},
									"host_header": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Host Header that request origin receives.",
									},
									"hotlink_protection": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the hotlink protection feature.",
									},
									"id": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Identifier of the action parameter to modify.",
									},
									"increment": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Identifier of the action parameter to modify.",
									},
									"mirage": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off Cloudflare Mirage of the Cloudflare Speed app.",
									},
									"opportunistic_encryption": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the Cloudflare Opportunistic Encryption feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.",
									},
									"phases": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s", utils.RenderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
									},
									"polish": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Apply options from the Polish feature of the Cloudflare Speed app.",
									},
									"products": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: fmt.Sprintf("Products to target with the actions. %s", utils.RenderAvailableDocumentationValuesStringSlice(cloudflare.RulesetActionParameterProductValues())),
									},
									"request_fields": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of request headers to include as part of custom fields logging, in lowercase.",
									},
									"respect_strong_etags": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Respect strong ETags.",
									},
									"response_fields": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of response headers to include as part of custom fields logging, in lowercase.",
									},
									"rocket_loader": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off Cloudflare Rocket Loader in the Cloudflare Speed app.",
									},
									"rules": schema.MapAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "Map of managed WAF rule ID to comma-delimited string of ruleset rule IDs. Example: `rules = { \"efb7b8c949ac4650a09736fc376e9aee\" = \"5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760\" }`",
									},
									"ruleset": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Which ruleset ID to target.",
									},
									"rulesets": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of managed WAF rule IDs to target. Only valid when the `\"action\"` is set to skip",
									},
									"security_level": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Control options for the Security Level feature from the Security app.",
									},
									"server_side_excludes": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the Server Side Excludes feature of the Cloudflare Scrape Shield app.",
									},
									"ssl": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Control options for the SSL feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.",
									},
									"status_code": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "HTTP status code of the custom error response",
									},
									"sxg": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the SXG feature.",
									},
									"version": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Version of the ruleset to deploy.",
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
						"ratelimit": schema.ListNestedBlock{
							MarkdownDescription: "List of parameters that configure HTTP rate limiting behaviour.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"characteristics": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of parameters that define how Cloudflare tracks the request rate for this rule.",
									},
									"period": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "The period of time to consider (in seconds) when evaluating the request rate.",
									},
									"requests_per_period": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "The number of requests over the period of time that will trigger the Rate Limiting rule.",
									},
									"score_per_period": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "The maximum aggregate score over the period of time that will trigger Rate Limiting rule.",
									},
									"score_response_header_name": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Name of HTTP header in the response, set by the origin server, with the score for the current request.",
									},
									"mitigation_timeout": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Once the request rate is reached, the Rate Limiting rule blocks further requests for the period of time defined in this field.",
									},
									"counting_expression": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Criteria for counting HTTP requests to trigger the Rate Limiting action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
									},
									"requests_to_origin": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Whether to include requests to origin within the Rate Limiting count.",
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
						"exposed_credential_check": schema.ListNestedBlock{
							MarkdownDescription: "List of parameters that configure exposed credential checks.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"username_expression": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: `Firewall Rules expression language based on Wireshark display filters for where to check for the "username" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).`,
									},
									"password_expression": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: `Firewall Rules expression language based on Wireshark display filters for where to check for the "password" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).`,
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
						"logging": schema.ListNestedBlock{
							MarkdownDescription: "Override the default logging behavior when a rule is matched.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Whether the rule is active.",
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unique rule identifier.",
						},
						"version": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Version of the ruleset to deploy.",
						},
						"ref": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Rule reference.",
						},
						"enabled": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Whether the rule is active.",
						},
						"description": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Brief summary of the ruleset rule and its intended use.",
						},
						"expression": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Criteria for an HTTP request to trigger the ruleset rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
						},
						"action": schema.StringAttribute{
							MarkdownDescription: fmt.Sprintf("Action to perform in the ruleset rule. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionValues())),
							Validators: []validator.String{
								stringvalidator.OneOf(cloudflare.RulesetRuleActionValues()...),
							},
							Optional: true,
						},
						"last_updated": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "The most recent update to this rule.",
						},
					},
				},
			},
		},
	}
}

func (r *RulesetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *RulesetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RulesetResourceModel

	// Read Terraform plan data into the model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RulesetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RulesetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RulesetResourceModel

	// Read Terraform prior state data into the model.
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *RulesetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
