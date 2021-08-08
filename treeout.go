package treeout

import (
	"bytes"
	"strings"
)

// special character groups used in composing the heriarchy layout
const (
	BranchDelimiterBox = `└─ `

	BranchChainerBox = `│ `

	BranchSplitterBox = `├─ `

	Indent = "   "
)

type Tree struct {
	Docs     []string
	branches []Branches

	isRoot bool
	level  int
	parent Branches
	index  int
	prefix string
}

type Branches interface {
	Child(...string) Branches
	Add(Branches)
	ParentFunc(fn func(Branches))
	String() string

	padding() string
	children() []Branches
	prnt() Branches
	setPrefix(string)
	getPrefix() string
	selfIndex() int

	setBranches([]Branches)
	setLevel(level int)
	setParent(parent Branches)
	setIndex(index int)
}

func New(docs ...string) *Tree {
	return &Tree{
		Docs:   docs,
		isRoot: true,
		level:  0,
	}
}

func (t *Tree) setPrefix(s string) {
	t.prefix = s
}
func (t *Tree) getPrefix() string {
	return t.prefix
}

func (t *Tree) selfIndex() int {
	return t.index
}

func (t Tree) String() string {
	if t.isRoot {
		return t.padding() + fmtDocs(t.Docs) + "\n" + formatArr(t.branches)
	}
	return t.branchLn(t.Docs...) + formatArr(t.branches)
}

func (t *Tree) padding() string {
	var padding string
	for i := 0; i <= t.level; i++ {
		padding += Indent
	}
	return padding
}

func (t *Tree) branchLn(docs ...string) string {
	if t.selfIndex() < len(t.prnt().children())-1 {
		return strings.TrimSuffix(t.getPrefix(), BranchChainerBox) + BranchSplitterBox +
			fmtDocs(docs) + "\n"
	}
	if t.selfIndex() == len(t.prnt().children())-1 {
		return strings.TrimSuffix(t.getPrefix(), BranchChainerBox) + BranchDelimiterBox +
			fmtDocs(docs) + "\n"
	}
	return strings.TrimSuffix(t.getPrefix(), BranchChainerBox) + BranchDelimiterBox +
		fmtDocs(docs) + "\n"
}

func fmtDocs(docs []string) string {
	var s string
	s = strings.Join(docs, " ")
	return s
}
func (t *Tree) children() []Branches {
	return t.branches
}

func (t *Tree) prnt() Branches {
	return t.parent
}

func (t *Tree) Child(docs ...string) Branches {
	newT := &Tree{
		Docs:   docs,
		level:  t.level + 1,
		parent: t,
		index:  len(t.children()),
	}

	t.branches = append(t.branches, newT)
	return newT
}

func (t *Tree) ParentFunc(fn func(Branches)) {
	parent := t.branches[len(t.branches)-1]
	fn(parent)
}

func (t *Tree) Add(children Branches) {
	children.setParent(t)
	children.setIndex(len(t.children()))
	t.branches = append(t.branches, children)
}

func (t *Tree) setBranches(branches []Branches) {
	t.branches = branches
}
func (t *Tree) setLevel(level int) {
	t.level = level
}
func (t *Tree) setParent(parent Branches) {
	t.parent = parent
}
func (t *Tree) setIndex(index int) {
	t.index = index
}

func formatArr(arr []Branches) string {
	var accumulator bytes.Buffer
	for i, v := range arr {
		if len(v.prnt().children()) > i+1 {
			v.setPrefix(v.prnt().getPrefix() + Indent + BranchChainerBox)
		} else {
			if i == len(arr)-1 {
				v.setPrefix(v.prnt().getPrefix() + Indent)
			} else {
				v.setPrefix(v.prnt().getPrefix() + Indent)
			}
		}
		accumulator.WriteString(v.String())
	}
	return accumulator.String()
}
