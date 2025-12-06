package email_routing_dns_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// EmailRoutingSettingsWithSubdomains represents the full API response including subdomains
type EmailRoutingSettingsWithSubdomains struct {
	ID         string                       `json:"id"`
	Tag        string                       `json:"tag"`
	Name       string                       `json:"name"`
	Enabled    bool                         `json:"enabled"`
	Status     string                       `json:"status"`
	Subdomains []EmailRoutingDNSSubdomain   `json:"subdomains"`
}

type EmailRoutingDNSSubdomain struct {
	ID      string `json:"id"`
	Tag     string `json:"tag"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Status  string `json:"status"`
}

type EmailRoutingAPIResponse struct {
	Result  EmailRoutingSettingsWithSubdomains `json:"result"`
	Success bool                               `json:"success"`
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_dns", &resource.Sweeper{
		Name: "cloudflare_email_routing_dns",
		F: func(region string) error {
			ctx := context.Background()
			client := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if zoneID == "" {
				tflog.Info(ctx, "Skipping email routing DNS sweep: CLOUDFLARE_ZONE_ID not set")
				return nil
			}

			// First, get the full email routing settings including subdomains via raw API call
			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing", zoneID),
				nil,
			)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create request: %s", err))
				return fmt.Errorf("failed to create request: %w", err)
			}

			// Add authentication headers
			apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			apiKey := os.Getenv("CLOUDFLARE_API_KEY")
			apiEmail := os.Getenv("CLOUDFLARE_EMAIL")

			if apiToken != "" {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
			} else if apiKey != "" && apiEmail != "" {
				req.Header.Set("X-Auth-Key", apiKey)
				req.Header.Set("X-Auth-Email", apiEmail)
			} else {
				tflog.Error(ctx, "Missing authentication credentials")
				return fmt.Errorf("missing authentication credentials")
			}
			req.Header.Set("Content-Type", "application/json")

			// Execute the request
			httpClient := &http.Client{}
			resp, err := httpClient.Do(req)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to execute request: %s", err))
				return fmt.Errorf("failed to execute request: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to read response: %s", err))
				return fmt.Errorf("failed to read response: %w", err)
			}

			// Parse the response
			var apiResp EmailRoutingAPIResponse
			if err := json.Unmarshal(body, &apiResp); err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to parse response: %s", err))
				return fmt.Errorf("failed to parse response: %w", err)
			}

			settings := apiResp.Result
			tflog.Info(ctx, fmt.Sprintf("Found %d email routing DNS subdomains", len(settings.Subdomains)))

			// Clean up subdomains first by calling disable for each one
			deletedCount := 0
			skippedCount := 0
			for _, subdomain := range settings.Subdomains {
				if !subdomain.Enabled {
					skippedCount++
					continue
				}

				tflog.Info(ctx, fmt.Sprintf("Disabling subdomain: %s (zone: %s)", subdomain.Name, zoneID))

				// Call the disable endpoint with the subdomain name
				bodyJSON := fmt.Sprintf(`{"name":"%s"}`, subdomain.Name)
				disableReq, err := http.NewRequestWithContext(
					ctx,
					"POST",
					fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing/disable", zoneID),
					bytes.NewBufferString(bodyJSON),
				)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to create disable request for %s: %s", subdomain.Name, err))
					continue
				}

				// Add authentication headers
				if apiToken != "" {
					disableReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
				} else if apiKey != "" && apiEmail != "" {
					disableReq.Header.Set("X-Auth-Key", apiKey)
					disableReq.Header.Set("X-Auth-Email", apiEmail)
				}
				disableReq.Header.Set("Content-Type", "application/json")

				disableResp, err := httpClient.Do(disableReq)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to disable subdomain %s: %s", subdomain.Name, err))
					continue
				}
				disableResp.Body.Close()

				if disableResp.StatusCode >= 200 && disableResp.StatusCode < 300 {
					deletedCount++
					tflog.Info(ctx, fmt.Sprintf("Disabled subdomain: %s", subdomain.Name))
				} else {
					tflog.Error(ctx, fmt.Sprintf("Got status %d when disabling %s", disableResp.StatusCode, subdomain.Name))
				}
			}

			if deletedCount > 0 || skippedCount > 0 {
				tflog.Info(ctx, fmt.Sprintf("Disabled %d subdomains, skipped %d already disabled", deletedCount, skippedCount))
			}

			// Delete email routing DNS configuration (removes remaining DNS records)
			deletedRecords, err := client.EmailRouting.DNS.Delete(ctx, email_routing.DNSDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Info(ctx, fmt.Sprintf("Note: DNS delete returned error (might be expected): %v", err))
			} else if deletedRecords != nil && deletedRecords.Result != nil {
				tflog.Info(ctx, fmt.Sprintf("Deleted %d email routing DNS records", len(deletedRecords.Result)))
			}

			// Also disable main zone email routing if it's still enabled
			if settings.Enabled {
				tflog.Info(ctx, fmt.Sprintf("Disabling main zone email routing (zone: %s)", zoneID))
				_, err = client.EmailRouting.Disable(ctx, email_routing.EmailRoutingDisableParams{
					ZoneID: cloudflare.F(zoneID),
					Body:   map[string]interface{}{"name": settings.Name},
				})
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to disable main zone email routing: %s", err))
				} else {
					tflog.Info(ctx, "Disabled main zone email routing")
				}
			}

			return nil
		},
	})

	// Sweeper specifically for cleaning up subdomains (can be run separately if needed)
	resource.AddTestSweepers("cloudflare_email_routing_dns_subdomains", &resource.Sweeper{
		Name: "cloudflare_email_routing_dns_subdomains",
		F: func(region string) error {
			ctx := context.Background()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if zoneID == "" {
				tflog.Info(ctx, "Skipping email routing DNS subdomains sweep: CLOUDFLARE_ZONE_ID not set")
				return nil
			}

			// Make a raw HTTP GET request to get the full response with subdomains
			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing", zoneID),
				nil,
			)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create request: %s", err))
				return fmt.Errorf("failed to create request: %w", err)
			}

			// Add authentication headers
			apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			apiKey := os.Getenv("CLOUDFLARE_API_KEY")
			apiEmail := os.Getenv("CLOUDFLARE_EMAIL")

			if apiToken != "" {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
			} else if apiKey != "" && apiEmail != "" {
				req.Header.Set("X-Auth-Key", apiKey)
				req.Header.Set("X-Auth-Email", apiEmail)
			} else {
				tflog.Error(ctx, "Missing authentication credentials")
				return fmt.Errorf("missing authentication credentials")
			}
			req.Header.Set("Content-Type", "application/json")

			// Execute the request
			httpClient := &http.Client{}
			resp, err := httpClient.Do(req)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to execute request: %s", err))
				return fmt.Errorf("failed to execute request: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to read response: %s", err))
				return fmt.Errorf("failed to read response: %w", err)
			}

			// Parse the response
			var apiResp EmailRoutingAPIResponse
			if err := json.Unmarshal(body, &apiResp); err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to parse response: %s", err))
				return fmt.Errorf("failed to parse response: %w", err)
			}

			settings := apiResp.Result
			tflog.Info(ctx, fmt.Sprintf("Found %d email routing DNS subdomains to clean up", len(settings.Subdomains)))

			if len(settings.Subdomains) == 0 {
				tflog.Info(ctx, "No subdomains to clean up")
				return nil
			}

			// Iterate through each subdomain and disable it
			deletedCount := 0
			skippedCount := 0
			for _, subdomain := range settings.Subdomains {
				if !subdomain.Enabled {
					tflog.Info(ctx, fmt.Sprintf("Subdomain %s is already disabled, skipping", subdomain.Name))
					skippedCount++
					continue
				}

				tflog.Info(ctx, fmt.Sprintf("Disabling subdomain: %s (zone: %s)", subdomain.Name, zoneID))

				// Call the disable endpoint with the subdomain name
				bodyJSON := fmt.Sprintf(`{"name":"%s"}`, subdomain.Name)
				disableReq, err := http.NewRequestWithContext(
					ctx,
					"POST",
					fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing/disable", zoneID),
					bytes.NewBufferString(bodyJSON),
				)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to create disable request for %s: %s", subdomain.Name, err))
					continue
				}

				// Add authentication headers
				if apiToken != "" {
					disableReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
				} else if apiKey != "" && apiEmail != "" {
					disableReq.Header.Set("X-Auth-Key", apiKey)
					disableReq.Header.Set("X-Auth-Email", apiEmail)
				}
				disableReq.Header.Set("Content-Type", "application/json")

				disableResp, err := httpClient.Do(disableReq)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to disable subdomain %s: %s", subdomain.Name, err))
					continue
				}
				disableResp.Body.Close()

				if disableResp.StatusCode >= 200 && disableResp.StatusCode < 300 {
					deletedCount++
					tflog.Info(ctx, fmt.Sprintf("Successfully disabled subdomain: %s", subdomain.Name))
				} else {
					tflog.Error(ctx, fmt.Sprintf("Got status %d when disabling %s", disableResp.StatusCode, subdomain.Name))
				}
			}

			tflog.Info(ctx, fmt.Sprintf("Disabled %d email routing DNS subdomains, skipped %d already disabled", deletedCount, skippedCount))
			return nil
		},
	})
}

func testEmailRoutingDNSConfig(resourceID, zoneID string, subDomain string) string {
	return acctest.LoadTestCase("emailroutingdnsconfig.tf", resourceID, zoneID, subDomain)
}

func TestAccTestEmailRoutingDNS(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_dns." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	subDomain := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingDNSConfig(rnd, zoneID, subDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}
