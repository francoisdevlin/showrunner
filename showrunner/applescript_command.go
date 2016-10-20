package showrunner

import (
	//"flag"
	"fmt"
	//"io/ioutil"
	//"os"
	//"path/filepath"
)

type Applescript struct {
}

func (this *Applescript) Main(args []string) error {
	//fmt.Println("Hello")
	fmt.Println(lineBuilder([]string{"echo \"Hi\""}, 3))
	return nil
}
