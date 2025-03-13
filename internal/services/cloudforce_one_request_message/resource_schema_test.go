// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloudforce_one_request_message"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCloudforceOneRequestMessageModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*cloudforce_one_request_message.CloudforceOneRequestMessageModel)(nil)
  schema := cloudforce_one_request_message.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
