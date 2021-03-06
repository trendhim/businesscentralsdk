/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package businesscentralsdk

type CompaniesApiSalesQuotesApiSalesQuoteLinesSalesQuoteLineIdApiRequest struct {
	// The id property for the Dynamics 365 Business Central salesQuoteLine entity
	Id string `json:"id,omitempty"`
	// The documentId property for the Dynamics 365 Business Central salesQuoteLine entity
	DocumentId string `json:"documentId,omitempty"`
	// The sequence property for the Dynamics 365 Business Central salesQuoteLine entity
	Sequence int32 `json:"sequence,omitempty"`
	// The itemId property for the Dynamics 365 Business Central salesQuoteLine entity
	ItemId string `json:"itemId,omitempty"`
	// The accountId property for the Dynamics 365 Business Central salesQuoteLine entity
	AccountId string `json:"accountId,omitempty"`
	// The lineType property for the Dynamics 365 Business Central salesQuoteLine entity
	LineType string `json:"lineType,omitempty"`
	LineDetails *Documentlineobjectdetailstype `json:"lineDetails,omitempty"`
	// The description property for the Dynamics 365 Business Central salesQuoteLine entity
	Description string `json:"description,omitempty"`
	// The unitOfMeasureId property for the Dynamics 365 Business Central salesQuoteLine entity
	UnitOfMeasureId string `json:"unitOfMeasureId,omitempty"`
	UnitOfMeasure *Unitofmeasuretype `json:"unitOfMeasure,omitempty"`
	// The unitPrice property for the Dynamics 365 Business Central salesQuoteLine entity
	UnitPrice float64 `json:"unitPrice,omitempty"`
	// The quantity property for the Dynamics 365 Business Central salesQuoteLine entity
	Quantity float64 `json:"quantity,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesQuoteLine entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountPercent property for the Dynamics 365 Business Central salesQuoteLine entity
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesQuoteLine entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The amountExcludingTax property for the Dynamics 365 Business Central salesQuoteLine entity
	AmountExcludingTax float64 `json:"amountExcludingTax,omitempty"`
	// The taxCode property for the Dynamics 365 Business Central salesQuoteLine entity
	TaxCode string `json:"taxCode,omitempty"`
	// The taxPercent property for the Dynamics 365 Business Central salesQuoteLine entity
	TaxPercent float64 `json:"taxPercent,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesQuoteLine entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The amountIncludingTax property for the Dynamics 365 Business Central salesQuoteLine entity
	AmountIncludingTax float64 `json:"amountIncludingTax,omitempty"`
	// The netAmount property for the Dynamics 365 Business Central salesQuoteLine entity
	NetAmount float64 `json:"netAmount,omitempty"`
	// The netTaxAmount property for the Dynamics 365 Business Central salesQuoteLine entity
	NetTaxAmount float64 `json:"netTaxAmount,omitempty"`
	// The netAmountIncludingTax property for the Dynamics 365 Business Central salesQuoteLine entity
	NetAmountIncludingTax float64 `json:"netAmountIncludingTax,omitempty"`
}
