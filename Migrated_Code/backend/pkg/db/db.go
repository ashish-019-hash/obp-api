package db

import (
	"log"
	"sync"
)

type InMemoryDB struct {
	data map[string]map[string]interface{}
	mu   sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]map[string]interface{}),
	}
}

func (db *InMemoryDB) Set(table, key string, value interface{}) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.data[table] == nil {
		db.data[table] = make(map[string]interface{})
	}
	db.data[table][key] = value
}

func (db *InMemoryDB) Get(table, key string) (interface{}, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if tableData, exists := db.data[table]; exists {
		if value, exists := tableData[key]; exists {
			return value, true
		}
	}
	return nil, false
}

func (db *InMemoryDB) GetAll(table string) map[string]interface{} {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if tableData, exists := db.data[table]; exists {
		result := make(map[string]interface{})
		for k, v := range tableData {
			result[k] = v
		}
		return result
	}
	return make(map[string]interface{})
}

func (db *InMemoryDB) Delete(table, key string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if tableData, exists := db.data[table]; exists {
		delete(tableData, key)
	}
}

func Initialize() *InMemoryDB {
	log.Println("Initializing in-memory database...")
	return NewInMemoryDB()
}
