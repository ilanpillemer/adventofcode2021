package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "D2FE28"
var input2 = "38006F45291200"
var input3 = "EE00D40C823060"
var input4 = "8A004A801A8002F478"
var input5 = "8A004A801A8002F478"
var input6 = "620080001611562C8802118E34"
var input7 = "C0015000016115A2E0802F182340"
var input8 = "A0016C880162017C3686B18A3D4780"
var state = "V"
var current int
var program = ""

func main() {
	fmt.Println("day16")
	program = load(part1)
	//program = load(input8)
	fmt.Println(program)

	//parse
	fmt.Println("starting")
	total := int64(0)
	for {

		token, symbol := nextToken()
		//		fmt.Println(symbol, token)
		if symbol == "V" {

			version, err := strconv.ParseInt(token, 2, 64)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(token, version)
			total = total + version
		}

		if state == "EOF" {
			break
		}
	}
	fmt.Println(total)
	fmt.Println("\nfinished")
}
func load(input string) string {
	sb := strings.Builder{}

	for _, chr := range input {
		sb.WriteString(conv(chr))
	}
	return sb.String()
}

func nextToken() (string, string) {
	token := ""
	symbol := state
	switch state {
	case "LINPACKETS":
		token = program[current : current+11]
		current = current + 11
		state = "V"
	case "LINBITS":
		token = program[current : current+15]
		current = current + 15
		state = "V"
	case "I":
		token = program[current : current+1]
		current = current + 1
		switch token {
		case "0":
			state = "LINBITS"
		case "1":
			state = "LINPACKETS"
		default:
			panic(symbol)
		}
	case "A":
		token = program[current : current+5]
		current = current + 5
		if token[0] == '0' {
			state = "V"
		} else {
			state = "A"
		}
		token = token[1:]
	case "V": //version
		rest := program[current:]
		if len(rest) < 16 && atoi(rest) == 0 {
			symbol = "EOF"
			state = "EOF"
			break
		}
		token = program[current : current+3]
		current = current + 3
		state = "T"
	case "T": //version
		token = program[current : current+3]
		current = current + 3
		switch token {
		case "100":
			state = "A" // In Literal keep reading
		default:
			state = "I"
		}

	default:
		panic("unknown state " + state)
	}
	return token, symbol
}

func atoi(str string) int {
	x, _ := strconv.Atoi(str)
	return x
}

func conv(str rune) string {
	c, _ := strconv.ParseInt(string(str), 16, 64)
	//b := strconv.FormatInt(c, 2)
	b := fmt.Sprintf("%04b", c)
	//fmt.Println("..", b)
	return b
}

var part1 = "220D4B80491FE6FBDCDA61F23F1D9B763004A7C128012F9DA88CE27B000B30F4804D49CD515380352100763DC5E8EC000844338B10B667A1E60094B7BE8D600ACE774DF39DD364979F67A9AC0D1802B2A41401354F6BF1DC0627B15EC5CCC01694F5BABFC00964E93C95CF080263F0046741A740A76B704300824926693274BE7CC880267D00464852484A5F74520005D65A1EAD2334A700BA4EA41256E4BBBD8DC0999FC3A97286C20164B4FF14A93FD2947494E683E752E49B2737DF7C4080181973496509A5B9A8D37B7C300434016920D9EAEF16AEC0A4AB7DF5B1C01C933B9AAF19E1818027A00A80021F1FA0E43400043E174638572B984B066401D3E802735A4A9ECE371789685AB3E0E800725333EFFBB4B8D131A9F39ED413A1720058F339EE32052D48EC4E5EC3A6006CC2B4BE6FF3F40017A0E4D522226009CA676A7600980021F1921446700042A23C368B713CC015E007324A38DF30BB30533D001200F3E7AC33A00A4F73149558E7B98A4AACC402660803D1EA1045C1006E2CC668EC200F4568A5104802B7D004A53819327531FE607E118803B260F371D02CAEA3486050004EE3006A1E463858600F46D8531E08010987B1BE251002013445345C600B4F67617400D14F61867B39AA38018F8C05E430163C6004980126005B801CC0417080106005000CB4002D7A801AA0062007BC0019608018A004A002B880057CEF5604016827238DFDCC8048B9AF135802400087C32893120401C8D90463E280513D62991EE5CA543A6B75892CB639D503004F00353100662FC498AA00084C6485B1D25044C0139975D004A5EB5E52AC7233294006867F9EE6BA2115E47D7867458401424E354B36CDAFCAB34CBC2008BF2F2BA5CC646E57D4C62E41279E7F37961ACC015B005A5EFF884CBDFF10F9BFF438C014A007D67AE0529DED3901D9CD50B5C0108B13BAFD6070"
