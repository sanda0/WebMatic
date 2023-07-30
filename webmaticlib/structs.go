package webmaticlib

// type Goto struct {
// 	Url string `json:"url"`
// }

// type Getter struct {
// 	Target string `json:"target"`
// 	To     string `json:"to"`
// }

// type Setter struct {
// 	Target string `json:"target"`
// 	From   string `json:"from"`
// }

// type Click struct {
// 	Target string `json:"target"`
// }

// type Write struct {
// 	Target string `json:"target"`
// 	Text   string `json:"text"`
// }

// type Types interface {
// }

type Block struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Project struct {
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	CreateAt    string  `json:"createAt"`
	Headless    bool    `json:"headless"`
	Blocks      []Block `json:"blocks"`
}
