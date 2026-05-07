---
layout: "cloudflare"
page_title: "Migrating to Cloudflare Provider v5"
description: Terraform Cloudflare Provider v5 Migration Guide
---

# Terraform Cloudflare Provider v5 Migration Guide

This guide covers the recommended migration path to the latest version of the
Cloudflare Terraform Provider v5. It applies whether you are coming from v4 or
upgrading from an earlier v5 release.

Starting with v5.19, the provider includes **automatic state upgraders** that
handle Terraform state migration transparently. Combined with the
[tf-migrate] CLI tool for HCL configuration changes, the migration process is
significantly simpler than previous approaches.

~> **Grit-based migration is deprecated.** This guide supersedes the Grit-based
migration instructions in the [version 5 upgrade guide]. You do **not** need
Grit to migrate. Grit patterns are no longer supported and will be removed in
a future release. The [version 5 upgrade guide] remains a useful reference for
per-resource attribute change details if you prefer manual HCL changes.

## Quick Reference

| Your Current Version | What To Do |
|---|---|
| v4.x (not v4.52.5) | Upgrade to v4.52.5 first, then follow [Path A](#migration-path-a-from-v4-to-v5) |
| v4.52.5 | Follow [Path A](#migration-path-a-from-v4-to-v5) |
| v5.0 -- v5.16 | Follow [Path B: Users on v5.16 or earlier](#users-on-v516-or-earlier) |
| v5.17 -- v5.18 | Follow [Path B: Users on v5.17 or v5.18](#users-on-v517-or-v518) |
| v5.19+ | Normal minor version upgrade. No migration steps needed. |

## Prerequisites

~> **IMPORTANT: Back up your Terraform state before migrating.** The v5
provider's state upgraders will automatically transform your state when you
run `terraform plan` or `terraform apply`. While these transformations are
tested extensively, you should always have a backup in case you need to
rollback. Run `terraform state pull > terraform.tfstate.backup` before
starting, or ensure your remote backend has versioning enabled.

Before starting the migration:

- **Terraform 1.8+** is recommended if you use any
  [renamed resources](#resource-rename-reference). It enables `moved` blocks,
  which provide the smoothest migration experience. On older Terraform
  versions, you can use `terraform state mv` instead -- see
  [Using `terraform state mv` (Terraform < 1.8)](#using-terraform-state-mv-terraform--18).
- **tf-migrate** -- Install the [tf-migrate] CLI tool for automatic HCL
  configuration changes. Download the latest release for your platform from
  the [tf-migrate releases page](https://github.com/cloudflare/tf-migrate/releases):

  ```bash
  # Example for macOS (ARM64)
  curl -LO https://github.com/cloudflare/tf-migrate/releases/download/v1.0.0/tf-migrate_1.0.0_darwin_arm64.tar.gz
  tar -xzf tf-migrate_1.0.0_darwin_arm64.tar.gz
  chmod +x tf-migrate
  sudo mv tf-migrate /usr/local/bin/
  
  # Or build from source:
  git clone https://github.com/cloudflare/tf-migrate.git
  cd tf-migrate
  make
  # Binary available at ./bin/tf-migrate
  ```

  For a complete list of supported resources and data sources, see the
  [tf-migrate README][tf-migrate].

- **Back up** your Terraform state and configuration files. Use version control.
- **Lock state** -- Ensure no concurrent Terraform operations are running.

## Understanding the Migration

There are three categories of changes between v4 and v5:

1. **Attribute changes** -- Field renames, nesting changes, type changes
   (e.g., `enabled` to `value`, flat attributes to nested objects). These
   affect your HCL configuration files (`.tf`).

2. **State schema changes** -- Internal state representation differences
   between v4 (SDKv2) and v5 (Plugin Framework). These are handled
   **automatically** by the provider's built-in state upgraders starting in
   v5.19 for [supported resources](#resources-with-state-upgraders). No manual
   state file editing is required for these resources.

3. **Resource renames** -- 40+ resources have new type names (e.g.,
   `cloudflare_record` to `cloudflare_dns_record`). These require `moved`
   blocks in HCL (Terraform 1.8+). 23 of these renamed resources include
   `MoveState` handlers that automatically transform state during the move.
   The remaining renamed resources require `terraform state rm` +
   `terraform import`.

### How It Works

```
 HCL Configuration ──> tf-migrate rewrites .tf files
 State Migration   ──> Provider state upgraders (automatic on plan/apply)
 Resource Renames  ──> moved blocks (Terraform 1.8+)
```

### What tf-migrate Handles

`tf-migrate` automatically handles:

- **Resource type renames** -- Updates resource types in your HCL (e.g.,
  `cloudflare_record` to `cloudflare_dns_record`)
- **Attribute transformations** -- Renames, restructures, and removes deprecated
  attributes
- **`moved` block generation** -- Creates `moved` blocks for renamed resources
  so Terraform knows to preserve state
- **Cross-file reference updates** -- Updates references to renamed resources
  across all `.tf` files (e.g., `cloudflare_record.example.id` becomes
  `cloudflare_dns_record.example.id`)
- **Data source migrations** -- Transforms data source configurations and
  updates output attribute references (e.g., `data.cloudflare_zones.example.zones`
  becomes `data.cloudflare_zones.example.result`)

For a complete list of supported resources, see the [tf-migrate README][tf-migrate].

---

## Migration Path A: From v4 to v5

This path is for users currently on any v4 release.

### Step 1: Upgrade to v4.52.5

~> If you are on any v4 version earlier than v4.52.5, you **must** upgrade to
v4.52.5 first. This ensures transitional state updates are applied before
the v5 migration.

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "4.52.5"
    }
  }
}
```

```bash
terraform init -upgrade
terraform plan   # Verify no unexpected changes
terraform apply  # Apply any transitional state updates
```

### Step 2: Migrate HCL Configuration

Use `tf-migrate` to automatically rewrite your `.tf` files for v5 compatibility.

```bash
# Preview changes without modifying files
tf-migrate migrate --source-version v4 --target-version v5 --dry-run

# Apply the migration (creates .bak backups by default)
tf-migrate migrate --source-version v4 --target-version v5
```

For projects with `.tf` files in a specific directory:

```bash
tf-migrate migrate \
  --config-dir ./terraform \
  --source-version v4 \
  --target-version v5
```

To migrate only specific resources:

```bash
tf-migrate migrate \
  --resources dns_record,zero_trust_list \
  --source-version v4 \
  --target-version v5
```

**Manual alternative:** If you prefer to make HCL changes manually or
`tf-migrate` does not fully cover your configuration (e.g., complex modules
or dynamic blocks), refer to the [version 5 upgrade guide] for per-resource
attribute change documentation.

### Step 3: Update Provider Version

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}
```

```bash
terraform init -upgrade
```

On the next `terraform plan` or `terraform apply`, the provider's state
upgraders will automatically migrate your state from v4 format to v5. This is
transparent and requires no manual intervention.

### Step 4: Handle Resource Renames

`tf-migrate` handles resource renames automatically -- it rewrites resource
types in your HCL and generates the corresponding `moved` blocks. If you ran
`tf-migrate` in Step 2, no further action is needed. Proceed to Step 5.

If you are migrating HCL manually (without `tf-migrate`), you need to add
`moved` blocks so Terraform knows the renamed resource is the same object:

```hcl
moved {
  from = cloudflare_record.example
  to   = cloudflare_dns_record.example
}

moved {
  from = cloudflare_tunnel.example
  to   = cloudflare_zero_trust_tunnel_cloudflared.example
}
```

The `moved` blocks can be removed after a successful `terraform apply`.

For a complete list of resource renames, see the
[resource rename reference](#resource-rename-reference) below.

### Step 5: Verify

```bash
terraform plan
```

The plan should show no changes, or only expected computed attribute updates
that will resolve on the next apply. If you see unexpected resource
replacements, check:

- All renamed resources have corresponding `moved` blocks.
- HCL attribute changes are complete (compare against the
  [version 5 upgrade guide] if needed).

```bash
terraform apply
```

After `terraform apply`, run `terraform plan` again. The plan should now show
**no changes**, or only the
[perpetual computed field differences](#perpetual-computed-field-differences)
listed below.

For a detailed breakdown of what to expect on the first plan and what persists
after apply, see
[Expected Plan Changes After Migration](#expected-plan-changes-after-migration).

---

## Expected Plan Changes After Migration

After completing the migration steps, the first `terraform plan` will show
changes across several categories. All of these are expected and benign --
they reflect differences between the v4 and v5 provider internals, not actual
infrastructure changes.

### One-Time Changes (Resolve After First Apply)

These changes appear on the first `terraform plan` after migration but
disappear after a single `terraform apply`.

#### Computed value refreshes

The most common category. Many read-only attributes show
`(known after apply)` because the v5 provider defines them differently than
v4. The provider needs to re-read these values from the API.

```
+ warning_status = (known after apply)
+ deleted_at     = (known after apply)
~ version        = 28 -> (known after apply)
~ created_at     = "2025-12-16T17:36:32Z" -> (known after apply)
```

These affect most resources and can account for hundreds of attribute
changes. They are harmless and resolve completely after apply.

#### State upgrader gaps

The v5 provider's state upgraders cannot carry forward every field from v4
state. These attributes exist in your HCL or in the API but are missing from
the migrated state until the provider refreshes them:

| Resource | Attributes | Details |
|---|---|---|
| `cloudflare_account` | `settings.abuse_contact_email`, `unit`, `managed_by`, `created_on` | New v5 computed fields not present in v4 state. |
| `cloudflare_account_member` | `policies`, `user` | New v5 computed fields not present in v4 state. |
| `cloudflare_custom_ssl` | `keyless_server` | New v5 computed field not present in v4 state. |
| `cloudflare_observatory_scheduled_test` | `schedule`, `test` | New v5 computed fields not present in v4 state. |
| `cloudflare_page_rule` | `created_on`, `modified_on` | New v5 computed timestamp fields. |
| `cloudflare_r2_bucket` | `jurisdiction` | v5 adds `jurisdiction = "default"` to all R2 buckets. Not present in v4 state. |
| `cloudflare_zone` | `plan`, `meta`, `owner`, `tenant`, `tenant_unit` | v5 restructured these as computed-only nested objects. Will refresh from API. |
| `cloudflare_zero_trust_access_identity_provider` | `support_groups`, `conditional_access_enabled` | Azure/OIDC fields not carried forward by the state upgrader. |
| `cloudflare_zero_trust_access_identity_provider` | `idp_public_certs` | SAML providers: v4 `idp_public_cert` (string) becomes `idp_public_certs` (list) in v5. The state upgrader does not transform this. |
| `cloudflare_zero_trust_gateway_settings` | `settings.host_selector`, `settings.inspection`, `settings.sandbox` | New v5 settings fields not present in v4 state. |

#### Resource creation from splits

Some v4 resources split into multiple v5 resources. The new resources must be
created on the first apply:

| v4 Resource | New v5 Resource Created | Why |
|---|---|---|
| `cloudflare_argo` | `cloudflare_argo_tiered_caching` | v4 `cloudflare_argo` managed both smart routing and tiered caching. v5 splits them. The `cloudflare_argo_smart_routing` resource inherits the state via `moved`, but `cloudflare_argo_tiered_caching` must be created fresh. |
| `cloudflare_tiered_cache` (with `cache_type = "generic"`) | `cloudflare_argo_tiered_caching` | Similar split. The `cloudflare_tiered_cache` resource retains its state, but the companion `cloudflare_argo_tiered_caching` resource must be created. |

#### Precedence normalization

`cloudflare_zero_trust_gateway_policy` resources will show precedence values
changing from large internal values to the user-facing values:

```
~ precedence = 1400392 -> 1400
```

This is a one-time normalization applied by the state upgrader.

#### IP CIDR normalization

`cloudflare_list` resources with IP items may show CIDR notation being
stripped (e.g., `10.0.0.0/8` to `10.0.0.0` with the comment updated) to
match the v5 schema format.

#### Account settings restructure

`cloudflare_account` resources will show:

- `enforce_twofactor` moves from top-level to `settings.enforce_twofactor`
- New computed fields (`settings.abuse_contact_email`, `unit`, `managed_by`, `created_on`) appear as `(known after apply)`

#### Account member field renames

`cloudflare_account_member` resources will show field renames:

```
~ email_address = "user@example.com" -> null
+ email         = "user@example.com"
~ role_ids      = ["abc123"] -> null
+ roles         = ["abc123"]
```

New computed fields (`policies`, `user`) appear as `(known after apply)`.

#### Custom SSL restructure

`cloudflare_custom_ssl` resources will show:

- Nested `custom_ssl_options[0].{certificate,private_key,bundle_method,type}` hoisted to top-level attributes
- `geo_restrictions` changes from string to nested object with `label` attribute
- `custom_ssl_priority` is dropped (was write-only in v4)
- Timestamp fields change format

#### Page rule status default change

`cloudflare_page_rule` resources with no explicit `status` may show:

```
~ status = "" -> "active"
```

The v5 default changed from `"active"` to `"disabled"`. The state upgrader
preserves the v4 behavior by explicitly setting `"active"` for resources
without a status.

Deprecated fields `disable_railgun` and `minify` are dropped during migration.

#### Zone field restructure

`cloudflare_zone` resources will show:

- `zone` renamed to `name`
- `account_id` moves to `account.id` (nested object)
- `jump_start` is dropped (removed in v5)
- `plan` and `meta` become computed-only (will refresh from API)

#### Zero Trust Gateway Settings restructure

`cloudflare_zero_trust_gateway_settings` resources will show significant
restructuring:

- Flat booleans (`activity_log_enabled`, `tls_decrypt_enabled`, `protocol_detection_enabled`) move to `settings.*` nested objects
- `antivirus.notification_settings[0].message` renamed to `notification_settings.msg`
- `logging`, `proxy`, `ssh_session_log`, `payload_log` are dropped

#### Zero Trust Access Policy normalization

`cloudflare_zero_trust_access_policy` resources normalize `false` boolean
values to `null` for optional fields (`isolation_required`,
`purpose_justification_required`, `approval_required`). This prevents drift
since the API treats `false` and `null` as equivalent.

#### Load Balancer field renames

`cloudflare_load_balancer` resources will show field renames:

- `default_pool_ids` renamed to `default_pools`
- `fallback_pool_id` renamed to `fallback_pool`
- Type conversions: `ttl`, `session_affinity_ttl`, `drain_duration` change from Int64 to Float64

`cloudflare_load_balancer_pool` resources will show:

- `check_regions` changes from Set to List (may show reordering)
- `origins` array may show reordering on first plan
- `load_shedding` and `origin_steering` change from array blocks to nested objects

#### Turnstile Widget field changes

`cloudflare_turnstile_widget` resources will show:

- `domains` changes from Set to List (alphabetically sorted)
- `off_label` renamed to `offlabel` (case change)
- New computed fields (`created_on`, `modified_on`, `clearance_level`, `ephemeral_id`)

#### Schema type changes

Some attributes change representation between v4 and v5:

```
- description   = "" -> null       # Empty string becomes null
- pool_weights  = {} -> null       # Empty map becomes null
- enabled_entries = [] -> null     # Empty list becomes null
```

### Perpetual Computed Field Differences

A small number of changes persist even after `terraform apply`. These are
ongoing v5 provider behaviors where the plan output does not match the
applied state. They do **not** represent actual infrastructure changes and
are safe to ignore.

| Resource | Attribute | Behavior |
|---|---|---|
| `cloudflare_healthcheck` | `description` | The API returns `""` but the v5 schema represents it as `null`. Shows `"" -> null` on every plan for healthchecks with no description set. |
| `cloudflare_healthcheck` | `status`, `failure_reason`, `tcp_config`, `http_config` | Read-only fields that refresh to `(known after apply)` on every plan. `tcp_config` appears on HTTP healthchecks and `http_config` on TCP healthchecks (the "other" config type). |
| `cloudflare_zero_trust_gateway_policy` | `duration` | The provider normalizes `"24h"` to `"24h0m0s"` on read, causing a perpetual diff. |
| `cloudflare_zero_trust_gateway_policy` | Various rule settings (`read_only`, `schedule`, `source_account`, `override_host`, `sharable`, `warning_status`, `deleted_at`, `expiration`, and others) | Computed read-only fields that refresh to `(known after apply)` on every plan. Affects a subset of gateway policies (typically 5-6 resources). |
| `cloudflare_zero_trust_gateway_policy` | `audit_ssh.command_logging` | Computed default (`true`) that re-appears on every plan for resources with `audit_ssh` settings. |
| `cloudflare_zero_trust_gateway_policy` | `block_page_enabled`, `ip_categories` | Computed defaults (`false`) that appear on every plan for certain rule settings. |
| `cloudflare_zero_trust_dlp_predefined_profile` | `entries`, `name`, `open_access` | Read-only API fields that the provider recalculates on every plan. |
| `cloudflare_zero_trust_dlp_predefined_profile` | `enabled_entries` | API returns `[]` but config uses `null`. Shows `[] -> null` on every plan. |

~> These perpetual diffs are provider-level issues being tracked for
resolution. They do not affect your infrastructure and `terraform apply` will
succeed without making changes.

---

## Migration Path B: Upgrading Within v5

This path is for users already on a v5 release who want to upgrade to the
latest version.

### Users on v5.16 or Earlier

~> Users on v5.0 through v5.16 **must** upgrade to v5.17 or v5.18 as a
stepping stone before upgrading to v5.19+. This applies if you use **any** of
the [resources requiring stepping-stone upgrades](#resources-requiring-stepping-stone-upgrades).

**Why:** In v5.17, the provider bumped the internal `schema_version` from 0 to
1 for these resources with a safe no-op state upgrader. In v5.19, the schema
version jumps to 500, and the upgrader at slot 1 (which handles the 1 to 500
transition) expects state that has already been through the 0 to 1 bump. Skipping
the stepping stone means the wrong upgrader runs on your state.

**Step 1:** Upgrade to v5.17 or v5.18:

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5.18.0"
    }
  }
}
```

```bash
terraform init -upgrade
terraform plan
terraform apply  # Bumps schema_version from 0 to 1 for affected resources
```

**Step 2:** Upgrade to the latest:

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}
```

```bash
terraform init -upgrade
terraform plan   # State upgraders handle version 1 to 500 automatically
terraform apply
```

### Users on v5.17 or v5.18

You already have the stepping-stone state version. Upgrade directly to the
latest:

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = ">= 5.19.0"
    }
  }
}
```

```bash
terraform init -upgrade
terraform plan
terraform apply
```

The state upgraders will handle the version 1 to 500 transition automatically.

### Users on v5.19+

No migration steps are needed. Upgrade normally:

```bash
terraform init -upgrade
terraform plan
terraform apply
```

---

## Resource Rename Reference

The following resources were renamed between v4 and v5. Resources marked with
**Auto** have `MoveState` handlers that automatically transform state during
a `moved` block operation. Resources marked **Manual** require
`terraform state rm` + `terraform import` because they do not include
automatic state transformation for the rename.

~> Terraform 1.8+ is recommended for `moved` blocks. On older versions, use
`terraform state mv` instead -- see
[below](#using-terraform-state-mv-terraform--18).

| v4 Resource | v5 Resource | State |
|---|---|---|
| `cloudflare_access_application` | `cloudflare_zero_trust_access_application` | Auto |
| `cloudflare_access_ca_certificate` | `cloudflare_zero_trust_access_short_lived_certificate` | Manual |
| `cloudflare_access_custom_page` | `cloudflare_zero_trust_access_custom_page` | Manual |
| `cloudflare_access_group` | `cloudflare_zero_trust_access_group` | Auto |
| `cloudflare_access_identity_provider` | `cloudflare_zero_trust_access_identity_provider` | Auto |
| `cloudflare_access_keys_configuration` | `cloudflare_zero_trust_access_key_configuration` | Manual |
| `cloudflare_access_mutual_tls_certificate` | `cloudflare_zero_trust_access_mtls_certificate` | Auto |
| `cloudflare_access_mutual_tls_hostname_settings` | `cloudflare_zero_trust_access_mtls_hostname_settings` | Manual |
| `cloudflare_access_organization` | `cloudflare_zero_trust_organization` | Manual |
| `cloudflare_access_policy` | `cloudflare_zero_trust_access_policy` | Auto |
| `cloudflare_access_service_token` | `cloudflare_zero_trust_access_service_token` | Auto |
| `cloudflare_access_tag` | `cloudflare_zero_trust_access_tag` | Manual |
| `cloudflare_argo` | `cloudflare_argo_smart_routing` | Auto |
| `cloudflare_argo` / `cloudflare_tiered_cache` | `cloudflare_argo_tiered_caching` | Auto |
| `cloudflare_authenticated_origin_pulls` | `cloudflare_authenticated_origin_pulls_settings` | Auto |
| `cloudflare_authenticated_origin_pulls_certificate` (per-hostname) | `cloudflare_authenticated_origin_pulls_hostname_certificate` | Auto |
| `cloudflare_device_dex_test` | `cloudflare_zero_trust_dex_test` | Auto |
| `cloudflare_device_managed_networks` | `cloudflare_zero_trust_device_managed_networks` | Auto* |
| `cloudflare_device_policy_certificates` | `cloudflare_zero_trust_device_certificates` | Manual |
| `cloudflare_device_posture_integration` | `cloudflare_zero_trust_device_posture_integration` | Manual |
| `cloudflare_device_posture_rule` | `cloudflare_zero_trust_device_posture_rule` | Auto |
| `cloudflare_device_settings_policy` | `cloudflare_zero_trust_device_custom_profile` or `cloudflare_zero_trust_device_default_profile` | Manual |
| `cloudflare_dlp_custom_profile` | `cloudflare_zero_trust_dlp_custom_profile` | Auto |
| `cloudflare_dlp_predefined_profile` | `cloudflare_zero_trust_dlp_predefined_profile` | Auto |
| `cloudflare_dlp_profile` | `cloudflare_zero_trust_dlp_custom_profile` or `cloudflare_zero_trust_dlp_predefined_profile` | Auto (see [DLP profile migration](#cloudflare_dlp_profile)) |
| `cloudflare_fallback_domain` / `cloudflare_zero_trust_local_fallback_domain` | `cloudflare_zero_trust_device_default_profile_local_domain_fallback` or `cloudflare_zero_trust_device_custom_profile_local_domain_fallback` | Manual |
| `cloudflare_gateway_app_types` | `cloudflare_zero_trust_gateway_app_types` | Manual |
| `cloudflare_gre_tunnel` | `cloudflare_magic_wan_gre_tunnel` | Manual |
| `cloudflare_ipsec_tunnel` | `cloudflare_magic_wan_ipsec_tunnel` | Manual |
| `cloudflare_managed_headers` | `cloudflare_managed_transforms` | Auto |
| `cloudflare_record` | `cloudflare_dns_record` | Auto |
| `cloudflare_risk_behavior` | `cloudflare_zero_trust_risk_behavior` | Manual |
| `cloudflare_split_tunnel` | `cloudflare_zero_trust_device_default_profile` or `cloudflare_zero_trust_device_custom_profile` | Manual |
| `cloudflare_static_route` | `cloudflare_magic_wan_static_route` | Manual |
| `cloudflare_teams_account` | `cloudflare_zero_trust_gateway_settings` | Manual |
| `cloudflare_teams_list` | `cloudflare_zero_trust_list` | Auto |
| `cloudflare_teams_location` | `cloudflare_zero_trust_dns_location` | Manual |
| `cloudflare_teams_proxy_endpoint` | `cloudflare_zero_trust_gateway_proxy_endpoint` | Manual |
| `cloudflare_teams_rule` | `cloudflare_zero_trust_gateway_policy` | Auto |
| `cloudflare_tunnel` | `cloudflare_zero_trust_tunnel_cloudflared` | Auto |
| `cloudflare_tunnel_config` | `cloudflare_zero_trust_tunnel_cloudflared_config` | Manual |
| `cloudflare_tunnel_route` | `cloudflare_zero_trust_tunnel_cloudflared_route` | Manual |
| `cloudflare_tunnel_virtual_network` | `cloudflare_zero_trust_tunnel_cloudflared_virtual_network` | Manual |
| `cloudflare_worker_cron_trigger` | `cloudflare_workers_cron_trigger` | Manual |
| `cloudflare_worker_domain` | `cloudflare_workers_custom_domain` | Manual |
| `cloudflare_worker_route` | `cloudflare_workers_route` | Auto |
| `cloudflare_worker_script` | `cloudflare_workers_script` | Auto |
| `cloudflare_workers_for_platforms_namespace` | `cloudflare_workers_for_platforms_dispatch_namespace` | Auto |

\* Exception for Terraform < 1.8: `cloudflare_device_managed_networks` ->
`cloudflare_zero_trust_device_managed_networks` does not support
`terraform state mv`. Use Terraform 1.8+ `moved` blocks, or use
`terraform state rm` + `terraform import`.

### Using `moved` Blocks (Terraform 1.8+)

~> If you used `tf-migrate` in Step 2, it has already generated the `moved`
blocks for all renamed resources. No manual action is needed -- skip to
[Step 5](#step-5-verify).

If you are migrating HCL manually, add a `moved` block for each renamed
resource. For resources marked **Auto**, the provider's `MoveState` handler
will transform the state automatically:

```hcl
moved {
  from = cloudflare_record.example
  to   = cloudflare_dns_record.example
}
```

After a successful `terraform apply`, you can remove the `moved` block.

### Using `terraform state rm` and `import` (Manual Resources)

For resources marked **Manual**, remove the old resource from state and import
into the new resource type:

```bash
# Remove the old resource from state
terraform state rm cloudflare_access_ca_certificate.example

# Import into the new resource type
terraform import cloudflare_zero_trust_access_short_lived_certificate.example <account_id>/<certificate_id>
```

Refer to the individual [resource documentation](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs)
for the correct import string format.

### Using `terraform state mv` (Terraform < 1.8)

If you are on a Terraform version older than 1.8, you cannot use `moved`
blocks for cross-resource-type renames. Instead, use `terraform state mv` to
rename resources directly in the state. The provider's state upgraders will
automatically transform the state on the next `terraform plan` or
`terraform apply`.

For resources marked **Auto** in the rename table:

```bash
# Move the resource in state from the old type to the new type
terraform state mv cloudflare_record.example cloudflare_dns_record.example
```

~> Exception: `cloudflare_device_managed_networks` ->
`cloudflare_zero_trust_device_managed_networks` cannot be migrated via
`terraform state mv` on Terraform < 1.8. Use Terraform 1.8+ with `moved`
blocks, or use `terraform state rm` + `terraform import`.

Then update the resource type in your HCL to match the new name. On the next
`terraform plan`, the provider's state upgrader detects the old
`schema_version` and transforms the state automatically.

For resources marked **Manual**, use `terraform state rm` + `terraform import`
as described [above](#using-terraform-state-rm-and-import-manual-resources).

---

## Resources Requiring Manual Migration

Some resources cannot be handled by state upgraders alone and require
additional manual steps after running `tf-migrate` or updating your HCL.

### Application-Scoped Access Policies

~> **`cloudflare_access_policy` resources with `application_id` require manual
config migration.** `tf-migrate` can automatically generate a `removed` block
to drop the old state entry without destroying the remote policy, but you still
must rewrite the policy as inline `policies` on the application resource.

In v4, `cloudflare_access_policy` could be used for both account-level policies
and application-scoped policies (when `application_id` was set). These two types
use different API endpoints:

- **Account-level policies**: `POST /access/policies/`
- **Application-scoped policies**: `POST /access/apps/{app_id}/policies/`

In v5, application-scoped policies are no longer standalone resources. Instead,
they are defined inline within the `cloudflare_zero_trust_access_application`
resource using the `policies` attribute.

Do not use a `moved` block for this scenario. Application-scoped policies must
be dropped from standalone state and rewritten as inline application policies.

#### Migration Steps

**1.** Identify application-scoped policies in your configuration:

```hcl
# This is an application-scoped policy (has application_id)
resource "cloudflare_access_policy" "app_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  application_id = "abc123"  # <-- This makes it application-scoped
  name           = "Allow employees"
  decision       = "allow"
  precedence     = 1

  include {
    email_domain = ["example.com"]
  }
}
```

**2.** Drop the old standalone policy from Terraform state without destroying
the remote policy.

If you use `tf-migrate`, this is generated automatically:

```hcl
removed {
  from = cloudflare_access_policy.app_policy
  lifecycle {
    destroy = false
  }
}
```

If you are not using `tf-migrate` (or are on an older migration output), run:

```bash
terraform state rm cloudflare_access_policy.app_policy
```

**3.** Add the policy inline in your application resource:

```hcl
resource "cloudflare_zero_trust_access_application" "my_app" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "My Application"
  domain     = "app.example.com"
  type       = "self_hosted"

  # Policies are now inline
  policies = [
    {
      name       = "Allow employees"
      decision   = "allow"
      precedence = 1
      include    = [{ email_domain = { domain = "example.com" } }]
    }
  ]
}
```

**4.** Run `terraform apply` to update the application with the inline policy.

#### Key Differences

| v4 | v5 |
|---|---|
| Separate `cloudflare_access_policy` resource with `application_id` | Inline `policies` attribute in `cloudflare_zero_trust_access_application` |
| `include { email_domain = ["example.com"] }` (block with array) | `include = [{ email_domain = { domain = "example.com" } }]` (attribute with objects) |
| `precedence` as top-level attribute | `precedence` inside each policy object |

~> **Account-level policies** (without `application_id`) migrate normally using
`tf-migrate` and are renamed to `cloudflare_zero_trust_access_policy`.

### `cloudflare_zone_settings_override`

In v4, `cloudflare_zone_settings_override` was a single resource that managed
all zone settings in one block. In v5, this has been replaced by individual
`cloudflare_zone_setting` resources — one per setting.

#### With tf-migrate (recommended)

`tf-migrate` automatically handles the full transformation. Run it as part of
Step 2 if you have not already done so.

**What tf-migrate generates** for each `cloudflare_zone_settings_override` resource:

- One `cloudflare_zone_setting` resource per setting
- One `import` block per setting so Terraform adopts the existing API state
  rather than creating new resources (requires Terraform 1.7+)
- One `removed` block so Terraform drops the old state entry without destroying
  anything (requires Terraform 1.7+)

For example, this v4 configuration:

```hcl
resource "cloudflare_zone_settings_override" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"

  settings {
    always_online     = "on"
    brotli            = "on"
    browser_cache_ttl = 14400

    minify {
      css  = "on"
      html = "on"
      js   = "off"
    }

    security_header {
      enabled            = true
      max_age            = 86400
      include_subdomains = true
      preload            = true
      nosniff            = true
    }
  }
}
```

Becomes:

```hcl
resource "cloudflare_zone_setting" "example_always_online" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  setting_id = "always_online"
  value      = "on"
}
import {
  to = cloudflare_zone_setting.example_always_online
  id = "0da42c8d2132a9ddaf714f9e7c920711/always_online"
}

