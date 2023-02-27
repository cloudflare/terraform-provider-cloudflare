package utils

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

// GenerateRandomResourceName builds a unique-ish resource identifier to use in
// tests.
func GenerateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}
