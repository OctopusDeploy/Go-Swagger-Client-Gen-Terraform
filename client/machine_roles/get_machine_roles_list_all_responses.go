// Code generated by go-swagger; DO NOT EDIT.

package machine_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetMachineRolesListAllReader is a Reader for the GetMachineRolesListAll structure.
type GetMachineRolesListAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineRolesListAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineRolesListAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMachineRolesListAllOK creates a GetMachineRolesListAllOK with default headers values
func NewGetMachineRolesListAllOK() *GetMachineRolesListAllOK {
	return &GetMachineRolesListAllOK{}
}

/*GetMachineRolesListAllOK handles this case with default header values.

IEnumerable_of_String resource returned
*/
type GetMachineRolesListAllOK struct {
	Payload []string
}

func (o *GetMachineRolesListAllOK) Error() string {
	return fmt.Sprintf("[GET /api/machineroles/all][%d] getMachineRolesListAllOK  %+v", 200, o.Payload)
}

func (o *GetMachineRolesListAllOK) GetPayload() []string {
	return o.Payload
}

func (o *GetMachineRolesListAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}