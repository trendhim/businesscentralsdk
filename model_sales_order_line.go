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

type SalesOrderLine struct {
	// The id property for the Dynamics 365 Business Central salesOrderLine entity
	Id string `json:"id,omitempty"`
	// The documentId property for the Dynamics 365 Business Central salesOrderLine entity
	DocumentId string `json:"documentId,omitempty"`
	// The sequence property for the Dynamics 365 Business Central salesOrderLine entity
	Sequence int32 `json:"sequence,omitempty"`
	// The itemId property for the Dynamics 365 Business Central salesOrderLine entity
	ItemId string `json:"itemId,omitempty"`
	// The accountId property for the Dynamics 365 Business Central salesOrderLine entity
	AccountId string `json:"accountId,omitempty"`
	// The lineType property for the Dynamics 365 Business Central salesOrderLine entity
	LineType string `json:"lineType,omitempty"`
	LineDetails *Documentlineobjectdetailstype `json:"lineDetails,omitempty"`
	// The description property for the Dynamics 365 Business Central salesOrderLine entity
	Description string `json:"description,omitempty"`
	// The unitOfMeasureId property for the Dynamics 365 Business Central salesOrderLine entity
	UnitOfMeasureId string `json:"unitOfMeasureId,omitempty"`
	UnitOfMeasure *Unitofmeasuretype `json:"unitOfMeasure,omitempty"`
	// The quantity property for the Dynamics 365 Business Central salesOrderLine entity
	Quantity float64 `json:"quantity,omitempty"`
	// The unitPrice property for the Dynamics 365 Business Central salesOrderLine entity
	UnitPrice float64 `json:"unitPrice,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesOrderLine entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountPercent property for the Dynamics 365 Business Central salesOrderLine entity
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesOrderLine entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The amountExcludingTax property for the Dynamics 365 Business Central salesOrderLine entity
	AmountExcludingTax float64 `json:"amountExcludingTax,omitempty"`
	// The taxCode property for the Dynamics 365 Business Central salesOrderLine entity
	TaxCode string `json:"taxCode,omitempty"`
	// The taxPercent property for the Dynamics 365 Business Central salesOrderLine entity
	TaxPercent float64 `json:"taxPercent,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesOrderLine entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The amountIncludingTax property for the Dynamics 365 Business Central salesOrderLine entity
	AmountIncludingTax float64 `json:"amountIncludingTax,omitempty"`
	// The invoiceDiscountAllocation property for the Dynamics 365 Business Central salesOrderLine entity
	InvoiceDiscountAllocation float64 `json:"invoiceDiscountAllocation,omitempty"`
	// The netAmount property for the Dynamics 365 Business Central salesOrderLine entity
	NetAmount float64 `json:"netAmount,omitempty"`
	// The netTaxAmount property for the Dynamics 365 Business Central salesOrderLine entity
	NetTaxAmount float64 `json:"netTaxAmount,omitempty"`
	// The netAmountIncludingTax property for the Dynamics 365 Business Central salesOrderLine entity
	NetAmountIncludingTax float64 `json:"netAmountIncludingTax,omitempty"`
	// The shipmentDate property for the Dynamics 365 Business Central salesOrderLine entity
	ShipmentDate time.Time `json:"shipmentDate,omitempty"`
	// The shippedQuantity property for the Dynamics 365 Business Central salesOrderLine entity
	ShippedQuantity float64 `json:"shippedQuantity,omitempty"`
	// The invoicedQuantity property for the Dynamics 365 Business Central salesOrderLine entity
	InvoicedQuantity float64 `json:"invoicedQuantity,omitempty"`
	// The invoiceQuantity property for the Dynamics 365 Business Central salesOrderLine entity
	InvoiceQuantity float64 `json:"invoiceQuantity,omitempty"`
	// The shipQuantity property for the Dynamics 365 Business Central salesOrderLine entity
	ShipQuantity float64 `json:"shipQuantity,omitempty"`
	Item *Item `json:"item,omitempty"`
	Account *Account `json:"account,omitempty"`
}