resource "cloudflare_zone_setting" "example_brotli" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  setting_id = "brotli"
  value      = "on"
}
import {
  to = cloudflare_zone_setting.example_brotli
  id = "0da42c8d2132a9ddaf714f9e7c920711/brotli"
}

resource "cloudflare_zone_setting" "example_browser_cache_ttl" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  setting_id = "browser_cache_ttl"
  value      = 14400
}
import {
  to = cloudflare_zone_setting.example_browser_cache_ttl
  id = "0da42c8d2132a9ddaf714f9e7c920711/browser_cache_ttl"
}

resource "cloudflare_zone_setting" "example_minify" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  setting_id = "minify"
  value = {
    css  = "on"
    html = "on"
    js   = "off"
  }
}
import {
  to = cloudflare_zone_setting.example_minify
  id = "0da42c8d2132a9ddaf714f9e7c920711/minify"
}

resource "cloudflare_zone_setting" "example_security_header" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  setting_id = "security_header"
  value = {
    strict_transport_security = {
      enabled            = true
      include_subdomains = true
      max_age            = 86400
      nosniff            = true
      preload            = true
    }
  }
}
import {
  to = cloudflare_zone_setting.example_security_header
  id = "0da42c8d2132a9ddaf714f9e7c920711/security_header"
}

