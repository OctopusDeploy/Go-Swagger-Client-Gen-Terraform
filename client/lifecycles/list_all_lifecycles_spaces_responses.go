// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListAllLifecyclesSpacesReader is a Reader for the ListAllLifecyclesSpaces structure.
type ListAllLifecyclesSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllLifecyclesSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllLifecyclesSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllLifecyclesSpacesOK creates a ListAllLifecyclesSpacesOK with default headers values
func NewListAllLifecyclesSpacesOK() *ListAllLifecyclesSpacesOK {
	return &ListAllLifecyclesSpacesOK{}
}

/*ListAllLifecyclesSpacesOK handles this case with default header values.

Load all LifecycleResource
*/
type ListAllLifecyclesSpacesOK struct {
	Payload []*models.LifecycleResource
}

func (o *ListAllLifecyclesSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/lifecycles/all][%d] listAllLifecyclesSpacesOK  %+v", 200, o.Payload)
}

func (o *ListAllLifecyclesSpacesOK) GetPayload() []*models.LifecycleResource {
	return o.Payload
}

func (o *ListAllLifecyclesSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
