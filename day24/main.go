package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sknoslo/aoc2024/stacks"
	"sknoslo/aoc2024/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

type GateType byte

const (
	Constant GateType = iota
	And
	Or
	Xor
)

type Gate struct {
	t    GateType
	l, r string // for logic gates
	v    uint64 // for constant gates
}

func ConstantGate(val uint64) Gate { return Gate{t: Constant, v: val} }
func AndGate(l, r string) Gate     { return Gate{t: And, l: l, r: r} }
func OrGate(l, r string) Gate      { return Gate{t: Or, l: l, r: r} }
func XorGate(l, r string) Gate     { return Gate{t: Xor, l: l, r: r} }

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

var dot = flag.Bool("dot", false, "output dot file for graphiz")

func main() {
	flag.Parse()
	if *dot {
		genDot()
	}
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() (map[string]Gate, int) {
	out := make(map[string]Gate)

	constants, gates, _ := strings.Cut(input, "\n\n")

	zs := 0

	for _, constant := range strings.Split(constants, "\n") {
		wire, val, _ := strings.Cut(constant, ": ")
		out[wire] = ConstantGate(uint64(utils.MustAtoi(val)))
	}

	for _, gate := range strings.Split(gates, "\n") {
		parts := strings.Fields(strings.Replace(gate, " -> ", " ", 1))
		if parts[3][0] == 'z' {
			zs++
		}
		switch parts[1] {
		case "AND":
			out[parts[3]] = AndGate(parts[0], parts[2])
		case "OR":
			out[parts[3]] = OrGate(parts[0], parts[2])
		case "XOR":
			out[parts[3]] = XorGate(parts[0], parts[2])
		}
	}

	return out, zs
}

func pad(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}

	return strconv.Itoa(i)
}

func getbit(wire string, wires map[string]Gate) uint64 {
	gate := wires[wire]

	switch gate.t {
	case Constant:
		return gate.v
	case And:
		return getbit(gate.l, wires) & getbit(gate.r, wires)
	case Or:
		return getbit(gate.l, wires) | getbit(gate.r, wires)
	case Xor:
		return getbit(gate.l, wires) ^ getbit(gate.r, wires)
	default:
		panic("impossible")
	}
}

func partone() string {
	wires, zs := parseInput()
	result := uint64(0)
	for z := zs - 1; z >= 0; z-- {
		zwire := "z" + pad(z)
		bit := getbit(zwire, wires)
		result <<= 1
		result |= bit
	}
	return strconv.FormatUint(result, 10)
}

type Node struct {
	label, wire string
}

func genDot() {
	f, err := os.Create("circuit.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.WriteString(f, "digraph {\n")
	wires, zs := parseInput()
	seen := set.New[Node](64)
	for z := zs - 1; z >= 0; z-- {
		stack := stacks.New[Node](64)
		zwire := "z" + pad(z)
		stack.Push(Node{zwire, zwire})
		io.WriteString(f, "subgraph {\n")
		io.WriteString(f, "cluster=true;\n")
		fmt.Fprintf(f, "\"%s\";\n", zwire)
		for !stack.Empty() {
			n := stack.Pop()
			if seen.Contains(n) {
				continue
			}
			seen.Insert(n)
			gate := wires[n.wire]
			switch gate.t {
			case Constant:
				fmt.Fprintf(f, "%s -> %s\n", n.wire, n.label)
			case And:
				label := fmt.Sprintf("%s_AND_%s", gate.l, gate.r)
				fmt.Fprintf(f, "%s -> %s\n", label, n.label)
				stack.Push(Node{label, gate.l})
				stack.Push(Node{label, gate.r})
			case Or:
				label := fmt.Sprintf("%s_OR_%s", gate.l, gate.r)
				fmt.Fprintf(f, "%s -> %s\n", label, n.label)
				stack.Push(Node{label, gate.l})
				stack.Push(Node{label, gate.r})
			case Xor:
				label := fmt.Sprintf("%s_XOR_%s", gate.l, gate.r)
				fmt.Fprintf(f, "%s -> %s\n", label, n.label)
				stack.Push(Node{label, gate.l})
				stack.Push(Node{label, gate.r})
			}
		}
		io.WriteString(f, "\n}\n")
	}
	io.WriteString(f, "\n}\n")
}

func verifyAdder(out string, wires map[string]Gate) ([]string, bool) {
	bad := make([]string, 0, 2)
	good := true
	if out == "z00" || out == "z45" || out == "z01" {
		// should verify these for real, but the structure is different and I know they are right in my input
		return bad, good
	}

	xor0 := wires[out]

	if xor0.t != Xor {
		bad = append(bad, out)
		return bad, false
	}

	leftl, rightl := xor0.l, xor0.r
	xor1, or1 := wires[leftl], wires[rightl]
	if or1.t == Xor && xor1.t == Or || or1.t == Xor || xor1.t == Or {
		leftl, rightl = rightl, leftl
		or1, xor1 = xor1, or1
	}

	if xor1.t != Xor {
		bad = append(bad, leftl)
		good = false
	}

	if or1.t != Or {
		bad = append(bad, rightl)
		good = false
	}

	if !good {
		return bad, false
	}

	and21, and22 := wires[or1.l], wires[or1.r]

	if and21.t != And {
		bad = append(bad, or1.l)
		good = false
	}

	if and22.t != And {
		bad = append(bad, or1.r)
		good = false
	}

	return bad, good
}

func parttwo() string {
	wires, zs := parseInput()

	bad := make([]string, 0, 8)
	for z := range zs - 1 {
		zwire := "z" + pad(z)
		if badWires, ok := verifyAdder(zwire, wires); !ok {
			bad = append(bad, badWires...)
		}
	}

	slices.Sort(bad)
	return strings.Join(bad, ",")
}
