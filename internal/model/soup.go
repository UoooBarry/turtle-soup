package model

import (
	"log"

	"gorm.io/datatypes"
)

type Soup struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	SoupQuestion string         `json:"soup_question"`
	SoupAnswer   string         `json:"soup_answer"`
	SoupTag      datatypes.JSON `json:"tag" gorm:"type:text"`
	Enabled      bool           `json:"enabled" gorm:"default:true"`
}

// CreateSoup creates a new soup entry in the database.
func CreateSoup(soup *Soup) error {
	if err := DB.Create(soup).Error; err != nil {
		log.Printf("Failed to create soup: %v", err)
		return err
	}
	return nil
}

// GetSoupByID retrieves a soup entry by its ID.
func GetSoupByID(id string, soup *Soup) error {
	if err := DB.First(soup, id).Error; err != nil {
		log.Printf("Failed to fetch soup with ID %s: %v", id, err)
		return err
	}
	return nil
}

// UpdateSoup updates an existing soup entry in the database.
func UpdateSoup(soup *Soup) error {
	if err := DB.Save(soup).Error; err != nil {
		log.Printf("Failed to update soup with ID %d: %v", soup.ID, err)
		return err
	}
	return nil
}

// DeleteSoup deletes a soup entry by its ID.
func DeleteSoup(id string) error {
	if err := DB.Delete(&Soup{}, id).Error; err != nil {
		log.Printf("Failed to delete soup with ID %s: %v", id, err)
		return err
	}
	return nil
}
