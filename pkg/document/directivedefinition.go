package document

// DirectiveDefinition as specified in
// http://facebook.github.io/graphql/draft/#DirectiveDefinition
type DirectiveDefinition struct {
	Description         ByteSliceReference
	Name                ByteSliceReference
	ArgumentsDefinition []int
	DirectiveLocations  DirectiveLocations
}

func (d DirectiveDefinition) NodeValueType() ValueType {
	panic("implement me")
}

func (d DirectiveDefinition) NodeValueReference() int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeName() ByteSliceReference {
	return d.Name
}

func (d DirectiveDefinition) NodeAlias() ByteSliceReference {
	panic("implement me")
}

func (d DirectiveDefinition) NodeDescription() ByteSliceReference {
	return d.Description
}

func (d DirectiveDefinition) NodeArguments() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeArgumentsDefinition() []int {
	return d.ArgumentsDefinition
}

func (d DirectiveDefinition) NodeDirectives() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeEnumValuesDefinition() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeFields() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeFieldsDefinition() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeFragmentSpreads() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeInlineFragments() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeVariableDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeType() int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeOperationType() OperationType {
	panic("implement me")
}

func (d DirectiveDefinition) NodeValue() int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeDefaultValue() int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeImplementsInterfaces() []ByteSliceReference {
	panic("implement me")
}

func (d DirectiveDefinition) NodeSchemaDefinition() SchemaDefinition {
	panic("implement me")
}

func (d DirectiveDefinition) NodeScalarTypeDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeObjectTypeDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeInterfaceTypeDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeUnionTypeDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeEnumTypeDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeInputObjectTypeDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeDirectiveDefinitions() []int {
	panic("implement me")
}

func (d DirectiveDefinition) NodeUnionMemberTypes() []ByteSliceReference {
	panic("implement me")
}

// ContainsLocation returns if the $location is contained
func (d DirectiveDefinition) ContainsLocation(location DirectiveLocation) bool {
	for _, dirLoc := range d.DirectiveLocations {
		if dirLoc == location {
			return true
		}
	}

	return false
}

// DirectiveDefinitions is the plural of DirectiveDefinition
type DirectiveDefinitions []DirectiveDefinition
