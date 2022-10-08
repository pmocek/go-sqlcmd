package folder

import (
"sync"
)

var instance *fh
var once sync.Once

func GetInstance() *fh {
	once.Do(func() {
		instance = &fh{}
	})
	return instance
}
