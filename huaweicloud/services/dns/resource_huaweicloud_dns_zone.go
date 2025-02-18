package dns

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/common/tags"
	"github.com/chnsz/golangsdk/openstack/dns/v2/zones"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API DNS POST /v2/zones/{zone_id}/associaterouter
// @API DNS POST /v2/zones/{zone_id}/disassociaterouter
// @API DNS DELETE /v2/zones/{zone_id}
// @API DNS GET /v2/zones/{zone_id}
// @API DNS PATCH /v2/zones/{zone_id}
// @API DNS POST /v2/zones
// @API DNS POST /v2/{project_id}/DNS-public_zone/{resource_id}/tags/action
// @API DNS POST /v2/{project_id}/DNS-private_zone/{resource_id}/tags/action
// @API DNS GET /v2/{project_id}/DNS-public_zone/{resource_id}/tags
// @API DNS GET /v2/{project_id}/DNS-private_zone/{resource_id}/tags
// @API DNS PUT /v2/zones/{zone_id}/statuses
func ResourceDNSZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDNSZoneCreate,
		ReadContext:   resourceDNSZoneRead,
		UpdateContext: resourceDNSZoneUpdate,
		DeleteContext: resourceDNSZoneDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(_, oldVal, newVal string, _ *schema.ResourceData) bool {
					return strings.TrimSuffix(oldVal, ".") == strings.TrimSuffix(newVal, ".")
				},
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "public",
				ValidateFunc: validation.StringInSlice([]string{"public", "private"}, false),
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"router": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"router_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"router_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"enterprise_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The status of the zone.`,
			},
			"masters": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": common.TagsSchema(),
		},
	}
}

func resourceDNSRouter(d *schema.ResourceData, region string) *zones.RouterOpts {
	router := d.Get("router").(*schema.Set).List()
	if len(router) > 0 {
		routerOpts := zones.RouterOpts{}

		c := router[0].(map[string]interface{})
		if val, ok := c["router_id"]; ok {
			routerOpts.RouterID = val.(string)
		}
		if val, ok := c["router_region"]; ok {
			routerOpts.RouterRegion = val.(string)
		} else {
			routerOpts.RouterRegion = region
		}
		return &routerOpts
	}
	return nil
}

func resourceDNSZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	var dnsClient *golangsdk.ServiceClient

	dnsClient, err := cfg.DnsV2Client(region)
	if err != nil {
		return diag.Errorf("error creating DNS client: %s", err)
	}

	zoneType := d.Get("zone_type").(string)
	router := d.Get("router").(*schema.Set).List()

	// router is required when creating private zone
	if zoneType == "private" {
		if len(router) < 1 {
			return diag.Errorf("the argument (router) is required when creating DNS private zone")
		}
		// update the endpoint with region when creating private zone
		dnsClient, err = cfg.DnsWithRegionClient(cfg.GetRegion(d))
		if err != nil {
			return diag.Errorf("error creating DNS region client: %s", err)
		}
	}

	createOpts := zones.CreateOpts{
		Name:                d.Get("name").(string),
		TTL:                 d.Get("ttl").(int),
		Email:               d.Get("email").(string),
		Description:         d.Get("description").(string),
		ZoneType:            zoneType,
		EnterpriseProjectID: cfg.GetEnterpriseProjectID(d),
		Router:              resourceDNSRouter(d, region),
	}

	log.Printf("[DEBUG] Create options: %#v", createOpts)
	n, err := zones.Create(dnsClient, createOpts).Extract()
	if err != nil {
		return diag.Errorf("error creating DNS zone: %s", err)
	}

	d.SetId(n.ID)
	log.Printf("[DEBUG] Waiting for DNS zone (%s) to become available", n.ID)
	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE"},
		Pending:    []string{"PENDING"},
		Refresh:    waitForDNSZone(dnsClient, n.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf(
			"error waiting for DNS zone (%s) to become ACTIVE for creation: %s",
			n.ID, err)
	}

	// router length >1 when creating private zone
	if zoneType == "private" {
		// AssociateZone for the other routers
		routerList := getDNSRouters(d, region)
		if len(routerList) > 1 {
			for i := range routerList {
				// Skip the first router
				if i > 0 {
					log.Printf("[DEBUG] Creating associate zone options: %#v", routerList[i])
					_, err := zones.AssociateZone(dnsClient, n.ID, routerList[i]).Extract()
					if err != nil {
						return diag.Errorf("error associate zone: %s", err)
					}

					log.Printf("[DEBUG] Waiting for associate zone (%s) to router (%s) become ACTIVE",
						n.ID, routerList[i].RouterID)
					stateRouterConf := &resource.StateChangeConf{
						Target:     []string{"ACTIVE"},
						Pending:    []string{"PENDING"},
						Refresh:    waitForDNSZoneRouter(dnsClient, n.ID, routerList[i].RouterID),
						Timeout:    d.Timeout(schema.TimeoutCreate),
						Delay:      5 * time.Second,
						MinTimeout: 3 * time.Second,
					}

					_, err = stateRouterConf.WaitForStateContext(ctx)
					if err != nil {
						return diag.Errorf("error waiting for associate zone (%s) to router (%s) "+
							"become ACTIVE: %s", n.ID, routerList[i].RouterID, err)
					}
				} else {
					log.Printf("[DEBUG] First router options: %#v", routerList[i])
				}
			}
		}
	}

	// After zone is created, the status is ACTIVE (ENABLE).
	// This action cannot be called repeatedly.
	if v, ok := d.GetOk("status"); ok && v != "ENABLE" {
		if err := updateZoneStatus(ctx, d, dnsClient, d.Timeout(schema.TimeoutCreate)); err != nil {
			return diag.FromErr(err)
		}
	}

	// set tags
	tagRaw := d.Get("tags").(map[string]interface{})
	if len(tagRaw) > 0 {
		resourceType, err := utils.GetDNSZoneTagType(zoneType)
		if err != nil {
			return diag.Errorf("error getting resource type of DNS zone %s: %s", n.ID, err)
		}

		taglist := utils.ExpandResourceTags(tagRaw)
		if tagErr := tags.Create(dnsClient, resourceType, n.ID, taglist).ExtractErr(); tagErr != nil {
			return diag.Errorf("error setting tags of DNS zone %s: %s", n.ID, tagErr)
		}
	}

	log.Printf("[DEBUG] Created DNS zone %s: %#v", n.ID, n)
	return resourceDNSZoneRead(ctx, d, meta)
}

func resourceDNSZoneRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)

	// we can not get the corresponding client by zone type in import scene
	dnsClient, err := conf.DnsV2Client(region)
	if err != nil {
		return diag.Errorf("error creating DNS client: %s", err)
	}

	var zoneInfo *zones.Zone
	zoneInfo, err = zones.Get(dnsClient, d.Id()).Extract()
	if err != nil {
		log.Printf("[WARN] fetching zone failed with DNS global endpoint: %s", err)
		// an error occurred while fetching the zone with DNS global endpoint
		// try to fetch it again with DNS region endpoint
		var clientErr error
		dnsClient, clientErr = conf.DnsWithRegionClient(conf.GetRegion(d))
		if clientErr != nil {
			// it looks tricky as we return the fetching error rather than clientErr
			return common.CheckDeletedDiag(d, err, "zone")
		}

		zoneInfo, err = zones.Get(dnsClient, d.Id()).Extract()
		if err != nil {
			return common.CheckDeletedDiag(d, err, "zone")
		}
	}

	log.Printf("[DEBUG] Retrieved zone %s: %#v", d.Id(), zoneInfo)

	mErr := multierror.Append(nil,
		d.Set("name", zoneInfo.Name),
		d.Set("email", zoneInfo.Email),
		d.Set("description", zoneInfo.Description),
		d.Set("ttl", zoneInfo.TTL),
		d.Set("masters", zoneInfo.Masters),
		d.Set("region", region),
		d.Set("zone_type", zoneInfo.ZoneType),
		d.Set("enterprise_project_id", zoneInfo.EnterpriseProjectID),
		// The private zone also returns the "status" attribute.
		d.Set("status", parseZoneStatus(zoneInfo.Status)),
	)

	// save tags
	if resourceType, err := utils.GetDNSZoneTagType(zoneInfo.ZoneType); err == nil {
		resourceTags, err := tags.Get(dnsClient, resourceType, d.Id()).Extract()
		if err == nil {
			tagmap := utils.TagsToMap(resourceTags.Tags)
			mErr = multierror.Append(mErr, d.Set("tags", tagmap))
		} else {
			log.Printf("[WARN] Error fetching DNS zone tags: %s", err)
		}
	}

	if mErr.ErrorOrNil() != nil {
		return diag.Errorf("error setting resource: %s", mErr)
	}

	return nil
}

