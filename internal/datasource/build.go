package datasource

import (
	"errors"
	"fmt"
	"time"
)

type identifiable interface {
	getID() int
}

func build[T identifiable](items ...T) map[int]T {
	var errs []error
	m := make(map[int]T)
	for _, x := range items {
		id := x.getID()
		if _, found := m[id]; found {
			errs = append(errs, fmt.Errorf("item with id %d already exists", id))
			continue
		}
		m[id] = x
	}
	if len(errs) > 0 {
		panic(errors.Join(errs...))
	}
	return m
}

func DaysInMonth(month time.Month, year int) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
