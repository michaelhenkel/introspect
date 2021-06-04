package v1alpha1

import "encoding/xml"

func (b *VirtualMachineInterface) GetName() string {
	return ""
}

type ItfSandeshData struct {
	Text  string `xml:",chardata"`
	Index struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"index"`
	Name struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"name"`
	Uuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"uuid"`
	VrfName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"vrf_name"`
	Active struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"active"`
	Ipv4Active struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"ipv4_active"`
	L2Active struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"l2_active"`
	Ip6Active struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"ip6_active"`
	HealthCheckActive struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"health_check_active"`
	DhcpService struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"dhcp_service"`
	DnsService struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"dns_service"`
	Type struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"type"`
	Label struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"label"`
	L2Label struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"l2_label"`
	VxlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vxlan_id"`
	VnName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"vn_name"`
	VmUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"vm_uuid"`
	VmName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vm_name"`
	IpAddr struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"ip_addr"`
	MacAddr struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"mac_addr"`
	Policy struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"policy"`
	FipList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text                  string `xml:",chardata"`
			Type                  string `xml:"type,attr"`
			Size                  string `xml:"size,attr"`
			FloatingIpSandeshList []struct {
				Text   string `xml:",chardata"`
				IpAddr struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"ip_addr"`
				VrfName struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"vrf_name"`
				Installed struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"installed"`
				FixedIp struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"fixed_ip"`
				Direction struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"direction"`
				PortMapEnabled struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"port_map_enabled"`
				PortMap struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					List       struct {
						Text               string `xml:",chardata"`
						Type               string `xml:"type,attr"`
						Size               string `xml:"size,attr"`
						SandeshPortMapping []struct {
							Text     string `xml:",chardata"`
							Protocol struct {
								Text       string `xml:",chardata"`
								Type       string `xml:"type,attr"`
								Identifier string `xml:"identifier,attr"`
							} `xml:"protocol"`
							Port struct {
								Text       string `xml:",chardata"`
								Type       string `xml:"type,attr"`
								Identifier string `xml:"identifier,attr"`
							} `xml:"port"`
							NatPort struct {
								Text       string `xml:",chardata"`
								Type       string `xml:"type,attr"`
								Identifier string `xml:"identifier,attr"`
							} `xml:"nat_port"`
						} `xml:"SandeshPortMapping"`
					} `xml:"list"`
				} `xml:"port_map"`
			} `xml:"FloatingIpSandeshList"`
		} `xml:"list"`
	} `xml:"fip_list"`
	MdataIpAddr struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"mdata_ip_addr"`
	ServiceVlanList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"service_vlan_list"`
	OsIfindex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"os_ifindex"`
	FabricPort struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"fabric_port"`
	AllocLinklocalIp struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"alloc_linklocal_ip"`
	AnalyzerName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"analyzer_name"`
	ConfigName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"config_name"`
	SgUuidList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"sg_uuid_list"`
	StaticRouteList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text               string `xml:",chardata"`
			Type               string `xml:"type,attr"`
			Size               string `xml:"size,attr"`
			StaticRouteSandesh struct {
				Text    string `xml:",chardata"`
				VrfName struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"vrf_name"`
				IpAddr struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"ip_addr"`
				Prefix struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"prefix"`
				Communities struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					List       struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
						Size string `xml:"size,attr"`
					} `xml:"list"`
				} `xml:"communities"`
			} `xml:"StaticRouteSandesh"`
		} `xml:"list"`
	} `xml:"static_route_list"`
	VmProjectUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vm_project_uuid"`
	AdminState struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"admin_state"`
	FlowKeyIdx struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"flow_key_idx"`
	AllowedAddressPairList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"allowed_address_pair_list"`
	Ip6Addr struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"ip6_addr"`
	LocalPreference struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"local_preference"`
	TxVlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"tx_vlan_id"`
	RxVlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"rx_vlan_id"`
	ParentInterface struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"parent_interface"`
	Subnet struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"subnet"`
	SubType struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"sub_type"`
	VrfAssignAclUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"vrf_assign_acl_uuid"`
	VmiType struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vmi_type"`
	Transport struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"transport"`
	LogicalInterfaceUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"logical_interface_uuid"`
	FloodUnknownUnicast struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"flood_unknown_unicast"`
	PhysicalDevice struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"physical_device"`
	PhysicalInterface struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"physical_interface"`
	FixedIp4List struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			Size    string `xml:"size,attr"`
			Element string `xml:"element"`
		} `xml:"list"`
	} `xml:"fixed_ip4_list"`
	FixedIp6List struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"fixed_ip6_list"`
	FatFlowList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"fat_flow_list"`
	MetadataIpActive struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"metadata_ip_active"`
	AliasIpList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"alias_ip_list"`
	DropNewFlows struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"drop_new_flows"`
	BridgeDomainList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"bridge_domain_list"`
	VmiTagList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"vmi_tag_list"`
	PolicySetAclList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"policy_set_acl_list"`
	SloList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"slo_list"`
	VhostuserMode struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vhostuser_mode"`
	SiOtherEndVmi struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"si_other_end_vmi"`
	CfgIgmpEnable struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"cfg_igmp_enable"`
	IgmpEnabled struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"igmp_enabled"`
	MaxFlows struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"max_flows"`
	PolicySetFwaasList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"policy_set_fwaas_list"`
	BondInterfaceList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"bond_interface_list"`
	MacIpLearningEnable struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"mac_ip_learning_enable"`
	MacIpList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"mac_ip_list"`
	VhostsocketDir struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vhostsocket_dir"`
	VhostsocketFilename struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vhostsocket_filename"`
	ServiceHealthCheckIp struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"service_health_check_ip"`
}
type VirtualMachineInterface struct {
	XMLName xml.Name `xml:"__ItfResp_list"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
	ItfResp struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"type,attr"`
		ItfList struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
			List       struct {
				Text           string           `xml:",chardata"`
				Type           string           `xml:"type,attr"`
				Size           string           `xml:"size,attr"`
				ItfSandeshData []ItfSandeshData `xml:"ItfSandeshData"`
			} `xml:"list"`
		} `xml:"itf_list"`
		More struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
		} `xml:"more"`
	} `xml:"ItfResp"`
	Pagination struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Req  struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			Identifier  string `xml:"identifier,attr"`
			PageReqData struct {
				Text     string `xml:",chardata"`
				PrevPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"prev_page"`
				NextPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"next_page"`
				FirstPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"first_page"`
				All struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"all"`
				TableSize struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"table_size"`
				Entries struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"entries"`
			} `xml:"PageReqData"`
		} `xml:"req"`
		More struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
		} `xml:"more"`
	} `xml:"Pagination"`
}

func (b *VirtualNetwork) GetName() string {
	return b.VnListResp.VnList.List.VnSandeshData[0].Name.Text
}

type VnSandeshData struct {
	Text string `xml:",chardata"`
	Name struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"name"`
	Uuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"uuid"`
	AclUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"acl_uuid"`
	MirrorAclUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"mirror_acl_uuid"`
	MirrorCfgAclUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"mirror_cfg_acl_uuid"`
	VrfName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"vrf_name"`
	IpamData struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Size       string `xml:"size,attr"`
			VnIpamData struct {
				Text     string `xml:",chardata"`
				IpPrefix struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"ip_prefix"`
				PrefixLen struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"prefix_len"`
				Gateway struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"gateway"`
				IpamName struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"ipam_name"`
				DhcpEnable struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"dhcp_enable"`
				DnsServer struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"dns_server"`
			} `xml:"VnIpamData"`
		} `xml:"list"`
	} `xml:"ipam_data"`
	IpamHostRoutes struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text             string `xml:",chardata"`
			Type             string `xml:"type,attr"`
			Size             string `xml:"size,attr"`
			VnIpamHostRoutes struct {
				Text     string `xml:",chardata"`
				IpamName struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"ipam_name"`
				HostRoutes struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					List       struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
						Size string `xml:"size,attr"`
					} `xml:"list"`
				} `xml:"host_routes"`
			} `xml:"VnIpamHostRoutes"`
		} `xml:"list"`
	} `xml:"ipam_host_routes"`
	Layer2Forwarding struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"layer2_forwarding"`
	Ipv4Forwarding struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"ipv4_forwarding"`
	AdminState struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"admin_state"`
	VxlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vxlan_id"`
	ConfigVxlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"config_vxlan_id"`
	VnID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vn_id"`
	EnableRpf struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"enable_rpf"`
	Bridging struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"bridging"`
	FloodUnknownUnicast struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"flood_unknown_unicast"`
	PbbEtreeEnabled struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"pbb_etree_enabled"`
	Layer2ControlWord struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"layer2_control_word"`
	SloList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"slo_list"`
	CfgIgmpEnable struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"cfg_igmp_enable"`
	MaxFlows struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"max_flows"`
	MpList struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"mp_list"`
	MacIpLearningEnable struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"mac_ip_learning_enable"`
	HealthCheckUuid struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"health_check_uuid"`
}

type VirtualNetwork struct {
	XMLName    xml.Name `xml:"__VnListResp_list"`
	Text       string   `xml:",chardata"`
	Type       string   `xml:"type,attr"`
	VnListResp struct {
		Text   string `xml:",chardata"`
		Type   string `xml:"type,attr"`
		VnList struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
			List       struct {
				Text          string          `xml:",chardata"`
				Type          string          `xml:"type,attr"`
				Size          string          `xml:"size,attr"`
				VnSandeshData []VnSandeshData `xml:"VnSandeshData"`
			} `xml:"list"`
		} `xml:"vn_list"`
		More struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
		} `xml:"more"`
	} `xml:"VnListResp"`
	Pagination struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Req  struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			Identifier  string `xml:"identifier,attr"`
			PageReqData struct {
				Text     string `xml:",chardata"`
				PrevPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"prev_page"`
				NextPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"next_page"`
				FirstPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"first_page"`
				All struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"all"`
				TableSize struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"table_size"`
				Entries struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"entries"`
			} `xml:"PageReqData"`
		} `xml:"req"`
		More struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
		} `xml:"more"`
	} `xml:"Pagination"`
}

func (b *RoutingInstance) GetName() string {
	return ""
}

type VrfSandeshData struct {
	Text string `xml:",chardata"`
	Name struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"name"`
	Ucindex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"ucindex"`
	Mcindex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"mcindex"`
	L2index struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"l2index"`
	Source struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"source"`
	Uc6index struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"uc6index"`
	Vn struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"vn"`
	TableLabel struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"table_label"`
	VxlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vxlan_id"`
	Evpnindex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"evpnindex"`
	Brindex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"brindex"`
	Mplsindex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"mplsindex"`
	RD struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"RD"`
	MacAgingTime struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"mac_aging_time"`
	Layer2ControlWord struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"layer2_control_word"`
	ForwardingVrf struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"forwarding_vrf"`
	HbfRintf struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"hbf_rintf"`
	HbfLintf struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"hbf_lintf"`
}

type RoutingInstance struct {
	XMLName     xml.Name `xml:"__VrfListResp_list"`
	Text        string   `xml:",chardata"`
	Type        string   `xml:"type,attr"`
	VrfListResp struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"type,attr"`
		VrfList struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
			List       struct {
				Text           string           `xml:",chardata"`
				Type           string           `xml:"type,attr"`
				Size           string           `xml:"size,attr"`
				VrfSandeshData []VrfSandeshData `xml:"VrfSandeshData"`
			} `xml:"list"`
		} `xml:"vrf_list"`
		More struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
		} `xml:"more"`
	} `xml:"VrfListResp"`
	Pagination struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Req  struct {
			Text        string `xml:",chardata"`
			Type        string `xml:"type,attr"`
			Identifier  string `xml:"identifier,attr"`
			PageReqData struct {
				Text     string `xml:",chardata"`
				PrevPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"prev_page"`
				NextPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"next_page"`
				FirstPage struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"first_page"`
				All struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"all"`
				TableSize struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"table_size"`
				Entries struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"entries"`
			} `xml:"PageReqData"`
		} `xml:"req"`
		More struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Identifier string `xml:"identifier,attr"`
		} `xml:"more"`
	} `xml:"Pagination"`
}
