package dpfm_api_output_formatter

import (
	"data-platform-api-business-partner-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		pm := &requests.General{}

		i++
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Industry,
			&pm.LegalEntityRegistration,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.OrganizationBPName1,
			&pm.OrganizationBPName2,
			&pm.OrganizationBPName3,
			&pm.OrganizationBPName4,
			&pm.BPTag1,
			&pm.BPTag2,
			&pm.BPTag3,
			&pm.BPTag4,
			&pm.OrganizationFoundationDate,
			&pm.OrganizationLiquidationDate,
			&pm.BusinessPartnerBirthplaceName,
			&pm.BusinessPartnerDeathDate,
			&pm.GroupBusinessPartnerName1,
			&pm.GroupBusinessPartnerName2,
			&pm.AddressID,
			&pm.BusinessPartnerIDByExtSystem,
			&pm.BusinessPartnerIsBlocked,
			&pm.CertificateAuthorityChain,
			&pm.UsageControlChain,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		general = append(general, General{
			BusinessPartner:               data.BusinessPartner,
			BusinessPartnerFullName:       data.BusinessPartnerFullName,
			BusinessPartnerName:           data.BusinessPartnerName,
			Industry:                      data.Industry,
			LegalEntityRegistration:       data.LegalEntityRegistration,
			Country:                       data.Country,
			Language:                      data.Language,
			Currency:                      data.Currency,
			OrganizationBPName1:           data.OrganizationBPName1,
			OrganizationBPName2:           data.OrganizationBPName2,
			OrganizationBPName3:           data.OrganizationBPName3,
			OrganizationBPName4:           data.OrganizationBPName4,
			BPTag1:                        data.BPTag1,
			BPTag2:                        data.BPTag2,
			BPTag3:                        data.BPTag3,
			BPTag4:                        data.BPTag4,
			OrganizationFoundationDate:    data.OrganizationFoundationDate,
			OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
			BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
			BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
			GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
			GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
			AddressID:                     data.AddressID,
			BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
			BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
			CertificateAuthorityChain:	   data.CertificateAuthorityChain,
			UsageControlChain:		  	   data.UsageControlChain,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &general, nil
}

func ConvertToRole(rows *sql.Rows) (*[]Role, error) {
	defer rows.Close()
	role := make([]Role, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Role{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerRole,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &role, err
		}

		data := pm
		role = append(role, Role{
			BusinessPartner:     data.BusinessPartner,
			BusinessPartnerRole: data.BusinessPartnerRole,
			ValidityEndDate:     data.ValidityEndDate,
			ValidityStartDate:   data.ValidityStartDate,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &role, nil
	}

	return &role, nil
}

func ConvertToFinInst(rows *sql.Rows) (*[]FinInst, error) {
	defer rows.Close()
	finInst := make([]FinInst, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.FinInst{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.FinInstIdentification,
			&pm.FinInstCountry,
			&pm.FinInstCode,
			&pm.FinInstBranchCode,
			&pm.FinInstFullCode,
			&pm.FinInstName,
			&pm.FinInstBranchName,
			&pm.SWIFTCode,
			&pm.InternalFinInstCustomerID,
			&pm.InternalFinInstAccountID,
			&pm.FinInstControlKey,
			&pm.FinInstAccountName,
			&pm.FinInstAccount,
			&pm.HouseBank,
			&pm.HouseBankAccount,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &finInst, err
		}

		data := pm
		finInst = append(finInst, FinInst{
			BusinessPartner:           data.BusinessPartner,
			FinInstIdentification:     data.FinInstIdentification,
			FinInstCountry:            data.FinInstCountry,
			FinInstCode:               data.FinInstCode,
			FinInstBranchCode:         data.FinInstBranchCode,
			FinInstFullCode:           data.FinInstFullCode,
			FinInstName:               data.FinInstName,
			FinInstBranchName:         data.FinInstBranchName,
			SWIFTCode:                 data.SWIFTCode,
			InternalFinInstCustomerID: data.InternalFinInstCustomerID,
			InternalFinInstAccountID:  data.InternalFinInstAccountID,
			FinInstControlKey:         data.FinInstControlKey,
			FinInstAccountName:        data.FinInstAccountName,
			FinInstAccount:            data.FinInstAccount,
			HouseBank:                 data.HouseBank,
			HouseBankAccount:          data.HouseBankAccount,
			CreationDate:              data.CreationDate,
			LastChangeDate:            data.LastChangeDate,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &finInst, nil
	}

	return &finInst, nil
}

func ConvertToAccounting(rows *sql.Rows) (*[]Accounting, error) {
	defer rows.Close()
	accounting := make([]Accounting, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Accounting{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.ChartOfAccounts,
			&pm.FiscalYearVariant,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &accounting, err
		}

		data := pm
		accounting = append(accounting, Accounting{
			BusinessPartner:     data.BusinessPartner,
			ChartOfAccounts:     data.ChartOfAccounts,
			FiscalYearVariant:   data.FiscalYearVariant,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &accounting, nil
	}

	return &accounting, nil
}
