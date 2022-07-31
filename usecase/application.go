package usecase

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"mkstacks/entity"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

type Application struct {
	docker  entity.Docker
	compose entity.DockerCompose
	filer   entity.Storage
	cwd     string
}

func (app *Application) CheckDependency() error {
	if err := app.docker.IsInstalled(); err != nil {
		return fmt.Errorf("docker is not installed. please install docker-cli")
	}
	if err := app.compose.IsInstalled(); err != nil {
		return fmt.Errorf("docker-compose is not installed. please install docker-cli")
	}
	return nil
}

func (app *Application) ResolveFileNames(udofIDs []string) ([]string, error) {
	l := make([]string, 0)
	for _, udofID := range udofIDs {
		fullFileName := fmt.Sprintf("docker-compose.override-%s.yml", udofID)
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		fullPathName := path.Join(pwd, "overrides", fullFileName)
		l = append(l, fullPathName)
	}
	return l, nil
}

func (app *Application) ReadAllComposeFiles(files []string) ([]entity.DockerComposeYml, error) {

	composes := make([]entity.DockerComposeYml, 0)
	for _, f := range files {
		buf, err := app.filer.Read(f)
		if err != nil {
			return nil, err
		}
		compose := entity.DockerComposeYml{}
		if err := yaml.Unmarshal(buf, &compose); err != nil {
			return nil, err
		}
		composes = append(composes, compose)
	}
	return composes, nil
}

