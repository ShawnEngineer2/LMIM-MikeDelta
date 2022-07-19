package general

//This structure is used as a datagram for passing columns and values to be updated for database records
type DataGram struct {
	RecKeyColumn      string
	RecKeyValue       uint32
	UpdateColumn      string
	UpdateColumnValue string
}
