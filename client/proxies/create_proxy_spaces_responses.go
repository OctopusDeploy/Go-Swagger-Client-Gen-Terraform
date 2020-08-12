// Code generated by go-swagger; DO NOT EDIT.

package proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateProxySpacesReader is a Reader for the CreateProxySpaces structure.
type CreateProxySpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProxySpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProxySpacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProxySpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProxySpacesCreated creates a CreateProxySpacesCreated with default headers values
func NewCreateProxySpacesCreated() *CreateProxySpacesCreated {
	return &CreateProxySpacesCreated{}
}

/*CreateProxySpacesCreated handles this case with default header values.

ProxyResource Created
*/
type CreateProxySpacesCreated struct {
	Payload *models.ProxyResource
}

func (o *CreateProxySpacesCreated) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/proxies][%d] createProxySpacesCreated  %+v", 201, o.Payload)
}

func (o *CreateProxySpacesCreated) GetPayload() *models.ProxyResource {
	return o.Payload
}

func (o *CreateProxySpacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProxySpacesBadRequest creates a CreateProxySpacesBadRequest with default headers values
func NewCreateProxySpacesBadRequest() *CreateProxySpacesBadRequest {
	return &CreateProxySpacesBadRequest{}
}

/*CreateProxySpacesBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateProxySpacesBadRequest struct {
}

func (o *CreateProxySpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/proxies][%d] createProxySpacesBadRequest ", 400)
}

func (o *CreateProxySpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}