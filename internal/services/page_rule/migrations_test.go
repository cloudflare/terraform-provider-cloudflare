package page_rule_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCloudflarePageRule_Basic tests the most fundamental
// page rule migration scenario with minimal configuration. This test ensures that:
// 1. The actions block is converted to an attribute (actions {} → actions = {})
// 2. Boolean string values are properly converted ("on" → true, "off" → false)
// 3. State transformation handles schema_version reset to 0
// 4. Resources remain functional after migration without requiring manual intervention
func TestMigrateCloudflarePageRule_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, target, rnd string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflarePageRuleMigrationConfigV4Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zone := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_page_rule." + rnd
			target := fmt.Sprintf("%s.%s/*", rnd, zone)
			testConfig := tc.configFn(zoneID, target, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create page rule with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
							// Note: priority is auto-managed by Cloudflare API, so we don't validate it
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.NotNull()),
					}),
					{
						// Step 3: Apply the migrated configuration
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
							// Check that actions contains ssl
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ssl": knownvalue.StringExact("flexible"),
							})),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflarePageRule_CacheDeceptionArmor tests migration of
// cache deception armor settings, which is commonly used for security.
func TestMigrateCloudflarePageRule_CacheDeceptionArmor(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, target, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflarePageRuleMigrationConfigV4CacheDeceptionArmor,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zone := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_page_rule." + rnd
			target := fmt.Sprintf("*.%s.%s/*", rnd, zone)
			testConfig := tc.configFn(zoneID, target, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							// Note: priority is auto-managed by Cloudflare API, so we don't validate it
						},
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"cache_deception_armor": knownvalue.StringExact("on"),
						})),
					}),
					{
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"cache_deception_armor": knownvalue.StringExact("on"),
							})),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflarePageRule_CacheKeyFields tests migration of
// cache key fields configuration, including the ignore → exclude/include transformation.
func TestMigrateCloudflarePageRule_CacheKeyFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, target, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflarePageRuleMigrationConfigV4CacheKeyFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zone := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_page_rule." + rnd
			target := fmt.Sprintf("%s.%s/api/*", rnd, zone)
			testConfig := tc.configFn(zoneID, target, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
						},
					},
					{
						PreConfig: func() {
							// Write out config
							acctest.WriteOutConfig(t, testConfig, tmpDir)

							// Run migration
							acctest.RunMigrationCommand(t, testConfig, tmpDir)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"cache_level": knownvalue.StringExact("aggressive"),
							})),
						},
					},
					{
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"cache_level": knownvalue.StringExact("aggressive"),
								"cache_key_fields": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"user": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"device_type": knownvalue.Bool(false),
										"geo":         knownvalue.Bool(false),
										"lang":        knownvalue.Bool(true),
									}),
									"query_string": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"exclude": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("api_token"),
										}),
									}),
									"header": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"include": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("x-custom-header"),
											knownvalue.StringExact("x-another-header"),
										}),
									}),
									"cookie": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"include": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("session_id"),
										}),
									}),
									"host": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"resolved": knownvalue.Bool(false),
									}),
								}),
							})),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflarePageRule_ForwardingURL tests migration of
// forwarding URL configuration (redirect rules).
func TestMigrateCloudflarePageRule_ForwardingURL(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, target, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflarePageRuleMigrationConfigV4ForwardingURL,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zone := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_page_rule." + rnd
			target := fmt.Sprintf("%s.%s/old-path/*", rnd, zone)
			testConfig := tc.configFn(zoneID, target, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
						},
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"forwarding_url": knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"status_code": knownvalue.Int64Exact(301),
								"url":         knownvalue.StringExact("https://example.com/new-path/$1"),
							}),
						})),
					}),
					{
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"forwarding_url": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"status_code": knownvalue.Int64Exact(301),
									"url":         knownvalue.StringExact("https://example.com/new-path/$1"),
								}),
							})),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflarePageRule_CompleteSettings tests migration with
