---
name: check-deprecation
description: Checks whether a Cloudflare Terraform resource, data source, or attribute is deprecated. Use ONLY when deprecation, removal, replacement, or a Terraform deprecation warning is the main question. Biases towards Terraform Registry retrieval over pre-trained knowledge.
---

# Cloudflare Terraform - Check Deprecation

Deprecation markers change often. Prefer retrieved docs over pre-trained knowledge.

## Retrieval Sources

| Source | URL | Returns |
|--------|-----|---------|
| Registry JSON API - index (primary) | `https://registry.terraform.io/v1/providers/cloudflare/cloudflare` | Docs index - find the exact `path` by `title` + `category`, e.g. `docs/resources/dns_record.md` |
| Registry JSON API - docs (primary) | `https://registry.terraform.io/v1/providers/cloudflare/cloudflare/{version}/docs?path={url-encoded-path}` | Full markdown |
| GitHub raw docs (fallback) | `https://raw.githubusercontent.com/cloudflare/terraform-provider-cloudflare/{tag-or-main}/docs/{category}/{title}.md` | Use matching tag for pinned versions; use `main` only for latest/unpinned checks |
| Local provider repo (fallback when available) | `internal/services/<service>/*schema*.go` | Framework `DeprecationMessage` values that may not appear in generated docs |
| Cloudflare API deprecations (replacement context only) | `https://developers.cloudflare.com/fundamentals/api/reference/deprecations/` | Dated log of deprecated Cloudflare APIs/fields with replacement products. Use only for replacement guidance after the Terraform doc confirms deprecation — never to pick Terraform attribute names. |
| Cloudflare developer docs (replacement context only) | Whichever `developers.cloudflare.com/*` URL the deprecation description links to | Product docs for the replacement service. Same scoping as above. |

Do **not** fetch `https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/...`; it is a JavaScript SPA shell for non-browser clients.

## Procedure

1. **Identify** whether the target is a resource, data source, or unknown, plus the optional attribute name.
2. **Determine provider version** from user context when available: `.terraform.lock.hcl`, then `required_providers`, otherwise latest Registry version.
3. **Normalize** the name: parse addresses like `data.cloudflare_zone.example.permissions` into category `data-sources`, type `zone`, attribute `permissions`; strip `data.cloudflare_` or `cloudflare_`, lowercase, keep only `[a-z0-9_]`.
4. **Find docs path** in the Registry index. If category is known, search only `resources` or `data-sources`; otherwise check both and disambiguate if both exist.
5. **Fetch docs** from the Registry JSON API with the exact URL-encoded `path` from the index (for example `docs%2Fresources%2Fdns_record.md`) or use the GitHub fallback if the Registry is unreachable/truncated.
6. **If the target isn't in the docs index**, fetch `version-5-migration` / `version-5-upgrade` and quote the relevant rename/removal entry if present; otherwise say no replacement was found in docs.
7. **For whole-resource or whole-data-source questions**, scan page prose outside code blocks for `\bdeprecated\b` before checking attributes.
8. **For attribute questions**, scan the full `## Schema` attribute block for the named attribute, including nested schema sections and continuation lines, using both patterns below.
9. **If running inside this provider repo**, also check local schema files for `DeprecationMessage` before saying something is not deprecated. Some schema-level deprecations may be present in Go but missing from generated docs.
10. **If the user named a specific attribute** that doesn't exist in the schema, say so. The user may have a typo, wrong category, or wrong provider version.
11. **If the user asks what to use instead** (not just "is this deprecated?"), and only after the Terraform doc confirms deprecation: fetch any `developers.cloudflare.com/*` URL in the deprecation description, and grep the [API deprecations page](https://developers.cloudflare.com/fundamentals/api/reference/deprecations/) for the attribute or API name. Report the deprecation and end-of-life dates when present. If the replacement is another Terraform resource, confirm it in the Registry before naming it.
12. **Be explicit about what you found.** If nothing is deprecated, say so.

## How deprecation is marked

Scan for both explicit type markers and description-only field/block deprecations:

```
- `old_attribute` (String, Deprecated) Use new_attribute instead.
- `legacy_field` (String) This field is deprecated, use modern_field instead.
```

Match `\bdeprecated\b` (word-boundary, case-insensitive) in the attribute block so words like "undeprecated" don't false-match. For attributes, scan only schema attribute blocks - not code blocks, examples, or unrelated guide prose. An attribute block starts at a bullet like ``- `name` (...)`` and continues until the next attribute bullet or heading.

If text says only an enum value is deprecated, report that distinction. Example: `Note: UNKNOWN is deprecated` means the `UNKNOWN` value is deprecated, not necessarily the whole attribute.

For whole-resource or whole-data-source questions, scan non-code page prose separately and report page-level deprecation text only if it clearly refers to the resource/data source itself, not just an attribute inside the schema.

When you find a deprecated attribute, quote its description verbatim. Do not infer replacement resources or attributes unless the docs explicitly name them.

Answer deprecation status only. Do not suggest HCL edits, `moved` blocks, imports, or state commands unless fetched docs explicitly provide that guidance.

## Example

> **User:** "I'm getting a deprecation warning on `notification_email` for `cloudflare_load_balancer_pool`. What should I use instead?"
>
> Fetch the `load_balancer_pool` resource doc. The `## Schema` block shows `notification_email` is deprecated and points at `https://developers.cloudflare.com/fundamentals/notifications/`. Fetch that page for replacement context, and grep the API deprecations page for the 2023-04-03 date. Reply:
>
> > `notification_email` was deprecated on 2023-04-03. Cloudflare moved health-check notifications to their centralized Notifications service ([docs](https://developers.cloudflare.com/fundamentals/notifications/)) — configure notifications there instead. There is no drop-in Terraform attribute replacement on `cloudflare_load_balancer_pool`; the migration is a workflow change, not a rename.
