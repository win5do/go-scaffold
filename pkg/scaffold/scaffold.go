package scaffold

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr/v2"

	"github.com/win5do/go-scaffold/pkg/logi"
)

var (
	log = logi.Log.Sugar()
)

type scaffold struct {
	debug bool

	data *tplData
}

type tplData struct {
	ProjectAbsPath string // The Abs Gen Project Path
	ProjectName    string // The project name which want to generated
}

func New() *scaffold {
	return new(scaffold)
}

func (s *scaffold) Generate(command, modName string) error {
	currDir, err := filepath.Abs(filepath.Dir(command))
	if err != nil {
		return err
	}

	dirName := filepath.Base(modName)

	s.data = &tplData{
		ProjectAbsPath: filepath.Join(currDir, dirName),
		ProjectName:    modName,
	}

	if err = os.MkdirAll(s.data.ProjectAbsPath, 0755); err != nil {
		return err
	}

	if err := s.genFromTemplate(); err != nil {
		return err
	}
	return nil
}

func (s *scaffold) genFromTemplate() error {
	box := packr.New("tpl", "../template")

	for _, name := range box.List() {
		log.Debugf("name: %s", name)

		tpl, err := box.FindString(name)
		if err != nil {
			return err
		}

		if dir := filepath.Dir(name); dir != "." {
			if err := os.MkdirAll(filepath.Join(s.data.ProjectAbsPath, dir), 0755); err != nil {
				return err
			}
		}

		if strings.HasSuffix(name, ".tpl") {
			name = strings.TrimSuffix(name, ".tpl")
		}

		if err := s.writeFile(filepath.Join(s.data.ProjectAbsPath, name), tpl); err != nil {
			return err
		}
	}

	if err := s.generate(); err != nil {
		return err
	}
	return nil
}

func (s *scaffold) generate() error {
	cmd := exec.Command("go", "generate", "./...")
	cmd.Dir = s.data.ProjectAbsPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *scaffold) writeFile(path, tpl string) error {
	data, err := s.execTpl(tpl)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func (s *scaffold) execTpl(tpl string) ([]byte, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, s.data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
