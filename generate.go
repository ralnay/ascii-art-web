package piscine

import (
	"bufio"
	"os"
	"strings"
)

func Load(file string) map[rune][]string {
	font, error := os.Open(file)
	if error != nil {
		return nil
	}
	defer font.Close()
	scanner := bufio.NewScanner(font)
	table := make(map[rune][]string)
	lines := make([]string, 8)
	i := ' '
	for j := 0; scanner.Scan(); j++ {
		line := scanner.Text()
		if j == 0 {
			lines = append(lines, line)
		}
		if j != 0 && j%9 != 0 {
			lines = append(lines, line)
		}
		if j != 0 && j%9 == 0 {
			table[i] = lines
			lines = nil
			i++
		}
	}
	if len(lines) > 0 {
		table[i] = lines
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return table
}

func PrintOutput(t map[rune][]string, str string) string {
	output := ""
	i := 0
	words := strings.Split(str, "\n")
	for j := 0; j < len(str); j++ {
		if i < len(words) {
			output += Recursion(words[i], t)
		}
		i++
		if i < len(words)-1 {
			j = j + len(words[i])
		}
	}
	return output
}
func Recursion(word string, t2 map[rune][]string) string {
	out := ""
	if word == "" {
		out = "\n"
		return out
	}
	if word == "\r" {
		out = ""
		return out
	}
	for j := 0; j < 8; j++ {
		for _, ch := range word {
			if ch == ' ' {
				for n := 1; n <= 4; n++ {
					out += string(ch)
				}
				continue
			}
			if lines, ok := t2[ch]; ok {
				out += lines[j]
			} else {
				out += "        " // Placeholder for missing characters
			}
		}
		out += "\n"
	}
	return out
}
