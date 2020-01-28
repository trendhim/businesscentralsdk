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

type CompaniesApiSalesInvoicesApiSalesInvoiceLinesRequest struct {
	// The id property for the Dynamics 365 Business Central salesInvoiceLine entity
	Id string `json:"id,omitempty"`
	// The documentId property for the Dynamics 365 Business Central salesInvoiceLine entity
	DocumentId string `json:"documentId,omitempty"`
	// The sequence property for the Dynamics 365 Business Central salesInvoiceLine entity
	Sequence int32 `json:"sequence,omitempty"`
	// The itemId property for the Dynamics 365 Business Central salesInvoiceLine entity
	ItemId string `json:"itemId,omitempty"`
	// The accountId property for the Dynamics 365 Business Central salesInvoiceLine entity
	AccountId string `json:"accountId,omitempty"`
	// The lineType property for the Dynamics 365 Business Central salesInvoiceLine entity
	LineType string `json:"lineType,omitempty"`
	LineDetails *Documentlineobjectdetailstype `json:"lineDetails,omitempty"`
	// The description property for the Dynamics 365 Business Central salesInvoiceLine entity
	Description string `json:"description,omitempty"`
	// The unitOfMeasureId property for the Dynamics 365 Business Central salesInvoiceLine entity
	UnitOfMeasureId string `json:"unitOfMeasureId,omitempty"`
	UnitOfMeasure *Unitofmeasuretype `json:"unitOfMeasure,omitempty"`
	// The unitPrice property for the Dynamics 365 Business Central salesInvoiceLine entity
	UnitPrice float64 `json:"unitPrice,omitempty"`
	// The quantity property for the Dynamics 365 Business Central salesInvoiceLine entity
	Quantity float64 `json:"quantity,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesInvoiceLine entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountPercent property for the Dynamics 365 Business Central salesInvoiceLine entity
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesInvoiceLine entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The amountExcludingTax property for the Dynamics 365 Business Central salesInvoiceLine entity
	AmountExcludingTax float64 `json:"amountExcludingTax,omitempty"`
	// The taxCode property for the Dynamics 365 Business Central salesInvoiceLine entity
	TaxCode string `json:"taxCode,omitempty"`
	// The taxPercent property for the Dynamics 365 Business Central salesInvoiceLine entity
	TaxPercent float64 `json:"taxPercent,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesInvoiceLine entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The amountIncludingTax property for the Dynamics 365 Business Central salesInvoiceLine entity
	AmountIncludingTax float64 `json:"amountIncludingTax,omitempty"`
	// The invoiceDiscountAllocation property for the Dynamics 365 Business Central salesInvoiceLine entity
	InvoiceDiscountAllocation float64 `json:"invoiceDiscountAllocation,omitempty"`
	// The netAmount property for the Dynamics 365 Business Central salesInvoiceLine entity
	NetAmount float64 `json:"netAmount,omitempty"`
	// The netTaxAmount property for the Dynamics 365 Business Central salesInvoiceLine entity
	NetTaxAmount float64 `json:"netTaxAmount,omitempty"`
	// The netAmountIncludingTax property for the Dynamics 365 Business Central salesInvoiceLine entity
	NetAmountIncludingTax float64 `json:"netAmountIncludingTax,omitempty"`
	// The shipmentDate property for the Dynamics 365 Business Central salesInvoiceLine entity
	ShipmentDate time.Time `json:"shipmentDate,omitempty"`
}
