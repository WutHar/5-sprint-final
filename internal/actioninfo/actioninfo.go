package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга строки '%s': %v\n", data, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации для строки '%s': %v\n", data, err)
			continue
		}
		fmt.Println(info)
	}
}
