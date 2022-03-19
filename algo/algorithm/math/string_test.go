package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	ass:= assert.New(t)
	arr := []string{"m","mo","moc","moch","mocha","l","la","lat","latt","latte","c","ca","cat"}
	arr1 := []string{"t","ti","tig","tige","tiger","e","en","eng","engl","engli","englis","english","h","hi","his","hist","histo","histor","history"}
	ass.Equal("latte", longestWord(arr))
	ass.Equal("english", longestWord(arr1))

}

func TestReverseVowels(t *testing.T) {
	ass:= assert.New(t)
	ass.Equal("leotcede", reverseVowels("leetcode"))

}

