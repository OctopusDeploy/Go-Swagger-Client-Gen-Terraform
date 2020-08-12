// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetListEventDocumentTypesReader is a Reader for the GetListEventDocumentTypes structure.
type GetListEventDocumentTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListEventDocumentTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListEventDocumentTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListEventDocumentTypesOK creates a GetListEventDocumentTypesOK with default headers values
func NewGetListEventDocumentTypesOK() *GetListEventDocumentTypesOK {
	return &GetListEventDocumentTypesOK{}
}

/*GetListEventDocumentTypesOK handles this case with default header values.

IEnumerable_of_DocumentTypeDocument resource returned
*/
type GetListEventDocumentTypesOK struct {
	Payload []*models.DocumentTypeDocument
}

func (o *GetListEventDocumentTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/events/documenttypes][%d] getListEventDocumentTypesOK  %+v", 200, o.Payload)
}

func (o *GetListEventDocumentTypesOK) GetPayload() []*models.DocumentTypeDocument {
	return o.Payload
}

func (o *GetListEventDocumentTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}