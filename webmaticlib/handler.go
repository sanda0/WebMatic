package webmaticlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func CreateActionBlocks(p *Project, action []byte, selector string) {
	at, _, _, err := jsonparser.Get(action, "type")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch string(at) {
	case "action_click":

		p.Blocks = append(p.Blocks, Block{
			Type: "click",
			Data: map[string]interface{}{
				"target": selector,
			},
		})

	case "action_write":
		txt, _, _, err := jsonparser.Get(action, "fields", "txt")
		if err != nil {
			log.Println(err.Error())
			return
		}
		p.Blocks = append(p.Blocks, Block{
			Type: "write",
			Data: map[string]interface{}{
				"target": selector,
				"text":   string(txt),
			},
		})
		//TODO: add other actions
	}
	//check next
	next_action, _, _, err := jsonparser.Get(action, "next", "block")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	nt, _, _, err := jsonparser.Get(next_action, "type") //next action type
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if string(nt) == "action_click" || string(nt) == "action_write" || string(nt) == "action_wait" {
		CreateActionBlocks(p, next_action, selector)
	}

}

func CreateBlocksForCssSelectorElement(p *Project, element []byte) {
	t, _, _, err := jsonparser.Get(element, "type")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if string(t) == "element_by_css_selector" {
		css_selector_path, _, _, err := jsonparser.Get(element, "fields", "css_selector")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("")
		fmt.Println(string(css_selector_path))
		//get action blocks
		action, _, _, err := jsonparser.Get(element, "inputs", "actions", "block")
		if err != nil {
			fmt.Println(err.Error())

		} else {
			CreateActionBlocks(p, action, string(css_selector_path))
		}
		fmt.Println("")

		next_element, _, _, err := jsonparser.Get(element, "next", "block")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		nt, _, _, err := jsonparser.Get(next_element, "type") //next element type
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if string(nt) == "element_by_css_selector" {
			CreateBlocksForCssSelectorElement(p, next_element)
		}

	}
}

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
			p := Project{
				Headless: false,
			}
			p.Blocks = append(p.Blocks, Block{
				Type: "open",
				Data: map[string]interface{}{"url": string(url)},
			})

			//add elements and acctions
			elements, _, _, err := jsonparser.Get(value, "inputs", "elements", "block")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			CreateBlocksForCssSelectorElement(&p, elements)

			//create project runer
			l := launcher.New().Headless(p.Headless)
			u := l.MustLaunch()
			browser := rod.New().ControlURL(u).MustConnect()
			runer := NewProjectRunner(p, browser)
			runer.Start()
			time.Sleep(time.Minute * 2)
		}

	})
}
