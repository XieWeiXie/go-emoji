package emoji_fetch

type Objects struct {
	Request string
	Results Items
}

type Item struct {
	Url        string `json:"url"`
	CodePoints string `json:"code_points"`
	ShortCodes string `json:"short_codes"`
}

type Items []Item
