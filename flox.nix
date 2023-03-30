{
  # Environment variables that will allow you to setup and run acceptance tests.
  #
  # environmentVariables.CLOUDFLARE_EMAIL = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_API_KEY = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_API_TOKEN = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_DOMAIN = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_ZONE_ID = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_ACCOUNT_ID = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_ALT_ZONE_ID = "REPLACE_ME";
  # environmentVariables.CLOUDFLARE_ALT_DOMAIN = "REPLACE_ME";
  #
  # Configure Terraform debug output for operations.
  # environmentVariables.TF_LOG = "DEBUG|TRACE";
  #
  # Use a MITM proxy to view the HTTP interactions.
  # environmentVariables.HTTPS_PROXY = "REPLACE_ME";
  # environmentVariables.HTTP_PROXY = "REPLACE_ME";

  packages.nixpkgs-flox.go_1_20 = { };
}
