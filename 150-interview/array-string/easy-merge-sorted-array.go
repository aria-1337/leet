package main

import "slices"

// With sort
func merge(nums1 []int, m int, nums2 []int, n int) {
    for n > 0 {
        nums1[m] = nums2[n-1]
        n--
        m++
    }
    slices.Sort(nums1)
}
