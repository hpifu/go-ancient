package ancient

type GetCategoryQuery struct {
	Title   string `json:"title,omitempty" http:"param"`
	Author  string `json:"author,omitempty" http:"param"`
	Dynasty string `json:"dynasty,omitempty" http:"param"`
	Offset  int    `json:"offset,omitempty" http:"param"`
	Limit   int    `json:"limit,omitempty" http:"param"`
}

type GetCategoryRes struct {
	OK       bool       `json:"ok"`
	Ancients []*Ancient `json:"ancient"`
}
