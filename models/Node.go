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

func NewTreeList(name string, list string) (Node, error) {
	
	lines := strings.Split(list, "\n")
	root := Node{Name: name} // raiz fictícia

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, string(os.PathSeparator))
		current := &root
		for i, part := range parts {
			isFile := i == len(parts)-1
			if isFile {
				current.Files = append(current.Files, Node{Name: part})
				break
			}
			// procura se já existe a pasta
			found := false
			for j := range current.Dirs {
				if current.Dirs[j].Name == part {
					current = &current.Dirs[j]
					found = true
					break
				}
			}
			if !found {
				newDir := Node{Name: part}
				current.Dirs = append(current.Dirs, newDir)
				current = &current.Dirs[len(current.Dirs)-1]
			}
		}
	}

	return root, nil
}

func NewTreePath(path string) (Node, error) {
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
			childNode, e := NewTreePath(childPath)
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