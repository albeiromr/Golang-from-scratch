package main

import (
	"fmt"
	"time"
)


func main() {

	//obteniendo la fecha
	date := time.Now()

	// imprimiendo la fecha, muestra la fecha, hora milisegundos y zona horaria
    fmt.Println("la fecha actual es: ", date)



	// Format DD-MM-YYYY
	fmt.Println(date.Format("02-01-2006"))

	// Format MM-DD-YYYY
    fmt.Println(date.Format("01-02-2006"))
  
    // Format MM-DD-YYYY hh:mm:ss
    fmt.Println(date.Format("01-02-2006 15:04:05"))
  
    // With short weekday (Mon)
    fmt.Println(date.Format("01-02-2006 15:04:05 Mon"))
  
    // With weekday (Monday)
    fmt.Println(date.Format("01-02-2006 15:04:05 Monday"))
  
    // Include micro seconds
    fmt.Println(date.Format("01-02-2006 15:04:05.000000"))



	// usando constantes predefinidad. CONTROL + CLICK sobre RFC850 para ver su valor. es un string
	//el paquete "time" tiene muchas mas constantes predefinidas que se pueden usar
	fmt.Println(date.Format(time.RFC850))


	//elementos para construcción de formato, solo estos elementos son reconocidos
	/* 
		Year:               "2006" "06"
		Month:              "Jan" "January" "01" "1"
		Day of the week:    "Mon" "Monday"
		Day of the month:   "2" "_2" "02"
		Day of the year:    "__2" "002"
		Hour:               "15" "3" "03" (PM or AM)
		Minute:             "4" "04"
		Second:             "5" "05"
		AM/PM mark:         "PM" 
	*/

	//numeric timezone
	/* 
		"-0700"     ±hhmm
		"-07:00"    ±hh:mm
		"-07"       ±hh
		"-070000"   ±hhmmss
		"-07:00:00" ±hh:mm:ss 
	*/

	// ejemplo
    fmt.Println(date.Format("Mon _2 Jan 2006"))
	fmt.Println(date.Format("Mon _2 Jan 2006 15:04:05 PM "))
	fmt.Println(date.Format("Mon _2 Jan 2006 15:04:05 PM -0700"))
}
