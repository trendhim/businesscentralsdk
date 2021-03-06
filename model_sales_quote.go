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

type SalesQuote struct {
	// The id property for the Dynamics 365 Business Central salesQuote entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central salesQuote entity
	Number string `json:"number,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central salesQuote entity
	ExternalDocumentNumber string `json:"externalDocumentNumber,omitempty"`
	// The documentDate property for the Dynamics 365 Business Central salesQuote entity
	DocumentDate time.Time `json:"documentDate,omitempty"`
	// The dueDate property for the Dynamics 365 Business Central salesQuote entity
	DueDate time.Time `json:"dueDate,omitempty"`
	// The customerId property for the Dynamics 365 Business Central salesQuote entity
	CustomerId string `json:"customerId,omitempty"`
	// The contactId property for the Dynamics 365 Business Central salesQuote entity
	ContactId string `json:"contactId,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central salesQuote entity
	CustomerNumber string `json:"customerNumber,omitempty"`
	// The customerName property for the Dynamics 365 Business Central salesQuote entity
	CustomerName string `json:"customerName,omitempty"`
	// The billToName property for the Dynamics 365 Business Central salesQuote entity
	BillToName string `json:"billToName,omitempty"`
	// The billToCustomerId property for the Dynamics 365 Business Central salesQuote entity
	BillToCustomerId string `json:"billToCustomerId,omitempty"`
	// The billToCustomerNumber property for the Dynamics 365 Business Central salesQuote entity
	BillToCustomerNumber string `json:"billToCustomerNumber,omitempty"`
	// The shipToName property for the Dynamics 365 Business Central salesQuote entity
	ShipToName string `json:"shipToName,omitempty"`
	// The shipToContact property for the Dynamics 365 Business Central salesQuote entity
	ShipToContact string `json:"shipToContact,omitempty"`
	SellingPostalAddress *Postaladdresstype `json:"sellingPostalAddress,omitempty"`
	BillingPostalAddress *Postaladdresstype `json:"billingPostalAddress,omitempty"`
	ShippingPostalAddress *Postaladdresstype `json:"shippingPostalAddress,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central salesQuote entity
	CurrencyId string `json:"currencyId,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central salesQuote entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The paymentTermsId property for the Dynamics 365 Business Central salesQuote entity
	PaymentTermsId string `json:"paymentTermsId,omitempty"`
	// The shipmentMethodId property for the Dynamics 365 Business Central salesQuote entity
	ShipmentMethodId string `json:"shipmentMethodId,omitempty"`
	// The salesperson property for the Dynamics 365 Business Central salesQuote entity
	Salesperson string `json:"salesperson,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesQuote entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central salesQuote entity
	TotalAmountExcludingTax float64 `json:"totalAmountExcludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesQuote entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central salesQuote entity
	TotalAmountIncludingTax float64 `json:"totalAmountIncludingTax,omitempty"`
	// The status property for the Dynamics 365 Business Central salesQuote entity
	Status string `json:"status,omitempty"`
	// The sentDate property for the Dynamics 365 Business Central salesQuote entity
	SentDate time.Time `json:"sentDate,omitempty"`
	// The validUntilDate property for the Dynamics 365 Business Central salesQuote entity
	ValidUntilDate time.Time `json:"validUntilDate,omitempty"`
	// The acceptedDate property for the Dynamics 365 Business Central salesQuote entity
	AcceptedDate time.Time `json:"acceptedDate,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central salesQuote entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central salesQuote entity
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// The email property for the Dynamics 365 Business Central salesQuote entity
	Email string `json:"email,omitempty"`
	SalesQuoteLines []SalesQuoteLine `json:"salesQuoteLines,omitempty"`
	PdfDocument []PdfDocument `json:"pdfDocument,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
	ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
}
