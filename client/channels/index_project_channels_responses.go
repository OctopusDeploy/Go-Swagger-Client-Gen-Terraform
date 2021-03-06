// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexProjectChannelsReader is a Reader for the IndexProjectChannels structure.
type IndexProjectChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexProjectChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexProjectChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIndexProjectChannelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexProjectChannelsOK creates a IndexProjectChannelsOK with default headers values
func NewIndexProjectChannelsOK() *IndexProjectChannelsOK {
	return &IndexProjectChannelsOK{}
}

/*IndexProjectChannelsOK handles this case with default header values.

ResourceCollection_of_ChannelResource resource returned
*/
type IndexProjectChannelsOK struct {
	Payload *models.ResourceCollectionOfChannelResource
}

func (o *IndexProjectChannelsOK) Error() string {
	return fmt.Sprintf("[GET /api/projects/{id}/channels][%d] indexProjectChannelsOK  %+v", 200, o.Payload)
}

func (o *IndexProjectChannelsOK) GetPayload() *models.ResourceCollectionOfChannelResource {
	return o.Payload
}

func (o *IndexProjectChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfChannelResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIndexProjectChannelsBadRequest creates a IndexProjectChannelsBadRequest with default headers values
func NewIndexProjectChannelsBadRequest() *IndexProjectChannelsBadRequest {
	return &IndexProjectChannelsBadRequest{}
}

/*IndexProjectChannelsBadRequest handles this case with default header values.

No id parameter was provided.
*/
type IndexProjectChannelsBadRequest struct {
}

func (o *IndexProjectChannelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/projects/{id}/channels][%d] indexProjectChannelsBadRequest ", 400)
}

func (o *IndexProjectChannelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
