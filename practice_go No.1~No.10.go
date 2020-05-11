// No.1
// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
// It should look like this:
// Sam Harris => S.H
// Patrick Feeney => P.F

// Best
package kata

import "strings"

func AbbrevName(name string) string{
  words := strings.Split(name, " ")
  return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
}

// My
package kata

import (
  "strings"
)

func AbbrevName(name string) string{
  strs := strings.Split(name, " ")
  var result []string
  for _, s := range strs {
    result = append(result, s[:1])
  }
  return strings.ToUpper(strings.Join(result, "."))
}

// No.2
// If you can't sleep, just count sheep!!
// Task:
// Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.

// Best
package kata

import (
  "fmt"
  "strings"
)

func countSheep(num int) string {
  var sb strings.Builder

  for count := 1; count <= num; count++ {
        fmt.Fprintf(&sb, "%d sheep...", count)
  }
  
  return sb.String()
}

// My
package kata

import (
  "strconv"
  "strings"
)

func countSheep(num int) string {
  var result []string
  for i := 1; i <= num; i++ {
    result = append(result, strconv.Itoa(i) + " sheep...")
  }
  return strings.Join(result, "")
}

// No.3
// Write function isPalindrome that checks if a given string (case insensitive) is a palindrome.
// In Racket, the function is called palindrome?
// (palindrome? "nope") ; returns #f
// (palindrome? "Yay")  ; returns #t

// Best
func IsPalindrome(s string) bool {
  runes := []rune(strings.ToLower(s))
  length := len(runes)

  for i := 0; i < length/2; i++ {
    if runes[i] != runes[length-1-i] {
      return false
    }
  }

  return true
}

// No.4
// Yor task is to write function factorial

package kata

func Factorial(n int) int {
  if n < 2 {
    return 1
  }
  
  return n * Factorial(n - 1)
}

// No.5
// In a factory a printer prints labels for boxes. For one kind of boxes the printer has to use colors which, for the sake of simplicity, are named with letters from a to m.
// The colors used by the printer are recorded in a control string. For example a "good" control string would be aaabbbbhaijjjm meaning that the printer used three times color a, four times color b, one time color h then one time color a...
// Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is produced e.g. aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.
// You have to write a function printer_error which given a string will output the error rate of the printer as a string representing a rational whose numerator is the number of errors and the denominator the length of the control string. Don't reduce this fraction to a simpler expression.
// The string has a length greater or equal to one and contains only letters from ato z.
// #Examples:
// s="aaabbbbhaijjjm"
// error_printer(s) => "0/14"
// s="aaaxbbbbyyhwawiwjjjwwm"
// error_printer(s) => "8/22"

// My
package kata

import (
  "regexp"
  "strconv"
)

func PrinterError(s string) string {
  re := regexp.MustCompile(`[n-z]`)
  badStrings := re.FindAllString(s, -1)
  badNumbers := len(badStrings)
  totalNumbers := len([]rune(s))
  return strconv.Itoa(badNumbers) + "/" + strconv.Itoa(totalNumbers)
}

// Best
package kata

import "fmt"

func PrinterError(s string) string {
  errors := 0
  for _, c := range s {
    if string(c) > "m" {
      errors++
    }
  }
  return fmt.Sprintf("%d/%d", errors, len(s))
}

// No.6
// Input:
// a string strng
// an array of strings arr
// Output of function contain_all_rots(strng, arr) (or containAllRots or contain-all-rots):
// a boolean true if all rotations of strng are included in arr (C returns 1)
// false otherwise (C returns 0)
// Examples:
// contain_all_rots(
//   "bsjq", ["bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"]) -> true
// contain_all_rots(
//   "Ajylvpy", ["Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"]) -> false)
// Note:
// Though not correct in a mathematical sense
// we will consider that there are no rotations of strng == ""
// and for any array arr: contain_all_rots("", arr) --> true

// Best
package kata

func ContainAllRots(strng string, arr []string) bool { 
    counter := 0
    master := []string{}
    for i,_ := range strng{
				// 簡易スライス式でrotationを実現する
        master = append(master,strng[i:]+strng[:i])
    }
    for _,j := range master{
        for _,k := range arr{
            if j == k{
                counter ++
                break
            }
        }
    }
    return len(strng) == counter
}

// No.7
// The vowel substrings in the word codewarriors are o,e,a,io.
// The longest of these has a length of 2. Given a lowercase string that has alphabetic characters only and no spaces, return the length of the longest vowel substring.
// Vowels are any of aeiou.

// Best
func isVowel(r rune) bool {
    switch r {
    case 'a', 'e', 'i', 'o', 'u':
        return true
    }
    return false
}

func Solve(s string) int {
    substr_length := 0
    max_length := 0

    for _, c := range s {
      if isVowel(c) {
        substr_length += 1
        if max_length < substr_length {
          max_length = substr_length
        }
      } else {
        substr_length = 0
      }
    }
    return max_length  
}

// My
import (
  "regexp"
)

func Solve(s string) int {
    re := regexp.MustCompile(`[a e i o u]+`)
    results := re.FindAllString(s, -1)
    max := 0
    for _, s := range results {
      if max < len(s) {
        max = len(s)
      }
    }
    return max
}

// No.8
// Write a function that takes a positive integer n, sums all the cubed values from 1 to n, and returns that sum.
// Assume that the input n will always be a positive integer.
// Examples:
// SumCubes(2) == 9
// sum of the cubes of 1 and 2 is 1 + 8

// Best
func SumCubes(n int) int {
  var sum int
  for i := 1; i <= n; i++ {
    sum += i * i * i
  }

  return sum
}

// My(false)
import "math"

func SumCubes(n int) int {
  var res float64
  for i := 1; i <= n; i++ {
    res += math.Pow(float64(i), 3)
  }
  return int(res)
}

// No.9
// Count the number of occurrences of each character and return it as a list of tuples in order of appearance.
// For empty output return an empty list.
// Example:
// OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}
// // Where
// type Tuple struct {
//     Char  rune
//     Count int
// }

// Best
import "strings"

func OrderedCount(text string) []Tuple {
  counts := make([]Tuple, 0, len(text))
  counted := ""
  
  for _, r := range(text) {
    if strings.Count(counted, string(r)) == 0 {
      counts = append(counts, Tuple{Char: r, Count: strings.Count(text, string(r))})
      counted += string(r)
    }
  }
  return counts 
}