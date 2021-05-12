package v1alpha1

import "encoding/xml"

type BgpNeighborResp struct {
	Text         string `xml:",chardata"`
	InstanceName struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"instance_name"`
	Peer struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"peer"`
	Deleted struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"deleted"`
	PeerAddress struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"peer_address"`
	PeerID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"peer_id"`
	PeerAsn struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"peer_asn"`
	Encoding struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"encoding"`
	PeerType struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"peer_type"`
	State struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"state"`
	ClosedAt struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"closed_at"`
	RouterType struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"router_type"`
	AdminDown struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"admin_down"`
	Passive struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"passive"`
	AsOverride struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"as_override"`
	OriginOverride struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"origin_override"`
	RouteOrigin struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"route_origin"`
	As4Supported struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"as4_supported"`
	PrivateAsAction struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"private_as_action"`
	SendReady struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"send_ready"`
	PeerPort struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"peer_port"`
	TransportAddress struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"transport_address"`
	LocalAddress struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"local_address"`
	LocalID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"local_id"`
	LocalAsn struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"local_asn"`
	ClusterID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"cluster_id"`
	TaskInstance struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"task_instance"`
	PeerCloseInfo struct {
		Text          string `xml:",chardata"`
		Type          string `xml:"type,attr"`
		Identifier    string `xml:"identifier,attr"`
		PeerCloseInfo struct {
			Text  string `xml:",chardata"`
			State struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"state"`
			MembershipState struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"membership_state"`
			CloseAgain struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"close_again"`
			Graceful struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"graceful"`
			Init struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"init"`
			Close struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"close"`
			Nested struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"nested"`
			Deletes struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"deletes"`
			Stale struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"stale"`
			LlgrStale struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"llgr_stale"`
			Sweep struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"sweep"`
			GrTimer struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"gr_timer"`
			LlgrTimer struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"llgr_timer"`
		} `xml:"PeerCloseInfo"`
	} `xml:"peer_close_info"`
	SendState struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"send_state"`
	LastEvent struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"last_event"`
	LastState struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"last_state"`
	LastStateAt struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"last_state_at"`
	LastError struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"last_error"`
	AuthType struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"auth_type"`
	ConfiguredAddressFamilies struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"configured_address_families"`
	NegotiatedAddressFamilies struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"negotiated_address_families"`
	ConfiguredHoldTime struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"configured_hold_time"`
	NegotiatedHoldTime struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"negotiated_hold_time"`
	PrimaryPathCount struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"primary_path_count"`
	FlapCount struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"flap_count"`
	FlapTime struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"flap_time"`
	RoutingInstances struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text                       string `xml:",chardata"`
			Type                       string `xml:"type,attr"`
			Size                       string `xml:"size,attr"`
			BgpNeighborRoutingInstance []struct {
				Text string `xml:",chardata"`
				Name struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"name"`
				State struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"state"`
				Index struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"index"`
				ImportTargets struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					List       struct {
						Text    string   `xml:",chardata"`
						Type    string   `xml:"type,attr"`
						Size    string   `xml:"size,attr"`
						Element []string `xml:"element"`
					} `xml:"list"`
				} `xml:"import_targets"`
			} `xml:"BgpNeighborRoutingInstance"`
		} `xml:"list"`
	} `xml:"routing_instances"`
	RoutingTables struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text                    string `xml:",chardata"`
			Type                    string `xml:"type,attr"`
			Size                    string `xml:"size,attr"`
			BgpNeighborRoutingTable []struct {
				Text string `xml:",chardata"`
				Name struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"name"`
				CurrentState struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"current_state"`
				CurrentRequest struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"current_request"`
				PendingRequest struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"pending_request"`
			} `xml:"BgpNeighborRoutingTable"`
		} `xml:"list"`
	} `xml:"routing_tables"`
	RxProtoStats struct {
		Text           string `xml:",chardata"`
		Type           string `xml:"type,attr"`
		Identifier     string `xml:"identifier,attr"`
		PeerProtoStats struct {
			Text  string `xml:",chardata"`
			Total struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"total"`
			Open struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"open"`
			Keepalive struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"keepalive"`
			Notification struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"notification"`
			Update struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"update"`
			Close struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"close"`
		} `xml:"PeerProtoStats"`
	} `xml:"rx_proto_stats"`
	TxProtoStats struct {
		Text           string `xml:",chardata"`
		Type           string `xml:"type,attr"`
		Identifier     string `xml:"identifier,attr"`
		PeerProtoStats struct {
			Text  string `xml:",chardata"`
			Total struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"total"`
			Open struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"open"`
			Keepalive struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"keepalive"`
			Notification struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"notification"`
			Update struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"update"`
			Close struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"close"`
		} `xml:"PeerProtoStats"`
	} `xml:"tx_proto_stats"`
	RxUpdateStats struct {
		Text            string `xml:",chardata"`
		Type            string `xml:"type,attr"`
		Identifier      string `xml:"identifier,attr"`
		PeerUpdateStats struct {
			Text  string `xml:",chardata"`
			Total struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"total"`
			Reach struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"reach"`
			Unreach struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"unreach"`
			EndOfRib struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"end_of_rib"`
		} `xml:"PeerUpdateStats"`
	} `xml:"rx_update_stats"`
	TxUpdateStats struct {
		Text            string `xml:",chardata"`
		Type            string `xml:"type,attr"`
		Identifier      string `xml:"identifier,attr"`
		PeerUpdateStats struct {
			Text  string `xml:",chardata"`
			Total struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"total"`
			Reach struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"reach"`
			Unreach struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"unreach"`
			EndOfRib struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"end_of_rib"`
		} `xml:"PeerUpdateStats"`
	} `xml:"tx_update_stats"`
	RxSocketStats struct {
		Text            string `xml:",chardata"`
		Type            string `xml:"type,attr"`
		Identifier      string `xml:"identifier,attr"`
		PeerSocketStats struct {
			Text  string `xml:",chardata"`
			Bytes struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"bytes"`
			Calls struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"calls"`
			AverageBytes struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"average_bytes"`
			BlockedDuration struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"blocked_duration"`
			BlockedCount struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"blocked_count"`
			AverageBlockedDuration struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"average_blocked_duration"`
		} `xml:"PeerSocketStats"`
	} `xml:"rx_socket_stats"`
	TxSocketStats struct {
		Text            string `xml:",chardata"`
		Type            string `xml:"type,attr"`
		Identifier      string `xml:"identifier,attr"`
		PeerSocketStats struct {
			Text  string `xml:",chardata"`
			Bytes struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"bytes"`
			Calls struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"calls"`
			AverageBytes struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"average_bytes"`
			BlockedDuration struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"blocked_duration"`
			BlockedCount struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"blocked_count"`
			AverageBlockedDuration struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Identifier string `xml:"identifier,attr"`
			} `xml:"average_blocked_duration"`
		} `xml:"PeerSocketStats"`
	} `xml:"tx_socket_stats"`
	DscpValue struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"dscp_value"`
}

type BgpNeighbor struct {
	XMLName   xml.Name `xml:"BgpNeighborListResp"`
	Text      string   `xml:",chardata"`
	Type      string   `xml:"type,attr"`
	Neighbors struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text            string            `xml:",chardata"`
			Type            string            `xml:"type,attr"`
			Size            string            `xml:"size,attr"`
			BgpNeighborResp []BgpNeighborResp `xml:"BgpNeighborResp"`
		} `xml:"list"`
	} `xml:"neighbors"`
	More struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"more"`
}

