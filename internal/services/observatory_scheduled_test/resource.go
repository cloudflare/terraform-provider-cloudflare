// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/speed"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ObservatoryScheduledTestResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ObservatoryScheduledTestResource)(nil)
var _ resource.ResourceWithImportState = (*ObservatoryScheduledTestResource)(nil)

func NewResource() resource.Resource {
	return &ObservatoryScheduledTestResource{}
}

// ObservatoryScheduledTestResource defines the resource implementation.
type ObservatoryScheduledTestResource struct {
	client *cloudflare.Client
}

func (r *ObservatoryScheduledTestResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_observatory_scheduled_test"
}

func (r *ObservatoryScheduledTestResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ObservatoryScheduledTestResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ObservatoryScheduledTestModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := speed.ScheduleNewParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
	}

	if !data.Frequency.IsNull() && !data.Frequency.IsUnknown() {
		params.Frequency = cloudflare.F(speed.ScheduleNewParamsFrequency(data.Frequency.ValueString()))
	}

	if !data.Region.IsNull() && !data.Region.IsUnknown() {
		params.Region = cloudflare.F(speed.ScheduleNewParamsRegion(data.Region.ValueString()))
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}

	res := new(http.Response)
	env := ObservatoryScheduledTestResultEnvelope{*data}
	_, err = r.client.Speed.Schedule.New(
		ctx,
		data.URL.ValueString(),
		params,
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.URL

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ObservatoryScheduledTestResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *ObservatoryScheduledTestResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ObservatoryScheduledTestModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := speed.ScheduleGetParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
	}

	if !data.Region.IsNull() && !data.Region.IsUnknown() {
		params.Region = cloudflare.F(speed.ScheduleGetParamsRegion(data.Region.ValueString()))
	}

	res := new(http.Response)
	env := ObservatoryScheduledTestResultEnvelope{*data}
	_, err := r.client.Speed.Schedule.Get(
		ctx,
		data.URL.ValueString(),
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
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.URL

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ObservatoryScheduledTestResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ObservatoryScheduledTestModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := speed.ScheduleDeleteParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
	}

	if !data.Region.IsNull() && !data.Region.IsUnknown() {
		params.Region = cloudflare.F(speed.ScheduleDeleteParamsRegion(data.Region.ValueString()))
	}

	_, err := r.client.Speed.Schedule.Delete(
		ctx,
		data.URL.ValueString(),
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.URL

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// parseScheduleImportID splits an import ID of the form
//
//	<zone_id>/<url>
//	<zone_id>/<url>/<region>
//
// into its components. Both url and region may be percent-encoded
// (e.g. example.com%2Fpath); they are decoded before being returned.
// region is empty when the two-segment form is used.
func parseScheduleImportID(id string) (zoneID, rawURL, region string, err error) {
	firstSlash := strings.Index(id, "/")
	if firstSlash <= 0 {
		return "", "", "", fmt.Errorf("expected format <zone_id>/<url> or <zone_id>/<url>/<region>, got %q", id)
	}
	zoneID = id[:firstSlash]
	rest := id[firstSlash+1:]

	// Default to no region; override only when a region segment is present.
	rawURL = rest
	if lastSlash := strings.LastIndex(rest, "/"); lastSlash >= 0 {
		rawURL = rest[:lastSlash]
		region = rest[lastSlash+1:]
	}

	// URL-decode both segments: import IDs may contain percent-encoded
	// characters such as example.com%2Fpath.
	rawURL, err = url.PathUnescape(rawURL)
	if err != nil {
		return "", "", "", fmt.Errorf("url segment is not valid percent-encoding: %w", err)
	}
	if region != "" {
		region, err = url.PathUnescape(region)
		if err != nil {
			return "", "", "", fmt.Errorf("region segment is not valid percent-encoding: %w", err)
		}
	}

	if zoneID == "" || rawURL == "" {
		return "", "", "", fmt.Errorf("zone_id and url must be non-empty, got %q", id)
	}

	return zoneID, rawURL, region, nil
}

func (r *ObservatoryScheduledTestResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ObservatoryScheduledTestModel)

	path_zone_id, path_url, path_region, err := parseScheduleImportID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("invalid import ID", err.Error())
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.URL = types.StringValue(path_url)

	params := speed.ScheduleGetParams{
		ZoneID: cloudflare.F(path_zone_id),
	}
	if path_region != "" {
		data.Region = types.StringValue(path_region)
		params.Region = cloudflare.F(speed.ScheduleGetParamsRegion(path_region))
	}

	res := new(http.Response)
	env := ObservatoryScheduledTestResultEnvelope{*data}
	var apiErr error
	_, apiErr = r.client.Speed.Schedule.Get(
		ctx,
		path_url,
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if apiErr != nil {
		resp.Diagnostics.AddError("failed to make http request", apiErr.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	apiErr = apijson.Unmarshal(bytes, &env)
	if apiErr != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", apiErr.Error())
		return
	}
	data = &env.Result
	data.ID = data.URL

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ObservatoryScheduledTestResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
