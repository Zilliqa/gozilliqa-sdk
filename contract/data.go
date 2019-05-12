package contract

type Data struct {
	Tag Transition `json:"_tag"`
	Params []Value `json:"params"`
}
