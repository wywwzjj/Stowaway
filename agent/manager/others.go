package manager

import "Stowaway/share"

type fileManager struct {
	File *share.MyFile

	FileMessChan chan interface{}
}

func newFileManager(file *share.MyFile) *fileManager {
	manager := new(fileManager)
	manager.File = file
	manager.FileMessChan = make(chan interface{}, 5)
	return manager
}

type sshManager struct {
	SSHMessChan chan interface{}
}

func newSSHManager() *sshManager {
	manager := new(sshManager)
	manager.SSHMessChan = make(chan interface{}, 5)
	return manager
}

type sshTunnelManager struct {
	SSHTunnelMessChan chan interface{}
}

func newSSHTunnelManager() *sshTunnelManager {
	manager := new(sshTunnelManager)
	manager.SSHTunnelMessChan = make(chan interface{}, 5)
	return manager
}

type shellManager struct {
	ShellMessChan chan interface{}
}

func newShellManager() *shellManager {
	manager := new(shellManager)
	manager.ShellMessChan = make(chan interface{}, 5)
	return manager
}

type listenManager struct {
	ListenMessChan chan interface{}
	ChildUUIDChan  chan string
}

func newListenManager() *listenManager {
	manager := new(listenManager)
	manager.ListenMessChan = make(chan interface{}, 5)
	manager.ChildUUIDChan = make(chan string)
	return manager
}

type connectManager struct {
	ConnectMessChan chan interface{}
}

func newConnectManager() *connectManager {
	manager := new(connectManager)
	manager.ConnectMessChan = make(chan interface{}, 5)
	return manager
}

type offlineManager struct {
	OfflineMessChan chan interface{}
}

func newOfflineManager() *offlineManager {
	manager := new(offlineManager)
	manager.OfflineMessChan = make(chan interface{}, 5)
	return manager
}
