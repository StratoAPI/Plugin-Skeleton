package main

import (
	"fmt"

	"github.com/ResourceAPI/Plugin-Skeleton/skeleton"

	"github.com/ResourceAPI/Core/plugins"
)

type SkeletonPlugin string

func (SkeletonPlugin) Name() string {
	return "Skeleton"
}

func (SkeletonPlugin) Entrypoint() {
	plugins.RegisterFacade("skeleton", &skeleton.SkeletonFacade{})
	plugins.RegisterStorage("skeleton", &skeleton.SkeletonStorage{})
	plugins.RegisterFilter("skeleton", &skeleton.SkeletonFilter{})
	plugins.AssociateFilter("skeleton", "skeleton")
	fmt.Println("Skeleton Plugin Initialized")
}

var CorePlugin SkeletonPlugin
