// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAzureWebSitesListSpacesReader is a Reader for the GetAzureWebSitesListSpaces structure.
type GetAzureWebSitesListSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureWebSitesListSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureWebSitesListSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAzureWebSitesListSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAzureWebSitesListSpacesOK creates a GetAzureWebSitesListSpacesOK with default headers values
func NewGetAzureWebSitesListSpacesOK() *GetAzureWebSitesListSpacesOK {
	return &GetAzureWebSitesListSpacesOK{}
}

/*GetAzureWebSitesListSpacesOK handles this case with default header values.

IEnumerable_of_AzureWebSiteResource_of_AzureWebSitesListAction resource returned
*/
type GetAzureWebSitesListSpacesOK struct {
	Payload []*models.AzureWebSiteResourceOfAzureWebSitesListAction
}

func (o *GetAzureWebSitesListSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/accounts/{id}/websites][%d] getAzureWebSitesListSpacesOK  %+v", 200, o.Payload)
}

func (o *GetAzureWebSitesListSpacesOK) GetPayload() []*models.AzureWebSiteResourceOfAzureWebSitesListAction {
	return o.Payload
}

func (o *GetAzureWebSitesListSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureWebSitesListSpacesBadRequest creates a GetAzureWebSitesListSpacesBadRequest with default headers values
func NewGetAzureWebSitesListSpacesBadRequest() *GetAzureWebSitesListSpacesBadRequest {
	return &GetAzureWebSitesListSpacesBadRequest{}
}

/*GetAzureWebSitesListSpacesBadRequest handles this case with default header values.

Account must be an Azure account.
No id parameter was provided.
This operation is not supported using Azure Management Certificate authentication. Please try again using an Azure Service Account.
*/
type GetAzureWebSitesListSpacesBadRequest struct {
}

func (o *GetAzureWebSitesListSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/accounts/{id}/websites][%d] getAzureWebSitesListSpacesBadRequest ", 400)
}

func (o *GetAzureWebSitesListSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
