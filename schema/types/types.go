package types

import (
	"pintekid/config"
	"pintekid/libs/security"

	"github.com/graphql-go/graphql"
)

var ProductTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Products",
	Fields: graphql.Fields{
		"ID_PRO": &graphql.Field{
			Type: graphql.String,
		},
		"PRO_NAME": &graphql.Field{
			Type: graphql.String,
		},
		"QTE_IN_STOCK": &graphql.Field{
			Type: graphql.Int,
		},
		"PRICE": &graphql.Field{
			Type: graphql.Float,
		},
		"PRO_IMG": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserRegistrantTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserRegistrant",
	Fields: graphql.Fields{
		"registrant_id": &graphql.Field{
			Type: graphql.Int,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"fullname": &graphql.Field{
			Type: graphql.String,
		},
		"phone_no": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var LoanEduAppDetailTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoanEduDetail",
	Fields: graphql.Fields{
		"loan_app_id": &graphql.Field{
			Type: graphql.Int,
		},
		"student_name": &graphql.Field{
			Type: graphql.String,
		},
		"student_phone": &graphql.Field{
			Type: graphql.String,
		},
		"institution_name": &graphql.Field{
			Type: graphql.String,
		},
		"dp_to_pintek": &graphql.Field{
			Type: graphql.Int,
		},
		"application_type": &graphql.Field{
			Type: graphql.String,
		},
		"program_fee": &graphql.Field{
			Type: graphql.Float,
		},
		"mdr": &graphql.Field{
			Type: graphql.Float,
		},
		"actual_loan_principal": &graphql.Field{
			Type: graphql.Float,
		},
		"bank_code": &graphql.Field{
			Type: graphql.String,
		},
		"acc_number": &graphql.Field{
			Type: graphql.String,
		},
		"cof": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var LoanSmeAppDetailTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoanSmeDetail",
	Fields: graphql.Fields{
		"loan_app_id": &graphql.Field{
			Type: graphql.Int,
		},
		"student_name": &graphql.Field{
			Type: graphql.String,
		},
		"student_phone": &graphql.Field{
			Type: graphql.String,
		},
		"institution_name": &graphql.Field{
			Type: graphql.String,
		},
		"dp_to_pintek": &graphql.Field{
			Type: graphql.Int,
		},
		"application_type": &graphql.Field{
			Type: graphql.String,
		},
		"program_fee": &graphql.Field{
			Type: graphql.Float,
		},
		"mdr": &graphql.Field{
			Type: graphql.Float,
		},
		"actual_loan_principal": &graphql.Field{
			Type: graphql.Float,
		},
		"bank_code": &graphql.Field{
			Type: graphql.String,
		},
		"acc_number": &graphql.Field{
			Type: graphql.String,
		},
		"cof": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var EduLoanAppTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "EduLoan",
	Fields: graphql.Fields{
		"loan_app_id": &graphql.Field{
			Type: graphql.Int,
		},
		"app_num": &graphql.Field{
			Type: graphql.String,
		},
		"applicant_name": &graphql.Field{
			Type: graphql.String,
		},
		"student_name": &graphql.Field{
			Type: graphql.String,
		},
		"student_phone": &graphql.Field{
			Type: graphql.String,
		},
		"application_type": &graphql.Field{
			Type: graphql.Int,
		},
		"principal": &graphql.Field{
			Type: graphql.Float,
		},
		"risk_level": &graphql.Field{
			Type: graphql.Int,
		},
		"credit_score": &graphql.Field{
			Type: graphql.Int,
		},
		"loan_detail": &graphql.Field{
			Type: graphql.NewList(LoanEduAppDetailTypes),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				bapplicant, _ := p.Source.(EduLoan)
				//log.Printf("fetching comments of post with id: %d", bapplicant.ApplicantId)

				return LoanEduDetailResolve(bapplicant.LoanAppId, p)
			},
		},
	},
})

var SmeLoanAppTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "SmeLoan",
	Fields: graphql.Fields{
		"loan_app_id": &graphql.Field{
			Type: graphql.Int,
		},
		"app_num": &graphql.Field{
			Type: graphql.String,
		},
		"applicant_name": &graphql.Field{
			Type: graphql.String,
		},
		"student_name": &graphql.Field{
			Type: graphql.String,
		},
		"student_phone": &graphql.Field{
			Type: graphql.String,
		},
		"application_type": &graphql.Field{
			Type: graphql.Int,
		},
		"principal": &graphql.Field{
			Type: graphql.Float,
		},
		"risk_level": &graphql.Field{
			Type: graphql.Int,
		},
		"credit_score": &graphql.Field{
			Type: graphql.Int,
		},
		"loan_detail": &graphql.Field{
			Type: graphql.NewList(LoanSmeAppDetailTypes),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				bapplicant, _ := p.Source.(SmeLoan)
				//log.Printf("fetching comments of post with id: %d", bapplicant.ApplicantId)

				return LoanSmeDetailResolve(bapplicant.LoanAppId, p)
			},
		},
	},
})

