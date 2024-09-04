package controllers

import (
	"notifier/internal/app/usecases"
	"notifier/internal/constants"
	"notifier/internal/presenter/dtos"
	awssdk "notifier/pkg/aws"
	"notifier/pkg/http"
	"notifier/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func NotifierController(http *http.HttpServer,
	validate *validator.XValidator,
	aws *awssdk.SDKImplementation,
) {
	http.BasicAuthMiddleware()
	api := http.App().Group("/notifier")

	api.Post("/send-email", func(c *fiber.Ctx) error {

		var body dtos.SendEmailDTO

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		if err := validate.Validate(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Message,
			})
		}

		html, err := usecases.SendEmailUsecase(body.Type, body.Data)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error sending email",
			})
		}

		//GET SUBJECT FROM CONFIG
		subject, ok := constants.Subjects[body.Type]

		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error getting subject",
			})
		}

		go aws.SendEmail(awssdk.SendEmailDto{
			Recipient: body.To,
			Subject:   subject,
			HtmlBody:  html,
		})

		return c.JSON(fiber.Map{
			"message": "Email sent successfully",
			"data":    body,
		})
	})

}
