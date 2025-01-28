---
note: This workflow is used as input for the autogen pipeline. Use it to generate workflow.ts
---

## Terraform Provider Workflow

The goal of this workflow is to migrate between two Terraform provider versions for the `Cloudflare` API.
Your job is to generate a GritQL migration that can handle upgrading between the two provider versions.

We will be upgrading to the `v5` workflow. The migration should be named `cloudflare_terraform_v5`.

### 1. Attribute mapping

The respective Terraform provider schema diffs have been dumped to `new.json` and `old.json`.

Many of the resources have had `block` attributes converted to lists. In the old schema, such attributes will appear like this:

```
"cloudflare_access_application": {
          "version": 0,
          "block": {
            ...
            },
            "block_types": {
              "cors_headers": {
                "nesting_mode": "list",
                "block": {
                  ...
                },
                "max_items": 1
              }
            },
          }
        },
```

We will want to generate a GritQL migration for each such block. The `nesting_mode` will be `list` or `set`.

There are _two_ possible destination structures. You should determine this by inspecting the equivalent attribute in a `new_schema_file`.

If the new nested attribute is `"nesting_mode": "set"` or `"nesting_mode": "list"`, then the block should be converted to a list, like this:

```grit
language hcl

inline_cloudflare_block_to_list(`cors_headers`) as $block where { $block <: within `resource "cloudflare_access_application" $_ { $_ }` }
```

If the new nested attribute is `"nesting_mode": "single"`, then the block should be converted to a map, like this:

```grit
language hcl

inline_cloudflare_block_to_map(`cors_headers`) as $block where { $block <: within `resource "cloudflare_access_application" $_ { $_ }` }
```

Make sure to look recursively for _all_ blocks in the schema. Eliminate all duplicates. We should use `any` to combine all the blocks into a single migration.

You must carefully look inside the `new` schema to find the correct nesting mode for each block. If the attribute can't be found in the new schema, this is an error and your recursion logic is wrong.
