package abstraction

type Entity struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`
}

type Filter struct {
	
}


type EntityImportedWorksheetDetail struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;"`
}