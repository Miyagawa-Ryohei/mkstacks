package entity

import "time"

type UDOFUpCondition struct {
	UpDelay       time.Duration
	Depends       string
	WaitCondition string
}

type UDOFConfiguration struct {
	UDOF        string            `yaml:"udof,omitempty"`
	Timing      []UDOFUpCondition `yaml:"timing,omitempty"`
	DependsUDOF []string          `yaml:"depends,omitempty"`
}

type StackConfiguration struct {
	Parameter   map[string]string            `yaml:"parameters,omitempty"`
	Hosts       map[string]string            `yaml:"hosts,omitempty"`
	DefaultUDOF string                       `yaml:"defaultUDOF,omitempty"`
	UDOFs       map[string]UDOFConfiguration `yaml:"udofs"`
}
