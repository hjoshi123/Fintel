package models

import (
	"encoding/json"
	"fmt"
)

type TopContentInfo struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func (tc *TopContentInfo) Scan() ([]byte, error) {
	if tc == nil {
		return nil, fmt.Errorf("top content info is nil")
	}

	return json.Marshal(tc)
}

func (tc *TopContentInfo) Value(s []byte) error {
	if s == nil {
		return fmt.Errorf("top content info is nil")
	}

	if err := json.Unmarshal(s, tc); err != nil {
		return fmt.Errorf("error unmarshalling top content info %v", err.Error())
	}

	return nil
}
