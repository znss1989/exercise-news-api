package main

import "fmt"

func queryChannels() ([]Channel, error) {
	var channels []Channel

	rows, err := db.Query("SELECT id, title FROM channels")
	if err != nil {
		return nil, fmt.Errorf("queryChannels: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chn Channel
		if err := rows.Scan(&chn.ID, &chn.Title); err != nil {
			return nil, fmt.Errorf("queryChannels: %v", err)
		}
		channels = append(channels, chn)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryChannels: %v", err)
	}
	return channels, nil
}

func insertChannel(chn Channel) (int64, error) {
	result, err := db.Exec("INSERT INTO channels (title) VALUES (?)", chn.Title)
	if err != nil {
		return 0, fmt.Errorf("insertChannel: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertChannel: %v", err)
	}
	return id, nil
}
