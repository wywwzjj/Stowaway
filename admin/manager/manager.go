/*
 * @Author: ph4ntom
 * @Date: 2021-03-23 19:01:26
 * @LastEditors: ph4ntom
 * @LastEditTime: 2021-04-02 17:24:21
 */
package manager

import (
	"Stowaway/share"
)

// Manager is used to maintain all status and keep different parts connected
// Really complicated when i trying to keep Stowaway running stable
// So there may be some bugs in manager
// Plz let me know if stowaway panic when you are using forward/stopforward/backward/stopbackward/socks/stopsocks or node offline
type Manager struct {
	ConsoleManager   *consoleManager
	FileManager      *fileManager
	SocksManager     *socksManager
	ForwardManager   *forwardManager
	BackwardManager  *backwardManager
	SSHManager       *sshManager
	SSHTunnelManager *sshTunnelManager
	ShellManager     *shellManager
	InfoManager      *infoManager
	ListenManager    *listenManager
	ConnectManager   *connectManager
	ChildrenManager  *childrenManager
}

func NewManager(file *share.MyFile) *Manager {
	manager := new(Manager)
	manager.ConsoleManager = newConsoleManager()
	manager.FileManager = newFileManager(file)
	manager.SocksManager = newSocksManager()
	manager.ForwardManager = newForwardManager()
	manager.BackwardManager = newBackwardManager()
	manager.SSHManager = newSSHManager()
	manager.SSHTunnelManager = newSSHTunnelManager()
	manager.ShellManager = newShellManager()
	manager.InfoManager = newInfoManager()
	manager.ListenManager = newListenManager()
	manager.ConnectManager = newConnectManager()
	manager.ChildrenManager = newchildrenManager()
	return manager
}

func (manager *Manager) Run() {
	go manager.SocksManager.run()
	go manager.ForwardManager.run()
	go manager.BackwardManager.run()
}
