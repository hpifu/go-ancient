package ancient

type Ancient struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Dynasty string `json:"dynasty,omitempty"`
	Content string `json:"content,omitempty"`
}

type GetAncientReq struct {
	ID int `json:"id"`
}

type GetAncientRes struct {
	OK      bool     `json:"ok"`
	Ancient *Ancient `json:"ancient"`
}
