package models

import (
	"encoding/json"
	"plant-manager/internal/types"
)

func placeFromRow(row map[string][]byte) (place *types.Place, err error) {
	place = new(types.Place)

	if val, exists := row["id"]; exists {
		if err = json.Unmarshal(val, &place.ID); err != nil {
			return
		}
	}

	if val, exists := row["name"]; exists {
		place.Name = string(val)
	}

	if val, exists := row["rows"]; exists {
		if err = json.Unmarshal(val, &place.Rows); err != nil {
			return
		}
	}

	if val, exists := row["columns"]; exists {
		if err = json.Unmarshal(val, &place.Columns); err != nil {
			return
		}
	}

	if val, exists := row["customer_id"]; exists {
		if err = json.Unmarshal(val, &place.CustomerID); err != nil {
			return
		}
	}

	if val, exists := row["slots"]; exists {
		if err = json.Unmarshal(val, &place.Slots); err != nil {
			return
		}
	}

	return
}
