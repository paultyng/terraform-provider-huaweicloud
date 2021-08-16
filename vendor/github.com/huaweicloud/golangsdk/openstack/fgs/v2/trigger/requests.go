package trigger

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOpts is a struct which will be used to create a new trigger.
type CreateOpts struct {
	// Trigger type, which support:
	//   TIMER
	//   APIG
	//   CTS
	//   DDS
	//   DMS
	//   DIS
	//   LTS
	//   OBS
	//   KAFKA
	//   SMN
	TriggerTypeCode string `json:"trigger_type_code" required:"true"`
	// Trigger status, which support:
	//   ACTIVE
	//   DISABLED
	TriggerStatus string `json:"trigger_status,omitempty"`
	// Message code.
	EventTypeCode string `json:"event_type_code" required:"true"`
	// Event struct.
	EventData map[string]interface{} `json:"event_data" required:"true"`
}

// CreateOptsBuilder is an interface used to support the construction of the request body for trigger creates.
type CreateOptsBuilder interface {
	ToCreateTriggerMap() (map[string]interface{}, error)
}

// ToCreateTriggerMap is a method which to build a request body by the CreateOpts.
func (opts CreateOpts) ToCreateTriggerMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method to create trigger through function urn and CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder, urn string) (r CreateResult) {
	b, err := opts.ToCreateTriggerMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c, urn), b, &r.Body, nil)
	return
}

// List is a method to obtain an array of one or more trigger for function graph according to the urn.
func List(c *golangsdk.ServiceClient, urn string) pagination.Pager {
	url := listURL(c, urn)
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return TriggerPage{pagination.SinglePageBase(r)}
	})
}

// Get is a method to obtain a trigger through function urn, function code and trigger ID.
func Get(c *golangsdk.ServiceClient, urn, triggerType, triggerId string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, urn, triggerType, triggerId), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOpts is a struct which will be used to update existing trigger.
type UpdateOpts struct {
	TriggerStatus string `json:"trigger_status" required:"true"`
}

// UpdateOptsBuilder is an interface used to support the construction of the request body for trigger updates.
type UpdateOptsBuilder interface {
	ToUpdateTriggerMap() (map[string]interface{}, error)
}

// ToUpdateTriggerMap is a method which to build a request body by the UpdateOpts.
func (opts UpdateOpts) ToUpdateTriggerMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update is a method to update existing trigger through function urn, function code, trigger ID and UpdateOpts.
func Update(c *golangsdk.ServiceClient, opts UpdateOptsBuilder, urn, triggerType, triggerId string) (r UpdateResult) {
	b, err := opts.ToUpdateTriggerMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, urn, triggerType, triggerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete is a method to delete existing trigger.
func Delete(c *golangsdk.ServiceClient, urn, triggerType, triggerId string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, urn, triggerType, triggerId), &golangsdk.RequestOpts{
		OkCodes: []int{204, 200},
	})
	return
}

// DeleteAll is a method to delete all triggers from a function
func DeleteAll(c *golangsdk.ServiceClient, urn string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteAllURL(c, urn), nil)
	return
}
