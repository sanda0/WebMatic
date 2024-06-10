package webmaticlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// SaveStructToJSON saves a struct to a JSON file
func SaveStructToJSON(filePath string, data interface{}) error {
	// Marshal the struct to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Write the JSON data to a file
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

// ReadStructFromJSON reads a JSON file into a struct
func ReadStructFromJSON(filePath string, result interface{}) error {
	// Read the JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return nil
}

// GenerateRandomString generates a random string of the given length
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func SaveMatic(db *gorm.DB, name string, author string) (string, error) {
	matic := Project{
		Name:        name,
		Author:      author,
		CreateAt:    time.Now().String(),
		Description: "",
		Headless:    false,
		Blocks:      []Block{},
	}
	fileName := GenerateRandomString(5) + ".matic.json"
	err := SaveStructToJSON("matics/"+fileName, matic)
	if err != nil {
		return "", err
	}

	projectMap := ProjectMap{
		Name:     name,
		Author:   author,
		FileName: fileName,
	}
	db.Create(&projectMap)
	return fileName, nil
}

func GetAllMatics(db *gorm.DB) ([]ProjectMap, error) {
	matics := []ProjectMap{}

	res := db.Find(&matics)
	if res.Error != nil {
		return nil, res.Error
	}
	return matics, nil

}
