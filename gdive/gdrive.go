package gdive

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	// Save file to root directory:
	return c.SaveFile(file, fmt.Sprintf("./files/%d-%s", time.Now().UnixMilli(), file.Filename))
}
