/* Day 16: Packet Decoder */
package main

import (
	"AoC2021/utils/args"
	"AoC2021/utils/bits"
	"AoC2021/utils/color"
	"AoC2021/utils/files"
	"AoC2021/utils/log"
	"AoC2021/utils/numbers"
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	argMap := args.Parse(os.Args)
	log.SetLogLevelFromArgs(argMap)
	selectedPart, partSpecified := argMap["part"]

	log.Println(log.NORMAL, "--- Day 16: Packet Decoder ---")
	timer.Start()

	if selectedPart == "parse" {
		parseInputFile()
	}
	if !partSpecified || selectedPart == "1" {
		part1()
		timer.Tick()
		log.Println(log.NORMAL)
	}

	if !partSpecified || selectedPart == "2" {
		part2()
		timer.Tick()
		log.Println(log.NORMAL)
	}
}

func part1() {
	log.Println(log.NORMAL, "* Part 1 *")
	log.Println(log.NORMAL, " Goal: For now, parse the hierarchy of the packets throughout the transmission. Decode the structure of your hexadecimal-encoded BITS transmission")
	log.Println(log.NORMAL, " Answer: what do you get if you add up the version numbers in all packets?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	bitReader := NewBitReader(input)
	p, err := parsePacket(&bitReader, 0)
	check(err)

	scanner := bufio.NewScanner(bitReader.file)
	if scanner.Scan() {
		bytesLeft := len(scanner.Bytes())
		log.Printf(log.NORMAL, color.Yellow+"Completed reading from file with %d bytes remaining!\n"+color.Reset, bytesLeft)
	}

	log.Printf(log.NORMAL, "\nTotal of all versions is %d\n", p.version+p.subVersions)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Do the same thing but perform operations and stuff")
	log.Println(log.NORMAL, " Answer: What do you get if you evaluate the expression represented by your hexadecimal-encoded BITS transmission?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	bitReader := NewBitReader(input)
	p, err := parsePacket(&bitReader, 0)
	check(err)

	scanner := bufio.NewScanner(bitReader.file)
	if scanner.Scan() {
		bytesLeft := len(scanner.Bytes())
		log.Printf(log.NORMAL, color.Yellow+"\nCompleted reading from file with %d bytes remaining!\n"+color.Reset, bytesLeft)
	}

	log.Printf(log.NORMAL, "\nTotal of all versions is %d\n", p.value)
}

func parseInputFile() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	output, err := files.CreateOrReplace("./output.txt")
	check(err)
	defer output.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		for _, b := range bytes {
			outBits, err := bits.HexByteToBits(b)
			if err != nil {
				panic(fmt.Sprintf("Error parsing byte %s", string(b)))
			}
			for _, ob := range outBits {
				if ob {
					output.WriteString("1")
				} else {
					output.WriteString("0")
				}
			}
		}
	}
}

type Packet struct {
	id             int
	value          int
	version        int
	subVersions    int
	pType          int
	subpacketCount int
	bitCount       int
	subBitCount    int
}

const (
	SUM         = 0
	PRODUCT     = 1
	MINIMUM     = 2
	MAXIMUM     = 3
	LITERAL     = 4
	GREATERTHAN = 5
	LESSTHAN    = 6
	EQUALTO     = 7
)

var packetNum = 0

