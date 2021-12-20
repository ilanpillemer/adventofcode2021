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
var minimum = "880086C3E88112"
var maximum = "CE00C43D881120"
var lessthan = "D8005AC2A8F0"
var greaterthan = "F600BC2D8F"
var equals = "9C005AC2F8F0"
var subequals = "9C0141080250320F1802104A08"

// sum
var sum = "C200B40A82"
var product = "04005AC33890"

//
var input9 = "9C0141080250320F1802104A08"
var state = "V"
var current, counter int64
var program = ""

type node struct {
	parent   *node
	typ      string
	children []*node
	val      int64
	op       string
}

func (n node) String() string {
	sb := strings.Builder{}
	switch n.typ {
	case "LIT":
		sb.WriteString(fmt.Sprintf("[%s %v]", n.typ, n.val))
	default:
		sb.WriteString(fmt.Sprintf("[%s %s %v]", n.typ, n.op, n.children))
	}
	return sb.String()
}

func main() {
	fmt.Println("day16")
	program = load(part1)
	//parse
	fmt.Println("starting")
	packets := getPacketStream()
	packets = firstPass(packets)
	packets = secondPass(packets)
	packets = thirdPass(packets)
	_ = packets
	for _, packet := range packets {
		fmt.Println(packet)
		_ = packet
	}
	root := &node{}
	ast, _ := generateAST(root, packets[1:])
	//log.Println(ast)
	fmt.Println("****")
	fmt.Println("answer", Eval(ast))
	fmt.Println("****")

}

func Eval(current *node) int64 {
	switch current.typ {
	case "LIT":
		return current.val
	default:
		switch current.op {
		case "+":
			total := int64(0)
			for _, child := range current.children {
				total = total + Eval(child)
			}
			return total

		case "*":
			total := int64(1)
			for _, child := range current.children {
				total = total * Eval(child)
			}
			return total
		case "max":
			total := int64(0)
			for _, child := range current.children {
				total = max(total, Eval(child))
			}
			return total
		case "min":
			total := int64(9223372036854775807)
			for _, child := range current.children {
				total = min(total, Eval(child))
			}
			return total
		case "=":
			if Eval(current.children[0]) == Eval(current.children[1]) {
				return 1
			}
			return 0
		case ">":
			if Eval(current.children[0]) > Eval(current.children[1]) {
				return 1
			}
			return 0
		case "<":
			if Eval(current.children[0]) < Eval(current.children[1]) {
				return 1
			}
			return 0

		}
	}
	panic("illegal program")
}

func generateAST(parent *node, packets []Packet) (*node, []Packet) {
	packet := packets[0]
	current := &node{}
	current.parent = parent
	switch packet.typ {
	case "LIT":
		current.typ = "LIT"
		current.val, _ = strconv.ParseInt(packet.val, 2, 64)
	default:
		current.typ = "OP"
		current.op = packet.val
		children := []*node{}
		for i := 0; i < int(packet.args); i++ {
			packets = packets[1:]
			var child *node
			child, packets = generateAST(current, packets)
			children = append(children, child)
		}
		current.children = children
	}
	return current, packets
}

type Packet struct {
	counter  int64
	version  int64
	typ      string
	bits     string
	val      string
	number   int64
	args     int64
	npackets int64
	offset   int64
}

func (p Packet) distantSubPackets(packets []Packet) int64 {
	switch p.typ {
	case "LIT":
		return 0
	case "OP":
		if p.args != 0 {
			return p.args
		}
		total := p.npackets
		for i := 0; i < int(p.npackets); i++ {
			total = total - packets[i+1].distantSubPackets(packets[i+1:])
		}
		return total
	}
	panic(packets)
}

func firstPass(packets []Packet) []Packet {

	for i, packet := range packets {
		if packet.offset != 0 && packet.typ == "OP" {
			count := numPackets(packets, packet.counter, packet.counter+packet.offset)
			packet.npackets = count
			switch packet.val {
			case "<", ">", "=":
				packet.args = 2
			}
			packet.offset = 0
			packets[i] = packet
		}
	}

	return packets

}

func secondPass(packets []Packet) []Packet {

	for i, packet := range packets {
		if packet.typ == "LIT" {
			j, _ := strconv.ParseInt(packet.val, 2, 64)
			packet.number = j
			packets[i] = packet
		}

	}

	return packets

}

