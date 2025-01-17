package hcl

import (
	"github.com/hashicorp/hcl2/hcl"
)

type hclNameDescription struct {
	Name    string         `hcl:"name,label"`
	Default hcl.Attributes `hcl:"default,remain"`
}
type hclVariable hclNameDescription
type hclOutput hclNameDescription
type hclProvider hclNameDescription
type hclModule hclNameDescription

type hclNameType struct {
	Type   string   `hcl:"type,label"`
	Name   string   `hcl:"name,label"`
	Config hcl.Body `hcl:",remain"`
}
type hclResource hclNameType
type hclData hclNameType

type configOnly struct {
	Config hcl.Attributes `hcl:",remain"`
}
type hclLocals configOnly

type backend struct {
	Type   string   `hcl:"type,label"`
	Config hcl.Body `hcl:",remain"`
}
type hclTerraform struct {
	RequiredVersion *string  `hcl:"required_version,attr"`
	Backend         *backend `hcl:"backend,block"`
}

type hclRoot struct {
	Variables []*hclVariable `hcl:"variable,block"`
	Outputs   []*hclOutput   `hcl:"output,block"`
	Resources []*hclResource `hcl:"resource,block"`
	Locals    []*hclLocals   `hcl:"locals,block"`
	Data      []*hclData     `hcl:"data,block"`
	Providers []*hclProvider `hcl:"provider,block"`
	Terraform *hclTerraform  `hcl:"terraform,block"`
	Modules   []*hclModule   `hcl:"module,block"`
}
