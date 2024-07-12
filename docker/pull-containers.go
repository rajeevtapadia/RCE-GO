/*
 * utility functions to pull containers for the following languages
 * c/cpp
 * python
 * javascript
 */

package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"rce-go/utils"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullAllContainers() {
	ctx := context.Background()

	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer dockerCli.Close()

	pullContainer(ctx, dockerCli, utils.NodeImage)
}

func pullContainer(ctx context.Context, client *client.Client, name string) {
	out, err := client.ImagePull(ctx, name, image.PullOptions{})
	if err != nil {
		fmt.Println("error while pulling images")
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
}
