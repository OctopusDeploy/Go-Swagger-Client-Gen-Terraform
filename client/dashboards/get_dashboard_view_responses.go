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

// GetDashboardViewReader is a Reader for the GetDashboardView structure.
type GetDashboardViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDashboardViewOK creates a GetDashboardViewOK with default headers values
func NewGetDashboardViewOK() *GetDashboardViewOK {
	return &GetDashboardViewOK{}
}

/*GetDashboardViewOK handles this case with default header values.

DashboardResource resource returned
*/
type GetDashboardViewOK struct {
	Payload *models.DashboardResource
}

func (o *GetDashboardViewOK) Error() string {
	return fmt.Sprintf("[GET /api/dashboard][%d] getDashboardViewOK  %+v", 200, o.Payload)
}

func (o *GetDashboardViewOK) GetPayload() *models.DashboardResource {
	return o.Payload
}

func (o *GetDashboardViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
