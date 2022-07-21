package geo

type City struct {
	Geocityid     uint32 `gorm:"autoincrement"`
	Cityid        uint32
	Stateid       uint32
	Countryid     uint32
	Cityname      string
	Citylatitude  float32
	Citylongitude float32
}
