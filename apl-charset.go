// ISO/IEC 2000 APL Required Character Set
// see also /usr/share/X11/xkb/symbols/apl
package main

import "fmt"

var charmap = map[rune]string{
	'⍺': "alpha",
	'↓': "down arrow",
	'←': "left arrow",
	'→': "right arrow",
	'↑': "up arrow",
	'-': "bar",
	' ': "blank",
	'{': "left brace",
	'}': "right brace",
	'[': "left bracket",
	']': "right bracket",
	'∨': "down caret",
	'⍱': "down caret tilde",
	'<': "left caret",
	'>': "right caret",
	'∧': "up caret",
	'⍲': "up caret tilde",
	'○': "circle",
	'⍉': "circle backslash",
	'⊖': "circle bar",
	'⍟': "circle star",
	'⌽': "circle stile",
	':': "colon",
	',': "comma",
	'⍪': "comma bar",
	'∇': "del",
	'⍒': "del stile",
	'⍫': "del tilde",
	'∆': "delta",
	'⍋': "delta stile",
	'⍙': "delta underbar",
	'¨': "diaeresis",
	'⋄': "diamond",
	'÷': "divide",
	'$': "dollar sign",
	'.': "dot",
	'∊': "epsion",
	'=': "equal",
	'≥': "greater-than or equal",
	'⍳': "iota",
	'∘': "jot",
	'≤': "less-than or equal",
	'×': "miltiply",
	'≠': "not equal",
	'¯': "overbar",
	'(': "left parenthesis",
	')': "right parenthesis",
	'+': "plus",
	'⎕': "quad",
	'⌹': "quad divide",
	'?': "query",
	'⍤': "diaeresis jot",
	'\'': "quote",
	'!': "quote dot",
	'⍞': "quate quad",
	'⍴': "rho",
	';': "semicolon",
	'∪': "down shoe",
	'⊂': "left shoe",
	'⊃': "right shoe",
	'∩': "up shoe",
	'⍝': "up shoe jot",
	'/': "slash",
	'\\': "back slash",
	'⌿': "slash bar",
	'⍀': "back slash bar",
	'*': "star",
	'|': "stile",
	'⌊': "down stile",
	'⌈': "up stile",
	'⊥': "up tack",
	'⍎': "up tack jot",
	'⊢': "right tack",
	'⊣': "left tack",
	'⊤': "down tack",
	'⍕': "down tack jot",
	'~': "tilde",
	'_': "underbar",
	'⍨': "diaeresis tilde",
	'⍵': "omega",
	'≡': "equal underbar",
}

func scan(chars string) {
	for _, c := range chars {
		if name, ok := charmap[c]; ok {
			fmt.Println(string(c), name)
		} else {
			fmt.Println(string(c), "... missing")
		}
	}
}

func main() {
	scan("⋄⌶⍫⍒⍋⌽⍉⊖⍟⍱⍲!⌹?⍵⍷⍴⍨↑↓⍸⍥⍣⍞⍬⍺⌈⌊_∇∆⍤'⌷≡≢⊣⍕⊂⊃∩∪⊥⊤|⍪⍙⍠")
	scan("⋄¨¯<≤=≥>≠∨∧×÷?⍵∊⍴~↑↓⍳○*←→⍺⌈⌊_∇∆∘'⎕⍎⍕⊢⍎⊂⊃∩∪⊥⊤|⍝⍀⌿")
}
