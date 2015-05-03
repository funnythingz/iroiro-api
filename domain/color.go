package domain

type Color struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

var (
	Red500        = Color{Name: "Red500", Code: "#F44336"}
	Pink500       = Color{Name: "Pink500", Code: "#E91E63"}
	Purple500     = Color{Name: "Purple500", Code: "#9C27B0"}
	DeepPurple500 = Color{Name: "DeepPurple500", Code: "#673AB7"}
	Indigo500     = Color{Name: "Indigo500", Code: "#3F51B5"}
	Blue500       = Color{Name: "Blue500", Code: "#2196F3"}
	LightBlue500  = Color{Name: "LightBlue500", Code: "#03A9F4"}
	Cyan500       = Color{Name: "Cyan500", Code: "#00BCD4"}
	Teal500       = Color{Name: "Teal500", Code: "#009688"}
	Green500      = Color{Name: "Green500", Code: "#4CAF50"}
	LightGreen500 = Color{Name: "LightGreen500", Code: "#8BC34A"}
	Lime500       = Color{Name: "Lime500", Code: "#CDDC39"}
	Yellow500     = Color{Name: "Yellow500", Code: "#FFEB3B"}
	Amber500      = Color{Name: "Amber500", Code: "#FFC107"}
	Orange500     = Color{Name: "Orange500", Code: "#FF9800"}
	DeepOrange500 = Color{Name: "DeepOrange500", Code: "#FF5722"}
	Brawn500      = Color{Name: "Brawn500", Code: "#795548"}
	Grey500       = Color{Name: "Grey500", Code: "#9E9E9E"}
	BlueGrey500   = Color{Name: "BlueGrey500", Code: "#607D8B"}
	Black500      = Color{Name: "Black500", Code: "#000000"}
	White500      = Color{Name: "White500", Code: "#FFFFFF"}
)
