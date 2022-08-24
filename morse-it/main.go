// Morseit package
// Author: @GeekAmita
// Date: Aug 22, 2022
// Experiment: morse to ascii or ascii to morse

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

// required maps
var asciiToMorseMap map[string]string
var morseToAsciiMap map[string]string

func init() {
	// init is called when main is run
	asciiToMorseMap = map[string]string{
		"A":  ".-",
		"B":  "-...",
		"C":  "-.-.",
		"D":  "-..",
		"E":  ".",
		"F":  "..-.",
		"G":  "--.",
		"H":  "....",
		"I":  "..",
		"J":  ".---",
		"K":  "-.-",
		"L":  ".-..",
		"M":  "--",
		"N":  "-.",
		"O":  "---",
		"P":  ".--.",
		"Q":  "--.-",
		"R":  ".-.",
		"S":  "...",
		"T":  "-",
		"U":  "..-",
		"V":  "...-",
		"W":  ".--",
		"X":  "-..-",
		"Y":  "-.--",
		"Z":  "--..",
		" ":  "/",
		"0":  "-----",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		".":  ".-.-.-",
		",":  "--..--",
		":":  "---...",
		"?":  "..--..",
		"'":  ".----.",
		"-":  "-....-",
		"/":  "-..-.",
		"(":  "-.--.-",
		"@":  ".--.-.",
		"=":  "-...-",
		"\"": ".-..-.",
	}
	morseToAsciiMap = reversedMap(asciiToMorseMap)
}

// Function to create a reversed map, provided an original map
func reversedMap(original_map map[string]string) map[string]string {
	reversed_map := make(map[string]string)
	for key, value := range original_map {
		reversed_map[value] = key
	}
	return reversed_map
}

// Mehod for morse to ascii conversion
func convertMorseIntoAscii(input string) string {
	splitted := strings.Split(input, " ")
	var result string = ""
	for i := 0; i < len(splitted); i++ {
		if val, ok := morseToAsciiMap[splitted[i]]; ok {
			result = result + string(val)
		} else {
			var errorString string = "  [ERROR] Couldn't match \"" + string(splitted[i]) + "\" to ascii"
			return errorString
		}
	}
	return result
}

// Method for ascii to morse conversion
func convertAsciiIntoMorse(input string) string {
	var result string = ""
	for i := 0; i < len(input); i++ {
		if val, ok := asciiToMorseMap[strings.ToUpper(string(input[i]))]; ok {
			result = result + string(val) + " "
		} else {
			var errorString string = "  [ERROR] Couldn't match \"" + string(input[i]) + "\" to morse"
			return errorString
		}
	}
	return result
}

func main() {
	//[Usage] morseit ["ascii or morse input string"]
	//
	//    Example 1: morseit "hello world"
	//    > .... . .-.. .-.. --- / .-- --- .-. .-.. -..
	//
	//    Example 2: morseit ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."
	//    > HELLO WORLD
	//
	//    read more about morse code here
	//    https://en.wikipedia.org/wiki/Morse_code
	//
	if len(os.Args) == 2 {
		var inputString string = os.Args[1]
		var converted string
		if matched, _ := regexp.MatchString("^[\\/.\\-\\s]*$", inputString); matched {
			fmt.Println("Morse string detected ... ")
			fmt.Println("The corresponding ascii string is below")
			converted = convertMorseIntoAscii(inputString)
		} else {
			fmt.Println("Ascii string detected ... ")
			fmt.Println("The corresponding morse string is below")
			converted = convertAsciiIntoMorse(inputString)
		}
		fmt.Println(converted)
		clipboard.WriteAll(converted)
		fmt.Println("converted string copied to your clipboard")
	} else {
		fmt.Println("Usage: morseit [ascii or morse string here]")
	}
}
