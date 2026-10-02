resource "cloudflare_zero_trust_dlp_predefined_profile" "example_zero_trust_dlp_predefined_profile"   {
  profile_id = "e91a2360-da51-4fdf-9711-bcdecd462614"
  account_id = "account_id"
  ocr_enabled = true
  // Entries in this predefined profile we want to enable. Any entries not included will be disabled
  enabled_entries = [
    "56a8c060-01bb-4f89-ba1e-3ad42770a342",
    "7f575e6d-039a-465e-85cf-175bda88d4f2",
    "03ebabfd-ce7e-45ed-8061-65e28f0a6e53",
    "2d9c356d-b5a3-482a-b01e-0363e0de7458",
    "2f3657af-c39b-4899-9a98-22f7d187dd28",
    "753a16f9-f533-4208-a5b8-6319b201e9fb",
    "ebcea2c4-335a-457c-853b-f7ae7cc74e07",
    "3f5c4c83-f34c-4d17-81c7-3028385737b3",
    "d1a84fde-c375-4d3c-8a27-8c4eaa33cf60",
    "6dbe5604-d3a3-4c3e-905c-57985704bea7",
    "55ba2c6c-8ef4-4b2e-9148-e75e8b6ccac1",
    "5b1d5035-8c53-4bc9-a151-404eb32b34b4",
    "acf28d88-2daf-4bc4-aa36-5ac1fac0540a"
  ]
}