removed {
  from = cloudflare_zone_settings_override.example
  lifecycle {
    destroy = false
  }
}
```

~> `tf-migrate` skips the deprecated `universal_ssl` setting and maps
`zero_rtt` to the v5 setting ID `0rtt`. The resource name remains
`{name}_zero_rtt` but the `setting_id` is set to `"0rtt"`.

**After running tf-migrate, two manual steps are required:**

**Step 1: Remove the old state entry.**

The v5 provider has no schema for `cloudflare_zone_settings_override`, so
Terraform cannot read the old state entry to process the `removed` block.
Remove it before running `terraform plan`:

```bash
terraform state rm cloudflare_zone_settings_override.example
```

If the resource is inside a Terraform module:

```bash
terraform state rm module.<module_name>.cloudflare_zone_settings_override.example
```

~> `tf-migrate` prints the exact `terraform state rm` command for each
resource in its output. Look for the `⚠️ Action required` warnings.

**Step 2: Move import blocks to the root module (if using modules).**

`import` blocks are only valid in the root Terraform module. If the migrated
file is used as a child module, move the `import` blocks to your root module
and prefix the `to` address with the module path:

```hcl
# In your root module (e.g. main.tf), not inside the child module:
import {
  to = module.<module_name>.cloudflare_zone_setting.example_always_online
  id = "<zone_id>/always_online"
}
# ... repeat for each setting
```

If the migrated file is your root module (no module wrapping), the import
blocks work as-is and no action is needed.

**Step 3: Apply.**

```bash
terraform plan   # Shows imports and no creates/destroys
terraform apply  # Imports existing settings into state
terraform plan   # Should show no changes
```

#### Without tf-migrate

**1.** Remove the v4 resource from state:

```bash
terraform state rm cloudflare_zone_settings_override.example
```

**2.** Replace the `cloudflare_zone_settings_override` resource in your HCL
with individual `cloudflare_zone_setting` resources. Create one resource per
setting you were managing:

```hcl
resource "cloudflare_zone_setting" "example_always_online" {
  zone_id    = "<zone_id>"
  setting_id = "always_online"
  value      = "on"
}

