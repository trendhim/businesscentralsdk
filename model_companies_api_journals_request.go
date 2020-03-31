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

type CompaniesApiJournalsRequest struct {
	// The id property for the Dynamics 365 Business Central journal entity
	Id string `json:"id,omitempty"`
	// The code property for the Dynamics 365 Business Central journal entity
	Code string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central journal entity
	DisplayName string `json:"displayName,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central journal entity
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The balancingAccountId property for the Dynamics 365 Business Central journal entity
	BalancingAccountId string `json:"balancingAccountId,omitempty"`
	// The balancingAccountNumber property for the Dynamics 365 Business Central journal entity
	BalancingAccountNumber string `json:"balancingAccountNumber,omitempty"`
}
