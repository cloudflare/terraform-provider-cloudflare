package zero_trust_access_application

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func targetRuleTags(values ...string) *map[string]*[]types.String {
	items := make([]types.String, len(values))
	for i, value := range values {
		items[i] = types.StringValue(value)
	}
	return &map[string]*[]types.String{"environment": &items}
}

func TestTargetCriteriaRuleValidation(t *testing.T) {
	ctx := context.Background()
	attrs := &ZeroTrustAccessApplicationTargetRuleAttributesModel{Hostname: &[]types.String{types.StringValue("web-*")}}
	tagRule := &ZeroTrustAccessApplicationTargetRuleModel{Tags: targetRuleTags("production")}
	hostRule := &ZeroTrustAccessApplicationTargetRuleModel{TargetAttributes: attrs}
	legacy := &map[string]*[]types.String{"hostname": &[]types.String{types.StringValue("server")}}
	tooManyKeys := map[string]*[]types.String{}
	for i := 0; i < 51; i++ {
		tooManyKeys[fmt.Sprintf("key%d", i)] = &[]types.String{types.StringValue("value")}
	}
	tooManyHosts := make([]types.String, 5001)
	for i := range tooManyHosts {
		tooManyHosts[i] = types.StringValue("server")
	}
	tests := []struct {
		name      string
		criteria  ZeroTrustAccessApplicationTargetCriteriaModel
		wantError bool
	}{
		{name: "legacy", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{TargetAttributes: legacy}},
		{name: "legacy non-hostname map deferred to API", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{TargetAttributes: &map[string]*[]types.String{"environment": &[]types.String{types.StringValue("production")}}}},
		{name: "tags only", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: tagRule}},
		{name: "hostnames only", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: hostRule}},
		{name: "all rules", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: hostRule, Require: tagRule, Exclude: tagRule}},
		{name: "combined selectors", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{TargetAttributes: attrs, Tags: targetRuleTags("production")}}},
		{name: "missing selectors", wantError: true},
		{name: "require without include", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Require: tagRule}, wantError: true},
		{name: "exclude with legacy", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{TargetAttributes: legacy, Exclude: tagRule}, wantError: true},
		{name: "require with legacy", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{TargetAttributes: legacy, Require: tagRule}, wantError: true},
		{name: "exclude without include", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Exclude: tagRule}, wantError: true},
		{name: "empty tag map", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{Tags: &map[string]*[]types.String{}}}, wantError: true},
		{name: "empty tag key", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{Tags: &map[string]*[]types.String{"": &[]types.String{types.StringValue("production")}}}}, wantError: true},
		{name: "too many tag keys", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{Tags: &tooManyKeys}}, wantError: true},
		{name: "empty tag values", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{Tags: targetRuleTags()}}, wantError: true},
		{name: "empty tag string", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{Tags: targetRuleTags("")}}, wantError: true},
		{name: "too many tag values", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{Tags: targetRuleTags(strings.Split(strings.Repeat("value,", 50)+"value", ",")...)}}, wantError: true},
		{name: "empty hostnames", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{TargetAttributes: &ZeroTrustAccessApplicationTargetRuleAttributesModel{Hostname: &[]types.String{}}}}, wantError: true},
		{name: "empty hostname string", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{TargetAttributes: &ZeroTrustAccessApplicationTargetRuleAttributesModel{Hostname: &[]types.String{types.StringValue("")}}}}, wantError: true},
		{name: "too many hostnames", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{TargetAttributes: &ZeroTrustAccessApplicationTargetRuleAttributesModel{Hostname: &tooManyHosts}}}, wantError: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := validateTargetCriteriaRules(t, ctx, tt.criteria)
			require.Equal(t, tt.wantError, diags.HasError(), "%v", diags)
		})
	}
}

func TestTargetCriteriaMixedFormatsDiagnostic(t *testing.T) {
	criteria := ZeroTrustAccessApplicationTargetCriteriaModel{
		TargetAttributes: &map[string]*[]types.String{"hostname": &[]types.String{types.StringValue("server")}},
		Include:          &ZeroTrustAccessApplicationTargetRuleModel{Tags: targetRuleTags("production")},
	}
	diags := validateTargetCriteriaRules(t, context.Background(), criteria)
	require.Len(t, diags.Errors(), 1, "%v", diags)
}

func TestTargetCriteriaEmptyRuleDiagnostics(t *testing.T) {
	ctx := context.Background()
	tagRule := &ZeroTrustAccessApplicationTargetRuleModel{Tags: targetRuleTags("production")}
	tests := []struct {
		name     string
		criteria ZeroTrustAccessApplicationTargetCriteriaModel
	}{
		{name: "include", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: &ZeroTrustAccessApplicationTargetRuleModel{}}},
		{name: "require", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: tagRule, Require: &ZeroTrustAccessApplicationTargetRuleModel{}}},
		{name: "exclude", criteria: ZeroTrustAccessApplicationTargetCriteriaModel{Include: tagRule, Exclude: &ZeroTrustAccessApplicationTargetRuleModel{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := validateTargetCriteriaRules(t, ctx, tt.criteria)
			require.Len(t, diags.Errors(), 1, "%v", diags)
			diagnostic, ok := diags.Errors()[0].(diag.DiagnosticWithPath)
			require.True(t, ok, "expected a diagnostic with an attribute path, got %T", diags.Errors()[0])
			require.True(t, path.Root("target_criteria").AtListIndex(0).AtName(tt.name).Equal(diagnostic.Path()), "%v", diagnostic.Path())
		})
	}
}

