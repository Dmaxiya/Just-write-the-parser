package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicTokenization(t *testing.T) {
	a := assert.New(t)

	a.Equal(+ (8 - 12/3), ParsingByASinglePass("8-12/3"))
	a.Equal(2*(6+4)*5+5, ParsingByASinglePass("2 * ( 6 + 4 ) * 5 + 5 "))
	a.Equal((10+97)*(25*5), ParsingByASinglePass("(   10 +  97  ) * (  25 * 5 )"))
	a.Equal((10+97)*(25*5)+55, ParsingByASinglePass("(   10 +  97  	) * (  25 * 5 ) + 55"))
	a.Equal(92*(12)*(1+(1)), ParsingByASinglePass("92 * ( 12 ) * ( 1 + ( 1 ))"))
	a.Equal((10+97)*(25*5)+5, ParsingByASinglePass(` (   10 +  97  )
			* (  25 * 5 ) +
			5  `))
	a.Equal((36)*(74)*(34+33), ParsingByASinglePass(`// hello world
			( 36 # this is a number 36
			// and bracket
			# emmmmmm.... and some meaningless comments
			) * ( 74     	)
			*  (
			# hello world // hello world again
			# hello world again and again
			34+
			/*
			*/
			33)
			// over`))
	a.Equal(8, ParsingByASinglePass(`2 /* first comment /* second comment */ * 2 * 2`))
	a.Equal(64*10*(72+44*86), ParsingByASinglePass(`64
			/*hello /*hello /*hello // world3 world2 world1*/
			# /*    //
			/* /* */ *10 // hello world */
			*(72+44* /* there is a delimited comment between '*' and 86 // */ 86)`))
	a.Equal(12+(56/7)*10/16+3, ParsingByASinglePass(`12+(56/7)*10/16+3`))
}
