// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetSubscriptionByIDSpacesReader is a Reader for the GetSubscriptionByIDSpaces structure.
type GetSubscriptionByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubscriptionByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSubscriptionByIDSpacesOK creates a GetSubscriptionByIDSpacesOK with default headers values
func NewGetSubscriptionByIDSpacesOK() *GetSubscriptionByIDSpacesOK {
	return &GetSubscriptionByIDSpacesOK{}
}

/*GetSubscriptionByIDSpacesOK handles this case with default header values.

Load SubscriptionResource by ID
*/
type GetSubscriptionByIDSpacesOK struct {
	Payload *models.SubscriptionResource
}

func (o *GetSubscriptionByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/subscriptions/{id}][%d] getSubscriptionByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionByIDSpacesOK) GetPayload() *models.SubscriptionResource {
	return o.Payload
}

func (o *GetSubscriptionByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionByIDSpacesBadRequest creates a GetSubscriptionByIDSpacesBadRequest with default headers values
func NewGetSubscriptionByIDSpacesBadRequest() *GetSubscriptionByIDSpacesBadRequest {
	return &GetSubscriptionByIDSpacesBadRequest{}
}

/*GetSubscriptionByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetSubscriptionByIDSpacesBadRequest struct {
}

func (o *GetSubscriptionByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/subscriptions/{id}][%d] getSubscriptionByIdSpacesBadRequest ", 400)
}

func (o *GetSubscriptionByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}