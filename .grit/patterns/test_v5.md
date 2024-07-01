# Test Cloudflare v5

```grit
language hcl

terraform_cloudflare_v5()
```

## test: basic rewrite

```hcl
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  require {
    any_valid_service_token = true
  }
}
```

```hcl
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  require =[{
    any_valid_service_token = true
  }]
}
```

## test: list collapsing

Multiple blocks should be collapsed into a single list attribute.

```hcl
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  include {
    email = ["test@example.com"]
  }

  include {
    email = ["someone@example.com"]
  }

  exclude {
    email = ["bad@other.com"]
  }
}
```

```hcl
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  include = [{
    email = ["test@example.com"]
  },
  {
    email = ["someone@example.com"]
  }]


  exclude = [{
    email = ["bad@other.com"]
  }]
}
```

## test: nested blocks

Nested blocks must also be rewritten.

```hcl
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  require {
    azure {
      id = ["1234"]
    }
  }
}
```

```hcl
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  require = [{
    azure = [{
      id = ["1234"]
    }]
  }]
}
```
