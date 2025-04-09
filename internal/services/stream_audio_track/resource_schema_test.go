// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_audio_track_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_audio_track"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamAudioTrackModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*stream_audio_track.StreamAudioTrackModel)(nil)
  schema := stream_audio_track.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
