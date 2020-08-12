// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateFeedReader is a Reader for the CreateFeed structure.
type CreateFeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFeedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFeedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFeedCreated creates a CreateFeedCreated with default headers values
func NewCreateFeedCreated() *CreateFeedCreated {
	return &CreateFeedCreated{}
}

/*CreateFeedCreated handles this case with default header values.

FeedResource Created
*/
type CreateFeedCreated struct {
	Payload *models.FeedResource
}

func (o *CreateFeedCreated) Error() string {
	return fmt.Sprintf("[POST /api/feeds][%d] createFeedCreated  %+v", 201, o.Payload)
}

func (o *CreateFeedCreated) GetPayload() *models.FeedResource {
	return o.Payload
}

func (o *CreateFeedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedBadRequest creates a CreateFeedBadRequest with default headers values
func NewCreateFeedBadRequest() *CreateFeedBadRequest {
	return &CreateFeedBadRequest{}
}

/*CreateFeedBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateFeedBadRequest struct {
}

func (o *CreateFeedBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/feeds][%d] createFeedBadRequest ", 400)
}

func (o *CreateFeedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}