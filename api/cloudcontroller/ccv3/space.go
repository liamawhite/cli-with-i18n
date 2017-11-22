package ccv3

import (
	"net/url"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3/internal"
)

// Space represents a Cloud Controller V3 Space.
type Space struct {
	Name string `json:"name"`
	GUID string `json:"guid"`
}

// GetSpaces lists spaces with optional filters.
func (client *Client) GetSpaces(query url.Values) ([]Space, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetSpacesRequest,
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullSpacesList []Space
	warnings, err := client.paginate(request, Space{}, func(item interface{}) error {
		if space, ok := item.(Space); ok {
			fullSpacesList = append(fullSpacesList, space)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   Space{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullSpacesList, warnings, err
}
