package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const MONITORING_QUANTITY = 2
const MONITORING_DELAY = 5

func main() {
	registerLog("test", true)

	showIntroduction()

	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			initMonitoring()
		case 2:
			printLogs()
			fmt.Println()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Printf("Commando %d desconhecido", command)
			fmt.Println()
			os.Exit(13)
		}
	}
}

func showIntroduction() {
	name := "Fabricio"
	var version = 1.2
	var nameTyped string = "Test"
	fmt.Println("Olá sr. ", name, "-", nameTyped)
	fmt.Println("Program version ", version)
	fmt.Println("Type of version", reflect.TypeOf(version))
	fmt.Println()
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exiber logs")
	fmt.Println("0- Sair")
}

func readCommand() int {
	var command int
	fmt.Scanf("%d", &command)

	fmt.Println("Commando escolhido", command)
	return command
}

func initMonitoring() {
	fmt.Println("Monitorando ...")
	fmt.Println()

	sites := loadSitesFromFile()

	for x := 0; x < MONITORING_QUANTITY; x++ {
		fmt.Println("Monitoring sites step", x)
		for i, site := range sites {
			fmt.Println("Pos", i, "site", site)
			checkSite(site)
		}
		fmt.Println()
		time.Sleep(MONITORING_DELAY * time.Second)
	}
}

func checkSite(site string) {
	res, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode == 200 {
		fmt.Printf("Site %s ok\n", site)
		registerLog(site, true)
	} else {
		//log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
		fmt.Printf("Response failed with status code: %d\n", res.StatusCode)
		registerLog(site, false)
	}
}

func loadSitesFromFile() []string {

	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error to load sites file:", err)
		return sites
	}

	fileReader := bufio.NewReader(file)

	for {
		line, err := fileReader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {

	data, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
