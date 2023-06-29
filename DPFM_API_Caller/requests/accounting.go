package requests

type Accounting struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
