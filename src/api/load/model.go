package load

import (
	httperrors "github.com/myrachanto/erroring"
)

type Load struct {
	Url  string `json:"url,omitempty"`
	Size int    `json:"size,omitempty"`
}

func (l Load) Validate() httperrors.HttpErr {
	if l.Url == "" {
		return httperrors.NewBadRequestError("Url must not be empty")
	}
	return nil
}
