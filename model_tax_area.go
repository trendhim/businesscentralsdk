/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package businesscentralsdk

import (
	"time"
)

type TaxArea struct {
	// The id property for the Dynamics 365 Business Central taxArea entity
	Id string `json:"id,omitempty"`
	// The code property for the Dynamics 365 Business Central taxArea entity
	Code string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central taxArea entity
	DisplayName string `json:"displayName,omitempty"`
	// The taxType property for the Dynamics 365 Business Central taxArea entity
	TaxType string `json:"taxType,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central taxArea entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
}