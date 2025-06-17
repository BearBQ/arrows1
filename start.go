package main

import "fmt"

func lesson1(arr1 []string) {

	slice1 := arr1[3:]                        //делаю ноавый срез
	slice1 = append(slice1, "sssdd", "dddss") //добавляю новые элементы
	lenOfSlice1 := len(slice1)                //длина нового массива
	fmt.Println("Новый массив ", slice1, "Его длина ", lenOfSlice1)
}

func main() {
	arrow1 := []string{"asss", "as", "asssss", "sasssss", "aaasss", "as"}
	lesson1(arrow1)

}
