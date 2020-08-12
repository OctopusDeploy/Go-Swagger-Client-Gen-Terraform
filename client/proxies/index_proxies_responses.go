// Code generated by go-swagger; DO NOT EDIT.

package proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexProxiesReader is a Reader for the IndexProxies structure.
type IndexProxiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexProxiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexProxiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexProxiesOK creates a IndexProxiesOK with default headers values
func NewIndexProxiesOK() *IndexProxiesOK {
	return &IndexProxiesOK{}
}

/*IndexProxiesOK handles this case with default header values.

ResourceCollection_of_ProxyResource resource returned
*/
type IndexProxiesOK struct {
	Payload *models.ResourceCollectionOfProxyResource
}

func (o *IndexProxiesOK) Error() string {
	return fmt.Sprintf("[GET /api/proxies][%d] indexProxiesOK  %+v", 200, o.Payload)
}

func (o *IndexProxiesOK) GetPayload() *models.ResourceCollectionOfProxyResource {
	return o.Payload
}

func (o *IndexProxiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfProxyResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}