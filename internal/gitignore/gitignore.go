package gitignore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
	"github.com/pkg/errors"
)

// GetTemplates retuens a list of templates
func GetTemplates() (Templates, error) {
	res, err := http.Get(templatesURL)
	if err != nil {
		return nil, errors.Wrap(err, "error getting templates")
	}
	defer res.Body.Close()

	var tt Templates
	err = json.NewDecoder(res.Body).Decode(&tt)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding response body")
	}

	return tt, nil
}

// DownloadTemplates list
func DownloadTemplates(langs []string) (string, error) {
	var buff bytes.Buffer

	templates, err := GetTemplates()
	if err != nil {
		return "", err
	}

	tt := templates.FilterByName(langs)

	for _, t := range tt {
		log.WithField("name", t.String()).Info("downloading template")
		content, err := DownloadTemplate(t.DownloadURL)
		if err != nil {
			return "", err
		}

		buff.WriteString(fmt.Sprintf("# %s\n", t.String()))
		buff.Write(content)
		buff.WriteString("\n")
	}

	return buff.String(), nil
}

// DownloadTemplate returns the template content
func DownloadTemplate(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error downloading templates")
	}
	defer res.Body.Close()

	template, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading template content")
	}

	return template, nil
}
