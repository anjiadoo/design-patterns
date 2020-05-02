package component

import (
	"design-patterns/component/demo1/dir"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type NodeHandler func(http.ResponseWriter, *http.Request)

type DirComponent interface {
	Add(DirComponent) bool
	Remove(string) bool
	GetChild() ([]DirComponent, bool)
	GetName() string
	Handler(w http.ResponseWriter, req *http.Request)
	IsDir() bool
}

func AddDir(url string, node DirComponent, root DirComponent) bool {
	if node == nil || root == nil || node.IsDir() {
		return false
	}

	dirs := strings.Split(url, "/")
	dirDeth := len(dirs)
	rootTmp := root

	for i, d := range dirs {
		if i == 0 {
			continue
		}

		if i == dirDeth-1 {
			sons, ok := rootTmp.GetChild()
			if ok {
				for _, dir := range sons {
					if d == dir.GetName() {
						return false
					}
				}
			}
			fmt.Println("bbb")
			rootTmp.Add(node)
			return true
		}
		//获取子节点
		sons, ok := rootTmp.GetChild()
		if !ok {
			newdir := dir.NewDir(d)
			rootTmp.Add(newdir)
			sons, _ = rootTmp.GetChild()
		}

		for _, dir := range sons {
			if d == dir.GetName() {
				rootTmp = dir
				break
			}
		}
	}
	return true
}

func DelDir(url string, root DirComponent) bool {
	if root == nil {
		return false
	}

	dirs := strings.Split(url, "/")
	dirDeth := len(dirs)
	rootTmp := root
	for i, d := range dirs {
		if i == 0 {
			continue
		}

		sons, ok := rootTmp.GetChild()
		if !ok {
			return false
		}

		for _, dir := range sons {
			if d == dir.GetName() {
				if i == dirDeth-1 {
					fmt.Println("last")
					return rootTmp.Remove(d)
				}
				rootTmp = dir
				break
			}
		}

	}

	return false
}

func UpdateDir(url string, node DirComponent, root DirComponent) bool {
	if DelDir(url, root) {
		return AddDir(url, node, root)
	}
	return false
}

func FindNode(url string, root DirComponent) (DirComponent, error) {
	if root == nil {
		return nil, errors.New("root is nil")
	}

	dirs := strings.Split(url, "/")
	dirDeth := len(dirs)
	rootTmp := root
	for i, d := range dirs {
		if i == 0 {
			continue
		}
		sons, ok := rootTmp.GetChild()
		if !ok {
			return nil, errors.New("dir no exist!")
		}

		for _, dir := range sons {
			if d == dir.GetName() {
				if i == dirDeth-1 {

					return dir, nil
				}
				rootTmp = dir
				break
			}
		}
	}
	return nil, errors.New("no find!")
}
