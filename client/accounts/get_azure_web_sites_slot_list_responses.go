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

// GetAzureWebSitesSlotListReader is a Reader for the GetAzureWebSitesSlotList structure.
type GetAzureWebSitesSlotListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureWebSitesSlotListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureWebSitesSlotListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAzureWebSitesSlotListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAzureWebSitesSlotListOK creates a GetAzureWebSitesSlotListOK with default headers values
func NewGetAzureWebSitesSlotListOK() *GetAzureWebSitesSlotListOK {
	return &GetAzureWebSitesSlotListOK{}
}

/*GetAzureWebSitesSlotListOK handles this case with default header values.

IEnumerable_of_AzureWebSiteSlotResource resource returned
*/
type GetAzureWebSitesSlotListOK struct {
	Payload []*models.AzureWebSiteSlotResource
}

func (o *GetAzureWebSitesSlotListOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/{resourceGroupName}/websites/{webSiteName}/slots][%d] getAzureWebSitesSlotListOK  %+v", 200, o.Payload)
}

func (o *GetAzureWebSitesSlotListOK) GetPayload() []*models.AzureWebSiteSlotResource {
	return o.Payload
}

func (o *GetAzureWebSitesSlotListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureWebSitesSlotListBadRequest creates a GetAzureWebSitesSlotListBadRequest with default headers values
func NewGetAzureWebSitesSlotListBadRequest() *GetAzureWebSitesSlotListBadRequest {
	return &GetAzureWebSitesSlotListBadRequest{}
}

/*GetAzureWebSitesSlotListBadRequest handles this case with default header values.

Account must be an Azure account.
No id parameter was provided.
No resourceGroupName parameter was provided.
No webSiteName parameter was provided.
This operation is not supported using Azure Management Certificate authentication. Please try again using an Azure Service Account.
*/
type GetAzureWebSitesSlotListBadRequest struct {
}

func (o *GetAzureWebSitesSlotListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/{resourceGroupName}/websites/{webSiteName}/slots][%d] getAzureWebSitesSlotListBadRequest ", 400)
}

func (o *GetAzureWebSitesSlotListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}