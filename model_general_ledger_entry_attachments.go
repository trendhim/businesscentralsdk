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

type GeneralLedgerEntryAttachments struct {
	// The generalLedgerEntryNumber property for the Dynamics 365 Business Central generalLedgerEntryAttachments entity
	GeneralLedgerEntryNumber int32 `json:"generalLedgerEntryNumber,omitempty"`
	// The id property for the Dynamics 365 Business Central generalLedgerEntryAttachments entity
	Id string `json:"id,omitempty"`
	// The fileName property for the Dynamics 365 Business Central generalLedgerEntryAttachments entity
	FileName string `json:"fileName,omitempty"`
	// The byteSize property for the Dynamics 365 Business Central generalLedgerEntryAttachments entity
	ByteSize int32 `json:"byteSize,omitempty"`
	// The content property for the Dynamics 365 Business Central generalLedgerEntryAttachments entity
	Content string `json:"content,omitempty"`
	// The createdDateTime property for the Dynamics 365 Business Central generalLedgerEntryAttachments entity
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	GeneralLedgerEntry *GeneralLedgerEntry `json:"generalLedgerEntry,omitempty"`
}
