package aby

import (
	"log"
	"os/exec"
)

type ABY struct {
	*exec.Cmd
}

func NewABY() *ABY {
	return &ABY{
		Cmd: exec.Command("go"),
	}
}

func (aby *ABY) Client(param string) ([]byte, error) {
	aby.Args = append(aby.Args, param)
	log.Println(aby.Args)
	return aby.Output()
}
