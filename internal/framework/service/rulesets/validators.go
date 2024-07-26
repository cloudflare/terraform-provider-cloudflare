package rulesets

import (
	"context"
	"fmt"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type sbfmDeprecationWarningValidator struct{}

func (v sbfmDeprecationWarningValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Cloudflare is going to change the way Super Bot Fight Mode managed rules are configured through Terraform and our API. No action is required at this time. " +
		" Please follow updates to our documentation regarding this here: https://developers.cloudflare.com/bots/get-started/biz-and-ent/#ruleset-engine")
}

func (v sbfmDeprecationWarningValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Cloudflare is going to change the way Super Bot Fight Mode managed rules are configured through Terraform and our API. **No action is required at this time**. " +
		" Please follow updates to our documentation regarding this [here](https://developers.cloudflare.com/bots/get-started/biz-and-ent/#ruleset-engine)")
}

func (v sbfmDeprecationWarningValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	if req.ConfigValue.ValueString() == string(cfv1.RulesetPhaseHTTPRequestSBFM) {
		resp.Diagnostics.AddAttributeWarning(
			req.Path,
			fmt.Sprintf(`%q phase will soon be deprecated in the "cloudflare_ruleset" resource`, string(cfv1.RulesetPhaseHTTPRequestSBFM)),
			v.Description(ctx),
		)

		return
	}
}

type RulesetActionParameterEdgeTTL struct {
	Mode       basetypes.StringValue `tfsdk:"mode"`
	Default    basetypes.Int64Value  `tfsdk:"default"`
	StatusCode basetypes.ListValue   `tfsdk:"status_code_ttl"`
}

type EdgeTTLValidator struct{}

func (v EdgeTTLValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("ttl values are required when the override_origin mode is set")
}

func (v EdgeTTLValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("ttl values are required when the override_origin mode is set")
}

func (v EdgeTTLValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	var parameter RulesetActionParameterEdgeTTL

	diag := req.ConfigValue.As(ctx, &parameter, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	if parameter.Mode.ValueString() == "override_origin" {
		if parameter.Default.ValueInt64() <= 0 {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				errInvalidConfiguration,
				fmt.Sprintf("using mode '%s' requires setting a default for ttl", parameter.Mode.ValueString()),
			)
		}
	} else if parameter.Mode.ValueString() == "bypass_by_default" {
		if !parameter.Default.IsNull() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				errInvalidConfiguration,
				fmt.Sprintf("cannot set default ttl when using mode '%s'", parameter.Mode.ValueString()),
			)
		}
	}
}

type RulesetActionParameterBrowserTTL struct {
	Mode    basetypes.StringValue `tfsdk:"mode"`
	Default basetypes.Int64Value  `tfsdk:"default"`
}

type BrowserTTLValidator struct{}

func (v BrowserTTLValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("ttl values are required when the override_origin mode is set")
}

func (v BrowserTTLValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("ttl values are required when the override_origin mode is set")
}

func (v BrowserTTLValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	var parameter RulesetActionParameterBrowserTTL

	diag := req.ConfigValue.As(ctx, &parameter, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	if parameter.Mode.ValueString() == "override_origin" {
		if parameter.Default.ValueInt64() <= 0 {

			resp.Diagnostics.AddAttributeError(
				req.Path,
				errInvalidConfiguration,
				fmt.Sprintf("using mode '%s' requires setting a default for ttl", parameter.Mode.ValueString()),
			)
		}
	} else if parameter.Mode.ValueString() == "bypass" {
		if !parameter.Default.IsNull() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				errInvalidConfiguration,
				fmt.Sprintf("cannot set default ttl when using mode '%s'", parameter.Mode.ValueString()),
			)
		}
	}
}
