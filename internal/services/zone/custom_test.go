package zone_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ---------------------------------------------------------------------------
// Resource helpers
// ---------------------------------------------------------------------------

// makeZoneState builds a tfsdk.State whose ID field is set to id. All other
// fields are null so that State.Get succeeds without real API data.
func makeZoneState(t *testing.T, ctx context.Context, id types.String) tfsdk.State {
	t.Helper()

	model := zone.ZoneModel{
		ID:                  id,
		Name:                types.StringValue("example.com"),
		Account:             &zone.ZoneAccountModel{ID: types.StringValue("account-123")},
		Paused:              types.BoolValue(false),
		Type:                types.StringValue("full"),
		VanityNameServers:   customfield.NullList[types.String](ctx),
		ActivatedOn:         timetypes.NewRFC3339Null(),
		CNAMESuffix:         types.StringNull(),
		CreatedOn:           timetypes.NewRFC3339Null(),
		DevelopmentMode:     types.Float64Null(),
		ModifiedOn:          timetypes.NewRFC3339Null(),
		OriginalDnshost:     types.StringNull(),
		OriginalRegistrar:   types.StringNull(),
		Status:              types.StringNull(),
		VerificationKey:     types.StringNull(),
		NameServers:         customfield.NullList[types.String](ctx),
		OriginalNameServers: customfield.NullList[types.String](ctx),
		Permissions:         customfield.NullList[types.String](ctx),
		Meta:                customfield.NullObject[zone.ZoneMetaModel](ctx),
		Owner:               customfield.NullObject[zone.ZoneOwnerModel](ctx),
		Plan:                customfield.NullObject[zone.ZonePlanModel](ctx),
		Tenant:              customfield.NullObject[zone.ZoneTenantModel](ctx),
		TenantUnit:          customfield.NullObject[zone.ZoneTenantUnitModel](ctx),
	}

	schema := zone.ResourceSchema(ctx)
	state := tfsdk.State{Schema: schema}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("makeZoneState: failed to set state: %v", diags)
	}
	return state
}

// newZoneResource returns a ZoneResource whose cloudflare client is pointed at
// baseURL. Pass an empty string to leave the client nil (safe when the guard
// under test fires before any API call is attempted).
func newZoneResource(t *testing.T, baseURL string) *zone.ZoneResource {
	t.Helper()
	r := zone.NewResource().(*zone.ZoneResource)
	if baseURL == "" {
		return r
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
	)
	req := resource.ConfigureRequest{ProviderData: client}
	resp := &resource.ConfigureResponse{}
	r.Configure(context.Background(), req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("newZoneResource: Configure failed: %v", resp.Diagnostics)
	}
	return r
}

// ---------------------------------------------------------------------------
// Data source helpers
// ---------------------------------------------------------------------------

// makeZoneDataSourceConfig builds a tfsdk.Config whose ZoneID field is set to
// zoneID. All other fields are null and Filter is nil.
func makeZoneDataSourceConfig(t *testing.T, ctx context.Context, zoneID types.String) tfsdk.Config {
	t.Helper()

	model := zone.ZoneDataSourceModel{
		ID:                  types.StringNull(),
		ZoneID:              zoneID,
		ActivatedOn:         timetypes.NewRFC3339Null(),
		CNAMESuffix:         types.StringNull(),
		CreatedOn:           timetypes.NewRFC3339Null(),
		DevelopmentMode:     types.Float64Null(),
		ModifiedOn:          timetypes.NewRFC3339Null(),
		Name:                types.StringNull(),
		OriginalDnshost:     types.StringNull(),
		OriginalRegistrar:   types.StringNull(),
		Paused:              types.BoolNull(),
		Status:              types.StringNull(),
		Type:                types.StringNull(),
		VerificationKey:     types.StringNull(),
		NameServers:         customfield.NullList[types.String](ctx),
		OriginalNameServers: customfield.NullList[types.String](ctx),
		Permissions:         customfield.NullList[types.String](ctx),
		VanityNameServers:   customfield.NullList[types.String](ctx),
		Account:             customfield.NullObject[zone.ZoneAccountDataSourceModel](ctx),
		Meta:                customfield.NullObject[zone.ZoneMetaDataSourceModel](ctx),
		Owner:               customfield.NullObject[zone.ZoneOwnerDataSourceModel](ctx),
		Plan:                customfield.NullObject[zone.ZonePlanDataSourceModel](ctx),
		Tenant:              customfield.NullObject[zone.ZoneTenantDataSourceModel](ctx),
		TenantUnit:          customfield.NullObject[zone.ZoneTenantUnitDataSourceModel](ctx),
		Filter:              nil,
	}

	// tfsdk.Config has no Set method; use tfsdk.State (which does) with the
	// data source schema as a proxy to obtain the raw tftypes.Value.
	dsSchema := zone.DataSourceSchema(ctx)
	state := tfsdk.State{Schema: dsSchema}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("makeZoneDataSourceConfig: failed to set state: %v", diags)
	}
	return tfsdk.Config{Schema: dsSchema, Raw: state.Raw}
}

