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

type UnitOfMeasure struct {
	// The id property for the Dynamics 365 Business Central unitOfMeasure entity
	Id string `json:"id,omitempty"`
	// The code property for the Dynamics 365 Business Central unitOfMeasure entity
	Code string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central unitOfMeasure entity
	DisplayName string `json:"displayName,omitempty"`
	// The internationalStandardCode property for the Dynamics 365 Business Central unitOfMeasure entity
	InternationalStandardCode string `json:"internationalStandardCode,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central unitOfMeasure entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
}