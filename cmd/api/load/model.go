package load

import (
	httperrors "github.com/myrachanto/erroring"
)

type Load struct {
	ResponseCode int `json:"response_code"`
}

func (l Load) Validate() httperrors.HttpErr {
	if l.ResponseCode == 0 {
		return httperrors.NewBadRequestError("Response code!")
	}
	return nil
}
