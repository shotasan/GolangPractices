// No.11
// Our football team finished the championship. The result of each match look like "x:y". Results of all matches are recorded in the collection.
// For example: ["3:1", "2:2", "0:1", ...]
// Write a function that takes such collection and counts the points of our team in the championship. Rules for counting points for each match:
// if x>y - 3 points
// if x<y - 0 point
// if x=y - 1 point
// Notes:
// there are 10 matches in the championship
// 0 <= x <= 4
// 0 <= y <= 4

// My
func Points(games []string) int {
  var result int
  for _, v := range games {
    if v[0] > v[2] {
      result += 3
    } else if v[0] == v[2] {
      result += 1
    }
  }
  return result;
}

// Best
import ( "strings" )

func Points(games []string) int {
  result := 0
  for _, game := range games {
    str := strings.Split(game, ":")
    x, y := str[0], str[1]
    switch {
      case x > y:
        result += 3
      case x == y:
        result += 1
    }
  }
  return result
}