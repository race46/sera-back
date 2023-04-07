package gdive

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString("error uploading file")
	}

	path := fmt.Sprintf("./files/%d-%s", time.Now().UnixMilli(), file.Filename)
	c.Locals("file_name", file.Filename)
	c.Locals("file_path", path)
	err = c.SaveFile(file, path)
	if err != nil {
		return c.Status(400).SendString("error saving file")
	}
	return c.Next()
}
