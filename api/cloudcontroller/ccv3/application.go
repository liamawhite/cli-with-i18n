package ccv3

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3/internal"
)

type AppLifecycleType string

const (
	BuildpackAppLifecycleType AppLifecycleType = "buildpack"
	DockerAppLifecycleType    AppLifecycleType = "docker"
)

// Application represents a Cloud Controller V3 Application.
type Application struct {
	Name          string        `json:"name,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	GUID          string        `json:"guid,omitempty"`
	State         string        `json:"state,omitempty"`
	Lifecycle     AppLifecycle  `json:"lifecycle,omitempty"`
}

type AppLifecycle struct {
	Type AppLifecycleType `json:"type,omitempty"`
	Data AppLifecycleData `json:"data,omitempty"`
}

type AppLifecycleData struct {
	Buildpacks []string `json:"buildpacks,omitempty"`
}

func (a Application) MarshalJSON() ([]byte, error) {
	var ccApp struct {
		Name          string                 `json:"name,omitempty"`
		Relationships Relationships          `json:"relationships,omitempty"`
		Lifecycle     map[string]interface{} `json:"lifecycle,omitempty"`
	}

	ccApp.Name = a.Name
	ccApp.Relationships = a.Relationships

	switch a.Lifecycle.Type {
	case BuildpackAppLifecycleType:
		if len(a.Lifecycle.Data.Buildpacks) > 0 {
			switch a.Lifecycle.Data.Buildpacks[0] {
			case "default", "null":
				ccApp.Lifecycle = map[string]interface{}{
					"type": a.Lifecycle.Type,
					"data": map[string]interface{}{
						"buildpacks": nil,
					},
				}
			default:
				ccApp.Lifecycle = map[string]interface{}{
					"type": a.Lifecycle.Type,
					"data": map[string]interface{}{
						"buildpacks": a.Lifecycle.Data.Buildpacks,
					},
				}
			}
		}
	case DockerAppLifecycleType:
		ccApp.Lifecycle = map[string]interface{}{
			"type": a.Lifecycle.Type,
			"data": map[string]interface{}{},
		}
	}

	return json.Marshal(ccApp)
}

// DropletRelationship represents the relationship between a V3 Droplet and its
// V3 Application
type DropletRelationship struct {
	Relationship Relationship `json:"data"`
	Links        APILinks     `json:"links"`
}

// GetApplications lists applications with optional filters.
func (client *Client) GetApplications(query url.Values) ([]Application, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetAppsRequest,
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullAppsList []Application
	warnings, err := client.paginate(request, Application{}, func(item interface{}) error {
		if app, ok := item.(Application); ok {
			fullAppsList = append(fullAppsList, app)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   Application{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullAppsList, warnings, err
}

// CreateApplication creates an application with the given settings
func (client *Client) CreateApplication(app Application) (Application, Warnings, error) {
	bodyBytes, err := json.Marshal(app)
	if err != nil {
		return Application{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostApplicationRequest,
		Body:        bytes.NewReader(bodyBytes),
	})
	if err != nil {
		return Application{}, nil, err
	}

	var responseApp Application
	response := cloudcontroller.Response{
		Result: &responseApp,
	}
	err = client.connection.Make(request, &response)

	return responseApp, response.Warnings, err
}

func (client *Client) DeleteApplication(appGUID string) (string, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteApplicationRequest,
		URIParams:   internal.Params{"app_guid": appGUID},
	})
	if err != nil {
		return "", nil, err
	}

	response := cloudcontroller.Response{}
	err = client.connection.Make(request, &response)

	return response.ResourceLocationURL, response.Warnings, err
}

// UpdateApplication updates an application with the given settings
func (client *Client) UpdateApplication(app Application) (Application, Warnings, error) {
	bodyBytes, err := json.Marshal(app)
	if err != nil {
		return Application{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PatchApplicationRequest,
		Body:        bytes.NewReader(bodyBytes),
		URIParams:   map[string]string{"app_guid": app.GUID},
	})
	if err != nil {
		return Application{}, nil, err
	}

	var responseApp Application
	response := cloudcontroller.Response{
		Result: &responseApp,
	}
	err = client.connection.Make(request, &response)

	return responseApp, response.Warnings, err
}

func (client *Client) SetApplicationDroplet(appGUID string, dropletGUID string) (Relationship, Warnings, error) {
	relationship := Relationship{GUID: dropletGUID}
	bodyBytes, err := json.Marshal(relationship)
	if err != nil {
		return Relationship{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PatchApplicationCurrentDropletRequest,
		URIParams:   map[string]string{"app_guid": appGUID},
		Body:        bytes.NewReader(bodyBytes),
	})
	if err != nil {
		return Relationship{}, nil, err
	}

	var responseRelationship Relationship
	response := cloudcontroller.Response{
		Result: &responseRelationship,
	}
	err = client.connection.Make(request, &response)

	return responseRelationship, response.Warnings, err
}

func (client *Client) StopApplication(appGUID string) (Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostApplicationStopRequest,
		URIParams:   map[string]string{"app_guid": appGUID},
	})
	if err != nil {
		return nil, err
	}

	response := cloudcontroller.Response{}
	err = client.connection.Make(request, &response)

	return response.Warnings, err
}

func (client *Client) StartApplication(appGUID string) (Application, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostApplicationStartRequest,
		URIParams:   map[string]string{"app_guid": appGUID},
	})
	if err != nil {
		return Application{}, nil, err
	}

	var responseApp Application
	response := cloudcontroller.Response{
		Result: &responseApp,
	}
	err = client.connection.Make(request, &response)

	return responseApp, response.Warnings, err
}
