//line :1
package main

import (
	"log"
	"net"
	"runtime"

	"Stowaway/agent/initial"
	"Stowaway/agent/process"
	"Stowaway/global"
	"Stowaway/protocol"
	"Stowaway/share"
)

func init() {
	/*line :1*/ runtime.GOMAXPROCS( /*line :1*/ runtime.NumCPU())
}

func main() {
	options := /*line :1*/ initial.ParseOptions()

	agent := /*line :1*/ process.NewAgent(options)

	/*line :1*/ protocol.DecideType(options.Upstream, options.Downstream)

	var conn net.Conn
	switch options.Mode {
	case initial.NORMAL_PASSIVE:
		conn, agent.UUID = /*line :1*/ initial.NormalPassive(options)
	case initial.NORMAL_RECONNECT_ACTIVE:
		fallthrough
	case initial.NORMAL_ACTIVE:
		conn, agent.UUID = /*line :1*/ initial.NormalActive(options, nil)
	case initial.PROXY_RECONNECT_ACTIVE:
		fallthrough
	case initial.PROXY_ACTIVE:
		proxy := /*line :1*/ share.NewProxy(options.Connect, options.Proxy, options.ProxyU, options.ProxyP)
		conn, agent.UUID = /*line :1*/ initial.NormalActive(options, proxy)
	case initial.IPTABLES_REUSE_PASSIVE:
		defer /*line :1*/ initial.DeletePortReuseRules(options.Listen, options.ReusePort)
		conn, agent.UUID = /*line :1*/ initial.IPTableReusePassive(options)
	case initial.SO_REUSE_PASSIVE:
		conn, agent.UUID = /*line :1*/ initial.SoReusePassive(options)
	default:
		/*line :1*/ log.Fatal( /*line :1*/ func() string {
		fullData := /*line :1*/ []byte("\xcfWr\x1cF8\xbf\x14\xe8H\xeby\xe7\xd09\x17\xa4*\x1ayT\x1e\xb3\n\x13\x8c#\xb2(\xbf\x0f/")
		var data []byte
		data = /*line :1*/ append(data, fullData[27]-fullData[1], fullData[14]-fullData[30], fullData[10]+fullData[2], fullData[8]+fullData[5], fullData[7]-fullData[6], fullData[12]-fullData[19], fullData[21]-fullData[22], fullData[20]+fullData[18], fullData[24]-fullData[16], fullData[9]+fullData[31], fullData[15]^fullData[11], fullData[17]-fullData[23], fullData[3]-fullData[0], fullData[29]^fullData[13], fullData[25]-fullData[28], fullData[4]^fullData[26])
		return /*line :1*/ string(data)
	}())
	}

	/*line :1*/ global.InitialGComponent(conn, options.Secret, agent.UUID)

	/*line :1*/ agent.Run()
}
