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

type Employee struct {
	// The id property for the Dynamics 365 Business Central employee entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central employee entity
	Number string `json:"number,omitempty"`
	// The displayName property for the Dynamics 365 Business Central employee entity
	DisplayName string `json:"displayName,omitempty"`
	// The givenName property for the Dynamics 365 Business Central employee entity
	GivenName string `json:"givenName,omitempty"`
	// The middleName property for the Dynamics 365 Business Central employee entity
	MiddleName string `json:"middleName,omitempty"`
	// The surname property for the Dynamics 365 Business Central employee entity
	Surname string `json:"surname,omitempty"`
	// The jobTitle property for the Dynamics 365 Business Central employee entity
	JobTitle string `json:"jobTitle,omitempty"`
	Address *Postaladdresstype `json:"address,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central employee entity
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// The mobilePhone property for the Dynamics 365 Business Central employee entity
	MobilePhone string `json:"mobilePhone,omitempty"`
	// The email property for the Dynamics 365 Business Central employee entity
	Email string `json:"email,omitempty"`
	// The personalEmail property for the Dynamics 365 Business Central employee entity
	PersonalEmail string `json:"personalEmail,omitempty"`
	// The employmentDate property for the Dynamics 365 Business Central employee entity
	EmploymentDate time.Time `json:"employmentDate,omitempty"`
	// The terminationDate property for the Dynamics 365 Business Central employee entity
	TerminationDate time.Time `json:"terminationDate,omitempty"`
	// The status property for the Dynamics 365 Business Central employee entity
	Status string `json:"status,omitempty"`
	// The birthDate property for the Dynamics 365 Business Central employee entity
	BirthDate time.Time `json:"birthDate,omitempty"`
	// The statisticsGroupCode property for the Dynamics 365 Business Central employee entity
	StatisticsGroupCode string `json:"statisticsGroupCode,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central employee entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	Picture []Picture `json:"picture,omitempty"`
	DefaultDimensions []DefaultDimensions `json:"defaultDimensions,omitempty"`
	TimeRegistrationEntries []TimeRegistrationEntry `json:"timeRegistrationEntries,omitempty"`
}
