package workers_script

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func writeFileBytes(partName string, filename string, contentType string, content io.Reader, writer *multipart.Writer) error {
	h := make(textproto.MIMEHeader)
	header := "form-data"

	if escapeQuotes(partName) != "" {
		header += fmt.Sprintf(`; name="%s"`, escapeQuotes(partName))
	}

	if escapeQuotes(filename) != "" {
		header += fmt.Sprintf(`; filename="%s"`, escapeQuotes(filename))
	}

	h.Set("Content-Disposition", header)
	h.Set("Content-Type", contentType)
	filewriter, err := writer.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = io.Copy(filewriter, content)
	return err
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func readFile(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		dirname, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("could not expand home directory in path %s: %w", path, err)
		}
		path = filepath.Join(dirname, path[2:])
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read file %s: %w", path, err)
	}

	return string(content), nil
}

func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func calculateStringHash(content string) (string, error) {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:]), nil
}

var _ validator.String = &contentSHA256Validator{}

type contentSHA256Validator struct {
	ContentPath     string
	ContentFilePath string
}

func (v contentSHA256Validator) Description(_ context.Context) string {
	return fmt.Sprintf("Validates that the provided value matches the SHA-256 hash of content in either `content` or `content_file`.")
}

func (v contentSHA256Validator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v contentSHA256Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	providedHash := req.ConfigValue.ValueString()

	var config WorkersScriptModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var hasContent, hasContentFile bool

	if !config.Content.IsNull() {
		hasContent = true
	}

	if !config.ContentFile.IsNull() {
		hasContentFile = true
	}

	if !hasContent && !hasContentFile {
		resp.Diagnostics.AddError("Missing required attributes", "One of `content` or `content_file` is required")
		return
	}

	var actualHash string
	var err error

	if hasContent {
		actualHash, err = calculateStringHash(config.Content.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Hash Calculation Error",
				fmt.Sprintf("Failed to calculate SHA-256 hash of content: %s", err.Error()),
			)
			return
		}
	} else if hasContentFile {
		actualHash, err = calculateFileHash(config.ContentFile.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Hash Calculation Error",
				fmt.Sprintf("Failed to calculate SHA-256 hash of file '%s': %s", config.ContentFile.ValueString(), err.Error()),
			)
			return
		}
	}

	if providedHash != actualHash {
		var source string
		if hasContent {
			source = "content"
		} else if hasContentFile {
			source = fmt.Sprintf("content_file (%s)", config.ContentFile.ValueString())
		}

		resp.Diagnostics.AddAttributeError(
			req.Path,
			"SHA-256 Hash Mismatch",
			fmt.Sprintf("The provided SHA-256 hash '%s' does not match the actual hash '%s' of %s",
				providedHash, actualHash, source),
		)
	}
}

func ValidateContentSHA256() validator.String {
	return contentSHA256Validator{
		ContentPath:     "content",
		ContentFilePath: "content_file",
	}
}
