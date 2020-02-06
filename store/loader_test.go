package store

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	dsn := "user:pwd@tcp(127.0.0.1:3306)/db"

	nodeTypes := []string{
		"Neos.NodeTypes:Page",
		"Neos.Neos:Shortcut",
		"Foomo.Neos.Shop:ShopCategory",
		"Foomo.Neos.Shop:ShopRootCategory",
		"Foomo.Neos.Site:App",
		"Foomo.Neos.Site:External",
		"Foomo.Neos.Site:Folder",
	}

	start := time.Now()
	nodes, _, _, errNodes := Load(dsn, nodeTypes)
	fmt.Println("time to load:", time.Since(start))

	assert.NoError(t, errNodes)
	assert.NotEqual(t, 0, len(nodes))

	for _, node := range nodes {
		fmt.Println("node: ", node.Identifier, *node.Workspace, node.DimensionsHash, node.Version, node.Removed)
	}

	fmt.Println("nodes loaded:", len(nodes))
}
