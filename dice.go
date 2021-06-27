package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
    "time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/argot42/dice/parser"
)

func main() {
    var r *bufio.Reader

    if len(os.Args) > 1 {
        args := strings.Join(os.Args[1:], "\n")
        r = bufio.NewReader(strings.NewReader(args))
    } else {
        r = bufio.NewReader(os.Stdin)
    }

    for {
        line, err := r.ReadString('\n')
        if err != nil {
            fmt.Fprint(os.Stderr, err)
            os.Exit(1)
        }

        n := roll(line)

        fmt.Println(n)
    }
}

func roll(in string) int {
    is := antlr.NewInputStream(in)

    lexer := parser.NewDiceLexer(is)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

    p := parser.NewDiceParser(stream)

    var listener DiceListener
    listener.Setup()

    antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

    return listener.Stack.Pop()
}

type DiceListener struct {
    *parser.BaseDiceListener
    Stack Stack
    Random *rand.Rand
}

type Stack []int

func (s *Stack) Push(n int) {
    *s = append(*s, n)
}

func (s *Stack) Pop() (n int) {
    if len(*s) < 1 {
        panic("error: Stack is empty")
    }

    n = (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]

    return
}

func (l *DiceListener) Setup() {
    src := rand.NewSource(time.Now().UnixNano())
    l.Random = rand.New(src)
}

func (l *DiceListener) ExitDic(ctx *parser.DicContext) {
    dType, dAmount := l.Stack.Pop(), l.Stack.Pop()

    n := 0
    for i := 0; i < dAmount; i++ {
        n += l.Random.Intn(dType-1) + 1
    }

    l.Stack.Push(n)
}

func (l *DiceListener) ExitMulDiv(ctx *parser.MulDivContext) {
    right, left := l.Stack.Pop(), l.Stack.Pop()

    switch ctx.GetOp().GetTokenType() {
    case parser.DiceParserMUL:
        l.Stack.Push(left * right)
    case parser.DiceParserDIV:
        l.Stack.Push(left / right)
    default:
        panic(fmt.Sprintf("error: unexpected op: %s", ctx.GetOp().GetText()))
    }
}

func (l *DiceListener) ExitAddSub(ctx *parser.AddSubContext) {
    right, left := l.Stack.Pop(), l.Stack.Pop()

    switch ctx.GetOp().GetTokenType() {
    case parser.DiceParserADD:
        l.Stack.Push(left + right)
    case parser.DiceParserSUB:
        l.Stack.Push(left - right)
    default:
        panic(fmt.Sprintf("error: unexpected op: %s", ctx.GetOp().GetText()))
    }
}

func (l *DiceListener) ExitNumber(ctx *parser.NumberContext) {
    n, err := strconv.Atoi(ctx.GetText())
    if err != nil {
        panic(err)
    }

    l.Stack.Push(n)
}
