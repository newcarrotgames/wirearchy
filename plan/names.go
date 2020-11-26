package plan

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/newcarrotgames/wirearchy/gen"
)

// silly naming scheme
var words []string

func Name() string {
	if words == nil {
		file, err := os.Open("/tmp/codewords.txt")
		if err != nil {
			fmt.Printf("error %+v\n\n", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			fword := fmt.Sprintf("%s%s", strings.ToUpper(word[:1]), strings.ToLower(word[1:]))
			words = append(words, fword)
		}
	}

	// get 3 random words
	i, j, k := gen.RND.Intn(len(words)), gen.RND.Intn(len(words)), gen.RND.Intn(len(words))
	return fmt.Sprintf("%s%s%s", words[i], words[j], words[k])
}
