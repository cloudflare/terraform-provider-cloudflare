package logpush_dataset_field

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogpushDatasetFieldResultDataSourceEnvelope(t *testing.T) {
	ctx := context.Background()
	env := LogpushDatasetFieldResultDataSourceEnvelope{
		Result: customfield.NullMap[types.String](ctx),
	}

	err := apijson.UnmarshalComputed([]byte(`{"result":{"ClientIP":"The IP address of the client.","Datetime":"The date and time the event occurred."},"success":true}`), &env)
	if err != nil {
		t.Fatalf("UnmarshalComputed() error = %v", err)
	}

	fields, diags := env.Result.Value(ctx)
	if diags.HasError() {
		t.Fatalf("Result.Value() diagnostics = %v", diags.Errors())
	}
	if got := fields["ClientIP"].ValueString(); got != "The IP address of the client." {
		t.Fatalf("Result[ClientIP] = %q, want %q", got, "The IP address of the client.")
	}
	if got := fields["Datetime"].ValueString(); got != "The date and time the event occurred." {
		t.Fatalf("Result[Datetime] = %q, want %q", got, "The date and time the event occurred.")
	}
}
