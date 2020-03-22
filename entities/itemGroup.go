package entities

type ItemGroup struct {
	ID       string `json:"id"          bson:"id"         yaml:"id"`
	Name     string `json:"name"        bson:"name"       validate:"required" yaml:"name"`
	Icon     string `json:"icon"        bson:"icon"       yaml:"icon"`
	ShopType string `json:"shop_type"   bson:"shop_type"  yaml:"shop_type"`
}
