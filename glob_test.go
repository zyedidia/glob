package glob

import (
	"regexp"
	"testing"
)

func assertEquals(t *testing.T, expected string, result *Glob) {
	r, _ := regexp.Compile(expected)
	if r.String() != result.Regexp.String() {
		t.Fatal("Expected ", expected, ", got ", result.Regexp, "\n")
	}
}

func TestStarBecomesDotStar(t *testing.T) {
	result, _ := Compile("gl*b")
	assertEquals(t, "gl.*b", result)
}

func TestEscapedStarIsUnchanged(t *testing.T) {
	result, _ := Compile("gl\\*b")
	assertEquals(t, "gl\\*b", result)
}

func TestQuestionMarkBecomesDot(t *testing.T) {
	result, _ := Compile("gl?b")
	assertEquals(t, "gl.b", result)
}

func TestEscapedQuestionMRkIsUnchanged(t *testing.T) {
	result, _ := Compile("gl\\?b")
	assertEquals(t, "gl\\?b", result)
}

func TestCharacterClassesDontNeedConversion(t *testing.T) {
	result, _ := Compile("gl[-o]b")
	assertEquals(t, "gl[-o]b", result)
}

func TestEscapedClassesAreUnchanged(t *testing.T) {
	result, _ := Compile("gl\\[-o\\]b")
	assertEquals(t, "gl\\[-o\\]b", result)
}

func TestNegationInCharacterClasses(t *testing.T) {
	result, _ := Compile("gl[!a-n!p-z]b")
	assertEquals(t, "gl[^a-n!p-z]b", result)
}

func TestNestedNegationInCharacterClasses(t *testing.T) {
	result, _ := Compile("gl[[!a-n]!p-z]b")
	assertEquals(t, "gl[[^a-n]!p-z]b", result)
}

func TestEscapeCaratIfItIsTheFirstCharInACharacterClass(t *testing.T) {
	result, _ := Compile("gl[^o]b")
	assertEquals(t, "gl[\\^o]b", result)
}

func TestMetacharsAreEscaped(t *testing.T) {
	result, _ := Compile("gl?*.()+|^$@%b")
	assertEquals(t, "gl..*\\.\\(\\)\\+\\|\\^\\$\\@\\%b", result)
}

func TestMetacharsInCharacterClassesDontNeedEscaping(t *testing.T) {
	result, _ := Compile("gl[?*.()+|^$@%]b")
	assertEquals(t, "gl[?*.()+|^$@%]b", result)
}

func TestEscapedBackslashIsUnchanged(t *testing.T) {
	result, _ := Compile("gl\\\\b")
	assertEquals(t, "gl\\\\b", result)
}

func TestSlashQAndSlashEAreEscaped(t *testing.T) {
	result, _ := Compile("\\Qglob\\E")
	assertEquals(t, "\\\\Qglob\\\\E", result)
}

func TestBracesAreTurnedIntoGroups(t *testing.T) {
	result, _ := Compile("{glob,regex}")
	assertEquals(t, "(glob|regex)", result)
}

func TestEscapedBracesAreUnchanged(t *testing.T) {
	result, _ := Compile("\\{glob\\}")
	assertEquals(t, "\\{glob\\}", result)
}

func TestCommasDontNeedEscaping(t *testing.T) {
	result, _ := Compile("{glob\\,regex},")
	assertEquals(t, "(glob,regex),", result)
}