resource "cloudflare_zone_setting" "example_brotli" {
  zone_id    = "<zone_id>"
  setting_id = "brotli"
  value      = "on"
}
# ... one resource per setting
```

Key differences from v4:

- Each setting is its own resource with a `setting_id` and `value` attribute.
- Nested blocks like `minify { css = "on" }` become `value = { css = "on" }`.
- The `security_header` block must be wrapped:
  `value = { strict_transport_security = { enabled = true, ... } }`.
- The `universal_ssl` setting has been removed. Do not create a resource for it.
- The v4 setting name `zero_rtt` maps to the v5 setting ID `0rtt`.

**3.** Import each new resource into state (format: `<zone_id>/<setting_id>`):

```bash
terraform import cloudflare_zone_setting.example_always_online <zone_id>/always_online
terraform import cloudflare_zone_setting.example_brotli <zone_id>/brotli
# ... repeat for each setting
```

**4.** Verify:

```bash
terraform plan
# Should show no changes
```

### `cloudflare_dlp_profile`

In v4, `cloudflare_dlp_profile` (or `cloudflare_zero_trust_dlp_profile`) was a
single resource that handled both custom and predefined DLP profiles via a
`type` attribute. In v5, these are separate resources:

- `type = "custom"` becomes `cloudflare_zero_trust_dlp_custom_profile`
- `type = "predefined"` becomes `cloudflare_zero_trust_dlp_predefined_profile`

`tf-migrate` handles the HCL transformation and generates `moved` blocks.
State migration is handled automatically by the provider's `MoveState`
handlers.

#### Custom profiles

Custom profiles are fully handled by `tf-migrate`. The key schema changes are:

- `type` attribute removed
- `entry` blocks converted to an `entries` list attribute
- `entry.id` removed
- `pattern` blocks converted to `pattern` attribute objects

No additional user intervention is required unless you use `dynamic "entry"`
blocks. Dynamic blocks cannot be automatically migrated because v5 uses
`entries` as a list attribute which does not support dynamic blocks. You must
manually convert them to a list comprehension:

```hcl
# Before (v4 - dynamic blocks)
dynamic "entry" {
  for_each = var.patterns
  content {
    name    = entry.value.name
    enabled = entry.value.enabled
    pattern {
      regex = entry.value.regex
    }
  }
}

