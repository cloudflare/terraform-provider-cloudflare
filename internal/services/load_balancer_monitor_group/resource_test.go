package load_balancer_monitor_group_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/load_balancers"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// Acceptance tests for cloudflare_load_balancer_monitor_group.

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_load_balancer_monitor_group", &resource.Sweeper{
		Name: "cloudflare_load_balancer_monitor_group",
		F:    testSweepCloudflareLoadBalancerMonitorGroups,
	})
}

func testSweepCloudflareLoadBalancerMonitorGroups(_ string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	client := acctest.SharedClient()
	iter := client.LoadBalancers.MonitorGroups.ListAutoPaging(ctx, load_balancers.MonitorGroupListParams{
		AccountID: cloudflare.F(accountID),
	})

	for iter.Next() {
		group := iter.Current()
		if !strings.HasPrefix(group.Description, "tf-acc") {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer Monitor Group ID: %s", group.ID))
		_, err := client.LoadBalancers.MonitorGroups.Delete(ctx, group.ID, load_balancers.MonitorGroupDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare Load Balancer Monitor Group %q: %s", group.ID, err))
		}
	}
	if err := iter.Err(); err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list Cloudflare Load Balancer Monitor Groups: %s", err))
	}
	return nil
}

// ---------------------------------------------------------------------------
// Schema validation tests
// ---------------------------------------------------------------------------

func TestAccCloudflareLoadBalancerMonitorGroup_MissingDescription(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccLBMonitorGroupMissingDescription(rnd, accountID),
				ExpectError: regexp.MustCompile(`(?s)The argument "description" is required`),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitorGroup_MissingMembers(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccLBMonitorGroupMissingMembers(rnd, accountID),
				ExpectError: regexp.MustCompile(`(?s)The argument "members" is required`),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Lifecycle tests
// ---------------------------------------------------------------------------

// TestAccCloudflareLoadBalancerMonitorGroup_Basic exercises the full lifecycle:
//
//  1. Create with 1 member — verifies state.
//  2. Import — verifies imported state matches.
func TestAccCloudflareLoadBalancerMonitorGroup_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor_group." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBMonitorGroupSingleMember(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("tf-acc basic "+rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members").AtSliceIndex(0),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"enabled":         knownvalue.Bool(true),
							"monitoring_only": knownvalue.Bool(false),
							"must_be_healthy": knownvalue.Bool(true),
						}),
					),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{
					"created_on", "modified_on",
					"members.0.created_at", "members.0.updated_at",
				},
			},
		},
	})
}

// TestAccCloudflareLoadBalancerMonitorGroup_AddRemoveMember verifies adding
// and removing members via in-place updates:
//
//  1. Create with 1 member.
//  2. Add a second member (1→2) — verifies the apply succeeds and state has
//     2 members. Uses ExpectNonEmptyPlan because the API returns members
//     ordered by created_at DESC, which may differ from config order and
//     cause a spurious refresh diff with ListNestedAttribute.
//  3. Remove a member (2→1) — verifies in-place update back to 1 member.
func TestAccCloudflareLoadBalancerMonitorGroup_AddRemoveMember(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor_group." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorGroupDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with 1 member.
			{
				Config: testAccLBMonitorGroupSingleMember(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
				},
			},
			// Step 2: Add a second member (1→2).
			{
				Config: testAccLBMonitorGroupBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "members.#", "2"),
				),
				ExpectNonEmptyPlan: true,
			},
			// Step 3: Remove a member (2→1).
			{
				Config: testAccLBMonitorGroupSingleMember(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members").AtSliceIndex(0),
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"enabled":         knownvalue.Bool(true),
							"monitoring_only": knownvalue.Bool(false),
							"must_be_healthy": knownvalue.Bool(true),
						}),
					),
				},
			},
		},
	})
}