func parseZoneStatus(status string) string {
	if status == "ACTIVE" {
		return "ENABLE"
	}
	return status
}

func resourceDNSZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)
	var dnsClient *golangsdk.ServiceClient

	dnsClient, err := conf.DnsV2Client(region)
	if err != nil {
		return diag.Errorf("error creating DNS client: %s", err)
	}

	zoneType := d.Get("zone_type").(string)
	router := d.Get("router").(*schema.Set).List()

	// router is required when updating private zone
	if zoneType == "private" {
		if len(router) < 1 {
			return diag.Errorf("the argument (router) is required when updating DNS private zone")
		}
		// update the endpoint with region when creating private zone
		dnsClient, err = conf.DnsWithRegionClient(conf.GetRegion(d))
		if err != nil {
			return diag.Errorf("error creating DNS region client: %s", err)
		}
	}

	if d.HasChanges("description", "ttl", "email") {
		if err := updateDNSZone(ctx, d, dnsClient); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("status") {
		if err := updateZoneStatus(ctx, d, dnsClient, d.Timeout(schema.TimeoutUpdate)); err != nil {
			return diag.FromErr(err)
		}
	}

	// This operation is supported only when the zone status is enabled.
	if d.HasChange("router") && zoneType == "private" {
		if err := updateDNSZoneRouters(ctx, d, dnsClient, region); err != nil {
			return diag.FromErr(err)
		}
	}

	// update tags
	resourceType, err := utils.GetDNSZoneTagType(zoneType)
	if err != nil {
		return diag.Errorf("error getting resource type of DNS zone %s: %s", d.Id(), err)
	}

	tagErr := utils.UpdateResourceTags(dnsClient, d, resourceType, d.Id())
	if tagErr != nil {
		return diag.Errorf("error updating tags of DNS zone %s: %s", d.Id(), tagErr)
	}

	return resourceDNSZoneRead(ctx, d, meta)
}

func updateDNSZone(ctx context.Context, d *schema.ResourceData, client *golangsdk.ServiceClient) error {
	var updateOpts zones.UpdateOpts
	if d.HasChange("email") {
		updateOpts.Email = d.Get("email").(string)
	}
	if d.HasChange("ttl") {
		updateOpts.TTL = d.Get("ttl").(int)
	}
	if d.HasChange("description") {
		updateOpts.Description = d.Get("description").(string)
	}

	log.Printf("[DEBUG] Updating zone %s with options: %#v", d.Id(), updateOpts)
	_, err := zones.Update(client, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("error updating DNS zone: %s", err)
	}

	log.Printf("[DEBUG] Waiting for DNS zone (%s) to update", d.Id())
	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE", "DISABLE"},
		Pending:    []string{"PENDING"},
		Refresh:    waitForDNSZone(client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for DNS zone (%s) to become ACTIVE for update: %s", d.Id(), err)
	}
	return nil
}

