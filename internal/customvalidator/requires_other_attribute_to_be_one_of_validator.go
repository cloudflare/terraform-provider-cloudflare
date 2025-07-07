package customvalidator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

func RequiresOtherStringAttributeToBeOneOf(pathExpr path.Expression, wantStrValues ...string) requiresOtherAttributeToBeOneOfValidator {
	var wantValues []attr.Value
	for _, v := range wantStrValues {
		wantValues = append(wantValues, types.StringValue(v))
	}
	return requiresOtherAttributeToBeOneOfValidator{
		pathExpr,
		wantValues,
	}
}

func RequiresOtherStringAttributeToNullOrBeOneOf(pathExpr path.Expression, wantStrValues ...string) requiresOtherAttributeToBeOneOfValidator {
	v := RequiresOtherStringAttributeToBeOneOf(pathExpr, wantStrValues...)
	v.wantValues = append(v.wantValues, types.StringNull())
	return v
}

type requiresOtherAttributeToBeOneOfValidator struct {
	pathExpr   path.Expression
	wantValues []attr.Value
}

func (i requiresOtherAttributeToBeOneOfValidator) Description(ctx context.Context) string {
	var wantValuesAsStrings []string
	for _, v := range i.wantValues {
		wantValuesAsStrings = append(wantValuesAsStrings, v.String())
	}
	return fmt.Sprintf("can only be set if %q is one of: %s", i.pathExpr, strings.Join(wantValuesAsStrings, ", "))
}

func (i requiresOtherAttributeToBeOneOfValidator) MarkdownDescription(ctx context.Context) string {
	return i.Description(ctx)
}

func (i requiresOtherAttributeToBeOneOfValidator) validateAny(ctx context.Context, cfg *tfsdk.Config, attrPathExpr path.Expression, attrPath path.Path, value attr.Value, resDiagnostics *diag.Diagnostics) {
	if value.IsNull() || value.IsUnknown() {
		return
	}

	expression := attrPathExpr.Merge(i.pathExpr)
	matchedPaths, diags := cfg.PathMatches(ctx, expression)
	resDiagnostics.Append(diags...)

	for _, mp := range matchedPaths {
		// If the user specifies the same attribute this validator is applied to,
		// also as part of the input, skip it
		if mp.Equal(attrPath) {
			continue
		}

		var mpVal attr.Value
		resDiagnostics.Append(cfg.GetAttribute(ctx, mp, &mpVal)...)

		// Collect all errors
		if diags.HasError() {
			continue
		}

		// Delay validation until all involved attribute have a known value
		if mpVal.IsUnknown() {
			return
		}

		foundMatch := false
		for _, wantValue := range i.wantValues {
			if mpVal.Equal(wantValue) {
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			description := fmt.Sprintf("%q %s", attrPath, i.Description(ctx))
			resDiagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
				attrPath,
				description,
			))
		}
	}
}

func (i requiresOtherAttributeToBeOneOfValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, res *validator.BoolResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeOneOfValidator) ValidateString(ctx context.Context, req validator.StringRequest, res *validator.StringResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeOneOfValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, res *validator.ObjectResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeOneOfValidator) ValidateList(ctx context.Context, req validator.ListRequest, res *validator.ListResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}
