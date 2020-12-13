package main

import(
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
	"bufio"
	"strings"
)

var phrases []string

func main(){
	settings := ReadSettings()
	if(settings == "err"){
		fmt.Print("Couldn't open Settings file, check existance of file in folder files, or reinstall whole programm. If it still doesn't work contact latvianprogame@gmail.com")
		return
	}
	var lang int
	if(settings[0] == '0'){
		lang = SelectLang(0)
	}else{
		lang = int(settings[1])-48
	}
	Clear()
	//writeSettings(lang)	
	OpenLangFile(lang)
	//nextTask := Welcome(lang)
	nextTask := "gen"
	//fmt.Print(nextTask)
	if(nextTask == "gen"){
		Generator(lang)
	}
	Clear()
	End()
}

func ReadSettings()(out string){
	sttIn, err := os.Open("./files/Settings.txt")
	if(err != nil){	
		out = "err"
	}else{
		fmt.Fscan(sttIn, &out)
	}
	sttIn.Close()
	return
}

func SelectLang(strt int)(lang int){
	var text [][]string
	text = append(text, []string{"Select language\n\n", "Available languages:\n","EN\nRU\nLV\n",
		"\nTo select language write one of the options above and press Enter\n\n", "There is no such option provided, please try again\n\n"})
	text = append(text, []string{"Выберите язык\n\n", "Доступные языки:\n", "EN\nRU\nLV\n",
		"\nЧтобы выбрать язык напишите один из предложенных выше вариантов и нажмите Enter\n\n", "Такой опции нету, пожалуйста, попробуйте ещё раз\n\n"})
	text = append(text, []string{"Izvēlēties valodu\n\n", "Pieejamas valodas\n", "EN\nRU\nLV\n",
		"\nLai izvēlētos valodu uzrakstiet vienu no augšā norādītām opcijām  un nospiediet Enter\n\n", "Tādas opcijas nav, lūdzu, pamēģiniet vēlreiz\n\n"})
	
	fmt.Print(text[strt][0])
	for{
		fmt.Print(text[strt][1])
		fmt.Print(text[strt][2])
		fmt.Print(text[strt][3])
		
		var langStr string
		fmt.Scan(&langStr)
		langStr = strings.ToUpper(langStr)
		switch{
			case langStr == "EN":
				lang = 0
				return
				
			case langStr == "RU":
				lang = 1
				return
				
			case langStr == "LV":
				lang = 2
				return
				
			default:
				fmt.Print(text[strt][4])
		}
	}
}

func writeSettings(lang int){
	sttOut, _ := os.Create("./files/Settings.txt")
	fmt.Fprint(sttOut, "1", lang)
	sttOut.Close()
}

func OpenLangFile(langInt int){
	lang := "./languages/"
	switch{
		case langInt == 0:
			lang += "EN.txt"
		case langInt == 1:
			lang += "RU.txt"
		case langInt == 2:
			lang += "LV.txt"
	}
	langIn, _ := os.Open(lang)
	defer langIn.Close()
	scanner := bufio.NewScanner(langIn)
	for scanner.Scan() {
		phrases = append(phrases, scanner.Text())
	}
	fileNewLine()
}

func fileNewLine(){
	for phr := 0; phr < len(phrases); phr++{
		prune := []rune(phrases[phr])
		for i := 0; i < len(prune)-1; i++{
			if(prune[i] == 92 && prune[i+1] == 110){
				var res1, res2 []rune
				for j := 0; j < i; j++{
					res1 = append(res1, prune[j])
				}
				for j := i+2; j < len(prune); j++{
					res2 = append(res2, prune[j])
				}
				var temp string
				temp = string(res1) + "\n" + string(res2)
				prune = []rune(temp)
			}
		}
		phrases[phr] = string(prune)
	}
	return
}

func Welcome(lang int)(i string){
	return
}

func Generator(lang int){
	pof := pofAsk()
	Clear()
	fileName := getFileName(lang)
	//pof := 2
	Contestants := openContFile(fileName)
	//fmt.Println(Contestants)
	var result []string
	for len(Contestants) > 0{
		rand.Seed(time.Now().UnixNano())
		c := rand.Intn(len(Contestants))
		result = append(result, Contestants[c])
		Contestants = remove(Contestants, c)
	}
	/*
	fmt.Println(Contestants)
	fmt.Println(result)
	*/
	switch pof {
		case 2: Res2(result)
		case 3: Res3(result)
		case 4: Res4(result)
	}
}

func getFileName(lang int)string{
	var fileNameIn, txt string
	fmt.Print(phrases[0])
	fmt.Scan(&fileNameIn)
	if(len(fileNameIn) > 3){
		for i := len(fileNameIn)-4; i < len(fileNameIn); i++{
			txt += string(fileNameIn[i])
		}
	}
	if(txt != ".txt"){
		fileNameIn += ".txt"
	}
	return fileNameIn
}

func pofAsk()(pof int){
	fmt.Print(phrases[2])
	fmt.Scan(&pof)
	switch {
		case pof == 2 || pof == 3 || pof == 4:
		default :
			Clear()
			fmt.Print("Try again!\n")
			pofAsk()
	}
	return
}

func openContFile(fileName string)(contestants []string){
	fileName = "./Input/" + fileName
	contIn, _ := os.Open(fileName)
	defer contIn.Close()
	scanner := bufio.NewScanner(contIn)
	for scanner.Scan() {
		contestants = append(contestants, scanner.Text())
	}
	return
}

func Res2(results []string){
	fileName := "./Output/Results.txt"
	resOut, _ := os.Create(fileName)
	defer resOut.Close()
	for i := 0; len(results) > 0; i++{
		fmt.Fprint(resOut, results[0], "\n")
		results = remove(results, 0)
		fmt.Fprint(resOut, results[0], "\n")
		results = remove(results, 0)
		if(len(results) == 1){
			fmt.Fprint(resOut, results[0], "\n")
			results = remove(results, 0)
		}else{
			fmt.Fprint(resOut, "\n")
		}
	}
}

func Res3(results []string){
	fileName := "./Output/Results.txt"
	resOut, _ := os.Create(fileName)
	defer resOut.Close()
	for i := 0; len(results) > 0; i++{
		fmt.Fprint(resOut, results[0], "\n")
		results = remove(results, 0)
		fmt.Fprint(resOut, results[0], "\n")
		results = remove(results, 0)
		if(len(results) == 2){
			fmt.Fprint(resOut, "\n")
		}else if(len(results) == 0){
		}else{
			fmt.Fprint(resOut, results[0], "\n")
			results = remove(results, 0)
			fmt.Fprint(resOut, "\n")
		}
	}
}

func Res4(results []string){
	fileName := "./Output/Results.txt"
	resOut, _ := os.Create(fileName)
	defer resOut.Close()
}

func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func sleep(sleepingTime int){
	time.Sleep(time.Duration(sleepingTime) * time.Millisecond)
}

func Clear(){
	cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func End(){
	fmt.Print(phrases[1])
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func processingBar(){
	mx := 2
	fmt.Print("Processing")
	for i := 0; i < mx; i++{
		for j := 0; j < 3; j++{
			fmt.Print(".")
			time.Sleep(time.Millisecond * 500)
		}
		fmt.Print("\b\b\b   \b\b\b")
		if(i != mx-1){
			time.Sleep(time.Millisecond * 500)
		}
	}
	fmt.Print("...")
	fmt.Print("\nPress Enter to continue . . . ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
//"%USERPROFILE%\Documents\"
