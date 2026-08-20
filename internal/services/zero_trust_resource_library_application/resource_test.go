package zero_trust_resource_library_application_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZeroTrustResourceLibraryApplication_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_zero_trust_resource_library_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustResourceLibraryApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustResourceLibraryApplicationConfig(rnd, accountID, "basic"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("human_id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("version"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemAttr(resourceName, "hostnames.*", "basic.example.com"),
					resource.TestCheckTypeSetElemAttr(resourceName, "ip_subnets.*", "192.0.2.0/24"),
					resource.TestCheckTypeSetElemAttr(resourceName, "port_protocols.*", "tcp/443"),
					resource.TestCheckTypeSetElemAttr(resourceName, "support_domains.*", "support-basic.example.com"),
				),
			},
			{
				Config: testAccCloudflareZeroTrustResourceLibraryApplicationConfig(rnd, accountID, "updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemAttr(resourceName, "hostnames.*", "updated.example.com"),
					resource.TestCheckTypeSetElemAttr(resourceName, "ip_subnets.*", "198.51.100.0/24"),
					resource.TestCheckTypeSetElemAttr(resourceName, "port_protocols.*", "udp/53"),
					resource.TestCheckTypeSetElemAttr(resourceName, "support_domains.*", "support-updated.example.com"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustResourceLibraryApplicationDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_resource_library_application" {
			continue
		}

		baseURL := os.Getenv(consts.BaseURLEnvVarKey)
		if baseURL == "" {
			baseURL = "https://api.cloudflare.com/client/v4"
		}
		url := fmt.Sprintf(
			"%s/accounts/%s/resource-library/applications/%s",
			strings.TrimRight(baseURL, "/"),
			rs.Primary.Attributes[consts.AccountIDSchemaKey],
			rs.Primary.ID,
		)
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		if token := os.Getenv(consts.APITokenEnvVarKey); token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		} else if serviceKey := os.Getenv(consts.APIUserServiceKeyEnvVarKey); serviceKey != "" {
			req.Header.Set("X-Auth-User-Service-Key", serviceKey)
		} else {
			req.Header.Set("X-Auth-Key", os.Getenv(consts.APIKeyEnvVarKey))
			req.Header.Set("X-Auth-Email", os.Getenv(consts.EmailEnvVarKey))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			return fmt.Errorf("expected deleted Resource Library application %s to return 404, got %d", rs.Primary.ID, resp.StatusCode)
		}
	}

	return nil
}

func testAccCloudflareZeroTrustResourceLibraryApplicationConfig(rnd, accountID, matcher string) string {
	subnet := "192.0.2.0/24"
	portProtocol := "tcp/443"
	if matcher == "updated" {
		subnet = "198.51.100.0/24"
		portProtocol = "udp/53"
	}

	return acctest.LoadTestCase("basic.tf", rnd, accountID, matcher, subnet, portProtocol)
}
