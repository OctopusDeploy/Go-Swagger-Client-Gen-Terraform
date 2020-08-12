// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAccountByIDSpacesReader is a Reader for the GetAccountByIDSpaces structure.
type GetAccountByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountByIDSpacesOK creates a GetAccountByIDSpacesOK with default headers values
func NewGetAccountByIDSpacesOK() *GetAccountByIDSpacesOK {
	return &GetAccountByIDSpacesOK{}
}

/*GetAccountByIDSpacesOK handles this case with default header values.

Load AccountResource by ID
*/
type GetAccountByIDSpacesOK struct {
	Payload *models.AccountResource
}

func (o *GetAccountByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/accounts/{id}][%d] getAccountByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetAccountByIDSpacesOK) GetPayload() *models.AccountResource {
	return o.Payload
}

func (o *GetAccountByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountByIDSpacesBadRequest creates a GetAccountByIDSpacesBadRequest with default headers values
func NewGetAccountByIDSpacesBadRequest() *GetAccountByIDSpacesBadRequest {
	return &GetAccountByIDSpacesBadRequest{}
}

/*GetAccountByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetAccountByIDSpacesBadRequest struct {
}

func (o *GetAccountByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/accounts/{id}][%d] getAccountByIdSpacesBadRequest ", 400)
}

func (o *GetAccountByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
