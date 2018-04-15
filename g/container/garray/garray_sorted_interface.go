// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package garray

//import "sync"
//
//type IntArray struct {
//    mu           sync.RWMutex           // 互斥锁
//    array        []int                  // 底层数组
//    lessThanFunc func(v1, v2 int) bool  // 比较函数，如何判断v1比v2小
//}
//
//func NewIntArray(size int, cap ... int) *IntArray {
//    a := &IntArray{}
//    if len(cap) > 0 {
//        a.array = make([]int, size, cap[0])
//    } else {
//        a.array = make([]int, size)
//    }
//    a.lessThanFunc = func(v1, v2 int) bool {
//        return v1 < v2
//    }
//    return a
//}
//
//// 获取指定索引的数据项, 调用方注意判断数组边界
//func (a *IntArray) Get(index int) int {
//    a.mu.RLock()
//    value := a.array[index]
//    a.mu.RUnlock()
//    return value
//}
//
//// 设置指定索引的数据项, 调用方注意判断数组边界
//func (a *IntArray) Set(index int, value int) {
//    a.mu.Lock()
//    a.array[index] = value
//    a.mu.Unlock()
//}
//
//// 删除指定索引的数据项, 调用方注意判断数组边界
//func (a *IntArray) Remove(index int) {
//    a.mu.Lock()
//    a.array = append(a.array[ : index], a.array[index + 1 : ]...)
//    a.mu.RUnlock()
//}
//
//// 追加数据项
//func (a *IntArray) Append(value int) {
//    a.mu.Lock()
//    a.array = append(a.array, value)
//    a.mu.Unlock()
//}
//
//// 数组长度
//func (a *IntArray) Len() int {
//    a.mu.RLock()
//    length := len(a.array)
//    a.mu.RUnlock()
//    return length
//}