package exporter

import (
	"github.com/foomo/contentserver/content"
	"github.com/foomo/neoscontentrepository/model"
)

// --------------------------------------------------------------------------------------------
// ~ Constants
// --------------------------------------------------------------------------------------------

const NodeTypePage NodeType = "Neos.NodeTypes:Page"
const MimeTypePage = "application/neos+page"

type Page struct{}

// --------------------------------------------------------------------------------------------
// ~ Public methods
// --------------------------------------------------------------------------------------------

func (p *Page) GetRepoNode(node *model.NodeData) (repoNode *content.RepoNode, export bool, err error) {

	repoNode, properties, errRepoNode := createRepoNode(node)
	if errRepoNode != nil {
		err = errRepoNode
		return
	}

	// is not visible if node is marked as "hidden" or "hiddenBeforeDateTime" and "hiddenAfterDateTime" did not match the current time
	// if(!$node->isVisible()) {
	// 	return false;
	// }
	// if(!parent::map($node)) {
	// 	return false;
	// }

	// if ('' != $prop = $node->getProperty('layout')) {
	// 	$layout = $prop;
	// } else if ('' != $prop = RepoNode::getParentProperty($node, 'subpageLayout')) {
	// 	$layout = $prop;
	// } else {
	// 	$layout = 'default';
	// }
	// $this->addData('layout', $layout);

	// add data: youtubeID
	youtubeID := GetPropertyString("youtubeId", properties)
	if youtubeID != "" {
		repoNode.Data["youtubeId"] = youtubeID
	}

	export = true
	return
}
