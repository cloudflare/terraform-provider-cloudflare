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
			client := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			ctx := context.Background()

			// First, get the full email routing settings including subdomains via raw API call
			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing", zoneID),
				nil,
			)
			if err != nil {
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
				return fmt.Errorf("missing authentication credentials")
			}
			req.Header.Set("Content-Type", "application/json")

			// Execute the request
			httpClient := &http.Client{}
			resp, err := httpClient.Do(req)
			if err != nil {
				return fmt.Errorf("failed to execute request: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response: %w", err)
			}

			// Parse the response
			var apiResp EmailRoutingAPIResponse
			if err := json.Unmarshal(body, &apiResp); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			settings := apiResp.Result
			fmt.Printf("Found %d email routing DNS subdomains\n", len(settings.Subdomains))

			// Clean up subdomains first by calling disable for each one
			deletedCount := 0
			skippedCount := 0
			for _, subdomain := range settings.Subdomains {
				if !subdomain.Enabled {
					skippedCount++
					continue
				}

				fmt.Printf("Disabling subdomain: %s\n", subdomain.Name)

				// Call the disable endpoint with the subdomain name
				bodyJSON := fmt.Sprintf(`{"name":"%s"}`, subdomain.Name)
				disableReq, err := http.NewRequestWithContext(
					ctx,
					"POST",
					fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing/disable", zoneID),
					bytes.NewBufferString(bodyJSON),
				)
				if err != nil {
					fmt.Printf("Warning: failed to create disable request for %s: %v\n", subdomain.Name, err)
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
					fmt.Printf("Warning: failed to disable subdomain %s: %v\n", subdomain.Name, err)
					continue
				}
				disableResp.Body.Close()

				if disableResp.StatusCode >= 200 && disableResp.StatusCode < 300 {
					deletedCount++
				} else {
					fmt.Printf("Warning: got status %d when disabling %s\n", disableResp.StatusCode, subdomain.Name)
				}
			}

			if deletedCount > 0 || skippedCount > 0 {
				fmt.Printf("Disabled %d subdomains, skipped %d already disabled\n", deletedCount, skippedCount)
			}

			// Delete email routing DNS configuration (removes remaining DNS records)
			deletedRecords, err := client.EmailRouting.DNS.Delete(ctx, email_routing.DNSDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				fmt.Printf("Note: DNS delete returned error (might be expected): %v\n", err)
			} else if deletedRecords != nil && deletedRecords.Result != nil {
				fmt.Printf("Deleted %d email routing DNS records\n", len(deletedRecords.Result))
			}

			// Also disable main zone email routing if it's still enabled
			if settings.Enabled {
				_, err = client.EmailRouting.Disable(ctx, email_routing.EmailRoutingDisableParams{
					ZoneID: cloudflare.F(zoneID),
					Body:   map[string]interface{}{"name": settings.Name},
				})
				if err != nil {
					fmt.Printf("Warning: failed to disable main zone email routing: %v\n", err)
				} else {
					fmt.Println("Disabled main zone email routing")
				}
			}

			return nil
		},
	})

	// Sweeper specifically for cleaning up subdomains (can be run separately if needed)
	resource.AddTestSweepers("cloudflare_email_routing_dns_subdomains", &resource.Sweeper{
		Name: "cloudflare_email_routing_dns_subdomains",
		F: func(region string) error {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			ctx := context.Background()

			// Make a raw HTTP GET request to get the full response with subdomains
			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing", zoneID),
				nil,
			)
			if err != nil {
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
				return fmt.Errorf("missing authentication credentials")
			}
			req.Header.Set("Content-Type", "application/json")

			// Execute the request
			httpClient := &http.Client{}
			resp, err := httpClient.Do(req)
			if err != nil {
				return fmt.Errorf("failed to execute request: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response: %w", err)
			}

			// Parse the response
			var apiResp EmailRoutingAPIResponse
			if err := json.Unmarshal(body, &apiResp); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			settings := apiResp.Result
			fmt.Printf("Found %d email routing DNS subdomains to clean up\n", len(settings.Subdomains))

			if len(settings.Subdomains) == 0 {
				fmt.Println("No subdomains to clean up")
				return nil
			}

			// Iterate through each subdomain and disable it
			deletedCount := 0
			skippedCount := 0
			for _, subdomain := range settings.Subdomains {
				if !subdomain.Enabled {
					fmt.Printf("Subdomain %s is already disabled, skipping\n", subdomain.Name)
					skippedCount++
					continue
				}

				fmt.Printf("Disabling subdomain: %s\n", subdomain.Name)

				// Call the disable endpoint with the subdomain name
				bodyJSON := fmt.Sprintf(`{"name":"%s"}`, subdomain.Name)
				disableReq, err := http.NewRequestWithContext(
					ctx,
					"POST",
					fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/email/routing/disable", zoneID),
					bytes.NewBufferString(bodyJSON),
				)
				if err != nil {
					fmt.Printf("Warning: failed to create disable request for %s: %v\n", subdomain.Name, err)
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
					fmt.Printf("Warning: failed to disable subdomain %s: %v\n", subdomain.Name, err)
					continue
				}
				disableResp.Body.Close()

				if disableResp.StatusCode >= 200 && disableResp.StatusCode < 300 {
					deletedCount++
					fmt.Printf("Successfully disabled subdomain: %s\n", subdomain.Name)
				} else {
					fmt.Printf("Warning: got status %d when disabling %s\n", disableResp.StatusCode, subdomain.Name)
				}
			}

			fmt.Printf("Disabled %d email routing DNS subdomains, skipped %d already disabled\n", deletedCount, skippedCount)
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