func validateTargetCriteriaRules(t *testing.T, ctx context.Context, criteria ZeroTrustAccessApplicationTargetCriteriaModel) diag.Diagnostics {
	t.Helper()
	criteria.Port = types.Int64Value(22)
	criteria.Protocol = types.StringValue("SSH")
	criteriaSchema, ok := ResourceSchema(ctx).Attributes["target_criteria"].(schema.ListNestedAttribute)
	require.True(t, ok, "expected target_criteria to be a ListNestedAttribute")
	state := tfsdk.State{Schema: schema.Schema{Attributes: map[string]schema.Attribute{"target_criteria": criteriaSchema}}}
	diags := state.Set(ctx, &struct {
		TargetCriteria []*ZeroTrustAccessApplicationTargetCriteriaModel `tfsdk:"target_criteria"`
	}{[]*ZeroTrustAccessApplicationTargetCriteriaModel{&criteria}})
	require.False(t, diags.HasError(), "%v", diags)
	config := tfsdk.Config{Raw: state.Raw, Schema: state.Schema}
	for _, name := range []string{"include", "require", "exclude"} {
		ruleSchema, ok := criteriaSchema.NestedObject.Attributes[name].(schema.SingleNestedAttribute)
		require.True(t, ok, "expected %s to be a SingleNestedAttribute", name)
		p := path.Root("target_criteria").AtListIndex(0).AtName(name)
		expr := path.MatchRoot("target_criteria").AtListIndex(0).AtName(name)
		diags.Append(validateTargetRule(t, ctx, config, ruleSchema, p, expr)...)
	}
	return diags
}

func validateTargetRule(t *testing.T, ctx context.Context, config tfsdk.Config, ruleSchema schema.SingleNestedAttribute, p path.Path, expr path.Expression) diag.Diagnostics {
	t.Helper()
	var diags diag.Diagnostics
	var value types.Object
	require.False(t, config.GetAttribute(ctx, p, &value).HasError())
	for _, v := range ruleSchema.Validators {
		resp := validator.ObjectResponse{}
		v.ValidateObject(ctx, validator.ObjectRequest{Config: config, ConfigValue: value, Path: p, PathExpression: expr}, &resp)
		diags.Append(resp.Diagnostics...)
	}
	if value.IsNull() || value.IsUnknown() {
		return diags
	}
	var selectors types.Object
	require.False(t, config.GetAttribute(ctx, p.AtName("target_attributes"), &selectors).HasError())
	if !selectors.IsNull() && !selectors.IsUnknown() {
		selectorSchema, ok := ruleSchema.Attributes["target_attributes"].(schema.SingleNestedAttribute)
		require.True(t, ok, "expected target_attributes to be a SingleNestedAttribute")
		hostSchema, ok := selectorSchema.Attributes["hostname"].(schema.ListAttribute)
		require.True(t, ok, "expected hostname to be a ListAttribute")
		var hosts types.List
		require.False(t, config.GetAttribute(ctx, p.AtName("target_attributes").AtName("hostname"), &hosts).HasError())
		for _, v := range hostSchema.Validators {
			resp := validator.ListResponse{}
			v.ValidateList(ctx, validator.ListRequest{Config: config, ConfigValue: hosts, Path: p.AtName("target_attributes").AtName("hostname"), PathExpression: expr.AtName("target_attributes").AtName("hostname")}, &resp)
			diags.Append(resp.Diagnostics...)
		}
	}
	var tags types.Map
	require.False(t, config.GetAttribute(ctx, p.AtName("tags"), &tags).HasError())
	tagSchema, ok := ruleSchema.Attributes["tags"].(schema.MapAttribute)
	require.True(t, ok, "expected tags to be a MapAttribute")
	for _, v := range tagSchema.Validators {
		resp := validator.MapResponse{}
		v.ValidateMap(ctx, validator.MapRequest{Config: config, ConfigValue: tags, Path: p.AtName("tags"), PathExpression: expr.AtName("tags")}, &resp)
		diags.Append(resp.Diagnostics...)
	}
	return diags
}

