// Code generated from Dice.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Dice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiceListener is a complete listener for a parse tree produced by DiceParser.
type DiceListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterDic is called when entering the Dic production.
	EnterDic(c *DicContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitDic is called when exiting the Dic production.
	ExitDic(c *DicContext)
}
