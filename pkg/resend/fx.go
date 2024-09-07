package resend

import "go.uber.org/fx"

var Module = fx.Module("resend", fx.Provide(NewResendImplementation))
