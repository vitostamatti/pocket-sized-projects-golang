package cache_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	cache "github.com/vitostamatti/genericcache"
)

func TestCache(t *testing.T) {
	c := cache.New[int, string](5, time.Millisecond*100)

	c.Upsert(5, "fünf")

	v, found := c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "fünf", v)

	c.Upsert(5, "pum")

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	v, found = c.Read(5)
	assert.False(t, found)
	assert.Equal(t, "", v)

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)
}

func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string](5, time.Millisecond*100)

	const parallelTasks = 10

	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := range parallelTasks {
		go func(j int) {
			defer wg.Done()
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}
}

func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string](5, time.Millisecond*100)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})
	t.Run("write kuss", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](5, time.Millisecond*100)
	c.Upsert("Norwegian", "Blue")

	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "Blue", got)

	time.Sleep(time.Millisecond * 200)

	got, found = c.Read("Norwegian")
	assert.False(t, found)
	assert.Equal(t, "", got)
}

func TestCache_MaxSize(t *testing.T) {
	t.Parallel()

	c := cache.New[int, int](3, time.Minute)
	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, found := c.Read(1)
	assert.True(t, found)
	assert.Equal(t, 1, got)

	c.Upsert(1, 10)

	c.Upsert(4, 4)

	got, found = c.Read(2)
	assert.False(t, found)
	assert.Equal(t, 0, got)
}
