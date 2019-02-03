package proto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type DependsTasks []DependsTask

func (d *DependsTasks) Scan(v interface{}) error {
	switch val := v.(type) {
	case string:
		return json.Unmarshal([]byte(val), d)
	case []byte:
		return json.Unmarshal(val, d)
	default:
		return errors.New("not support")
	}

}

func (d DependsTasks) Value() (driver.Value, error) {
	bts, err := json.Marshal(d)
	return string(bts), err
}

type CrontabArgs struct {
	Weekday string
	Month   string
	Day     string
	Hour    string
	Minute  string
}

func (c *CrontabArgs) Scan(v interface{}) error {
	switch val := v.(type) {
	case string:
		return json.Unmarshal([]byte(val), c)
	case []byte:
		return json.Unmarshal(val, c)
	default:
		return errors.New("not support")
	}

}

func (c CrontabArgs) Value() (driver.Value, error) {
	bts, err := json.Marshal(c)
	return string(bts), err
}

type DependsTask struct {
	Name       string
	Dest       string
	From       string
	ProcessID  int
	ID         string `json:"-"`
	JobEntryID int    `json:"-"`
	Commands   []string
	Timeout    int64
	Err        string `json:"-"`
	LogContent []byte `json:"-"`
}

type PipeComamnds [][]string

func (p *PipeComamnds) Scan(v interface{}) error {
	switch val := v.(type) {
	case string:
		return json.Unmarshal([]byte(val), p)
	case []byte:
		return json.Unmarshal(val, p)
	default:
		return errors.New("not support")
	}

}

func (p PipeComamnds) Value() (driver.Value, error) {
	bts, err := json.Marshal(p)
	return string(bts), err
}