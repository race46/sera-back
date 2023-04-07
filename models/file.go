package models

import (
	"gorm.io/gorm"
	"sera-back/gdive"
)

type File struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	IDOnDrive string `json:"id_on_drive"`
	PublicUrl string `json:"public_url"`
	Status    int    `json:"status"`
}

func (file *File) BeforeDelete(tx *gorm.DB) (err error) {
	if file.IDOnDrive != "" {
		srv, _ := gdive.GetService()
		err := gdive.DeleteFile(srv, file.IDOnDrive)
		if err != nil {
			file.Status = 3
			tx.Save(&file)

			return err
		}
	}
	return
}
