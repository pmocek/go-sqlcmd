// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package container

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
	repository, err := client.NewRepository(
		repo,
		baseURL,
		http.DefaultTransport,
	)
	checkErr(err)
	tagService := repository.Tags(ctx)
	tags, err := tagService.All(ctx)
	checkErr(err)

	return tags
}
