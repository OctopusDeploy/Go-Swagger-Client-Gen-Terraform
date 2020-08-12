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

// GetLibraryVariableSetByIDSpacesReader is a Reader for the GetLibraryVariableSetByIDSpaces structure.
type GetLibraryVariableSetByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibraryVariableSetByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLibraryVariableSetByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLibraryVariableSetByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLibraryVariableSetByIDSpacesOK creates a GetLibraryVariableSetByIDSpacesOK with default headers values
func NewGetLibraryVariableSetByIDSpacesOK() *GetLibraryVariableSetByIDSpacesOK {
	return &GetLibraryVariableSetByIDSpacesOK{}
}

/*GetLibraryVariableSetByIDSpacesOK handles this case with default header values.

Load LibraryVariableSetResource by ID
*/
type GetLibraryVariableSetByIDSpacesOK struct {
	Payload *models.LibraryVariableSetResource
}

func (o *GetLibraryVariableSetByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/libraryvariablesets/{id}][%d] getLibraryVariableSetByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetLibraryVariableSetByIDSpacesOK) GetPayload() *models.LibraryVariableSetResource {
	return o.Payload
}

func (o *GetLibraryVariableSetByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LibraryVariableSetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryVariableSetByIDSpacesBadRequest creates a GetLibraryVariableSetByIDSpacesBadRequest with default headers values
func NewGetLibraryVariableSetByIDSpacesBadRequest() *GetLibraryVariableSetByIDSpacesBadRequest {
	return &GetLibraryVariableSetByIDSpacesBadRequest{}
}

/*GetLibraryVariableSetByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetLibraryVariableSetByIDSpacesBadRequest struct {
}

func (o *GetLibraryVariableSetByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/libraryvariablesets/{id}][%d] getLibraryVariableSetByIdSpacesBadRequest ", 400)
}

func (o *GetLibraryVariableSetByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
