package eps

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"
)

// ParseQueryError403 is a method used to parse whether a 403 error message should be ignored.
// For the EPS service, there are some known 403 error codes:
// + EPS.0004: The current user lacks permission for this service.
func ParseQueryError403(err error, specErrors []string, msg ...string) error {
	var err403 golangsdk.ErrDefault403
	if errors.As(err, &err403) {
		var apiError interface{}
		if jsonErr := json.Unmarshal(err403.Body, &apiError); jsonErr != nil {
			return err
		}

		errCode, searchErr := jmespath.Search("error_code", apiError)
		if searchErr != nil {
			return err
		}

		for _, v := range specErrors {
			if errCode == v {
				if len(msg) > 0 {
					// Record the custom message to the log.
					log.Println(msg[0])
				}
				return nil
			}
		}
	}
	return err
}