func (b *BgpNeighbor) GetAddress() string {
	return b.Neighbors.List.BgpNeighborResp[0].PeerAddress.Text
}

func (b *BgpNeighbor) GetState() string {
	return b.Neighbors.List.BgpNeighborResp[0].State.Text
}

func (b *BgpNeighbor) GetName() string {
	return b.Neighbors.List.BgpNeighborResp[0].Peer.Text
}

type RoutingInstance struct {
	XMLName   xml.Name `xml:"ShowRoutingInstanceResp"`
	Text      string   `xml:",chardata"`
	Type      string   `xml:"type,attr"`
	Instances struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text                string                `xml:",chardata"`
			Type                string                `xml:"type,attr"`
			Size                string                `xml:"size,attr"`
			ShowRoutingInstance []ShowRoutingInstance `xml:"ShowRoutingInstance"`
		} `xml:"list"`
	} `xml:"instances"`
	More struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"more"`
}

type ShowRoutingInstance struct {
	Text string `xml:",chardata"`
	Name struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		Link       string `xml:"link,attr"`
	} `xml:"name"`
	VirtualNetwork struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"virtual_network"`
	VnIndex struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vn_index"`
	VxlanID struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"vxlan_id"`
	ImportTarget struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text    string   `xml:",chardata"`
			Type    string   `xml:"type,attr"`
			Size    string   `xml:"size,attr"`
			Element []string `xml:"element"`
		} `xml:"list"`
	} `xml:"import_target"`
	ExportTarget struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text    string   `xml:",chardata"`
			Type    string   `xml:"type,attr"`
			Size    string   `xml:"size,attr"`
			Element []string `xml:"element"`
		} `xml:"list"`
	} `xml:"export_target"`
	AlwaysSubscribe struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"always_subscribe"`
	AllowTransit struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"allow_transit"`
	PbbEvpnEnable struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"pbb_evpn_enable"`
	Deleted struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"deleted"`
	DeletedAt struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
	} `xml:"deleted_at"`
	Tables struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text                     string `xml:",chardata"`
			Type                     string `xml:"type,attr"`
			Size                     string `xml:"size,attr"`
			ShowRoutingInstanceTable []struct {
				Text string `xml:",chardata"`
				Name struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
					Link       string `xml:"link,attr"`
				} `xml:"name"`
				Deleted struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"deleted"`
				Prefixes struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"prefixes"`
				Paths struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"paths"`
				PrimaryPaths struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"primary_paths"`
				SecondaryPaths struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"secondary_paths"`
				InfeasiblePaths struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"infeasible_paths"`
				StalePaths struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"stale_paths"`
				LlgrStalePaths struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"llgr_stale_paths"`
				WalkRequests struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"walk_requests"`
				WalkAgainRequests struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"walk_again_requests"`
				ActualWalks struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"actual_walks"`
				WalkCompletes struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"walk_completes"`
				WalkCancels struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"walk_cancels"`
				PendingUpdates struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"pending_updates"`
				Markers struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"markers"`
				Listeners struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"listeners"`
				Walkers struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Identifier string `xml:"identifier,attr"`
				} `xml:"walkers"`
				Membership struct {
					Text                    string `xml:",chardata"`
					Type                    string `xml:"type,attr"`
					Identifier              string `xml:"identifier,attr"`
					ShowTableMembershipInfo struct {
						Text     string `xml:",chardata"`
						Requests struct {
							Text       string `xml:",chardata"`
							Type       string `xml:"type,attr"`
							Identifier string `xml:"identifier,attr"`
						} `xml:"requests"`
						Walks struct {
							Text       string `xml:",chardata"`
							Type       string `xml:"type,attr"`
							Identifier string `xml:"identifier,attr"`
						} `xml:"walks"`
						Peers struct {
							Text       string `xml:",chardata"`
							Type       string `xml:"type,attr"`
							Identifier string `xml:"identifier,attr"`
							List       struct {
								Text                   string `xml:",chardata"`
								Type                   string `xml:"type,attr"`
								Size                   string `xml:"size,attr"`
								ShowMembershipPeerInfo struct {
									Text string `xml:",chardata"`
									Peer struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr"`
										Identifier string `xml:"identifier,attr"`
									} `xml:"peer"`
									RiboutRegistered struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr"`
										Identifier string `xml:"identifier,attr"`
									} `xml:"ribout_registered"`
									RibinRegistered struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr"`
										Identifier string `xml:"identifier,attr"`
									} `xml:"ribin_registered"`
									InstanceID struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr"`
										Identifier string `xml:"identifier,attr"`
									} `xml:"instance_id"`
									GenerationID struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr"`
										Identifier string `xml:"identifier,attr"`
									} `xml:"generation_id"`
								} `xml:"ShowMembershipPeerInfo"`
							} `xml:"list"`
						} `xml:"peers"`
					} `xml:"ShowTableMembershipInfo"`
				} `xml:"membership"`
			} `xml:"ShowRoutingInstanceTable"`
		} `xml:"list"`
	} `xml:"tables"`
	RoutingPolicies struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"routing_policies"`
	Neighbors struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Identifier string `xml:"identifier,attr"`
		List       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Size string `xml:"size,attr"`
		} `xml:"list"`
	} `xml:"neighbors"`
}
