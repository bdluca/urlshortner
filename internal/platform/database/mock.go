package database

type LocalDB struct {
	storage map[string]string
}

func NewMockDB() *LocalDB {
	return &LocalDB{
		storage: map[string]string{},
	}
}

func (db LocalDB) SaveItem(id string, item string) {
	db.storage[id] = item
}

func (db LocalDB) GetItem(id string) string {
	item, ok := db.storage[id]
	if !ok {
		return ""
	}

	return item
}
