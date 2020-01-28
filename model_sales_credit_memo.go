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

type SalesCreditMemo struct {
	// The id property for the Dynamics 365 Business Central salesCreditMemo entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central salesCreditMemo entity
	Number string `json:"number,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central salesCreditMemo entity
	ExternalDocumentNumber string `json:"externalDocumentNumber,omitempty"`
	// The creditMemoDate property for the Dynamics 365 Business Central salesCreditMemo entity
	CreditMemoDate time.Time `json:"creditMemoDate,omitempty"`
	// The dueDate property for the Dynamics 365 Business Central salesCreditMemo entity
	DueDate time.Time `json:"dueDate,omitempty"`
	// The customerId property for the Dynamics 365 Business Central salesCreditMemo entity
	CustomerId string `json:"customerId,omitempty"`
	// The contactId property for the Dynamics 365 Business Central salesCreditMemo entity
	ContactId string `json:"contactId,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central salesCreditMemo entity
	CustomerNumber string `json:"customerNumber,omitempty"`
	// The customerName property for the Dynamics 365 Business Central salesCreditMemo entity
	CustomerName string `json:"customerName,omitempty"`
	// The billToName property for the Dynamics 365 Business Central salesCreditMemo entity
	BillToName string `json:"billToName,omitempty"`
	// The billToCustomerId property for the Dynamics 365 Business Central salesCreditMemo entity
	BillToCustomerId string `json:"billToCustomerId,omitempty"`
	// The billToCustomerNumber property for the Dynamics 365 Business Central salesCreditMemo entity
	BillToCustomerNumber string `json:"billToCustomerNumber,omitempty"`
	SellingPostalAddress *Postaladdresstype `json:"sellingPostalAddress,omitempty"`
	BillingPostalAddress *Postaladdresstype `json:"billingPostalAddress,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central salesCreditMemo entity
	CurrencyId string `json:"currencyId,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central salesCreditMemo entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The paymentTermsId property for the Dynamics 365 Business Central salesCreditMemo entity
	PaymentTermsId string `json:"paymentTermsId,omitempty"`
	// The shipmentMethodId property for the Dynamics 365 Business Central salesCreditMemo entity
	ShipmentMethodId string `json:"shipmentMethodId,omitempty"`
	// The salesperson property for the Dynamics 365 Business Central salesCreditMemo entity
	Salesperson string `json:"salesperson,omitempty"`
	// The pricesIncludeTax property for the Dynamics 365 Business Central salesCreditMemo entity
	PricesIncludeTax bool `json:"pricesIncludeTax,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesCreditMemo entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesCreditMemo entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central salesCreditMemo entity
	TotalAmountExcludingTax float64 `json:"totalAmountExcludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesCreditMemo entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central salesCreditMemo entity
	TotalAmountIncludingTax float64 `json:"totalAmountIncludingTax,omitempty"`
	// The status property for the Dynamics 365 Business Central salesCreditMemo entity
	Status string `json:"status,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central salesCreditMemo entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The invoiceId property for the Dynamics 365 Business Central salesCreditMemo entity
	InvoiceId string `json:"invoiceId,omitempty"`
	// The invoiceNumber property for the Dynamics 365 Business Central salesCreditMemo entity
	InvoiceNumber string `json:"invoiceNumber,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central salesCreditMemo entity
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// The email property for the Dynamics 365 Business Central salesCreditMemo entity
	Email string `json:"email,omitempty"`
	SalesCreditMemoLines []SalesCreditMemoLine `json:"salesCreditMemoLines,omitempty"`
	PdfDocument []PdfDocument `json:"pdfDocument,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
	ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
}
