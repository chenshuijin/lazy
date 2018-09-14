package generator

import "fmt"

// Target describes a target in a makefile
type Target struct {
	Name    string
	Deps    []string
	Cmd     []string
	Comment []string
}

// Param describes a paramater in a makefile
type Param struct {
	Name, Value string
	Comment     []string
}

// Make describes a makefile
type Make struct {
	Params  []Param
	Targets []Target
}

func (m *Make) String() string {
	c := []byte{}
	for _, v := range m.Params {
		c = append(c, []byte(v.String())...)
	}
	for _, v := range m.Targets {
		c = append(c, []byte(v.String())...)
	}
	return string(c)
}

func (p *Param) String() string {
	c := []byte{}
	for _, v := range p.Comment {
		c = append(c, []byte(fmt.Sprintf("# %s\n", v))...)
	}
	c = append(c, []byte(fmt.Sprintf("%s=%s\n", p.Name, p.Value))...)
	return string(c)
}

func (t *Target) String() string {
	c := []byte{}
	for _, v := range t.Comment {
		c = append(c, []byte(fmt.Sprintf("# %s\n", v))...)
	}

	d := ""
	for _, v := range t.Deps {
		d += " " + v
	}
	c = append(c, []byte(fmt.Sprintf("%s:%s\n", t.Name, d))...)

	d = ""
	for _, v := range t.Cmd {
		c = append(c, []byte(fmt.Sprintf("\t%s\n", v))...)
	}

	return string(c)
}
