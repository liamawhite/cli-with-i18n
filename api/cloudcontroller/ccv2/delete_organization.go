// generated from codetemplates/delete_async_by_guid.go.template

package ccv2

import (
	"net/url"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv2/internal"
)

// DeleteOrganization deletes the Organization associated with the provided
// GUID. It will return the Cloud Controller job that is assigned to the
// Organization deletion.
func (client *Client) DeleteOrganization(guid string) (Job, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteOrganizationRequest,
		URIParams:   Params{"organization_guid": guid},
		Query: url.Values{
			"recursive": {"true"},
			"async":     {"true"},
		},
	})
	if err != nil {
		return Job{}, nil, err
	}

	var job Job
	response := cloudcontroller.Response{
		Result: &job,
	}

	err = client.connection.Make(request, &response)
	return job, response.Warnings, err
}
