package requests

type FinInst struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstCode               *string `json:"FinInstCode"`
	FinInstBranchCode         *string `json:"FinInstBranchCode"`
	FinInstFullCode           *string `json:"FinInstFullCode"`
	FinInstName               *string `json:"FinInstName"`
	FinInstBranchName         *string `json:"FinInstBranchName"`
	SWIFTCode                 *string `json:"SWIFTCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         *string `json:"FinInstControlKey"`
	FinInstAccountName        *string `json:"FinInstAccountName"`
	FinInstAccount            *string `json:"FinInstAccount"`
	HouseBank                 *string `json:"HouseBank"`
	HouseBankAccount          *string `json:"HouseBankAccount"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
