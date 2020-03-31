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

type CompaniesApiJournalsApiJournalLinesRequest struct {
	// The id property for the Dynamics 365 Business Central journalLine entity
	Id string `json:"id,omitempty"`
	// The journalDisplayName property for the Dynamics 365 Business Central journalLine entity
	JournalDisplayName string `json:"journalDisplayName,omitempty"`
	// The lineNumber property for the Dynamics 365 Business Central journalLine entity
	LineNumber int32 `json:"lineNumber,omitempty"`
	// The accountType property for the Dynamics 365 Business Central journalLine entity
	AccountType string `json:"accountType,omitempty"`
	// The accountId property for the Dynamics 365 Business Central journalLine entity
	AccountId string `json:"accountId,omitempty"`
	// The accountNumber property for the Dynamics 365 Business Central journalLine entity
	AccountNumber string `json:"accountNumber,omitempty"`
	// The postingDate property for the Dynamics 365 Business Central journalLine entity
	PostingDate *time.Time `json:"postingDate,omitempty"`
	// The documentNumber property for the Dynamics 365 Business Central journalLine entity
	DocumentNumber string `json:"documentNumber,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central journalLine entity
	ExternalDocumentNumber string `json:"externalDocumentNumber,omitempty"`
	// The amount property for the Dynamics 365 Business Central journalLine entity
	Amount float64 `json:"amount,omitempty"`
	// The description property for the Dynamics 365 Business Central journalLine entity
	Description string `json:"description,omitempty"`
	// The comment property for the Dynamics 365 Business Central journalLine entity
	Comment string `json:"comment,omitempty"`
	Dimensions []Dimensiontype `json:"dimensions,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central journalLine entity
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
