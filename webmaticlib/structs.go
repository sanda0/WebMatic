package webmaticlib

type Goto struct {
	Url string `json:"url"`
}

type Getter struct {
	Target string `json:"target"`
	To     string `json:"to"`
}

type Setter struct {
	Target string `json:"target"`
	From   string `json:"from"`
}

type Click struct {
	Target string `json:"target"`
}

type Write struct {
	Target string `json:"target"`
	Text   string `json:"text"`
}
