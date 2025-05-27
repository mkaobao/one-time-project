package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type stringKey struct {
	Section string
	Key     string
}

func TestStructKeyMap(iterations int) {
	m := make(map[stringKey]string, iterations)
	for i := 0; i < iterations; i++ {
		key := stringKey{Section: "Common", Key: "Key_" + strconv.Itoa(i)}
		m[key] = "Data " + strconv.Itoa(i)
	}

	defer PrintDurationWithTime("struct key map", iterations, time.Now())

	for i := 0; i < iterations; i++ {
		key := stringKey{Section: "Common", Key: "Key_" + strconv.Itoa(i)}
		_ = m[key]
	}
}

func TestStringKeyMap(iterations int) {
	m := make(map[string]string, iterations)

	for i := 0; i < iterations; i++ {
		key := "Common_Key_" + strconv.Itoa(i)
		m[key] = "Data " + strconv.Itoa(i)
	}

	defer PrintDurationWithTime("string key map", iterations, time.Now())

	for i := 0; i < iterations; i++ {
		key := "Common_Key_" + strconv.Itoa(i)
		_ = m[key]
	}
}

func TestReadOnly(iterations int) {
	m := map[string]string{}
	m["key"] = "data"

	defer PrintDurationWithTime("read only", iterations, time.Now())

	for i := 0; i < iterations; i++ {
		_ = m["key"]
	}
}

func TestReadWithLock(iterations int) {
	var mutex sync.RWMutex
	m := map[string]string{}
	m["key"] = "data"

	defer PrintDurationWithTime("read with lock", iterations, time.Now())

	for i := 0; i < iterations; i++ {
		mutex.RLock()
		_ = m["key"]
		mutex.RUnlock()
	}
}

func TestWriteOnly(iterations int) {
	m := map[string]string{}
	m["key"] = "data"

	defer PrintDurationWithTime("write only", iterations, time.Now())

	for i := 0; i < iterations; i++ {
		m["key"] = "data 2"
	}
}

func TestWriteWithLock(iterations int) {
	var mutex sync.RWMutex
	m := map[string]string{}
	m["key"] = "data"

	defer PrintDurationWithTime("write with lock", iterations, time.Now())

	for i := 0; i < iterations; i++ {
		mutex.Lock()
		m["key"] = "data 2"
		mutex.Unlock()
	}
}

func TestKeyWithStringCombine(iterations int) {
	defer PrintDurationWithTime("key with string combine", iterations, time.Now())

	section := "Common"
	key := "Key_500"
	lang := "EN"

	for i := 0; i < iterations; i++ {
		_ = section + ":" + key + ":" + lang
	}
}

func TestKeyWithFmtFormat(iterations int) {
	defer PrintDurationWithTime("key with fmt format", iterations, time.Now())

	section := "Common"
	key := "Key_500"
	lang := "EN"

	for i := 0; i < iterations; i++ {
		_ = fmt.Sprintf("%s:%s:%s", section, key, lang)
	}
}

func outputResult(label string, duration time.Duration) {
	fmt.Printf("%-30s %10s\n", label, duration.Round(time.Millisecond))
}
