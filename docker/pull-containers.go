/*
 * functions that pull containers for the following languages
 * c/cpp
 * python
 * javascript
 */

package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullAllContainers(ctx context.Context, client *client.Client) {
	pullNodeContainer(ctx, client)
}

func pullGccContainer(ctx context.Context, client *client.Client) {
	out, err := client.ImagePull(ctx, "gcc:14", image.PullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

func pullNodeContainer(ctx context.Context, client *client.Client) {
	out, err := client.ImagePull(ctx, "node:20-alpine", image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(os.Stdout, out)
}