// a comprehensive set of page rule settings including cache, security, and performance options.
func TestMigrateCloudflarePageRule_CompleteSettings(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, target, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflarePageRuleMigrationConfigV4Complete,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zone := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_page_rule." + rnd
			target := fmt.Sprintf("%s.%s/assets/*", rnd, zone)
			testConfig := tc.configFn(zoneID, target, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config:                 testConfig,
						ImportStateConfigExact: true,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("priority"), knownvalue.NotNull()),
						},
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"cache_level": knownvalue.StringExact("aggressive"),
						})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"browser_cache_ttl": knownvalue.Float64Exact(3600),
						})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"edge_cache_ttl": knownvalue.Float64Exact(7200),
						})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"security_level": knownvalue.StringExact("high"),
						})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"ssl": knownvalue.StringExact("flexible"),
						})),
					}),
					{
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target"), knownvalue.StringExact(target)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"cache_level": knownvalue.StringExact("aggressive"),
							})),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"browser_cache_ttl": knownvalue.Float64Exact(3600),
							})),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"edge_cache_ttl": knownvalue.Float64Exact(7200),
							})),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"security_level": knownvalue.StringExact("high"),
							})),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("actions"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ssl": knownvalue.StringExact("flexible"),
							})),
						},
					},
				},
			})
		})
	}
}

// V4 Configuration Functions

func testAccCloudflarePageRuleMigrationConfigV4Basic(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  priority = 1
  status   = "active"

  actions {
    ssl = "flexible"
  }
}`, zoneID, target, rnd)
}

func testAccCloudflarePageRuleMigrationConfigV4CacheDeceptionArmor(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  priority = 1
  status   = "active"

  actions {
    cache_deception_armor = "on"
  }
}`, zoneID, target, rnd)
}

func testAccCloudflarePageRuleMigrationConfigV4CacheKeyFields(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  priority = 1
  status   = "active"

  actions {
    cache_level = "aggressive"

    cache_key_fields {
	  user {
        lang = true
      }
      query_string {
        exclude = ["api_token"]
      }
      header {
        include = ["x-custom-header", "x-another-header"]
      }
      cookie {
        include = ["session_id"]
      }
	  host {}
    }
  }
}`, zoneID, target, rnd)
}

func testAccCloudflarePageRuleMigrationConfigV4ForwardingURL(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  priority = 1
  status   = "active"

  actions {
    forwarding_url {
      status_code = 301
      url         = "https://example.com/new-path/$1"
    }
  }
}`, zoneID, target, rnd)
}

func testAccCloudflarePageRuleMigrationConfigV4Complete(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  priority = 1
  status   = "active"

  actions {
    cache_level       = "aggressive"
    browser_cache_ttl = 3600
    edge_cache_ttl    = 7200
    security_level    = "high"
    ssl              = "flexible"

    cache_ttl_by_status {
      codes = "200"
      ttl   = 86400
    }

    cache_ttl_by_status {
      codes = "404"
      ttl   = 300
    }

    minify {
      html = "on"
      css  = "on"
      js   = "off"
    }
  }
}`, zoneID, target, rnd)
}

// V5 Configuration Functions

func testAccCloudflarePageRuleMigrationConfigV5Basic(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  status   = "active"

  actions = {
    ssl = "flexible"
  }
}`, zoneID, target, rnd)
}

func testAccCloudflarePageRuleMigrationConfigV5CacheDeceptionArmor(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  status   = "active"

  actions = {
    cache_deception_armor = "on"
  }
}`, zoneID, target, rnd)
}

func testAccCloudflarePageRuleMigrationConfigV5ForwardingURL(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
  zone_id  = "%[1]s"
  target   = "%[2]s"
  status   = "active"

  actions = {
    forwarding_url = {
      status_code = 301
      url         = "https://example.com/new-path/$1"
    }
  }
}`, zoneID, target, rnd)
}
