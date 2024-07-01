package model

type UITagA struct {
	Uitag string `json:"uitag"`
	Text  string `json:"text"`
	Href  string `json:"href"`
}

type UITagImg struct {
	Uitag    string `json:"uitag"`            //img
	Src      string `json:"src"`              //img src
	Position string `json:"position"`         // head foot
	Height   string `json:"height,omitempty"` //img height，大于0生效，否则表示没设置将用默认值
	Href     string `json:"href"`             // img href
}
