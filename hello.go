package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	introduction()
	registerLog("site-falso", false)

	for {

		menu()

		option := readCommand()

		switch option {
		case 1:
			startMonitoring()

		case 2:
			fmt.Println("Exibindo logs...")
			printLogs()

		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

}

func introduction() {
	name := "Eduardo"
	version := 1.1
	fmt.Println("Olá, Sr.", name)
	fmt.Println("Este programa está na versão", version)

}

func menu() {
	fmt.Println("1-Inicializar monitoramento")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair do programa")

}

func readCommand() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("A opção escolhida foi", option)

	return option

}
func startMonitoring() {
	fmt.Println("Monitorando...")
	// sites := []string{}

	sites := readArchiveWebsite()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}
func testSite(site string) {

	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", res.StatusCode)
		registerLog(site, false)
	}

}

func readArchiveWebsite() []string {

	var sites []string

	archives, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(archives)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)
		if err == io.EOF {
			break
		}

	}

	archives.Close()
	return sites
}
func registerLog(site string, status bool) {

	archives, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	archives.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	archives.Close()
}
func printLogs() {

	archives, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archives))
}
