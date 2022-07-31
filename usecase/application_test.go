package usecase_test

import (
	"mkstacks/adapter/cli"
	"mkstacks/adapter/gateway"
	"mkstacks/entity"
	"mkstacks/usecase"
	"os"
	"path"
	"strings"
	"testing"
)

func TestApplicationOnSuccess(t *testing.T) {
	app := usecase.NewApplication(cli.DockerSuccessMock{}, cli.DockerComposeSuccessMock{}, gateway.LocalFS{})
	conf := entity.StackConfiguration{
		Parameter: map[string]string{
			"dummy_env1": "dummy_env_val1",
			"dummy_env2": "dummy_env_val2",
			"dummy_env3": "dummy_env_val3",
		},
		Hosts: map[string]string{
			"dummy_host1": "192.168.11.1",
			"dummy_host2": "192.168.11.2",
			"dummy_host3": "192.168.11.3",
		},
		DefaultUDOF: "",
		UDOFs: map[string]entity.UDOFConfiguration{
			"a": {
				UDOF:   "a",
				Timing: nil,
				DependsUDOF: []string{
					"a-1",
				},
			},
			"a-2": {
				UDOF:        "a-2",
				Timing:      nil,
				DependsUDOF: []string{},
			},
			"a-1": {
				UDOF:   "a-1",
				Timing: nil,
				DependsUDOF: []string{
					"a-2",
				},
			},
		},
	}

	t.Run("ResolveFileNames", func(t *testing.T) {
		pwd, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		expect := []string{
			path.Join(pwd, "overrides", "docker-compose.override-a.yml"),
			path.Join(pwd, "overrides", "docker-compose.override-b.yml"),
			path.Join(pwd, "overrides", "docker-compose.override-c.yml"),
		}
		act, err := app.ResolveFileNames([]string{"a", "b", "c"})
		if err != nil {
			t.Fatal(err)
		}
		for i, v := range act {
			if expect[i] != v {
				t.Fatalf("expect is %s, but actual is %s", expect[i], v)
			}
		}
	})

	t.Run("GetDependencies", func(t *testing.T) {
		depends := app.GetDependencies(&conf, []string{"a"})
		m := map[string]bool{}
		for _, depend := range depends {
			m[depend] = true
		}
		if _, ok := m["a-1"]; !ok {
			t.Fatalf("expect [a-1], but actual [%s]", strings.Join(depends, ","))
		}
	})

	t.Run("GetAllDependencies", func(t *testing.T) {
		depends := app.GetAllDependencies(&conf, []string{"a"})
		expect := []string{
			"a-1",
			"a-2",
		}
		m := map[string]bool{}
		for _, depend := range depends {
			m[depend] = true
		}

		for _, k := range expect {
			if _, ok := m[k]; !ok {
				t.Fatalf("expect [\"a-1\", \"a-2\"], but actual [%s]", strings.Join(depends, ","))
			}
		}
	})

	t.Run("CreateInjectionConfig", func(t *testing.T) {
		services := []string{
			"dummy_service1",
			"dummy_service2",
			"dummy_service3",
		}
		compose := app.CreateInjectionConfig(services, conf)
		for _, s := range services {
			if service, ok := compose.Services[s]; ok {
				em := map[entity.EnvironmentItem]bool{}
				for _, v := range service.Environment {
					em[v] = true
				}
				if _, ok := em["dummy_env1=dummy_env_val1"]; !ok {
					t.Fatalf("env expected \n    [dummy_env1=dummy_env_val1,dummy_env2=dummy_env_val2,dummy_env3=dummy_env_val3]\nbut actual\n     %s", service.Environment)
				}
				if _, ok := em["dummy_env2=dummy_env_val2"]; !ok {
					t.Fatalf("env expected \n    [dummy_env1=dummy_env_val1,dummy_env2=dummy_env_val2,dummy_env3=dummy_env_val3]\nbut actual\n     %s", service.Environment)
				}
				if _, ok := em["dummy_env3=dummy_env_val3"]; !ok {
					t.Fatalf("env expected \n    [dummy_env1=dummy_env_val1,dummy_env2=dummy_env_val2,dummy_env3=dummy_env_val3]\nbut actual\n     %s", service.Environment)
				}

				for k, v := range conf.Hosts {
					if ip, ok := service.ExtraHosts[k]; ok {
						if string(ip) != v {
							t.Fatalf("external hosts expected")
						}
					} else {
						t.Fatalf("external hosts expected")
					}
				}
			}
		}
	})
	t.Run("GetServiceList", func(t *testing.T) {

	})
	t.Run("WriteInjectComposeFile", func(t *testing.T) {

	})

	t.Run("GetAllRunComposeFileList", func(t *testing.T) {

	})

	t.Run("Up", func(t *testing.T) {
		if err := app.Up([]string{"opensearch", "mine"}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestApplicationOnDockerError(t *testing.T) {
	//app := usecase.NewApplication(cli.DockerErrorMock{}, cli.DockerComposeSuccessMock{})

	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
}

func TestApplicationOnComposeError(t *testing.T) {
	//app := usecase.NewApplication(cli.DockerSuccessMock{}, cli.DockerComposeErrorMock{})

	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
	t.Run("", func(t *testing.T) {

	})
}
