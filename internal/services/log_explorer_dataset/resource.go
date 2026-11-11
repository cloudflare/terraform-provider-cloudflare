package log_explorer_dataset

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/logs"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*LogExplorerDatasetResource)(nil)
var _ resource.ResourceWithModifyPlan = (*LogExplorerDatasetResource)(nil)
var _ resource.ResourceWithImportState = (*LogExplorerDatasetResource)(nil)

func NewResource() resource.Resource {
	return &LogExplorerDatasetResource{}
}

// LogExplorerDatasetResource defines the resource implementation.
type LogExplorerDatasetResource struct {
	client *cloudflare.Client
}

func (r *LogExplorerDatasetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_log_explorer_dataset"
}

func (r *LogExplorerDatasetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LogExplorerDatasetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *LogExplorerDatasetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	configuredEnabled, configuredProtection := data.Enabled, data.DeletionProtection
	configuredFields := data.Fields
	request, err := createDatasetRequest(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize request", err.Error())
		return
	}
	params := logs.LogExplorerDatasetNewParams{CreateRequest: request}
	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}
	res := new(http.Response)
	env := LogExplorerDatasetResultEnvelope{*data}
	_, err = r.client.Logs.LogExplorer.Datasets.New(ctx, params,
		option.WithResponseBodyInto(&res), option.WithMiddleware(logging.Middleware(ctx)))
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, err := io.ReadAll(res.Body)
	if err == nil {
		err = apijson.Unmarshal(bytes, &env)
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize response", err.Error())
		return
	}
	data = &env.Result
	if data.Filter.IsNull() {
		data.Filter = types.StringValue("")
	}
	data.ID = data.DatasetID
	if err := normalizeResponseFields(ctx, data, configuredFields); err != nil {
		resp.Diagnostics.AddError("failed to read fields", err.Error())
		return
	}
	// Keep the created ID if the update-only settings cannot be applied.
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	update := *data
	update.Fields = configuredFields
	if !configuredEnabled.IsNull() && !configuredEnabled.IsUnknown() {
		update.Enabled = configuredEnabled
	}
	if !configuredProtection.IsNull() && !configuredProtection.IsUnknown() {
		update.DeletionProtection = configuredProtection
	}
	if !update.Enabled.Equal(data.Enabled) || !update.DeletionProtection.Equal(data.DeletionProtection) {
		// Use the full update request so the API does not reset selected fields.
		request, err := updateDatasetRequest(ctx, &update)
		if err == nil {
			data, err = r.updateDataset(ctx, &update, request)
		}
		if err != nil {
			resp.Diagnostics.AddError("failed to apply dataset settings", err.Error())
			return
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *LogExplorerDatasetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *LogExplorerDatasetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request, err := updateDatasetRequest(ctx, data)
	if err == nil {
		data, err = r.updateDataset(ctx, data, request)
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to update dataset", err.Error())
		return
	}
	data.ID = data.DatasetID
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func createDatasetRequest(ctx context.Context, data *LogExplorerDatasetModel) (logs.CreateRequestParam, error) {
	request := logs.CreateRequestParam{Dataset: cloudflare.F(data.Dataset.ValueString())}
	if !data.Filter.IsNull() && !data.Filter.IsUnknown() {
		request.Filter = cloudflare.F(data.Filter.ValueString())
	}
	if !data.FilterAttackTraffic.IsNull() && !data.FilterAttackTraffic.IsUnknown() {
		request.FilterAttackTraffic = cloudflare.F(data.FilterAttackTraffic.ValueBool())
	}
	if !data.Fields.IsNullOrUnknown() {
		fields, diags := data.Fields.AsStructSliceT(ctx)
		if diags.HasError() {
			return request, fmt.Errorf("read fields: %v", diags)
		}
		values := make([]logs.CreateRequestFieldParam, len(fields))
		for i, field := range fields {
			values[i] = logs.CreateRequestFieldParam{Name: cloudflare.F(field.Name.ValueString()), Enabled: cloudflare.F(field.Enabled.ValueBool())}
		}
		request.Fields = cloudflare.F(values)
	}
	return request, nil
}

func updateDatasetRequest(ctx context.Context, data *LogExplorerDatasetModel) (logs.UpdateRequestParam, error) {
	if data.Enabled.IsNull() || data.Enabled.IsUnknown() {
		return logs.UpdateRequestParam{}, errors.New("enabled must be known")
	}
	request := logs.UpdateRequestParam{Enabled: cloudflare.F(data.Enabled.ValueBool())}
	if !data.DeletionProtection.IsNull() && !data.DeletionProtection.IsUnknown() {
		request.DeletionProtection = cloudflare.F(data.DeletionProtection.ValueBool())
	}
	if !data.Filter.IsNull() && !data.Filter.IsUnknown() {
		request.Filter = cloudflare.F(data.Filter.ValueString())
	}
	if !data.FilterAttackTraffic.IsNull() && !data.FilterAttackTraffic.IsUnknown() {
		request.FilterAttackTraffic = cloudflare.F(data.FilterAttackTraffic.ValueBool())
	}
	if !data.Fields.IsNullOrUnknown() {
		fields, diags := data.Fields.AsStructSliceT(ctx)
		if diags.HasError() {
			return request, fmt.Errorf("read fields: %v", diags)
		}
		values := make([]logs.UpdateRequestFieldParam, len(fields))
		for i, field := range fields {
			values[i] = logs.UpdateRequestFieldParam{Name: cloudflare.F(field.Name.ValueString()), Enabled: cloudflare.F(field.Enabled.ValueBool())}
		}
		request.Fields = cloudflare.F(values)
	}
	return request, nil
}

func (r *LogExplorerDatasetResource) updateDataset(ctx context.Context, data *LogExplorerDatasetModel, request logs.UpdateRequestParam) (*LogExplorerDatasetModel, error) {
	params := logs.LogExplorerDatasetUpdateParams{UpdateRequest: request}
	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}
	res := new(http.Response)
	env := LogExplorerDatasetResultEnvelope{*data}
	_, err := r.client.Logs.LogExplorer.Datasets.Update(ctx, data.DatasetID.ValueString(), params,
		option.WithResponseBodyInto(&res), option.WithMiddleware(logging.Middleware(ctx)))
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(res.Body)
	if err == nil {
		err = apijson.Unmarshal(bytes, &env)
	}
	if err == nil {
		if env.Result.Filter.IsNull() {
			env.Result.Filter = types.StringValue("")
		}
		err = normalizeResponseFields(ctx, &env.Result, data.Fields)
	}
	return &env.Result, err
}

func (r *LogExplorerDatasetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *LogExplorerDatasetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := LogExplorerDatasetResultEnvelope{*data}
	params := logs.LogExplorerDatasetGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.Logs.LogExplorer.Datasets.Get(
		ctx,
		data.DatasetID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err == nil {
		if env.Result.Filter.IsNull() {
			env.Result.Filter = types.StringValue("")
		}
		err = normalizeResponseFields(ctx, &env.Result, data.Fields)
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.DatasetID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogExplorerDatasetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *LogExplorerDatasetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := logs.LogExplorerDatasetDeleteParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.Logs.LogExplorer.Datasets.Delete(
		ctx,
		data.DatasetID.ValueString(),
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		var apiErr *cloudflare.Error
		if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusConflict {
			for _, apiError := range apiErr.Errors {
				if apiError.Code == 21002 {
					resp.Diagnostics.AddError("Dataset deletion is protected", "Set `deletion_protection = false` and run `terraform apply` before destroying or replacing this resource.")
					return
				}
			}
		}
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.DatasetID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogExplorerDatasetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(LogExplorerDatasetModel)
	params := logs.LogExplorerDatasetGetParams{}

	path_accounts_or_zones, path_account_id_or_zone_id := "", ""
	path_dataset_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<{accounts|zones}/{account_id|zone_id}>/<dataset_id>",
		&path_accounts_or_zones,
		&path_account_id_or_zone_id,
		&path_dataset_id,
	)
	resp.Diagnostics.Append(diags...)
	switch path_accounts_or_zones {
	case "accounts":
		params.AccountID = cloudflare.F(path_account_id_or_zone_id)
		data.AccountID = types.StringValue(path_account_id_or_zone_id)
	case "zones":
		params.ZoneID = cloudflare.F(path_account_id_or_zone_id)
		data.ZoneID = types.StringValue(path_account_id_or_zone_id)
	default:
		resp.Diagnostics.AddError("invalid discriminator segment - <{accounts|zones}/{account_id|zone_id}>", "expected discriminator to be one of {accounts|zones}")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	data.DatasetID = types.StringValue(path_dataset_id)

	res := new(http.Response)
	env := LogExplorerDatasetResultEnvelope{*data}
	_, err := r.client.Logs.LogExplorer.Datasets.Get(
		ctx,
		path_dataset_id,
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	if data.Filter.IsNull() {
		data.Filter = types.StringValue("")
	}
	data.ID = data.DatasetID
	// The API cannot distinguish a configured selection from its defaults.
	data.Fields = customfield.NullObjectList[LogExplorerDatasetFieldsModel](ctx)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogExplorerDatasetResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var plan, config, state LogExplorerDatasetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Filter.IsNull() {
		if state.Filter.IsNull() || state.Filter.ValueString() == "" {
			plan.Filter = state.Filter
		} else {
			plan.Filter = types.StringValue("")
		}
	}
	if config.Fields.IsNull() {
		plan.Fields = customfield.NullObjectList[LogExplorerDatasetFieldsModel](ctx)
	} else if !config.Fields.IsUnknown() && !state.Fields.IsNullOrUnknown() {
		equal, err := fieldsSemanticallyEqual(ctx, config.Fields, state.Fields)
		if err != nil {
			resp.Diagnostics.AddError("failed to compare fields", err.Error())
			return
		}
		if equal {
			plan.Fields = state.Fields
		}
	}
	// A case-only edit does not change the server's modification time.
	if plan.AccountID.Equal(state.AccountID) && plan.ZoneID.Equal(state.ZoneID) &&
		plan.Dataset.Equal(state.Dataset) && plan.Enabled.Equal(state.Enabled) &&
		plan.DeletionProtection.Equal(state.DeletionProtection) &&
		plan.FilterAttackTraffic.Equal(state.FilterAttackTraffic) &&
		plan.Filter.Equal(state.Filter) && plan.Fields.Equal(state.Fields) {
		plan.UpdatedAt = state.UpdatedAt
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func normalizeResponseFields(ctx context.Context, data *LogExplorerDatasetModel, configured customfield.NestedObjectList[LogExplorerDatasetFieldsModel]) error {
	if configured.IsNullOrUnknown() {
		data.Fields = customfield.NullObjectList[LogExplorerDatasetFieldsModel](ctx)
		return nil
	}
	fields, diags := configured.AsStructSliceT(ctx)
	if diags.HasError() {
		return fmt.Errorf("read configured fields: %v", diags)
	}
	response, diags := data.Fields.AsStructSliceT(ctx)
	if diags.HasError() {
		return fmt.Errorf("read response fields: %v", diags)
	}
	enabled := make(map[string]types.Bool, len(response))
	for _, field := range response {
		enabled[strings.ToLower(field.Name.ValueString())] = field.Enabled
	}
	for i := range fields {
		if value, ok := enabled[strings.ToLower(fields[i].Name.ValueString())]; ok {
			fields[i].Enabled = value
		}
	}
	data.Fields, diags = customfield.NewObjectList(ctx, fields)
	if diags.HasError() {
		return fmt.Errorf("store fields: %v", diags)
	}
	return nil
}

func fieldsSemanticallyEqual(ctx context.Context, left, right customfield.NestedObjectList[LogExplorerDatasetFieldsModel]) (bool, error) {
	a, diags := left.AsStructSliceT(ctx)
	if diags.HasError() {
		return false, fmt.Errorf("read configured fields: %v", diags)
	}
	b, diags := right.AsStructSliceT(ctx)
	if diags.HasError() {
		return false, fmt.Errorf("read previous fields: %v", diags)
	}
	if len(a) != len(b) {
		return false, nil
	}
	for i := range a {
		if a[i].Name.IsNull() || a[i].Name.IsUnknown() || b[i].Name.IsNull() || b[i].Name.IsUnknown() ||
			!strings.EqualFold(a[i].Name.ValueString(), b[i].Name.ValueString()) || !a[i].Enabled.Equal(b[i].Enabled) {
			return false, nil
		}
	}
	return true, nil
}
