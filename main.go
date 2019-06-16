package main

import(
	"fmt"
	"log"
	"os"

	"github.com/otus-05-2019-golang/stm"
)

func main() {
	source := "a4bc2d5e"
	//source := "abcd"
	//source := "45"
	//source := `qwe\4\5`
	//source := `qwe\45`
	//source := `qwe\\5`
	logger := log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime)
	

	state := stm.State{}
	state.Init(source)

	fmt.Printf("%s: ", source)
	for  state.CurrentState() != stm.TerminatedState {
		result := state.Next()
		if false == result {
			logger.Println("Недопустимый символ в последовательности")
			os.Exit(stm.ErrorCode)
		}
	}

	state.Rewind()
}
