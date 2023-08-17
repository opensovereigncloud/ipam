// Copyright 2023 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"encoding/json"
	"net/netip"
)

func IPAddrFromString(ipString string) (*IPAddr, error) {
	ip, err := netip.ParseAddr(ipString)
	if err != nil {
		return nil, err
	}
	return &IPAddr{Net: ip}, nil
}

// +kubebuilder:validation:Type=string
type IPAddr struct {
	Net netip.Addr `json:"-"`
}

func (in IPAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(in.String())
}

func (in *IPAddr) UnmarshalJSON(b []byte) error {
	stringVal := string(b)
	if stringVal == "null" {
		return nil
	}
	if err := json.Unmarshal(b, &stringVal); err != nil {
		return err
	}
	pIP, err := netip.ParseAddr(stringVal)
	if err != nil {
		return err
	}
	in.Net = pIP
	return nil
}

func (in *IPAddr) String() string {
	return in.Net.String()
}

func (in *IPAddr) Equal(other *IPAddr) bool {
	return in.Net.Compare(other.Net) == 0
}

func (in *IPAddr) AsCidr() *CIDR {
	var cidrRange uint8 = 32
	if in.Net.Is6() {
		cidrRange = 128
	}

	ipNet := netip.PrefixFrom(in.Net, int(cidrRange))
	return CIDRFromNet(ipNet)
}

// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddr) DeepCopyInto(out *IPAddr) {
	if in != nil {
		out.Net = in.Net
	}
}
