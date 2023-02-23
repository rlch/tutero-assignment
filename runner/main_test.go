package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
)

var (
	caseSimple string = "simple"
)

type logConsumer []byte

func (l *logConsumer) Accept(log testcontainers.Log) {
	fmt.Println(log)
	if log.LogType == testcontainers.StdoutLog {
		*l = append(*l, log.Content...)
	}
}

func runBinary(t *testing.T, ctx context.Context, caseName string) string {
	request := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       *DockerContext,
			Dockerfile:    *Dockerfile,
			PrintBuildLog: true,
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      fmt.Sprintf("cases/%s.txt", caseName),
				ContainerFilePath: "/usr/src/app/input.txt",
				FileMode:          0444,
			},
		},
		// WaitingFor: wait.ForExit().WithExitTimeout(time.Minute * 2),
	}
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		t.Fatalf("unable to start container: %s", err)
	}
	defer func() {
		if err = container.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	c := logConsumer{}
	if err = container.StartLogProducer(ctx); err != nil {
		t.Fatal(err)
	}
	container.FollowOutput(&c)
	for {
		time.Sleep(time.Millisecond * 20)
		if running := container.IsRunning(); running {
			break
		}
	}
	defer func() {
		if err = container.StopLogProducer(); err != nil {
			t.Fatal(err)
		}
	}()
	return string(c)
}

func TestRunner(t *testing.T) {
	ctx := context.Background()
	t.Parallel()
	t.Log(*DockerContext, *Dockerfile)

	testCases := []struct {
		name     string
		caseFile string
	}{
		{
			name:     "...",
			caseFile: caseSimple,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			output := runBinary(t, ctx, tc.caseFile)
			t.Error(output)
			// TODO:
		})
	}
}

func BenchmarkRunner(b *testing.B) {
	// TODO:
}