func updateDNSZoneRouters(ctx context.Context, d *schema.ResourceData, client *golangsdk.ServiceClient,
	region string) error {
	associateList, disassociateList, err := resourceGetDNSRouters(client, d, region)
	if err != nil {
		return fmt.Errorf("error getting DNS zone router: %s", err)
	}
	if len(associateList) > 0 {
		// AssociateZone
		for i := range associateList {
			log.Printf("[DEBUG] Updating associate zone options: %#v", associateList[i])
			_, err := zones.AssociateZone(client, d.Id(), associateList[i]).Extract()
			if err != nil {
				return fmt.Errorf("error associate zone: %s", err)
			}

			log.Printf("[DEBUG] Waiting for associate zone (%s) to router (%s) become ACTIVE",
				d.Id(), associateList[i].RouterID)
			stateRouterConf := &resource.StateChangeConf{
				Target:     []string{"ACTIVE"},
				Pending:    []string{"PENDING"},
				Refresh:    waitForDNSZoneRouter(client, d.Id(), associateList[i].RouterID),
				Timeout:    d.Timeout(schema.TimeoutUpdate),
				Delay:      5 * time.Second,
				MinTimeout: 3 * time.Second,
			}

			_, err = stateRouterConf.WaitForStateContext(ctx)
			if err != nil {
				return fmt.Errorf("error waiting for associate zone (%s) to router (%s) become ACTIVE: %s",
					d.Id(), associateList[i].RouterID, err)
			}
		}
	}
	if len(disassociateList) > 0 {
		// DisassociateZone
		for j := range disassociateList {
			log.Printf("[DEBUG] Updating disassociate zone options: %#v", disassociateList[j])
			_, err := zones.DisassociateZone(client, d.Id(), disassociateList[j]).Extract()
			if err != nil {
				return fmt.Errorf("error disassociate zone: %s", err)
			}

			log.Printf("[DEBUG] Waiting for disassociate zone (%s) to router (%s) become DELETED",
				d.Id(), disassociateList[j].RouterID)
			stateRouterConf := &resource.StateChangeConf{
				Target:     []string{"DELETED"},
				Pending:    []string{"ACTIVE", "PENDING", "ERROR"},
				Refresh:    waitForDNSZoneRouter(client, d.Id(), disassociateList[j].RouterID),
				Timeout:    d.Timeout(schema.TimeoutUpdate),
				Delay:      5 * time.Second,
				MinTimeout: 3 * time.Second,
			}

			_, err = stateRouterConf.WaitForStateContext(ctx)
			if err != nil {
				return fmt.Errorf("error waiting for disassociate zone (%s) to router (%s) become DELETED: %s",
					d.Id(), disassociateList[j].RouterID, err)
			}
		}
	}
	return nil
}

func updateZoneStatus(ctx context.Context, d *schema.ResourceData, client *golangsdk.ServiceClient, timeout time.Duration) error {
	opts := zones.UpdateStatusOpts{
		ZoneId: d.Id(),
		Status: d.Get("status").(string),
	}
	err := zones.UpdateZoneStatus(client, opts)
	if err != nil {
		return fmt.Errorf("error updating the status of the zone: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE", "DISABLE", "FREEZE"},
		Pending:    []string{"PENDING"},
		Refresh:    waitForDNSZone(client, d.Id()),
		Timeout:    timeout,
		MinTimeout: 10 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for updating the zone status completed: %s", err)
	}

	return nil
}

func resourceDNSZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	var dnsClient *golangsdk.ServiceClient
	var err error

	zoneType := d.Get("zone_type").(string)
	// update the endpoint with region when creating private zone
	if zoneType == "private" {
		dnsClient, err = conf.DnsWithRegionClient(conf.GetRegion(d))
	} else {
		dnsClient, err = conf.DnsV2Client(conf.GetRegion(d))
	}
	if err != nil {
		return diag.Errorf("error creating DNS client: %s", err)
	}

	_, err = zones.Delete(dnsClient, d.Id()).Extract()
	if err != nil {
		return diag.Errorf("error deleting DNS zone: %s", err)
	}

	log.Printf("[DEBUG] Waiting for DNS zone (%s) to become DELETED", d.Id())
	stateConf := &resource.StateChangeConf{
		Target: []string{"DELETED"},
		// we allow to try to delete ERROR zone
		Pending:    []string{"ACTIVE", "PENDING", "ERROR"},
		Refresh:    waitForDNSZone(dnsClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf(
			"error waiting for DNS zone (%s) to delete: %s",
			d.Id(), err)
	}

	d.SetId("")
	return nil
}