func parsePacket(bitReader *BitReader, depth int) (Packet, error) {
	packetNum++
	p := Packet{packetNum, 0, 0, 0, 0, 0, 0, 0}
	versionBits, err := bitReader.ReadBits(3)
	p.bitCount += 3
	if err != nil {
		return p, fmt.Errorf("%s for version bits", err.Error())
	}
	p.version = bits.ToInt(versionBits)
	typeBits, err := bitReader.ReadBits(3)
	p.bitCount += 3
	if err != nil {
		return p, fmt.Errorf("%s for type bits", err.Error())
	}
	p.pType = bits.ToInt(typeBits)

	// log.Printf(log.DIAGNOSTIC, "\nP: %3d %s", packetNum, strings.Repeat(" ", depth*2))
	// log.Printf(log.DIAGNOSTIC, "Ver %d, Type %d", p.version, p.pType)

	switch p.pType {
	case LITERAL:
		// log.Printf(log.DIAGNOSTIC, ": L")
		var litBits []bool
		var bitVals []bool
		// read sets of 5 bits until the first bit of 5 is not a '1'
		for bitVals, err = bitReader.ReadBits(5); bitVals[0]; bitVals, err = bitReader.ReadBits(5) {
			p.bitCount += 5
			if err != nil {
				return p, fmt.Errorf("%s for literal bits", err.Error())
			}
			litBits = append(litBits, bitVals[1:]...)
		}
		// add the last set to the literal value
		p.bitCount += 5
		litBits = append(litBits, bitVals[1:]...)
		p.value = bits.ToInt(litBits)
		log.Printf(log.DIAGNOSTIC, "%d", p.value)

		return p, nil
	case SUM, PRODUCT, MINIMUM, MAXIMUM, GREATERTHAN, LESSTHAN, EQUALTO:
		// log.Printf(log.DIAGNOSTIC, ": O")
		log.Printf(log.DIAGNOSTIC, operString(p.pType))
		hasSubpackets, err := bitReader.ReadBit()
		p.bitCount += 1
		if err != nil {
			return p, fmt.Errorf("%s for subpacket bit", err.Error())
		}
		if hasSubpackets {
			countBits, err := bitReader.ReadBits(11)
			p.bitCount += 11
			if err != nil {
				return p, fmt.Errorf("%s for subpacket count", err.Error())
			}
			p.subpacketCount = bits.ToInt(countBits)
			// log.Printf(log.DIAGNOSTIC, " (count: %d) {", p.subpacketCount)

			for subpacketTotal, accum := 0, getAccum(p.pType); subpacketTotal < p.subpacketCount; {
				sp, err := parsePacket(bitReader, depth+1)
				// log.Printf(log.DIAGNOSTIC, " [S: %d]", subpacketTotal+1)
				if err != nil {
					return p, fmt.Errorf("%s for subpacket %d of id#%d", err.Error(), subpacketTotal+1, p.id)
				}

				if subpacketTotal == 0 && (p.pType == GREATERTHAN || p.pType == LESSTHAN || p.pType == EQUALTO) {
					accum = sp.value
				} else {
					p.value = mungeValue(accum, p.pType, sp.value)
				}

				subpacketTotal++
				p.subVersions += sp.version + sp.subVersions
				p.subBitCount += sp.bitCount + sp.subBitCount

				if subpacketTotal < p.subpacketCount {
					log.Printf(log.DIAGNOSTIC, ", ")
				}
			}
		} else {
			lengthBits, err := bitReader.ReadBits(15)
			p.bitCount += 15
			if err != nil {
				return p, fmt.Errorf("%s for subpacket length", err.Error())
			}
			p.subBitCount = bits.ToInt(lengthBits)
			//log.Printf(log.DIAGNOSTIC, " (length: %d) {", p.subBitCount)
			subpacketBits := 0
			for accum := getAccum(p.pType); subpacketBits < p.subBitCount; {
				sp, err := parsePacket(bitReader, depth+1)
				//log.Printf(log.DIAGNOSTIC, " [S: %d]", p.subpacketCount+1)
				if err != nil {
					return p, fmt.Errorf("%s for subpacket of id#%d", err.Error(), p.id)
				}

				if subpacketBits == 0 && (p.pType == GREATERTHAN || p.pType == LESSTHAN || p.pType == EQUALTO) {
					accum = sp.value
				} else {
					p.value = mungeValue(accum, p.pType, sp.value)
				}

				subpacketBits += sp.bitCount + sp.subBitCount
				p.subVersions += sp.version + sp.subVersions
				p.subpacketCount += 1 + sp.subpacketCount

				if subpacketBits < p.subBitCount {
					log.Printf(log.DIAGNOSTIC, ", ")
				}
			}
			if subpacketBits != p.subBitCount {
				return p, fmt.Errorf("packet #%d should parse %d bits but parsed %d", p.id, p.subBitCount, subpacketBits)
			}
		}
		log.Printf(log.DIAGNOSTIC, ")")
	}

	// if p.pType != 4 {
	// 	log.Printf(log.DIAGNOSTIC, "\n%s}", strings.Repeat(" ", 6+depth*2))
	// }
	return p, nil
}

