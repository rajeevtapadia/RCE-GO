package docker

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func StartContainer(command string, image string) {
	ctx := context.Background()

	// create connection
	fmt.Println("creating docker client")
	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer dockerCli.Close()

	resp, err := dockerCli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   []string{"sh", "-c", command},
		Tty:   false,
	}, nil, nil, nil, "")

	if err != nil {
		panic(err)
	}

	fmt.Println("starting cont")
	if err := dockerCli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := dockerCli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	fmt.Println("getting logs")
	out, err := dockerCli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	if err := dockerCli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
		panic(err)
	}
	fmt.Println("container removed")

}
