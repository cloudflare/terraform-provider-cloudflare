package managed_transforms_test

import (
	"context"
	"fmt"
	"regexp"

	"os"
	"strconv"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}


func init() {
	resource.AddTestSweepers("cloudflare_managed_headers", &resource.Sweeper{
		Name: "cloudflare_managed_headers",
		F:    testSweepCloudflareManagedTransforms,
	})
}

func testSweepCloudflareManagedTransforms(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	if client == nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client"))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	err := client.ManagedTransforms.Delete(
		ctx,
		managed_transforms.ManagedTransformDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		},
	)

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare managed transforms: %s", err))
	}

	return nil
}

func cleanup(t *testing.T) {
	err := testSweepCloudflareManagedTransforms("")

	if err != nil {
		t.Fatal("failed to cleanup resource for testing")
	}
}

func makeTransform(id string, enabled bool) map[string]string {
	return map[string]string{
		"id": id,
		"enabled": strconv.FormatBool(enabled),
	}
}

func TestAccCloudflareManagedHeaders(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransforms(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_visitor_location_headers", true),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransformsReorder(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_visitor_location_headers", true),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransformsRemovedHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransforms(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_visitor_location_headers", true),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCheckCloudflareManagedTransformsDisabledHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_visitor_location_headers", false),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransformsEnabledReorderHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_visitor_location_headers", true),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransformsDisabledReorderHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_visitor_location_headers", false),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransformsRemovedHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_request_headers.*",
						makeTransform("add_true_client_ip_headers", true),
					),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "managed_response_headers.*",
						makeTransform("add_security_headers", true),
					),
				),
			},
			// See "Note about state import checks" comment in the end of the file for why we don't test that here.
		},
	})
}

func TestAccCloudflareManagedHeaders_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsBasic(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCheckCloudflareManagedTransformsRemovedHeader(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						// Verify the updated managed_request_headers (should have only 1 item now)
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("managed_request_headers"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":      knownvalue.StringExact("add_true_client_ip_headers"),
									"enabled": knownvalue.Bool(true),
								}),
							}),
						),
						// Verify managed_response_headers remains unchanged
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("managed_response_headers"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":      knownvalue.StringExact("add_security_headers"),
									"enabled": knownvalue.Bool(true),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_Disabled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsDisabled(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
				},
			},
			// See "Note about state import checks" comment in the end of the file for why we don't test that here.
		},
	})
}

func TestAccCloudflareManagedHeaders_RequestOnly(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsRequestOnly(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_ResponseOnly(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsResponseOnly(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_VisitorLocationHeaders(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsVisitorLocation(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_LeakedCredentialsCheck(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsMixedRequestResponse(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_RemoveVisitorIP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsRemoveVisitorIP(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("remove_visitor_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_ResponseHeaderDisabled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsResponseHeaderDisabled(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
				},
			},
			// See "Note about state import checks" comment in the end of the file for why we don't test that here.
		},
	})
}

func TestAccCloudflareManagedHeaders_MultipleRequestHeaders(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsMultipleRequest(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_MixedEnabledDisabled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsMixedState(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			// See "Note about state import checks" comment in the end of the file for why we don't test that here.
		},
	})
}

func TestAccCloudflareManagedHeaders_UpdateTransforms(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsVisitorLocation(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
				},
			},
			{
				Config: testAccCheckCloudflareManagedTransformsResponseHeaderDisabled(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						// Both managed_request_headers and managed_response_headers should change
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("managed_request_headers"),
							knownvalue.ListSizeExact(0),
						),
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("managed_response_headers"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":      knownvalue.StringExact("add_security_headers"),
									"enabled": knownvalue.Bool(false),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_security_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
				},
			},
			// See "Note about state import checks" comment in the end of the file for why we don't test that here.
		},
	})
}

func TestAccCloudflareManagedHeaders_ConflictDetection(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsConflictTest(rnd, zoneID),
				ExpectError: regexp.MustCompile("404"),
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_ConflictingTransforms(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareManagedTransformsConflictingHeaders(rnd, zoneID),
				ExpectError: regexp.MustCompile("403|Forbidden|conflict"),
			},
		},
	})
}

func TestAccCloudflareManagedHeaders_NonConflictingTransforms(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsNonConflictingHeaders(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// This test will only run if the environment has Enterprise plan access
func TestAccCloudflareManagedHeaders_Enterprise(t *testing.T) {
	// Skip if not Enterprise environment
	//if os.Getenv("CLOUDFLARE_ENTERPRISE_TEST") == "" {
	//	t.Skip("Skipping Enterprise test - set CLOUDFLARE_ENTERPRISE_TEST=1 to run")
	//}

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanup(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareManagedTransformsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransformsEnterprise(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_true_client_ip_headers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareManagedTransforms(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransforms.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsReorder(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsreorder.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsRemovedHeader(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsremovedheader.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsDisabledHeader(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsdisabledheader.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsEnabledReorderHeader(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsenabledreorderheader.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsDisabledReorderHeader(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsdisabledreorderheader.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsDisabled(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsdisabled.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsRequestOnly(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsrequestonly.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsResponseOnly(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsresponseonly.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsBasic(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransforms.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsVisitorLocation(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsvisitorlocation.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsMixedRequestResponse(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformstlsauth.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsRemoveVisitorIP(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsremovevisitorip.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsResponseHeaderDisabled(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsremovexpoweredby.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsMultipleRequest(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsmultiplerequest.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsMixedState(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsmixedstate.tf", rnd, zoneID)
}

// Test for potential conflicting transforms - this uses transforms that should be safe to use together
func testAccCheckCloudflareManagedTransformsConflictTest(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsconflicttest.tf", rnd, zoneID)
}

// Test data for Enterprise-only transforms
func testAccCheckCloudflareManagedTransformsEnterprise(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsenterprise.tf", rnd, zoneID)
}

// Test data for conflicting headers
func testAccCheckCloudflareManagedTransformsConflictingHeaders(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsconflictingheaders.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsNonConflictingHeaders(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsnonconflictingheaders.tf", rnd, zoneID)
}

// Destroy verification function
func testAccCheckCloudflareManagedTransformsDestroy(s *terraform.State) error {

	client := cloudflare.NewClient(
		option.WithAPIKey(os.Getenv("CLOUDFLARE_API_KEY")),
		option.WithAPIEmail(os.Getenv("CLOUDFLARE_EMAIL")),
	)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_managed_transforms" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		managedHeaders, err := client.ManagedTransforms.List(context.Background(), managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.String(zoneID),
		})
		if err != nil {
			return fmt.Errorf("error listing managed headers: %w", err)
		}

		// Check if any headers are still enabled
		for _, h := range managedHeaders.ManagedRequestHeaders {
			if h.Enabled {
				return fmt.Errorf("managed request header %s is still enabled", h.ID)
			}
		}
		for _, h := range managedHeaders.ManagedResponseHeaders {
			if h.Enabled {
				return fmt.Errorf("managed response header %s is still enabled", h.ID)
			}
		}
	}

	return nil
}

// Note about state import checks:
//
// We can't test the state import when there are disabled transformations: those won't exist in
// the new state (and there's no way to change that in `ImportState()` because it doesn't have access
// to the previous state), so terraform would see a state diff from an import.
