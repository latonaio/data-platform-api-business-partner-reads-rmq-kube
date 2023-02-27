package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"APIType"`
	General           General  `json:"BusinessPartnerGeneral"`
	Generals          Generals `json:"BusinessPartnerGenerals"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type General struct {
	BusinessPartner               int          `json:"BusinessPartner"`
	BusinessPartnerFullName       *string      `json:"BusinessPartnerFullName"`
	BusinessPartnerName           *string      `json:"BusinessPartnerName"`
	CreationDate                  *string      `json:"CreationDate"`
	CreationTime                  *string      `json:"CreationTime"`
	Industry                      *string      `json:"Industry"`
	LegalEntityRegistration       *string      `json:"LegalEntityRegistration"`
	Country                       *string      `json:"Country"`
	Language                      *string      `json:"Language"`
	Currency                      *string      `json:"Currency"`
	LastChangeDate                *string      `json:"LastChangeDate"`
	LastChangeTime                *string      `json:"LastChangeTime"`
	OrganizationBPName1           *string      `json:"OrganizationBPName1"`
	OrganizationBPName2           *string      `json:"OrganizationBPName2"`
	OrganizationBPName3           *string      `json:"OrganizationBPName3"`
	OrganizationBPName4           *string      `json:"OrganizationBPName4"`
	BPTag1                        *string      `json:"BPTag1"`
	BPTag2                        *string      `json:"BPTag2"`
	BPTag3                        *string      `json:"BPTag3"`
	BPTag4                        *string      `json:"BPTag4"`
	OrganizationFoundationDate    *string      `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string      `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string      `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string      `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      *bool        `json:"BusinessPartnerIsBlocked"`
	GroupBusinessPartnerName1     *string      `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string      `json:"GroupBusinessPartnerName2"`
	AddressID                     *int         `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string      `json:"BusinessPartnerIDByExtSystem"`
	IsMarkedForDeletion           *bool        `json:"IsMarkedForDeletion"`
	Role                          []Role       `json:"Role"`
	FinInst                       []FinInst    `json:"FinInst"`
	Accounting                    []Accounting `json:"Accounting"`
}
type Generals struct {
	BusinessPartners []int `json:"BusinessPartners"`
}

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityStartDate   string `json:"ValidityStartDate"`
}

type FinInst struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	ValidityEndDate           string  `json:"ValidityEndDate"`
	ValidityStartDate         string  `json:"ValidityStartDate"`
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
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type Accounting struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
