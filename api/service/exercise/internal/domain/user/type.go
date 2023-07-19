package user

type (
	Type struct {
		ID          TypeID
		Name        TypeName
		Description TypeDescription
	}

	TypeID          int
	TypeName        string
	TypeDescription string
)

// TODO: sampleから実際の値に変更する
const (
	TypeSampleID TypeID = iota + 1

	TypeSampleName        TypeName        = "sample_type"
	TypeSampleDescription TypeDescription = "sample_description"
)

var id2Type = map[TypeID]Type{
	TypeSampleID: {
		ID:          TypeSampleID,
		Name:        TypeSampleName,
		Description: TypeSampleDescription,
	},
}

func NewType(id TypeID) Type {
	return id2Type[id]
}
