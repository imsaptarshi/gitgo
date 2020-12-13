package utils

import "fmt"

func AsciiBanner() {
	fmt.Println(LBlue("                _  _                 "))
	fmt.Println(LBlue("               (_)| |                "))
	fmt.Println(LBlue("          __ _  _ | |_  __ _   ___   "))
	fmt.Println(LBlue("         / _` || || __|/ _` | / _ \\  "))
	fmt.Println(LBlue("        | (_| || || |_| (_| || (_) | "))
	fmt.Println(LBlue("         \\__, ||_| \\__|\\__, | \\___/  "))
	fmt.Println(LBlue("          __/ |         __/ |        "))
	fmt.Println(LBlue("         |___/         |___/         "))
	fmt.Println()
}
