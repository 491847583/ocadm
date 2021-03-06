package compute

import (
	"yunion.io/x/pkg/util/netutils"
)

const (
	VPC_EXTERNAL_ACCESS_MODE_DISTGW     = "distgw"     // distgw only
	VPC_EXTERNAL_ACCESS_MODE_EIP_DISTGW = "eip-distgw" // eip when available, distgw otherwise
	VPC_EXTERNAL_ACCESS_MODE_EIP        = "eip"        // eip only
)

var (
	VPC_EXTERNAL_ACCESS_MODES = []string{
		VPC_EXTERNAL_ACCESS_MODE_DISTGW,
		VPC_EXTERNAL_ACCESS_MODE_EIP_DISTGW,
		VPC_EXTERNAL_ACCESS_MODE_EIP,
	}
)

const (
	sVpcMappedCidr      = "100.64.0.0/17"
	VpcMappedIPMask     = 17
	sVpcMappedGatewayIP = "100.64.0.1"
	VpcMappedGatewayMac = "ee:ee:ee:ee:ee:ee"

	// reserved: [100.64.0.2, 100.64.0.127]

	// [128, 2176 (128+2048)]
	sVpcMappedHostIPStart = "100.64.0.128"
	sVpcMappedHostIPEnd   = "100.64.8.128"

	// reserved: [10.64.8.129, 10.64.8.255]

	// [2304, 32511], 30207
	sVpcMappedIPStart = "100.64.9.0"
	sVpcMappedIPEnd   = "100.64.126.255"

	// reserved: [10.64.127.0 , 10.64.127.255]
)

const (
	sVpcEipGatewayCidr  = "100.64.128.0/17"
	VpcEipGatewayIPMask = 17
	sVpcEipGatewayIP    = "100.64.128.2"
	VpcEipGatewayMac    = "ee:ee:ee:ee:ee:ef"

	sVpcEipGatewayIP3 = "100.64.128.3"
	VpcEipGatewayMac3 = "ee:ee:ee:ee:ee:f0"
)

var (
	vpcMappedCidr      netutils.IPV4Prefix
	vpcMappedGatewayIP netutils.IPV4Addr

	vpcEipGatewayCidr netutils.IPV4Prefix
	vpcEipGatewayIP   netutils.IPV4Addr
	vpcEipGatewayIP3  netutils.IPV4Addr

	vpcMappedHostIPStart netutils.IPV4Addr
	vpcMappedHostIPEnd   netutils.IPV4Addr

	vpcMappedIPStart netutils.IPV4Addr
	vpcMappedIPEnd   netutils.IPV4Addr
)

func init() {
	mi := func(v netutils.IPV4Addr, err error) netutils.IPV4Addr {
		if err != nil {
			panic(err.Error())
		}
		return v
	}
	mp := func(v netutils.IPV4Prefix, err error) netutils.IPV4Prefix {
		if err != nil {
			panic(err.Error())
		}
		return v
	}

	vpcMappedCidr = mp(netutils.NewIPV4Prefix(sVpcMappedCidr))
	vpcMappedGatewayIP = mi(netutils.NewIPV4Addr(sVpcMappedGatewayIP))

	vpcEipGatewayCidr = mp(netutils.NewIPV4Prefix(sVpcEipGatewayCidr))
	vpcEipGatewayIP = mi(netutils.NewIPV4Addr(sVpcEipGatewayIP))
	vpcEipGatewayIP3 = mi(netutils.NewIPV4Addr(sVpcEipGatewayIP3))

	vpcMappedHostIPStart = mi(netutils.NewIPV4Addr(sVpcMappedHostIPStart))
	vpcMappedHostIPEnd = mi(netutils.NewIPV4Addr(sVpcMappedHostIPEnd))

	vpcMappedIPStart = mi(netutils.NewIPV4Addr(sVpcMappedIPStart))
	vpcMappedIPEnd = mi(netutils.NewIPV4Addr(sVpcMappedIPEnd))
}

func VpcMappedCidr() netutils.IPV4Prefix {
	return vpcMappedCidr
}

func VpcMappedGatewayIP() netutils.IPV4Addr {
	return vpcMappedGatewayIP
}

func VpcEipGatewayCidr() netutils.IPV4Prefix {
	return vpcEipGatewayCidr
}

func VpcEipGatewayIP() netutils.IPV4Addr {
	return vpcEipGatewayIP
}

func VpcEipGatewayIP3() netutils.IPV4Addr {
	return vpcEipGatewayIP3
}

func VpcMappedHostIPStart() netutils.IPV4Addr {
	return vpcMappedHostIPStart
}

func VpcMappedHostIPEnd() netutils.IPV4Addr {
	return vpcMappedHostIPEnd
}

func VpcMappedIPStart() netutils.IPV4Addr {
	return vpcMappedIPStart
}

func VpcMappedIPEnd() netutils.IPV4Addr {
	return vpcMappedIPEnd
}
