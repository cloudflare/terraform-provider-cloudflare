---
page_title: "cloudflare_zero_trust_dlp_dataset Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_dlp_dataset (Resource)



## Example Usage

```terraform
resource "cloudflare_zero_trust_dlp_dataset" "example_zero_trust_dlp_dataset" {
  account_id = "account_id"
  name = "name"
  description = "description"
  encoding_version = 0
  secret = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)
- `name` (String)

### Optional

- `dataset_id` (String)
- `description` (String) The description of the dataset
- `encoding_version` (Number) Dataset encoding version

Non-secret custom word lists with no header are always version 1.
Secret EDM lists with no header are version 1.
Multicolumn CSV with headers are version 2.
Omitting this field provides the default value 0, which is interpreted
the same as 1.
- `secret` (Boolean) Generate a secret dataset.

If true, the response will include a secret to use with the EDM encoder.
If false, the response has no secret and the dataset is uploaded in plaintext.

### Read-Only

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--columns))
- `created_at` (String)
- `dataset` (Attributes) (see [below for nested schema](#nestedatt--dataset))
- `id` (String) The ID of this resource.
- `max_cells` (Number)
- `num_cells` (Number)
- `status` (String)
- `updated_at` (String) When the dataset was last updated.

This includes name or description changes as well as uploads.
- `uploads` (Attributes List) (see [below for nested schema](#nestedatt--uploads))
- `version` (Number) The version to use when uploading the dataset.

<a id="nestedatt--columns"></a>
### Nested Schema for `columns`

Read-Only:

- `entry_id` (String)
- `header_name` (String)
- `num_cells` (Number)
- `upload_status` (String)


<a id="nestedatt--dataset"></a>
### Nested Schema for `dataset`

Read-Only:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--dataset--columns))
- `created_at` (String)
- `description` (String) The description of the dataset
- `encoding_version` (Number)
- `id` (String)
- `name` (String)
- `num_cells` (Number)
- `secret` (Boolean)
- `status` (String)
- `updated_at` (String) When the dataset was last updated.

This includes name or description changes as well as uploads.
- `uploads` (Attributes List) (see [below for nested schema](#nestedatt--dataset--uploads))

<a id="nestedatt--dataset--columns"></a>
### Nested Schema for `dataset.columns`

Read-Only:

- `entry_id` (String)
- `header_name` (String)
- `num_cells` (Number)
- `upload_status` (String)


<a id="nestedatt--dataset--uploads"></a>
### Nested Schema for `dataset.uploads`

Read-Only:

- `num_cells` (Number)
- `status` (String)
- `version` (Number)



<a id="nestedatt--uploads"></a>
### Nested Schema for `uploads`

Read-Only:

- `num_cells` (Number)
- `status` (String)
- `version` (Number)

