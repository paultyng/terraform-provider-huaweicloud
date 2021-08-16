package trigger

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// CreateResult represents a result of the Create method.
type CreateResult struct {
	commonResult
}

// GetResult represents a result of the Update method.
type GetResult struct {
	commonResult
}

// UpdateResult represents a result of the Update method.
type UpdateResult struct {
	golangsdk.ErrResult
}

// DeleteResult represents a result of the Delete method.
type DeleteResult struct {
	golangsdk.ErrResult
}

// Trigger is a struct that represents the result of Create and Get methods.
type Trigger struct {
	TriggerId       string                 `json:"trigger_id"`
	TriggerTypeCode string                 `json:"trigger_type_code"`
	EventData       map[string]interface{} `json:"event_data"`
	EventTypeCode   string                 `json:"event_type_code"`
	Status          string                 `json:"trigger_status"`
	LastUpdatedTime string                 `json:"last_updated_time"`
	CreatedTime     string                 `json:"created_time"`
	LastError       string                 `json:"last_error"`
}

func (r commonResult) Extract() (*Trigger, error) {
	var s Trigger
	err := r.ExtractInto(&s)
	return &s, err
}

// TriggerPage represents the response pages of the List method.
type TriggerPage struct {
	pagination.SinglePageBase
}

// ExtractList is a method which to extract the response to a trigger list.
func ExtractList(r pagination.Page) ([]Trigger, error) {
	var s []Trigger
	err := (r.(TriggerPage)).ExtractInto(&s)
	return s, err
}
