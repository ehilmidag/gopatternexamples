package main

import "fmt"

var todoCount = 0

type Todo struct {
	todos []string
}

func main() {
	t := Todo{}
	t.AddTodo("hello")
	t.AddTodo("from")
	t.AddTodo("single responsibility")
	fmt.Println(t.todos[0])
	fmt.Println(t.todos[1])
	fmt.Println(t.todos[2])

}

// örnekte görüldüğü gibi to do structunun metodları to do objesini manüpile etmekle sorumlu olmalı
func (t *Todo) AddTodo(text string) int {
	todoCount++
	todoelement := fmt.Sprintf("%d:%s", todoCount, text)
	t.todos = append(t.todos, todoelement)
	return todoCount
}

// bir dosyaya kaydetme işleminin bu structa ait olması single responsiblity'e aykırı çünkü kaydetme onun görevi değil
// ve sistem içinde bir çok dosya kayıt yapmak isteyebilir. O yüzden kaydetme yükleme işlemleri için ayrı bir tip belirlenip
// tüm bu kaydetme yükleme sorumlukları ona yüklenmelidir to do structuna değil
func (t *Todo) SaveToFile(path string) {
	//dosyaya kaydetme işlemi
}
