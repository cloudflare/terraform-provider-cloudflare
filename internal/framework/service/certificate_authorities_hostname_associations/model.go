package certificate_authorities_hostname_associations

import "github.com/hashicorp/terraform-plugin-framework/types"

type HostnameAssociation = types.String

type CertificateAuthoritiesHostnameAssociationsModel struct {
	ZoneID            types.String          `tfsdk:"zone_id"`
	MTLSCertificateID types.String          `tfsdk:"mtls_certificate_id"`
	Hostnames         []HostnameAssociation `tfsdk:"hostnames"`
}
