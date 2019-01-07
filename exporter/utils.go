package exporter

import (
	"encoding/json"
	"time"

	"github.com/foomo/contentserver/content"
	"github.com/foomo/neoscontentrepository/model"
)

const RepoNodeDefaultLayout = "default"

func createRepoNode(node *model.NodeData) (repoNode *content.RepoNode, properties map[string]interface{}, err error) {

	var errProperties error
	properties, errProperties = parseProperties(node.Properties)

	repoNode = content.NewRepoNode()
	now := time.Now()

	// set base data
	repoNode.ID = node.Identifier
	repoNode.MimeType = MimeTypePage
	repoNode.Data["createdAt"] = node.CreationDatetime.Unix()
	repoNode.Data["updatedAt"] = node.LastModificationDatetime.Unix()

	repoNode.Hidden = node.HiddenInIndex || node.Hidden ||
		(node.HiddenBeforeDatetime != nil && node.HiddenBeforeDatetime.Unix() > 0 && now.Unix() < node.HiddenBeforeDatetime.Unix()) ||
		(node.HiddenAfterDatetime != nil && node.HiddenAfterDatetime.Unix() > 0 && now.Unix() > node.HiddenAfterDatetime.Unix())

	// add localized properties
	// $this->URI = $this->getNodeUri($node);

	// handle property errors
	if errProperties != nil {
		err = errProperties
		return
	}

	// node name
	repoNode.Name = GetPropertyString("title", properties)

	// layout
	repoNode.Data["layout"] = RepoNodeDefaultLayout
	layout := GetPropertyString("layout", properties)
	if layout != "" {
		repoNode.Data["layout"] = layout
	}

	return
}

// GetPropertyString will return properties value for given key or an empty string
func GetPropertyString(key string, properties map[string]interface{}) string {
	if value, ok := properties[key]; ok {
		return value.(string)
	}
	return ""
}

func parseProperties(props string) (properties map[string]interface{}, err error) {
	errUnmarshall := json.Unmarshal([]byte(props), &properties)
	if errUnmarshall != nil {
		err = errUnmarshall
		return
	}
	return
}
