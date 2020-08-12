// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetOctopusServerNodeDetailsReader is a Reader for the GetOctopusServerNodeDetails structure.
type GetOctopusServerNodeDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOctopusServerNodeDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOctopusServerNodeDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOctopusServerNodeDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOctopusServerNodeDetailsOK creates a GetOctopusServerNodeDetailsOK with default headers values
func NewGetOctopusServerNodeDetailsOK() *GetOctopusServerNodeDetailsOK {
	return &GetOctopusServerNodeDetailsOK{}
}

/*GetOctopusServerNodeDetailsOK handles this case with default header values.

OctopusServerNodeDetailsResource resource returned
*/
type GetOctopusServerNodeDetailsOK struct {
	Payload *models.OctopusServerNodeDetailsResource
}

func (o *GetOctopusServerNodeDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/octopusservernodes/{id}/details][%d] getOctopusServerNodeDetailsOK  %+v", 200, o.Payload)
}

func (o *GetOctopusServerNodeDetailsOK) GetPayload() *models.OctopusServerNodeDetailsResource {
	return o.Payload
}

func (o *GetOctopusServerNodeDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OctopusServerNodeDetailsResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOctopusServerNodeDetailsBadRequest creates a GetOctopusServerNodeDetailsBadRequest with default headers values
func NewGetOctopusServerNodeDetailsBadRequest() *GetOctopusServerNodeDetailsBadRequest {
	return &GetOctopusServerNodeDetailsBadRequest{}
}

/*GetOctopusServerNodeDetailsBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetOctopusServerNodeDetailsBadRequest struct {
}

func (o *GetOctopusServerNodeDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/octopusservernodes/{id}/details][%d] getOctopusServerNodeDetailsBadRequest ", 400)
}

func (o *GetOctopusServerNodeDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