// newZoneDataSource returns a ZoneDataSource whose cloudflare client is pointed
// at baseURL.
func newZoneDataSource(t *testing.T, baseURL string) *zone.ZoneDataSource {
	t.Helper()
	d := zone.NewZoneDataSource().(*zone.ZoneDataSource)
	if baseURL == "" {
		return d
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
	)
	req := datasource.ConfigureRequest{ProviderData: client}
	resp := &datasource.ConfigureResponse{}
	d.Configure(context.Background(), req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("newZoneDataSource: Configure failed: %v", resp.Diagnostics)
	}
	return d
}

// ---------------------------------------------------------------------------
// ZoneResource — ModifyPlan
// ---------------------------------------------------------------------------

func TestZoneResource_ModifyPlan_EmptyID_FailsPlan(t *testing.T) {
	ctx := context.Background()
	r := newZoneResource(t, "") // ModifyPlan never touches the client

	req := resource.ModifyPlanRequest{State: makeZoneState(t, ctx, types.StringValue(""))}
	resp := &resource.ModifyPlanResponse{}

	r.ModifyPlan(ctx, req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for empty zone ID, got none")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// During a Create the prior state is null — the guard must not fire.
func TestZoneResource_ModifyPlan_NullPriorState_Passes(t *testing.T) {
	ctx := context.Background()
	r := newZoneResource(t, "")

	schema := zone.ResourceSchema(ctx)
	// tfsdk.State with zero-value Raw: IsNull() returns true (value == nil).
	req := resource.ModifyPlanRequest{State: tfsdk.State{Schema: schema}}
	resp := &resource.ModifyPlanResponse{}

	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("expected no error for null prior state (create path), got: %v", resp.Diagnostics)
	}
}

func TestZoneResource_ModifyPlan_ValidID_Passes(t *testing.T) {
	ctx := context.Background()
	r := newZoneResource(t, "")

	req := resource.ModifyPlanRequest{State: makeZoneState(t, ctx, types.StringValue("valid-zone-id-abc123"))}
	resp := &resource.ModifyPlanResponse{}

	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("expected no error for valid zone ID, got: %v", resp.Diagnostics)
	}
}

// ---------------------------------------------------------------------------
// ZoneResource — Read (prior state guard)
// ---------------------------------------------------------------------------

func TestZoneResource_Read_EmptyID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringValue(""))

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for empty zone ID, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with empty zone ID, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// ---------------------------------------------------------------------------
// ZoneResource — Update (prior state guard)
// ---------------------------------------------------------------------------

