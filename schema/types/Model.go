package types

type Product struct {
	IdPro    string `json:"ID_PRO"`
	ProName  string `json:"PRO_NAME"`
	QteStock int    `json:"QTE_IN_STOCK"`
	ProImg   string `json:"PRO_IMG"`
}

type UserRegistrant struct {
	RegistrantId int64  `json:"registrant_id"`
	Email        string `json:"email" validate:"required,email"`
	Fullname     string `json:"fullname" validate:"required"`
	PhoneNo      string `json:"phone_no" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type EduLoan struct {
	AppNum          string  `json:"app_num"`
	LoanAppId       int     `json:"loan_app_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	StudentName     string  `json:"student_name"`
	StudentPhone    int     `json:"student_phone"`
	Principal       float64 `json:"principal"`
	InstitutionName string  `json:"institution_name"`
	ApplicationType string  `json:"application_type"`
	Status          string  `json:"status"`
	ApplicantName   string  `json:"applicant_name"`
	RiskLevel       string  `json:"risk_level"`
	CreditScore     float64 `json:"credit_score"`
}

type SmeLoan struct {
	AppNum          string  `json:"app_num"`
	LoanAppId       int     `json:"loan_app_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	StudentName     string  `json:"student_name"`
	StudentPhone    int     `json:"student_phone"`
	Principal       float64 `json:"principal"`
	InstitutionName string  `json:"institution_name"`
	ApplicationType string  `json:"application_type"`
	Status          string  `json:"status"`
	ApplicantName   string  `json:"applicant_name"`
	RiskLevel       string  `json:"risk_level"`
	CreditScore     float64 `json:"credit_score"`
}

type LoanEduDetail struct {
	LoanAppId           int     `json:"loan_app_id"`
	StudentName         string  `json:"student_name"`
	StudentPhone        string  `json:"student_phone"`
	InstitutionName     string  `json:"institution_name"`
	DpToPintek          int64   `json:"dp_to_pintek"`
	ApplicationType     string  `json:"application_type"`
	ProgramFee          float64 `json:"program_fee"`
	Mdr                 float64 `json:"mdr"`
	ActualLoanPrincipal float64 `json:"actual_loan_principal"`
	Status              string  `json:"status"`
	AccNumber           string  `json:"acc_number"`
	NameOnBank          string  `json:"name_on_bank"`
	BankName            string  `json:"bank_name"`
	BankCode            string  `json:"bank_code"`
	Cof                 float64 `json:"cof"`
	Tenure              int     `json:"tenure"`
}

type LoanSmeDetail struct {
	LoanAppId           int     `json:"loan_app_id"`
	StudentName         string  `json:"student_name"`
	StudentPhone        string  `json:"student_phone"`
	InstitutionName     string  `json:"institution_name"`
	DpToPintek          int64   `json:"dp_to_pintek"`
	ApplicationType     string  `json:"application_type"`
	ProgramFee          float64 `json:"program_fee"`
	Mdr                 float64 `json:"mdr"`
	ActualLoanPrincipal float64 `json:"actual_loan_principal"`
	Status              string  `json:"status"`
	AccNumber           string  `json:"acc_number"`
	NameOnBank          string  `json:"name_on_bank"`
	BankName            string  `json:"bank_name"`
	BankCode            string  `json:"bank_code"`
	Cof                 float64 `json:"cof"`
	Tenure              int     `json:"tenure"`
}

/*
type BorrowerApplicant struct {
	Id            int     `json:"id"`
	RegistrantId  int     `json:"registrant_id"`
	FullName      string  `json:"fullname"`
	Email         string  `json:"email"`
	Phone         float64 `json:"phone"`
	EmailVerified int     `json:"email_verified"`
	PhoneVerified float64 `json:"phone_verified"`
}
*/
