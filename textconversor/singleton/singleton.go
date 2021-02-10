package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	SupportedFormats []string
}

var singleInstance *Single

func GetInstances() *Single {

	//This is to prevent the expensive lock operations every time
	//getinstance() method is called. If this check fails then
	//it means that singleInstance is already created
	if singleInstance == nil {

		//The singleInstance is created inside the lock.
		lock.Lock()
		defer lock.Unlock()

		//This is to make sure that if more than one goroutine bypass
		//the first check then only one goroutine is able to create the
		//singleton instance otherwise each of the goroutine will create its
		//own instance of the single struct.
		if singleInstance == nil {
			fmt.Println("Creando instancias")
			singleInstance = &Single{}
			fillSupportedFormats(singleInstance)
		} else {
			fmt.Println("Las instancias ya existen")
		}
	} else {
		fmt.Println("Las instancias ya existen")
	}

	return singleInstance
}

func fillSupportedFormats(s *Single) {
	s.SupportedFormats = append(s.SupportedFormats, "text")
	s.SupportedFormats = append(s.SupportedFormats, "binary")
	s.SupportedFormats = append(s.SupportedFormats, "morse")
}
