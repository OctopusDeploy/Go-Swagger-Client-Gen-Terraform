// Code generated by go-swagger; DO NOT EDIT.

package server_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetExtensionStatsReader is a Reader for the GetExtensionStats structure.
type GetExtensionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionStatsOK creates a GetExtensionStatsOK with default headers values
func NewGetExtensionStatsOK() *GetExtensionStatsOK {
	return &GetExtensionStatsOK{}
}

/*GetExtensionStatsOK handles this case with default header values.

IEnumerable_of_ExtensionsInfoResource resource returned
*/
type GetExtensionStatsOK struct {
	Payload []*models.ExtensionsInfoResource
}

func (o *GetExtensionStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/serverstatus/extensions][%d] getExtensionStatsOK  %+v", 200, o.Payload)
}

func (o *GetExtensionStatsOK) GetPayload() []*models.ExtensionsInfoResource {
	return o.Payload
}

func (o *GetExtensionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
