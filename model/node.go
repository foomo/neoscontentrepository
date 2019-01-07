package model

import (
	"strings"
	"time"
)

/*
+-------------------------------+---------------+------+-----+---------+-------+
| Field                         | Type          | Null | Key | Default | Extra |
+-------------------------------+---------------+------+-----+---------+-------+
| persistence_object_identifier | varchar(40)   | NO   | PRI | NULL    |       |
| workspace                     | varchar(255)  | YES  | MUL | NULL    |       |
| contentobjectproxy            | varchar(40)   | YES  | MUL | NULL    |       |
| movedto                       | varchar(40)   | YES  | UNI | NULL    |       |
| version                       | int(11)       | NO   |     | 1       |       |
| pathhash                      | varchar(32)   | NO   | MUL | NULL    |       |
| path                          | varchar(4000) | NO   |     | NULL    |       |
| parentpathhash                | varchar(32)   | NO   | MUL | NULL    |       |
| parentpath                    | varchar(4000) | NO   | MUL | NULL    |       |
| identifier                    | varchar(255)  | NO   | MUL | NULL    |       |
| sortingindex                  | int(11)       | YES  |     | NULL    |       |
| removed                       | tinyint(1)    | NO   |     | NULL    |       |
| dimensionshash                | varchar(32)   | NO   |     | NULL    |       |
| lastmodificationdatetime      | datetime      | NO   |     | NULL    |       |
| lastpublicationdatetime       | datetime      | YES  |     | NULL    |       |
| hiddenbeforedatetime          | datetime      | YES  |     | NULL    |       |
| hiddenafterdatetime           | datetime      | YES  |     | NULL    |       |
| dimensionvalues               | longtext      | NO   |     | NULL    |       |
| properties                    | longtext      | NO   |     | NULL    |       |
| nodetype                      | varchar(255)  | NO   | MUL | NULL    |       |
| creationdatetime              | datetime      | NO   |     | NULL    |       |
| hidden                        | tinyint(1)    | NO   |     | NULL    |       |
| hiddeninindex                 | tinyint(1)    | NO   |     | NULL    |       |
| accessroles                   | longtext      | NO   |     | NULL    |       |
+-------------------------------+---------------+------+-----+---------+-------+
*/

type NodeData struct {
	PersistenceObjectIdentifier string     `json:"persistence_object_identifier"`
	Identifier                  string     `json:"identifier"`
	Version                     uint       `json:"version"`
	Workspace                   *string    `json:"workspace"`
	ContentObjectProxy          *string    `json:"contentobjectproxy"`
	MovedTo                     *string    `json:"movedto"`
	PathHash                    string     `json:"pathhash"`
	Path                        string     `json:"path"`
	ParentPathHash              string     `json:"parentpathhash"`
	ParentPath                  string     `json:"parentpath"`
	SortingIndex                *uint      `json:"sortingindex"`
	Removed                     bool       `json:"removed"`
	DimensionsHash              string     `json:"dimensionshash"`
	CreationDatetime            time.Time  `json:"creationdatetime"`
	LastModificationDatetime    time.Time  `json:"lastmodificationdatetime"`
	LastPublicationDatetime     *time.Time `json:"lastpublicationdatetime"`
	HiddenBeforeDatetime        *time.Time `json:"hiddenbeforedatetime"`
	HiddenAfterDatetime         *time.Time `json:"hiddenafterdatetime"`
	DimensionValues             string     `json:"dimensionvalues"`
	Properties                  string     `json:"properties"`
	NodeType                    string     `json:"nodetype"`
	Hidden                      bool       `json:"hidden"`
	HiddenInIndex               bool       `json:"hiddeninindex"`
	AccessRoles                 string     `json:"accessroles"`
}

type NodeProperties struct {
	Title          string `json:"title,omitempty"`
	URIPathSegment string `json:"uriPathSegment"`

	Layout        string `json:"layout,omitempty"`
	SubPageLayout string `json:"subpageLayout,omitempty"`
}

func (node *NodeData) IsVisible() bool {
	// @todo @implementme
	return false
}

func (node *NodeData) GetHash() string {
	return strings.Join([]string{
		node.Identifier,
		*node.Workspace,
		node.DimensionsHash,
	}, "___")
}
