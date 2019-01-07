package exporter

import (
	"errors"

	"github.com/foomo/contentserver/content"
	"github.com/foomo/neoscontentrepository/model"
)

type NodeType string

type NodeExporter interface {
	GetRepoNode(node *model.NodeData) (repoNode *content.RepoNode, export bool, err error)
}

var nodeExporterMap = map[string]NodeExporter{}
var ErrorUnsupportedNodeType = errors.New("unsupported node type: no exporter registered for given node type")

func RegisterExporter(nodeType NodeType, exporter NodeExporter) {
	nodeExporterMap[string(nodeType)] = exporter
}

func Export(node *model.NodeData) (repoNode *content.RepoNode, ok bool, err error) {
	nodeExporter, nodeExporterOK := nodeExporterMap[node.NodeType]
	if !nodeExporterOK {
		err = ErrorUnsupportedNodeType
		return
	}
	return nodeExporter.GetRepoNode(node)
}
