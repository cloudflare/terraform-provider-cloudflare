package zero_trust_device_posture_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

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
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func TestAccCloudflareDevicePostureRule_OsVersion(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("mac")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact("==")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_OsVersionExtra(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersionExtra(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("mac")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact("==")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("os_version_extra"), knownvalue.StringExact("(a)")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_LinuxOsDistro(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigLinuxDistro(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("linux")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("1.0.0")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("os_distro_name"), knownvalue.StringExact("ubuntu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("os_distro_revision"), knownvalue.StringExact("1.0.0")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_DomainJoined(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("domain_joined")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("domain"), knownvalue.StringExact("example.com")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_Firewall(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("firewall")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("firewall description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_DiskEncryption_RequireAll(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDiskEncryptionRequireAll(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("disk_encryption")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("mac")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("require_all"), knownvalue.Bool(true)),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_DiskEncryption_CheckDisks(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDiskEncryptionCheckDisks(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("disk_encryption")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("mac")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("require_all"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("check_disks").AtSliceIndex(0), knownvalue.StringExact("C")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("check_disks").AtSliceIndex(1), knownvalue.StringExact("D")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigosversion.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigOsVersionExtra(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigosversionextra.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigLinuxDistro(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfiglinuxdistro.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigdomainjoined.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryptionRequireAll(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigdiskencryptionrequireall.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryptionCheckDisks(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigdiskencryptioncheckdisks.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigfirewall.tf", rnd, accountID)
}

// Test for File posture rule type
func TestAccCloudflareDevicePostureRule_File(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigFile(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("file")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("File posture rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("1h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("path"), knownvalue.StringExact("C:\\Program Files\\Test\\test.exe")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("exists"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("sha256"), knownvalue.StringExact("abcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Test for Application posture rule type
func TestAccCloudflareDevicePostureRule_Application(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigApplication(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("application")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Application posture rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("mac")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("path"), knownvalue.StringExact("/Applications/Test.app")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("thumbprint"), knownvalue.StringExact("abcd1234567890abcdef1234567890abcdef1234")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Test for Client Certificate posture rule type
func TestAccCloudflareDevicePostureRule_ClientCertificate(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	// Skip this test as it requires a valid certificate ID which is not available in CI
	if os.Getenv("CLOUDFLARE_CERTIFICATE_ID") == "" {
		t.Skip("CLOUDFLARE_CERTIFICATE_ID not set, skipping client certificate test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigClientCertificate(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("client_certificate_v2")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Client certificate posture rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("certificate_id"), knownvalue.StringExact("12345678-1234-1234-1234-123456789abc")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("cn"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("check_private_key"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("extended_key_usage").AtSliceIndex(0), knownvalue.StringExact("clientAuth")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("extended_key_usage").AtSliceIndex(1), knownvalue.StringExact("emailProtection")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("locations").AtMapKey("trust_stores").AtSliceIndex(0), knownvalue.StringExact("system")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("locations").AtMapKey("trust_stores").AtSliceIndex(1), knownvalue.StringExact("user")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("subject_alternative_names").AtSliceIndex(0), knownvalue.StringExact("test.example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("subject_alternative_names").AtSliceIndex(1), knownvalue.StringExact("alt.example.com")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Test for SentinelOne posture rule type
func TestAccCloudflareDevicePostureRule_SentinelOne(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	// Skip this test as it requires a valid SentinelOne S2S integration connection ID
	if os.Getenv("CLOUDFLARE_SENTINELONE_CONNECTION_ID") == "" {
		t.Skip("CLOUDFLARE_SENTINELONE_CONNECTION_ID not set, skipping SentinelOne test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigSentinelOne(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("sentinelone_s2s")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("SentinelOne posture rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("10m")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("2h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operating_system"), knownvalue.StringExact("windows")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("active_threats"), knownvalue.Float64Exact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact("==")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("infected"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("is_active"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("network_status"), knownvalue.StringExact("connected")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operational_state"), knownvalue.StringExact("na")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Test for Tanium posture rule type
func TestAccCloudflareDevicePostureRule_Tanium(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	// Skip this test as it requires a valid Tanium integration connection ID
	if os.Getenv("CLOUDFLARE_TANIUM_CONNECTION_ID") == "" {
		t.Skip("CLOUDFLARE_TANIUM_CONNECTION_ID not set, skipping Tanium test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigTanium(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("tanium_s2s")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Tanium posture rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("linux")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("connection_id"), knownvalue.StringExact("12345678-1234-1234-1234-123456789abc")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("eid_last_seen"), knownvalue.StringExact("2023-01-01T00:00:00Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("risk_level"), knownvalue.StringExact("low")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("score_operator"), knownvalue.StringExact(">=")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("total_score"), knownvalue.Float64Exact(85.5)),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Test for Serial Number posture rule type
func TestAccCloudflareDevicePostureRule_SerialNumber(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigSerialNumber(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("serial_number")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Serial number posture rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("mac")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("id"), knownvalue.StringExact("ABCD1234567890")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Test for Update scenarios with ConfigPlanChecks
func TestAccCloudflareDevicePostureRule_Update(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.1")),
				},
			},
			// Update step with plan validation
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersionUpdated(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						// Validate only changed fields appear in plan
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated description")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("1h")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("48h")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("11.0.1")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact(">=")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("os_version_extra"), knownvalue.StringExact("(updated)")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedule"), knownvalue.StringExact("1h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expiration"), knownvalue.StringExact("48h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("11.0.1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact(">=")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("input").AtMapKey("os_version_extra"), knownvalue.StringExact("(updated)")),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Helper functions for new test configurations
func testAccCloudflareDevicePostureRuleConfigFile(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigfile.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigApplication(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigapplication.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigClientCertificate(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigclientcertificate.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigSentinelOne(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigsentinelone.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigTanium(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigtanium.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigSerialNumber(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigserialnumber.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigOsVersionUpdated(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigosversion_updated.tf", rnd, accountID)
}

func testAccCheckCloudflareDevicePostureRuleDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_device_posture_rule" {
			continue
		}

		_, err := client.DevicePostureRule(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Posture Rule still exists")
		}
	}

	return nil
}
