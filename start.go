package main

import "fmt"

func lesson1(arr1 []string) []string {

	slice1 := arr1[3:]                        //делаю ноавый срез
	slice1 = append(slice1, "sssdd", "dddss") //добавляю новые элементы
	lenOfSlice1 := len(slice1)                //длина нового массива
	fmt.Println("Новый срез ", slice1, "Его длина ", lenOfSlice1)
	return slice1
}

func lesson2(slice2 []string) {
	if len(slice2)%2 == 0 { //проверю, есть ли вообще центральный
		fmt.Println("Срез не имеет центрального элемента")
	} else {
		sliceLesson2 := make([]string, 0)      //Создам пустой срез
		centerOfArrow := (len(slice2) - 1) / 2 //центральный элемент среза
		firstPart := slice2[:centerOfArrow]
		secondPart := slice2[centerOfArrow+1:]
		sliceLesson2 = append(sliceLesson2, firstPart...)
		sliceLesson2 = append(sliceLesson2, secondPart...)
		fmt.Println("Срез без центрального элемента", sliceLesson2)
	}
}

func main() {
	arrow1 := []string{"asss", "as", "asssss", "sasssss", "aaasss", "as"}
	resultOfLesson1 := lesson1(arrow1)
	lesson2(resultOfLesson1)
}
