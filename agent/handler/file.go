package handler

import (
	"Stowaway/agent/manager"
	"Stowaway/protocol"
	"Stowaway/share"
)

func DispatchFileMess(mgr *manager.Manager) {
	for {
		message := <-mgr.FileManager.FileMessChan

		switch mess := message.(type) {
		case *protocol.FileStatReq:
			mgr.FileManager.File.FileName = mess.Filename
			mgr.FileManager.File.SliceNum = mess.SliceNum
			err := mgr.FileManager.File.CheckFileStat(protocol.TEMP_ROUTE, protocol.ADMIN_UUID, share.AGENT)
			if err == nil {
				go mgr.FileManager.File.Receive(protocol.TEMP_ROUTE, protocol.ADMIN_UUID, share.AGENT)
			}
		case *protocol.FileStatRes:
			if mess.OK == 1 {
				go mgr.FileManager.File.Upload(protocol.TEMP_ROUTE, protocol.ADMIN_UUID, share.AGENT)
			} else {
				mgr.FileManager.File.Handler.Close()
			}
		case *protocol.FileDownReq:
			mgr.FileManager.File.FilePath = mess.FilePath
			mgr.FileManager.File.FileName = mess.Filename
			go mgr.FileManager.File.SendFileStat(protocol.TEMP_ROUTE, protocol.ADMIN_UUID, share.AGENT)
		case *protocol.FileData:
			mgr.FileManager.File.DataChan <- mess.Data
		case *protocol.FileErr:
			mgr.FileManager.File.ErrChan <- true
		case *protocol.DirStatReq:
			mgr.FileManager.File.FilePath = mess.DirName
			go func() {
				err := mgr.FileManager.File.SendDirStat(protocol.TEMP_ROUTE, protocol.ADMIN_UUID)
				if err != nil {
					mgr.FileManager.File.ErrChan <- true
				}
			}()
		}
	}
}
