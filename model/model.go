package model

type SrcRoot struct {
	Name  string   `json:"name"`
	Num   int      `json:"num"`
	Sites []string `json:"sites"`
}

type StepMethod struct {
	Method []string `json:"method"`
}

type Rule struct {
	Root SrcRoot    `json:"root"`
	Step StepMethod `json:"method"`
}
