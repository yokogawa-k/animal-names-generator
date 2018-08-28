package animalnamesgenerator // import "github.com/yokogawa-k/animal-names-generator"

import (
	"fmt"
	"math/rand"
)

var (
	left = [...]string{
		"airashii",
		"akippoi",
		"hakanai",
		"kashikoi",
		"kawaii",
		"marui",
		"menkoi",
		"monoganashii",
		"nemui",
		"ohkii",
		"osoroshii",
		"ririshii",
		"takedakeshii",
		"tsuyoi",
		"utsukushii",
		"zurugashikoi",
		"yasashii",
	}

	right = [...]string{
		"hebi",
		"hitsuji",
		"inoshishi",
		"inu",
		"kitsune",
		"neko",
		"nezumi",
		"niwatori",
		"saru",
		"tora",
		"uma",
		"usagi",
		"ushi",
	}
)

// GetRandomName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_animalname". For example 'kawaii_neko'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetRandomName(retry int) string {
	name := fmt.Sprintf("%s_%s", left[rand.Intn(len(left))], right[rand.Intn(len(right))])
	if retry > 0 {
		name = fmt.Sprintf("%s%d", name, rand.Intn(10))
	}
	return name
}
