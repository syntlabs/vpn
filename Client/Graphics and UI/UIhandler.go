package SysOps

import (
	"bufio"
	"fmt"
	"github.com/wailsapp/wails"
	"os"
	"runtime"
	"strings"
)

const uiDimsPath = ""

type UISet struct {
	name string
	v    map[string]map[string]interface{} //nested interface mapping
}

func (subset UISet) init() {

	file, err := os.Open(uiDimsPath)
	if err != nil {
		panic(runtime.PanicNilError{})
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		branch := strings.Split(line, "=")
		nestedMap := make(map[string]interface{})

		chuncks := strings.Split(branch[1], ",")

		for _, piece := range chuncks {

			keyval := strings.Split(piece, ":")

			nestedMap[keyval[0]] = keyval[1]
		}
		//button=width:20200,height:202020,color:blue
		subset.name = branch[0]
		subset.v = map[string]map[string]interface{}{
			subset.name: nestedMap,
		}
	}
}

func main() {}

type myView struct{}

func (m *myView) WailsInit(runtime *wails.Runtime) error {
	runtime.On("myFunction", myFunction)
	return nil
}

func (m *myView) Render() string {
	return `
        <button onclick="myFunction()">Click me!</button>
    `
}

func myFunction() {
	// This function will be called when the button is clicked
	fmt.Println("Button clicked!")
}
