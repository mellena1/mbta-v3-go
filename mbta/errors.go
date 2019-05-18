package mbta

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrRateLimitExceeded = errors.New("you have exceeded your allowed usage rate")
	ErrForbidden         = errors.New("forbidden")
	ErrMustSpecifyID     = errors.New("must specify an id (cannot be an empty string)")
)

// BadRequestError error type holding the returned info about the bad request
type BadRequestError struct {
	SourceParameter string // The name of parameter that caused the error
	Detail          string // A short, human-readable summary of the problem
	Code            string // An application-specific error code
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("parameter \"%s\" caused error [%s]: (%s)", e.SourceParameter, e.Code, e.Detail)
}

func getBadRequestError(body io.Reader) BadRequestError {
	var jsonErr struct {
		Errors []struct {
			Status string `json:"status"`
			Source struct {
				Parameter string `json:"parameter"`
			} `json:"source"`
			Detail string `json:"title"`
			Code   string `json:"code"`
		} `json:"errors"`
	}
	json.NewDecoder(body).Decode(&jsonErr)

	return BadRequestError{
		SourceParameter: jsonErr.Errors[0].Source.Parameter,
		Detail:          jsonErr.Errors[0].Detail,
		Code:            jsonErr.Errors[0].Code,
	}
}

func getSpecialError(resp *http.Response, err error) error {
	switch resp.StatusCode {
	case 400:
		return getBadRequestError(resp.Body)
	case 403:
		return ErrForbidden
	case 404:
		return getBadRequestError(resp.Body)
	case 429:
		return ErrRateLimitExceeded
	default:
		return err
	}
}
