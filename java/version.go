package java

type Version int

const (
	JDK8 Version = iota
	JDK9
	JDK10
	JDK11
	JDK12
)

type VersionManager struct {
}
