package workers_cron_trigger_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_cron_trigger"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestWorkersCronTriggerReadRefreshesSchedules(t *testing.T) {
	for _, test := range []struct {
		name  string
		crons []string
	}{
		{"changed", []string{"0 12 * * *"}},
		{"added", []string{"0 0 * * *", "0 12 * * *"}},
		{"removed", []string{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			r, closeServer := cronResource(t, http.StatusOK, test.crons)
			defer closeServer()
			schedules := []*workers_cron_trigger.WorkersCronTriggerSchedulesModel{{Cron: types.StringValue("0 0 * * *")}}
			model := workers_cron_trigger.WorkersCronTriggerModel{
				ID: types.StringValue("test-worker"), AccountID: types.StringValue("account-123"), ScriptName: types.StringValue("test-worker"), Schedules: &schedules,
			}
			state := tfsdk.State{Schema: workers_cron_trigger.ResourceSchema(ctx)}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatal(diags)
			}
			resp := resource.ReadResponse{State: state}
			r.Read(ctx, resource.ReadRequest{State: state}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatal(resp.Diagnostics)
			}
			assertCronState(t, resp.State, test.crons)
		})
	}
}

func TestWorkersCronTriggerImportReadsSchedules(t *testing.T) {
	r, closeServer := cronResource(t, http.StatusOK, []string{"0 12 * * *"})
	defer closeServer()
	resp := resource.ImportStateResponse{State: tfsdk.State{Schema: workers_cron_trigger.ResourceSchema(context.Background())}}
	r.ImportState(context.Background(), resource.ImportStateRequest{ID: "account-123/test-worker"}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	assertCronState(t, resp.State, []string{"0 12 * * *"})
}

func TestWorkersCronTriggerReadNotFound(t *testing.T) {
	for _, status := range []int{http.StatusNotFound, http.StatusForbidden} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			ctx := context.Background()
			r, closeServer := cronResource(t, status, nil)
			defer closeServer()
			schedules := []*workers_cron_trigger.WorkersCronTriggerSchedulesModel{}
			model := workers_cron_trigger.WorkersCronTriggerModel{ID: types.StringValue("test-worker"), AccountID: types.StringValue("account-123"), ScriptName: types.StringValue("test-worker"), Schedules: &schedules}
			state := tfsdk.State{Schema: workers_cron_trigger.ResourceSchema(ctx)}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatal(diags)
			}
			resp := resource.ReadResponse{State: state}
			r.Read(ctx, resource.ReadRequest{State: state}, &resp)
			if status == http.StatusNotFound {
				if resp.Diagnostics.HasError() || !resp.State.Raw.IsNull() {
					t.Fatalf("expected removal from state: %v", resp.Diagnostics)
				}
			} else if !resp.Diagnostics.HasError() || resp.State.Raw.IsNull() {
				t.Fatal("expected API error without removing state")
			}
		})
	}
}

func cronResource(t *testing.T, status int, crons []string) (*workers_cron_trigger.WorkersCronTriggerResource, func()) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-123/workers/scripts/test-worker/schedules" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if status != http.StatusOK {
			fmt.Fprint(w, `{"success":false,"errors":[{"code":10000,"message":"request failed"}],"result":null}`)
			return
		}
		fmt.Fprint(w, `{"success":true,"errors":[],"messages":[],"result":{"schedules":[`)
		for i, cron := range crons {
			if i > 0 {
				fmt.Fprint(w, ",")
			}
			fmt.Fprintf(w, `{"cron":%q,"created_on":"2026-11-01T00:00:00Z","modified_on":"2026-11-02T00:00:00Z"}`, cron)
		}
		fmt.Fprint(w, `]}}`)
	}))
	r := workers_cron_trigger.NewResource().(*workers_cron_trigger.WorkersCronTriggerResource)
	client := cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"), option.WithMaxRetries(0))
	resp := resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: client}, &resp)
	if resp.Diagnostics.HasError() {
		server.Close()
		t.Fatal(resp.Diagnostics)
	}
	return r, server.Close
}

func assertCronState(t *testing.T, state tfsdk.State, want []string) {
	t.Helper()
	var got workers_cron_trigger.WorkersCronTriggerModel
	if diags := state.Get(context.Background(), &got); diags.HasError() {
		t.Fatal(diags)
	}
	if got.AccountID.ValueString() != "account-123" || got.ScriptName.ValueString() != "test-worker" || got.ID.ValueString() != "test-worker" {
		t.Fatalf("resource identity changed: %+v", got)
	}
	if got.Schedules == nil {
		t.Fatal("schedules is null")
	}
	crons := []string{}
	for _, schedule := range *got.Schedules {
		crons = append(crons, schedule.Cron.ValueString())
	}
	if !reflect.DeepEqual(crons, want) {
		t.Fatalf("schedules = %v, want %v", crons, want)
	}
	for _, schedule := range *got.Schedules {
		if schedule.CreatedOn.ValueString() != "2026-11-01T00:00:00Z" || schedule.ModifiedOn.ValueString() != "2026-11-02T00:00:00Z" {
			t.Fatalf("timestamps not refreshed: %+v", schedule)
		}
	}
}
