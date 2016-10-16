package models

type SingularitySandbox struct {
	SlaveHostname    string                   `json:"slaveHostname"`
	Files            []SingularitySandboxFile `json:"files"`
	CurrentDirectory string                   `json:"currentDirectory"`
	FullPathToRoot   string                   `json:"fullPathToRoot"`
}

type SingularitySandboxFile struct {
	Size  uint64 `json:"size"`
	Mode  string `json:"mode"`
	Mtime uint64 `json:"mtime"`
	Name  string `json:"name"`
}
