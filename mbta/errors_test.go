package mbta

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func errorsBadRequestSetup() (io.ReadCloser, BadRequestError) {
	body := `{"errors":[{"code":"not_found","source":{"parameter":"id"},"status":"404","title":"Resource Not Found"}],"jsonapi":{"version":"1.0"}}`
	expected := BadRequestError{
		SourceParameter: "id",
		Detail:          "Resource Not Found",
		Code:            "not_found",
	}
	return ioutil.NopCloser(strings.NewReader(body)), expected
}

func Test_BadRequestError_Error(t *testing.T) {
	_, testErr := errorsBadRequestSetup()
	expected := `parameter "id" caused error [not_found]: (Resource Not Found)`
	actual := testErr.Error()
	equals(t, expected, actual)
}

func Test_getBadRequestError(t *testing.T) {
	body, expected := errorsBadRequestSetup()
	actual := getBadRequestError(body)
	equals(t, expected, actual)
}

func Test_getSpecialError(t *testing.T) {
	body, expectedBadRequest := errorsBadRequestSetup()
	inputErr := errors.New("hi")
	testCases := []struct {
		statusCode  int
		expectedErr error
	}{
		{400, expectedBadRequest},
		{403, ErrForbidden},
		{404, expectedBadRequest},
		{429, ErrRateLimitExceeded},
		{500, inputErr},
	}

	for _, testCase := range testCases {
		resp := http.Response{StatusCode: testCase.statusCode, Body: body}
		actual := getSpecialError(&resp, inputErr)
		equals(t, testCase.expectedErr, actual)

		// Make a new body every time
		body.Close()
		body, _ = errorsBadRequestSetup()
	}
}
