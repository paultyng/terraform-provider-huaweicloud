package huaweicloud

import (
	"reflect"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandCdmClusterV1CreateAutoRemind(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	email, _ := navigateValue(d, []string{"email"}, nil)
	e, eok := email.([]interface{})

	phone, _ := navigateValue(d, []string{"phone_num"}, nil)
	p, pok := phone.([]interface{})

	return (eok && len(e) > 0) || (pok && len(p) > 0), nil
}

func expandCdmClusterV1CreateClusterIsScheduleBootOff(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	on, _ := navigateValue(d, []string{"schedule_boot_time"}, nil)
	on1, ok1 := on.(string)

	off, _ := navigateValue(d, []string{"schedule_off_time"}, nil)
	off1, ok2 := off.(string)

	return (ok1 && on1 != "") || (ok2 && off1 != ""), nil
}

func checkCssClusterV1ExtendClusterFinished(data interface{}) bool {
	//actions --- the behaviors on a cluster
	v, err := navigateValue(data, []string{"actions"}, nil)
	e, err := isEmptyValue(reflect.ValueOf(v))
	if err == nil && !e {
		return false
	}

	//actionProgress --- indicates the progress in percentage
	v, err = navigateValue(data, []string{"actionProgress"}, nil)
	e, err = isEmptyValue(reflect.ValueOf(v))
	if err == nil && !e {
		return false
	}

	instances, err := navigateValue(data, []string{"instances"}, nil)
	if err != nil {
		return false
	}
	if v, ok := instances.([]interface{}); ok {
		if len(v) == 0 {
			return false
		}
		for _, item := range v {
			status, err := navigateValue(item, []string{"status"}, nil)
			if err != nil {
				return false
			}
			if s, ok := status.(string); !ok || "200" != s {
				return false
			}
		}
		return true
	}
	return false
}

func expandCssClusterV1ExtendClusterNodeNum(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	t, _ := navigateValue(d, []string{"terraform_resource_data"}, nil)
	rd := t.(*schema.ResourceData)

	oldv, newv := rd.GetChange("expect_node_num")
	v := newv.(int) - oldv.(int)
	if v < 0 {
		return 0, fmtp.Errorf("expect_node_num only supports to be extended")
	}
	return v, nil
}

func expandCssClusterV1ExtendClusterVolumeSize(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	t, _ := navigateValue(d, []string{"terraform_resource_data"}, nil)
	rd := t.(*schema.ResourceData)

	//volume size location: reference to the Schema of css_cluster_v1
	oldv, newv := rd.GetChange("node_config.0.volume.0.size")
	v := newv.(int) - oldv.(int)
	if v < 0 {
		return 0, fmtp.Errorf("volume size only supports to be extended")
	}
	return v, nil
}

func expandDisStreamV2CreateAutoCaleEnable(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	max, _ := navigateValue(d, []string{"auto_scale_max_partition_count"}, nil)
	max1, ok1 := max.(int)

	min, _ := navigateValue(d, []string{"auto_scale_min_partition_count"}, nil)
	min1, ok2 := min.(int)

	return (ok1 && max1 > 0) && (ok2 && min1 > 0), nil
}
