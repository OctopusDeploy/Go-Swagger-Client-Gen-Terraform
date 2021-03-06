// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListAllEnvironmentsSpacesReader is a Reader for the ListAllEnvironmentsSpaces structure.
type ListAllEnvironmentsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllEnvironmentsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllEnvironmentsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllEnvironmentsSpacesOK creates a ListAllEnvironmentsSpacesOK with default headers values
func NewListAllEnvironmentsSpacesOK() *ListAllEnvironmentsSpacesOK {
	return &ListAllEnvironmentsSpacesOK{}
}

/*ListAllEnvironmentsSpacesOK handles this case with default header values.

Load all EnvironmentResource
*/
type ListAllEnvironmentsSpacesOK struct {
	Payload []*models.EnvironmentResource
}

func (o *ListAllEnvironmentsSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/environments/all][%d] listAllEnvironmentsSpacesOK  %+v", 200, o.Payload)
}

func (o *ListAllEnvironmentsSpacesOK) GetPayload() []*models.EnvironmentResource {
	return o.Payload
}

func (o *ListAllEnvironmentsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
