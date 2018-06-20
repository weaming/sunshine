package main

import "encoding/json"

func JSON(p interface{}) ([]byte, error) {
	var b []byte
	var err error

	b, err = json.MarshalIndent(p, "", "  ")

	if err != nil {
		return []byte{}, nil
	}
	return b, nil
}
