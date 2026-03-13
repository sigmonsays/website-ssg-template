package site

type FrontMatter struct {
	Title     string   `yaml:"title"`
	Summary   string   `yaml:"summary"`
	Date      string   `yaml:"date"`
	Draft     bool     `yaml:"draft"`
	SkipIndex bool     `yaml:"skip_index"`
	Tags      []string `yaml:"tags"`
	Layout    string
	Directory string
}

func (me *FrontMatter) SetDefaults() {
	if me.Layout == "" {
		me.Layout = "page"
	}
}
