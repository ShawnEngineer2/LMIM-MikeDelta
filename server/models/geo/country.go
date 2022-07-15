package geo

type Country struct {
	Geocountryid   uint32 `gorm:"autoIncrement"`
	Countryid      uint32
	Countryname    string
	Iso3           string
	Iso2           string
	Countrycode    string
	Phonecode      string
	Capitalcity    string
	Currencycode   string
	Internetcode   string
	Georegionid    uint32
	Geosubregionid uint32
}
