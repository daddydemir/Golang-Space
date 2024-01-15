package singleton

import "log"

var SingletonInstance *Singleton

type Singleton struct{}

func GetInstance(key string) *Singleton {

	if SingletonInstance == nil {
		SingletonInstance = &Singleton{}
		log.Println("Singleton instance is created.", "KEY :", key)
	}
	return SingletonInstance
}
