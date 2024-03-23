package response

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func ModifyJSONResponse() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Next handler in chain
		if err := ctx.Next(); err != nil {
			return err
		}

		// Check if the response is JSON
		contentType := string(ctx.Response().Header.ContentType())
		if contentType == "application/json" {
			// Get the response body
			body := ctx.Response().Body()

			// Unmarshal JSON into an interface{} value
			var data interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				return err
			}

			// Wrap the data in a new JSON object
			modifiedData := map[string]interface{}{
				"data":    data,
				"success": true,
			}

			// Marshal the modified JSON data back to bytes
			modifiedBody, err := json.Marshal(modifiedData)
			if err != nil {
				return err
			}

			// Set the modified JSON body
			ctx.Response().SetBody(modifiedBody)
		}

		return nil
	}
}
