// Code generated by go-swagger; DO NOT EDIT.

package invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetInvitationByIDReader is a Reader for the GetInvitationByID structure.
type GetInvitationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvitationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvitationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInvitationByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvitationByIDOK creates a GetInvitationByIDOK with default headers values
func NewGetInvitationByIDOK() *GetInvitationByIDOK {
	return &GetInvitationByIDOK{}
}

/*GetInvitationByIDOK handles this case with default header values.

Load InvitationResource by ID
*/
type GetInvitationByIDOK struct {
	Payload *models.InvitationResource
}

func (o *GetInvitationByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/users/invitations/{id}][%d] getInvitationByIdOK  %+v", 200, o.Payload)
}

func (o *GetInvitationByIDOK) GetPayload() *models.InvitationResource {
	return o.Payload
}

func (o *GetInvitationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvitationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvitationByIDBadRequest creates a GetInvitationByIDBadRequest with default headers values
func NewGetInvitationByIDBadRequest() *GetInvitationByIDBadRequest {
	return &GetInvitationByIDBadRequest{}
}

/*GetInvitationByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetInvitationByIDBadRequest struct {
}

func (o *GetInvitationByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/users/invitations/{id}][%d] getInvitationByIdBadRequest ", 400)
}

func (o *GetInvitationByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
