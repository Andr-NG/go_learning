package main

import "fmt"



func ReadData() {
	pr := PeerState{
		m: make(map[string]*Result),
	}

	pr.m["123"] = &Result{
		AudioDone: true,
		VideoDone: false,
	}
	pr.m["321"] = &Result{
		AudioDone: false,
		VideoDone: true,
	}
	pr.m["2"] = &Result{
		AudioDone: true,
		VideoDone: true,
	}
	pr.m["77"] = &Result{
		AudioDone: false,
		VideoDone: false,
	}

    pr.Update("77", func(r *Result){
        r.AudioDone = true
        r.VideoDone = true
    })

	totalCount := len(pr.m)
	succesCount := 0
	failedCount := 0
	onlyVideo := 0
	onlyAudio := 0

	fmt.Println(totalCount)
	for key, val := range pr.m {
		if val.AudioDone && val.VideoDone {
			succesCount++
			fmt.Printf("Peer: %s is connected successfully\n", key)
		} else if !val.AudioDone && !val.VideoDone {
			failedCount++
			fmt.Printf("Peer: %s failed to connect\n", key)
		} else {
			if val.AudioDone {
				onlyAudio++
				fmt.Printf("Peer: %s has only audio\n", key)
			} else {
				onlyVideo++
				fmt.Printf("Peer: %s has only video\n", key)
			}
		}
		}
	}

type Result struct {
	AudioDone bool
	VideoDone bool
}

type PeerState struct {
	m map[string]*Result
}

func (pr *PeerState) Update(id string, fn func(*Result)) {
	res, ok := pr.m[id]
    if !ok || res == nil {
        res = &Result{}
        pr.m[id] = res
    }
	fn(res)
}
