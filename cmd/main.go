package main

import (
	"flag"
	"fmt"
	"log"
	"passwordgenerator"
	"strconv"
)

func main() {
	// pwdLenght := flag.Int("s", 0, "size")
	// if *pwdLenght <= 0 {
	// 	log.Fatal("pwdLenght invalido")
	// }

	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("especifique o tamanho da senha")
	}

	if len(flag.Args()) > 1 {
		log.Fatal("comando invalido")
	}

	pwdLenght, err := strconv.Atoi(flag.Arg(0))
	if err != nil || pwdLenght <= 0 {
		log.Fatal("pwdLenght invalido")
	}

	pwd, err := passwordgenerator.Generate(pwdLenght)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pwd)
}
