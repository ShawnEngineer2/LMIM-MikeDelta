package customer

type Customertype struct {
	Recid      uint32 `gorm:"autoincrement"`
	Custtypeid uint32
	Custtype   string
}