func LoanEduDetailResolve(loan_app_id int, param graphql.ResolveParams) (interface{}, error) {
	_, err := security.ValidateJWT(param.Context.Value("token").(string))
	if err != nil {
		//return nil, err
	}

	var a LoanEduDetail
	var b []LoanEduDetail

	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	sql_single := `select loan_app.id as loan_app_id, student.fullname as student_name, 
					student.phone as student_phone, institution.name as institution_name, 
					institution.dp_to_pintek as dp_to_pintek, 
					loan_app_details.app_type as application_type, 
					loan_app_details.program_fee as program_fee, 
					loan_app_details.mdr as mdr, 
					coalesce(loan_app_details.actual_loan_principal,0) as actual_loan_principal, 
					setting_application_status.backoffice_desc as status,
					IFNULL(institution_bank_acc.acc_number,'') as acc_number, 
					IFNULL(institution_bank_acc.name_on_bank,'') as name_on_bank, 
					IFNULL(master_bank.name,'')  as bank_name, 
					IFNULL(master_bank.ojk_code,'')  as bank_code, 
					IFNULL(master_cos_of_funds.value,0) as cof,
					loan_app.tenure
					from loan_app 
					left join loan_app_details on loan_app.id = loan_app_details.loan_app_id 
					inner join loan_app_status on loan_app.id = loan_app_status.loan_app_id 
					left join student on loan_app_details.student_id = student.id 
					left join institution on loan_app_details.institution_id = institution.id 
					left join setting_institution_category on institution.institution_cat_id = setting_institution_category.id 
					inner join setting_application_status on setting_application_status.id = loan_app_status.status_id 
					left join institution_bank_acc on institution_bank_acc.institution_id = institution.id 
					left join master_bank on master_bank.id = institution_bank_acc.bank_id 
					left join loan_lender on loan_lender.loan_app_id = loan_app.id 
					left join master_cos_of_funds on loan_lender.cof_id = master_cos_of_funds.id 
					where loan_app.id = ? limit 1`

	//var xQ = "Select id, registrant_id, fullname, email, phone FROM applicant  where registrant_id = ?"

	result, err := db.Query(sql_single, loan_app_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.LoanAppId, &a.StudentName, &a.StudentPhone, &a.InstitutionName, &a.DpToPintek,
			&a.ApplicationType, &a.ProgramFee, &a.Mdr, &a.ActualLoanPrincipal, &a.Status,
			&a.AccNumber, &a.NameOnBank, &a.BankName, &a.BankCode, &a.Cof, &a.Tenure)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}

func LoanSmeDetailResolve(loan_app_id int, param graphql.ResolveParams) (interface{}, error) {
	_, err := security.ValidateJWT(param.Context.Value("token").(string))
	if err != nil {
		//return nil, err
	}

	var a LoanEduDetail
	var b []LoanEduDetail

	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	sql_single := `select loan_app.id as loan_app_id, student.fullname as student_name, 
					student.phone as student_phone, institution.name as institution_name, 
					institution.dp_to_pintek as dp_to_pintek, 
					loan_app_details.app_type as application_type, 
					loan_app_details.program_fee as program_fee, 
					loan_app_details.mdr as mdr, 
					coalesce(loan_app_details.actual_loan_principal,0) as actual_loan_principal, 
					setting_application_status.backoffice_desc as status,
					IFNULL(institution_bank_acc.acc_number,'') as acc_number, 
					IFNULL(institution_bank_acc.name_on_bank,'') as name_on_bank, 
					IFNULL(master_bank.name,'')  as bank_name, 
					IFNULL(master_bank.ojk_code,'')  as bank_code, 
					IFNULL(master_cos_of_funds.value,0) as cof,
					loan_app.tenure
					from loan_app 
					left join loan_app_details on loan_app.id = loan_app_details.loan_app_id 
					inner join loan_app_status on loan_app.id = loan_app_status.loan_app_id 
					left join student on loan_app_details.student_id = student.id 
					left join institution on loan_app_details.institution_id = institution.id 
					left join setting_institution_category on institution.institution_cat_id = setting_institution_category.id 
					inner join setting_application_status on setting_application_status.id = loan_app_status.status_id 
					left join institution_bank_acc on institution_bank_acc.institution_id = institution.id 
					left join master_bank on master_bank.id = institution_bank_acc.bank_id 
					left join loan_lender on loan_lender.loan_app_id = loan_app.id 
					left join master_cos_of_funds on loan_lender.cof_id = master_cos_of_funds.id 
					where loan_app.id = ? limit 1`

	//var xQ = "Select id, registrant_id, fullname, email, phone FROM applicant  where registrant_id = ?"

	result, err := db.Query(sql_single, loan_app_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.LoanAppId, &a.StudentName, &a.StudentPhone, &a.InstitutionName, &a.DpToPintek,
			&a.ApplicationType, &a.ProgramFee, &a.Mdr, &a.ActualLoanPrincipal, &a.Status,
			&a.AccNumber, &a.NameOnBank, &a.BankName, &a.BankCode, &a.Cof, &a.Tenure)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}
