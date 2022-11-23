package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

// Создание двумерного массива с заданными размерами
func NewUniverse() Universe{
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Отображение массива
func (u Universe) Show(){
	for i := range u{
		for j:=0;j<len(u[i]);j++{
			if u[i][j] {
				fmt.Printf( "%c", '*')
			} else {fmt.Printf("%c", ' ')}
		}
		fmt.Println()
	}
}

//Заполнение начального массива(звездочек примерно 25%)
func (u Universe) Seed(){
	for i := range u{
		for j:=0;j<len(u[i]);j++ {
			if rand.Intn(5) + 1 == 1{
				u[i][j] = true
			}else {
				u[i][j] = false
			}
		}
		}
}

// Обработка координат( для избежания ошибок полей вне массива)
func (u Universe) Alive(x, y int) bool{
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// Кол-во соседей
func (u Universe) Neighbors(x, y int) int{
	neighbours := 0
	if u.Alive(x, y+1){ neighbours++}
	if u.Alive(x, y-1){ neighbours++}
	if u.Alive(x+1, y){ neighbours++}
	if u.Alive(x-1, y){ neighbours++}
	if u.Alive(x+1, y+1){ neighbours++}
	if u.Alive(x+1, y-1){ neighbours++}
	if u.Alive(x-1, y+1){ neighbours++}
	if u.Alive(x-1, y-1){ neighbours++}
	return neighbours
}

// Узнать какая клетка будет на следующем тике
func (u Universe) Next(x, y int) bool{
	if (u.Alive(x, y) && u.Neighbors(x, y) == 3) || u.Neighbors(x, y) == 2 || (u.Alive(x, y) == false && u.Neighbors(x, y) == 3){
		return true
	} else{
		return false
	}
}

// Записование следующего положения живых полей в новый массив
func Step(a, b Universe){
	for i:= range b{
		for j:=0;j<len(b[i]);j++{
			a[i][j] = b.Next(i, j)
		}
	}
}

func main() {
	game, a := NewUniverse(), NewUniverse()
	game.Seed()
	for i:=0;i<100;i++{
		fmt.Print("\033[H\033[2J")
		Step(a, game)
		a.Show()
		time.Sleep(time.Second*2)
		a, game = game, a
	}
}