// get the starting accumulator for a given operation
func getAccum(oper int) int {
	switch oper {
	case SUM:
		return 0
	case PRODUCT:
		return 1
	case MINIMUM:
		return -math.MaxInt
	case MAXIMUM:
		return math.MinInt
	case GREATERTHAN:
		return 1
	case LESSTHAN:
		return 1
	case EQUALTO:
		return 1
	default:
		panic(fmt.Sprintf("Got invalid operation %d", oper))
	}
}

// perform an operation on the given packet and return a new value
func mungeValue(oldVal, oper, newVal int) int {
	switch oper {
	case SUM:
		return oldVal + newVal
	case PRODUCT:
		return oldVal * newVal
	case MINIMUM:
		return numbers.Min(oldVal, newVal)
	case MAXIMUM:
		return numbers.Max(oldVal, newVal)
	case GREATERTHAN:
		return numbers.BoolToInt(oldVal > newVal)
	case LESSTHAN:
		return numbers.BoolToInt(oldVal < newVal)
	case EQUALTO:
		return numbers.BoolToInt(oldVal == newVal)
	default:
		panic(fmt.Sprintf("Got invalid operation %d", oper))
	}
}
func operString(oper int) string {
	switch oper {
	case SUM:
		return "SUM("
	case PRODUCT:
		return "PROD("
	case MINIMUM:
		return "MIN("
	case MAXIMUM:
		return "MAX("
	case GREATERTHAN:
		return "GT("
	case LESSTHAN:
		return "LT("
	case EQUALTO:
		return "EQ("
	default:
		panic(fmt.Sprintf("Got invalid operation %d", oper))
	}
}

type BitReader struct {
	file   *os.File
	values []bool
}

func NewBitReader(file *os.File) BitReader {
	return BitReader{file, []bool{}}
}

var totalBytesRead = 0

func (b *BitReader) ReadBits(num int) ([]bool, error) {
	if len(b.values) < num {
		// each char is worth 4 bits
		byteCount := int(math.Ceil(float64(num-len(b.values)) / 4))
		bytes := make([]byte, byteCount)
		b.file.Read(bytes)
		log.Printf(log.NEVER, "\n [Reading in %d more byte(s) to fulfill %d bit read]", byteCount, num)
		totalBytesRead += byteCount
		err := b.AddToValues(bytes)
		if err != nil {
			return nil, fmt.Errorf("%s while reading %d bits", err.Error(), num)
		}
	}

	// read the rest from input and set any remainder
	bits := b.values[:num]
	b.values = b.values[num:]
	log.Printf(log.NEVER, " \n [Read %d bits from reader and left %d remainder values]", len(bits), len(b.values))

	return bits, nil
}

func (b *BitReader) ReadBit() (bool, error) {
	bits, err := b.ReadBits(1)
	if err == nil {
		return bits[0], nil
	} else {
		return false, err
	}
}

func (b *BitReader) AddToValues(bytes []byte) error {
	for bIdx, byte := range bytes {
		newBits, err := bits.HexByteToBits(byte)
		if err != nil {
			return fmt.Errorf("%s @ byte #%d", err.Error(), totalBytesRead-len(bytes)+bIdx+1)
		}
		b.values = append(b.values, newBits...)
	}
	return nil
}
