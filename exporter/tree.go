package exporter

import (
	"errors"
	"sync"

	"github.com/foomo/contentserver/content"
	"github.com/foomo/neoscontentrepository/model"
)

func BuildTree(node *model.NodeData, tree map[string][]*model.NodeData) (repoNode *content.RepoNode, err error) {
	repoNode, exportRootNodeOk, errRepoNode := Export(node)
	if errRepoNode != nil {
		err = errRepoNode
		return
	}
	if !exportRootNodeOk {
		err = errors.New("root node export not ok")
		return
	}

	// childNodes, childNodesOk := tree[node.Path]
	// if !childNodesOk {
	// 	err = errors.New("no nodes for path found: " + node.Path)
	// 	return
	// }

	wg := &sync.WaitGroup{}
	lock := &sync.RWMutex{}
	for _, childNode := range tree[node.Path] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			childRepoNode, errChildRepoNode := BuildTree(childNode, tree)
			if errChildRepoNode != nil {
				err = errChildRepoNode
				return
			}
			// @todo keep sorting index order
			lock.Lock()
			repoNode.Index = append(repoNode.Index, childRepoNode.ID)
			repoNode.AddNode(childRepoNode.ID, childRepoNode)
			lock.Unlock()
		}()
	}
	wg.Wait()
	return
}
