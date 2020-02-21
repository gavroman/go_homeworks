package main

import (
    "fmt"
    "sort"
    "strings"
    "sync"
)

func doJob(waiter *sync.WaitGroup, jobToDo job, in, out chan interface{}) {
    jobToDo(in, out)
    waiter.Done()
    close(out)
}

func ExecutePipeline(jobs ...job) {
    wg := &sync.WaitGroup{}
    in := make(chan interface{}, MaxInputDataLen)
    for i := range jobs {
        wg.Add(1)
        out := make(chan interface{}, MaxInputDataLen)
        go doJob(wg, jobs[i], in, out)
        in = out
    }
    wg.Wait()
}

func SingleHash(in, out chan interface{}) {
    mutexMD5signer := &sync.Mutex{}
    mainWG := &sync.WaitGroup{}
    for data := range in {
        number := fmt.Sprintf("%d", data.(int))
        str := fmt.Sprintf("%s", number)
        mainWG.Add(1)
        go func() {
            outHash1 := make(chan string)
            go func(output chan<- string) {
                output <- DataSignerCrc32(str)
            }(outHash1)

            outHash2 := make(chan string)
            go func(output chan<- string) {
                mutexMD5signer.Lock()
                md5 := DataSignerMd5(str)
                mutexMD5signer.Unlock()
                output <- DataSignerCrc32(md5)
            }(outHash2)

            hash1 := <-outHash1
            hash2 := <-outHash2
            close(outHash1)
            close(outHash2)

            out <- hash1 + "~" + hash2
            mainWG.Done()
        }()
    }
    mainWG.Wait()
}

func MultiHash(in, out chan interface{}) {
    mainWG := &sync.WaitGroup{}
    for data := range in {
        str := fmt.Sprintf("%s", data.(string))
        mainWG.Add(1)
        go func() {
            stringSlice := make([]string, 6)
            localWG := &sync.WaitGroup{}
            for th := 0; th != 6; th++ {
                localWG.Add(1)
                go func(resultArray *[]string, index int) {
                    (*resultArray)[index] = DataSignerCrc32(fmt.Sprintf("%d", index) + str)
                    localWG.Done()
                }(&stringSlice, th)
            }
            localWG.Wait()

            result := strings.Join(stringSlice, "")
            out <- result
            mainWG.Done()
        }()
    }
    mainWG.Wait()
}

func CombineResults(in, out chan interface{}) {
    stringSlice := make([]string, 0, len(in))
    for data := range in {
        stringSlice = append(stringSlice, data.(string))
    }

    sort.Slice(stringSlice, func(i, j int) bool {
        return stringSlice[i] < stringSlice[j]
    })

    out <- strings.Join(stringSlice, "_")
}
