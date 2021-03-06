// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Server struct {
	Id             int64       `db:"id" json:"id"`
	HostName       string      `db:"host_name" json:"hostName"`
	DomainName     string      `db:"domain_name" json:"domainName"`
	TcpPort        null.Int    `db:"tcp_port" json:"tcpPort"`
	XmppId         null.String `db:"xmpp_id" json:"xmppId"`
	XmppPasswd     null.String `db:"xmpp_passwd" json:"xmppPasswd"`
	InterfaceName  string      `db:"interface_name" json:"interfaceName"`
	IpAddress      string      `db:"ip_address" json:"ipAddress"`
	IpNetmask      string      `db:"ip_netmask" json:"ipNetmask"`
	IpGateway      string      `db:"ip_gateway" json:"ipGateway"`
	Ip6Address     null.String `db:"ip6_address" json:"ip6Address"`
	Ip6Gateway     null.String `db:"ip6_gateway" json:"ip6Gateway"`
	InterfaceMtu   int64       `db:"interface_mtu" json:"interfaceMtu"`
	Rack           null.String `db:"rack" json:"rack"`
	UpdPending     int64       `db:"upd_pending" json:"updPending"`
	MgmtIpAddress  null.String `db:"mgmt_ip_address" json:"mgmtIpAddress"`
	MgmtIpNetmask  null.String `db:"mgmt_ip_netmask" json:"mgmtIpNetmask"`
	MgmtIpGateway  null.String `db:"mgmt_ip_gateway" json:"mgmtIpGateway"`
	IloIpAddress   null.String `db:"ilo_ip_address" json:"iloIpAddress"`
	IloIpNetmask   null.String `db:"ilo_ip_netmask" json:"iloIpNetmask"`
	IloIpGateway   null.String `db:"ilo_ip_gateway" json:"iloIpGateway"`
	IloUsername    null.String `db:"ilo_username" json:"iloUsername"`
	IloPassword    null.String `db:"ilo_password" json:"iloPassword"`
	RouterHostName null.String `db:"router_host_name" json:"routerHostName"`
	RouterPortName null.String `db:"router_port_name" json:"routerPortName"`
	LastUpdated    time.Time   `db:"last_updated" json:"lastUpdated"`
	Links          ServerLinks `json:"_links" db:-`
}

type ServerLinks struct {
	Self             string           `db:"self" json:"_self"`
	TypeLink         TypeLink         `json:"type" db:-`
	StatusLink       StatusLink       `json:"status" db:-`
	ProfileLink      ProfileLink      `json:"profile" db:-`
	CdnLink          CdnLink          `json:"cdn" db:-`
	PhysLocationLink PhysLocationLink `json:"phys_location" db:-`
	CachegroupLink   CachegroupLink   `json:"cachegroup" db:-`
}

type ServerLink struct {
	ID  int64  `db:"server" json:"id"`
	Ref string `db:"server_id_ref" json:"_ref"`
}

