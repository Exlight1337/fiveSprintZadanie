package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает методы для парсинга и формирования инф-ии об активности
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Info обрабатывает слайс строк с данными и выводит инф-ю об активностях
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
