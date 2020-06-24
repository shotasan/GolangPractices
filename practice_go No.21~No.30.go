// No.21
// Consider the word "abode". We can see that the letter a is in position 1 and b is in position 2. In the alphabet, a and b are also in positions 1 and 2. Notice also that d and e in abode occupy the positions they would occupy in the alphabet, which are positions 4 and 5.
// Given an array of words, return an array of the number of letters that occupy their positions in the alphabet for each word. For example,
// solve(["abode","ABc","xyzD"]) = [4, 3, 1]
// See test cases for more examples.
// Input will consist of alphabet characters, both uppercase and lowercase. No spaces.

// My
import "strings"

func solve(slice []string) []int {
  res := make([]int, len(slice))
  for i, v := range slice {
    tmp := 0
    for j, s := range strings.ToLower(v) {
      if j + 1 == int(s) - 96 {
        tmp++
      }
    }
    res[i] = tmp
  }
  return res
}

// Best
func solve(slice []string) []int {
  result := make([]int, len(slice))

  for i, str := range slice {
    count := 0

    for i := 0; i < len(str); i++ {
      if int(str[i]) & 31 == i+1 {
        count++
      }
    }

    result[i] = count
  }

  return result
}