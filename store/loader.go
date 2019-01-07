package store

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/foomo/neoscontentrepository/model"

	// mysql driver import
	_ "github.com/go-sql-driver/mysql"
)

func Load(dsn string, nodeTypes []string) (nodes map[string]*model.NodeData, tree map[string][]*model.NodeData, rootPath string, err error) {

	nodesLock := &sync.RWMutex{}
	nodes = map[string]*model.NodeData{}

	// connect with DB
	db, errDB := sql.Open("mysql", dsn+"?parseTime=true")
	if errDB != nil {
		err = errDB
		return
	}
	defer db.Close()

	nodeTypesSlice := []string{}
	for _, nodeType := range nodeTypes {
		nodeTypesSlice = append(nodeTypesSlice, `"`+nodeType+`"`)
	}
	nodeTypesString := strings.Join(nodeTypesSlice, ",")

	// sqlStatement, errStatement := db.Prepare(`SELECT persistence_object_identifier, identifier, version, workspace, contentobjectproxy, movedto, pathhash, path, parentpathhash, parentpath, sortingindex, removed, dimensionshash, creationdatetime, lastmodificationdatetime, lastpublicationdatetime, hiddenbeforedatetime, hiddenafterdatetime, dimensionvalues, properties, nodetype, hidden, hiddeninindex, accessroles FROM neos_contentrepository_domain_model_nodedata WHERE nodetype IN (?` + strings.Repeat(`,?`, len(nodeTypes)-1) + `)`)
	sqlStatement, errStatement := db.Prepare(fmt.Sprintf(`SELECT persistence_object_identifier, identifier, version, workspace, contentobjectproxy, movedto, pathhash, path, parentpathhash, parentpath, sortingindex, removed, dimensionshash, creationdatetime, lastmodificationdatetime, lastpublicationdatetime, hiddenbeforedatetime, hiddenafterdatetime, dimensionvalues, properties, nodetype, hidden, hiddeninindex, accessroles FROM neos_contentrepository_domain_model_nodedata WHERE nodetype IN (%s)`, nodeTypesString))
	if errStatement != nil {
		err = errStatement
		return
	}

	// execute sql query
	rows, errRows := sqlStatement.Query()
	if errRows != nil {
		err = errRows
		return
	}

	idsLock := &sync.RWMutex{}
	ids := map[string]bool{}
	tree = map[string][]*model.NodeData{}

	wg := &sync.WaitGroup{}
	for rows.Next() {
		wg.Add(1)

		nodeData := &model.NodeData{}
		// for each row, scan the result into a node data object
		errScan := rows.Scan(
			&nodeData.PersistenceObjectIdentifier,
			&nodeData.Identifier,
			&nodeData.Version,
			&nodeData.Workspace,
			&nodeData.ContentObjectProxy,
			&nodeData.MovedTo,
			&nodeData.PathHash,
			&nodeData.Path,
			&nodeData.ParentPathHash,
			&nodeData.ParentPath,
			&nodeData.SortingIndex,
			&nodeData.Removed,
			&nodeData.DimensionsHash,
			&nodeData.CreationDatetime,
			&nodeData.LastModificationDatetime,
			&nodeData.LastPublicationDatetime,
			&nodeData.HiddenBeforeDatetime,
			&nodeData.HiddenAfterDatetime,
			&nodeData.DimensionValues,
			&nodeData.Properties,
			&nodeData.NodeType,
			&nodeData.Hidden,
			&nodeData.HiddenInIndex,
			&nodeData.AccessRoles,
		)
		if errScan != nil {
			err = errScan
			return
		}

		// find shortest path
		if rootPath == "" || len(rootPath) > len(nodeData.Path) {
			rootPath = nodeData.ParentPath
		}

		go func() {
			defer wg.Done()

			// skip other workspaces
			if *nodeData.Workspace != "live" && *nodeData.Workspace != "stage" {
				return
			}

			// skip removed nodes
			if nodeData.Removed {
				return
			}

			hash := nodeData.GetHash()

			// collect nodes
			nodesLock.Lock()
			nodes[hash] = nodeData
			if _, ok := tree[nodeData.ParentPath]; !ok {
				tree[nodeData.ParentPath] = []*model.NodeData{}
			}
			tree[nodeData.ParentPath] = append(tree[nodeData.ParentPath], nodeData)
			nodesLock.Unlock()

			idsLock.RLock()
			_, ok := ids[hash]
			idsLock.RUnlock()
			if !ok {
				idsLock.Lock()
				ids[hash] = true
				idsLock.Unlock()
			} else {
				fmt.Println("duplicate ID", nodeData.GetHash())
			}
		}()
	}

	wg.Wait()

	return
}
