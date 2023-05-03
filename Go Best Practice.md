1- use make to create new array, which is most efficient way to create compare to other method.
Ex - 
myMap := make(map[string]int)
myMap := map[string]int{
    "apple":  1,
    "banana": 2,
    "orange": 3,
}

2- for loop usage in golang
The for loop in Go is a fundamental control flow construct used to repeatedly execute a block of statements until a specified condition is met.
The basic syntax of a for loop in Go is as follows:
for initialization; condition; post {
    // statement(s)
}


3- Creating array in go
var myArray [5]int
myArray := [5]int{1, 2, 3, 4, 5}
myArray := [...]int{1, 2, 3, 4, 5}

4- In Go, you can declare and initialize a variable using the := shortcut notation. This shorthand declaration syntax can be used only within functions, as it declares and initializes a variable at the same time.
Ex - func example() {
    count := 0 // shorthand declaration and initialization
    // use count here
}

#with var,
func example() {
    var count int // declaration
    count = 0 // initialization
    // use count here
}

