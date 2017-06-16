// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net"
	"sync"

	"github.com/vishvananda/netlink"
)

type NetlinkAdapter struct {
	LinkByNameStub        func(string) (netlink.Link, error)
	linkByNameMutex       sync.RWMutex
	linkByNameArgsForCall []struct {
		arg1 string
	}
	linkByNameReturns struct {
		result1 netlink.Link
		result2 error
	}
	linkByNameReturnsOnCall map[int]struct {
		result1 netlink.Link
		result2 error
	}
	ParseAddrStub        func(string) (*netlink.Addr, error)
	parseAddrMutex       sync.RWMutex
	parseAddrArgsForCall []struct {
		arg1 string
	}
	parseAddrReturns struct {
		result1 *netlink.Addr
		result2 error
	}
	parseAddrReturnsOnCall map[int]struct {
		result1 *netlink.Addr
		result2 error
	}
	AddrAddScopeLinkStub        func(netlink.Link, *netlink.Addr) error
	addrAddScopeLinkMutex       sync.RWMutex
	addrAddScopeLinkArgsForCall []struct {
		arg1 netlink.Link
		arg2 *netlink.Addr
	}
	addrAddScopeLinkReturns struct {
		result1 error
	}
	addrAddScopeLinkReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetHardwareAddrStub        func(netlink.Link, net.HardwareAddr) error
	linkSetHardwareAddrMutex       sync.RWMutex
	linkSetHardwareAddrArgsForCall []struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}
	linkSetHardwareAddrReturns struct {
		result1 error
	}
	linkSetHardwareAddrReturnsOnCall map[int]struct {
		result1 error
	}
	NeighAddPermanentIPv4Stub        func(index int, destIP net.IP, hwAddr net.HardwareAddr) error
	neighAddPermanentIPv4Mutex       sync.RWMutex
	neighAddPermanentIPv4ArgsForCall []struct {
		index  int
		destIP net.IP
		hwAddr net.HardwareAddr
	}
	neighAddPermanentIPv4Returns struct {
		result1 error
	}
	neighAddPermanentIPv4ReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetARPOffStub        func(netlink.Link) error
	linkSetARPOffMutex       sync.RWMutex
	linkSetARPOffArgsForCall []struct {
		arg1 netlink.Link
	}
	linkSetARPOffReturns struct {
		result1 error
	}
	linkSetARPOffReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetNameStub        func(netlink.Link, string) error
	linkSetNameMutex       sync.RWMutex
	linkSetNameArgsForCall []struct {
		arg1 netlink.Link
		arg2 string
	}
	linkSetNameReturns struct {
		result1 error
	}
	linkSetNameReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetUpStub        func(netlink.Link) error
	linkSetUpMutex       sync.RWMutex
	linkSetUpArgsForCall []struct {
		arg1 netlink.Link
	}
	linkSetUpReturns struct {
		result1 error
	}
	linkSetUpReturnsOnCall map[int]struct {
		result1 error
	}
	LinkDelStub        func(netlink.Link) error
	linkDelMutex       sync.RWMutex
	linkDelArgsForCall []struct {
		arg1 netlink.Link
	}
	linkDelReturns struct {
		result1 error
	}
	linkDelReturnsOnCall map[int]struct {
		result1 error
	}
	LinkAddStub        func(netlink.Link) error
	linkAddMutex       sync.RWMutex
	linkAddArgsForCall []struct {
		arg1 netlink.Link
	}
	linkAddReturns struct {
		result1 error
	}
	linkAddReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetNsFdStub        func(netlink.Link, int) error
	linkSetNsFdMutex       sync.RWMutex
	linkSetNsFdArgsForCall []struct {
		arg1 netlink.Link
		arg2 int
	}
	linkSetNsFdReturns struct {
		result1 error
	}
	linkSetNsFdReturnsOnCall map[int]struct {
		result1 error
	}
	RouteAddStub        func(route *netlink.Route) error
	routeAddMutex       sync.RWMutex
	routeAddArgsForCall []struct {
		route *netlink.Route
	}
	routeAddReturns struct {
		result1 error
	}
	routeAddReturnsOnCall map[int]struct {
		result1 error
	}
	QdiscAddStub        func(qdisc netlink.Qdisc) error
	qdiscAddMutex       sync.RWMutex
	qdiscAddArgsForCall []struct {
		qdisc netlink.Qdisc
	}
	qdiscAddReturns struct {
		result1 error
	}
	qdiscAddReturnsOnCall map[int]struct {
		result1 error
	}
	FilterAddStub        func(netlink.Filter) error
	filterAddMutex       sync.RWMutex
	filterAddArgsForCall []struct {
		arg1 netlink.Filter
	}
	filterAddReturns struct {
		result1 error
	}
	filterAddReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetlinkAdapter) LinkByName(arg1 string) (netlink.Link, error) {
	fake.linkByNameMutex.Lock()
	ret, specificReturn := fake.linkByNameReturnsOnCall[len(fake.linkByNameArgsForCall)]
	fake.linkByNameArgsForCall = append(fake.linkByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LinkByName", []interface{}{arg1})
	fake.linkByNameMutex.Unlock()
	if fake.LinkByNameStub != nil {
		return fake.LinkByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.linkByNameReturns.result1, fake.linkByNameReturns.result2
}

func (fake *NetlinkAdapter) LinkByNameCallCount() int {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return len(fake.linkByNameArgsForCall)
}

func (fake *NetlinkAdapter) LinkByNameArgsForCall(i int) string {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return fake.linkByNameArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkByNameReturns(result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	fake.linkByNameReturns = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) LinkByNameReturnsOnCall(i int, result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	if fake.linkByNameReturnsOnCall == nil {
		fake.linkByNameReturnsOnCall = make(map[int]struct {
			result1 netlink.Link
			result2 error
		})
	}
	fake.linkByNameReturnsOnCall[i] = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) ParseAddr(arg1 string) (*netlink.Addr, error) {
	fake.parseAddrMutex.Lock()
	ret, specificReturn := fake.parseAddrReturnsOnCall[len(fake.parseAddrArgsForCall)]
	fake.parseAddrArgsForCall = append(fake.parseAddrArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ParseAddr", []interface{}{arg1})
	fake.parseAddrMutex.Unlock()
	if fake.ParseAddrStub != nil {
		return fake.ParseAddrStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseAddrReturns.result1, fake.parseAddrReturns.result2
}

func (fake *NetlinkAdapter) ParseAddrCallCount() int {
	fake.parseAddrMutex.RLock()
	defer fake.parseAddrMutex.RUnlock()
	return len(fake.parseAddrArgsForCall)
}

func (fake *NetlinkAdapter) ParseAddrArgsForCall(i int) string {
	fake.parseAddrMutex.RLock()
	defer fake.parseAddrMutex.RUnlock()
	return fake.parseAddrArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) ParseAddrReturns(result1 *netlink.Addr, result2 error) {
	fake.ParseAddrStub = nil
	fake.parseAddrReturns = struct {
		result1 *netlink.Addr
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) ParseAddrReturnsOnCall(i int, result1 *netlink.Addr, result2 error) {
	fake.ParseAddrStub = nil
	if fake.parseAddrReturnsOnCall == nil {
		fake.parseAddrReturnsOnCall = make(map[int]struct {
			result1 *netlink.Addr
			result2 error
		})
	}
	fake.parseAddrReturnsOnCall[i] = struct {
		result1 *netlink.Addr
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) AddrAddScopeLink(arg1 netlink.Link, arg2 *netlink.Addr) error {
	fake.addrAddScopeLinkMutex.Lock()
	ret, specificReturn := fake.addrAddScopeLinkReturnsOnCall[len(fake.addrAddScopeLinkArgsForCall)]
	fake.addrAddScopeLinkArgsForCall = append(fake.addrAddScopeLinkArgsForCall, struct {
		arg1 netlink.Link
		arg2 *netlink.Addr
	}{arg1, arg2})
	fake.recordInvocation("AddrAddScopeLink", []interface{}{arg1, arg2})
	fake.addrAddScopeLinkMutex.Unlock()
	if fake.AddrAddScopeLinkStub != nil {
		return fake.AddrAddScopeLinkStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addrAddScopeLinkReturns.result1
}

func (fake *NetlinkAdapter) AddrAddScopeLinkCallCount() int {
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return len(fake.addrAddScopeLinkArgsForCall)
}

func (fake *NetlinkAdapter) AddrAddScopeLinkArgsForCall(i int) (netlink.Link, *netlink.Addr) {
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return fake.addrAddScopeLinkArgsForCall[i].arg1, fake.addrAddScopeLinkArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) AddrAddScopeLinkReturns(result1 error) {
	fake.AddrAddScopeLinkStub = nil
	fake.addrAddScopeLinkReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrAddScopeLinkReturnsOnCall(i int, result1 error) {
	fake.AddrAddScopeLinkStub = nil
	if fake.addrAddScopeLinkReturnsOnCall == nil {
		fake.addrAddScopeLinkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addrAddScopeLinkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddr(arg1 netlink.Link, arg2 net.HardwareAddr) error {
	fake.linkSetHardwareAddrMutex.Lock()
	ret, specificReturn := fake.linkSetHardwareAddrReturnsOnCall[len(fake.linkSetHardwareAddrArgsForCall)]
	fake.linkSetHardwareAddrArgsForCall = append(fake.linkSetHardwareAddrArgsForCall, struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}{arg1, arg2})
	fake.recordInvocation("LinkSetHardwareAddr", []interface{}{arg1, arg2})
	fake.linkSetHardwareAddrMutex.Unlock()
	if fake.LinkSetHardwareAddrStub != nil {
		return fake.LinkSetHardwareAddrStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetHardwareAddrReturns.result1
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrCallCount() int {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return len(fake.linkSetHardwareAddrArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrArgsForCall(i int) (netlink.Link, net.HardwareAddr) {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return fake.linkSetHardwareAddrArgsForCall[i].arg1, fake.linkSetHardwareAddrArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturns(result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	fake.linkSetHardwareAddrReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturnsOnCall(i int, result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	if fake.linkSetHardwareAddrReturnsOnCall == nil {
		fake.linkSetHardwareAddrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetHardwareAddrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) NeighAddPermanentIPv4(index int, destIP net.IP, hwAddr net.HardwareAddr) error {
	fake.neighAddPermanentIPv4Mutex.Lock()
	ret, specificReturn := fake.neighAddPermanentIPv4ReturnsOnCall[len(fake.neighAddPermanentIPv4ArgsForCall)]
	fake.neighAddPermanentIPv4ArgsForCall = append(fake.neighAddPermanentIPv4ArgsForCall, struct {
		index  int
		destIP net.IP
		hwAddr net.HardwareAddr
	}{index, destIP, hwAddr})
	fake.recordInvocation("NeighAddPermanentIPv4", []interface{}{index, destIP, hwAddr})
	fake.neighAddPermanentIPv4Mutex.Unlock()
	if fake.NeighAddPermanentIPv4Stub != nil {
		return fake.NeighAddPermanentIPv4Stub(index, destIP, hwAddr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.neighAddPermanentIPv4Returns.result1
}

func (fake *NetlinkAdapter) NeighAddPermanentIPv4CallCount() int {
	fake.neighAddPermanentIPv4Mutex.RLock()
	defer fake.neighAddPermanentIPv4Mutex.RUnlock()
	return len(fake.neighAddPermanentIPv4ArgsForCall)
}

func (fake *NetlinkAdapter) NeighAddPermanentIPv4ArgsForCall(i int) (int, net.IP, net.HardwareAddr) {
	fake.neighAddPermanentIPv4Mutex.RLock()
	defer fake.neighAddPermanentIPv4Mutex.RUnlock()
	return fake.neighAddPermanentIPv4ArgsForCall[i].index, fake.neighAddPermanentIPv4ArgsForCall[i].destIP, fake.neighAddPermanentIPv4ArgsForCall[i].hwAddr
}

func (fake *NetlinkAdapter) NeighAddPermanentIPv4Returns(result1 error) {
	fake.NeighAddPermanentIPv4Stub = nil
	fake.neighAddPermanentIPv4Returns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) NeighAddPermanentIPv4ReturnsOnCall(i int, result1 error) {
	fake.NeighAddPermanentIPv4Stub = nil
	if fake.neighAddPermanentIPv4ReturnsOnCall == nil {
		fake.neighAddPermanentIPv4ReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.neighAddPermanentIPv4ReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetARPOff(arg1 netlink.Link) error {
	fake.linkSetARPOffMutex.Lock()
	ret, specificReturn := fake.linkSetARPOffReturnsOnCall[len(fake.linkSetARPOffArgsForCall)]
	fake.linkSetARPOffArgsForCall = append(fake.linkSetARPOffArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkSetARPOff", []interface{}{arg1})
	fake.linkSetARPOffMutex.Unlock()
	if fake.LinkSetARPOffStub != nil {
		return fake.LinkSetARPOffStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetARPOffReturns.result1
}

func (fake *NetlinkAdapter) LinkSetARPOffCallCount() int {
	fake.linkSetARPOffMutex.RLock()
	defer fake.linkSetARPOffMutex.RUnlock()
	return len(fake.linkSetARPOffArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetARPOffArgsForCall(i int) netlink.Link {
	fake.linkSetARPOffMutex.RLock()
	defer fake.linkSetARPOffMutex.RUnlock()
	return fake.linkSetARPOffArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkSetARPOffReturns(result1 error) {
	fake.LinkSetARPOffStub = nil
	fake.linkSetARPOffReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetARPOffReturnsOnCall(i int, result1 error) {
	fake.LinkSetARPOffStub = nil
	if fake.linkSetARPOffReturnsOnCall == nil {
		fake.linkSetARPOffReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetARPOffReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetName(arg1 netlink.Link, arg2 string) error {
	fake.linkSetNameMutex.Lock()
	ret, specificReturn := fake.linkSetNameReturnsOnCall[len(fake.linkSetNameArgsForCall)]
	fake.linkSetNameArgsForCall = append(fake.linkSetNameArgsForCall, struct {
		arg1 netlink.Link
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LinkSetName", []interface{}{arg1, arg2})
	fake.linkSetNameMutex.Unlock()
	if fake.LinkSetNameStub != nil {
		return fake.LinkSetNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetNameReturns.result1
}

func (fake *NetlinkAdapter) LinkSetNameCallCount() int {
	fake.linkSetNameMutex.RLock()
	defer fake.linkSetNameMutex.RUnlock()
	return len(fake.linkSetNameArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetNameArgsForCall(i int) (netlink.Link, string) {
	fake.linkSetNameMutex.RLock()
	defer fake.linkSetNameMutex.RUnlock()
	return fake.linkSetNameArgsForCall[i].arg1, fake.linkSetNameArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) LinkSetNameReturns(result1 error) {
	fake.LinkSetNameStub = nil
	fake.linkSetNameReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetNameReturnsOnCall(i int, result1 error) {
	fake.LinkSetNameStub = nil
	if fake.linkSetNameReturnsOnCall == nil {
		fake.linkSetNameReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetNameReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetUp(arg1 netlink.Link) error {
	fake.linkSetUpMutex.Lock()
	ret, specificReturn := fake.linkSetUpReturnsOnCall[len(fake.linkSetUpArgsForCall)]
	fake.linkSetUpArgsForCall = append(fake.linkSetUpArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkSetUp", []interface{}{arg1})
	fake.linkSetUpMutex.Unlock()
	if fake.LinkSetUpStub != nil {
		return fake.LinkSetUpStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetUpReturns.result1
}

func (fake *NetlinkAdapter) LinkSetUpCallCount() int {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return len(fake.linkSetUpArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetUpArgsForCall(i int) netlink.Link {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return fake.linkSetUpArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkSetUpReturns(result1 error) {
	fake.LinkSetUpStub = nil
	fake.linkSetUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetUpReturnsOnCall(i int, result1 error) {
	fake.LinkSetUpStub = nil
	if fake.linkSetUpReturnsOnCall == nil {
		fake.linkSetUpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetUpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkDel(arg1 netlink.Link) error {
	fake.linkDelMutex.Lock()
	ret, specificReturn := fake.linkDelReturnsOnCall[len(fake.linkDelArgsForCall)]
	fake.linkDelArgsForCall = append(fake.linkDelArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkDel", []interface{}{arg1})
	fake.linkDelMutex.Unlock()
	if fake.LinkDelStub != nil {
		return fake.LinkDelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkDelReturns.result1
}

func (fake *NetlinkAdapter) LinkDelCallCount() int {
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return len(fake.linkDelArgsForCall)
}

func (fake *NetlinkAdapter) LinkDelArgsForCall(i int) netlink.Link {
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return fake.linkDelArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkDelReturns(result1 error) {
	fake.LinkDelStub = nil
	fake.linkDelReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkDelReturnsOnCall(i int, result1 error) {
	fake.LinkDelStub = nil
	if fake.linkDelReturnsOnCall == nil {
		fake.linkDelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkDelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkAdd(arg1 netlink.Link) error {
	fake.linkAddMutex.Lock()
	ret, specificReturn := fake.linkAddReturnsOnCall[len(fake.linkAddArgsForCall)]
	fake.linkAddArgsForCall = append(fake.linkAddArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkAdd", []interface{}{arg1})
	fake.linkAddMutex.Unlock()
	if fake.LinkAddStub != nil {
		return fake.LinkAddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkAddReturns.result1
}

func (fake *NetlinkAdapter) LinkAddCallCount() int {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return len(fake.linkAddArgsForCall)
}

func (fake *NetlinkAdapter) LinkAddArgsForCall(i int) netlink.Link {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return fake.linkAddArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkAddReturns(result1 error) {
	fake.LinkAddStub = nil
	fake.linkAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkAddReturnsOnCall(i int, result1 error) {
	fake.LinkAddStub = nil
	if fake.linkAddReturnsOnCall == nil {
		fake.linkAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetNsFd(arg1 netlink.Link, arg2 int) error {
	fake.linkSetNsFdMutex.Lock()
	ret, specificReturn := fake.linkSetNsFdReturnsOnCall[len(fake.linkSetNsFdArgsForCall)]
	fake.linkSetNsFdArgsForCall = append(fake.linkSetNsFdArgsForCall, struct {
		arg1 netlink.Link
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("LinkSetNsFd", []interface{}{arg1, arg2})
	fake.linkSetNsFdMutex.Unlock()
	if fake.LinkSetNsFdStub != nil {
		return fake.LinkSetNsFdStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetNsFdReturns.result1
}

func (fake *NetlinkAdapter) LinkSetNsFdCallCount() int {
	fake.linkSetNsFdMutex.RLock()
	defer fake.linkSetNsFdMutex.RUnlock()
	return len(fake.linkSetNsFdArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetNsFdArgsForCall(i int) (netlink.Link, int) {
	fake.linkSetNsFdMutex.RLock()
	defer fake.linkSetNsFdMutex.RUnlock()
	return fake.linkSetNsFdArgsForCall[i].arg1, fake.linkSetNsFdArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) LinkSetNsFdReturns(result1 error) {
	fake.LinkSetNsFdStub = nil
	fake.linkSetNsFdReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetNsFdReturnsOnCall(i int, result1 error) {
	fake.LinkSetNsFdStub = nil
	if fake.linkSetNsFdReturnsOnCall == nil {
		fake.linkSetNsFdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetNsFdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) RouteAdd(route *netlink.Route) error {
	fake.routeAddMutex.Lock()
	ret, specificReturn := fake.routeAddReturnsOnCall[len(fake.routeAddArgsForCall)]
	fake.routeAddArgsForCall = append(fake.routeAddArgsForCall, struct {
		route *netlink.Route
	}{route})
	fake.recordInvocation("RouteAdd", []interface{}{route})
	fake.routeAddMutex.Unlock()
	if fake.RouteAddStub != nil {
		return fake.RouteAddStub(route)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.routeAddReturns.result1
}

func (fake *NetlinkAdapter) RouteAddCallCount() int {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return len(fake.routeAddArgsForCall)
}

func (fake *NetlinkAdapter) RouteAddArgsForCall(i int) *netlink.Route {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return fake.routeAddArgsForCall[i].route
}

func (fake *NetlinkAdapter) RouteAddReturns(result1 error) {
	fake.RouteAddStub = nil
	fake.routeAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) RouteAddReturnsOnCall(i int, result1 error) {
	fake.RouteAddStub = nil
	if fake.routeAddReturnsOnCall == nil {
		fake.routeAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.routeAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) QdiscAdd(qdisc netlink.Qdisc) error {
	fake.qdiscAddMutex.Lock()
	ret, specificReturn := fake.qdiscAddReturnsOnCall[len(fake.qdiscAddArgsForCall)]
	fake.qdiscAddArgsForCall = append(fake.qdiscAddArgsForCall, struct {
		qdisc netlink.Qdisc
	}{qdisc})
	fake.recordInvocation("QdiscAdd", []interface{}{qdisc})
	fake.qdiscAddMutex.Unlock()
	if fake.QdiscAddStub != nil {
		return fake.QdiscAddStub(qdisc)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.qdiscAddReturns.result1
}

func (fake *NetlinkAdapter) QdiscAddCallCount() int {
	fake.qdiscAddMutex.RLock()
	defer fake.qdiscAddMutex.RUnlock()
	return len(fake.qdiscAddArgsForCall)
}

func (fake *NetlinkAdapter) QdiscAddArgsForCall(i int) netlink.Qdisc {
	fake.qdiscAddMutex.RLock()
	defer fake.qdiscAddMutex.RUnlock()
	return fake.qdiscAddArgsForCall[i].qdisc
}

func (fake *NetlinkAdapter) QdiscAddReturns(result1 error) {
	fake.QdiscAddStub = nil
	fake.qdiscAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) QdiscAddReturnsOnCall(i int, result1 error) {
	fake.QdiscAddStub = nil
	if fake.qdiscAddReturnsOnCall == nil {
		fake.qdiscAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.qdiscAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) FilterAdd(arg1 netlink.Filter) error {
	fake.filterAddMutex.Lock()
	ret, specificReturn := fake.filterAddReturnsOnCall[len(fake.filterAddArgsForCall)]
	fake.filterAddArgsForCall = append(fake.filterAddArgsForCall, struct {
		arg1 netlink.Filter
	}{arg1})
	fake.recordInvocation("FilterAdd", []interface{}{arg1})
	fake.filterAddMutex.Unlock()
	if fake.FilterAddStub != nil {
		return fake.FilterAddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.filterAddReturns.result1
}

func (fake *NetlinkAdapter) FilterAddCallCount() int {
	fake.filterAddMutex.RLock()
	defer fake.filterAddMutex.RUnlock()
	return len(fake.filterAddArgsForCall)
}

func (fake *NetlinkAdapter) FilterAddArgsForCall(i int) netlink.Filter {
	fake.filterAddMutex.RLock()
	defer fake.filterAddMutex.RUnlock()
	return fake.filterAddArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) FilterAddReturns(result1 error) {
	fake.FilterAddStub = nil
	fake.filterAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) FilterAddReturnsOnCall(i int, result1 error) {
	fake.FilterAddStub = nil
	if fake.filterAddReturnsOnCall == nil {
		fake.filterAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.filterAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	fake.parseAddrMutex.RLock()
	defer fake.parseAddrMutex.RUnlock()
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	fake.neighAddPermanentIPv4Mutex.RLock()
	defer fake.neighAddPermanentIPv4Mutex.RUnlock()
	fake.linkSetARPOffMutex.RLock()
	defer fake.linkSetARPOffMutex.RUnlock()
	fake.linkSetNameMutex.RLock()
	defer fake.linkSetNameMutex.RUnlock()
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	fake.linkSetNsFdMutex.RLock()
	defer fake.linkSetNsFdMutex.RUnlock()
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	fake.qdiscAddMutex.RLock()
	defer fake.qdiscAddMutex.RUnlock()
	fake.filterAddMutex.RLock()
	defer fake.filterAddMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NetlinkAdapter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
