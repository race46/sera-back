package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"sera-back/database"
	"sera-back/gdive"
	"sera-back/models"
)

func UpdatePP(c *fiber.Ctx) error {
	filename := c.Locals("file_name").(string)
	filepath := c.Locals("file_path").(string)
	username := c.Locals("user").(string)

	file := new(models.File)
	user := new(models.User)

	file.ID = uuid.New().String()
	file.Name = filename
	file.Path = filepath

	srv, err := gdive.GetService()

	if err == nil {
		ID, err := gdive.UploadToDrive(srv, filepath, filename)
		if err == nil {
			file.IDOnDrive = ID
			file.Status = 1
			url, err := gdive.PublicUrl(srv, ID)
			if err == nil {
				file.Status = 2
				file.PublicUrl = url
				os.Remove(filepath)
			}
		}
	}

	database.Connection.Create(&file)

	database.Connection.First(&user, "username=?", username)

	oldFile := user.FileId.String

	user.FileId = sql.NullString{
		String: file.ID,
		Valid:  true,
	}

	database.Connection.Save(&user)
	println(oldFile)
	if oldFile != "" {
		of := new(models.File)
		database.Connection.First(&of, "id=?", oldFile)
		database.Connection.Delete(&of)
	}

	return c.JSON(file)

}
