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

// No.22
// Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.
// Note: no empty arrays will be given.
// Examples
// [12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
// [12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
// [12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3

// Best
func HighestRank(nums []int) int {
  mii, maxK, maxV := map[int]int{}, 0, 0
  for _, v := range nums {
    mii[v]++
    if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
      maxK = v
      maxV = mii[v]
    }
  }
  return maxK
}

// No.23
// A Cartesian coordinate system is a coordinate system that specifies each point uniquely in a plane by a pair of numerical coordinates, which are the signed distances to the point from two fixed perpendicular directed lines, measured in the same unit of length.
// The сoordinates of a point in the grid are written as (x,y). Each point in a coordinate system has eight neighboring points. Provided that the grid step = 1.
// It is necessary to write a function that takes a coordinate on the x-axis and y-axis and returns a list of all the neighboring points. Points inside list don't have to been sorted (any order is valid).
// For Example:
// CartesianNeighbor(2,2) -> {{1,1},{1,2},{1,3},{2,1},{2,3},{3,1},{3,2},{3,3}}
// CartesianNeighbor(5,7) -> {{6,7},{6,6},{6,8},{4,7},{4,6},{4,8},{5,6},{5,8}}

// My
func CartesianNeighbor(x,y int) [][]int{
  var res [][]int
  for i := -1; i <= 1; i++ {
    newX := x + i
    for j := -1; j <= 1; j++ {
      newY := y + j
      if newX == x && newY == y {
        continue
      }
      res = append(res, []int{newX, newY})
    }
  }
  return res
}

// Best
func CartesianNeighbor(x,y int) [][]int{
  res := make ([][]int, 0, 8)
  for xd := -1; xd <= 1; xd++ {
    for yd := -1; yd <= 1; yd++ {
      if xd != 0 || yd != 0 {
        res = append(res, []int{x+xd, y+yd})
      }
    }
  }
  return res
}

// No.24
// Complete the function so that it finds the mean of the three scores passed to it and returns the letter value associated with that grade.
// Numerical Score	Letter Grade
// 90 <= score <= 100	'A'
// 80 <= score < 90	'B'
// 70 <= score < 80	'C'
// 60 <= score < 70	'D'
// 0 <= score < 60	'F'
// Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.

// My
func GetGrade(a,b,c int) rune {
    switch avg := (a + b + c) / 3; {
      case 90 <= avg && avg <= 100:
        return 'A'
      case 80 <= avg && avg < 90:
        return 'B'
      case 70 <= avg && avg < 80:
        return 'C'
      case 60 <= avg && avg < 70:
        return 'D'
      default:
        return 'F'
    } 
}

// Best
func GetGrade(a, b, c int) rune {
    switch ((a + b + c) / 30) {
        case 10: return 'A'
        case 9: return 'A'
        case 8: return 'B'
        case 7: return 'C'
        case 6: return 'D'
        default: return 'F'
    }
}