func thirdPass(packets []Packet) []Packet {

	for i, packet := range packets {
		if packet.typ == "OP" && packet.args == 0 {
			extras := int64(0)
			rest := packets[i:]
			for i := int64(0); i < packet.npackets; i++ {
				extras = extras + rest[i+1].distantSubPackets(rest[i+1:])
			}
			packet.args = packet.npackets - int64(extras)
			packets[i] = packet
		}
	}
	for i, packet := range packets {
		if packet.typ == "LIT" {
			packet.npackets = 1
			packets[i] = packet
		}
		if packet.typ == "OP" && packet.npackets == 0 {
			extras := int64(0)
			rest := packets[i:]
			for i := 0; i < int(packet.npackets); i++ {
				extras = extras + rest[i+1].distantSubPackets(rest[i+1:])
			}
			packet.npackets = int64(extras) + packet.args
			packets[i] = packet
		}
	}

	return packets

}

func numPackets(packets []Packet, start int64, end int64) int64 {
	count := int64(0)
	for _, packet := range packets {
		if packet.counter > start && packet.counter < end {
			count++
		}
	}
	return count
}

func getPacketStream() []Packet {
	packets := []Packet{}
	packet := Packet{}
	for {

		token, symbol, counter := nextToken()
		switch symbol {
		case "V":
			packets = append(packets, packet)
			packet = Packet{}
			version, _ := strconv.ParseInt(token, 2, 64)
			packet.version = version
			packet.counter = int64(counter)
		case "A":
			packet.val = packet.val + token
		case "LINBITS":
			offset, _ := strconv.ParseInt(token, 2, 64)
			packet.offset = offset + 15 + 7
		case "LINPACKETS":
			args, _ := strconv.ParseInt(token, 2, 64)
			packet.args = args
		case "T":
			switch token {
			case "100":
				packet.typ = "LIT"
			default:
				packet.typ = "OP"
				switch token {
				case "000":
					packet.val = "+"
				case "001":
					packet.val = "*"
				case "010":
					packet.val = "min"
				case "011":
					packet.val = "max"
				case "101":
					packet.val = ">"
				case "110":
					packet.val = "<"
				case "111":
					packet.val = "="

				default:
					panic("abandon")
				}

			}
		}

		if state == "EOF" {
			packets = append(packets, packet)
			return packets
		}
	}

}

func load(input string) string {
	sb := strings.Builder{}

	for _, chr := range input {
		sb.WriteString(conv(chr))
	}
	return sb.String()
}

func nextToken() (string, string, int64) {
	token := ""
	symbol := state
	counter = current
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
	case "T":
		token = program[current : current+3]
		current = current + 3
		state = "I"
		switch token {

		case "100":
			state = "A" // In Literal keep reading
		default:

		}

	default:
		panic("unknown state " + state)
	}
	return token, symbol, counter
}

func atoi(str string) int {
	x, _ := strconv.Atoi(str)
	return x
}

func conv(str rune) string {
	c, _ := strconv.ParseInt(string(str), 16, 64)
	b := fmt.Sprintf("%04b", c)
	return b
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

var part1 = "220D4B80491FE6FBDCDA61F23F1D9B763004A7C128012F9DA88CE27B000B30F4804D49CD515380352100763DC5E8EC000844338B10B667A1E60094B7BE8D600ACE774DF39DD364979F67A9AC0D1802B2A41401354F6BF1DC0627B15EC5CCC01694F5BABFC00964E93C95CF080263F0046741A740A76B704300824926693274BE7CC880267D00464852484A5F74520005D65A1EAD2334A700BA4EA41256E4BBBD8DC0999FC3A97286C20164B4FF14A93FD2947494E683E752E49B2737DF7C4080181973496509A5B9A8D37B7C300434016920D9EAEF16AEC0A4AB7DF5B1C01C933B9AAF19E1818027A00A80021F1FA0E43400043E174638572B984B066401D3E802735A4A9ECE371789685AB3E0E800725333EFFBB4B8D131A9F39ED413A1720058F339EE32052D48EC4E5EC3A6006CC2B4BE6FF3F40017A0E4D522226009CA676A7600980021F1921446700042A23C368B713CC015E007324A38DF30BB30533D001200F3E7AC33A00A4F73149558E7B98A4AACC402660803D1EA1045C1006E2CC668EC200F4568A5104802B7D004A53819327531FE607E118803B260F371D02CAEA3486050004EE3006A1E463858600F46D8531E08010987B1BE251002013445345C600B4F67617400D14F61867B39AA38018F8C05E430163C6004980126005B801CC0417080106005000CB4002D7A801AA0062007BC0019608018A004A002B880057CEF5604016827238DFDCC8048B9AF135802400087C32893120401C8D90463E280513D62991EE5CA543A6B75892CB639D503004F00353100662FC498AA00084C6485B1D25044C0139975D004A5EB5E52AC7233294006867F9EE6BA2115E47D7867458401424E354B36CDAFCAB34CBC2008BF2F2BA5CC646E57D4C62E41279E7F37961ACC015B005A5EFF884CBDFF10F9BFF438C014A007D67AE0529DED3901D9CD50B5C0108B13BAFD6070"
