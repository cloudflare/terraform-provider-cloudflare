package api_shield_operation_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cfv3 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_api_shield_operation", &resource.Sweeper{
		Name: "cloudflare_api_shield_operation",
		F:    testSweepCloudflareAPIShieldOperations,
	})
}

func testSweepCloudflareAPIShieldOperations(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping API Shield operations sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	operations, err := client.APIGateway.Operations.List(ctx, api_gateway.OperationListParams{
		ZoneID: cfv3.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch API Shield operations: %s", err))
		return fmt.Errorf("failed to fetch API Shield operations: %w", err)
	}

	if len(operations.Result) == 0 {
		tflog.Info(ctx, "No API Shield operations to sweep")
		return nil
	}

	for _, operation := range operations.Result {
		// Use standard filtering helper on the endpoint field
		if !utils.ShouldSweepResource(operation.Endpoint) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting API Shield operation: %s (endpoint: %s, zone: %s)", operation.OperationID, operation.Endpoint, zoneID))
		_, err := client.APIGateway.Operations.Delete(ctx, operation.OperationID, api_gateway.OperationDeleteParams{
			ZoneID: cfv3.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete API Shield operation %s: %s", operation.OperationID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted API Shield operation: %s", operation.OperationID))
	}

	return nil
}

func TestAccCloudflareAPIShieldOperation_Create(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAPIShieldOperationDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: domain, Endpoint: "/example/path"}),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("host"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("endpoint"), knownvalue.StringExact("/example/path")),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("last_updated"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:        resourceID,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			},
		},
	})
}

func TestAccCloudflareAPIShieldOperation_RequiresReplace(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	var operationID1, operationID2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAPIShieldOperationDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: domain, Endpoint: "/example/path"}),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("host"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("endpoint"), knownvalue.StringExact("/example/path")),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("last_updated"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testCheckAPIShieldOperationID(resourceID, &operationID1),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "POST", Host: domain, Endpoint: "/example/path"}),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("method"), knownvalue.StringExact("POST")),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("host"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("endpoint"), knownvalue.StringExact("/example/path")),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("last_updated"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testCheckAPIShieldOperationID(resourceID, &operationID2),
					testCheckAPIShieldOperationRecreated(&operationID1, &operationID2),
				),
			},
			{
				ResourceName:        resourceID,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			},
		},
	})
}

func TestAccCloudflareAPIShieldOperation_AllMethods(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// Test all HTTP methods supported by the API
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			resourceID := "cloudflare_api_shield_operation." + rnd

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck_Credentials(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckAPIShieldOperationDelete,
				Steps: []resource.TestStep{
					{
						Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{
							Method:   method,
							Host:     domain,
							Endpoint: fmt.Sprintf("/api/%s/test", method),
						}),
						ConfigStateChecks: []statecheck.StateCheck{
							// Required attributes
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("method"), knownvalue.StringExact(method)),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("host"), knownvalue.StringExact(domain)),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("endpoint"), knownvalue.StringExact(fmt.Sprintf("/api/%s/test", method))),
							// Computed attributes
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("last_updated"), knownvalue.NotNull()),
						},
					},
					{
						ResourceName:        resourceID,
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
					},
				},
			})
		})
	}
}

func TestAccCloudflareAPIShieldOperation_DifferentEndpoints(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	testCases := []struct {
		name     string
		endpoint string
	}{
		{
			name:     "SimpleEndpoint",
			endpoint: "/api/users",
		},
		{
			name:     "EndpointWithSingleParameter",
			endpoint: "/api/users/{var1}",
		},
		{
			name:     "EndpointWithMultipleParameters",
			endpoint: "/api/users/{var1}/posts/{var2}",
		},
		{
			name:     "EndpointWithTrailingSlash",
			endpoint: "/api/resources/",
		},
		{
			name:     "RootEndpoint",
			endpoint: "/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resourceName := rnd + "_" + tc.name
			resourceID := "cloudflare_api_shield_operation." + resourceName

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck_Credentials(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckAPIShieldOperationDelete,
				Steps: []resource.TestStep{
					{
						Config: testAccCloudflareAPIShieldOperation(resourceName, zoneID, cloudflare.APIShieldBasicOperation{
							Method:   "GET",
							Host:     domain,
							Endpoint: tc.endpoint,
						}),
						ConfigStateChecks: []statecheck.StateCheck{
							// Required attributes
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("host"), knownvalue.StringExact(domain)),
							// Note: endpoint may be normalized by Cloudflare, so we just check it's not null
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("endpoint"), knownvalue.NotNull()),
							// Computed attributes
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceID, tfjsonpath.New("last_updated"), knownvalue.NotNull()),
						},
					},
					{
						ResourceName:        resourceID,
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
					},
				},
			})
		})
	}
}

func testAccCheckAPIShieldOperationDelete(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_api_shield_operation" {
			continue
		}

		_, err := client.APIGateway.Operations.Get(
			context.Background(),
			rs.Primary.Attributes["operation_id"],
			api_gateway.OperationGetParams{
				ZoneID: cfv3.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
			},
		)
		if err == nil {
			return fmt.Errorf("operation still exists")
		}

		var notFoundError *cfv3.Error
		if !errors.As(err, &notFoundError) {
			return fmt.Errorf("expected not found error but got: %w", err)
		}
	}

	return nil
}

func testCheckAPIShieldOperationID(resourceName string, operationID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		*operationID = rs.Primary.Attributes["operation_id"]
		if *operationID == "" {
			return fmt.Errorf("operation_id is empty")
		}

		return nil
	}
}

func testCheckAPIShieldOperationRecreated(operationID1, operationID2 *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *operationID1 == "" || *operationID2 == "" {
			return fmt.Errorf("operation IDs not captured")
		}

		if *operationID1 == *operationID2 {
			return fmt.Errorf("resource was not recreated: operation_id remained the same (%s)", *operationID1)
		}

		return nil
	}
}

func testAccCloudflareAPIShieldOperation(resourceName, zone string, op cloudflare.APIShieldBasicOperation) string {
	return acctest.LoadTestCase("apishieldoperation.tf", resourceName, zone, op.Method, op.Host, op.Endpoint)
}
