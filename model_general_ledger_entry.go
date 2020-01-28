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

type GeneralLedgerEntry struct {
	// The id property for the Dynamics 365 Business Central generalLedgerEntry entity
	Id int32 `json:"id,omitempty"`
	// The postingDate property for the Dynamics 365 Business Central generalLedgerEntry entity
	PostingDate time.Time `json:"postingDate,omitempty"`
	// The documentNumber property for the Dynamics 365 Business Central generalLedgerEntry entity
	DocumentNumber string `json:"documentNumber,omitempty"`
	// The documentType property for the Dynamics 365 Business Central generalLedgerEntry entity
	DocumentType string `json:"documentType,omitempty"`
	// The accountId property for the Dynamics 365 Business Central generalLedgerEntry entity
	AccountId string `json:"accountId,omitempty"`
	// The accountNumber property for the Dynamics 365 Business Central generalLedgerEntry entity
	AccountNumber string `json:"accountNumber,omitempty"`
	// The description property for the Dynamics 365 Business Central generalLedgerEntry entity
	Description string `json:"description,omitempty"`
	// The debitAmount property for the Dynamics 365 Business Central generalLedgerEntry entity
	DebitAmount float64 `json:"debitAmount,omitempty"`
	// The creditAmount property for the Dynamics 365 Business Central generalLedgerEntry entity
	CreditAmount float64 `json:"creditAmount,omitempty"`
	Dimensions []Dimensiontype `json:"dimensions,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central generalLedgerEntry entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	Account *Account `json:"account,omitempty"`
}