// @Title getServerById
// @Description retrieves the server information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Server
// @Resource /api/2.0
// @Router /api/2.0/server/{id} [get]
func getServerById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []Server{}
	arg := Server{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "server/', id) as self "
	queryStr += ", concat('" + API_PATH + "phys_location/', phys_location) as phys_location_id_ref"
	queryStr += ", concat('" + API_PATH + "cachegroup/', cachegroup) as cachegroup_id_ref"
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += ", concat('" + API_PATH + "status/', status) as status_id_ref"
	queryStr += ", concat('" + API_PATH + "profile/', profile) as profile_id_ref"
	queryStr += ", concat('" + API_PATH + "cdn/', cdn) as cdn_id_ref"
	queryStr += " from server where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getServers
// @Description retrieves the server
// @Accept  application/json
// @Success 200 {array}    Server
// @Resource /api/2.0
// @Router /api/2.0/server [get]
func getServers(db *sqlx.DB) (interface{}, error) {
	ret := []Server{}
	queryStr := "select *, concat('" + API_PATH + "server/', id) as self "
	queryStr += ", concat('" + API_PATH + "phys_location/', phys_location) as phys_location_id_ref"
	queryStr += ", concat('" + API_PATH + "cachegroup/', cachegroup) as cachegroup_id_ref"
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += ", concat('" + API_PATH + "status/', status) as status_id_ref"
	queryStr += ", concat('" + API_PATH + "profile/', profile) as profile_id_ref"
	queryStr += ", concat('" + API_PATH + "cdn/', cdn) as cdn_id_ref"
	queryStr += " from server"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postServer
// @Description enter a new server
// @Accept  application/json
// @Param                 Body body     Server   true "Server object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/server [post]
func postServer(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Server
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO server("
	sqlString += "host_name"
	sqlString += ",domain_name"
	sqlString += ",tcp_port"
	sqlString += ",xmpp_id"
	sqlString += ",xmpp_passwd"
	sqlString += ",interface_name"
	sqlString += ",ip_address"
	sqlString += ",ip_netmask"
	sqlString += ",ip_gateway"
	sqlString += ",ip6_address"
	sqlString += ",ip6_gateway"
	sqlString += ",interface_mtu"
	sqlString += ",phys_location"
	sqlString += ",rack"
	sqlString += ",cachegroup"
	sqlString += ",type"
	sqlString += ",status"
	sqlString += ",upd_pending"
	sqlString += ",profile"
	sqlString += ",cdn"
	sqlString += ",mgmt_ip_address"
	sqlString += ",mgmt_ip_netmask"
	sqlString += ",mgmt_ip_gateway"
	sqlString += ",ilo_ip_address"
	sqlString += ",ilo_ip_netmask"
	sqlString += ",ilo_ip_gateway"
	sqlString += ",ilo_username"
	sqlString += ",ilo_password"
	sqlString += ",router_host_name"
	sqlString += ",router_port_name"
	sqlString += ") VALUES ("
	sqlString += ":host_name"
	sqlString += ",:domain_name"
	sqlString += ",:tcp_port"
	sqlString += ",:xmpp_id"
	sqlString += ",:xmpp_passwd"
	sqlString += ",:interface_name"
	sqlString += ",:ip_address"
	sqlString += ",:ip_netmask"
	sqlString += ",:ip_gateway"
	sqlString += ",:ip6_address"
	sqlString += ",:ip6_gateway"
	sqlString += ",:interface_mtu"
	sqlString += ",:phys_location"
	sqlString += ",:rack"
	sqlString += ",:cachegroup"
	sqlString += ",:type"
	sqlString += ",:status"
	sqlString += ",:upd_pending"
	sqlString += ",:profile"
	sqlString += ",:cdn"
	sqlString += ",:mgmt_ip_address"
	sqlString += ",:mgmt_ip_netmask"
	sqlString += ",:mgmt_ip_gateway"
	sqlString += ",:ilo_ip_address"
	sqlString += ",:ilo_ip_netmask"
	sqlString += ",:ilo_ip_gateway"
	sqlString += ",:ilo_username"
	sqlString += ",:ilo_password"
	sqlString += ",:router_host_name"
	sqlString += ",:router_port_name"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putServer
// @Description modify an existing serverentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Server   true "Server object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/server/{id}  [put]
func putServer(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Server
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE server SET "
	sqlString += "host_name = :host_name"
	sqlString += ",domain_name = :domain_name"
	sqlString += ",tcp_port = :tcp_port"
	sqlString += ",xmpp_id = :xmpp_id"
	sqlString += ",xmpp_passwd = :xmpp_passwd"
	sqlString += ",interface_name = :interface_name"
	sqlString += ",ip_address = :ip_address"
	sqlString += ",ip_netmask = :ip_netmask"
	sqlString += ",ip_gateway = :ip_gateway"
	sqlString += ",ip6_address = :ip6_address"
	sqlString += ",ip6_gateway = :ip6_gateway"
	sqlString += ",interface_mtu = :interface_mtu"
	sqlString += ",phys_location = :phys_location"
	sqlString += ",rack = :rack"
	sqlString += ",cachegroup = :cachegroup"
	sqlString += ",type = :type"
	sqlString += ",status = :status"
	sqlString += ",upd_pending = :upd_pending"
	sqlString += ",profile = :profile"
	sqlString += ",cdn = :cdn"
	sqlString += ",mgmt_ip_address = :mgmt_ip_address"
	sqlString += ",mgmt_ip_netmask = :mgmt_ip_netmask"
	sqlString += ",mgmt_ip_gateway = :mgmt_ip_gateway"
	sqlString += ",ilo_ip_address = :ilo_ip_address"
	sqlString += ",ilo_ip_netmask = :ilo_ip_netmask"
	sqlString += ",ilo_ip_gateway = :ilo_ip_gateway"
	sqlString += ",ilo_username = :ilo_username"
	sqlString += ",ilo_password = :ilo_password"
	sqlString += ",router_host_name = :router_host_name"
	sqlString += ",router_port_name = :router_port_name"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delServerById
// @Description deletes server information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Server
// @Resource /api/2.0
// @Router /api/2.0/server/{id} [delete]
func delServer(id int, db *sqlx.DB) (interface{}, error) {
	arg := Server{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM server WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
