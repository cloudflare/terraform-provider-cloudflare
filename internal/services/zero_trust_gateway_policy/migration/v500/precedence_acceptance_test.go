package v500_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

// The first check runs both migration handlers on real v4 state before a v5
// refresh can hide corruption. The migration then follows the documented manual
// review: reconcile all configured precedences with the captured live values.
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_PrecedenceOrder(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	// Avoid a chance CRC32 suffix collision in the two intentionally equal
	// configured precedences. The live values must be distinct for this case.
	for acctest.GatewayPolicyV4PrecedenceChange(gatewayPrecedenceRuleName(rnd, "allow_one"), 10116).Before ==
		acctest.GatewayPolicyV4PrecedenceChange(gatewayPrecedenceRuleName(rnd, "allow_two"), 10116).Before {
		rnd = utils.GenerateRandomResourceName()
	}
	configured := map[string]int64{"allow_one": 10116, "allow_two": 10116, "block": 10117}
	labels := []string{"allow_one", "allow_two", "block"}
	live := make(map[string]cloudflare.TeamsRule)
	var initialOrder []string
	tmpDir := t.TempDir()
	v4Config := gatewayPrecedenceConfig(rnd, accountID, configured)

	capture := gatewayPrecedenceStateCheck(func(ctx context.Context, state *tfjson.State) error {
		client, err := acctest.SharedV1Client()
		if err != nil {
			return err
		}
		for _, label := range labels {
			address := "cloudflare_teams_rule." + rnd + "_" + label
			attrs, err := gatewayPrecedenceAttributes(state, address)
			if err != nil {
				return err
			}
			raw, err := json.Marshal(attrs)
			if err != nil {
				return err
			}
			for _, migrationPath := range []string{"MoveState", "UpgradeFromVersion0"} {
				migrated, err := migrateGatewayPolicyRawState(ctx, migrationPath, &tfprotov6.RawState{JSON: raw})
				if err != nil {
					return fmt.Errorf("%s %s: %w", address, migrationPath, err)
				}
				if migrated.Precedence.IsNull() || migrated.Precedence.IsUnknown() || migrated.Precedence.ValueInt64() != configured[label] {
					return fmt.Errorf("%s %s changed source precedence %d to %s before refresh", address, migrationPath, configured[label], migrated.Precedence)
				}
			}
			id, ok := attrs["id"].(string)
			if !ok || id == "" {
				return fmt.Errorf("%s has no rule ID", address)
			}
			rule, err := client.TeamsRule(ctx, accountID, id)
			if err != nil {
				return err
			}
			expected := acctest.GatewayPolicyV4PrecedenceChange(gatewayPrecedenceRuleName(rnd, label), configured[label]).Before
			if rule.Precedence != uint64(expected) {
				return fmt.Errorf("%s: v4 API precedence = %d, expected %d", address, rule.Precedence, expected)
			}
			live[label] = rule
		}
		initialOrder = gatewayPrecedenceOrder(live)
		if live["allow_one"].Precedence == live["allow_two"].Precedence || initialOrder[len(initialOrder)-1] != "block" {
			return fmt.Errorf("expected two distinct allow rules before block, got order %v", initialOrder)
		}
		t.Logf("Captured live order %v; both handlers preserved real v4 source precedences before refresh", initialOrder)
		return nil
	})

	verify := gatewayPrecedenceStateCheck(func(ctx context.Context, state *tfjson.State) error {
		client, err := acctest.SharedV1Client()
		if err != nil {
			return err
		}
		current := make(map[string]cloudflare.TeamsRule)
		for _, label := range labels {
			address := "cloudflare_zero_trust_gateway_policy." + rnd + "_" + label
			attrs, err := gatewayPrecedenceAttributes(state, address)
			if err != nil {
				return err
			}
			before := live[label]
			precedence, err := json.Marshal(attrs["precedence"])
			if err != nil {
				return err
			}
			if attrs["id"] != before.ID || string(precedence) != fmt.Sprint(before.Precedence) {
				return fmt.Errorf("%s: migrated ID or precedence does not match the live snapshot", address)
			}
			after, err := client.TeamsRule(ctx, accountID, before.ID)
			if err != nil {
				return err
			}
			if after.ID != before.ID || after.Name != before.Name || after.Action != before.Action || after.Precedence != before.Precedence {
				return fmt.Errorf("%s: migration changed live rule identity, action, or precedence (%d -> %d)", address, before.Precedence, after.Precedence)
			}
			current[label] = after
		}
		if order := gatewayPrecedenceOrder(current); !reflect.DeepEqual(order, initialOrder) {
			return fmt.Errorf("live order changed from %v to %v", initialOrder, order)
		}
		return nil
	})

	migration := acctest.MigrationV2TestStepForGatewayPolicy(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", false, []statecheck.StateCheck{verify}, nil)
	migration.PreConfig = func() {
		// tf-migrate cannot access live API order. Simulate a user reviewing all
		// policies consistently, without collapsing the v4 hash suffixes.
		reconciled := make(map[string]int64)
		for _, label := range labels {
			reconciled[label] = int64(live[label].Precedence)
		}
		cfg := gatewayPrecedenceConfig(rnd, accountID, reconciled)
		acctest.WriteOutConfig(t, cfg, tmpDir)
		acctest.RunMigrationV2Command(t, cfg, tmpDir, "v4", "v5")
		if err := os.Remove(filepath.Join(tmpDir, "provider.tf")); err != nil {
			t.Fatal(err)
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: acctest.GetLastV4Version()},
				},
				Config:            v4Config,
				ConfigStateChecks: []statecheck.StateCheck{capture},
			},
			migration,
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				PlanOnly:                 true,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}

func gatewayPrecedenceRuleName(rnd, label string) string {
	return "tf-test-precedence-" + rnd + "-" + label
}

func gatewayPrecedenceConfig(rnd, accountID string, precedences map[string]int64) string {
	var config strings.Builder
	for _, label := range []string{"allow_one", "allow_two", "block"} {
		action, traffic := "allow", "net.dst.ip in {192.0.2.1 192.0.2.2}"
		if label == "block" {
			action, traffic = "block", "net.dst.ip in {192.0.2.0/24}"
		}
		fmt.Fprintf(&config, `
resource "cloudflare_teams_rule" %q {
  account_id = %q
  name = %q
  description = "Gateway precedence migration regression"
  precedence = %d
  enabled = true
  action = %q
  filters = ["l4"]
  traffic = %q
}
`, rnd+"_"+label, accountID, gatewayPrecedenceRuleName(rnd, label), precedences[label], action, traffic)
	}
	return config.String()
}

type gatewayPrecedenceStateCheck func(context.Context, *tfjson.State) error

func (check gatewayPrecedenceStateCheck) CheckState(ctx context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
	resp.Error = check(ctx, req.State)
}

func gatewayPrecedenceAttributes(state *tfjson.State, address string) (map[string]interface{}, error) {
	if state != nil && state.Values != nil && state.Values.RootModule != nil {
		for _, resource := range state.Values.RootModule.Resources {
			if resource.Address == address {
				return resource.AttributeValues, nil
			}
		}
	}
	return nil, fmt.Errorf("missing state for %s", address)
}

func gatewayPrecedenceOrder(rules map[string]cloudflare.TeamsRule) []string {
	labels := make([]string, 0, len(rules))
	for label := range rules {
		labels = append(labels, label)
	}
	sort.Slice(labels, func(i, j int) bool { return rules[labels[i]].Precedence < rules[labels[j]].Precedence })
	return labels
}
