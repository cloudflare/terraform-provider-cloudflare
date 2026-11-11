package log_explorer_dataset_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	testresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestLogExplorerDatasetTerraformFilterRemoval(t *testing.T) {
	var mu sync.Mutex
	exists, filter := false, ""
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		switch req.Method {
		case http.MethodPost:
			body := decodeRequestBody(t, req)
			filter, _ = body["filter"].(string)
			exists = true
		case http.MethodPut:
			body := decodeRequestBody(t, req)
			if value, present := body["filter"].(string); present {
				filter = value
			}
		case http.MethodDelete:
			exists = false
			fmt.Fprint(w, `{"success":true,"result":{},"errors":[],"messages":[]}`)
			return
		case http.MethodGet:
			if !exists {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, `{"success":false,"result":null,"errors":[{"code":404,"message":"not found"}],"messages":[]}`)
				return
			}
		}
		var responseFilter *string
		if filter != "" {
			responseFilter = &filter
		}
		writeDatasetResponse(t, w, datasetResponse{Enabled: true, DeletionProtection: false, Filter: responseFilter, ObjectType: "zone", ObjectID: testZoneID})
	}))
	defer server.Close()
	t.Setenv("CLOUDFLARE_BASE_URL", server.URL)
	t.Setenv("CLOUDFLARE_API_TOKEN", strings.Repeat("x", 40))
	t.Setenv("CLOUDFLARE_API_KEY", "")
	t.Setenv("CLOUDFLARE_EMAIL", "")
	config := `resource "cloudflare_log_explorer_dataset" "test" {
  zone_id = "zone-123"
  dataset = "dns_logs"
  deletion_protection = false
  %s
}`
	testresource.UnitTest(t, testresource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []testresource.TestStep{
			{Config: fmt.Sprintf(config, ""), Check: testresource.TestCheckResourceAttr("cloudflare_log_explorer_dataset.test", "filter", "")},
			{Config: fmt.Sprintf(config, ""), PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: fmt.Sprintf(config, `filter = "configured"`), Check: testresource.TestCheckResourceAttr("cloudflare_log_explorer_dataset.test", "filter", "configured")},
			{Config: fmt.Sprintf(config, ""), Check: testresource.TestCheckResourceAttr("cloudflare_log_explorer_dataset.test", "filter", "")},
			{Config: fmt.Sprintf(config, ""), PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: fmt.Sprintf(config, `filter = ""`), Check: testresource.TestCheckResourceAttr("cloudflare_log_explorer_dataset.test", "filter", "")},
		},
	})
}
