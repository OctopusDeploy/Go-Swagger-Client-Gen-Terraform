// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetDynamicDashboardViewReader is a Reader for the GetDynamicDashboardView structure.
type GetDynamicDashboardViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDynamicDashboardViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDynamicDashboardViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDynamicDashboardViewOK creates a GetDynamicDashboardViewOK with default headers values
func NewGetDynamicDashboardViewOK() *GetDynamicDashboardViewOK {
	return &GetDynamicDashboardViewOK{}
}

/*GetDynamicDashboardViewOK handles this case with default header values.

DashboardResource resource returned
*/
type GetDynamicDashboardViewOK struct {
	Payload *models.DashboardResource
}

func (o *GetDynamicDashboardViewOK) Error() string {
	return fmt.Sprintf("[GET /api/dashboard/dynamic][%d] getDynamicDashboardViewOK  %+v", 200, o.Payload)
}

func (o *GetDynamicDashboardViewOK) GetPayload() *models.DashboardResource {
	return o.Payload
}

func (o *GetDynamicDashboardViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
