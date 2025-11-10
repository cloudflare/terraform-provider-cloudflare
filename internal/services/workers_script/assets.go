package workers_script

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AssetUploadSessionRequestBody struct {
	Manifest AssetManifest `json:"manifest"`
}

type AssetManifest map[string]AssetManifestEntry

type AssetManifestEntry struct {
	Filepath string `json:"-"`
	Hash     string `json:"hash"`
	Size     int64  `json:"size"`
}

type Bucket []AssetManifestEntry

func (b Bucket) MarshalMultipart() (data []byte, formDataContentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)

	for _, entry := range b {
		reader, err := os.Open(entry.Filepath)
		if err != nil {
			return nil, "", err
		}
		defer reader.Close()

		contentType := mime.TypeByExtension(filepath.Ext(entry.Filepath))

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, entry.Hash, entry.Filepath))
		h.Set("Content-Type", contentType)
		filewriter, err := writer.CreatePart(h)
		if err != nil {
			return nil, "", err
		}

		// Stream base64 encoding directly to the form field
		encoder := base64.NewEncoder(base64.StdEncoding, filewriter)
		_, err = io.Copy(encoder, reader)
		if err != nil {
			return nil, "", err
		}

		err = encoder.Close()
		if err != nil {
			return nil, "", err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

func getAssetManifest(directory string) (AssetManifest, error) {
	// Convert to absolute path to handle relative paths properly
	absBasePath, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}

	manifest := make(AssetManifest)

	// Scan directory and generate manifest
	err = filepath.WalkDir(absBasePath, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Calculate relative path from the base directory
		relPath, err := filepath.Rel(absBasePath, filePath)
		if err != nil {
			return fmt.Errorf("failed to calculate relative path for %s: %w", filePath, err)
		}

		// Normalize path separators for consistent keys
		relPath = filepath.ToSlash(relPath)

		// Add leading slash
		relPath = fmt.Sprintf("/%s", relPath)

		// Calculate SHA256 hash
		hash, err := calculateFileHash(filePath)
		if err != nil {
			return fmt.Errorf("failed to calculate hash for %s: %w", filePath, err)
		}

		// Get file info for size
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %w", filePath, err)
		}

		manifest[relPath] = AssetManifestEntry{
			Filepath: filePath,
			Hash:     hash[:32],
			Size:     info.Size(),
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return manifest, nil
}

func getAssetManifestHash(manifest AssetManifest) (string, error) {
	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		return "", err
	}

	return calculateStringHash(string(manifestBytes))
}

func handleAssets(ctx context.Context, client *cloudflare.Client, data *WorkersScriptModel) error {
	// TODO: Asset handling is not currently implemented
	// The Assets.Directory field is not exposed in the WorkersScriptMetadataAssetsModel
	// This function needs to be reimplemented once the schema is updated
	return nil
}

// =========================== Plan modifiers ===========================

func ComputeSHA256HashOfAssetManifest() planmodifier.String {
	return computeSHA256HashOfAssetManifestModifier{}
}

var _ planmodifier.String = &computeSHA256HashOfAssetManifestModifier{}

type computeSHA256HashOfAssetManifestModifier struct{}

func (c computeSHA256HashOfAssetManifestModifier) Description(_ context.Context) string {
	return "Calculates the SHA-256 hash of the manifest of asset files in the specified directory."
}

func (c computeSHA256HashOfAssetManifestModifier) MarkdownDescription(ctx context.Context) string {
	return c.Description(ctx)
}

func (c computeSHA256HashOfAssetManifestModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Don't modify during destroy
	if req.Config.Raw.IsNull() {
		return
	}

	directoryPath := req.Path.ParentPath().AtName("directory")

	var directory types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, directoryPath, &directory)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if directory.IsNull() || directory.IsUnknown() {
		return
	}

	manifest, err := getAssetManifest(directory.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Error reading asset files", err.Error())
		return
	}

	manifestHash, err := getAssetManifestHash(manifest)
	if err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Error computing SHA-256 hash of asset manifest", err.Error())
		return
	}

	resp.PlanValue = types.StringValue(manifestHash)
}
