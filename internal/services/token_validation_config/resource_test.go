package token_validation_config_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_token_validation_config", &resource.Sweeper{
		Name: "cloudflare_token_validation_config",
		F:    testSweepCloudflareTokenValidationConfig,
	})
}

func testSweepCloudflareTokenValidationConfig(r string) error {
	ctx := context.Background()
	// Token Validation Config is API Shield JWT validation configuration.
	// These are managed as part of API Shield configuration.
	// No sweeping required.
	tflog.Info(ctx, "Token Validation Config doesn't require sweeping (API Shield configuration)")
	return nil
}

type JWK struct {
	Alg string  `json:"alg"`
	Kid string  `json:"kid"`
	Kty string  `json:"kty"`
	Crv *string `json:"crv,omitempty"`
	X   *string `json:"x,omitempty"`
	Y   *string `json:"y,omitempty"`
	E   *string `json:"e,omitempty"`
	N   *string `json:"n,omitempty"`
}
type JWKS struct {
	Keys []JWK `json:"keys"`
}

// Returns a new JWKS from this JWKS with only the keys matching the algorithms
// The keys are sorted by algorithm
func (jwks *JWKS) CloneFiltered(algs ...string) JWKS {
	var filteredKeys []JWK
	for _, jwk := range jwks.Keys {
		if slices.Contains(algs, jwk.Alg) {
			filteredKeys = append(filteredKeys, jwk)
		}
	}
	slices.SortStableFunc(filteredKeys, func(a JWK, b JWK) int {
		return strings.Compare(a.Alg, b.Alg)
	})
	return JWKS{Keys: filteredKeys}
}

func TestAccCloudflareTokenValidationConfig(t *testing.T) {
	rndResourceName := utils.GenerateRandomResourceName()

	// resourceName is resourceIdentifier . resourceName
	resourceName := "cloudflare_token_validation_config." + rndResourceName
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// load a series of test keys
	var jwks JWKS
	require.NoError(t, json.Unmarshal([]byte(acctest.LoadTestCase("test-keys.json")), &jwks))

	resource.Test(t, resource.TestCase{
		IsUnitTest:               false,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// create a new token config
			{
				Config: testAccCloudflareTokenConfig(rndResourceName, zoneID, "title", "description", []string{`http.request.headers["x-auth"][0]`}, jwks.CloneFiltered("ES256")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "title", "title"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "token_type", "JWT"),
					resource.TestCheckResourceAttr(resourceName, "token_sources.0", "http.request.headers[\"x-auth\"][0]"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.alg", "ES256"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kid", "es256-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kty", "EC"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.x", checkHasField("x")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.y", checkHasField("y")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.crv", checkHasField("crv")),
				),
			},
			// edit that config
			{
				Config: testAccCloudflareTokenConfig(rndResourceName, zoneID, "title2", "description2", []string{`http.request.headers["x-auth"][0]`, `http.request.cookies["auth"][0]`}, jwks.CloneFiltered("ES256", "PS256")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "title", "title2"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "token_type", "JWT"),
					resource.TestCheckResourceAttr(resourceName, "token_sources.0", "http.request.headers[\"x-auth\"][0]"),
					resource.TestCheckResourceAttr(resourceName, "token_sources.1", "http.request.cookies[\"auth\"][0]"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.alg", "ES256"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kid", "es256-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kty", "EC"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.x", checkHasField("x")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.y", checkHasField("y")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.crv", checkHasField("crv")),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.alg", "PS256"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.kid", "ps256-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.kty", "RSA"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.1.e", checkHasField("e")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.1.n", checkHasField("n")),
				),
			},

			// ensure all other supported keys are accepted, above already tested es256 and ps256, we can at most supply 4 keys per config
			{
				Config: testAccCloudflareTokenConfig(rndResourceName, zoneID, "title", "description", []string{`http.request.headers["x-auth"][0]`}, jwks.CloneFiltered("ES384", "PS384", "PS512", "RS256")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "title", "title"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "token_type", "JWT"),
					resource.TestCheckResourceAttr(resourceName, "token_sources.0", "http.request.headers[\"x-auth\"][0]"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.#", "4"),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.alg", "ES384"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kid", "es384-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kty", "EC"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.x", checkHasField("x")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.y", checkHasField("y")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.crv", checkHasField("crv")),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.alg", "PS384"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.kid", "ps384-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.kty", "RSA"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.1.e", checkHasField("e")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.1.n", checkHasField("n")),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.2.alg", "PS512"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.2.kid", "ps512-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.2.kty", "RSA"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.2.e", checkHasField("e")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.2.n", checkHasField("n")),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.3.alg", "RS256"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.3.kid", "rs256-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.3.kty", "RSA"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.3.e", checkHasField("e")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.3.n", checkHasField("n")),
				),
			},

			// ensure all other supported keys are accepted
			{
				Config: testAccCloudflareTokenConfig(rndResourceName, zoneID, "title", "description", []string{`http.request.headers["x-auth"][0]`}, jwks.CloneFiltered("RS384", "RS512")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "title", "title"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "token_type", "JWT"),
					resource.TestCheckResourceAttr(resourceName, "token_sources.0", "http.request.headers[\"x-auth\"][0]"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.alg", "RS384"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kid", "rs384-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.0.kty", "RSA"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.e", checkHasField("e")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.0.n", checkHasField("n")),

					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.alg", "RS512"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.kid", "rs512-kid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.keys.1.kty", "RSA"),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.1.e", checkHasField("e")),
					resource.TestCheckResourceAttrWith(resourceName, "credentials.keys.1.n", checkHasField("n")),
				),
			},

			// deletes are implicitly tested

			// ensure import works
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("not found: %s", resourceName)
					}
					return fmt.Sprintf("%s/%s", zoneID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccCloudflareTokenConfig(resourceName, zone string, title string, description string, tokenSources []string, credentials JWKS) string {
	tokenSourcesStrings := make([]string, 0, len(tokenSources))
	for _, tokenSource := range tokenSources {
		tokenSourcesStrings = append(tokenSourcesStrings, fmt.Sprintf(`"%s"`, strings.ReplaceAll(tokenSource, `"`, `\"`)))
	}

	keys := []string{}
	for _, key := range credentials.Keys {
		if key.Kty == "EC" {
			keys = append(keys, acctest.LoadTestCase("ec_key.tf", key.Alg, key.Kid, *key.X, *key.Y, *key.Crv))
		} else {
			keys = append(keys, acctest.LoadTestCase("rsa_key.tf", key.Alg, key.Kid, *key.E, *key.N))
		}
	}

	return acctest.LoadTestCase("config.tf", resourceName, zone, title, description, strings.Join(tokenSourcesStrings, ", "), strings.Join(keys, ",\n"))
}

func checkHasField(name string) resource.CheckResourceAttrWithFunc {
	return func(value string) error {
		if len(value) > 0 {
			return nil
		}
		return fmt.Errorf("%s is empty", name)
	}
}
