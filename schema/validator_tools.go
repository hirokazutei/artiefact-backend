package schema

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type validatorer interface {
	Validate(interface{}) error
}

type Error struct {
	Code   string `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
	Errors []struct {
		Code       string `json:"code,omitempty"`
		Detail     string `json:"detail,omitempty"`
		Field      string `json:"field,omitempty"`
		UserDetail string `json:"user_detail,omitempty"`
	} `json:"errors,omitempty"`
	Status     int64  `json:"status"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type"`
	UserDetail string `json:"user_detail,omitempty"`
	UserTitle  string `json:"user_title,omitempty"`
}

// Validate validates the JSON
func Validate(req interface{}, validator validatorer, r *http.Request) (*Error, error) {
	body, err := ioutil.ReadAll(r.Body)
	if bytes.Compare(body, []byte("")) == 0 {
		body = []byte("{}")
	}
	if err != nil {
		e := Error{
			Status: 500,
			Type:   "JSON parse error",
		}
		return &e, errors.Wrap(err, "failed to decode json")
	}
	if err := json.Unmarshal(body, req); err != nil {
		e := Error{
			Status: 500,
			Type:   "JSON parse error",
		}
		return &e, errors.Wrap(err, "failed to decode json")
	}
	var param map[string]interface{}
	if err := json.Unmarshal(body, &param); err != nil {
		e := Error{
			Status: 500,
			Type:   "JSON parse error",
		}
		return &e, errors.Wrap(err, "failed to decode json")
	}
	if err := validator.Validate(param); err != nil {
		e := Error{
			Status: 500,
			Type:   "JSON parse error",
		}
		return &e, errors.Wrap(err, "failed to validate params")
	}
	return nil, nil
}
