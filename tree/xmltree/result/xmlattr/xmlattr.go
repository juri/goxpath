package xmlattr

import (
	"encoding/xml"

	"github.com/ChrisTrenkamp/goxpath/parser/pathexpr"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/xconst"
)

//XMLAttr is an implementation of XPRes for XML attributes
type XMLAttr struct {
	xml.Attr
	Parent tree.XPResEle
}

//GetParent returns the parent node
func (a *XMLAttr) GetParent() tree.XPResEle {
	return a.Parent
}

//String returns the string value of the attribute
func (a *XMLAttr) String() string {
	return a.Attr.Value
}

//XMLPrint prints the attribute as a processing-instruction.
func (a *XMLAttr) XMLPrint(e *xml.Encoder) error {
	str := a.Attr.Name.Local + `="` + a.Attr.Value + `"`

	if a.Attr.Name.Space != "" {
		str += ` xmlns="` + a.Attr.Name.Space + `"`
	}

	pi := xml.ProcInst{
		Target: "attribute",
		Inst:   ([]byte)(str),
	}

	return e.EncodeToken(pi)
}

//EvalPath evaluates the XPath path instruction on the element
func (a *XMLAttr) EvalPath(p *pathexpr.PathExpr) bool {
	if p.NodeType == "" {
		if p.Name.Space != "*" {
			if a.Attr.Name.Space != p.NS[p.Name.Space] {
				return false
			}
		}

		if p.Name.Local == "*" && p.Axis == xconst.AxisAttribute {
			return true
		}

		if p.Name.Local == a.Attr.Name.Local {
			return true
		}
	} else {
		if p.NodeType == xconst.NodeTypeNode {
			return true
		}
	}

	return false
}