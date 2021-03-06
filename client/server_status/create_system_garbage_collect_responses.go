// Code generated by go-swagger; DO NOT EDIT.

package server_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateSystemGarbageCollectReader is a Reader for the CreateSystemGarbageCollect structure.
type CreateSystemGarbageCollectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSystemGarbageCollectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSystemGarbageCollectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSystemGarbageCollectOK creates a CreateSystemGarbageCollectOK with default headers values
func NewCreateSystemGarbageCollectOK() *CreateSystemGarbageCollectOK {
	return &CreateSystemGarbageCollectOK{}
}

/*CreateSystemGarbageCollectOK handles this case with default header values.

OK - Default
*/
type CreateSystemGarbageCollectOK struct {
}

func (o *CreateSystemGarbageCollectOK) Error() string {
	return fmt.Sprintf("[POST /api/serverstatus/gc-collect][%d] createSystemGarbageCollectOK ", 200)
}

func (o *CreateSystemGarbageCollectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
