package request

type MyRequest struct {
	Arg1 string `json:"this"`
	Arr  []struct {
		A1 int    `json:"a1,omitempty"`
		B1 string `json:"b1,omitempty"`
	} `json:"arr"`
}