package entities

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FilterType string

const (
	BusynessFilter  FilterType = "busyness"
	AvailableFilter            = "available"
)

var ErrUnknownFilterType = errors.New("unknown filter type, available filter types: busyness, available")

type Filter struct {
	Type  FilterType
	Value string
}

func (f *Filter) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"type\":\"%s\",\"value\":\"%s\"}", f.Type, f.Value)), nil
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	var tempStore struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	err := json.Unmarshal(data, &tempStore)
	if err != nil {
		return err
	}

	if tempStore.Type != string(BusynessFilter) && tempStore.Type != string(AvailableFilter) {
		return ErrUnknownFilterType
	}

	f.Type = FilterType(tempStore.Type)
	f.Value = tempStore.Value
	return nil
}
