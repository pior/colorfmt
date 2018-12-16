package colorfmt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func runIt(enableColor bool, text string, a ...interface{}) string {
	buf := new(bytes.Buffer)
	New(buf, enableColor).Printf(text, a...)
	return buf.String()
}

func TestTags(t *testing.T) {
	type testRun struct {
		text     string
		expected string
	}

	tests := []testRun{
		testRun{"_", "\x1b[0m_"},

		testRun{"{yellow}_", "\x1b[0m\x1b[0;33m_\x1b[0m"},
		testRun{"{yellow+h:blue+h}_", "\x1b[0m\x1b[0;93;104m_\x1b[0m"},

		testRun{"{notAColor}_", "\x1b[0m\x1b[0;30m_\x1b[0m"},
		testRun{"{with space}_", "\x1b[0m\x1b[0;30m_\x1b[0m"},

		testRun{"{link}_", "\x1b[0m\x1b[0;1;4;92m_\x1b[0m"},
	}

	for _, run := range tests {
		require.Equal(t, run.expected, runIt(true, run.text), "with color, format='%s'", run.text)
		require.Equal(t, "_", runIt(false, run.text), "without color, format='%s'", run.text)
	}
}

func TestVariadic(t *testing.T) {
	require.Equal(t,
		"\x1b[0m\x1b[0;33m_1_\x1b[0;32m_2_\x1b[0m",
		runIt(true, "{yellow}%s{green}%s", "_1_", "_2_"),
	)
}

func TestEscapeBracket(t *testing.T) {
	require.Equal(t, "\x1b[0m_{yellow}}_", runIt(true, "_{{yellow}}_"))
}

func TestMultipleTags(t *testing.T) {
	require.Equal(t, "\x1b[0m_\x1b[0;33m_\x1b[0;32m_\x1b[0m", runIt(true, "_{yellow}_{green}_"))
}

func TestReset(t *testing.T) {
	require.Equal(t, "\x1b[0m\x1b[0;33m_\x1b[0m_\x1b[0m", runIt(true, "{yellow}_{reset}_"))
}

func TestForegroundColors(t *testing.T) {
	require.Equal(t, "\x1b[0m\x1b[0;30m\x1b[0m", runIt(true, "{NOT A COLOR}"))
	require.Equal(t, "\x1b[0m\x1b[0;30m\x1b[0m", runIt(true, "{black}"))
	require.Equal(t, "\x1b[0m\x1b[0;31m\x1b[0m", runIt(true, "{red}"))
	require.Equal(t, "\x1b[0m\x1b[0;32m\x1b[0m", runIt(true, "{green}"))
	require.Equal(t, "\x1b[0m\x1b[0;33m\x1b[0m", runIt(true, "{yellow}"))
	require.Equal(t, "\x1b[0m\x1b[0;34m\x1b[0m", runIt(true, "{blue}"))
	require.Equal(t, "\x1b[0m\x1b[0;35m\x1b[0m", runIt(true, "{magenta}"))
	require.Equal(t, "\x1b[0m\x1b[0;36m\x1b[0m", runIt(true, "{cyan}"))
	require.Equal(t, "\x1b[0m\x1b[0;37m\x1b[0m", runIt(true, "{white}"))
}

func TestForegroundStyles(t *testing.T) {
	require.Equal(t, "\x1b[0m\x1b[0;1;30m\x1b[0m", runIt(true, "{+b}"))
	require.Equal(t, "\x1b[0m\x1b[0;1;37m\x1b[0m", runIt(true, "{white+b}"))
	require.Equal(t, "\x1b[0m\x1b[0;4;37m\x1b[0m", runIt(true, "{white+u}"))
	require.Equal(t, "\x1b[0m\x1b[0;5;37m\x1b[0m", runIt(true, "{white+B}"))
	require.Equal(t, "\x1b[0m\x1b[0;7;37m\x1b[0m", runIt(true, "{white+i}"))
	require.Equal(t, "\x1b[0m\x1b[0;9;37m\x1b[0m", runIt(true, "{white+s}"))
	require.Equal(t, "\x1b[0m\x1b[0;97m\x1b[0m", runIt(true, "{white+h}"))
}

func TestBackroundStyles(t *testing.T) {
	require.Equal(t, "\x1b[0m\x1b[0;30m\x1b[0m", runIt(true, "{:+h}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;107m\x1b[0m", runIt(true, "{:white+h}"))
}

func TestBackgroundColors(t *testing.T) {
	require.Equal(t, "\x1b[0m\x1b[0;30;40m\x1b[0m", runIt(true, "{:invalid}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;40m\x1b[0m", runIt(true, "{:black}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;41m\x1b[0m", runIt(true, "{:red}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;42m\x1b[0m", runIt(true, "{:green}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;43m\x1b[0m", runIt(true, "{:yellow}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;44m\x1b[0m", runIt(true, "{:blue}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;45m\x1b[0m", runIt(true, "{:magenta}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;46m\x1b[0m", runIt(true, "{:cyan}"))
	require.Equal(t, "\x1b[0m\x1b[0;30;47m\x1b[0m", runIt(true, "{:white}"))
}

func TestForgroundAndBackgroundColors(t *testing.T) {
	require.Equal(t, "\x1b[0m\x1b[0;37;40m\x1b[0m", runIt(true, "{white:black}"))
}
