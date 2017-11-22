package ccv2

import (
	"encoding/json"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv2/internal"
)

// ServiceInstanceType is the type of the Service Instance.
type ServiceInstanceType string

const (
	// UserProvidedService is a Service Instance that is created by a user.
	UserProvidedService ServiceInstanceType = "user_provided_service_instance"

	// ManagedService is a Service Instance that is managed by a service broker.
	ManagedService ServiceInstanceType = "managed_service_instance"
)

// ServiceInstance represents a Cloud Controller Service Instance.
type ServiceInstance struct {
	GUID      string
	Name      string
	SpaceGUID string
	Type      ServiceInstanceType
}

// UnmarshalJSON helps unmarshal a Cloud Controller Service Instance response.
func (serviceInstance *ServiceInstance) UnmarshalJSON(data []byte) error {
	var ccServiceInstance struct {
		Metadata internal.Metadata
		Entity   struct {
			Name      string `json:"name"`
			SpaceGUID string `json:"space_guid"`
			Type      string `json:"type"`
		}
	}
	err := json.Unmarshal(data, &ccServiceInstance)
	if err != nil {
		return err
	}

	serviceInstance.GUID = ccServiceInstance.Metadata.GUID
	serviceInstance.Name = ccServiceInstance.Entity.Name
	serviceInstance.SpaceGUID = ccServiceInstance.Entity.SpaceGUID
	serviceInstance.Type = ServiceInstanceType(ccServiceInstance.Entity.Type)
	return nil
}

// UserProvided returns true if the Service Instance is a user provided
// service.
func (serviceInstance ServiceInstance) UserProvided() bool {
	return serviceInstance.Type == UserProvidedService
}

// Managed returns true if the Service Instance is a managed service.
func (serviceInstance ServiceInstance) Managed() bool {
	return serviceInstance.Type == ManagedService
}

// GetServiceInstance returns the service instance with the given GUID. This
// service can be either a managed or user provided.
func (client *Client) GetServiceInstance(serviceInstanceGUID string) (ServiceInstance, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceInstanceRequest,
		URIParams:   Params{"service_instance_guid": serviceInstanceGUID},
	})
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	var serviceInstance ServiceInstance
	response := cloudcontroller.Response{
		Result: &serviceInstance,
	}

	err = client.connection.Make(request, &response)
	return serviceInstance, response.Warnings, err
}

// GetServiceInstances returns back a list of *managed* Service Instances based
// off of the provided queries.
func (client *Client) GetServiceInstances(queries ...Query) ([]ServiceInstance, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceInstancesRequest,
		Query:       FormatQueryParameters(queries),
	})
	if err != nil {
		return nil, nil, err
	}

	var fullInstancesList []ServiceInstance
	warnings, err := client.paginate(request, ServiceInstance{}, func(item interface{}) error {
		if instance, ok := item.(ServiceInstance); ok {
			fullInstancesList = append(fullInstancesList, instance)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceInstance{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullInstancesList, warnings, err
}

// GetSpaceServiceInstances returns back a list of Service Instances based off
// of the space and queries provided. User provided services will be included
// if includeUserProvidedServices is set to true.
func (client *Client) GetSpaceServiceInstances(spaceGUID string, includeUserProvidedServices bool, queries ...Query) ([]ServiceInstance, Warnings, error) {
	query := FormatQueryParameters(queries)

	if includeUserProvidedServices {
		query.Add("return_user_provided_service_instances", "true")
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetSpaceServiceInstancesRequest,
		URIParams:   map[string]string{"guid": spaceGUID},
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullInstancesList []ServiceInstance
	warnings, err := client.paginate(request, ServiceInstance{}, func(item interface{}) error {
		if instance, ok := item.(ServiceInstance); ok {
			fullInstancesList = append(fullInstancesList, instance)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceInstance{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullInstancesList, warnings, err
}
