package models

import (
	"strings"
	"os"
	"sync"
	"path/filepath"
)

type Node struct {
	Name  string
	Dirs  []Node
	Files []Node
}

func (this *Node) Stringfy(prefix string, includeName bool, last bool, level []bool) string {
	padding := "\n"
	for i, b := range level {
		if i == len(level)-1 {
			break
		}
		if b {
			padding += "│" + prefix
			continue
		}
		padding += prefix
	}

	out := strings.Builder{}
	if includeName {
		if len(level) == 0 {
			out.WriteString(this.Name)
		} else {
			if last {
				out.WriteString(padding + "╰  " + this.Name)
			} else {
				out.WriteString(padding + "├  " + this.Name)
			}
		}
	}

	all := append(this.Dirs, this.Files...)
	for i, child := range all {
		if i+1 == len(all) {
			out.WriteString(child.Stringfy(prefix, true, true, append(level, false)))
			continue
		}
		out.WriteString(child.Stringfy(prefix, true, false, append(level, true)))
	}
	return out.String()
}

func (this *Node) ToString() string {
	return this.Stringfy("  ", true, false, []bool{})
}

func NewTree(path string) (Node, error) {
	info, err := os.Stat(path)
	if err != nil {
		return Node{}, err
	}

	node := Node{Name: info.Name()}
	if !info.IsDir() {
		return node, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return Node{}, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var goRerr error

	for _, entry := range entries {
		entry := entry
		childPath := filepath.Join(path, entry.Name())

		if !entry.IsDir() {
			node.Files = append(node.Files, Node{Name: entry.Name()})
			continue
		}
		
		wg.Add(1)
		go func() {
			childNode, e := NewTree(childPath)
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			if e != nil {
				if goRerr == nil {
					goRerr = e
				}
				return
			}
			node.Dirs = append(node.Dirs, childNode)
		}()
	}

	wg.Wait()
	return node, goRerr
}