// This file is intentionally NOT code-generated.

package stream_download_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_download"
)

// Every tfsdk tag on the model must be a schema attribute. Top-level
// percent_complete, status and url fields made every Read fail with a Value
// Conversion Error (APIX-1934), and the parity test did not catch it.
func TestStreamDownloadModelFitsState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	s := stream_download.ResourceSchema(ctx)
	state := tfsdk.State{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}
	if diags := state.Set(ctx, &stream_download.StreamDownloadModel{}); diags.HasError() {
		t.Fatalf("model does not fit the resource schema: %v", diags)
	}
}
