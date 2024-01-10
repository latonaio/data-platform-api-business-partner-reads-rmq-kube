package requests

type General struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	BusinessPartnerFullName       *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName           string  `json:"BusinessPartnerName"`
	Industry                      *string `json:"Industry"`
	LegalEntityRegistration       *string `json:"LegalEntityRegistration"`
	Country                       string  `json:"Country"`
	Language                      string  `json:"Language"`
	Currency                      string  `json:"Currency"`
	OrganizationBPName1           *string `json:"OrganizationBPName1"`
	OrganizationBPName2           *string `json:"OrganizationBPName2"`
	OrganizationBPName3           *string `json:"OrganizationBPName3"`
	OrganizationBPName4           *string `json:"OrganizationBPName4"`
	BPTag1                        *string `json:"BPTag1"`
	BPTag2                        *string `json:"BPTag2"`
	BPTag3                        *string `json:"BPTag3"`
	BPTag4                        *string `json:"BPTag4"`
	OrganizationFoundationDate    *string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string `json:"GroupBusinessPartnerName2"`
	AddressID                     *int    `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
	CertificateAuthorityChain     *string `json:"CertificateAuthorityChain"`
	UsageControlChain        	  *string `json:"UsageControlChain"`
	CreationDate                  string  `json:"CreationDate"`
	LastChangeDate                string  `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