func (app *Application) ReadStackConfiguration() (*entity.StackConfiguration, error) {

	b, err := app.filer.Read(path.Join(app.cwd, ".stack.yml"))
	if err != nil {
		return nil, err
	}
	conf := &entity.StackConfiguration{}

	if err := yaml.Unmarshal(b, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func (app *Application) GetDependencies(conf *entity.StackConfiguration, udofs []string) []string {

	udofMap := map[string]bool{}

	for _, u := range udofs {
		udofMap[u] = true
	}

	udofSetting := conf.UDOFs
	m := map[string]bool{}
	dependsList := make([]string, 0)

	for k, v := range udofSetting {
		// チェックするudofセッティングが引数で指定されていないものならスキップ
		if _, ok := udofMap[k]; !ok {
			continue
		}

		for _, u := range v.DependsUDOF {
			// すでにコマンドライン引数に指定されていた場合は除外
			if _, ok := udofMap[u]; ok {
				continue
			}

			if _, ok := m[u]; !ok {
				dependsList = append([]string{u}, dependsList...)
				m[u] = true
			}
		}

	}
	return dependsList
}

func (app *Application) GetAllDependencies(conf *entity.StackConfiguration, udofs []string) []string {
	dependencies := make([]string, 0)
	l := append(make([]string, 0), udofs...)
	m := map[string]bool{}

	for len(l) != 0 {
		dl := app.GetDependencies(conf, l)
		l = make([]string, 0)
		for _, d := range dl {
			if _, ok := m[d]; !ok {
				m[d] = true
				dependencies = append(dependencies, d)
				l = append(l, d)
			}
		}
	}
	return dependencies
}

func (app *Application) getCommonParameter(param map[string]string) []entity.EnvironmentItem {
	overrideEnvs := make([]entity.EnvironmentItem, 0)
	for k, v := range param {
		overrideEnvs = append(overrideEnvs, (entity.EnvironmentItem)(fmt.Sprintf("%s=%s", k, v)))
	}
	return overrideEnvs
}

func (app *Application) getCommonHosts(param map[string]string) map[string]entity.ExtraHost {
	overrideHosts := map[string]entity.ExtraHost{}
	for k, v := range param {
		overrideHosts[k] = (entity.ExtraHost)(v)
	}
	return overrideHosts
}

func (app *Application) CreateInjectionConfig(services []string, inject entity.StackConfiguration) *entity.DockerComposeYml {
	serviceMap := map[string]entity.Service{}
	for _, s := range services {
		service := entity.Service{}
		overrideEnvs := app.getCommonParameter(inject.Parameter)
		if len(overrideEnvs) > 0 {
			service.Environment = overrideEnvs
		}

		overrideHosts := app.getCommonHosts(inject.Hosts)
		if len(overrideEnvs) > 0 {
			service.ExtraHosts = overrideHosts
		}

		if len(service.ExtraHosts)+len(service.Environment) > 0 {
			serviceMap[s] = service
		}
	}
	if len(serviceMap) == 0 {
		return nil
	}
	return &entity.DockerComposeYml{
		Services: serviceMap,
	}
}

func (app *Application) CreateRecreateConfig(services []string) entity.DockerComposeYml {
	serviceMap := map[string]entity.Service{}
	for _, s := range services {
		service := entity.Service{}
		id := uuid.New().String()
		service.Environment = []entity.EnvironmentItem{
			(entity.EnvironmentItem)(fmt.Sprintf("RECREATE_ID=%s", id)),
		}
		serviceMap[s] = service
	}
	return entity.DockerComposeYml{
		Services: serviceMap,
	}
}

func (app *Application) GetServiceList(composes []entity.DockerComposeYml) []string {
	services := make([]string, 0)
	m := map[string]bool{}
	for _, compose := range composes {
		for k, _ := range compose.Services {
			if ok := m[k]; !ok {
				services = append(services, k)
				m[k] = true
			}
		}
	}
	return services
}

func (app *Application) WriteInjectComposeFile(composes entity.DockerComposeYml) (string, error) {
	f := path.Join(app.cwd, "overrides", ".inject.yml")
	return f, app.filer.WriteAsYaml(composes, f)
}

func (app *Application) WriteRecreateComposeFile(composes entity.DockerComposeYml) (string, error) {
	f := path.Join(app.cwd, "overrides", ".recreate.yml")
	return f, app.filer.WriteAsYaml(composes, f)
}

func (app *Application) Up(udofIDs []string) error {
	conf, err := app.ReadStackConfiguration()
	if err != nil {
		return err
	}

	depends := app.GetDependencies(conf, udofIDs)
	overrides, err := app.ResolveFileNames(append(depends, udofIDs...))
	if err != nil {
		return err
	}
	filePaths := append(
		[]string{path.Join(app.cwd, "docker-compose.yml")},
		overrides...,
	)
	composes, err := app.ReadAllComposeFiles(filePaths)
	if err != nil {
		return err
	}

	services := app.GetServiceList(composes)
	injectCompose := app.CreateInjectionConfig(services, *conf)
	if injectCompose != nil {
		n, err := app.WriteInjectComposeFile(*injectCompose)
		if err != nil {
			return err
		}
		//fileListのマージ
		filePaths = append(filePaths, n)
	}

	buf, err := app.compose.Up(filePaths)

	if err != nil {
		return err
	}

	if err := app.filer.Write(buf, path.Join(app.cwd, "overrides", ".compose.yml")); err != nil {
		log.Println("failed to write .compose.yml")
		log.Println(err)
	}

	if err := app.filer.Write(bytes.NewBufferString(strings.Join(filePaths, "\n")).Bytes(),
		path.Join(app.cwd, "overrides", ".do_not_edit")); err != nil {
		log.Println("failed to write .do_not_edit")
		log.Println(err)
	}

	return nil
}

func (app *Application) Down() error {
	buf, err := os.ReadFile(path.Join(app.cwd, "overrides", ".do_not_edit"))
	if err != nil {
		return err
	}
	filePath := strings.Split(string(buf), "\n")
	return app.compose.Down(filePath)
}

func (app *Application) Recreate(recreate []string) error {
	buf, err := os.ReadFile(path.Join(app.cwd, "overrides", ".do_not_edit"))
	if err != nil {
		return err
	}
	filePath := strings.Split(string(buf), "\n")

	overrides, err := app.ResolveFileNames(recreate)
	if err != nil {
		return err
	}
	composes, err := app.ReadAllComposeFiles(overrides)
	if err != nil {
		return err
	}
	services := app.GetServiceList(composes)
	compose := app.CreateRecreateConfig(services)
	fname, err := app.WriteRecreateComposeFile(compose)
	b, err := app.compose.Up(append(filePath, fname))
	if err != nil {
		return err
	}

	return app.filer.Write(b, path.Join(app.cwd, "overrides", ".compose.yml"))
}

func NewApplication(docker entity.Docker, compose entity.DockerCompose, filer entity.Storage) *Application {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Application{
		docker,
		compose,
		filer,
		cwd,
	}
}
