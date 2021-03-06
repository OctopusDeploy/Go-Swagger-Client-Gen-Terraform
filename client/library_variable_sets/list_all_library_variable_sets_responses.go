// Code generated by go-swagger; DO NOT EDIT.

package library_variable_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListAllLibraryVariableSetsReader is a Reader for the ListAllLibraryVariableSets structure.
type ListAllLibraryVariableSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllLibraryVariableSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllLibraryVariableSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllLibraryVariableSetsOK creates a ListAllLibraryVariableSetsOK with default headers values
func NewListAllLibraryVariableSetsOK() *ListAllLibraryVariableSetsOK {
	return &ListAllLibraryVariableSetsOK{}
}

/*ListAllLibraryVariableSetsOK handles this case with default header values.

Load all LibraryVariableSetResource
*/
type ListAllLibraryVariableSetsOK struct {
	Payload []*models.LibraryVariableSetResource
}

func (o *ListAllLibraryVariableSetsOK) Error() string {
	return fmt.Sprintf("[GET /api/libraryvariablesets/all][%d] listAllLibraryVariableSetsOK  %+v", 200, o.Payload)
}

func (o *ListAllLibraryVariableSetsOK) GetPayload() []*models.LibraryVariableSetResource {
	return o.Payload
}

func (o *ListAllLibraryVariableSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
