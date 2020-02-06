package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/foomo/neoscontentrepository/exporter"
	"github.com/foomo/neoscontentrepository/store"
)

func main() {

	flagDSN := flag.String("dsn", "user:pwd@tcp(127.0.0.1:3306)/db", "mysql DSN")
	flagMimeTypes := flag.String("mimes", "Foomo.Neos.Shop:ShopCategory,Foomo.Neos.Shop:ShopRootCategory,Foomo.Neos.Site:App", "comma separated list of node types to search for")
	flag.Parse()

	dsn := *flagDSN
	nodeTypes := strings.Split(*flagMimeTypes, ",")

	start := time.Now()
	nodes, tree, rootPath, errNodes := store.Load(dsn, nodeTypes)
	since := time.Since(start)
	if errNodes != nil {
		log.Fatalln("failed to load nodes:", errNodes)
	}

	exporterPage := &exporter.Page{}
	for _, nodeType := range nodeTypes {
		exporter.RegisterExporter(exporter.NodeType(nodeType), exporterPage)
	}

	convertDuration := time.Now()

	rootNodes, rootNodeOk := tree[rootPath]
	if !rootNodeOk || len(rootNodes) <= 0 {
		log.Fatalln("no nodes for path found", rootPath, rootNodes)
	}

	for _, rootNode := range rootNodes {
		repoNode, errRepoNode := exporter.BuildTree(rootNode, tree)
		if errRepoNode != nil {
			log.Fatalln(errRepoNode)
		}
		fmt.Println("repo node", repoNode)
	}

	fmt.Println("time to load:", since)
	fmt.Println("duration to convert", time.Since(convertDuration))
	fmt.Println("nodes loaded:", len(nodes))
	fmt.Println("root path:", rootPath)
}
