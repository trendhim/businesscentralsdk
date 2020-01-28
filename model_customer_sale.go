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

type CustomerSale struct {
	// The customerId property for the Dynamics 365 Business Central customerSale entity
	CustomerId string `json:"customerId,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central customerSale entity
	CustomerNumber string `json:"customerNumber,omitempty"`
	// The name property for the Dynamics 365 Business Central customerSale entity
	Name string `json:"name,omitempty"`
	// The totalSalesAmount property for the Dynamics 365 Business Central customerSale entity
	TotalSalesAmount float64 `json:"totalSalesAmount,omitempty"`
	// The dateFilter_FilterOnly property for the Dynamics 365 Business Central customerSale entity
	DateFilterFilterOnly time.Time `json:"dateFilter_FilterOnly,omitempty"`
}