# After (v5 - list comprehension)
entries = [
  for pattern in var.patterns : {
    name    = pattern.name
    enabled = pattern.enabled
    pattern = {
      regex = pattern.regex
    }
  }
]
```

#### Predefined profiles

~> After running `tf-migrate`, you must manually add `profile_id` to each
predefined profile resource. `tf-migrate` cannot derive this value because it
was a computed attribute in v4.

The `profile_id` is the UUID of the predefined profile. You can find it in your
existing Terraform state:

```bash
terraform state show cloudflare_dlp_profile.example | grep '"id"'
```

Or via the Cloudflare API:

```bash
curl -s https://api.cloudflare.com/client/v4/accounts/{account_id}/dlp/profiles \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" | jq '.result[] | select(.type == "predefined") | {name, id}'
```

Add the `profile_id` to each migrated predefined profile resource:

```hcl
resource "cloudflare_zero_trust_dlp_predefined_profile" "aws_keys" {
  account_id      = var.cloudflare_account_id
  profile_id      = "c8932cc4-3312-4152-8041-f3f257122dc4"  # Add this manually
  allowed_match_count = 3
  enabled_entries = ["aws-access-key-id", "aws-secret-key-id"]
}
```

Other predefined profile changes handled by `tf-migrate`:

- `type` attribute removed
- `name` attribute removed (read-only in v5)
- `id` renamed to `profile_id` (if present in v4 config)
- `entry` blocks replaced with `enabled_entries` list (only IDs of enabled
  entries)

---

## Resources Requiring Stepping-Stone Upgrades

The following resources require users on v5.16 or earlier to upgrade to
v5.17 or v5.18 before upgrading to v5.19+. This is because these resources
have a two-phase state upgrade: version 0 to 1 (applied in v5.17/v5.18), then
version 1 to 500 (applied in v5.19+).

| Resource |
|---|
| `cloudflare_account` |
| `cloudflare_api_shield` |
| `cloudflare_argo_smart_routing` |
| `cloudflare_argo_tiered_caching` |
| `cloudflare_authenticated_origin_pulls` |
| `cloudflare_authenticated_origin_pulls_certificate` |
| `cloudflare_certificate_pack` |
| `cloudflare_custom_hostname_fallback_origin` |
| `cloudflare_healthcheck` |
| `cloudflare_list` |
| `cloudflare_list_item` |
| `cloudflare_load_balancer_monitor` |
| `cloudflare_load_balancer_pool` |
| `cloudflare_logpull_retention` |
| `cloudflare_logpush_job` |
| `cloudflare_logpush_ownership_challenge` |
| `cloudflare_managed_transforms` |
| `cloudflare_mtls_certificate` |
| `cloudflare_notification_policy` |
| `cloudflare_notification_policy_webhooks` |
| `cloudflare_page_rule` |
| `cloudflare_pages_domain` |
| `cloudflare_pages_project` |
| `cloudflare_queue` |
| `cloudflare_r2_bucket` |
| `cloudflare_regional_hostname` |
| `cloudflare_spectrum_application` |
| `cloudflare_tiered_cache` |
| `cloudflare_turnstile_widget` |
| `cloudflare_url_normalization_settings` |
| `cloudflare_workers_custom_domain` |
| `cloudflare_workers_kv` |
| `cloudflare_workers_kv_namespace` |
| `cloudflare_zero_trust_access_mtls_certificate` |
| `cloudflare_zero_trust_access_service_token` |
| `cloudflare_zero_trust_device_custom_profile` |
| `cloudflare_zero_trust_device_custom_profile_local_domain_fallback` |
| `cloudflare_zero_trust_device_default_profile` |
| `cloudflare_zero_trust_device_default_profile_local_domain_fallback` |
| `cloudflare_zero_trust_device_managed_networks` |
| `cloudflare_zero_trust_device_posture_rule` |
| `cloudflare_zero_trust_dlp_custom_profile` |
| `cloudflare_zero_trust_dlp_predefined_profile` |
| `cloudflare_zero_trust_gateway_policy` |
| `cloudflare_zero_trust_list` |
| `cloudflare_zero_trust_tunnel_cloudflared` |
| `cloudflare_zero_trust_tunnel_cloudflared_config` |
| `cloudflare_zero_trust_tunnel_cloudflared_virtual_network` |
| `cloudflare_zone` |
| `cloudflare_zone_dnssec` |

**If you do not use any of these resources**, you can upgrade directly from any
v5 version to v5.19+.

### What Happens If You Skip the Stepping Stone

If you upgrade from v5.16 or earlier directly to v5.19+ while using one of
these resources, Terraform will attempt to run the wrong state upgrader on your
state data. This may result in an error like:

```
Error: Failed to upgrade resource state
```

**Recovery:** Pin back to v5.18, run `terraform apply` to apply the
intermediate state upgrade, then upgrade to v5.19+.

```hcl
# Temporary: pin to v5.18 to apply the stepping-stone upgrade
version = "~> 5.18.0"
```

```bash
terraform init -upgrade
terraform apply
```

Then update your version constraint to `>= 5.19.0` and run
`terraform init -upgrade` again.

---

## Resources with State Upgraders

The following resources have built-in state upgraders that automatically
transform v4 (SDKv2) state to v5 (Plugin Framework) state when you run
`terraform plan` or `terraform apply` on v5.19+. Resources **not** listed here
do not have automatic state migration and may require `terraform state rm` +
`terraform import` after upgrading.

| Product | Resource |
|---------|----------|
| **Zones** | `cloudflare_zone` |
| | `cloudflare_zone_setting` |
| | `cloudflare_zone_subscription` |
| **DNS** | `cloudflare_dns_record` |
| | `cloudflare_zone_dnssec` |
| **Load Balancers** | `cloudflare_load_balancer` |
| | `cloudflare_load_balancer_monitor` |
| | `cloudflare_load_balancer_pool` |
| **Rulesets** | `cloudflare_ruleset` |
| **Page Rules** | `cloudflare_page_rule` |
| **Managed Transforms** | `cloudflare_managed_transforms` |
| **URL Normalization** | `cloudflare_url_normalization_settings` |
| **Snippets** | `cloudflare_snippet` |
| | `cloudflare_snippet_rules` |
| **Workers** | `cloudflare_workers_script` |
| | `cloudflare_workers_route` |
| **KV** | `cloudflare_workers_kv` |
| | `cloudflare_workers_kv_namespace` |
| **Pages** | `cloudflare_pages_project` |
| **Cache** | `cloudflare_tiered_cache` |
| **Argo** | `cloudflare_argo_smart_routing` |
| | `cloudflare_argo_tiered_caching` |
| **Spectrum** | `cloudflare_spectrum_application` |
| **Addressing** | `cloudflare_regional_hostname` |
| | `cloudflare_byo_ip_prefix` |
| **Bot Management** | `cloudflare_bot_management` |
| **Healthchecks** | `cloudflare_healthcheck` |
| **Custom Pages** | `cloudflare_custom_pages` |
| **Rules** | `cloudflare_list` |
| | `cloudflare_list_item` |
| **Logpush** | `cloudflare_logpush_job` |
| | `cloudflare_logpush_ownership_challenge` |
| **Logs** | `cloudflare_logpull_retention` |
| **Alerting** | `cloudflare_notification_policy` |
| | `cloudflare_notification_policy_webhooks` |
| **R2** | `cloudflare_r2_bucket` |
| **User** | `cloudflare_api_token` |
| **Account** | `cloudflare_account` |
| | `cloudflare_account_member` |
| | `cloudflare_account_token` |
| **SSL/TLS** | `cloudflare_authenticated_origin_pulls` |
| | `cloudflare_authenticated_origin_pulls_certificate` |
| | `cloudflare_authenticated_origin_pulls_hostname_certificate` |
| | `cloudflare_authenticated_origin_pulls_settings` |
| | `cloudflare_certificate_pack` |
| | `cloudflare_custom_hostname` |
| | `cloudflare_custom_hostname_fallback_origin` |
| | `cloudflare_custom_ssl` |
| | `cloudflare_mtls_certificate` |
| | `cloudflare_origin_ca_certificate` |
| **Security** | `cloudflare_access_rule` |
| | `cloudflare_leaked_credential_check` |
| | `cloudflare_leaked_credential_check_rule` |
| **API Shield** | `cloudflare_api_shield` |
| | `cloudflare_api_shield_operation` |
| **Pages** | `cloudflare_pages_project` |
| | `cloudflare_pages_domain` |
| **Cache** | `cloudflare_regional_tiered_cache` |
| | `cloudflare_tiered_cache` |
| **Images** | `cloudflare_image_variant` |
| **Turnstile** | `cloudflare_turnstile_widget` |
| **Observatory** | `cloudflare_observatory_scheduled_test` |
| **Queues** | `cloudflare_queue` |
| | `cloudflare_queue_consumer` |
| **Workers** | `cloudflare_workers_custom_domain` |
| | `cloudflare_workers_for_platforms_dispatch_namespace` |
| **Zero Trust** | `cloudflare_zero_trust_access_application` |
| | `cloudflare_zero_trust_access_group` |
| | `cloudflare_zero_trust_access_identity_provider` |
| | `cloudflare_zero_trust_access_mtls_certificate` |
| | `cloudflare_zero_trust_access_mtls_hostname_settings` |
| | `cloudflare_zero_trust_access_policy` |
| | `cloudflare_zero_trust_access_service_token` |
| | `cloudflare_zero_trust_device_custom_profile` |
| | `cloudflare_zero_trust_device_custom_profile_local_domain_fallback` |
| | `cloudflare_zero_trust_device_default_profile` |
| | `cloudflare_zero_trust_device_default_profile_local_domain_fallback` |
| | `cloudflare_zero_trust_device_managed_networks` |
| | `cloudflare_zero_trust_device_posture_integration` |
| | `cloudflare_zero_trust_device_posture_rule` |
| | `cloudflare_zero_trust_dex_test` |
| | `cloudflare_zero_trust_dlp_custom_profile` |
| | `cloudflare_zero_trust_dlp_custom_entry` |
| | `cloudflare_zero_trust_dlp_predefined_profile` |
| | `cloudflare_zero_trust_dlp_predefined_entry` |
| | `cloudflare_zero_trust_dlp_integration_entry` |
| | `cloudflare_zero_trust_gateway_certificate` |
| | `cloudflare_zero_trust_gateway_policy` |
| | `cloudflare_zero_trust_gateway_settings` |
| | `cloudflare_zero_trust_list` |
| | `cloudflare_zero_trust_organization` |
| | `cloudflare_zero_trust_tunnel_cloudflared` |
| | `cloudflare_zero_trust_tunnel_cloudflared_config` |
| | `cloudflare_zero_trust_tunnel_cloudflared_route` |
| | `cloudflare_zero_trust_tunnel_cloudflared_virtual_network` |

---

## Migrating Data Sources

Data sources are simpler to migrate than resources because their state is
refreshed on every `terraform plan` -- there is no persistent state to upgrade.
You only need to update the HCL configuration.

`tf-migrate` handles the following data source migrations automatically:

### `data.cloudflare_zone`

The `name` and `account_id` attributes are now wrapped in a `filter`
attribute:

```hcl
# Before (v4)
data "cloudflare_zone" "example" {
  name       = "example.com"
  account_id = "f037e56e89293a057740de681ac9abbe"
}

