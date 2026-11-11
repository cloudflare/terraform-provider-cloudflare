package logpush_job

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogpushJobDestinationConfNoRefresh(t *testing.T) {
	t.Parallel()

	responses := []struct {
		name string
		body string
	}{
		{"present", `{"result":{"id":42,"name":"refreshed","destination_conf":"api-value"}}`},
		{"absent", `{"result":{"id":42,"name":"refreshed"}}`},
		{"null", `{"result":{"id":42,"name":"refreshed","destination_conf":null}}`},
	}
	states := []struct {
		name  string
		value types.String
	}{
		{"refresh", types.StringValue("configured-value")},
		{"import", types.StringNull()},
	}

	for _, state := range states {
		for _, response := range responses {
			t.Run(state.name+"/"+response.name, func(t *testing.T) {
				t.Parallel()

				// Read seeds the envelope with prior state; import starts with no destination.
				env := LogpushJobResultEnvelope{Result: LogpushJobModel{
					DestinationConf: state.value,
					Name:            types.StringValue("previous-name"),
				}}
				if err := apijsoncustom.Unmarshal([]byte(response.body), &env); err != nil {
					t.Fatal(err)
				}
				if !env.Result.DestinationConf.Equal(state.value) {
					t.Fatal("refresh must ignore the response destination_conf and preserve prior state")
				}
				if env.Result.ID.ValueInt64() != 42 || env.Result.Name.ValueString() != "refreshed" {
					t.Fatal("other response fields must still refresh")
				}
			})
		}
	}
}

func TestLogpushJobDestinationConfCreateUpdate(t *testing.T) {
	t.Parallel()

	for _, operation := range []string{"create", "update"} {
		for _, response := range []string{
			`{"result":{"id":42,"destination_conf":"api-value"}}`,
			`{"result":{"id":42}}`,
			`{"result":{"id":42,"destination_conf":null}}`,
		} {
			t.Run(operation+"/"+response, func(t *testing.T) {
				t.Parallel()

				plan := LogpushJobModel{DestinationConf: types.StringValue("configured-value")}
				var body []byte
				var err error
				if operation == "create" {
					body, err = plan.MarshalJSON()
				} else {
					state := LogpushJobModel{DestinationConf: types.StringValue("previous-value")}
					body, err = plan.MarshalJSONForUpdate(state)
				}
				if err != nil {
					t.Fatal(err)
				}
				var request map[string]any
				if err := json.Unmarshal(body, &request); err != nil {
					t.Fatal(err)
				}
				if request["destination_conf"] != plan.DestinationConf.ValueString() {
					t.Fatal("no_refresh must not suppress destination_conf in requests")
				}

				// Both Create and Update decode computed fields into the planned model.
				env := LogpushJobResultEnvelope{Result: plan}
				if err := apijsoncustom.UnmarshalComputed([]byte(response), &env); err != nil {
					t.Fatal(err)
				}
				if !env.Result.DestinationConf.Equal(plan.DestinationConf) {
					t.Fatal("create/update response must preserve the planned destination_conf")
				}
				if env.Result.ID.ValueInt64() != 42 {
					t.Fatal("create/update must still decode computed response fields")
				}
			})
		}
	}
}