func waitForDNSZone(dnsClient *golangsdk.ServiceClient, zoneId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		zone, err := zones.Get(dnsClient, zoneId).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return zone, "DELETED", nil
			}

			return nil, "", err
		}

		log.Printf("[DEBUG] DNS zone (%s) current status: %s", zone.ID, zone.Status)
		return zone, parseStatus(zone.Status), nil
	}
}

func getDNSRouters(d *schema.ResourceData, region string) []zones.RouterOpts {
	router := d.Get("router").(*schema.Set).List()
	if len(router) == 0 {
		return nil
	}

	res := make([]zones.RouterOpts, len(router))
	for i := range router {
		ro := zones.RouterOpts{}
		c := router[i].(map[string]interface{})
		if val, ok := c["router_id"]; ok {
			ro.RouterID = val.(string)
		}
		if val, ok := c["router_region"]; ok {
			ro.RouterRegion = val.(string)
		} else {
			ro.RouterRegion = region
		}

		res[i] = ro
	}
	return res
}

func waitForDNSZoneRouter(dnsClient *golangsdk.ServiceClient, zoneId string, routerId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		zone, err := zones.Get(dnsClient, zoneId).Extract()
		if err != nil {
			return nil, "", err
		}
		for i := range zone.Routers {
			if routerId == zone.Routers[i].RouterID {
				log.Printf("[DEBUG] DNS zone (%s) router (%s) current status: %s",
					zoneId, routerId, zone.Routers[i].Status)
				return zone, parseStatus(zone.Routers[i].Status), nil
			}
		}
		return zone, "DELETED", nil
	}
}

func resourceGetDNSRouters(dnsClient *golangsdk.ServiceClient, d *schema.ResourceData,
	region string) ([]zones.RouterOpts, []zones.RouterOpts, error) {
	// get zone info from api
	n, err := zones.Get(dnsClient, d.Id()).Extract()
	if err != nil {
		return nil, nil, common.CheckDeleted(d, err, "zone")
	}
	// get routers from local
	localRouters := getDNSRouters(d, region)

	// get associateMap
	associateMap := make(map[string]zones.RouterOpts)
	for _, local := range localRouters {
		// Check if local is found in api
		found := false
		for _, raw := range n.Routers {
			if local.RouterID == raw.RouterID {
				found = true
				break
			}
		}
		// If local is not found in api
		if !found {
			associateMap[local.RouterID] = local
		}
	}

	// convert associateMap to associateList
	associateList := make([]zones.RouterOpts, len(associateMap))
	var i = 0
	for _, associateRouter := range associateMap {
		associateList[i] = associateRouter
		i++
	}

	// get disassociateMap
	disassociateMap := make(map[string]zones.RouterOpts)
	for _, raw := range n.Routers {
		// Check if api is found in local
		found := false
		for _, local := range localRouters {
			if raw.RouterID == local.RouterID {
				found = true
				break
			}
		}
		// If api is not found in local
		if !found {
			disassociateMap[raw.RouterID] = zones.RouterOpts{
				RouterID:     raw.RouterID,
				RouterRegion: raw.RouterRegion,
			}
		}
	}

	// convert disassociateMap to disassociateList
	disassociateList := make([]zones.RouterOpts, len(disassociateMap))
	var j = 0
	for _, disassociateRouter := range disassociateMap {
		disassociateList[j] = disassociateRouter
		j++
	}

	return associateList, disassociateList, nil
}
