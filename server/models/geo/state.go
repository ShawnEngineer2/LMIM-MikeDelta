package geo

type State struct {
	Recid          uint32 `gorm:"autoIncrement"`
	Countryid      uint32
	Stateid        uint32
	Statename      string
	Statecode      string
	Statetype      string
	Statelatitude  float32
	Statelongitude float32
}
