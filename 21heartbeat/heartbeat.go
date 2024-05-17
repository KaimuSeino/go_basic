package heartbeat

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func Heartbeat() { //main goroutine側でtask goroutineが正常に動作しているかを確認する
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.LstdFlags) // MultiWriter(fileの内容と標準出力で表示)
	ctx, cancel := context.WithTimeout(context.Background(), 7100*time.Millisecond)
	defer cancel()
	const wdtTimeout = 800 * time.Millisecond
	const beatInterval = 500 * time.Millisecond
	heartbeat, v := task(ctx, beatInterval)

loop:
	for {
		select {
		case _, ok := <-heartbeat: // heartbeatが500ms毎に書き込まれるからこちらが実行する
			if !ok {
				break loop
			}
			fmt.Println("beat pulse ⚡️")
		case r, ok := <-v:
			if !ok {
				break loop
			}
			t := strings.Split(r.String(), "m=") // "m=の部分でスライスをしてm=の後の値を出力する"
			fmt.Printf("value: %v [s]\n", t[1])
		case <-time.After(wdtTimeout): // time.Afterで設定しているwdtTimeoutの時間が経過した場合(time.Afterはfor文を抜けるとリセットされる)
			errorLogger.Println("doTask groutine's heartbeat stopped")
			break loop
		}
	}
}

func task(
	ctx context.Context,
	beatInterval time.Duration,
) (<-chan struct{}, <-chan time.Time) { // <-chan struct{}: データなしの読み取り専用ch
	heartbeat := make(chan struct{}) // heartbeatは値を持たない通信用のchなのでからの構造体にする
	out := make(chan time.Time)      // Newtickerの時刻の値を出力する
	go func() {
		defer close(heartbeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(2 * beatInterval)
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		sendValue := func(t time.Time) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-pulse.C:
					sendPulse()
				case out <- t:
					return
				}
			}
		}
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			case <-pulse.C:
				if i == 3 {
					time.Sleep(1000 * time.Millisecond)
				}
				sendPulse()
				i++
			case t := <-task.C:
				sendValue(t)
			}
		}
	}()
	return heartbeat, out
}
