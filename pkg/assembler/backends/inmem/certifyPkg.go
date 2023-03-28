//
// Copyright 2023 The GUAC Authors.
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

package inmem

import (
	"context"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/exp/slices"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

type certifyPkgList []*certifyPkgStruct
type certifyPkgStruct struct {
	id            uint32
	pkgs          []uint32
	justification string
	origin        string
	collector     string
}

func (n *certifyPkgStruct) ID() uint32          { return n.id }
func (n *certifyPkgStruct) Neighbors() []uint32 { return n.pkgs }

func (n *certifyPkgStruct) BuildModelNode(c *demoClient) (model.Node, error) {
	return c.convCertifyPkg(n), nil
}

// func registerAllCertifyPkg(client *demoClient) error {

// 	// pkg:conan/openssl.org/openssl@3.0.3?user=bincrafters&channel=stable
// 	//	("conan", "openssl.org", "openssl", "3.0.3", "", "user=bincrafters", "channel=stable")

// 	selectedType := "conan"
// 	selectedNameSpace := "openssl.org"
// 	selectedName := "openssl"
// 	selectedVersion := "3.0.3"
// 	selectedSubPath := ""
// 	qualifierA := "bincrafters"
// 	qualifierB := "stable"
// 	selectedQualifiers := []*model.PackageQualifierSpec{{Key: "user", Value: &qualifierA}, {Key: "channel", Value: &qualifierB}}
// 	selectedPkgSpec := &model.PkgSpec{Type: &selectedType, Namespace: &selectedNameSpace, Name: &selectedName, Version: &selectedVersion, Subpath: &selectedSubPath, Qualifiers: selectedQualifiers}
// 	selectedPackage1, err := client.Packages(context.TODO(), selectedPkgSpec)
// 	if err != nil {
// 		return err
// 	}

// 	// pkg:conan/openssl@3.0.3
// 	//	("conan", "", "openssl", "3.0.3", "")
// 	selectedType = "conan"
// 	selectedNameSpace = ""
// 	selectedName = "openssl"
// 	selectedVersion = "3.0.3"
// 	selectedSubPath = ""
// 	selectedPkgSpec = &model.PkgSpec{Type: &selectedType, Namespace: &selectedNameSpace, Name: &selectedName, Version: &selectedVersion, Subpath: &selectedSubPath}
// 	selectedPackage2, err := client.Packages(context.TODO(), selectedPkgSpec)
// 	if err != nil {
// 		return err
// 	}
// 	client.registerCertifyPkg([]*model.Package{selectedPackage1[0], selectedPackage2[0]}, "these two opnessl packages are the same", "inmem backend", "inmem backend")

// 	// pkg:pypi/django@1.11.1
// 	// client.registerPackage("pypi", "", "django", "1.11.1", "")

// 	selectedType = "pypi"
// 	selectedNameSpace = ""
// 	selectedName = "django"
// 	selectedVersion = "1.11.1"
// 	selectedSubPath = ""
// 	selectedPkgSpec = &model.PkgSpec{Type: &selectedType, Namespace: &selectedNameSpace, Name: &selectedName, Version: &selectedVersion, Subpath: &selectedSubPath}
// 	selectedPackage3, err := client.Packages(context.TODO(), selectedPkgSpec)
// 	if err != nil {
// 		return err
// 	}

// 	// pkg:pypi/django@1.11.1#subpath
// 	// client.registerPackage("pypi", "", "django", "1.11.1", "subpath")

// 	selectedType = "pypi"
// 	selectedNameSpace = ""
// 	selectedName = "django"
// 	selectedVersion = "1.11.1"
// 	selectedSubPath = "subpath"
// 	selectedPkgSpec = &model.PkgSpec{Type: &selectedType, Namespace: &selectedNameSpace, Name: &selectedName, Version: &selectedVersion, Subpath: &selectedSubPath}
// 	selectedPackage4, err := client.Packages(context.TODO(), selectedPkgSpec)
// 	if err != nil {
// 		return err
// 	}
// 	client.registerCertifyPkg([]*model.Package{selectedPackage3[0], selectedPackage4[0]}, "these two pypi packages are the same", "inmem backend", "inmem backend")

// 	return nil
// }

// Ingest CertifyPkg

func (c *demoClient) convCertifyPkg(in *certifyPkgStruct) *model.CertifyPkg {
	out := &model.CertifyPkg{
		ID:            nodeID(in.id),
		Justification: in.justification,
		Origin:        in.origin,
		Collector:     in.collector,
	}
	for _, id := range in.pkgs {
		p, _ := c.buildPackageResponse(id, nil)
		out.Packages = append(out.Packages, p)
	}
	return out
}

func (c *demoClient) IngestCertifyPkg(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, certifyPkg model.CertifyPkgInputSpec) (*model.CertifyPkg, error) {
	var pmt model.MatchFlags
	pmt.Pkg = model.PkgMatchTypeSpecificVersion

	pIDs := make([]uint32, 0, 2)
	for _, pi := range []model.PkgInputSpec{pkg, depPkg} {
		pid, err := getPackageIDFromInput(c, pi, pmt)
		if err != nil {
			return nil, gqlerror.Errorf("IngestCertifyPkg :: %v", err)
		}
		pIDs = append(pIDs, pid)
	}
	slices.Sort(pIDs)

	ps := make([]*pkgVersionNode, 0, 2)
	for _, pID := range pIDs {
		p, _ := byID[*pkgVersionNode](pID, c)
		ps = append(ps, p)
	}

	for _, id := range ps[0].certifyPkgs {
		cp, _ := byID[*certifyPkgStruct](id, c)
		if slices.Equal(cp.pkgs, pIDs) &&
			cp.justification == certifyPkg.Justification &&
			cp.origin == certifyPkg.Origin &&
			cp.collector == certifyPkg.Collector {
			return c.convCertifyPkg(cp), nil
		}
	}

	cp := &certifyPkgStruct{
		id:            c.getNextID(),
		pkgs:          pIDs,
		justification: certifyPkg.Justification,
		origin:        certifyPkg.Origin,
		collector:     certifyPkg.Collector,
	}
	c.index[cp.id] = cp
	for _, p := range ps {
		p.setCertifyPkgs(cp.id)
	}
	c.certifyPkgs = append(c.certifyPkgs, cp)

	return c.convCertifyPkg(cp), nil
}

// Query CertifyPkg

func (c *demoClient) CertifyPkg(ctx context.Context, filter *model.CertifyPkgSpec) ([]*model.CertifyPkg, error) {
	funcName := "CertifyPkg"
	if filter.ID != nil {
		id64, err := strconv.ParseUint(*filter.ID, 10, 32)
		if err != nil {
			return nil, gqlerror.Errorf("%v :: invalid ID %s", funcName, err)
		}
		id := uint32(id64)
		link, err := byID[*certifyPkgStruct](id, c)
		if err != nil {
			// Not found
			return nil, nil
		}
		// If found by id, ignore rest of fields in spec and return as a match
		return []*model.CertifyPkg{c.convCertifyPkg(link)}, nil
	}

	var search []uint32
	foundOne := false
	for _, p := range filter.Packages {
		if !foundOne {
			exactPackage, err := c.exactPackageVersion(p)
			if err != nil {
				return nil, gqlerror.Errorf("%v :: %v", funcName, err)
			}
			if exactPackage != nil {
				search = append(search, exactPackage.certifyPkgs...)
				foundOne = true
				break
			}
		}
	}

	var out []*model.CertifyPkg
	if foundOne {
		for _, id := range search {
			link, err := byID[*certifyPkgStruct](id, c)
			if err != nil {
				return nil, gqlerror.Errorf("%v :: %v", funcName, err)
			}
			out, err = c.addCPIfMatch(out, filter, link)
			if err != nil {
				return nil, gqlerror.Errorf("%v :: %v", funcName, err)
			}
		}
	} else {
		for _, link := range c.certifyPkgs {
			var err error
			out, err = c.addCPIfMatch(out, filter, link)
			if err != nil {
				return nil, gqlerror.Errorf("%v :: %v", funcName, err)
			}
		}
	}
	return out, nil
}

func (c *demoClient) addCPIfMatch(out []*model.CertifyPkg,
	filter *model.CertifyPkgSpec, link *certifyPkgStruct) (
	[]*model.CertifyPkg, error) {

	if noMatch(filter.Justification, link.justification) ||
		noMatch(filter.Origin, link.origin) ||
		noMatch(filter.Collector, link.collector) {
		return out, nil
	}
	for _, ps := range filter.Packages {
		if ps == nil {
			continue
		}
		found := false
		for _, pid := range link.pkgs {
			p, err := c.buildPackageResponse(pid, ps)
			if err != nil {
				return nil, err
			}
			if p != nil {
				found = true
			}
		}
		if !found {
			return out, nil
		}
	}
	return append(out, c.convCertifyPkg(link)), nil
}