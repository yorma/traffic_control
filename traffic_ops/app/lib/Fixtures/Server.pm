package Fixtures::Server;
#
# Copyright 2015 Comcast Cable Communications Management, LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
use Moose;
extends 'DBIx::Class::EasyFixture';
use namespace::autoclean;

my %definition_for = (
	server_edge1 => {
		new   => 'Server',
		using => {
			id               => 1,
			host_name        => 'atlanta-edge-01',
			domain_name      => 'ga.atlanta.kabletown.net',
			tcp_port         => 80,
			xmpp_id          => 'atlanta-edge-01\@ocdn.kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.1',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.1',
			ip6_address      => '127.0.0.1',
			ip6_gateway      => '127.0.0.1',
			interface_mtu    => 9000,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 1,
			status           => 2,
			profile          => 1,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
	server_mid1 => {
		new   => 'Server',
		using => {
			id               => 2,
			host_name        => 'atlanta-mid-01',
			domain_name      => 'ga.atlanta.kabletown.net',
			tcp_port         => 80,
			xmpp_id          => 'atlanta-mid-01\@ocdn.kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.2',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.2',
			ip6_address      => '127.0.0.2',
			ip6_gateway      => '127.0.0.2',
			interface_mtu    => 9000,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 2,
			status           => 2,
			profile          => 2,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
	redis_server => {
		new   => 'Server',
		using => {
			id               => 3,
			host_name        => 'redis01',
			domain_name      => 'kabletown.net',
			tcp_port         => 6379,
			xmpp_id          => 'redis\@kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.3',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.3',
			ip6_address      => '127.0.0.3',
			ip6_gateway      => '127.0.0.3',
			interface_mtu    => 9000,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 25,
			status           => 2,
			profile          => 4,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
	rascal_server => {
		new   => 'Server',
		using => {
			id               => 4,
			host_name        => 'rascal01',
			domain_name      => 'kabletown.net',
			tcp_port         => 81,
			xmpp_id          => 'rascal\@kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.4',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.4',
			ip6_address      => '127.0.0.4',
			ip6_gateway      => '127.0.0.4',
			interface_mtu    => 9000,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 14,
			status           => 2,
			profile          => 6,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
	riak_server1 => {
		new   => 'Server',
		using => {
			id               => 5,
			host_name        => 'riak01',
			domain_name      => 'kabletown.net',
			tcp_port         => 8088,
			xmpp_id          => '',
			xmpp_passwd      => '',
			interface_name   => 'eth1',
			ip_address       => '127.0.0.5',
			ip_netmask       => '255.255.252.0',
			ip_gateway       => '127.0.0.5',
			ip6_address      => '',
			ip6_gateway      => '',
			interface_mtu    => 1500,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 31,
			status           => 2,
			profile          => 5,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
	rascal_server2 => {
		new   => 'Server',
		using => {
			id               => 6,
			host_name        => 'rascal02',
			domain_name      => 'kabletown.net',
			tcp_port         => 81,
			xmpp_id          => 'rascal\@kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.6',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.6',
			ip6_address      => '127.0.0.6',
			ip6_gateway      => '127.0.0.6',
			interface_mtu    => 9000,
			rack             => 'RR 119.05',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 14,
			status           => 2,
			profile          => 7,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
	server_edge2 => {
		new   => 'Server',
		using => {
			id               => 7,
			host_name        => 'atlanta-edge-02',
			domain_name      => 'ga.atlanta.kabletown.net',
			tcp_port         => 80,
			xmpp_id          => 'atlanta-edge-02\@ocdn.kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.7',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.7',
			ip6_address      => '127.0.0.7',
			ip6_gateway      => '127.0.0.7',
			interface_mtu    => 9000,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 1,
			status           => 2,
			profile          => 1,
			cachegroup       => 2,
			phys_location    => 1,
		},
	},
	server_mid2 => {
		new   => 'Server',
		using => {
			id               => 8,
			host_name        => 'atlanta-mid-02',
			domain_name      => 'ga.atlanta.kabletown.net',
			tcp_port         => 80,
			xmpp_id          => 'atlanta-mid-02\@ocdn.kabletown.net',
			xmpp_passwd      => 'X',
			interface_name   => 'bond0',
			ip_address       => '127.0.0.8',
			ip_netmask       => '255.255.255.252',
			ip_gateway       => '127.0.0.8',
			ip6_address      => '127.0.0.8',
			ip6_gateway      => '127.0.0.8',
			interface_mtu    => 9000,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 2,
			status           => 2,
			profile          => 2,
			cachegroup       => 2,
			phys_location    => 1,
		},
	},
	riak_server2 => {
		new   => 'Server',
		using => {
			id               => 9,
			host_name        => 'riak02',
			domain_name      => 'kabletown.net',
			tcp_port         => 8088,
			xmpp_id          => '',
			xmpp_passwd      => '',
			interface_name   => 'eth1',
			ip_address       => '127.0.0.9',
			ip_netmask       => '255.255.252.0',
			ip_gateway       => '127.0.0.9',
			ip6_address      => '127.0.0.9',
			ip6_gateway      => '127.0.0.9',
			interface_mtu    => 1500,
			rack             => 'RR 119.02',
			mgmt_ip_address  => '',
			mgmt_ip_netmask  => '',
			mgmt_ip_gateway  => '',
			ilo_ip_address   => '',
			ilo_ip_netmask   => '',
			ilo_ip_gateway   => '',
			ilo_username     => '',
			ilo_password     => '',
			router_host_name => '',
			router_port_name => '',
			type             => 31,
			status           => 2,
			profile          => 5,
			cachegroup       => 1,
			phys_location    => 1,
		},
	},
);

sub get_definition {
	my ( $self, $name ) = @_;
	return $definition_for{$name};
}

sub all_fixture_names {
	return keys %definition_for;
}

__PACKAGE__->meta->make_immutable;

1;