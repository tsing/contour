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

package fixture

// nolint:revive
const (
	// CERTIFICATE generated by
	// openssl genrsa -out example-key.pem 2048
	// openssl req -new -x509 -days 18250 -key example-key.pem -sha256 -subj "/CN=www.example.com" -out example.pem
	CERTIFICATE = `-----BEGIN CERTIFICATE-----
MIIDFzCCAf+gAwIBAgIUZULFakfIJl0qaJXAVPCz2nzvB38wDQYJKoZIhvcNAQEL
BQAwGjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMCAXDTIyMDgxOTExMDkxNVoY
DzIwNzIwODA2MTEwOTE1WjAaMRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20wggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDS9S4d/ea6wqiib8UeyHMptoks
w+q2DNuF75NQHLh5Z2rUnE/N8/KVhpIx81QdId1maWS0b3392hCRRFY3sDlMpRk/
1uoQgzLdk8pjw1JiqoDpvTiKZsADmVuUcCdHLNEzYtcLWBv0VyyNyE5pdrnVnMbx
w1aiQ8w2lcBCJQ8Y4DAc5oKlvBu49aUsvfFZwjL6Cr1qafQiYylqQcz7zqGBYjXc
iMzN+4fE1XQlw1iy6XmVZiHQr8Sb7EBI+g0iJapgNv7tBunzywSvAYK8N42QQOll
1sKEVf7thoNEmJTIUFo6m57Fys7LQ/B8in5JwBU+1FjqNWLJ1Gj+zIc93oc3AgMB
AAGjUzBRMB0GA1UdDgQWBBS/BZ2Uu1Y0//Um8bOqyyWz9LnPvzAfBgNVHSMEGDAW
gBS/BZ2Uu1Y0//Um8bOqyyWz9LnPvzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
DQEBCwUAA4IBAQCTZ6ZDi7aU7NjZdGWNLrRCBEt+FcD+mdvtRcaSp2K7m+WObnWl
rDM7V/s8ziu8ffwfwEbaBKYVLO7Mww8ke0WclBp1sq6A5AWy1sBCQYJuPCdOJNY0
fLaZObhUSQvNGw1wAXkgczrsOa/5QII356UsLiqhninXWYTMvNehab4+QW6Dldqo
EyxKgX2Ls984ZN5CDvvXfRnkeQW1/K705ReZq8qmtmCwU5wHYy0IoJGNapeX45VY
6s2n5I5CpH4L9Ua4NLgqphjC/QYK4q71GHTZD89mfTsmE+0flgFDS+wrv5SusK8u
CY2iW9j8VptZU8LVs9FrhgecEtfXbTA3MeSo
-----END CERTIFICATE-----`

	CERTIFICATE_WITH_TEXT = CERTIFICATE + "\t\r\n"

	RSA_PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA0vUuHf3musKoom/FHshzKbaJLMPqtgzbhe+TUBy4eWdq1JxP
zfPylYaSMfNUHSHdZmlktG99/doQkURWN7A5TKUZP9bqEIMy3ZPKY8NSYqqA6b04
imbAA5lblHAnRyzRM2LXC1gb9FcsjchOaXa51ZzG8cNWokPMNpXAQiUPGOAwHOaC
pbwbuPWlLL3xWcIy+gq9amn0ImMpakHM+86hgWI13IjMzfuHxNV0JcNYsul5lWYh
0K/Em+xASPoNIiWqYDb+7Qbp88sErwGCvDeNkEDpZdbChFX+7YaDRJiUyFBaOpue
xcrOy0PwfIp+ScAVPtRY6jViydRo/syHPd6HNwIDAQABAoIBAEf36RHGSu6v9gPk
iaUk0VULtuSUuf/9hu68es873Rtd0q5R3U/vx3SHglyUHMALi5Kipf6AgsUVnc1R
OPCqqAGj2WdUFGopuDKrdsJuIi8S6APVz/I3d45CxWFwmZXIjl4vfBmcp3zGOKbu
DQIhxOhBIgXclDOrWYHNuNdX+TyMr9cBe9bPtsWmN6Wl8V26kLzboQ3bIsM9taL+
sizVNrLL4U7oqbsKVhmI1m78+VaucQeK+XuyvuT6toOxY/Mgidcp4Mwq06aOTFug
N/u+VyXNen3Vk6/kMTxbaM9aQ05cnOrKqrgJAhiFGb+lcqvvGlwGxNqGcBABYkqf
ejhNdEECgYEA/3shOSl4GmEGZZXTARo0bDnelSpSJgmlStBeAGraq+lm/4IrTCCB
2tmJkEs6uiCEH7vkbbI/lLlhNjM2gKDovP3UwTQHGfiHlcIYYICMIpbp7+Ar9+ex
4KPXmgqqhTKC4xQYGoUE6INlpZcQ4blA4AtQbhgcC4Cp8FD/QbXp0xECgYEA02Ll
EAxzHZBNoK/5oPL0LiDUW/zPdOWCQCW8oGW5gDUvCLQ/lntegL8Qzubt7PZ4xgeg
m2ENTDcp1Zfn9s0T1V2T9Sba8gShUCvm9nLVCj4OQ3lwDufwHFuhKFpj1BVnHD5y
9yhXfyrFgvhamepEjq6LZUCrL1HxZqgOezv6JccCgYBqM1oFNArcFFcfZV+YRrdi
AdBX+4a4jyvp5KIe1ExgSB7rucWb2KuCOQmpNMyN0LR7qJR1UTKC9WjGqhVO9RSq
c228fo8xKZHbHBscCnO2cTt/3pUIcYUM167pNuPZiLzF/nVimMcIjI51fk2jN2oT
eECP82+9DFgYMONbAm7XsQKBgH/Y3ztenDz0Ks8Vv3/FkUNY3bco5vwHV0ieyj+k
ZpYRFHpKMe88fEKXzH2mk53uz8rNkCiJgTZoYqfpcQUGsYkpSLRLpL4daMcJVm4V
s523PH84sjqBsuojzQuP57K8oxkk9/ld79VctAprVLikRISbMnmxrBc5kywIVoHY
G4m/AoGBAPV0wuytURR3CbPHz28wrKRh/xnHrW+Fp3ooRfj8Pr/4zDVsqqNz2gA3
yVb6kAEpg2ON9NOWSfoUfC+THitOnRm8pKL1QL7oiq/2+s3IiK+jSevEU7TUfjio
1LwtUqv1MbKdv7TgkU0YQ99iLocWF4F4oWF6AX86/BL9y7gcbE0y
-----END RSA PRIVATE KEY-----`

	// EC_CERTIFICATE sample elliptical curve data generated
	// openssl ecparam -name prime256v1 -genkey -out ec_key.pem
	// openssl req -new -x509 -key ec_key.pem -out ec_crt.pem -days 3650 \
	//     -subj "/C=US/ST=CA/O=Acme" \
	//     -reqexts SAN -extensions SAN -config <(ca t /etc/ssl/openssl.cnf \
	//     <(printf "\n[SAN]\nsubjectAltName=DNS:example.com,DNS:www.example.com"))
	EC_CERTIFICATE = `-----BEGIN CERTIFICATE-----
MIIBfzCCASWgAwIBAgIUZ8EBxJShrhAiO9bG6aRVcJdlEJowCgYIKoZIzj0EAwIw
KTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMQ0wCwYDVQQKDARBY21lMB4XDTE5
MTIwNTAxNTg0NFoXDTI5MTIwMjAxNTg0NFowKTELMAkGA1UEBhMCVVMxCzAJBgNV
BAgMAkNBMQ0wCwYDVQQKDARBY21lMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
zCdqvU5dSKxzDAVakEi97epIazdkUKRT2XZtUk41Hp2H4xy8EzR1Re3r9AdJRsJn
sGrHGbIg2r7OUNYgeN4ot6MrMCkwJwYDVR0RBCAwHoILZXhhbXBsZS5jb22CD3d3
dy5leGFtcGxlLmNvbTAKBggqhkjOPQQDAgNIADBFAiAYFlD2n/uWWxTqi8WcWvb1
CUDxSzF2/jLe1PIFkwNk7wIhAP9kMCO1ys050JNvlVZg3xvPvCHKCkWcSachE5fC
5hc6
-----END CERTIFICATE-----`

	EC_PRIVATE_KEY = `-----BEGIN EC PARAMETERS-----
BggqhkjOPQMBBw==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAM3LdZrzZk8Hn4VqBDNTgOuh9E772M4sgEYvZMNOy4moAoGCCqGSM49
AwEHoUQDQgAEzCdqvU5dSKxzDAVakEi97epIazdkUKRT2XZtUk41Hp2H4xy8EzR1
Re3r9AdJRsJnsGrHGbIg2r7OUNYgeN4otw==
-----END EC PRIVATE KEY-----`

	EC_PARAMETERS = `-----BEGIN EC PARAMETERS-----
BggqhkjOPQMBBw==
-----END EC PARAMETERS-----`

	// MISSING_CN_CERT issue #1965
	// certificate with no CN field.
	// openssl req -new -newkey rsa:2048 -days 365 -nodes -subj="/DC=com/DC=domain/DC=my" -x509 -keyout server.key -out server.crt -sha256
	MISSING_CN_CERT = `-----BEGIN CERTIFICATE-----
MIIDYzCCAkugAwIBAgIUcNI/oD/y3dZ2Rmyvz9Xb4BsC0nswDQYJKoZIhvcNAQEL
BQAwQTETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBmRvbWFp
bjESMBAGCgmSJomT8ixkARkWAm15MB4XDTE5MTIwMzA0MDYyM1oXDTIwMTIwMjA0
MDYyM1owQTETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBmRv
bWFpbjESMBAGCgmSJomT8ixkARkWAm15MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAxblHSlzhyDY88rhILpn+rCALqK1ELcbL5T3FygLEdUlE/da3VpSz
wBaZc9UIhFpwnkvUy8eKYbLy9jIk3C0aJasrX828MkF4lInSJf0BREbddxmwDovl
KZoo7xz2rY6LPTopZ1GdszzsPxvBFvc1gWl0UEXkxDAZjtnDuVB+hBv6RR7BpSdN
Fxds3OVles1mgasW79gMCb22gFa9vJKkxVJz5IIstrShjCfNvPvULf3aeMJjy8fM
x8kHBiRSs6HC4dFR3cR1uuZCdnkeR76X3gAn9A18VBMvA25JehhfLYnVJi80UgGV
PAxKgjU8dw8auvbmopzkguyFTfW6sYfDBQIDAQABo1MwUTAdBgNVHQ4EFgQUITS5
7JyH/S2wUrZnUjQXAp1nTPQwHwYDVR0jBBgwFoAUITS57JyH/S2wUrZnUjQXAp1n
TPQwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAn5nxNyI6+3MD
X5hzg5G0gTZafL8qN5bJGlFRWeyU2CvWsXbHCYtzADkwGCcCyoA1nmb460kwZ4vk
fOz9FYqZoh4IIegs5bB559Ze26Kjl5AgAFARXxV6s0cPDQ2O1XFkNIQiLFHHPfRR
RJtQgNaLppCe8TFNPjUDqNfoWPQvnlEYMYsvKto7pNw+HGreXhGF5CUCGIqT747P
zYmBGZEE2q1L45nErZFK4d74XzGu/4Kc73zEFS6GT71Zu6Ec6wUCKdciyNtFHFmh
4l3l5YC+1zK6ZeOBxdoMKpldD9EV4GQqW8aE2Adpd5RcUUQJnqWM+/ysod/AnHL/
fm35MjBa+w==
-----END CERTIFICATE-----`
	MISSING_CN_KEY = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDFuUdKXOHINjzy
uEgumf6sIAuorUQtxsvlPcXKAsR1SUT91rdWlLPAFplz1QiEWnCeS9TLx4phsvL2
MiTcLRolqytfzbwyQXiUidIl/QFERt13GbAOi+UpmijvHPatjos9OilnUZ2zPOw/
G8EW9zWBaXRQReTEMBmO2cO5UH6EG/pFHsGlJ00XF2zc5WV6zWaBqxbv2AwJvbaA
Vr28kqTFUnPkgiy2tKGMJ828+9Qt/dp4wmPLx8zHyQcGJFKzocLh0VHdxHW65kJ2
eR5HvpfeACf0DXxUEy8Dbkl6GF8tidUmLzRSAZU8DEqCNTx3Dxq69uainOSC7IVN
9bqxh8MFAgMBAAECggEAZVvBjAFpTPVg8Rw/BIS05Q/YajPIS58pTu8HwbMBew6A
a4/ylFPOgZ4UNCj1IMQsDznYcE5uRf8yRbsW4jfMu5qvtYEGBM1DPwIX1JmKWLHr
Pe7RLePRKi545Xr3iakU/+Ic73YLXaLRiNh1d4xqxViF49CwoVH3CB/iEdGNybKX
CyWxNUS9Tguzg26zIqvtjjqzP3ULt/P+fi/NTIKswv4Btp6hI8QDA2RL6RFJYIL2
jVpRFUSOhl5N7fg8ZKIAaQFetenzM/yF29qa6nlIhXFsep6bRaLqw9Hni2YB6yCh
7z4I/ygSTirPW1yk+NPiLjc6mTZVf7EvdFeApOOSUQKBgQDq71Hb0hciDyFkdq3R
nxskJeiSk8b141tarlrUTvoPRYwOaeKcDN2Lzj7Y9Pqxllb8CSjoOR0B/pfwvfxI
J2HGoNf7PK3rn166QiFPPFxwa5bMpCcbEoljToB4nCRDsYRlUrd/hAHAb1ptLqK2
OHc2pbOo0mmYORstIFccrKJDYwKBgQDXc9FA78KjZ7iS2GCs3yXOdS2U1TD5hZMr
MDXMA0CMbxTYmLRLD0x41hXzH/2UsfQNCYQZCOqWhW3FJHqAaBz0w2n+dOMkWeSq
EI5ghDQyu/HM2V7xk3tYYXhZ/FxOzcd2l/DUzKZsPJuxsXsNaUn9+ZOE44xHRjNf
wYvNv17QdwKBgQCujZO/hLAlYSKJV1g8OD/dMsFDLsMT7JHypTrdJbTLZfvytZ9m
HHT7LAkr/5DII5CLgG7BY7X2xmezuiTYo1IVV2pBw8rhFy81qm6/RXTVHksTzx8z
ESm8/BWeBz02go2BDt1BxB3dEZ8ZIh5Iz1lb4+/BjlxgeoWDmNTAfE+vSwKBgFGq
I7HSb1tSsEJw48wC1Si5f6p/WI3r1Im1P17yCKByZltnHkepJ9pRg4ZhJNQc052x
crGukIS3VJE6L3jGfdtEysNZeNNJg4P2vJDW65YjaRa1eehld4Zbg6vQHQj9tNI9
61otrBMwse8bj8HYm+Q5mnHvcjd943EzQpOdKwonAoGBAKGymiC7RgrdaxbKf7oE
dsipT/w+MFInBLLiJcqLAfh0647vhjCFNGG4MZ6YGVwiIdvjZ1dH2Ujsdo8e0+Wi
8cnx8JvuAJdr5HzMI6fvnMDzjzAskMgYUNhOUhM2g223JuoyyLY2/DL7dOYkFeSn
b5qYn0JNERfPYdLwXNV1HCM9
-----END PRIVATE KEY-----`
	// Contour CA Cert and key
	// Generated using:
	// openssl req -x509 -new -nodes -keyout certs/CAkey.pem -sha256 -days 1825 -out certs/CAcert.pem -subj "/O=Project Contour/CN=Contour CA"
	CA_CERT = `-----BEGIN CERTIFICATE-----
MIIDOTCCAiGgAwIBAgIUZ8+/J2epHC3MvTNnwjElp4/nSzAwDQYJKoZIhvcNAQEL
BQAwLDEYMBYGA1UECgwPUHJvamVjdCBDb250b3VyMRAwDgYDVQQDDAdjb250b3Vy
MB4XDTIxMTEwODEwMjY0M1oXDTI2MTEwNzEwMjY0M1owLDEYMBYGA1UECgwPUHJv
amVjdCBDb250b3VyMRAwDgYDVQQDDAdjb250b3VyMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAmh4d9TeixoYNP2L8byVmfsuHi7e9DNdy9qCoVPfVmpoI
E95FWxgPmA9b1oAvVUOdR57G+LQLdVQoRTvypXPkibV/S3GHTfKftPTLX26dGoVy
RM0imQfYqtHdqt7WLPZZ2shwUgkRu7P7N4S8YicyFz4jCwMa1VzEL2+bP5SJFnKo
rLacn+/r+PFEuTwACIoBTH85pbF03DnC8CA8KlHYr3FqNks41zjrSvGbtR6vb8++
WUAlS5+9e7uXdTE72mLixyJzx84v9mFEGjHehHSsWzpzyTOr54umLL7hNtZVwe3q
YI8yC7//dWiSW7gurWZy6dtZ/vJxGEE9YMtEc0KEHQIDAQABo1MwUTAdBgNVHQ4E
FgQUdMZ9tM59+GSQioDV6oju/yuBYIkwHwYDVR0jBBgwFoAUdMZ9tM59+GSQioDV
6oju/yuBYIkwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAJTBN
7CvK6rhDkPbBdS8vKakpJiK3X9QASXDKdsDLP3QGqkNJvDxHNes2vMJxkrMzMfDR
v2Hlfa9rHoPzgz3w1D7jCpKP9bldwk2WxFcZRFW/EF1PXxD6cS3+FxLaDEvs5PIb
AGKJQauk9OiLcqoMFzeTnEMjfnFiLSZxVOIETak0ChrIdrtd4nQXAYNQ0PSX2cg+
2nnF9tHUufzQ9rYDovx8AKa00p0nsP/afxNCTDYQdyyhgf/uYo6qm2dneFDzPLcc
QAfYWEVsJAZ2VW578dO2/wbu2Ey3Zy71xu/7AfY6O8eqLng/q6nqOwS1AGc24+Yx
kB6w0LNcvz2KNbJBRQ==
-----END CERTIFICATE-----`
	CA_KEY = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCaHh31N6LGhg0/
YvxvJWZ+y4eLt70M13L2oKhU99WamggT3kVbGA+YD1vWgC9VQ51Hnsb4tAt1VChF
O/Klc+SJtX9LcYdN8p+09Mtfbp0ahXJEzSKZB9iq0d2q3tYs9lnayHBSCRG7s/s3
hLxiJzIXPiMLAxrVXMQvb5s/lIkWcqistpyf7+v48US5PAAIigFMfzmlsXTcOcLw
IDwqUdivcWo2SzjXOOtK8Zu1Hq9vz75ZQCVLn717u5d1MTvaYuLHInPHzi/2YUQa
Md6EdKxbOnPJM6vni6YsvuE21lXB7epgjzILv/91aJJbuC6tZnLp21n+8nEYQT1g
y0RzQoQdAgMBAAECggEAJ1oVUC3alFJaQ5sk1cR4/Vs+ywApd4PzyOGQ2dFGa2l1
Mo2IrA/iu3MAgi7M5zqSiF0UdSCT5GuOtM7v0CtdaYQ8cqkUYWTYPr5ax41Y9BKq
8MJoWSbVvhhEP4MSsPxeSSzw8SF7s6/rL2vnJjDX63NkQ0WRrnU3at1WiM+GiYJr
74VVwekqswo4U2G37hHtrnZEN3Poer6awY/sGPCfttJxXTzkNUeW9SlUxI/3ItGx
NlRB8J9p1aO/1zLZ+IguOuPz5cz+IBIAxKzmRPOVPGU+IuE2RCYPiZ3ki4qSqUui
7iyE5YVmzVkwjJRSPqyGH05YQkpQ+3FAbE6DrekuwQKBgQDMwtcNsxQboLauU5sz
P0qlvULi+KbL7NgN093vvkmGlUmdyW0/rjAw2Y+CBbAuczWadcID4AVsEbtWCD5r
3umdLxhmz3Ax8iI6qdOz8vpLYGVKsr+WrlT1cb1OHJmqkjE0WA4kcgz2UnGogwAw
oQRToCV2IVyCJy2dIUGU4Dd9NQKBgQDArwRYoM/+isKRSyGHjLrhC6wJHk6855Ah
PEh0H3gj+KZnTx+ca4Xs0tvgA0EqJGU1PIlLSxa3SLAZ2ZXvOwz6xPO0B75XZymT
cumUX3SINGhh15Z8fv++eSQN5PtvILWzd01UL4Z7fpH6P/AISG6nkR2G18N40A9A
s9IEEzqQSQKBgDDPCQBE/bgQOTuYxlgS8DJNyRHdf7UAggfZxv+M9m3UTDYGe7pA
RVB/q8usm3CXGeH1zAescN5v454AXRDNlBZZ0ZFARVhChZJCRSqR8sPg+IQz06OR
bsMLF7r06lRUgNIfY0+guCJHSLyIbZUp7KQaepOMJEcKG79m3AEm6VM5AoGBALpd
jXqbsBGm7KHTZU7SEz7wJyl6ovWbk5BR2lJWzcEdEvYBVbXZWePB3uFAxJqaMTnY
n4Kv0fzc1VdrWNXW66rxLDYI91VadqAjGUFlQjUrW3qnJuqre3kjenfl7juCC7zV
u3mdrhlsRo0homyggkaY2VI/BtWnh9kxFATVBGPhAoGBAMT78LKXMnPzKPo20Qgn
kaVf1mn5ZhtHUXVfiaSHmJOHDgPkT2FxssrV3lyrTgHPw68L6gBGh6Y8HZY/Ej82
/ZXgsGHEgo0+7p3F24fqemBZlqW4nstDhXu9kQJJJwb0eEUMEl+UvzjH3lBByQkY
hnHvbOWaTdW3tkGPgHjDi7S8
-----END PRIVATE KEY-----`
	// Contour CA Cert and key, without a CN
	// Generated using:
	// openssl req -x509 -new -nodes -keyout certs/CAkey.pem -sha256 -days 1825 -out certs/CAcert.pem -subj "/O=Project Contour"
	CA_CERT_NO_CN = `-----BEGIN CERTIFICATE-----
MIIDFTCCAf2gAwIBAgIUHDO4xZB226JzfA3Jy2uoKXyjAvYwDQYJKoZIhvcNAQEL
BQAwGjEYMBYGA1UECgwPUHJvamVjdCBDb250b3VyMB4XDTIxMTEwODA0MDc0M1oX
DTI2MTEwNzA0MDc0M1owGjEYMBYGA1UECgwPUHJvamVjdCBDb250b3VyMIIBIjAN
BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsmM4yuIn9zfJUaBJXkti7HmfPp2D
nBf8d40eo/koSC3lWi2qwSpYs3OxgDOzQf14plPKEETuKgbsXk/wIc/HP45I0jih
iiN3npEbYWBY++6XYTYTXPtR9D/Z8BcqzlTZpPlBbqz2ZXL7s87Kl0bi/z7LBb9i
1ttae64l8qd1DBHR1aHyzqIPER709pIwcQfPWaVgavin6z6PAhTG5BzutuI4zOlo
hrFfMA+cBtFUg7NveFkRyQZXKWKYrxO/8iRQkQ5WElx+K5Cty0zC60MG4wU8e8jK
Z2xGxgDspCx/YVQ7hKNpci79jcPguO1auxXQfYlRC9KO1F8IuA2+P9vPvQIDAQAB
o1MwUTAdBgNVHQ4EFgQU88BbxgMIfwrghVgV81OSmjGrG+wwHwYDVR0jBBgwFoAU
88BbxgMIfwrghVgV81OSmjGrG+wwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0B
AQsFAAOCAQEANFZRv6ojPVTsbTk2hVyXN45pkrzZfcqahpUEAt/8Xf3AsL0Vp8pm
7g8p8/+HIXgy6Gx6yenRdN4BVQxFNwjMYItm/gqmBWneQh+mc7CpgXB35o/BVIo9
3DdugX8oMsJrcQicVGrUn3K2QlZNHDvrV3nfjzOKYn8vWOC5onvLi6dQi54MPeou
hTB5vgr0cSRBj2WxWgetoIaLaba+qd4W0+G9Ktbjit8v2o4tkoskFWQ5Txu9C8gI
Hn1un+00VZOPKjK+M94R//baEQPtqyiOPSF+9HBBOAAMVgpJiPtbEM7+w4vnJ5RS
CO3wUlg5gKd07zXF216MXaLL4nwmIUCSSA==
-----END CERTIFICATE-----`
	CA_KEY_NO_CN = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCyYzjK4if3N8lR
oEleS2LseZ8+nYOcF/x3jR6j+ShILeVaLarBKlizc7GAM7NB/XimU8oQRO4qBuxe
T/Ahz8c/jkjSOKGKI3eekRthYFj77pdhNhNc+1H0P9nwFyrOVNmk+UFurPZlcvuz
zsqXRuL/PssFv2LW21p7riXyp3UMEdHVofLOog8RHvT2kjBxB89ZpWBq+KfrPo8C
FMbkHO624jjM6WiGsV8wD5wG0VSDs294WRHJBlcpYpivE7/yJFCRDlYSXH4rkK3L
TMLrQwbjBTx7yMpnbEbGAOykLH9hVDuEo2lyLv2Nw+C47Vq7FdB9iVEL0o7UXwi4
Db4/28+9AgMBAAECggEAIpDdYVq/OcUiimGPtejPq1jJxZ1J8kYnkpbSuB1Ac/QI
vBebPcDR3TvuyhO6iW/vH/OedNaWe6hJxuCgfDzBtAwDClEU5CqJND2T26OKuo55
4MlUqTC9qaFxvncOwVpprfDqf+Kd//xuz2GjmfqPY7wsJNn38gAhIVQWYFWYdZTv
iUQT6jnpyisGpkzjwa3STTTRFz9yUyhTMAYYJ79JFVm3jbh07JlZ3Fueg8jGYdP+
hHTG4BG4C28UAvhle2YbCU/W2emVtt6hXAqkXR7n21hmiGJDzQOrAXuZ2ASBS1CQ
oX1YiOMnt1SWLl+GhiSNsTg6400m7DmpMXMNhfanQQKBgQDaIbegY27XbQiH34Ym
09QyzuwCaelUl52GLLL7jNoPzIjvUQIq69V5HJe9ajh2FBz/IbqxGz9DXmzKZJY/
8Po+t6YZQbwdH0dPt1+Zda3I/HTb/IQ94NIoyVJ7H17WDGD2nF62TUHP07hH6Rr5
pM/UYDJWFRsKcXPIc1AEwiCurQKBgQDRWy7RnchakzKxXLneOGUFkfBLb7PAYAOG
VjurfiOgq+uQLEt/tJOPCPnRR0czfrgyw3OX6vLMz0WQ91NrPq2O/QK2lKokw2Jr
q/InOfeZLQPFGGFFmb5iBVZsH7LX9qD7H78gFELc1rm9pKrJ9g7AluaH57zlJG8p
3pw9rH8XUQKBgQDY79OJxZoE0cE8MEdg8icF6Ni7hc2yfZ8CRA41Mt3j/IzrRWuc
eDM0YgVfMfa5KEO2UTs0zF5ch9o2Z1cufGDEYyzjzYZiiCsJ0ttW6bxOORWXe0mh
qKgaPC61mHte8ay1VsqzD562LGAH1IMpaInxM0Kcbh0Yg8CcyGH6eX67eQKBgEdC
VR5OKVsDDVcna7CC24fY08SpgbJyO+DmyyqlJbVTrzuEiDIwoaUm9flQ3KBfeoz0
0AyekQQ7bL9aH8+6JtMrtVe2JG+c/YwOm77UUT2W/9h/YuSQ3yd4D6F/7JzruoGf
natqmhOXFuPA/8z8bqsI8fDCZR5Nl3TV67JmW1/xAoGBAJPJ/NRMSEjlWxAb/P5m
JSpHjmcnw5+ozB0DPRmoSHw5P5junVyrWAw7GepIWMTw6I1Hfp+baP0aW0LHPquT
8rPqgpTohfWpl/Tl3mfi9umjJYL6ZCl955s0YhnijeZeBe/Uaz8HHa3JuVyGWRTs
6AEtsFGX9Fm0SE7tQc7JwlKv
-----END PRIVATE KEY-----`
	// Wildcard certificate and key
	// See #3496
	// Generated using Contour CA
	// openssl genrsa -out certs/wildcardkey.pem 2048
	// openssl req -new -key certs/wildcardkey.pem -out certs/wildcard.csr -subj "/O=Example/CN=*.example.com"
	// openssl x509 -req -in certs/wildcard.csr -CA certs/CAcert.pem -CAkey certs/CAkey.pem -CAcreateserial -out certs/wildcardcert.pem -days 1825 -sha256 -extfile certs/wildcard.ext
	WILDCARD_CERT = `-----BEGIN CERTIFICATE-----
MIIDSDCCAjCgAwIBAgIUGNo2zuxwkgG4niQkxt+F3w7rd2AwDQYJKoZIhvcNAQEL
BQAwLDEYMBYGA1UECgwPUHJvamVjdCBDb250b3VyMRAwDgYDVQQDDAdjb250b3Vy
MB4XDTIxMTEwODEwNTU1MVoXDTI2MTEwNzEwNTU1MVowKjEQMA4GA1UECgwHRXhh
bXBsZTEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBALermKPmS+cytKbiIrwNuSOrpSzja8twZWHXQ7TuM/x2lVlz
dypNtkV6oKkEKHkkUQsrHFqrNf3oWSIMwc5XNUYIid0p1GOCN2nDMWEleBa+4ijX
5qlMrtOj0TS23wHShlAn5ppXbPiUeGwvMUnRQgW4mqQUsN69YGj0wi83uwVHP1mQ
UlExFdbIAwlVyMF+bqo/nRxzrdCcGYzMfVg5ux5VZipzrLb3MjsluqBN1cgfP4Eg
N4XeuHpPcsWEuRFuGkxyGbrLKwDe2aq/CXynGYDt3UVuDJXZuhpMuWUEdoJEhpql
XDA6VjFKbYRaeEbN0Tiwo3mwXwXdLek+G0G2e3ECAwEAAaNkMGIwHwYDVR0jBBgw
FoAUdMZ9tM59+GSQioDV6oju/yuBYIkwCQYDVR0TBAIwADALBgNVHQ8EBAMCBPAw
JwYDVR0RBCAwHoINKi5leGFtcGxlLm9yZ4INKi5leGFtcGxlLm5ldDANBgkqhkiG
9w0BAQsFAAOCAQEATgOChe4EbQsSSuEFpOvTgSO1aK3cQcMnR3lk3y/JOjRwoqf3
/KTm5E/9vWIBRjuVnUR0l/zIvSsXSsbMqf/7sR4X6YKfmLhAwJylUhvkCKHf/IGD
B+YRd1cSgegRwXsweOQA2lDHC9UCQ5YO44Cotj1z57OZrw0u1MO/BKpTKO4W7D8K
JRx5kCJol7oWe0BcnRT77gM3Zs0P/9WwAweWZX4E6IiLuxs2i/hI9RDak2IfqKHn
mUIzGeroN/YNwxHDDr6NUuRWSR05hwAIGACxbcLgbLkQVNe9i5UanE9S1cZkkCSJ
bg9coUE54ebAFjaWagQjTfds+NXBrfoFCXwtHQ==
-----END CERTIFICATE-----`
	WILDCARD_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAt6uYo+ZL5zK0puIivA25I6ulLONry3BlYddDtO4z/HaVWXN3
Kk22RXqgqQQoeSRRCyscWqs1/ehZIgzBzlc1RgiJ3SnUY4I3acMxYSV4Fr7iKNfm
qUyu06PRNLbfAdKGUCfmmlds+JR4bC8xSdFCBbiapBSw3r1gaPTCLze7BUc/WZBS
UTEV1sgDCVXIwX5uqj+dHHOt0JwZjMx9WDm7HlVmKnOstvcyOyW6oE3VyB8/gSA3
hd64ek9yxYS5EW4aTHIZussrAN7Zqr8JfKcZgO3dRW4Mldm6Gky5ZQR2gkSGmqVc
MDpWMUpthFp4Rs3ROLCjebBfBd0t6T4bQbZ7cQIDAQABAoIBAQCUWY8yYuLP3M65
NFGl3te5Dfyd4xq/+QqKGlrfmr52njlHRNHtY1NjjwaRPYCLSGAFS07dlQhJCxBj
RjaLuTiYVnE0J4Ma1KR5UmVh9+elR4xYPPpbqEpVOy3RDBvX1vkRaQkDhLXJmfIi
P4PqAZD4GgEdonCxl/h1DayCgtsWCPTrMCcVCib3m8CuehBmyc5MgK6psgm0WSJs
H5+9RZ5A7eJ9zdkxC03ulflZ6WHNR/ojN8+fyogX6O7n/xLLDT+Q5kUv1QHN6cbz
QvNiUI3CfN7koiVbEu8IhCj6+/rYHw3QLow1MQFLLD0zLPRHhhCAavgtzx9lpR8r
lMdYCt+5AoGBAOhmVj6c6hq2zQ3v3wi7lQvIDU8chvEEuS9CmyhcCBrn+Vz6vXcl
kCuqbCGqwjxHELN0qeACtqh/f0M1ykSyDFDYf6BwJRx6sz/l61VzaCCtvj/drLFZ
aNNGwyBma9m7RKbQXfPrcdVdPMd5w6rYkJAOw8FwNkH5MyHEvvlvzSyzAoGBAMpS
clmNfCGBmb+IUN5e++dYjZ3LnOd6nVjvrY444bloBtgPyAzQUbPxrBDqxr90nAdF
AeulvBIEu287Aff5ltZRB8+MvUgvxzxFK7qWgJlX6uesLEj6dFDFYYb4IVyoiqKT
vAGnvGQZsm+UKzp5pqyVVxLXHbd3Sl+hquyjE5FLAoGBAKwGLGrWfddeXrSlTQPK
zk3VPDzGEgDQDT+1XuJmFT3NXmD73UkXjfs0gCv/mR1DQVKE1cXSXGCnV41pgJqO
7NzfFMGuoVnOXWa8CxHKhlZCJWRi8xSn6RcS9xbBma2ml++epx0Jt85G4NmLAVIy
UffAZhiGtjLpgnbJis9aeZZPAoGAVqM3lGGh5jc5P2uANbk25bpl6kxmxDkdaBg6
mcyB3INPavZXFCWg1w2GQThII3Qr3HEQgXhXMOAV42vBTA71KJFKZvY9l8vd5VE5
iI0qRIABd3OjAx088dmUCdf3cVY7B6N7vrm1UqguYNlyKXguh3jr8IVtlELg9lKk
fvDWUEECgYEAzOk24/WyJSaJ2yQDcRURxQ9chrhFRBSnzgNyG14TPZRNljSUoxCz
3tcXbHydB69NQY5iTc5COS5AksJJRQL08Aejq++ROuNH8nHv3i8td3n5f+0QFc1c
8qhtEsV8xM2J5UYUKZH5aE89vCPRFVOGEx+F8sqHt0rjBd6j8jIkuLo=
-----END RSA PRIVATE KEY-----`

	PKCS8_PRIVATE_KEY = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDD/6ThXz72H5T6
Pdi8MHXQN90HAAWV0pxBfo1E5fyoqNjFt3y2DlnZH56xlj/oH08z9so5x5DNqH2p
zxrUgIdv5o2RHANrV3n0jQ4mh5eKw0VFwKkS24kXL5hFp7yLQ46QM9ZxzzntFKO5
rxDMOL+mrmnipRrBpq8J/7TuDIPJWeSy8o/nRDZjKeIp+O25DyI7UFKWrcsgpRxm
LGuA2GrKaCBSOrm8KJcHMGlYIAqL97p0dAl+P52EwK78LWHhGQQWBgJsbJgZzZzW
+vpsExCp/XKOIBAjfcGfZ5Hd50KTNn23rnrgaf86PW483EIRAFf4G1Jq7jto+fbc
ooj8ktlTAgMBAAECggEAFECyVOd2phAVYNZLMwb+5YkgObzy1sxwgQQGrPvm47au
z59qL5XECv+NcmRILkAt+vyhYN8SQ24JriZMa3Dsjc+RCVGPuOPNIkwBDlIsgpYy
7kY+MkilHdQ5M9llTi6o6JWyrwfuwGTFlkZUwgWLKDoVjunoneIPbEUT+RUp+QA2
qjzX1olzaVcH2+VeHQ7RoWK2qPfYwqTIhbaI+XC4Y58NMj4hNsnQeyyI2LAk7a44
eLP3GYIQM7n3X8tHvnpRee+sbk7VfapIMGetLX6AKYw7v228f590M7oySILGrom6
1X2hXe1mBhMbhmLQNR0KTWNd2sHRc8VddmKIonULAQKBgQDh1Y+RACP3ux/fmjOS
9OrGcP+6aiRgf2yAnyRuqsrDu8mfsXLsRWPzbtofYV1BjQlK5GkYrNpVaaiacgg0
WcWfcpdZlbUGxkfKdmUsUPpNuAkJj2hFdwTMVGpF6icptQnyHgoKMlqnHHvimE5/
0SI14dzJPUQoHqDzN0SyU/qz4QKBgQDeLdl3cPtGkKkJ3Dy8NJuP80er1+pHAEY/
vzQw2oaNVp3d4FCEi6wE8AA/yVwz+/B4rQ8H5cPgAU5mIH/r/OPNHZEyhFW96iMp
hj3kSiKtxhavn4zZvlEyQwkqlVrQ/jcUBZbf1i1v4FJYU2qUb3Oz6eBcMD9QBGG8
44OtpJlzswKBgAgbs9Df8WqmlElh7vIaywce7X+S6gVw7cNcc2FpzmfiqSWLqHWl
NCZ8V2aDB2cgGjJPKxrQsfIZHg2VhIfkpTzJT3fO8TXyDofNYmqJ6cK6ldB7DqJX
2JtBeKNT+d38ewPvYirlje2sDyKel0rthk3q8xNimLzcBKlW1aeVbk3hAoGAIqTN
kwVH+SHXadUwlzwwwYpys4XwXu/ZBFmhHrkEYAeQiCdRBsnsobx2nXudZ/Gx9m8t
t9KI72XkvqH0sGx3bOVeXZ0q7XkT0IbBIda5f5UytmHPiahWatrvE/fD+HCbUADr
HK392cMFtAYVyV1zkis2s8lGJgj44IYtYycWUmcCgYAbHW80j+z1KxHL1BCDLBb+
udSJB7soL4PO6X2CCIhs+yTFUwGFOyPwsyLgPcMn65e18yx7TXCvbhTG+hB4XUeS
7FU14RZjeXVkH23QcRwDcIZcK7peQffIwYh6V1xoeEHnb9KpAXYqWOUb/BwBXBVq
inUYhH6ekAq0R0dljOzhgQ==
-----END PRIVATE KEY-----`

	CRL = `-----BEGIN X509 CRL-----
MIHlMIGMAgEBMAoGCCqGSM49BAMCMBsxGTAXBgNVBAMTEGNsaWVudC1yb290LWNh
LTEXDTIyMDYyMTA5NDQ0NVoXDTIyMDYyODA5NDQ0NVowGzAZAggW+pmWu/XnExcN
MjIwNjIxMDk0NDQ1WqAjMCEwHwYDVR0jBBgwFoAULRlmjBtfjbzwV2WeO9Vj5pWO
h5gwCgYIKoZIzj0EAwIDSAAwRQIhANVFCqByuASAcbz6ovyvi5KCtPfNjHjxVaNT
x69LFPN1AiA5pF5rqHy1FBctZBTW+3LTEEX35j3p1++zcNu8oHMO/w==
-----END X509 CRL-----`
)
