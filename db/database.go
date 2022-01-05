package db

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"time"
	"todoapp"
)

type DataBaseService struct {
	DB *bolt.DB
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func init() {
	var database, _ = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer database.Close()
	_ = database.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Task"))
		return err
	})
	_ = database.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Completed"))
		return err
	})
}

func (dbService DataBaseService) AddTask(task string) (int, error) {
	var id int
	err := dbService.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dbService DataBaseService) GetTask(key int) string {
	var value string
	dbService.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		byteValue := b.Get(itob(key))
		if byteValue != nil {
			value = string(byteValue)
		}
		return nil
	})
	return value
}

func (dbService DataBaseService) ListTasks() ([]todo.Task, error) {
	var taskList []todo.Task
	err := dbService.DB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Task"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			taskList = append(taskList, todo.Task{Key: btoi(k), Value: string(v)})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return taskList, err
}

func (dbService DataBaseService) RemoveTask(key int) error {
	err := dbService.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		return b.Delete(itob(key))
	})
	return err
}

func (dbService DataBaseService) DoTask(key int) (string, error) {
	var value string
	err := dbService.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Completed"))
		if value = dbService.GetTask(key); value != "" {
			return b.Put(itob(key), []byte(value))
		}
		return nil
	})
	err = dbService.RemoveTask(key)
	return value, err
}

func (dbService DataBaseService) CompletedTasks() ([]todo.Task, error) {
	var completedTaskList []todo.Task
	err := dbService.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Completed"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			completedTaskList = append(completedTaskList, todo.Task{Key: btoi(k), Value: string(v)})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return completedTaskList, err
}
