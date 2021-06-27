// Code generated from Dice.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Dice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiceListener is a complete listener for a parse tree produced by DiceParser.
type BaseDiceListener struct{}

var _ DiceListener = &BaseDiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDiceListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDiceListener) ExitStart(ctx *StartContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseDiceListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseDiceListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseDiceListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseDiceListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseDiceListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseDiceListener) ExitAddSub(ctx *AddSubContext) {}

// EnterDic is called when production Dic is entered.
func (s *BaseDiceListener) EnterDic(ctx *DicContext) {}

// ExitDic is called when production Dic is exited.
func (s *BaseDiceListener) ExitDic(ctx *DicContext) {}
