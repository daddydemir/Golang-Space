package safe

import (
	"log"
	"sync"
)

var SingletonInstance *ThreadSafeSingleton
var once sync.Once

type ThreadSafeSingleton struct{}

func GetInstance() *ThreadSafeSingleton {

	once.Do(func() {
		SingletonInstance = &ThreadSafeSingleton{}
		log.Println("Singleton instance is created.")
	})

	return SingletonInstance
}
