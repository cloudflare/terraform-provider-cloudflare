package log_explorer_dataset_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogExplorerDatasetReadMatchesFieldNamesCaseInsensitively(t *testing.T) {
	ctx := context.Background()
	model := knownDatasetModel(ctx)
	model.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
		{Name: types.StringValue("QueryName"), Enabled: types.BoolValue(false)},
		{Name: types.StringValue("retired_field"), Enabled: types.BoolValue(false)},
	})
	state := datasetState(t, ctx, model)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeDatasetResponse(t, w, datasetResponse{
			Enabled: true, DeletionProtection: true, ObjectType: "zone", ObjectID: testZoneID,
			Fields: []map[string]any{
				{"name": "queryname", "enabled": true},
				{"name": "timestamp", "enabled": true},
			},
		})
	}))
	defer ts.Close()

	resp := &resource.ReadResponse{State: state}
	newDatasetResource(t, ts.URL).Read(ctx, resource.ReadRequest{State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read returned diagnostics: %v", resp.Diagnostics)
	}
	var got log_explorer_dataset.LogExplorerDatasetModel
	resp.State.Get(ctx, &got)
	gotFields, diags := got.Fields.AsStructSliceT(ctx)
	if diags.HasError() {
		t.Fatalf("failed to read fields: %v", diags)
	}
	if len(gotFields) != 2 || gotFields[0].Name.ValueString() != "QueryName" || !gotFields[0].Enabled.ValueBool() {
		t.Errorf("case-insensitive API field did not refresh configured projection: %#v", gotFields)
	}
	if gotFields[1].Name.ValueString() != "retired_field" || gotFields[1].Enabled.ValueBool() {
		t.Errorf("unknown disabled field must remain ignored in state: %#v", gotFields[1])
	}
}

func TestLogExplorerDatasetModifyPlanSuppressesFieldCaseOnlyChange(t *testing.T) {
	ctx := context.Background()
	stateModel := knownDatasetModel(ctx)
	stateModel.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
		{Name: types.StringValue("queryname"), Enabled: types.BoolValue(true)},
		{Name: types.StringValue("timestamp"), Enabled: types.BoolValue(true)},
	})
	configModel := stateModel
	configModel.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
		{Name: types.StringValue("QueryName"), Enabled: types.BoolValue(true)},
		{Name: types.StringValue("timestamp"), Enabled: types.BoolValue(true)},
	})
	state := datasetState(t, ctx, stateModel)
	plan := datasetPlan(t, ctx, configModel)
	configState := datasetState(t, ctx, configModel)
	config := tfsdk.Config{Schema: configState.Schema, Raw: configState.Raw}

	resp := &resource.ModifyPlanResponse{Plan: plan}
	log_explorer_dataset.NewResource().(*log_explorer_dataset.LogExplorerDatasetResource).ModifyPlan(
		ctx,
		resource.ModifyPlanRequest{Config: config, Plan: plan, State: state},
		resp,
	)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ModifyPlan returned diagnostics: %v", resp.Diagnostics)
	}

	var got log_explorer_dataset.LogExplorerDatasetModel
	resp.Plan.Get(ctx, &got)
	if !got.Fields.Equal(stateModel.Fields) {
		t.Fatalf("case-only field change remained in plan: %#v", got.Fields)
	}
}

func TestLogExplorerDatasetUpdateNullFieldsRestoresAPISelection(t *testing.T) {
	ctx := context.Background()
	stateModel := knownDatasetModel(ctx)
	stateModel.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{{
		Name: types.StringValue("timestamp"), Enabled: types.BoolValue(true),
	}})
	planModel := stateModel
	planModel.Fields = customfield.NullObjectList[log_explorer_dataset.LogExplorerDatasetFieldsModel](ctx)
	state := datasetState(t, ctx, stateModel)
	plan := datasetPlan(t, ctx, planModel)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := decodeRequestBody(t, r)
		if fields := body["fields"]; fields != nil {
			t.Errorf("got fields %#v, want omitted or null reset", fields)
		}
		writeDatasetResponse(t, w, datasetResponse{
			Enabled: true, DeletionProtection: true, ObjectType: "zone", ObjectID: testZoneID,
			Fields: []map[string]any{
				{"name": "timestamp", "enabled": true},
				{"name": "queryname", "enabled": true},
			},
		})
	}))
	defer ts.Close()

	resp := &resource.UpdateResponse{State: state}
	newDatasetResource(t, ts.URL).Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Update returned diagnostics: %v", resp.Diagnostics)
	}
	var got log_explorer_dataset.LogExplorerDatasetModel
	resp.State.Get(ctx, &got)
	if !got.Fields.IsNull() {
		t.Errorf("fields are %v, want null after API resets its full selection", got.Fields)
	}
}

