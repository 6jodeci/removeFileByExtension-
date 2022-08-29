package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("\t\t\tДобро пожаловать!")
	fmt.Println("   Приложение удаляет файлы с конкретным расширением из выбранной папки")
	fmt.Print("\t\t     Введите путь к папке: > ")
	var dir string
	var ext string
	fmt.Scan(&dir)
	fmt.Print("\t\t   Введите расширение файла: > ")
	fmt.Scan(&ext)

	dir = dir + string(filepath.Separator)

	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Считанная директория: " + dir)

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ext {
				os.Remove(file.Name())
				fmt.Println("Удалено: ", file.Name())
			}
		}
	}
}
