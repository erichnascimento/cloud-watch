package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

// Disk represents a disk configuragion
type Disk struct {
	Path      string
	Label     string
	Threshold float64
}

type SMTP struct {
	Email string
}

type Notification struct {
	SMTP SMTP `yaml:"SMTP"`
}

// Config represents the main YAML configuration
type Config struct {
	Disks        []Disk
	Interval     time.Duration
	Notification Notification
}

// New return a configuration from file
func New(file string) (*Config, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return NewString(string(b))
}

// NewString Create a new Config from string
func NewString(str string) (*Config, error) {
	c := new(Config)

	// unmarshal
	err := yaml.Unmarshal([]byte(str), &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
