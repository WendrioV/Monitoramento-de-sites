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

const monitoramentos = 3
const delay = 5
func main (){
	 /*
	 :=
	 tipo =
	 */
	 exibeIntroducao()
	 registralog("site-falso",false)
	 for {
		


		exibeMenu()
		

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	
	case 2:
		fmt.Println("exibindo log...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)

	default:
		fmt.Println("Não conheco este comando")	
		os.Exit(-1)
	}

	}
}
	

func exibeIntroducao(){
	nome := "Wendrio"
	 versao := 1.1
	fmt.Println("Olá, sr.",nome)
	fmt.Println("este programa está na versão",versao)
}

func exibeMenu (){
	fmt.Println("1- iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Scan("o comando escolhido foi ",&comandoLido)

	return comandoLido
}


func iniciarMonitoramento(){
	fmt.Println("monitorando...")
	
	sites :=leSitesDoArquivo()
	for i:=0;i < monitoramentos; i++{
		for i,site:= range sites{
			fmt.Println("testando site",i,":",site)
			testaSite(site)
	}
	time.Sleep(delay * time.Second)
}
	fmt.Println("")
	
	
	
	
}

func testaSite(site string){
	resp,err := http.Get(site)
	if err != nil {
		fmt.Println("ocorreu um erro:",err)
	}
	if resp.StatusCode ==200 {
		fmt.Println("site:",site,"foi carregado com sucesso!")
		registralog(site,true)
		fmt.Println("")
	}else {
		fmt.Println("o site",site,"está com problemas. status code:", resp.StatusCode)
		registralog(site,false)
		fmt.Println("")
	}
}

func leSitesDoArquivo()[]string{
	var sites []string
arquivo,err := os.Open("sites.txt")

if err != nil {
	fmt.Println("ocorreu um erro:",err)
	
}
leitor := bufio.NewReader(arquivo)
	for{
	linha, err := leitor.ReadString('\n')
	linha = strings.TrimSpace(linha)
	sites = append(sites, linha)
	if err == io.EOF {
		break
		
	}
}

return sites
}

func registralog(site string,status bool){

	arquivo,err :=os.OpenFile("log.txt",os.O_RDWR | os.O_CREATE| os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05")+"-"  + site + "-online:" + strconv.FormatBool(status)+ "\n")
	arquivo.Close()
}

func imprimeLogs(){
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}