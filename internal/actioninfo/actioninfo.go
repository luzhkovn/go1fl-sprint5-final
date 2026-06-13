package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, m := range dataset {
		err := dp.Parse(m)
		if err != nil {
			fmt.Printf("ошибка: %v\n", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка:", err)
		}
		fmt.Println(info)
	}

}
