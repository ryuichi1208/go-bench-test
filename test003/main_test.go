package main

import (
	"sync"
	"testing"
)

func BenchmarkMemAllocOndemand(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]string, 0)
		for j := 0; j < n; j++ {
			s = append(s, "alice")
		}
	}
}

func BenchmarkMemAllocAllBeforeUsing(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]string, 0, n)
		for j := 0; j < n; j++ {
			s = append(s, "alice")
		}
	}
}

func BenchmarkFillSliceByAppend(b *testing.B) {
	n := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, n)
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkFillSliceByIndex(b *testing.B) {
	n := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		for j := 0; j < n; j++ {
			s[j] = j
		}
	}
}

func BenchmarkSyncWithChannel(b *testing.B) {
	n := 10
	c := make(chan struct{}, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			go func() {
				// do something.
				c <- struct{}{}
			}()
		}
		for j := 0; j < n; j++ {
			<-c
		}
	}
}

func BenchmarkSyncWithWaitGroup(b *testing.B) {
	n := 10
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(n)
		for j := 0; j < n; j++ {
			go func() {
				// do something.
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
