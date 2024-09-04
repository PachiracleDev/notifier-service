package dtos

const (
	EmailVerification      = "EMAIL_VERIFICATION"
	PasswordReset          = "PASSWORD_RESET"
	WithdrawalNotification = "WITHDRAWAL_NOTIFICATION"
	DepositNotification    = "DEPOSIT_NOTIFICATION"
	BanNotification        = "BAN_NOTIFICATION"
)

type SendEmailDTO struct {
	//Email address of the recipient
	To string `json:"to" validate:"required,email"`

	//Services Available: Email verification, Password reset, Withdrawal notification, Deposit notification, Ban notification, Unban notification
	Type string `json:"type" validate:"required,oneof=EMAIL_VERIFICATION PASSWORD_RESET WITHDRAWAL_NOTIFICATION DEPOSIT_NOTIFICATION BAN_NOTIFICATION"`

	//Data to be sent in the email
	Data interface{} `json:"data"`
}
