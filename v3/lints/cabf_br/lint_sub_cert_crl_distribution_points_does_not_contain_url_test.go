package cabf_br

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCrlNoUrl(t *testing.T) {
	inputPath := "subCrlDistNoURL.pem"
	expected := lint.Error
	out := test.TestLint("e_sub_cert_crl_distribution_points_does_not_contain_url", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCrlContainsUrl(t *testing.T) {
	inputPath := "subCrlDistURL.pem"
	expected := lint.Pass
	out := test.TestLint("e_sub_cert_crl_distribution_points_does_not_contain_url", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCrlContainsUrlInCompoundFullName(t *testing.T) {
	// Re: https://github.com/zmap/zlint/issues/223
	// Previously, we only grabbed the first entry in the fullName of each
	// DistributionPoint, whereas multiple names are allowed (these are
	// interpreted as different names for the same underlying CRL, i.e.
	// providing an LDAP URI and an HTTP URI -- see section 4.2.1.13 of
	// lint.RFC5280).
	inputPath := "subCrlDistURLInCompoundFullName.pem"
	expected := lint.Pass
	out := test.TestLint("e_sub_cert_crl_distribution_points_does_not_contain_url", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
