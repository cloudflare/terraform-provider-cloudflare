package certificate_pack

import "fmt"

const certificatePackStatusDeletionTimedOut = "deletion_timed_out"

func isDeletedCertificatePackStatus(status string) bool {
	switch status {
	case "deleted", "pending_deletion", certificatePackStatusDeletionTimedOut:
		return true
	default:
		return false
	}
}

// deletedCertificatePackWarning returns the summary and detail of the warning
// Read adds when it drops a deleted pack from state. A pack in
// deletion_timed_out never finished deleting, so it may still exist.
func deletedCertificatePackWarning(zoneID, id, status string) (string, string) {
	if status != certificatePackStatusDeletionTimedOut {
		return "Resource deleted", "The resource is deleted on the server and will be removed from state."
	}
	return "Certificate pack deletion timed out",
		fmt.Sprintf("Cloudflare reports certificate pack %[2]s in zone %[1]s with status %[3]q. "+
			"Its deletion timed out on Cloudflare's side, so the pack may still exist. "+
			"It will be removed from Terraform state. If it is still in your configuration, the next apply creates a new certificate pack. "+
			"Check for the old pack on the Edge Certificates page of the Cloudflare dashboard or with GET /zones/%[1]s/ssl/certificate_packs/%[2]s, "+
			"and delete it from the dashboard or with DELETE /zones/%[1]s/ssl/certificate_packs/%[2]s.",
			zoneID, id, status)
}
