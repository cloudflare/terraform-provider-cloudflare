package rulesets

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
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

	if req.ConfigValue.ValueString() == string(cloudflare.RulesetPhaseSuperBotFightMode) {
		resp.Diagnostics.AddAttributeWarning(
			req.Path,
			fmt.Sprintf(`%q phase will soon be deprecated in the "cloudflare_ruleset" resource`, string(cloudflare.RulesetPhaseSuperBotFightMode)),
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
	}
}

type InvalidWildCardValidator struct{}

func (v InvalidWildCardValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("full wildcards should use the ignore field instead")
}

func (v InvalidWildCardValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("full wildcards should use the ignore field instead")
}

func (v InvalidWildCardValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	_, ok := req.ConfigValue.ElementType(ctx).(basetypes.StringTypable)

	if !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Validator for Element Type",
			"While performing schema-based validation, an unexpected error occurred. "+
				"The attribute declares a String values validator, however its values do not implement types.StringType or the types.StringTypable interface for custom String types. "+
				"Use the appropriate values validator that matches the element type. "+
				"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
				fmt.Sprintf("Path: %s\n", req.Path.String())+
				fmt.Sprintf("Element Type: %T\n", req.ConfigValue.ElementType(ctx)),
		)

		return
	}

	for _, element := range req.ConfigValue.Elements() {
		// ignore okay check, error handling above should catch it
		elementValuable, _ := element.(basetypes.StringValuable)

		elementValue, diags := elementValuable.ToStringValue(ctx)
		resp.Diagnostics.Append(diags...)

		if elementValue.ValueString() == "*" {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				errInvalidValue,
				fmt.Sprintf("full wildcards should use the ignore field instead, value: %s", elementValue.ValueString()),
			)
		}
	}
}
