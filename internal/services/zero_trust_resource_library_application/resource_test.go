package zero_trust_resource_library_application_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
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
						plancheck.ExpectNonEmptyPlan(),
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("category_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostnames"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("basic.example.com")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ip_subnets"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("192.0.2.0/24")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port_protocols"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("tcp/443")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_domains"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("support-basic.example.com")})),
				},
			},
			{
				Config: testAccCloudflareZeroTrustResourceLibraryApplicationConfig(rnd, accountID, "updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostnames"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("updated.example.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ip_subnets"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("198.51.100.0/24")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("port_protocols"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("udp/53")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_domains"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("support-updated.example.com")})),
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("category_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostnames"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("updated.example.com")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ip_subnets"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("198.51.100.0/24")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port_protocols"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("udp/53")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_domains"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("support-updated.example.com")})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"version"},
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustResourceLibraryApplicationDestroy(s *terraform.State) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := &http.Client{Timeout: 10 * time.Second}
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
		err := retry.RetryContext(ctx, 30*time.Second, func() *retry.RetryError {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return retry.NonRetryableError(err)
			}
			if token := os.Getenv(consts.APITokenEnvVarKey); token != "" {
				req.Header.Set("Authorization", "Bearer "+token)
			} else if serviceKey := os.Getenv(consts.APIUserServiceKeyEnvVarKey); serviceKey != "" {
				req.Header.Set("X-Auth-User-Service-Key", serviceKey)
			} else {
				req.Header.Set("X-Auth-Key", os.Getenv(consts.APIKeyEnvVarKey))
				req.Header.Set("X-Auth-Email", os.Getenv(consts.EmailEnvVarKey))
			}

			resp, err := client.Do(req)
			if err != nil {
				return retry.RetryableError(err)
			}
			defer resp.Body.Close()
			_, _ = io.Copy(io.Discard, resp.Body)

			switch resp.StatusCode {
			case http.StatusNotFound:
				return nil
			case http.StatusOK:
				return retry.RetryableError(fmt.Errorf("resource library application %s still exists", rs.Primary.ID))
			default:
				if resp.StatusCode == http.StatusRequestTimeout || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= http.StatusInternalServerError {
					return retry.RetryableError(fmt.Errorf("temporary status %d while checking Resource Library application %s", resp.StatusCode, rs.Primary.ID))
				}
				return retry.NonRetryableError(fmt.Errorf("expected deleted Resource Library application %s to return 404, got %d", rs.Primary.ID, resp.StatusCode))
			}
		})
		if err != nil {
			return fmt.Errorf("error checking if Resource Library application %s was destroyed: %w", rs.Primary.ID, err)
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
