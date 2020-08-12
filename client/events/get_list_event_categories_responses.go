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

// GetListEventCategoriesReader is a Reader for the GetListEventCategories structure.
type GetListEventCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListEventCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListEventCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListEventCategoriesOK creates a GetListEventCategoriesOK with default headers values
func NewGetListEventCategoriesOK() *GetListEventCategoriesOK {
	return &GetListEventCategoriesOK{}
}

/*GetListEventCategoriesOK handles this case with default header values.

IEnumerable_of_EventCategoryResource resource returned
*/
type GetListEventCategoriesOK struct {
	Payload []*models.EventCategoryResource
}

func (o *GetListEventCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/events/categories][%d] getListEventCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetListEventCategoriesOK) GetPayload() []*models.EventCategoryResource {
	return o.Payload
}

func (o *GetListEventCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}