func TestTargetRuleSelectorsValidationDeferred(t *testing.T) {
	hostType := types.ObjectType{AttrTypes: map[string]attr.Type{"hostname": types.ListType{ElemType: types.StringType}}}
	tagType := types.MapType{ElemType: types.ListType{ElemType: types.StringType}}
	ruleTypes := map[string]attr.Type{"target_attributes": hostType, "tags": tagType}
	tests := []struct {
		name  string
		value types.Object
	}{
		{name: "null rule", value: types.ObjectNull(ruleTypes)},
		{name: "unknown rule", value: types.ObjectUnknown(ruleTypes)},
		{name: "unknown hostname selector", value: types.ObjectValueMust(ruleTypes, map[string]attr.Value{
			"target_attributes": types.ObjectUnknown(hostType.AttrTypes),
			"tags":              types.MapNull(tagType.ElemType),
		})},
		{name: "unknown tags selector", value: types.ObjectValueMust(ruleTypes, map[string]attr.Value{
			"target_attributes": types.ObjectNull(hostType.AttrTypes),
			"tags":              types.MapUnknown(tagType.ElemType),
		})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := validator.ObjectResponse{}
			targetRuleSelectorsValidator{}.ValidateObject(context.Background(), validator.ObjectRequest{
				ConfigValue: tt.value,
				Path:        path.Root("target_criteria").AtListIndex(0).AtName("include"),
			}, &resp)
			require.Empty(t, resp.Diagnostics)
		})
	}
}

func TestTargetCriteriaRulesJSON(t *testing.T) {
	for _, protocol := range []string{"SSH", "RDP"} {
		t.Run(protocol, func(t *testing.T) {
			body := `{"target_criteria":[{"port":22,"protocol":"` + protocol + `","include":{"target_attributes":{"hostname":["web-*"]},"tags":{"environment":["production","staging"]}},"require":{"tags":{"team":["engineering"]}},"exclude":{"tags":{"environment":["development"]}}}]}`
			var model ZeroTrustAccessApplicationModel
			require.NoError(t, apijson.Unmarshal([]byte(body), &model))
			encoded, err := model.MarshalJSON()
			require.NoError(t, err)
			var result map[string]json.RawMessage
			require.NoError(t, json.Unmarshal(encoded, &result))
			var expected map[string]json.RawMessage
			require.NoError(t, json.Unmarshal([]byte(body), &expected))
			require.JSONEq(t, string(expected["target_criteria"]), string(result["target_criteria"]))

			// Create and Update decode computed fields into the planned model.
			envelope := ZeroTrustAccessApplicationResultEnvelope{Result: model}
			require.NoError(t, apijson.UnmarshalComputed([]byte(`{"result":`+body+`}`), &envelope))
			computedBody, err := envelope.Result.MarshalJSON()
			require.NoError(t, err)
			require.NoError(t, json.Unmarshal(computedBody, &result))
			require.JSONEq(t, string(expected["target_criteria"]), string(result["target_criteria"]))

			var single ZeroTrustAccessApplicationDataSourceModel
			require.NoError(t, apijson.UnmarshalComputed([]byte(body), &single))
			var listItem ZeroTrustAccessApplicationsResultDataSourceModel
			require.NoError(t, apijson.UnmarshalComputed([]byte(body), &listItem))
			require.False(t, single.TargetCriteria.IsNull())
			require.False(t, listItem.TargetCriteria.IsNull())
			singleValue, err := single.TargetCriteria.ToTerraformValue(context.Background())
			require.NoError(t, err)
			listValue, err := listItem.TargetCriteria.ToTerraformValue(context.Background())
			require.NoError(t, err)
			require.True(t, singleValue.Equal(listValue))
			require.Len(t, single.TargetCriteria.Elements(), 1)
			criteria, ok := single.TargetCriteria.Elements()[0].(types.Object)
			require.True(t, ok, "expected target criteria to be an Object, got %T", single.TargetCriteria.Elements()[0])
			for _, name := range []string{"include", "require", "exclude"} {
				require.False(t, criteria.Attributes()[name].IsNull(), name)
			}
		})
	}
}

func TestTargetCriteriaFormatUpdates(t *testing.T) {
	legacy := `{"target_criteria":[{"port":22,"protocol":"SSH","target_attributes":{"hostname":["server"]}}]}`
	rules := `{"target_criteria":[{"port":22,"protocol":"SSH","include":{"tags":{"environment":["production"]}}}]}`
	for _, tt := range []struct{ name, before, after string }{
		{"legacy to rules", legacy, rules},
		{"rules to legacy", rules, legacy},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var state, plan ZeroTrustAccessApplicationModel
			require.NoError(t, apijson.Unmarshal([]byte(tt.before), &state))
			require.NoError(t, apijson.Unmarshal([]byte(tt.after), &plan))
			body, err := plan.MarshalJSONForUpdate(state)
			require.NoError(t, err)
			var expected, actual map[string]json.RawMessage
			require.NoError(t, json.Unmarshal([]byte(tt.after), &expected))
			require.NoError(t, json.Unmarshal(body, &actual))
			require.JSONEq(t, string(expected["target_criteria"]), string(actual["target_criteria"]))
		})
	}
}
