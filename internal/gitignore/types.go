package gitignore

import (
	"sort"
	"strings"
)

type (
	// Template for gitignore
	Template struct {
		Name        string `json:"name,omitempty"`
		DownloadURL string `json:"download_url,omitempty"`
	}

	// Templates is an array of template
	Templates []Template
)

const (
	templatesURL = "https://api.github.com/repos/toptal/gitignore/contents/templates?ref=master"
)

func (t Template) String() string {
	return strings.ReplaceAll(t.Name, ".gitignore", "")
}

// Sort templates alphabetically
func (tt Templates) Sort() {
	sort.Slice(tt, func(i, j int) bool {
		return tt[i].Name < tt[j].Name
	})
}

// FilterByName all templates
func (tt Templates) FilterByName(names []string) Templates {
	var templates Templates
	for _, name := range names {
		for _, t := range tt {
			if strings.ToLower(t.String()) == strings.ToLower(name) {
				templates = append(templates, t)
				break
			}
		}
	}

	return templates
}
