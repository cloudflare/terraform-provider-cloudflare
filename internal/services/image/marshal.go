package image

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

// marshalCreateMultipart builds the multipart/form-data body for the upload
// endpoint (POST /images/v1). Localized replacement for the generated
// MarshalMultipart, which emits an empty JSON part for unset `metadata` and
// breaks create with API code 5400 (Bug A). Optional null/unknown fields are
// omitted; only image_basic_upload fields are written.
func (r ImageModel) marshalCreateMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)

	fail := func(e error) ([]byte, string, error) {
		if c := writer.Close(); c != nil {
			e = errors.Join(e, c)
		}
		return nil, "", e
	}

	// id is Required, so always known.
	if err = writer.WriteField("id", r.ID.ValueString()); err != nil {
		return fail(err)
	}

	// file is base64 (filebase64()); decode and send as a file part with an
	// image content-type. Writing it as a plain field caused 415/5455 (Bug B,
	// issue #6176).
	if !r.File.IsNull() && !r.File.IsUnknown() {
		raw, decodeErr := base64.StdEncoding.DecodeString(r.File.ValueString())
		if decodeErr != nil {
			return fail(fmt.Errorf("`file` must be base64-encoded image data, e.g. filebase64(\"logo.png\"): %w", decodeErr))
		}

		contentType := detectImageContentType(raw)
		filename := "upload" + extensionForImageContentType(contentType)

		header := textproto.MIMEHeader{}
		header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
		header.Set("Content-Type", contentType)

		var part io.Writer
		part, err = writer.CreatePart(header)
		if err != nil {
			return fail(err)
		}
		if _, err = part.Write(raw); err != nil {
			return fail(err)
		}
	}

	if !r.URL.IsNull() && !r.URL.IsUnknown() {
		if err = writer.WriteField("url", r.URL.ValueString()); err != nil {
			return fail(err)
		}
	}

	if !r.Creator.IsNull() && !r.Creator.IsUnknown() {
		if err = writer.WriteField("creator", r.Creator.ValueString()); err != nil {
			return fail(err)
		}
	}

	if !r.RequireSignedURLs.IsNull() && !r.RequireSignedURLs.IsUnknown() {
		value := "false"
		if r.RequireSignedURLs.ValueBool() {
			value = "true"
		}
		if err = writer.WriteField("requireSignedURLs", value); err != nil {
			return fail(err)
		}
	}

	// metadata: written only when set, as a JSON part. Omitting it when null
	// is the Bug A fix.
	if !r.Metadata.IsNull() && !r.Metadata.IsUnknown() {
		header := textproto.MIMEHeader{}
		header.Set("Content-Disposition", `form-data; name="metadata"`)
		header.Set("Content-Type", "application/json")

		var part io.Writer
		part, err = writer.CreatePart(header)
		if err != nil {
			return fail(err)
		}
		if _, err = part.Write([]byte(r.Metadata.ValueString())); err != nil {
			return fail(err)
		}
	}

	if err = writer.Close(); err != nil {
		return nil, "", err
	}

	return buf.Bytes(), writer.FormDataContentType(), nil
}

// marshalEditJSON builds the application/json body for the edit endpoint
// (PATCH /images/v1/{image_id}). The generated Update sent multipart to this
// JSON endpoint, causing API code 5400 (Bug C). Only edit fields (creator,
// requireSignedURLs, metadata) are emitted; null/unknown are omitted so unset
// values are left unchanged.
func (r ImageModel) marshalEditJSON() ([]byte, error) {
	body := map[string]any{}

	if !r.Creator.IsNull() && !r.Creator.IsUnknown() {
		body["creator"] = r.Creator.ValueString()
	}

	if !r.RequireSignedURLs.IsNull() && !r.RequireSignedURLs.IsUnknown() {
		body["requireSignedURLs"] = r.RequireSignedURLs.ValueBool()
	}

	// metadata is a JSON object; embed as raw JSON, not a quoted string.
	if !r.Metadata.IsNull() && !r.Metadata.IsUnknown() {
		body["metadata"] = json.RawMessage(r.Metadata.ValueString())
	}

	return json.Marshal(body)
}

// detectImageContentType returns the image MIME type for raw bytes.
// http.DetectContentType handles png/jpeg/gif/webp; SVG needs a heuristic.
func detectImageContentType(raw []byte) string {
	trimmed := bytes.TrimSpace(raw)
	if bytes.HasPrefix(trimmed, []byte("<svg")) ||
		(bytes.HasPrefix(trimmed, []byte("<?xml")) && bytes.Contains(trimmed, []byte("<svg"))) {
		return "image/svg+xml"
	}

	contentType := http.DetectContentType(raw)
	// Drop params like "; charset=utf-8".
	if i := strings.IndexByte(contentType, ';'); i >= 0 {
		contentType = strings.TrimSpace(contentType[:i])
	}
	return contentType
}

// extensionForImageContentType returns a cosmetic filename extension.
func extensionForImageContentType(contentType string) string {
	switch contentType {
	case "image/png":
		return ".png"
	case "image/jpeg":
		return ".jpg"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/svg+xml":
		return ".svg"
	case "image/heic":
		return ".heic"
	default:
		return ""
	}
}
