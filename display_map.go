package deejdsp

// Pretty much ripped from origonal deej repo with changes to suit the displays
import (
	"fmt"
	"sync"
)

type displayMap struct {
	m    map[int][]string
	lock sync.Locker
}

func newDisplayMap() *displayMap {
	return &displayMap{
		m:    make(map[int][]string),
		lock: &sync.Mutex{},
	}
}

func (m *displayMap) iterate(f func(int, []string)) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for key, value := range m.m {
		f(key, value)
	}
}

func (m *displayMap) Get(key int) ([]string, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	value, ok := m.m[key]
	return value, ok
}

func (m *displayMap) set(key int, value []string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.m[key] = value
}

func (m *displayMap) Length() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.m)
}

func (m *displayMap) String() string {
	m.lock.Lock()
	defer m.lock.Unlock()

	displayCount := 0
	targetCount := 0

	for _, value := range m.m {
		displayCount++
		targetCount += len(value)
	}

	return fmt.Sprintf("<%d pictures mapped to %d displays>", displayCount, targetCount)
}
