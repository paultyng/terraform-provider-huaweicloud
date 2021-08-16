package trigger

import "github.com/huaweicloud/golangsdk"

const (
	FGS     = "fgs"
	TRIGGER = "triggers"
)

func listURL(c *golangsdk.ServiceClient, functionUrn string) string {
	return c.ServiceURL(FGS, TRIGGER, functionUrn)
}

func createURL(c *golangsdk.ServiceClient, functionUrn string) string {
	return listURL(c, functionUrn)
}

func deleteAllURL(c *golangsdk.ServiceClient, functionUrn string) string {
	return listURL(c, functionUrn)
}

func deleteURL(c *golangsdk.ServiceClient, functionUrn, triggerTypeCode, triggerId string) string {
	return resourceURL(c, functionUrn, triggerTypeCode, triggerId)
}

func resourceURL(c *golangsdk.ServiceClient, functionUrn, triggerTypeCode, triggerId string) string {
	return c.ServiceURL(FGS, TRIGGER, functionUrn, triggerTypeCode, triggerId)
}
