package usecases

import (
	"fmt"
	"notifier/internal/presenter/dtos"
	"os"

	"github.com/aymerick/raymond"
)

// EmailVerification      = "EMAIL_VERIFICATION"
// PasswordReset          = "PASSWORD_RESET"
// WithdrawalNotification = "WITHDRAWAL_NOTIFICATION"
// DepositNotification    = "DEPOSIT_NOTIFICATION"
// BanNotification        = "BAN_NOTIFICATION"
// UnbanNotification      = "UNBAN_NOTIFICATION"

func SendEmailUsecase(serviceType string, data interface{}) (string, error) {
	templates := map[string]string{
		dtos.EmailVerification:      "./templates/verify-email.ejs",
		dtos.PasswordReset:          "./templates/recovery-password.ejs",
		dtos.WithdrawalNotification: "./templates/withdrawal-changes.ejs",
		dtos.BanNotification:        "./templates/suspend-account.ejs",
		dtos.DepositNotification:    "./templates/deposit-notification.ejs",
	}

	templateFile := templates[serviceType]

	htmlTemplate, err := os.ReadFile(templateFile)
	if err != nil {
		return "", fmt.Errorf("error leyendo el archivo de plantilla: %v", err)
	}

	emailHtml, err := raymond.Render(string(htmlTemplate), data)
	if err != nil {
		return "", fmt.Errorf("error renderizando la plantilla: %v", err)
	}

	return emailHtml, nil
}
