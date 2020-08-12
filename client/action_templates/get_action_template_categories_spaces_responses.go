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

// GetActionTemplateCategoriesSpacesReader is a Reader for the GetActionTemplateCategoriesSpaces structure.
type GetActionTemplateCategoriesSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionTemplateCategoriesSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionTemplateCategoriesSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetActionTemplateCategoriesSpacesOK creates a GetActionTemplateCategoriesSpacesOK with default headers values
func NewGetActionTemplateCategoriesSpacesOK() *GetActionTemplateCategoriesSpacesOK {
	return &GetActionTemplateCategoriesSpacesOK{}
}

/*GetActionTemplateCategoriesSpacesOK handles this case with default header values.

IEnumerable_of_ActionTemplateCategoryResource resource returned
*/
type GetActionTemplateCategoriesSpacesOK struct {
	Payload []*models.ActionTemplateCategoryResource
}

func (o *GetActionTemplateCategoriesSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/actiontemplates/categories][%d] getActionTemplateCategoriesSpacesOK  %+v", 200, o.Payload)
}

func (o *GetActionTemplateCategoriesSpacesOK) GetPayload() []*models.ActionTemplateCategoryResource {
	return o.Payload
}

func (o *GetActionTemplateCategoriesSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
