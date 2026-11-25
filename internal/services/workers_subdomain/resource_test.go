package workers_subdomain_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// The workers.dev subdomain is an account singleton that other tests rely on,
// so this test re-applies the current value and restores it after destroy.
func TestAccCloudflareWorkersSubdomain_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_subdomain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	var subdomain string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			current, err := acctest.SharedClient().Workers.Subdomains.Get(context.Background(), workers.SubdomainGetParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				t.Skipf("account has no workers.dev subdomain: %s", err)
			}
			subdomain = current.Subdomain
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			_, err := acctest.SharedClient().Workers.Subdomains.Update(context.Background(), workers.SubdomainUpdateParams{
				AccountID: cloudflare.F(accountID),
				Subdomain: cloudflare.F(subdomain),
			})
			return err
		},
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("workerssubdomain.tf", rnd, accountID, subdomain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("subdomain"), knownvalue.StringExact(subdomain)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
		},
	})
}
