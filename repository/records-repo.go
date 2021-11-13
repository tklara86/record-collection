package repository


type RecordRepository interface {
	Save(record string)
	FindAll([]string, error)
}

