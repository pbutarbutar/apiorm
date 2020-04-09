package query

import (
	"pintekid/config"
	"pintekid/libs/security"
	"pintekid/schema/types"

	"github.com/graphql-go/graphql"
)

/*func ProductResolve(param graphql.ResolveParams) (interface{}, error) {

	_, err := security.ValidateJWT(param.Context.Value("token").(string))
	if err != nil {
		//return nil, err
	}

	var a types.Product
	var b []types.Product
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,ifnull(PRO_IMG,'') from Products")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.IdPro, &a.ProName, &a.QteStock, &a.ProImg)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}

func ProductSingleResolve(param graphql.ResolveParams) (interface{}, error) {
	_, err := security.ValidateJWT(param.Context.Value("token").(string))
	if err != nil {
		return nil, err
	}
	var a types.Product
	var b []types.Product
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	id_pro, _ := param.Args["ID_PRO"].(string)

	result, err := db.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,ifnull(PRO_IMG,'') from Products where ID_PRO = ? limit  5", id_pro)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.IdPro, &a.ProName, &a.QteStock, &a.ProImg)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}*/

func EduLoanResolve(param graphql.ResolveParams, status_data int) (interface{}, error) {
	_, err := security.ValidateJWT(param.Context.Value("token").(string))
	if err != nil {
		return nil, err
	}

	var a types.EduLoan
	var b []types.EduLoan
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	//status_id := []int{1, 2, 3}

	q := `
				select loan_app.app_num, loan_app.id as loan_app_id, loan_app.created_at, loan_app.updated_at, 
				student.fullname as student_name, student.phone as student_phone, loan_app.principal, 
				institution.name as institution_name, loan_app_details.app_type as application_type, 
				setting_application_status.backoffice_desc as status, applicant.fullname as applicant_name, 
				coalesce(loan_app_analysis.q_46,0) as risk_level, coalesce(loan_app_analysis.q_45,0) as credit_score
				from loan_app left join loan_app_details on loan_app.id = loan_app_details.loan_app_id 
				inner join loan_app_status on loan_app.id = loan_app_status.loan_app_id 
				inner join applicant on loan_app_details.applicant_id = applicant.id 
				left join student on loan_app_details.student_id = student.id 
				left join loan_app_analysis on loan_app.id = loan_app_analysis.loan_app_id 
				left join institution on loan_app_details.institution_id = institution.id 
				left join setting_institution_category on institution.institution_cat_id = setting_institution_category.id 
				inner join setting_application_status on setting_application_status.id = loan_app_status.status_id 
				where loan_app_status.status_id = ?  AND loan_app.app_num NOT LIKE "%SME%" 
				`

	result, err := db.Query(q, status_data)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.AppNum, &a.LoanAppId, &a.CreatedAt, &a.UpdatedAt,
			&a.StudentName, &a.StudentPhone, &a.Principal, &a.InstitutionName,
			&a.ApplicationType, &a.Status, &a.ApplicantName, &a.RiskLevel, &a.CreditScore)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}

func SmeLoanResolve(param graphql.ResolveParams, status_data int) (interface{}, error) {
	_, err := security.ValidateJWT(param.Context.Value("token").(string))
	if err != nil {
		return nil, err
	}

	var a types.SmeLoan
	var b []types.SmeLoan
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	q := `
				select loan_app.app_num, loan_app.id as loan_app_id, loan_app.created_at, loan_app.updated_at, 
				student.fullname as student_name, student.phone as student_phone, loan_app.principal, 
				institution.name as institution_name, loan_app_details.app_type as application_type, 
				setting_application_status.backoffice_desc as status, applicant.fullname as applicant_name, 
				coalesce(loan_app_analysis.q_46,0) as risk_level, coalesce(loan_app_analysis.q_45,0) as credit_score
				from loan_app left join loan_app_details on loan_app.id = loan_app_details.loan_app_id 
				inner join loan_app_status on loan_app.id = loan_app_status.loan_app_id 
				inner join applicant on loan_app_details.applicant_id = applicant.id 
				left join student on loan_app_details.student_id = student.id 
				left join loan_app_analysis on loan_app.id = loan_app_analysis.loan_app_id 
				left join institution on loan_app_details.institution_id = institution.id 
				left join setting_institution_category on institution.institution_cat_id = setting_institution_category.id 
				inner join setting_application_status on setting_application_status.id = loan_app_status.status_id 
				where loan_app_status.status_id = ? AND loan_app.app_num LIKE '%SME%'
				 order by loan_app.updated_at desc
				`

	result, err := db.Query(q, status_data)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.AppNum, &a.LoanAppId, &a.CreatedAt, &a.UpdatedAt,
			&a.StudentName, &a.StudentPhone, &a.Principal, &a.InstitutionName,
			&a.ApplicationType, &a.Status, &a.ApplicantName, &a.RiskLevel, &a.CreditScore)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}
