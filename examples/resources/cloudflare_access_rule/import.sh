# User level access rule import.
$ terraform import cloudflare_access_rule.default user/<user_id>/<rule_id>

# Zone level access rule import.
$ terraform import cloudflare_access_rule.default zone/<zone_id>/<rule_id>

# Account level access rule import.
$ terraform import cloudflare_access_rule.default account/<account_id>/<rule_id>
