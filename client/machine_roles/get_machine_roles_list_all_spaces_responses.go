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

// GetMachineRolesListAllSpacesReader is a Reader for the GetMachineRolesListAllSpaces structure.
type GetMachineRolesListAllSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineRolesListAllSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineRolesListAllSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMachineRolesListAllSpacesOK creates a GetMachineRolesListAllSpacesOK with default headers values
func NewGetMachineRolesListAllSpacesOK() *GetMachineRolesListAllSpacesOK {
	return &GetMachineRolesListAllSpacesOK{}
}

/*GetMachineRolesListAllSpacesOK handles this case with default header values.

IEnumerable_of_String resource returned
*/
type GetMachineRolesListAllSpacesOK struct {
	Payload []string
}

func (o *GetMachineRolesListAllSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/machineroles/all][%d] getMachineRolesListAllSpacesOK  %+v", 200, o.Payload)
}

func (o *GetMachineRolesListAllSpacesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetMachineRolesListAllSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}