package singleton

import (
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonIsSame(t *testing.T) {

	instance1 := GetInstance("")
	instance2 := GetInstance("")

	assert.Same(t, instance1, instance2)
}

func TestSingletonIsThreadSafe_ShouldFalse(t *testing.T) {

	var wg sync.WaitGroup
	instanceList := make([]*Singleton, 1)

	wg.Add(100)

	for i := 0; i < 100; i++ {

		go func() {
			defer wg.Done()
			instanceList = append(instanceList, GetInstance("concurrent"))
		}()

	}

	wg.Wait()
	instance := GetInstance("main")

	for _, item := range instanceList {

		if item != instance {
			log.Println("This two instance is not same.")
			assert.NotSame(t, item, instance)
		}
	}
}
