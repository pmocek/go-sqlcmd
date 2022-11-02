// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package docker

import (
	"context"
	"github.com/docker/distribution/reference"
	"github.com/docker/distribution/registry/client"
	"net/http"
)

func ListTags(path string, baseURL string) []string {
	ctx := context.Background()

	repo, err := reference.WithName(path)
	checkErr(err)
	r, err := client.NewRepository(
		repo,
		baseURL,
		http.DefaultTransport,
	)
	checkErr(err)
	ts := r.Tags(ctx)
	tags, err := ts.All(ctx)
	checkErr(err)

	return tags
}
