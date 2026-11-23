// This file is intentionally NOT code-generated.

package snippet

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The content endpoint names each multipart part after its file, as in
// `Content-Disposition: form-data; name="worker.js"; filename="worker.js"`.
// Read must keep those parts, or the state never holds the code and a change made
// outside Terraform never shows in a plan.
func TestUnmarshalMultipartKeepsPartsNamedAfterTheirFile(t *testing.T) {
	t.Parallel()

	const boundary = "snippet-boundary"
	body := "--" + boundary + "\r\n" +
		"Content-Disposition: form-data; name=\"main.js\"; filename=\"main.js\"\r\n" +
		"Content-Type: application/javascript+module\r\n\r\n" +
		"import { greet } from \"./helper.js\"\nexport default { async fetch(request) { return new Response(greet()) } }\r\n" +
		"--" + boundary + "\r\n" +
		"Content-Disposition: form-data; name=\"helper.js\"; filename=\"helper.js\"\r\n" +
		"Content-Type: application/javascript+module\r\n\r\n" +
		"export const greet = () => \"hello\"\r\n" +
		"--" + boundary + "--\r\n"

	var got SnippetModel
	if err := got.UnmarshalMultipart([]byte(body), "multipart/form-data; boundary="+boundary); err != nil {
		t.Fatalf("UnmarshalMultipart: %v", err)
	}
	if got.Files == nil || len(*got.Files) != 2 {
		t.Fatalf("files = %v, want two files", got.Files)
	}
	want := []struct{ name, content string }{
		{"main.js", "import { greet } from \"./helper.js\"\nexport default { async fetch(request) { return new Response(greet()) } }"},
		{"helper.js", "export const greet = () => \"hello\""},
	}
	for i, w := range want {
		attrs := (*got.Files)[i].Attributes()
		if !attrs["name"].Equal(types.StringValue(w.name)) || !attrs["content"].Equal(types.StringValue(w.content)) {
			t.Errorf("file %d = %v, want name %q with content %q", i, attrs, w.name, w.content)
		}
	}
}

// A part without a file name, such as a JSON metadata field, is not a file.
func TestUnmarshalMultipartSkipsPartsWithoutAFileName(t *testing.T) {
	t.Parallel()

	const boundary = "snippet-boundary"
	body := "--" + boundary + "\r\n" +
		"Content-Disposition: form-data; name=\"metadata\"\r\n\r\n" +
		"{\"main_module\":\"main.js\"}\r\n" +
		"--" + boundary + "\r\n" +
		"Content-Disposition: form-data; name=\"main.js\"; filename=\"main.js\"\r\n\r\n" +
		"export default {}\r\n" +
		"--" + boundary + "--\r\n"

	var got SnippetModel
	if err := got.UnmarshalMultipart([]byte(body), "multipart/form-data; boundary="+boundary); err != nil {
		t.Fatalf("UnmarshalMultipart: %v", err)
	}
	if got.Files == nil || len(*got.Files) != 1 || (*got.Files)[0].Name() != "main.js" {
		t.Fatalf("files = %v, want main.js only", got.Files)
	}
}
