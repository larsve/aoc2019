package main

type opCode byte

type memoryReaderWriter interface {
	lda(adress int) int
	sta(address, value int)
}

const (
	opErr opCode = iota
	opAdd
	opMul
	opExit
)

var opCodeStrings = map[opCode]string{opAdd: "ADD", opMul: "MUL", opErr: "ERR", opExit: "BRK"}

func (o opCode) String() string {
	if s, ok := opCodeStrings[o]; ok {
		return s
	}
	return "Unknown OpCode"
}

func (o opCode) execute(parameters []int, m memoryReaderWriter) error {
	if (o == opErr) || (o == opExit) {
		return nil
	}
	p1 := m.lda(parameters[0])
	p2 := m.lda(parameters[1])
	resultAddress := parameters[2]

	switch o {
	case opAdd:
		m.sta(resultAddress, p1+p2)
	case opMul:
		m.sta(resultAddress, p1*p2)
	}
	return nil
}
