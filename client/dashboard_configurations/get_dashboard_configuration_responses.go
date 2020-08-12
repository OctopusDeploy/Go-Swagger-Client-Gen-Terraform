// Code generated by go-swagger; DO NOT EDIT.

package dashboard_configurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetDashboardConfigurationReader is a Reader for the GetDashboardConfiguration structure.
type GetDashboardConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDashboardConfigurationOK creates a GetDashboardConfigurationOK with default headers values
func NewGetDashboardConfigurationOK() *GetDashboardConfigurationOK {
	return &GetDashboardConfigurationOK{}
}

/*GetDashboardConfigurationOK handles this case with default header values.

DashboardConfigurationResource resource returned
*/
type GetDashboardConfigurationOK struct {
	Payload *models.DashboardConfigurationResource
}

func (o *GetDashboardConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /api/dashboardconfiguration][%d] getDashboardConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetDashboardConfigurationOK) GetPayload() *models.DashboardConfigurationResource {
	return o.Payload
}

func (o *GetDashboardConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardConfigurationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
