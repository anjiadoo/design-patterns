package dir

import (
	"design-patterns/component/demo1/component"
	"net/http"
	"sync"
)

type Dir struct {
	Name  string
	Sons  []component.DirComponent
	mutex sync.Mutex
}

func NewDir(name string) *Dir {
	return &Dir{Name: name}
}

func (pD *Dir) Add(iDirComponent component.DirComponent) bool {
	pD.mutex.Lock()
	for _, dirComponent := range pD.Sons {
		if iDirComponent.GetName() == dirComponent.GetName() {
			return false
		}
	}
	pD.Sons = append(pD.Sons, iDirComponent)
	pD.mutex.Unlock()
	return true
}

func (pD *Dir) Remove(name string) bool {
	pD.mutex.Lock()
	for i, dirComponent := range pD.Sons {
		if name == dirComponent.GetName() {
			pD.Sons = append(pD.Sons[:i], pD.Sons[i+1:]...)
			pD.mutex.Unlock()
			return true
		}
	}
	pD.mutex.Unlock()
	return false
}

//dir in web must have child
func (pD *Dir) GetChild() (dirs []component.DirComponent, b bool) {
	if len(pD.Sons) == 0 {
		return dirs, false
	}
	return pD.Sons, true
}

func (pD *Dir) GetName() string {
	return pD.Name
}

func (pD *Dir) IsDir() bool {
	return true
}

//return error
func (pD *Dir) Handler(w http.ResponseWriter, req *http.Request) {
	return
}
