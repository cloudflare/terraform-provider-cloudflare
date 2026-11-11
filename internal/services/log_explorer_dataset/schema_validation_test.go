package log_explorer_dataset_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogExplorerDatasetScopeValidation(t *testing.T) {
	for name, tc := range map[string]struct {
		account, zone types.String
		invalid       bool
	}{
		"account": {account: types.StringValue("account"), zone: types.StringNull()},
		"zone":    {account: types.StringNull(), zone: types.StringValue("zone")},
		"neither": {account: types.StringNull(), zone: types.StringNull(), invalid: true},
		"both":    {account: types.StringValue("account"), zone: types.StringValue("zone"), invalid: true},
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			model := log_explorer_dataset.LogExplorerDatasetModel{
				AccountID: tc.account, ZoneID: tc.zone, Dataset: types.StringValue("dns_logs"),
				Fields: customfield.NullObjectList[log_explorer_dataset.LogExplorerDatasetFieldsModel](ctx),
			}
			state := tfsdk.State{Schema: log_explorer_dataset.ResourceSchema(ctx)}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatal(diags)
			}
			r := log_explorer_dataset.NewResource().(resource.ResourceWithConfigValidators)
			resp := &resource.ValidateConfigResponse{}
			for _, v := range r.ConfigValidators(ctx) {
				v.ValidateResource(ctx, resource.ValidateConfigRequest{Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw}}, resp)
			}
			if resp.Diagnostics.HasError() != tc.invalid {
				t.Fatalf("invalid = %t, want %t: %v", resp.Diagnostics.HasError(), tc.invalid, resp.Diagnostics)
			}
		})
	}
}

func TestLogExplorerDatasetsScopeValidation(t *testing.T) {
	for name, tc := range map[string]struct {
		account, zone types.String
		include       types.Bool
		invalid       bool
	}{
		"account":         {account: types.StringValue("account"), zone: types.StringNull(), include: types.BoolNull()},
		"account include": {account: types.StringValue("account"), zone: types.StringNull(), include: types.BoolValue(true)},
		"account exclude": {account: types.StringValue("account"), zone: types.StringNull(), include: types.BoolValue(false)},
		"zone":            {account: types.StringNull(), zone: types.StringValue("zone"), include: types.BoolNull()},
		"zone include":    {account: types.StringNull(), zone: types.StringValue("zone"), include: types.BoolValue(true), invalid: true},
		"zone exclude":    {account: types.StringNull(), zone: types.StringValue("zone"), include: types.BoolValue(false), invalid: true},
		"neither":         {account: types.StringNull(), zone: types.StringNull(), include: types.BoolNull(), invalid: true},
		"both":            {account: types.StringValue("account"), zone: types.StringValue("zone"), include: types.BoolNull(), invalid: true},
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			model := log_explorer_dataset.LogExplorerDatasetsDataSourceModel{
				AccountID: tc.account, ZoneID: tc.zone, IncludeZones: tc.include,
				Result: customfield.NullObjectList[log_explorer_dataset.LogExplorerDatasetsResultDataSourceModel](ctx),
			}
			state := tfsdk.State{Schema: log_explorer_dataset.ListDataSourceSchema(ctx)}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatal(diags)
			}
			d := log_explorer_dataset.NewLogExplorerDatasetsDataSource().(datasource.DataSourceWithConfigValidators)
			resp := &datasource.ValidateConfigResponse{}
			for _, v := range d.ConfigValidators(ctx) {
				v.ValidateDataSource(ctx, datasource.ValidateConfigRequest{Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw}}, resp)
			}
			if resp.Diagnostics.HasError() != tc.invalid {
				t.Fatalf("invalid = %t, want %t: %v", resp.Diagnostics.HasError(), tc.invalid, resp.Diagnostics)
			}
		})
	}
}
