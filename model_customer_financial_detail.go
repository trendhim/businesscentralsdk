/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package businesscentralsdk

type CustomerFinancialDetail struct {
	// The id property for the Dynamics 365 Business Central customerFinancialDetail entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central customerFinancialDetail entity
	Number string `json:"number,omitempty"`
	// The balance property for the Dynamics 365 Business Central customerFinancialDetail entity
	Balance float64 `json:"balance,omitempty"`
	// The totalSalesExcludingTax property for the Dynamics 365 Business Central customerFinancialDetail entity
	TotalSalesExcludingTax float64 `json:"totalSalesExcludingTax,omitempty"`
	// The overdueAmount property for the Dynamics 365 Business Central customerFinancialDetail entity
	OverdueAmount float64 `json:"overdueAmount,omitempty"`
}
