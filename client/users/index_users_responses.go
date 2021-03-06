// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexUsersReader is a Reader for the IndexUsers structure.
type IndexUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexUsersOK creates a IndexUsersOK with default headers values
func NewIndexUsersOK() *IndexUsersOK {
	return &IndexUsersOK{}
}

/*IndexUsersOK handles this case with default header values.

ResourceCollection_of_UserResource resource returned
*/
type IndexUsersOK struct {
	Payload *models.ResourceCollectionOfUserResource
}

func (o *IndexUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/users][%d] indexUsersOK  %+v", 200, o.Payload)
}

func (o *IndexUsersOK) GetPayload() *models.ResourceCollectionOfUserResource {
	return o.Payload
}

func (o *IndexUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfUserResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