# After (v5)
data "cloudflare_zone" "example" {
  filter = {
    name = "example.com"
    account = {
      id = "f037e56e89293a057740de681ac9abbe"
    }
  }
}
```

Data sources using only `zone_id` require no changes.

### `data.cloudflare_zones`

The `filter` block is removed and its fields are restructured:

- `filter.name` is hoisted to a top-level `name` attribute.
- `filter.account_id` becomes `account = { id = "..." }`.
- `filter.status` is hoisted to a top-level `status` attribute.
- `filter.lookup_type`, `filter.match`, and `filter.paused` are **removed**
  (no v5 equivalent). If you relied on these, you will need to filter results
  in your Terraform code.
- The output attribute `zones` is renamed to `result`. Update any references
  (e.g., `data.cloudflare_zones.example.zones` becomes
  `data.cloudflare_zones.example.result`).

```hcl
# Before (v4)
data "cloudflare_zones" "example" {
  filter {
    name       = "example.com"
    account_id = "f037e56e89293a057740de681ac9abbe"
    status     = "active"
  }
}

# After (v5)
data "cloudflare_zones" "example" {
  name   = "example.com"
  status = "active"
  account = {
    id = "f037e56e89293a057740de681ac9abbe"
  }
}
```

### `data.cloudflare_rulesets`

The `filter` block and `include_rules` attribute are removed (no v5
equivalent):

```hcl
# Before (v4)
data "cloudflare_rulesets" "example" {
  zone_id       = "0da42c8d2132a9ddaf714f9e7c920711"
  include_rules = true

  filter {
    name = "my ruleset"
  }
}

