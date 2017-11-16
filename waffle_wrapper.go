package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	enginePath := "C:\\Users\\mwata\\sdt5\\engines\\honey_waffle\\YaneuraOu-2017-early.exe"
	cmd := exec.Command(enginePath)

	cmdStdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cmdStdin.Close()

	// cmd.Stdin = os.Stdin
	scanner := bufio.NewScanner(os.Stdin)
	cmd.Stdout = os.Stdout
	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	quitCh := make(chan bool, 1)
	go func() {
		cmd.Wait()
		quitCh <- true
	}()

	go func() {
		<-quitCh
		os.Exit(0)
	}()

	isBlack := true
	isNoColor := true

	for scanner.Scan() {
		text := scanner.Text()
		if isNoColor {
			if strings.HasPrefix(text, "position") {
				// 先手なら"position startpos"だけ
				tmpArr := strings.Split(text, " ")
				if len(tmpArr) == 2 {
					isBlack = true
				} else {
					isBlack = false
				}
				isNoColor = false
			}
		}
		// 必要なら持ち時間を改ざんする。対抗形での長期戦を想定。
		if strings.HasPrefix(text, "go") {
			tmpArr := strings.Split(text, " ")
			// "go btime 3600000 wtime 3600000 byoyomi 0"
			// "go btime 900000 wtime 900000 byoyomi 10000"
			// "go btime 318000 wtime 312000 binc 10000 winc 10000" WCSCではこうだったが。
			if tmpArr[1] != "ponder" {
				// ponderの時にもやるのは怖すぎる
				btime, _ := strconv.Atoi(tmpArr[2])
				wtime, _ := strconv.Atoi(tmpArr[4])
				/*
					// byoyomi, _ := strconv.Atoi(tmpArr[6])
					// modByoyomi := 0
					if isBlack {
						if btime < 10000 {
							// 自分の持ち時間が少ない場合はまとめて使う
							modByoyomi = btime + byoyomi
						} else {
							if btime > wtime {
								// 相手より時間がある場合は20秒
								modByoyomi = 20000
							} else {
								// 時間がなければ10秒
								modByoyomi = 10000
							}
						}
					} else {
						if wtime < 10000 {
							// 自分の持ち時間が少ない場合はまとめて使う
							modByoyomi = wtime + byoyomi
						} else {
							if wtime > btime {
								// 相手より時間がある場合は20秒
								modByoyomi = 20000
							} else {
								// 時間がなければ10秒
								modByoyomi = 10000
							}
						}
					}
				*/
				if isBlack {
					if btime > 1800000 {
						if btime > wtime {
							btime = 1800000
						}
					}
				} else {
					if wtime > 1800000 {
						if wtime > btime {
							wtime = 1800000
						}
					}
				}
				// text = "go btime 0 wtime 0 byoyomi " + strconv.Itoa(modByoyomi)
				text = "go btime " + strconv.Itoa(btime) + " wtime " + strconv.Itoa(wtime) + " byoyomi 0"
			}
		}
		io.WriteString(cmdStdin, text+"\n")
	}
}
