// Copyright Project Contour Authors
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

//go:build e2e
// +build e2e

package gateway

import (
	"context"
	"crypto/tls"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	"github.com/projectcontour/contour/internal/gatewayapi"
	"github.com/projectcontour/contour/test/e2e"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gatewayapi_v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func testMultipleHTTPSListeners(namespace string) {
	Specify("Multiple HTTPS listeners work", func() {
		t := f.T()

		// Set up a backend and HTTPRoute for each listener.
		for _, tc := range []string{"1", "2", "3"} {
			f.Fixtures.Echo.Deploy(namespace, "echo-"+tc)

			route := &gatewayapi_v1beta1.HTTPRoute{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      "httproute-" + tc,
				},
				Spec: gatewayapi_v1beta1.HTTPRouteSpec{
					CommonRouteSpec: gatewayapi_v1beta1.CommonRouteSpec{
						ParentRefs: []gatewayapi_v1beta1.ParentReference{
							gatewayapi.GatewayListenerParentRef("", "multiple-https-listeners", "https-"+tc),
						},
					},
					Rules: []gatewayapi_v1beta1.HTTPRouteRule{
						{
							Matches:     gatewayapi.HTTPRouteMatch(gatewayapi_v1beta1.PathMatchPathPrefix, "/"),
							BackendRefs: gatewayapi.HTTPBackendRef("echo-"+tc, 80, 1),
						},
					},
				},
			}
			_, ok := f.CreateHTTPRouteAndWaitFor(route, httpRouteAccepted)
			require.True(t, ok, "expected HTTPRoute to be accepted")
		}

		// Make requests to each listener hostname and validate the response
		// and upstream service.
		for _, tc := range []string{"1", "2", "3"} {
			certSecret := &corev1.Secret{}
			key := client.ObjectKey{Namespace: namespace, Name: "tlscert-" + tc}
			require.NoError(t, f.Client.Get(context.Background(), key, certSecret))

			res, ok := f.HTTP.SecureRequestUntil(&e2e.HTTPSRequestOpts{
				Host: fmt.Sprintf("https-%s.gateway.projectcontour.io", tc),
				TLSConfigOpts: []func(*tls.Config){
					// Verify the server cert to ensure we're getting
					// the right one.
					e2e.VerifyTLSServerCert(certSecret.Data["ca.crt"]),
				},
				Condition: e2e.HasStatusCode(200),
			})
			require.NotNil(t, res, "request never succeeded")
			require.Truef(t, ok, "expected 200 response code, got %d", res.StatusCode)
			require.Equal(t, "echo-"+tc, f.GetEchoResponseBody(res.Body).Service)
		}
	})
}
