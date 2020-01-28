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

type SalesOrder struct {
	// The id property for the Dynamics 365 Business Central salesOrder entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central salesOrder entity
	Number string `json:"number,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central salesOrder entity
	ExternalDocumentNumber string `json:"externalDocumentNumber,omitempty"`
	// The orderDate property for the Dynamics 365 Business Central salesOrder entity
	OrderDate time.Time `json:"orderDate,omitempty"`
	// The customerId property for the Dynamics 365 Business Central salesOrder entity
	CustomerId string `json:"customerId,omitempty"`
	// The contactId property for the Dynamics 365 Business Central salesOrder entity
	ContactId string `json:"contactId,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central salesOrder entity
	CustomerNumber string `json:"customerNumber,omitempty"`
	// The customerName property for the Dynamics 365 Business Central salesOrder entity
	CustomerName string `json:"customerName,omitempty"`
	// The billToName property for the Dynamics 365 Business Central salesOrder entity
	BillToName string `json:"billToName,omitempty"`
	// The billToCustomerId property for the Dynamics 365 Business Central salesOrder entity
	BillToCustomerId string `json:"billToCustomerId,omitempty"`
	// The billToCustomerNumber property for the Dynamics 365 Business Central salesOrder entity
	BillToCustomerNumber string `json:"billToCustomerNumber,omitempty"`
	// The shipToName property for the Dynamics 365 Business Central salesOrder entity
	ShipToName string `json:"shipToName,omitempty"`
	// The shipToContact property for the Dynamics 365 Business Central salesOrder entity
	ShipToContact string `json:"shipToContact,omitempty"`
	SellingPostalAddress *Postaladdresstype `json:"sellingPostalAddress,omitempty"`
	BillingPostalAddress *Postaladdresstype `json:"billingPostalAddress,omitempty"`
	ShippingPostalAddress *Postaladdresstype `json:"shippingPostalAddress,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central salesOrder entity
	CurrencyId string `json:"currencyId,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central salesOrder entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The pricesIncludeTax property for the Dynamics 365 Business Central salesOrder entity
	PricesIncludeTax bool `json:"pricesIncludeTax,omitempty"`
	// The paymentTermsId property for the Dynamics 365 Business Central salesOrder entity
	PaymentTermsId string `json:"paymentTermsId,omitempty"`
	// The shipmentMethodId property for the Dynamics 365 Business Central salesOrder entity
	ShipmentMethodId string `json:"shipmentMethodId,omitempty"`
	// The salesperson property for the Dynamics 365 Business Central salesOrder entity
	Salesperson string `json:"salesperson,omitempty"`
	// The partialShipping property for the Dynamics 365 Business Central salesOrder entity
	PartialShipping bool `json:"partialShipping,omitempty"`
	// The requestedDeliveryDate property for the Dynamics 365 Business Central salesOrder entity
	RequestedDeliveryDate time.Time `json:"requestedDeliveryDate,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesOrder entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesOrder entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central salesOrder entity
	TotalAmountExcludingTax float64 `json:"totalAmountExcludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesOrder entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central salesOrder entity
	TotalAmountIncludingTax float64 `json:"totalAmountIncludingTax,omitempty"`
	// The fullyShipped property for the Dynamics 365 Business Central salesOrder entity
	FullyShipped bool `json:"fullyShipped,omitempty"`
	// The status property for the Dynamics 365 Business Central salesOrder entity
	Status string `json:"status,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central salesOrder entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central salesOrder entity
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// The email property for the Dynamics 365 Business Central salesOrder entity
	Email string `json:"email,omitempty"`
	SalesOrderLines []SalesOrderLine `json:"salesOrderLines,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
	ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
}