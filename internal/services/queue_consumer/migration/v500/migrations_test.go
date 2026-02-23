package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v5_http.tf
var v5HTTPConfig string

//go:embed testdata/v5_worker.tf
var v5WorkerConfig string

// TestMigrateQueueConsumer_HTTP tests migration of an HTTP pull queue consumer.
//
// cloudflare_queue_consumer is a v5-native resource; this test verifies that
// existing states at schema_version=0 (before explicit versioning) upgrade
// cleanly to schema_version=500 with no data loss.
func TestMigrateQueueConsumer_HTTP(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd string) string
	}{
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd string) string {
				return fmt.Sprintf(v5HTTPConfig, accountID, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			firstStep := resource.TestStep{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_queue_consumer."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_queue_consumer."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
					}),
				},
			})
		})
	}
}

// TestMigrateQueueConsumer_Worker tests migration of a worker-type queue consumer.
//
// cloudflare_queue_consumer is a v5-native resource; this test verifies that
// existing states at schema_version=0 (before explicit versioning) upgrade
// cleanly to schema_version=500 with no data loss.
func TestMigrateQueueConsumer_Worker(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd string) string
	}{
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd string) string {
				return fmt.Sprintf(v5WorkerConfig, accountID, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			firstStep := resource.TestStep{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_queue_consumer."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_queue_consumer."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					}),
				},
			})
		})
	}
}
