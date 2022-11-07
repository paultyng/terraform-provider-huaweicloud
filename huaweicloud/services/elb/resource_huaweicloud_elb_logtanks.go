package elb

import (
	"context"
	"log"
	"time"

	"github.com/chnsz/golangsdk/openstack/elb/v3/logtanks"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func ResourceLogTanksV3() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLogTanksV3Create,
		ReadContext:   resourceLogTanksV3Read,
		UpdateContext: resourceLogTanksV3Update,
		DeleteContext: resourceLogTanksV3Delete,
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
			"loadbalancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_topic_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceLogTanksV3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	elbClient, err := config.ElbV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating elb client: %s", err)
	}

	createOpts := logtanks.CreateOpts{
		LoadbalancerID: d.Get("loadbalancer_id").(string),
		LogGroupId:     d.Get("log_group_id").(string),
		LogTopicId:     d.Get("log_topic_id").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	logTank, err := logtanks.Create(elbClient, createOpts).Extract()
	if err != nil {
		return diag.Errorf("error creating logtank: %s", err)
	}

	d.SetId(logTank.ID)

	return resourceLogTanksV3Read(ctx, d, meta)
}

func resourceLogTanksV3Read(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	elbClient, err := config.ElbV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating elb client: %s", err)
	}

	logTank, err := logtanks.Get(elbClient, d.Id()).Extract()
	if err != nil {
		return common.CheckDeletedDiag(d, err, "logtanks")
	}

	log.Printf("[DEBUG] Retrieved logtank %s: %#v", d.Id(), logTank)

	mErr := multierror.Append(nil,
		d.Set("loadbalancer_id", logTank.LoadbalancerID),
		d.Set("log_group_id", logTank.LogGroupId),
		d.Set("log_topic_id", logTank.LogTopicId),
	)

	if err := mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("error setting Dedicated ELB logtank fields: %s", err)
	}

	return nil
}

func resourceLogTanksV3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	elbClient, err := config.ElbV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating elb client: %s", err)
	}

	var updateOpts logtanks.UpdateOpts
	if d.HasChange("log_group_id") {
		updateOpts.LogGroupId = d.Get("log_group_id").(string)
	}
	if d.HasChange("log_topic_id") {
		updateOpts.LogTopicId = d.Get("log_topic_id").(string)
	}

	log.Printf("[DEBUG] Updating logtank %s with options: %#v", d.Id(), updateOpts)
	_, err = logtanks.Update(elbClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return diag.Errorf("unable to update logtank %s: %s", d.Id(), err)
	}

	return resourceLogTanksV3Read(ctx, d, meta)
}

func resourceLogTanksV3Delete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	elbClient, err := config.ElbV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating elb client: %s", err)
	}

	log.Printf("[DEBUG] Attempting to delete logtank %s", d.Id())
	err = logtanks.Delete(elbClient, d.Id()).ExtractErr()
	if err != nil {
		return diag.Errorf("unable to delete logtank %s: %s", d.Id(), err)
	}
	return nil
}
