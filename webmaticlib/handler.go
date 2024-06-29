package webmaticlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"math/rand"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
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

func SaveMatic(db *gorm.DB, name string, author string) (uint, error) {

	project := Project{
		Name:        name,
		Author:      author,
		Description: "",
		Headless:    false,
		XMLData:     `<xml xmlns="http://www.w3.org/1999/xhtml"></xml>`,
	}
	db.Create(&project)
	return project.ID, nil
}

func GetAllMatics(db *gorm.DB) ([]Project, error) {
	matics := []Project{}

	res := db.Find(&matics)
	if res.Error != nil {
		return nil, res.Error
	}
	return matics, nil

}

func GetMaticById(db *gorm.DB, id uint) (*Project, error) {
	matic := &Project{}
	res := db.First(&matic, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return matic, nil
}

func SaveXML(db *gorm.DB, id uint, data string) error {
	matic := &Project{}
	res := db.First(&matic, id)
	if res.Error != nil {
		return res.Error
	}

	matic.XMLData = data
	db.Save(matic)

	return nil

}

// TODO: get opening project id as parameter
func RunMatic(jsonStr string) {
	// p := Project{}
	// fmt.Println(p)
	blocks, _, _, err := jsonparser.Get([]byte(jsonStr), "blocks", "blocks")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(blocks))
	jsonparser.ArrayEach(blocks, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		//get string blocks
		blockType, _, _, err := jsonparser.Get(value, "type")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if string(blockType) == "control_start" {
			fmt.Println("this is a start")
			url, _, _, err := jsonparser.Get(value, "fields", "base_url")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			bc := BlockContainer{
				Headless: false,
			}
			bc.Blocks = append(bc.Blocks, Block{
				Type: "open",
				Data: map[string]interface{}{"url": string(url)},
			})

			//add elements and acctions
			elements, _, _, err := jsonparser.Get(value, "inputs", "elements", "block")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			CreateBlocksForCssSelectorElement(&bc, elements)

			//create project runer
			l := launcher.New().Headless(bc.Headless)
			u := l.MustLaunch()
			browser := rod.New().ControlURL(u).MustConnect()
			runer := NewProjectRunner(bc, browser)
			runer.Start()
			time.Sleep(time.Minute * 2)
		}

	})
}
