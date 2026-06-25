---
name: get-provider-resources
description: Lists Cloudflare Terraform provider resources, data sources, and guides from the public Registry. Use when the user is looking for a `cloudflare_*` resource, doesn't remember its exact name, or needs to map a Cloudflare product to a Terraform resource.
---

# Cloudflare Terraform — List Provider Resources, Data Sources, and Guides

Your pre-trained knowledge of which Cloudflare resources exist is likely stale — resources are added, renamed, and consolidated across provider versions. **Prefer retrieval over pre-training.**

Provider repository: https://github.com/cloudflare/terraform-provider-cloudflare

## Retrieval Sources

**Registry JSON API (primary)**
- `https://registry.terraform.io/v1/providers/cloudflare/cloudflare`
- Returns the full `docs` array — every resource, data source, guide, and overview page.

**Pinned-version index**
- `https://registry.terraform.io/v1/providers/cloudflare/cloudflare/{version}`
- Same shape, scoped to a specific version. Use when the user is pinned to an older provider — substitute their version (e.g. `5.15.0`) into the URL.
- To discover the user's pinned version, check their `required_providers` block (in their `*.tf` files) for the version constraint, or `.terraform.lock.hcl` for the exact resolved version.

**GitHub repo `docs/` (fallback)**
- `https://github.com/cloudflare/terraform-provider-cloudflare/tree/main/docs`
- Subdirectories: `resources/`, `data-sources/`, `guides/`.
- Same content as the Registry, browseable when the Registry is unreachable.
- List files via the GitHub API: `https://api.github.com/repos/cloudflare/terraform-provider-cloudflare/contents/docs/resources`.
- Use the `main` branch (not `master` — `master` URLs will 404).

Do **not** fetch `https://registry.terraform.io/providers/cloudflare/cloudflare/latest/...` URLs — they serve a JavaScript-gated SPA and return an empty shell to non-browser clients.

## Procedure

1. GET the registry index.
2. Parse the JSON. Each `docs` entry has `{ id, title, path, slug, category, subcategory, language }`.
3. Filter by `category`:
   - `"resources"` — declarable, manageable resources (e.g. `cloudflare_dns_record`)
   - `"data-sources"` — read-only lookups (e.g. `data.cloudflare_zone`)
   - `"guides"` — long-form contributor docs (migration guides, upgrade guides, conventions)
   - `"overview"` — provider-level overview page (auth, getting started)
4. Match against `title` values. Titles omit the `cloudflare_` prefix — reattach it when presenting to the user.
5. Return at most 5-10 matches sorted by relevance. For broad queries ("what storage resources exist?"), summarize by namespace prefix instead of listing every entry.

## Resource discovery workflow

When the user asks "is there a Cloudflare resource for X?" and the obvious `title` search returns nothing, **don't conclude the resource doesn't exist** — Cloudflare consolidates many features into single resources keyed by a discriminator. Follow this procedure before answering "no":

1. **Direct search.** Look for `title` substrings matching the user's intent (e.g. "rate limiting" → search `rate`, `limit`).
2. **Consolidation check.** If no direct match, try the known consolidation points:
   - **WAF / rate limiting / redirects / caching / transforms** → `cloudflare_ruleset` with the matching `phase` (~25 phases — fetch the `ruleset` doc to see them all)
   - **Access apps / policies / groups / identity providers / tunnels** → `cloudflare_zero_trust_*` namespace
   - **Magic WAN tunnels / static routes** → `cloudflare_magic_wan_*` namespace
   - **Workers (script / route / custom domain / KV / cron)** → `cloudflare_workers_*` namespace (plural, not `worker_*` — that's v4)
   - **Zone-level settings** → `cloudflare_zone_setting` (singular, one resource per setting; `cloudflare_zone_settings_override` is v4 only)
3. **Migration check.** If the user named a resource that doesn't appear (e.g. `cloudflare_record`, `cloudflare_access_application`), it was likely renamed in v5. Fetch `version-5-migration` (in the `guides` category) and consult the rename table.
4. **Only then** conclude the resource doesn't exist. Be explicit that you checked direct names, consolidation points, and rename tables.

## The `guides` category is high-value

Whenever the user is doing anything beyond a single-resource question — upgrading provider versions, migrating from v4, troubleshooting "this used to work", or asking about provider auth — **scan the `guides` category first**. The guides cover material that isn't in any individual resource doc: cross-resource migration paths, renamed resources, the official `tf-migrate` CLI, and known-issue rescue procedures.

Two key guides serve different purposes and are easy to confuse:

- **`version-5-migration`** — the end-to-end migration workflow (state upgraders, `moved` blocks, `tf-migrate` CLI, the `Auto`/`Manual` resource-rename table, and the **known-issues / perpetual-diff table**). Start here for "how do I migrate?" and "why does my plan keep showing the same diff?"
- **`version-5-upgrade`** — per-resource attribute change reference. Start here for "what changed about resource X between v4 and v5?"

## Example

> **User:** "Is there a Cloudflare resource for WAF custom rules?"
>
> Filter resources, search for `waf` and `rule`. Closest match: `ruleset` (no dedicated `waf_custom_rule`).
>
> Fetch the `ruleset` doc from the Registry to confirm the `phase` value, then reply:
> > WAF custom rules use **`cloudflare_ruleset`** with `phase = "http_request_firewall_custom"`. The `ruleset` resource is consolidated — it also handles rate limiting, redirects, caching, and transforms, keyed by `phase`.
