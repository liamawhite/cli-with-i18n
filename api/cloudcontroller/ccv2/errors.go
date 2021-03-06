package ccv2

import (
	"encoding/json"
	"net/http"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
)

// errorWrapper is the wrapper that converts responses with 4xx and 5xx status
// codes to an error.
type errorWrapper struct {
	connection cloudcontroller.Connection
}

func newErrorWrapper() *errorWrapper {
	return new(errorWrapper)
}

// Wrap wraps a Cloud Controller connection in this error handling wrapper.
func (e *errorWrapper) Wrap(innerconnection cloudcontroller.Connection) cloudcontroller.Connection {
	e.connection = innerconnection
	return e
}

// Make converts RawHTTPStatusError, which represents responses with 4xx and
// 5xx status codes, to specific errors.
func (e *errorWrapper) Make(request *cloudcontroller.Request, passedResponse *cloudcontroller.Response) error {
	err := e.connection.Make(request, passedResponse)

	if rawHTTPStatusErr, ok := err.(ccerror.RawHTTPStatusError); ok {
		return convert(rawHTTPStatusErr)
	}
	return err
}

func convert(rawHTTPStatusErr ccerror.RawHTTPStatusError) error {
	// Try to unmarshal the raw error into a CC error. If unmarshaling fails,
	// return the raw error.
	var errorResponse ccerror.V2ErrorResponse
	err := json.Unmarshal(rawHTTPStatusErr.RawResponse, &errorResponse)
	if err != nil {
		if rawHTTPStatusErr.StatusCode == http.StatusNotFound {
			return ccerror.NotFoundError{Message: string(rawHTTPStatusErr.RawResponse)}
		}
		return rawHTTPStatusErr
	}

	switch rawHTTPStatusErr.StatusCode {
	case http.StatusBadRequest: // 400
		return handleBadRequest(errorResponse)
	case http.StatusUnauthorized: // 401
		return handleUnauthorized(errorResponse)
	case http.StatusForbidden: // 403
		return ccerror.ForbiddenError{Message: errorResponse.Description}
	case http.StatusNotFound: // 404
		return ccerror.ResourceNotFoundError{Message: errorResponse.Description}
	case http.StatusUnprocessableEntity: // 422
		return ccerror.UnprocessableEntityError{Message: errorResponse.Description}
	default:
		return ccerror.V2UnexpectedResponseError{
			V2ErrorResponse: errorResponse,
			RequestIDs:      rawHTTPStatusErr.RequestIDs,
			ResponseCode:    rawHTTPStatusErr.StatusCode,
		}
	}
}

func handleBadRequest(errorResponse ccerror.V2ErrorResponse) error {
	switch errorResponse.ErrorCode {
	case "CF-AppStoppedStatsError":
		return ccerror.ApplicationStoppedStatsError{Message: errorResponse.Description}
	case "CF-InstancesError":
		return ccerror.InstancesError{Message: errorResponse.Description}
	case "CF-InvalidRelation":
		return ccerror.InvalidRelationError{Message: errorResponse.Description}
	case "CF-NotStaged":
		return ccerror.NotStagedError{Message: errorResponse.Description}
	case "CF-ServiceBindingAppServiceTaken":
		return ccerror.ServiceBindingTakenError{Message: errorResponse.Description}
	default:
		return ccerror.BadRequestError{Message: errorResponse.Description}
	}
}

func handleUnauthorized(errorResponse ccerror.V2ErrorResponse) error {
	if errorResponse.ErrorCode == "CF-InvalidAuthToken" {
		return ccerror.InvalidAuthTokenError{Message: errorResponse.Description}
	}

	return ccerror.UnauthorizedError{Message: errorResponse.Description}
}
