package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var generals *[]dpfm_api_output_formatter.General
	var role *[]dpfm_api_output_formatter.Role
	var finInst *[]dpfm_api_output_formatter.FinInst
	var accounting *[]dpfm_api_output_formatter.Accounting
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "Generals":
			func() {
				generals = c.Generals(mtx, input, output, errs, log)
			}()
		case "Role":
			func() {
				role = c.Role(mtx, input, output, errs, log)
			}()
		case "FinInst":
			func() {
				finInst = c.FinInst(mtx, input, output, errs, log)
			}()
		case "Accounting":
			func() {
				accounting = c.Accounting(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:    general,
		Generals:   generals,
		Role:       role,
		FinInst:    finInst,
		Accounting: accounting,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	businessPartner := input.General.BusinessPartner

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner = ?;`, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Generals(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	var args []interface{}
	businessPartners := input.Generals.BusinessPartners
	cnt := 0
	for _, v := range businessPartners {
		args = append(args, v)
		cnt++
	}
	repeat := strings.Repeat("?,", cnt-1) + "?"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner IN ( `+repeat+` ) ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGenerals(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Role(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Role {
	var args []interface{}
	businessPartner := input.General.BusinessPartner
	role := input.General.Role

	cnt := 0
	for _, v := range role {
		args = append(args, businessPartner, v.BusinessPartnerRole, v.ValidityEndDate, v.ValidityStartDate)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT BusinessPartner, BusinessPartnerRole, ValidityEndDate, ValidityStartDate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_role_data
		WHERE (BusinessPartner, BusinessPartnerRole, ValidityEndDate, ValidityStartDate) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToRole(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FinInst(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FinInst {
	var args []interface{}
	businessPartner := input.General.BusinessPartner
	finInst := input.General.FinInst

	cnt := 0
	for _, v := range finInst {
		args = append(args, businessPartner, v.FinInstIdentification, v.ValidityEndDate, v.ValidityStartDate)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT BusinessPartner, FinInstIdentification, ValidityEndDate, ValidityStartDate, FinInstCountry, FinInstCode, 
		FinInstBranchCode, FinInstFullCode, FinInstName, FinInstBranchName, SWIFTCode, InternalFinInstCustomerID, 
		InternalFinInstAccountID, FinInstControlKey, FinInstAccountName, FinInstAccount, HouseBank, HouseBankAccount, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_fin_inst_data
		WHERE (BusinessPartner, FinInstIdentification, ValidityEndDate, ValidityStartDate)IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToFinInst(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Accounting(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	businessPartner := input.General.BusinessPartner

	rows, err := c.db.Query(
		`SELECT BusinessPartner, ChartOfAccounts, FiscalYearVariant, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_accounting_data
		WHERE BusinessPartner = ?;`, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAccounting(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
