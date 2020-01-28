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

type Item struct {
	// The id property for the Dynamics 365 Business Central item entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central item entity
	Number string `json:"number,omitempty"`
	// The displayName property for the Dynamics 365 Business Central item entity
	DisplayName string `json:"displayName,omitempty"`
	// The type property for the Dynamics 365 Business Central item entity
	Type_ string `json:"type,omitempty"`
	// The itemCategoryId property for the Dynamics 365 Business Central item entity
	ItemCategoryId string `json:"itemCategoryId,omitempty"`
	// The itemCategoryCode property for the Dynamics 365 Business Central item entity
	ItemCategoryCode string `json:"itemCategoryCode,omitempty"`
	// The blocked property for the Dynamics 365 Business Central item entity
	Blocked bool `json:"blocked,omitempty"`
	// The baseUnitOfMeasureId property for the Dynamics 365 Business Central item entity
	BaseUnitOfMeasureId string `json:"baseUnitOfMeasureId,omitempty"`
	BaseUnitOfMeasure *Unitofmeasuretype `json:"baseUnitOfMeasure,omitempty"`
	// The gtin property for the Dynamics 365 Business Central item entity
	Gtin string `json:"gtin,omitempty"`
	// The inventory property for the Dynamics 365 Business Central item entity
	Inventory float64 `json:"inventory,omitempty"`
	// The unitPrice property for the Dynamics 365 Business Central item entity
	UnitPrice float64 `json:"unitPrice,omitempty"`
	// The priceIncludesTax property for the Dynamics 365 Business Central item entity
	PriceIncludesTax bool `json:"priceIncludesTax,omitempty"`
	// The unitCost property for the Dynamics 365 Business Central item entity
	UnitCost float64 `json:"unitCost,omitempty"`
	// The taxGroupId property for the Dynamics 365 Business Central item entity
	TaxGroupId string `json:"taxGroupId,omitempty"`
	// The taxGroupCode property for the Dynamics 365 Business Central item entity
	TaxGroupCode string `json:"taxGroupCode,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central item entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	Picture []Picture `json:"picture,omitempty"`
	DefaultDimensions []DefaultDimensions `json:"defaultDimensions,omitempty"`
	ItemCategory *ItemCategory `json:"itemCategory,omitempty"`
}
