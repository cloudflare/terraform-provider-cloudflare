---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_bookmark"
description: Provides a Cloudflare Access Bookmark resource.
---

# cloudflare_access_bookmark

Provides a Cloudflare Access Bookmark resource. Access Bookmark
applications are not protected behind Access but are displayed in the App
Launcher.

## Example Usage

```hcl
resource "cloudflare_access_bookmark" "my_bookmark_app" {
  account_id           = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name                 = "My Bookmark App"
  domain               = "example.com"
  logo_url             = "https://path-to-logo.com/example.png"
  app_launcher_visible = true
}
```

## Argument Reference

The following arguments are supported:

-> **Note:** It's required that an `account_id` or `zone_id` is provided and in most cases using either is fine. However, if you're using a scoped access token, you must provide the argument that matches the token's scope. For example, an access token that is scoped to the "example.com" zone needs to use the `zone_id` argument.

- `account_id` - (Optional) The account to which the Access bookmark application should be added. Conflicts with `zone_id`.
- `zone_id` - (Optional) The DNS zone to which the Access bookmark application should be added. Conflicts with `account_id`.
- `name` - (Required) Name of the bookmark application.
- `domain` - (Required) The domain of the bookmark application. Can include subdomains, paths, or both.
- `logo_url` - (Optional) The image URL for the logo shown in the app
  launcher dashboard.
- `app_launcher_visible` - (Optional) Option to show/hide the bookmark in the app launcher. Defaults to `true`.

## Attributes Reference

The following additional attributes are exported:

- `id` - ID of the bookmark application
- `name` - Name of the bookmark application
- `domain` - Domain of the bookmark application
- `logo_url` - Logo URL of the bookmark application
- `app_launcher_visible` - The visibility status of the bookmark in app
  launcher.

## Import

Access Bookmarks can be imported using a composite ID formed of account
ID and bookmark ID.

```
$ terraform import cloudflare_access_bookmark.my_bookmark cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
