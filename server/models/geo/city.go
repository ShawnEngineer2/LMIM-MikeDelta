package geo

type City struct {
	Recid         uint32 `gorm:"autoincrement"`
	Countryid     uint32
	Stateid       uint32
	Cityid        uint32
	Cityname      string
	Citylatitude  float32
	Citylongitude float32
}
