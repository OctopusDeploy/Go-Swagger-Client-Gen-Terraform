// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetListTaskTypesReader is a Reader for the GetListTaskTypes structure.
type GetListTaskTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListTaskTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListTaskTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListTaskTypesOK creates a GetListTaskTypesOK with default headers values
func NewGetListTaskTypesOK() *GetListTaskTypesOK {
	return &GetListTaskTypesOK{}
}

/*GetListTaskTypesOK handles this case with default header values.

IEnumerable_of_TaskTypeResource resource returned
*/
type GetListTaskTypesOK struct {
	Payload []*models.TaskTypeResource
}

func (o *GetListTaskTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/tasks/tasktypes][%d] getListTaskTypesOK  %+v", 200, o.Payload)
}

func (o *GetListTaskTypesOK) GetPayload() []*models.TaskTypeResource {
	return o.Payload
}

func (o *GetListTaskTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
