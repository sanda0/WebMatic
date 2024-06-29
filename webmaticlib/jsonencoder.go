package webmaticlib

import (
	"fmt"
	"log"

	"github.com/buger/jsonparser"
)

func CreateActionBlocks(bc *BlockContainer, action []byte, selector string) {
	at, _, _, err := jsonparser.Get(action, "type")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch string(at) {
	case "action_click":

		bc.Blocks = append(bc.Blocks, Block{
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
		bc.Blocks = append(bc.Blocks, Block{
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
		CreateActionBlocks(bc, next_action, selector)
	}

}

func CreateBlocksForCssSelectorElement(bc *BlockContainer, element []byte) {
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
			CreateActionBlocks(bc, action, string(css_selector_path))
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
			CreateBlocksForCssSelectorElement(bc, next_element)
		} else if string(nt) == "control_navigate" {
			CreateNavigationBlock(bc, next_element)
		}

	}
}

func CreateNavigationBlock(bc *BlockContainer, element []byte) {
	url, _, _, err := jsonparser.Get(element, "fields", "url")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bc.Blocks = append(bc.Blocks, Block{
		Type: "navigate",
		Data: map[string]interface{}{
			"url": string(url),
		},
	})
}