// TestAccCloudflareLoadBalancerMonitorGroup_UpdateDescription creates a group,
// then updates only its description, verifying an in-place update that
// preserves the member.
func TestAccCloudflareLoadBalancerMonitorGroup_UpdateDescription(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor_group." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBMonitorGroupSingleMember(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("tf-acc basic "+rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				Config: testAccLBMonitorGroupUpdatedDescription(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("tf-acc updated "+rnd)),
					// Member unchanged.
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestAccCloudflareLoadBalancerMonitorGroup_UpdateMemberFlags creates a group
// with one member, toggles its boolean flags, then swaps the monitor_id to a
// different monitor — verifying both are in-place updates.
func TestAccCloudflareLoadBalancerMonitorGroup_UpdateMemberFlags(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor_group." + rnd
	monitorBName := "cloudflare_load_balancer_monitor." + rnd + "_b"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorGroupDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with monitor_a, all flags true/false/true.
			{
				Config: testAccLBMonitorGroupSingleMember(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled":         knownvalue.Bool(true),
								"monitoring_only": knownvalue.Bool(false),
								"must_be_healthy": knownvalue.Bool(true),
							}),
						}),
					),
				},
			},
			// Step 2: Toggle all boolean flags.
			{
				Config: testAccLBMonitorGroupMemberFlagsToggled(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled":         knownvalue.Bool(false),
								"monitoring_only": knownvalue.Bool(true),
								"must_be_healthy": knownvalue.Bool(false),
							}),
						}),
					),
				},
			},
			// Step 3: Swap monitor_id from monitor_a to monitor_b.
			{
				Config: testAccLBMonitorGroupSwapMonitor(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
					// Confirm the member now references monitor_b. We
					// can't hardcode the ID (it's server-generated), so
					// use a cross-resource comparison.
					statecheck.CompareValuePairs(
						resourceName, tfjsonpath.New("members").AtSliceIndex(0).AtMapKey("monitor_id"),
						monitorBName, tfjsonpath.New("id"),
						compare.ValuesSame(),
					),
				},
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Error case tests
// ---------------------------------------------------------------------------

// TestAccCloudflareLoadBalancerMonitorGroup_InvalidMonitorID verifies that
// referencing a nonexistent monitor ID is rejected by the API.
func TestAccCloudflareLoadBalancerMonitorGroup_InvalidMonitorID(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccLBMonitorGroupInvalidMonitorID(rnd, accountID),
				ExpectError: regexp.MustCompile(`(?i)(not found|foreign key|dependency|invalid|does not exist)`),
			},
		},
	})
}

// TestAccCloudflareLoadBalancerMonitorGroup_DuplicateMonitorID verifies that
// two members referencing the same monitor_id are rejected.
func TestAccCloudflareLoadBalancerMonitorGroup_DuplicateMonitorID(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccLBMonitorGroupDuplicateMonitorID(rnd, accountID),
				ExpectError: regexp.MustCompile(`(?i)(duplicate|unique|already exists|conflict)`),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// CheckDestroy
// ---------------------------------------------------------------------------

func testAccCheckCloudflareLoadBalancerMonitorGroupDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer_monitor_group" {
			continue
		}

		accountID := rs.Primary.Attributes["account_id"]
		_, err := client.LoadBalancers.MonitorGroups.Get(
			context.Background(),
			rs.Primary.ID,
			load_balancers.MonitorGroupGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("load balancer monitor group %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// Test fixture helpers
// ---------------------------------------------------------------------------

func testAccLBMonitorGroupMissingDescription(rnd, accountID string) string {
	return acctest.LoadTestCase("missing_description.tf", rnd, accountID)
}

func testAccLBMonitorGroupMissingMembers(rnd, accountID string) string {
	return acctest.LoadTestCase("missing_members.tf", rnd, accountID)
}

func testAccLBMonitorGroupBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}

func testAccLBMonitorGroupSingleMember(rnd, accountID string) string {
	return acctest.LoadTestCase("single_member.tf", rnd, accountID)
}

func testAccLBMonitorGroupUpdatedDescription(rnd, accountID string) string {
	return acctest.LoadTestCase("updated_description.tf", rnd, accountID)
}

func testAccLBMonitorGroupMemberFlagsToggled(rnd, accountID string) string {
	return acctest.LoadTestCase("member_flags_toggled.tf", rnd, accountID)
}

func testAccLBMonitorGroupSwapMonitor(rnd, accountID string) string {
	return acctest.LoadTestCase("swap_monitor.tf", rnd, accountID)
}

func testAccLBMonitorGroupInvalidMonitorID(rnd, accountID string) string {
	return acctest.LoadTestCase("invalid_monitor_id.tf", rnd, accountID)
}

func testAccLBMonitorGroupDuplicateMonitorID(rnd, accountID string) string {
	return acctest.LoadTestCase("duplicate_monitor_id.tf", rnd, accountID)
}
