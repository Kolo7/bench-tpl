/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/errgroup"
)

var (
	httpRequest struct {
		// 并发数量
		concurrency int
		// 单协程内请求间隔
		interval int
	}
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "发起网络请求",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, err := os.Open(input.dir)
		if err != nil {
			logx.Errorf("open file error: %v", err)
			return
		}
		defer inputFile.Close()
		dispatcher := NewWorkerDispatcher(httpRequest.concurrency, httpRequest.interval)
		dispatcher.Run()

		// 逐行读取文件内容
		fileReader := bufio.NewReader(inputFile)
		for {
			line, err := fileReader.ReadString('\n')
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				logx.Errorf("read file error: %v", err)
				break
			}
			// 发送请求
			dispatcher.Dispatch(line)
		}
	},
}

func execCommand(line string) error {
	cmd := exec.Command("/bin/bash", "-c", line)
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Join(err, fmt.Errorf("exec command error: %s", line))
	} else {
		fmt.Println(string(out))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&input.dir, "input", "i", "", "输入文件")
	httpCmd.Flags().IntVarP(&httpRequest.concurrency, "concurrency", "c", 1, "并发数量")
	httpCmd.Flags().IntVarP(&httpRequest.interval, "interval", "t", 1000, "单协程内请求间隔")
}

type WorkerDispatcher struct {
	concurrency int
	interval    int

	input     chan string
	workerCtx context.Context
	worker    []*Worker
	errg      *errgroup.Group
}

func NewWorkerDispatcher(concurrency, interval int) *WorkerDispatcher {
	ctx := context.Background()
	dispatcher := &WorkerDispatcher{
		concurrency: concurrency,
		interval:    interval,
		input:       make(chan string, concurrency*2),
		workerCtx:   ctx,
	}
	dispatcher.errg, ctx = errgroup.WithContext(dispatcher.workerCtx)
	dispatcher.workerCtx = ctx
	dispatcher.errg.SetLimit(concurrency)
	for i := 0; i < concurrency; i++ {
		w := NewWorker(i, interval, dispatcher.input)
		dispatcher.worker = append(dispatcher.worker, w)
	}
	return dispatcher
}

func (d *WorkerDispatcher) Dispatch(line string) {
	d.input <- line
}

func (d *WorkerDispatcher) Run() {
	for _, w := range d.worker {
		d.errg.Go(func() error {
			logx.Infof("worker %d started", w.id)
			w.run(d.workerCtx)
			return nil
		})
	}
	go func() {
		defer d.Stop()
		if err := d.errg.Wait(); err != nil {
			logx.Errorf("run dispatcher error: %v", err)
		}
	}()
}

func (d *WorkerDispatcher) Stop() {
	d.workerCtx.Done()
	close(d.input)
}

type Worker struct {
	id       int
	interval int
	input    <-chan string
}

func NewWorker(id int, interval int, input <-chan string) *Worker {
	return &Worker{
		id:       id,
		interval: interval,
		input:    input,
	}
}

func (w *Worker) run(ctx context.Context) error {
	for {
		ticker := time.NewTicker(time.Duration(w.interval) * time.Millisecond)
		defer ticker.Stop()
		select {
		case <-ctx.Done():
			logx.Infof("worker %d stopped, ctx done", w.id)
			return nil
		case <-ticker.C:
			line, ok := <-w.input
			if !ok {
				logx.Infof("worker %d stopped, input channel closed", w.id)
				return nil
			}
			// 执行命令
			if err := execCommand(line); err != nil {
				logx.Errorf("worker %d exec command error: %v", w.id, err)
				return err
			}
		}
	}
}
