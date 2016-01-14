package xmlcomm

import (
	"encoding/xml"

	"github.com/ChrisTrenkamp/goxpath/parser/pathexpr"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/xconst"
)

//XMLComm is an implementation of XPRes for XML attributes
type XMLComm struct {
	xml.Comment
	Parent tree.XPResEle
}

//GetParent returns the parent node
func (c *XMLComm) GetParent() tree.XPResEle {
	return c.Parent
}

//String returns the value of the comment
func (c *XMLComm) String() string {
	return string(c.Comment)
}

//XMLPrint prints the comment in XML form
func (c *XMLComm) XMLPrint(e *xml.Encoder) error {
	return e.EncodeToken(c.Comment)
}

//EvalPath evaluates the XPath path instruction on the element
func (c *XMLComm) EvalPath(p *pathexpr.PathExpr) bool {
	if p.NodeType == xconst.NodeTypeComment || p.NodeType == xconst.NodeTypeNode {
		return true
	}

	return false
}