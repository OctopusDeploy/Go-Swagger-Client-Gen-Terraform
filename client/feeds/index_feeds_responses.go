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

// IndexFeedsReader is a Reader for the IndexFeeds structure.
type IndexFeedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexFeedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexFeedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexFeedsOK creates a IndexFeedsOK with default headers values
func NewIndexFeedsOK() *IndexFeedsOK {
	return &IndexFeedsOK{}
}

/*IndexFeedsOK handles this case with default header values.

ResourceCollection_of_FeedResource resource returned
*/
type IndexFeedsOK struct {
	Payload *models.ResourceCollectionOfFeedResource
}

func (o *IndexFeedsOK) Error() string {
	return fmt.Sprintf("[GET /api/feeds][%d] indexFeedsOK  %+v", 200, o.Payload)
}

func (o *IndexFeedsOK) GetPayload() *models.ResourceCollectionOfFeedResource {
	return o.Payload
}

func (o *IndexFeedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfFeedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}