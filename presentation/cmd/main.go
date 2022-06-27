package main

import (
	"bufio"
	"fmt"
	"github.com/echocat/golang-kata-1/v1/src/application"
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"github.com/echocat/golang-kata-1/v1/src/infraestructure"
	"os"
	"strings"
)

type UsesCases struct {
	GetAllReadbleSortByTitle application.GetAllReadbleSortByTitle
	GetReadbleByAuthors      application.GetReadbleByAuthors
	GetReadbleByIsbn         application.GetReadbleByIsbn
}

type ConsoleApp struct {
	UsesCases UsesCases
}

func (ca *ConsoleApp) InitInMemoryConsoleApp() {
	repoBook, _ := infraestructure.NewMemoryBookRepository()
	repoMagazine, _ := infraestructure.NewMemoryMagazineRepository()

	ca.UsesCases = UsesCases{
		GetAllReadbleSortByTitle: application.GetAllReadbleSortByTitle{repoBook, repoMagazine},
		GetReadbleByAuthors:      application.GetReadbleByAuthors{repoBook, repoMagazine},
		GetReadbleByIsbn:         application.GetReadbleByIsbn{repoBook, repoMagazine},
	}
}

func (ca ConsoleApp) Start() {

	reader := bufio.NewReader(os.Stdin)
	printInstructions()
	fmt.Println("Select a option")

	for {
		mainOptionSelected := readInput(reader)
		switch mainOptionSelected {

		case "1":
			readbles, err := ca.UsesCases.GetAllReadbleSortByTitle.Invoke()
			if err != nil {
				println(err)
			}
			printReadblesInConsole(readbles...)
			fmt.Println()
			printInstructions()

		case "2":
			fmt.Println("Introduce el correo del autor que deseas buscar")
			autor := readInput(reader)
			readbles, err := ca.UsesCases.GetReadbleByAuthors.Invoke(autor)
			if err != nil {
				println(err.Error())
				printInstructions()
				break
			}
			printReadblesInConsole(readbles...)
			fmt.Println()
			printInstructions()

		case "3":
			fmt.Println("Introduce el ISBN del libro/revista que buscas")
			IsBn := readInput(reader)
			readble, err := ca.UsesCases.GetReadbleByIsbn.Invoke(IsBn)
			if err != nil {
				println(err.Error())
				printInstructions()
				break
			}
			printReadblesInConsole(readble)
			fmt.Println()
			printInstructions()

		default:
			printInstructions()
			println("select a valid option")
		}
	}
}

func printInstructions() {
	fmt.Println("Uses Cases")
	fmt.Println("1. for get all titles")
	fmt.Println("2. find by Author")
	fmt.Println("3. find by IsBn")
}

func readInput(reader *bufio.Reader) string {
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func printReadblesInConsole(readbles ...domain.IReadble) {
	for _, readble := range readbles {
		fmt.Println("--------------------------")
		lines := readble.GetPrintData()
		for _, line := range lines {
			fmt.Printf("    %v \n", line)
		}
	}
}

func main() {
	app := ConsoleApp{}
	app.InitInMemoryConsoleApp()
	app.Start()
}
