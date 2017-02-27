package main

import (
	"fmt"
  "strings"
)

func main() {
	fmt.Printf("hello\n")
  Sort([]string{"a =>", "b => c", "c => f", "d => a", "e => b", "f =>"})
}

func Sort(jobs []string) string {
  result := ""
  deps := make(map[string]string)
  for _, job := range jobs {
    job = strings.Replace(job, " ", "", -1)
    splitJob := strings.Split(job, "=>")
    if splitJob[1] != "" {
      fmt.Printf("Adding dep %s => '%s' len %s\n", splitJob[0], splitJob[1], len(splitJob))
      deps[splitJob[0]] = splitJob[1]
    }
    result += splitJob[0]
  }

  fmt.Printf("result %s\n", result)
  for l, r := range deps {
    if Before(result, l, r) {
      result = strings.Replace(result, l, "*", -1)
      result = strings.Replace(result, r, "+", -1)
      result = strings.Replace(result, "*", r, -1)
      result = strings.Replace(result, "+", l, -1)
    }
    fmt.Printf("result after %s => %s is %s\n", l, r, result)
  }
  return result
}

func Before(s, a, b string) bool {
  return strings.Index(s, a) < strings.Index(s, b)
}
