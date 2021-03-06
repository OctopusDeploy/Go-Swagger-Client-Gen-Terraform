// Code generated by go-swagger; DO NOT EDIT.

package action_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetActionTemplateCategoriesReader is a Reader for the GetActionTemplateCategories structure.
type GetActionTemplateCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionTemplateCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionTemplateCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetActionTemplateCategoriesOK creates a GetActionTemplateCategoriesOK with default headers values
func NewGetActionTemplateCategoriesOK() *GetActionTemplateCategoriesOK {
	return &GetActionTemplateCategoriesOK{}
}

/*GetActionTemplateCategoriesOK handles this case with default header values.

IEnumerable_of_ActionTemplateCategoryResource resource returned
*/
type GetActionTemplateCategoriesOK struct {
	Payload []*models.ActionTemplateCategoryResource
}

func (o *GetActionTemplateCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/actiontemplates/categories][%d] getActionTemplateCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetActionTemplateCategoriesOK) GetPayload() []*models.ActionTemplateCategoryResource {
	return o.Payload
}

func (o *GetActionTemplateCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
