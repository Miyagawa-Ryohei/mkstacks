package entity

// generated by "schematyper -o entity/docker-compose-yml.3.8.go --package entity ./schemata/docker-compose-yml.3.8.schema.json" -- DO NOT EDIT

type Build struct {
	Args       map[string]interface{} `yaml:"args,omitempty"`
	Context    string                 `yaml:"context,omitempty"`
	Dockerfile string                 `yaml:"dockerfile,omitempty"`
	Target     string                 `yaml:"target,omitempty"`
}

type Command string

type ContainerVolumes interface{}

type DependsOn interface{}

type DependsOnItem string

// yml schema for docker-compose
type DockerComposeYml struct {
	Network  map[string]interface{}      `yaml:"network,omitempty"`
	Services map[string]Service          `yaml:"services,omitempty"`
	Version  string                      `yaml:"version,omitempty"`
	Volumes  map[string]ContainerVolumes `yaml:"volumes,omitempty"`
}

type EnvironmentItem string

type ExtraHost string

type Logging struct {
	Driver string                 `yaml:"driver,omitempty"`
	Option map[string]interface{} `yaml:"option,omitempty"`
}

type Networks interface{}

type Port string

type PortsSetting interface{}

type Service struct {
	Build         Build                `yaml:"build,omitempty"`
	ContainerName string               `yaml:"container-name,omitempty"`
	DependsOn     []DependsOnItem      `yaml:"depends_on,omitempty"`
	Environment   []EnvironmentItem    `yaml:"environment,omitempty"`
	ExtraHosts    map[string]ExtraHost `yaml:"extra_hosts,omitempty"`
	Image         string               `yaml:"image,omitempty"`
	Logging       Logging              `yaml:"logging,omitempty"`
	Networks      `yaml:"networks,omitempty"`
	Ports         []Port   `yaml:"ports,omitempty"`
	Tty           bool     `yaml:"tty,omitempty"`
	Volumes       []Volume `yaml:"volumes,omitempty"`
}

type Volume string