func TestZoneResource_Update_EmptyID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringValue(""))

	// Plan carries the same (empty) ID as the prior state.
	req := resource.UpdateRequest{
		Plan:  tfsdk.Plan{Schema: state.Schema, Raw: state.Raw},
		State: state,
	}
	resp := &resource.UpdateResponse{State: state}

	r.Update(ctx, req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for empty zone ID, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with empty zone ID, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// ---------------------------------------------------------------------------
// ZoneResource — Delete (prior state guard)
// ---------------------------------------------------------------------------

func TestZoneResource_Delete_EmptyID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringValue(""))

	resp := &resource.DeleteResponse{}
	r.Delete(ctx, resource.DeleteRequest{State: state}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for empty zone ID, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with empty zone ID, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// ---------------------------------------------------------------------------
// ZoneDataSource — Read (empty zone_id guard)
// ---------------------------------------------------------------------------

func TestZoneDataSource_Read_EmptyZoneID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	d := newZoneDataSource(t, ts.URL)
	cfg := makeZoneDataSourceConfig(t, ctx, types.StringValue(""))

	resp := &datasource.ReadResponse{State: tfsdk.State{Schema: cfg.Schema}}
	d.Read(ctx, datasource.ReadRequest{Config: cfg}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for empty zone_id, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with empty zone_id, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone configuration: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// ---------------------------------------------------------------------------
// ZoneResource — null ID prior-state tests
// ---------------------------------------------------------------------------

func TestZoneResource_ModifyPlan_NullID_FailsPlan(t *testing.T) {
	ctx := context.Background()
	r := newZoneResource(t, "")

	req := resource.ModifyPlanRequest{State: makeZoneState(t, ctx, types.StringNull())}
	resp := &resource.ModifyPlanResponse{}

	r.ModifyPlan(ctx, req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for null zone ID, got none")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

func TestZoneResource_Read_NullID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringNull())

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for null zone ID, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with null zone ID, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

func TestZoneResource_Update_NullID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringNull())

	req := resource.UpdateRequest{
		Plan:  tfsdk.Plan{Schema: state.Schema, Raw: state.Raw},
		State: state,
	}
	resp := &resource.UpdateResponse{State: state}

	r.Update(ctx, req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for null zone ID, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with null zone ID, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

func TestZoneResource_Delete_NullID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringNull())

	resp := &resource.DeleteResponse{}
	r.Delete(ctx, resource.DeleteRequest{State: state}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for null zone ID, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with null zone ID, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone state: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// ---------------------------------------------------------------------------
// ZoneDataSource — null zone_id prior-state test
// ---------------------------------------------------------------------------

func TestZoneDataSource_Read_NullZoneID_FailsWithoutAPICall(t *testing.T) {
	ctx := context.Background()

	requestMade := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	d := newZoneDataSource(t, ts.URL)
	cfg := makeZoneDataSourceConfig(t, ctx, types.StringNull())

	resp := &datasource.ReadResponse{State: tfsdk.State{Schema: cfg.Schema}}
	d.Read(ctx, datasource.ReadRequest{Config: cfg}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic for null zone_id, got none")
	}
	if requestMade {
		t.Error("no API request should have been made with null zone_id, but one was")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != "Invalid zone configuration: empty or null zone ID" {
		t.Errorf("unexpected summary: %q", got)
	}
}

// ---------------------------------------------------------------------------
// Post-API response guard — API returns empty zone ID
// ---------------------------------------------------------------------------

// zoneAPIResponseWithID returns a minimal Cloudflare API JSON envelope that
// will unmarshal into ZoneResultEnvelope with the given zone id value.
func zoneAPIResponseWithID(id string) string {
	return `{"result":{"id":"` + id + `","name":"example.com"},"success":true,"errors":[],"messages":[]}`
}

const wantAPIReturnedEmptyIDSummary = "Zone API returned empty or null zone ID"

func TestZoneResource_Read_APIReturnsEmptyID_FailsWithoutWritingState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(zoneAPIResponseWithID(""))) //nolint:errcheck
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	state := makeZoneState(t, ctx, types.StringValue("zone-abc123"))

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic when API returns empty zone ID, got none")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != wantAPIReturnedEmptyIDSummary {
		t.Errorf("unexpected summary: %q", got)
	}
}

func TestZoneResource_Create_APIReturnsEmptyID_FailsWithoutWritingState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(zoneAPIResponseWithID(""))) //nolint:errcheck
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)

	// Reuse makeZoneState as a convenient source of a fully-populated raw value
	// for the plan. The guard fires after the API call, so the plan content
	// does not need to carry a valid ID.
	planState := makeZoneState(t, ctx, types.StringValue(""))
	plan := tfsdk.Plan{Schema: planState.Schema, Raw: planState.Raw}

	schema := zone.ResourceSchema(ctx)
	resp := &resource.CreateResponse{State: tfsdk.State{Schema: schema}}
	r.Create(ctx, resource.CreateRequest{Plan: plan}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic when API returns empty zone ID on Create, got none")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != wantAPIReturnedEmptyIDSummary {
		t.Errorf("unexpected summary: %q", got)
	}
}

func TestZoneResource_Update_APIReturnsEmptyID_FailsWithoutWritingState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(zoneAPIResponseWithID(""))) //nolint:errcheck
	}))
	defer ts.Close()

	r := newZoneResource(t, ts.URL)
	// Prior state carries a valid ID so the prior-state guard passes.
	state := makeZoneState(t, ctx, types.StringValue("zone-abc123"))
	plan := tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}

	resp := &resource.UpdateResponse{State: state}
	r.Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected error diagnostic when API returns empty zone ID on Update, got none")
	}
	errs := resp.Diagnostics.Errors()
	if len(errs) == 0 {
		t.Fatal("expected at least one error")
	}
	if got := errs[0].Summary(); got != wantAPIReturnedEmptyIDSummary {
		t.Errorf("unexpected summary: %q", got)
	}
}
