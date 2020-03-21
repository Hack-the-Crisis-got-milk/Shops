package entities

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Shop struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Loc     Location `json:"loc"`
	Address string   `json:"address"`
	OpenNow bool     `json:"open_now"`
	Photo   string   `json:"photo"`
}
