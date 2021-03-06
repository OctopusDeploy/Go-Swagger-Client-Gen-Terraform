// Code generated by go-swagger; DO NOT EDIT.

package interruptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexInterruptionsSpacesReader is a Reader for the IndexInterruptionsSpaces structure.
type IndexInterruptionsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexInterruptionsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexInterruptionsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexInterruptionsSpacesOK creates a IndexInterruptionsSpacesOK with default headers values
func NewIndexInterruptionsSpacesOK() *IndexInterruptionsSpacesOK {
	return &IndexInterruptionsSpacesOK{}
}

/*IndexInterruptionsSpacesOK handles this case with default header values.

ResourceCollection_of_InterruptionResource resource returned
*/
type IndexInterruptionsSpacesOK struct {
	Payload *models.ResourceCollectionOfInterruptionResource
}

func (o *IndexInterruptionsSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/interruptions][%d] indexInterruptionsSpacesOK  %+v", 200, o.Payload)
}

func (o *IndexInterruptionsSpacesOK) GetPayload() *models.ResourceCollectionOfInterruptionResource {
	return o.Payload
}

func (o *IndexInterruptionsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfInterruptionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
