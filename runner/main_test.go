package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func containerRequest(caseName string) testcontainers.ContainerRequest {
	return testcontainers.ContainerRequest{
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
		WaitingFor: wait.ForExit().WithExitTimeout(time.Minute * 2),
	}
}

type runOutput struct {
	output []string
	types.ContainerState
}

func runBinary(ctx context.Context, caseName string) (_ *runOutput, err error) {
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest(caseName),
		Started:          true,
	}
	var container testcontainers.Container
	container, err = testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to start container: %s", err)
	}
	containerLogs, err := container.Logs(ctx)
	if err != nil {
		return nil, err
	}
	c, err := io.ReadAll(containerLogs)
	if err != nil {
		return nil, (err)
	}
	defer func() {
		if cErr := container.Terminate(ctx); cErr != nil {
			err = cErr
		}
	}()

	output := []string{}
	rawOutput := strings.TrimSpace(string(c))
	for _, node := range strings.Split(rawOutput, "\n") {
		output = append(output, strings.TrimSpace(node))
	}

	state, err := container.State(ctx)
	if err != nil {
		return nil, err
	}
	return &runOutput{
		output:         output,
		ContainerState: *state,
	}, nil
}

func TestRunner(t *testing.T) {
	ctx := context.Background()
	t.Parallel()

	testCases := []struct {
		name     string
		caseFile string
		output   []string
	}{
		{
			name:     "Only edges",
			caseFile: "no_progress",
			output:   []string{"A", "B", "C", "D"},
		},
		{
			name:     "Only progress",
			caseFile: "no_edges",
			output:   []string{"C", "B", "A"},
		},
		{
			name:     "Some nodes with progress",
			caseFile: "with_progress",
			output:   []string{"A", "E", "B", "C", "D"},
		},
		{
			name:     "All nodes with progress",
			caseFile: "with_progress_two",
			output:   []string{"A", "C", "B", "D", "E", "F"},
		},
		{
			name:     "Multiple leaf nodes having same level",
			caseFile: "multiple_leaf_node_in_same_layer",
			output:   []string{"A", "B", "C", "D"},
		},
		{
			name:     "Multiple roots at same level",
			caseFile: "multiple_root_at_same_level",
			output:   []string{"A", "B", "C", "D"},
		},
		{
			name:     "Multiple intermediate nodes at same level",
			caseFile: "multiple_intermediate_nodes_at_same_level",
			output:   []string{"A", "B", "C", "D", "E"},
		},
		{
			name:     "Intermittent edge and progress",
			caseFile: "intermittent_edge_and_progress",
			output:   []string{"A", "C", "B", "D"},
		},
		{
			name:     "Few edges and few nodes with only progress",
			caseFile: "few_edges_and_progress",
			output:   []string{"A", "E", "F", "B", "C", "D"},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)
			result, err := runBinary(ctx, tc.caseFile)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(tc.output, result.output)
		})
	}
}

func BenchmarkRunner(t *testing.B) {
	t.StopTimer()
	numNodes := t.N
	numEdges := numNodes / 2
	var (
		buf   strings.Builder
		nodes []string

		curNode = "A"
	)
	for i := 0; i < numNodes; i++ {
		nodes = append(nodes, curNode)
		progress := rand.Float64()
		buf.WriteString(
			curNode + "=" + strconv.FormatFloat(progress, 'f', 1, 64) + "\n",
		)
		curNode = incrementID(curNode)
	}
	for i := 0; i < numEdges; i++ {
		ri := rand.Intn(numNodes - 1)
		// ri + (n > 0) enforces acyclicity as adjacency matrix will be lower
		// triangular
		rj := ri + 1 + rand.Intn(numNodes-ri-1)
		ni := nodes[ri]
		nj := nodes[rj]
		buf.WriteString(ni + "->" + nj + "\n")
	}
	err := os.WriteFile("cases/benchmark.txt", []byte(buf.String()), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("cases/benchmark.txt")
	t.StartTimer()
	if _, err := runBinary(context.Background(), "benchmark"); err != nil {
		t.Fatal(err)
	}
}

func incrementID(id string) string {
	for c := len(id) - 1; c >= 0; c-- {
		char := id[c]
		if char == 'Z' {
			continue
		}
		char += 1
		return id[:c] + string(char) + strings.Repeat("A", len(id)-c-1)
	}
	return strings.Repeat("A", len(id)+1)
}
