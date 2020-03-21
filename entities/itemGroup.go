package entities

type ItemGroup struct {
	ID   string  `json:"id"     bson:"id"    yaml:"id"`
	Name string  `json:"name"   bson:"name"  validate:"required" yaml:"name"`
	Photo string `json:"photo"  bson:"photo" yaml:"photo"`
}
