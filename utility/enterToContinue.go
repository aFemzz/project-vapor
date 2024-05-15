package utility

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EnterToContinue() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Press enter to continue ...")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
}
