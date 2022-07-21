package general

//This structure is used as a datagram for passing columns and values to be updated for database records
//requiring two Key column/value pairs

type DataGram02 struct {
	RecKeyColumn1     string
	RecKeyValue1      uint32
	RecKeyColumn2     string
	RecKeyValue2      uint32
	UpdateColumn      string
	UpdateColumnValue string
}
