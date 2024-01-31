// Single Responsibility Principle

package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// seperation of concerns
// Don't do God object (Anti pattern)

// Func below are breaking single responsibility principle
// Responsibility of the journal is to with management of the entries
// Not to do deal with persistance
// Put "Save", "Load", "LoadFromWeb" into separate components, such as separate
// package or separate type
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
}

func (j *Journal) LoadFromWeb(url *url.URL) {
}

// Separation option here
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// Another option of separation
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// j.Save() is not single responsibility principle

	// One option
	SaveToFile(&j, "journal.txt")

	// Another option
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
