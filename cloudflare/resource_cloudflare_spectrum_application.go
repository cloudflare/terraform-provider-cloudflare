package cloudflare

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareSpectrumApplication() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareSpectrumApplicationSchema(),
		CreateContext: resourceCloudflareSpectrumApplicationCreate,
		ReadContext:   resourceCloudflareSpectrumApplicationRead,
		UpdateContext: resourceCloudflareSpectrumApplicationUpdate,
		DeleteContext: resourceCloudflareSpectrumApplicationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareSpectrumApplicationImport,
		},
	}
}

func resourceCloudflareSpectrumApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	newSpectrumApp := applicationFromResource(d)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Creating Cloudflare Spectrum Application from struct: %+v", newSpectrumApp)

	r, err := client.CreateSpectrumApplication(ctx, zoneID, newSpectrumApp)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating spectrum application for zone"))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.ID)

	log.Printf("[INFO] Cloudflare Spectrum Application ID: %s", d.Id())

	return resourceCloudflareSpectrumApplicationRead(ctx, d, meta)
}

func resourceCloudflareSpectrumApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	application := applicationFromResource(d)

	log.Printf("[INFO] Updating Cloudflare Spectrum Application from struct: %+v", application)

	_, err := client.UpdateSpectrumApplication(ctx, zoneID, application.ID, application)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating spectrum application for zone"))
	}

	return resourceCloudflareSpectrumApplicationRead(ctx, d, meta)
}

func resourceCloudflareSpectrumApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	applicationID := d.Id()

	application, err := client.SpectrumApplication(ctx, zoneID, applicationID)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Spectrum application %s in zone %s not found", applicationID, zoneID)
			d.SetId("")
			return nil
		}
		return diag.FromErr(errors.Wrap(err,
			fmt.Sprintf("Error reading spectrum application resource from API for resource %s in zone %s", applicationID, zoneID)))
	}

	d.Set("protocol", application.Protocol)

	if err := d.Set("dns", flattenDNS(application.DNS)); err != nil {
		log.Printf("[WARN] Error setting dns on spectrum application %q: %s", d.Id(), err)
	}

	if len(application.OriginDirect) > 0 {
		if err := d.Set("origin_direct", application.OriginDirect); err != nil {
			log.Printf("[WARN] Error setting origin direct on spectrum application %q: %s", d.Id(), err)
		}
	}

	if application.OriginDNS != nil {
		if err := d.Set("origin_dns", flattenOriginDNS(application.OriginDNS)); err != nil {
			log.Printf("[WARN] Error setting origin dns on spectrum application %q: %s", d.Id(), err)
		}
	}

	if application.OriginPort != nil {
		if application.OriginPort.Port > 0 {
			d.Set("origin_port", int(application.OriginPort.Port))
		} else {
			if err := d.Set("origin_port_range", flattenOriginPortRange(application.OriginPort)); err != nil {
				log.Printf("[WARN] Error setting origin port range on spectrum application %q: %s", d.Id(), err)
			}
		}
	}

	if application.EdgeIPs != nil {
		if err := d.Set("edge_ips", flattenEdgeIPs(application.EdgeIPs)); err != nil {
			log.Printf("[WARN] Error setting Edge IPs on spectrum application %q: %s", d.Id(), err)
		}
		if err := d.Set("edge_ip_connectivity", application.EdgeIPs.Connectivity.String()); err != nil {
			log.Printf("[WARN] Error setting Edge IP connectivity on spectrum application %q: %s", d.Id(), err)
		}
	}

	d.Set("tls", application.TLS)
	d.Set("traffic_type", application.TrafficType)
	d.Set("ip_firewall", application.IPFirewall)
	d.Set("proxy_protocol", application.ProxyProtocol)
	d.Set("argo_smart_routing", application.ArgoSmartRouting)

	return nil
}

func resourceCloudflareSpectrumApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	applicationID := d.Id()

	log.Printf("[INFO] Deleting Cloudflare Spectrum Application: %s in zone: %s", applicationID, zoneID)

	err := client.DeleteSpectrumApplication(ctx, zoneID, applicationID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Spectrum Application: %w", err))
	}

	return nil
}

func resourceCloudflareSpectrumApplicationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var applicationID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		applicationID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/applicationID\"", d.Id())
	}

	d.Set("zone_id", zoneID)
	d.SetId(applicationID)

	resourceCloudflareSpectrumApplicationRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func expandDNS(d interface{}) cloudflare.SpectrumApplicationDNS {
	cfg := d.([]interface{})
	dns := cloudflare.SpectrumApplicationDNS{}

	m := cfg[0].(map[string]interface{})
	dns.Type = m["type"].(string)
	dns.Name = m["name"].(string)

	return dns
}

