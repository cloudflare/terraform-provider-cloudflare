package workers_cron_trigger_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccWorkersCronTriggerRefreshAndImport(t *testing.T) {
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("set TF_ACC=1 to run against a live Cloudflare account")
	}
	accountID, token := os.Getenv("CLOUDFLARE_ACCOUNT_ID"), os.Getenv("CLOUDFLARE_API_TOKEN")
	if accountID == "" || token == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_API_TOKEN are required")
	}
	client := cloudflare.NewClient(option.WithAPIToken(token), option.WithBaseURL("https://api.cloudflare.com/client/v4/"))
	random := make([]byte, 8)
	if _, err := rand.Read(random); err != nil {
		t.Fatal(err)
	}
	name := "tf-acctest-cron-refresh-" + hex.EncodeToString(random)
	ctx := context.Background()
	_, err := client.Workers.Scripts.Schedules.Get(ctx, name, workers.ScriptScheduleGetParams{AccountID: cloudflare.F(accountID)})
	var apiError *cloudflare.Error
	if !errors.As(err, &apiError) || apiError.StatusCode != http.StatusNotFound {
		t.Fatalf("expected unused Worker name (404), got: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if _, err := client.Workers.Scripts.Delete(cleanupCtx, name, workers.ScriptDeleteParams{AccountID: cloudflare.F(accountID)}); err != nil {
			t.Errorf("delete test Worker %s: %v", name, err)
			return
		}
		_, err := client.Workers.Scripts.Schedules.Get(cleanupCtx, name, workers.ScriptScheduleGetParams{AccountID: cloudflare.F(accountID)})
		var apiError *cloudflare.Error
		if !errors.As(err, &apiError) || apiError.StatusCode != http.StatusNotFound {
			t.Errorf("test Worker deletion was not confirmed: %v", err)
			return
		}
		t.Logf("Deleted dedicated test Worker %s; confirmed API returns 404", name)
	})
	// The fixture has no bindings, routes or external effects; scheduled invocations do nothing.
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	metadata, err := writer.CreateFormField("metadata")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Fprintf(metadata, `{"main_module":"worker.js","compatibility_date":%q}`, time.Now().UTC().Format("2006-01-02"))
	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", `form-data; name="worker.js"; filename="worker.js"`)
	header.Set("Content-Type", "application/javascript+module")
	module, err := writer.CreatePart(header)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Fprint(module, `export default { async fetch() { return new Response("acceptance test"); }, async scheduled() {} };`)
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Workers.Scripts.Update(ctx, name, workers.ScriptUpdateParams{AccountID: cloudflare.F(accountID)}, option.WithRequestBody(writer.FormDataContentType(), body.Bytes())); err != nil {
		t.Fatal(err)
	}
	t.Logf("Created dedicated test Worker %s", name)

	config := fmt.Sprintf(`resource "cloudflare_workers_cron_trigger" "test" {
  account_id = %q
  script_name = %q
  schedules = [{ cron = "0 0 1 1 *" }]
 }`, accountID, name)
	const resourceName = "cloudflare_workers_cron_trigger.test"
	steps := []resource.TestStep{
		{Config: config, Check: resource.TestCheckResourceAttr(resourceName, "schedules.0.cron", "0 0 1 1 *")},
		{Config: config, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
	}
	for _, crons := range [][]string{{"0 12 1 1 *"}, {}} {
		crons := crons
		steps = append(steps,
			resource.TestStep{
				Config: config, PlanOnly: true, ExpectNonEmptyPlan: true,
				PreConfig: func() {
					body := []workers.ScriptScheduleUpdateParamsBody{}
					for _, cron := range crons {
						body = append(body, workers.ScriptScheduleUpdateParamsBody{Cron: cloudflare.F(cron)})
					}
					if _, err := client.Workers.Scripts.Schedules.Update(ctx, name, workers.ScriptScheduleUpdateParams{AccountID: cloudflare.F(accountID), Body: body}); err != nil {
						t.Fatal(err)
					}
					actual, err := client.Workers.Scripts.Schedules.Get(ctx, name, workers.ScriptScheduleGetParams{AccountID: cloudflare.F(accountID)})
					if err != nil {
						t.Fatal(err)
					}
					if len(actual.Schedules) != len(crons) || (len(crons) > 0 && actual.Schedules[0].Cron != crons[0]) {
						t.Fatalf("API did not confirm external schedule mutation: %+v", actual.Schedules)
					}
					t.Logf("Confirmed external schedule mutation via live API: %v", crons)
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate)}},
			},
			resource.TestStep{Config: config, Check: resource.TestCheckResourceAttr(resourceName, "schedules.0.cron", "0 0 1 1 *")},
			resource.TestStep{Config: config, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
		)
	}
	steps = append(steps,
		resource.TestStep{ResourceName: resourceName, ImportState: true, ImportStateId: accountID + "/" + name, ImportStateVerify: true},
		resource.TestStep{Config: strings.ReplaceAll(config, "0 0 1 1 *", "0 6 1 1 *"), Check: resource.TestCheckResourceAttr(resourceName, "schedules.0.cron", "0 6 1 1 *")},
	)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"cloudflare": providerserver.NewProtocol6WithError(&cronTestProvider{client: client})},
		Steps:                    steps,
	})
}