# After (v5)
data "cloudflare_rulesets" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
}
```

### `data.cloudflare_load_balancer_pools`

The `filter` block is removed (v4's regex name filtering has no v5
equivalent). The output attribute `pools` is renamed to `result`:

```hcl
# Before (v4)
data "cloudflare_load_balancer_pools" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"

  filter {
    name = ".*prod.*"
  }
}

# After (v5)
data "cloudflare_load_balancer_pools" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
}
```

Update any references from `.pools` to `.result`.

### Other data sources

Data sources that were renamed to match their resource counterparts (e.g.,
`data.cloudflare_record` to `data.cloudflare_dns_record`) simply need the
type updated in your HCL. Since data source state is ephemeral, no state
migration or `moved` blocks are needed -- Terraform will re-read the data
source on the next plan.

If a data source reference is used elsewhere in your configuration, update
those references to the new name:

```hcl
# Before
output "zone_id" {
  value = data.cloudflare_record.example.zone_id
}

# After
output "zone_id" {
  value = data.cloudflare_dns_record.example.zone_id
}
```

Data sources removed from state can be cleaned up with:

```bash
terraform state rm data.cloudflare_record.example
```

This is optional -- Terraform will handle it on the next plan/apply.

---

## Troubleshooting

### "Error: Failed to upgrade resource state"

One of the 21 stepping-stone resources was upgraded without the intermediate
step. See [recovery instructions](#what-happens-if-you-skip-the-stepping-stone)
above.

### "Error: Resource type not found" after rename

You changed the resource type in HCL but did not add a `moved` block.
Terraform does not know the new resource is the same as the old one. Add the
appropriate `moved` block. For resources marked **Manual** in the rename
table, use `terraform state rm` + `terraform import` instead.

### "Error: no schema available for cloudflare_access_policy.<name> while reading state"

This usually means Terraform is still trying to decode v4
`cloudflare_access_policy` state with a v5 provider before the migration flow
is complete.

Most common causes:

- The workspace was upgraded to v5 before applying on v4.52.5.
- Config/state rename steps were partially applied (for example, stale state
  still references `cloudflare_access_policy`).
- Application-scoped policies (`application_id` set) were treated like normal
  renamed resources.

Fix:

1. Pin to v4.52.5 and run `terraform init -upgrade && terraform apply`.
2. Run `tf-migrate` and apply generated changes.
3. For `cloudflare_access_policy` with `application_id`, keep the generated
   `removed` block and migrate policy config inline to
   `cloudflare_zero_trust_access_application.policies`.
4. Upgrade provider to v5 and run `terraform plan` again.


### Unexpected plan diff after migration

Some computed attributes may show a diff on the first plan after migration.
This is expected as Terraform reconciles the state with the API. Run
`terraform apply` to resolve.

### tf-migrate does not handle my configuration

`tf-migrate` works on static `.tf` files. If you use complex module
structures, `for_each` with dynamic resource types, or heavy use of
`templatefile`, you may need to make some HCL changes manually. Refer to the
[version 5 upgrade guide] for per-resource attribute change details.

#### tf-migrate resource rename limitations

`tf-migrate` supports resource rename post-processing, but a small set of resources use conditional routing where one v4 type can map to multiple v5 types. In these cases, automatic global reference rewriting may require manual verification.
 `cloudflare_authenticated_origin_pulls`
- `cloudflare_authenticated_origin_pulls` may migrate to:
  - `cloudflare_authenticated_origin_pulls_settings` (when `hostname` is not set)
  - `cloudflare_authenticated_origin_pulls` (when `hostname` is set)
- The target type depends on resource shape, not just the v4 type name.
 `cloudflare_authenticated_origin_pulls_certificate`
- `cloudflare_authenticated_origin_pulls_certificate` may migrate to:
  - `cloudflare_authenticated_origin_pulls_certificate` (when `type = "per-zone"`)
  - `cloudflare_authenticated_origin_pulls_hostname_certificate` (when `type = "per-hostname"`)
- The target type depends on certificate mode, not just the v4 type name.
 `cloudflare_dlp_profile` / `cloudflare_zero_trust_dlp_profile`
- Both v4 resource names share the same migration logic and may migrate to:
  - `cloudflare_zero_trust_dlp_custom_profile` (when `type = "custom"`)
  - `cloudflare_zero_trust_dlp_predefined_profile` (when `type = "predefined"`)
- The target type depends on the profile `type` value.
~> For the resources above, review generated `moved` blocks and cross-resource references before apply.

### State locked during migration

Ensure no other Terraform processes are running. If using remote state with
locking (e.g., S3 + DynamoDB, Terraform Cloud), verify the lock is released
before retrying.

---

## FAQ

**Do I still need Grit?**

No. **Grit-based migration is deprecated and will be removed in a future
release.** `tf-migrate` replaces the Grit patterns for HCL migration, and
state upgraders handle state automatically. Do not use Grit for new migrations.

**Do I need to manually edit my state file?**

No. State upgraders handle all state transformations automatically when you
run `terraform plan` or `terraform apply`.

**Why does Terraform say an object "will no longer be managed" after migration?**

This is expected when `tf-migrate` generates a `removed` block (for example,
`cloudflare_access_policy` resources with `application_id`). A `removed` block
drops Terraform state tracking without deleting the remote object
(`destroy = false`).

Apply once to clear the old state entry, then continue with the manual config
migration steps for that resource (for application-scoped access policies,
rewrite the policy as inline `policies` on
`cloudflare_zero_trust_access_application`).

**Can I skip from v4 directly to v5.19?**

Yes, as long as you are on v4.52.5 first. The state upgraders handle the full
v4 to v5 state transformation in a single step.

**Can I skip from v5.14 directly to v5.19?**

Only if you do not use any of the
[21 stepping-stone resources](#resources-requiring-stepping-stone-upgrades). If
you do, upgrade to v5.18 first.

**What if I am on Terraform < 1.8?**

You can still migrate. `moved` blocks require Terraform 1.8+, but the
provider also supports migration via `terraform state mv` for most renamed
resources. When you run `terraform state mv cloudflare_record.x
cloudflare_dns_record.x`, the provider's state upgraders automatically
transform the state on the next plan/apply. See exceptions in
[Using `terraform state mv` (Terraform < 1.8)](#using-terraform-state-mv-terraform--18)
for details.

**What about `cloudflare_worker_secret`?**

This resource has been removed in v5. Migrate to one of:

- [Secrets Store](https://developers.cloudflare.com/secrets-store/) with the
  `secrets_store_secret` binding on `cloudflare_workers_script`
- The `secret_text` binding on `cloudflare_workers_script`
- The [Workers Secrets API](https://developers.cloudflare.com/api/resources/workers/subresources/scripts/subresources/secrets/)

**What about `cloudflare_zone_settings_override`?**

This resource has been removed in v5. Each zone setting is now managed by an
individual `cloudflare_zone_setting` resource. See the
[`cloudflare_zone_settings_override`](#cloudflare_zone_settings_override)
section under Resources Requiring Manual Migration.

**What if my Access application is replaced and JWT audience (`aud`) changes?**

Some Access application migrations may require replacement depending on your
configuration and API-managed values. If replacement occurs, the application's
`aud` value changes, which can affect JWT audience checks in downstream
services.

Plan for this during migration:

- schedule a maintenance window for Access app replacements,
- update downstream audience allow-lists to include the new `aud`, and
- validate auth flows immediately after apply.

**What if I use Terraform Cloud or remote state?**

The migration works the same way. The state upgraders run when the provider
executes `terraform plan` or `terraform apply`, regardless of where state is
stored (local, S3, Terraform Cloud, etc.). Just ensure you have exclusive
access to the state during migration.

**What if I have a mixed environment (some resources still on v4)?**

You cannot use mixed provider versions in the same Terraform workspace. You
must migrate all resources in a workspace at once. If you need gradual
migration, separate resources into different workspaces and migrate them
independently.

**What if migration fails midway?**

1. Restore from your backup (version control or `.bak` files from tf-migrate)
2. Identify the failing resource from the error message
3. Check if it requires manual migration (see resource-specific sections above)
4. If you believe it's a bug, open an issue at
   https://github.com/cloudflare/tf-migrate/issues with the error details

**What if I see "No changes" but plan still shows diffs?**

Some resources may show cosmetic differences after migration that persist
across multiple applies. Common causes:

- HTML entity encoding differences (`&#39;` vs `'`)
- Timestamp formatting changes
- Computed fields that refresh on every plan

See [Expected Plan Changes After Migration](#expected-plan-changes-after-migration)
for details. If the diff is truly just cosmetic, you can safely ignore it.

---

## Additional Resources

- [Version 5 Upgrade Guide][version 5 upgrade guide] -- Per-resource attribute
  change details and manual migration notes.
- [Migrating Renamed Resources](migrating-renamed-resources) -- Detailed
  guide for the import, state file, and two-phase swap approaches.
- [tf-migrate] -- Source code and documentation for the HCL migration tool.
- [Terraform moved blocks](https://developer.hashicorp.com/terraform/language/moved) --
  HashiCorp documentation on the `moved` block syntax.

[version 5 upgrade guide]: version-5-upgrade
[tf-migrate]: https://github.com/cloudflare/tf-migrate
[migrating renamed resources]: migrating-renamed-resources
