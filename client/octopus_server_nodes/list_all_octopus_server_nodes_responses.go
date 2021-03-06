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

// ListAllOctopusServerNodesReader is a Reader for the ListAllOctopusServerNodes structure.
type ListAllOctopusServerNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllOctopusServerNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllOctopusServerNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllOctopusServerNodesOK creates a ListAllOctopusServerNodesOK with default headers values
func NewListAllOctopusServerNodesOK() *ListAllOctopusServerNodesOK {
	return &ListAllOctopusServerNodesOK{}
}

/*ListAllOctopusServerNodesOK handles this case with default header values.

Load all OctopusServerNodeResource
*/
type ListAllOctopusServerNodesOK struct {
	Payload []*models.OctopusServerNodeResource
}

func (o *ListAllOctopusServerNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/octopusservernodes/all][%d] listAllOctopusServerNodesOK  %+v", 200, o.Payload)
}

func (o *ListAllOctopusServerNodesOK) GetPayload() []*models.OctopusServerNodeResource {
	return o.Payload
}

func (o *ListAllOctopusServerNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
