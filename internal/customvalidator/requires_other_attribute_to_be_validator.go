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
)

func RequiresOtherStringAttributeToBe(pathExpr path.Expression, wantStrValue string) requiresOtherAttributeToBeValidator {
	return requiresOtherAttributeToBeValidator{
		pathExpr,
		types.StringValue(wantStrValue),
	}
}

type requiresOtherAttributeToBeValidator struct {
	pathExpr  path.Expression
	wantValue attr.Value
}

func (i requiresOtherAttributeToBeValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("can only be set if %q is %s", i.pathExpr, i.wantValue)
}

func (i requiresOtherAttributeToBeValidator) MarkdownDescription(ctx context.Context) string {
	return i.Description(ctx)
}

func (i requiresOtherAttributeToBeValidator) validateAny(ctx context.Context, cfg *tfsdk.Config, attrPathExpr path.Expression, attrPath path.Path, value attr.Value, resDiagnostics *diag.Diagnostics) {
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

		if !mpVal.Equal(i.wantValue) {
			description := fmt.Sprintf("%q %s", attrPath, i.Description(ctx))
			resDiagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
				attrPath,
				description,
			))
		}
	}
}

func (i requiresOtherAttributeToBeValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, res *validator.BoolResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateFloat32(ctx context.Context, req validator.Float32Request, res *validator.Float32Response) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, res *validator.Float64Response) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateInt32(ctx context.Context, req validator.Int32Request, res *validator.Int32Response) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, res *validator.Int64Response) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateList(ctx context.Context, req validator.ListRequest, res *validator.ListResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateMap(ctx context.Context, req validator.MapRequest, res *validator.MapResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, res *validator.NumberResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, res *validator.ObjectResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateSet(ctx context.Context, req validator.SetRequest, res *validator.SetResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateString(ctx context.Context, req validator.StringRequest, res *validator.StringResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}

func (i requiresOtherAttributeToBeValidator) ValidateDynamic(ctx context.Context, req validator.DynamicRequest, res *validator.DynamicResponse) {
	i.validateAny(ctx, &req.Config, req.PathExpression, req.Path, req.ConfigValue, &res.Diagnostics)
}
