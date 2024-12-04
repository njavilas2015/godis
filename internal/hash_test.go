package internal

import (
	"fmt"
	"sync"
	"testing"
)

var hashStore *HashStore = NewHashStorage(10)

func TestHSetAndHGet(t *testing.T) {

	t.Run("Set and Get value successfully", func(t *testing.T) {
		key := "user:1"
		field := "name"
		value := "John Doe"

		// Set value
		setResponse := hashStore.AddJobHSet(key, field, value)
		if setResponse != "OK" {
			t.Errorf("Expected 'OK', got '%s'", setResponse)
		}

		// Get value
		getResponse := hashStore.AddJobHGet(key, field)
		if getResponse != value {
			t.Errorf("Expected '%s', got '%s'", value, getResponse)
		}
	})

	t.Run("Get non-existent key", func(t *testing.T) {
		key := "user:2"
		field := "name"

		// Attempt to get a non-existent key
		getResponse := hashStore.AddJobHGet(key, field)
		if getResponse != "NOT FOUND" {
			t.Errorf("Expected 'NOT FOUND', got '%s'", getResponse)
		}
	})

	t.Run("Get non-existent field in existing key", func(t *testing.T) {
		key := "user:1"
		field := "age"

		// Attempt to get a non-existent field
		getResponse := hashStore.AddJobHGet(key, field)
		if getResponse != "" {
			t.Errorf("Expected empty string, got '%s'", getResponse)
		}
	})
}
func TestConcurrentAccess(t *testing.T) {

	key := "concurrent:test"
	field := "counter"

	var wg sync.WaitGroup
	var mu sync.Mutex

	const goroutines = 10000

	expectedValue := goroutines

	hashStore.AddJobHSet(key, field, "0")

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for {
				mu.Lock()

				currentValue := hashStore.AddJobHGet(key, field)

				var currentIntValue int

				if currentValue == "" {
					currentIntValue = 0
				} else {

					n, err := fmt.Sscanf(currentValue, "%d", &currentIntValue)

					if err != nil {

						t.Errorf("Error reading current value: %v", err)

						mu.Unlock()

						return
					}

					if n != 1 {

						t.Errorf("Expected to read one integer, got %d", n)

						mu.Unlock()

						return
					}
				}

				newValue := currentIntValue + 1

				updated := hashStore.AddJobHSet(key, field, fmt.Sprintf("%d", newValue))

				if updated == "OK" {
					mu.Unlock()
					break
				}

				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	finalValue := hashStore.AddJobHGet(key, field)

	if finalValue != fmt.Sprintf("%d", expectedValue) {
		t.Errorf("Expected value '%d', got '%s'", expectedValue, finalValue)
	}
}

func TestUnknownCommand(t *testing.T) {
	response := HandlerHashStore("UNKNOWN", []string{"arg1", "arg2"})
	for msg := range response {
		if msg != "ERROR: Unknown command" {
			t.Errorf("Expected 'ERROR: Unknown command', got '%s'", msg)
		}
	}
}
