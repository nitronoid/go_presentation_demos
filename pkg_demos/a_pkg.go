package pkg_demos

func Exported() string {
	return "This is an Exported function!"
}

func Exported2() string {
	return "This is an Exported function calling " + nonExported()
}

func nonExported() string {
	return "a non-Exported function!"
}

type ExportedType struct {ExportedMember, nonExportedMember int}

type ExportedType2 struct {X nonExportedType}

type nonExportedType struct {x int}