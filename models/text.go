package models

// Text is a struct containing FR/EN text content
type Text struct {
	en string
	fr string
}

// getLang
func (t *Text) getLang(l string) string {

	rs := ""

	switch l {
	case "en":
		rs = t.en
	case "fr":
		rs = t.fr
	}
	return rs
}
