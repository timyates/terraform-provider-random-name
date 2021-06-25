package name

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

func read(fs embed.FS, filename string) []string {
	bytes, e := fs.ReadFile(filename)
	if e != nil {
		fmt.Print(e)
		panic("Nooooooo")
	}
	return strings.Split(string(bytes), "\n")
}

type NameGenerator interface {
	One([]string) string
}

func NewNameGenerator() NameGenerator {
	return &nameGenerator{}
}

type nameGenerator struct {
}

func (n *nameGenerator) One(a []string) string {
	return a[rand.Intn(len(a))]
}

type wordGenerator struct {
	words []string
	NameGenerator
}

func (n *wordGenerator) Choose() string {
	return n.NameGenerator.One(n.words)
}

func NewWordGenerator(fs embed.FS, filename string) *wordGenerator {
	return &wordGenerator{
		words:         read(fs, filename),
		NameGenerator: NewNameGenerator(),
	}
}
