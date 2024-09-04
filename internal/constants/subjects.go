package constants

import (
	"fmt"
	"notifier/internal/presenter/dtos"
	"os"
)

var CompanyName = os.Getenv("COMPANY_NAME")

var Subjects = map[string]string{
	dtos.EmailVerification:      fmt.Sprintf("Verifica tu correo electrónico - %s", CompanyName),
	dtos.PasswordReset:          fmt.Sprintf("Recupera tu contraseña - %s", CompanyName),
	dtos.WithdrawalNotification: fmt.Sprintf("Actualizaron tu petición de retiro - %s", CompanyName),
	dtos.BanNotification:        fmt.Sprintf("Tu cuenta ha sido suspendida - %s", CompanyName),
	dtos.DepositNotification:    fmt.Sprintf("Depósito exitoso - %s", CompanyName),
}
