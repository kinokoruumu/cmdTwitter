package Service

import (
	"bufio"
	"os"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
)

func ReadHandle(api *anaconda.TwitterApi) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := s.Text()
		if input == "exit"{
			break
		}
		Tweet(api, input)
		fmt.Print("Tweet: " + input +"\n")
	}
}
