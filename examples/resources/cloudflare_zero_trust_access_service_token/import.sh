# If you are importing an Access Service Token you will not have the
# client_secret available in the state for use. The client_secret is only
# available once, at creation. In most cases, it is better to just create a new
# resource should you need to reference it in other resources.
$ terraform import cloudflare_access_service_token.example <account_id>/<service_token_id>
