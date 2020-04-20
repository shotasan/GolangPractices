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