func TestLogExplorerDatasetImportKeepsFieldsUnmanaged(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/zones/"+testZoneID+"/logs/explorer/datasets/"+testDatasetID {
			t.Errorf("got %s %s, want imported dataset GET", r.Method, r.URL.Path)
		}
		writeDatasetResponse(t, w, datasetResponse{
			Enabled: true, DeletionProtection: true, ObjectType: "zone", ObjectID: testZoneID,
			Fields: []map[string]any{{"name": "timestamp", "enabled": true}},
		})
	}))
	defer ts.Close()

	resp := &resource.ImportStateResponse{State: tfsdk.State{Schema: log_explorer_dataset.ResourceSchema(ctx)}}
	newDatasetResource(t, ts.URL).ImportState(ctx, resource.ImportStateRequest{ID: "zones/" + testZoneID + "/" + testDatasetID}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ImportState returned diagnostics: %v", resp.Diagnostics)
	}
	var got log_explorer_dataset.LogExplorerDatasetModel
	resp.State.Get(ctx, &got)
	if got.ID.ValueString() != testDatasetID || got.DatasetID.ValueString() != testDatasetID || !got.Fields.IsNull() {
		t.Errorf("imported id=%q dataset_id=%q fields=%v", got.ID.ValueString(), got.DatasetID.ValueString(), got.Fields)
	}
}

func TestLogExplorerDatasetFieldNamesAreUnique(t *testing.T) {
	for _, tc := range []struct {
		name    string
		invalid bool
	}{
		{"queryname", true}, {"QueryName", true}, {"timestamp", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			fields := customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
				{Name: types.StringValue("queryname"), Enabled: types.BoolValue(true)},
				{Name: types.StringValue(tc.name), Enabled: types.BoolValue(false)},
			})
			attribute := log_explorer_dataset.ResourceSchema(ctx).Attributes["fields"].(schema.ListNestedAttribute)
			resp := &validator.ListResponse{}
			for _, v := range attribute.Validators {
				v.ValidateList(ctx, validator.ListRequest{Path: path.Root("fields"), ConfigValue: fields.ListValue}, resp)
			}
			if resp.Diagnostics.HasError() != tc.invalid {
				t.Fatalf("invalid = %t, want %t: %v", resp.Diagnostics.HasError(), tc.invalid, resp.Diagnostics)
			}
		})
	}
}

func TestLogExplorerDatasetFieldResponsesPreserveProjectionAndDrift(t *testing.T) {
	for _, operation := range []string{"create", "update", "read"} {
		t.Run(operation, func(t *testing.T) {
			ctx := context.Background()
			model := knownDatasetModel(ctx)
			model.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
				{Name: types.StringValue("QueryName"), Enabled: types.BoolValue(true)},
				{Name: types.StringValue("Timestamp"), Enabled: types.BoolValue(false)},
			})
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				writeDatasetResponse(t, w, datasetResponse{Enabled: true, DeletionProtection: true, ObjectType: "zone", ObjectID: testZoneID,
					Fields: []map[string]any{{"name": "timestamp", "enabled": true}, {"name": "querytype", "enabled": true}, {"name": "queryname", "enabled": false}},
				})
			}))
			defer server.Close()
			r := newDatasetResource(t, server.URL)
			state := datasetState(t, ctx, model)
			plan := datasetPlan(t, ctx, model)
			switch operation {
			case "create":
				resp := &resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
				r.Create(ctx, resource.CreateRequest{Plan: plan}, resp)
				if resp.Diagnostics.HasError() {
					t.Fatal(resp.Diagnostics)
				}
				state = resp.State
			case "update":
				resp := &resource.UpdateResponse{State: state}
				r.Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, resp)
				if resp.Diagnostics.HasError() {
					t.Fatal(resp.Diagnostics)
				}
				state = resp.State
			case "read":
				resp := &resource.ReadResponse{State: state}
				r.Read(ctx, resource.ReadRequest{State: state}, resp)
				if resp.Diagnostics.HasError() {
					t.Fatal(resp.Diagnostics)
				}
				state = resp.State
			}
			var got log_explorer_dataset.LogExplorerDatasetModel
			if diags := state.Get(ctx, &got); diags.HasError() {
				t.Fatal(diags)
			}
			fields, diags := got.Fields.AsStructSliceT(ctx)
			if diags.HasError() {
				t.Fatal(diags)
			}
			if len(fields) != 2 || fields[0].Name.ValueString() != "QueryName" || fields[0].Enabled.ValueBool() ||
				fields[1].Name.ValueString() != "Timestamp" || !fields[1].Enabled.ValueBool() {
				t.Fatalf("API drift or configured projection lost: %v", fields)
			}
		})
	}
}

func TestLogExplorerDatasetModifyPlanKeepsEnabledChange(t *testing.T) {
	ctx := context.Background()
	model := knownDatasetModel(ctx)
	model.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
		{Name: types.StringValue("queryname"), Enabled: types.BoolValue(true)},
	})
	state := datasetState(t, ctx, model)
	model.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
		{Name: types.StringValue("QueryName"), Enabled: types.BoolValue(false)},
	})
	plan := datasetPlan(t, ctx, model)
	resp := &resource.ModifyPlanResponse{Plan: plan}
	log_explorer_dataset.NewResource().(*log_explorer_dataset.LogExplorerDatasetResource).ModifyPlan(ctx,
		resource.ModifyPlanRequest{State: state, Plan: plan, Config: tfsdk.Config{Schema: plan.Schema, Raw: plan.Raw}}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	var got log_explorer_dataset.LogExplorerDatasetModel
	if diags := resp.Plan.Get(ctx, &got); diags.HasError() {
		t.Fatal(diags)
	}
	if !got.Fields.Equal(model.Fields) {
		t.Fatalf("enabled change was suppressed: %v", got.Fields)
	}
}
