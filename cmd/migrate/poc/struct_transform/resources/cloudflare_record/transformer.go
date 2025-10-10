package cloudflare_record

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/base"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// NewDNSRecordStructTransformer creates a complete DNS record transformer
// that implements the ResourceTransformer interface using struct-based approach
func NewDNSRecordStructTransformer() interfaces.ResourceTransformer {
	return base.NewBaseStructTransformer(
		"cloudflare_record", // v4 resource type
		NewHCLParser(),               // Parser: HCL -> v4 model
		NewDNSRecordTransformer(),    // Transformer: v4 -> v5 model
		NewDNSRecordGenerator(),      // Generator: v5 model -> HCL
	)
}

// DNSRecordV5StructTransformer handles already-migrated v5 DNS records
type DNSRecordV5StructTransformer struct {
	*base.BaseStructTransformer
}

// NewDNSRecordV5StructTransformer creates a transformer for v5 DNS records
func NewDNSRecordV5StructTransformer() interfaces.ResourceTransformer {
	// For v5 records, we just pass through without transformation
	return &DNSRecordV5StructTransformer{
		BaseStructTransformer: base.NewBaseStructTransformer(
			"cloudflare_dns_record",
			nil,
			nil,
			nil,
		),
	}
}

// TransformConfig for v5 records returns the block unchanged
func (t *DNSRecordV5StructTransformer) TransformConfig(block *hclwrite.Block) (*interfaces.TransformResult, error) {
	// Already v5, no transformation needed
	return &interfaces.TransformResult{
		Blocks:         []*hclwrite.Block{block},
		RemoveOriginal: false,
	}, nil
}