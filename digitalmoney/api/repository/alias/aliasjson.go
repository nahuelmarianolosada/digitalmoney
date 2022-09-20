package alias

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var values []string

func GetRandomAlias() (*string, error) {
	var randomAlias string
	pwd, _ := os.Getwd()
	absPath := pwd + "/api/repository/alias/alias.csv"
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, lines := range csvLines {
		values = append(values, lines[0])
	}

	for i := 0; i < 3; {
		index := rand.Intn(len(values))
		if !strings.Contains(randomAlias, values[index]) {
			if i == 2 {
				randomAlias += values[index]
			} else {
				randomAlias += fmt.Sprintf("%s.", values[index])
			}
			i++
		}
	}
	return &randomAlias, nil
}
