package main

import (
  "testing"
  // "strings"
  )

// func TestEmptyString(t *testing.T) {
//   order := Sort([]string{""})
//   if order != "" {
//     t.Error("Expected empty, got", order)
//   }
// }
//
// func TestSingleJob(t *testing.T) {
//   order := Sort([]string{"a =>"})
//   if order != "a" {
//     t.Error("Expected a, got", order)
//   }
// }
//
// func TestMultipleJobs(t *testing.T) {
//   order := Sort([]string{"a =>", "b =>", "c =>"})
//   if !(strings.Contains(order, "a") &&
//        strings.Contains(order, "b") &&
//        strings.Contains(order, "c") &&
//        len(order) == 3) {
//     t.Error("Expected abc, got", order)
//   }
// }
//
// func TestMultipleJobsSingleDependency(t *testing.T) {
//   order := Sort([]string{"a =>", "b => c", "c =>"})
//   if order != "abc" {
//     t.Error("Expected abc, got", order)
//   }
// }

func TestMultipleJobsMultipleDependencies(t *testing.T) {
  order := Sort([]string{"a =>", "b => c", "c => f", "d => a", "e => b", "f =>"})
  if !(Before(order, "f", "c") &&
       Before(order, "c", "b") &&
       Before(order, "b", "e") &&
       Before(order, "a", "d") &&
       len(order) == 6) {
        t.Error("Order wrong", order)
  }
}