func expandOriginDNS(d interface{}) *cloudflare.SpectrumApplicationOriginDNS {
	cfg := d.([]interface{})
	dns := &cloudflare.SpectrumApplicationOriginDNS{}

	m := cfg[0].(map[string]interface{})
	dns.Name = m["name"].(string)

	return dns
}

func expandOriginPortRange(d interface{}) *cloudflare.SpectrumApplicationOriginPort {
	cfg := d.([]interface{})
	port := &cloudflare.SpectrumApplicationOriginPort{}

	m := cfg[0].(map[string]interface{})
	port.Start = uint16(m["start"].(int))
	port.End = uint16(m["end"].(int))

	return port
}

func flattenDNS(dns cloudflare.SpectrumApplicationDNS) []map[string]interface{} {
	flattened := map[string]interface{}{}
	flattened["type"] = dns.Type
	flattened["name"] = dns.Name

	return []map[string]interface{}{flattened}
}

func flattenOriginDNS(dns *cloudflare.SpectrumApplicationOriginDNS) []map[string]interface{} {
	flattened := map[string]interface{}{}
	flattened["name"] = dns.Name

	return []map[string]interface{}{flattened}
}

func flattenOriginPortRange(port *cloudflare.SpectrumApplicationOriginPort) []map[string]interface{} {
	flattened := map[string]interface{}{}
	flattened["start"] = port.Start
	flattened["end"] = port.End

	return []map[string]interface{}{flattened}
}

func flattenEdgeIPs(edgeIPs *cloudflare.SpectrumApplicationEdgeIPs) []string {
	flattened := make([]string, 0)

	for _, ip := range edgeIPs.IPs {
		flattened = append(flattened, ip.String())
	}

	return flattened
}

func applicationFromResource(d *schema.ResourceData) cloudflare.SpectrumApplication {
	application := cloudflare.SpectrumApplication{
		ID:       d.Id(),
		Protocol: d.Get("protocol").(string),
		DNS:      expandDNS(d.Get("dns")),
	}

	if originDirect, ok := d.GetOk("origin_direct"); ok {
		application.OriginDirect = expandInterfaceToStringList(originDirect.([]interface{}))
	}

	if originDNS, ok := d.GetOk("origin_dns"); ok {
		application.OriginDNS = expandOriginDNS(originDNS)
	}

	if originPort, ok := d.GetOk("origin_port"); ok {
		application.OriginPort = &cloudflare.SpectrumApplicationOriginPort{Port: uint16(originPort.(int))}
	} else if originPortRange, ok := d.GetOk("origin_port_range"); ok {
		application.OriginPort = expandOriginPortRange(originPortRange)
	}

	if tls, ok := d.GetOk("tls"); ok {
		application.TLS = tls.(string)
	}

	if trafficType, ok := d.GetOk("traffic_type"); ok {
		application.TrafficType = trafficType.(string)
	}

	if ipFirewall, ok := d.GetOk("ip_firewall"); ok {
		application.IPFirewall = ipFirewall.(bool)
	}

	if proxyProtocol, ok := d.GetOk("proxy_protocol"); ok {
		application.ProxyProtocol = cloudflare.ProxyProtocol(proxyProtocol.(string))
	}

	if argoSmartRouting, ok := d.GetOk("argo_smart_routing"); ok {
		application.ArgoSmartRouting = argoSmartRouting.(bool)
	}

	connectivity := cloudflare.SpectrumApplicationConnectivity(cloudflare.SpectrumConnectivityAll)
	application.EdgeIPs = &cloudflare.SpectrumApplicationEdgeIPs{
		Type:         cloudflare.SpectrumEdgeTypeDynamic,
		Connectivity: &connectivity,
	}

	if edgeIPConnectivity, ok := d.GetOk("edge_ip_connectivity"); ok {
		connectivity = cloudflare.SpectrumApplicationConnectivity(edgeIPConnectivity.(string))
	}

	if edgeIPs, ok := d.GetOk("edge_ips"); ok {
		application.EdgeIPs = &cloudflare.SpectrumApplicationEdgeIPs{
			Type: cloudflare.SpectrumEdgeTypeStatic,
		}
		// connectivity = cloudflare.SpectrumConnectivityStatic
		for _, value := range edgeIPs.([]interface{}) {
			application.EdgeIPs.IPs = append(application.EdgeIPs.IPs, net.ParseIP(value.(string)))
		}
	}

	return application
}
