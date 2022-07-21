package geo

type State struct {
	Geostateid     uint32 `gorm:"autoIncrement"`
	Stateid        uint32
	Countryid      uint32
	Statename      string
	Statecode      string
	Statetype      string
	Statelatitude  float32
	Statelongitude float32
}
