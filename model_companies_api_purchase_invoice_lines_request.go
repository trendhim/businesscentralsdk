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

type CompaniesApiPurchaseInvoiceLinesRequest struct {
	// The id property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	Id string `json:"id,omitempty"`
	// The documentId property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	DocumentId string `json:"documentId,omitempty"`
	// The sequence property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	Sequence int32 `json:"sequence,omitempty"`
	// The itemId property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	ItemId string `json:"itemId,omitempty"`
	// The accountId property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	AccountId string `json:"accountId,omitempty"`
	// The lineType property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	LineType string `json:"lineType,omitempty"`
	LineDetails *Documentlineobjectdetailstype `json:"lineDetails,omitempty"`
	// The description property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	Description string `json:"description,omitempty"`
	UnitOfMeasure *Unitofmeasuretype `json:"unitOfMeasure,omitempty"`
	// The unitCost property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	UnitCost float64 `json:"unitCost,omitempty"`
	// The quantity property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	Quantity float64 `json:"quantity,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountPercent property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The amountExcludingTax property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	AmountExcludingTax float64 `json:"amountExcludingTax,omitempty"`
	// The taxCode property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	TaxCode string `json:"taxCode,omitempty"`
	// The taxPercent property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	TaxPercent float64 `json:"taxPercent,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The amountIncludingTax property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	AmountIncludingTax float64 `json:"amountIncludingTax,omitempty"`
	// The invoiceDiscountAllocation property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	InvoiceDiscountAllocation float64 `json:"invoiceDiscountAllocation,omitempty"`
	// The netAmount property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	NetAmount float64 `json:"netAmount,omitempty"`
	// The netTaxAmount property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	NetTaxAmount float64 `json:"netTaxAmount,omitempty"`
	// The netAmountIncludingTax property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	NetAmountIncludingTax float64 `json:"netAmountIncludingTax,omitempty"`
	// The expectedReceiptDate property for the Dynamics 365 Business Central purchaseInvoiceLine entity
	ExpectedReceiptDate time.Time `json:"expectedReceiptDate,omitempty"`
}
