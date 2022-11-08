package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

type CertInfo struct {
	URL                         string // URL地址
	CertEffectiveTime           string // 证书生效时间
	CertEffectiveTimestamp      int64  // 证书生效时间戳
	CertExpirationTime          string // 证书到期时间
	CertExpirationTimestamp     int64  // 证书到期时间戳
	CertExpirationDay           int    // 证书到期天数
	ApplyOrganization           string // 颁发组织
	ApplyOrganizationCommonName string // 颁发组成名称
	IssueOrganization           string // 颁发对象
	IssueOrganizationCommonName string // 颁发对象名称
}

func CertificatesCheck(url string) CertInfo {
	time.Now().Unix()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Second * 5}

	resp, err := client.Get("https://" + url)
	if err != nil {
		log.Printf("http get url:%s err:%s", url, err)
		return CertInfo{URL: url}
	}

	defer resp.Body.Close()

	var applyOrganizationInfo *x509.Certificate
	var issueOrganizationInfo *x509.Certificate

	if len(resp.TLS.PeerCertificates) != 0 {
		applyOrganizationInfo = resp.TLS.PeerCertificates[0]
	}

	if len(resp.TLS.PeerCertificates) != 0 {
		issueOrganizationInfo = resp.TLS.PeerCertificates[1]
	}

	return CertInfo{
		URL:                         url,
		CertEffectiveTime:           applyOrganizationInfo.NotBefore.Local().Format("2006-01-02 15:04:05"),
		CertEffectiveTimestamp:      applyOrganizationInfo.NotBefore.Local().Unix(),
		CertExpirationTime:          applyOrganizationInfo.NotAfter.Local().Format("2006-01-02 15:04:05"),
		CertExpirationTimestamp:     applyOrganizationInfo.NotAfter.Unix(),
		CertExpirationDay:           int(math.Floor(applyOrganizationInfo.NotAfter.Sub(time.Now()).Hours()/24 + 0.5)), // math.Floor 四舍五入天数
		ApplyOrganization:           strings.Join(issueOrganizationInfo.Subject.Organization, ""),
		ApplyOrganizationCommonName: issueOrganizationInfo.Subject.CommonName,
		IssueOrganization:           strings.Join(applyOrganizationInfo.Subject.Organization, ""),
		IssueOrganizationCommonName: applyOrganizationInfo.Subject.CommonName,
	}
}

func main() {
	CertInfo := CertificatesCheck("www.baidu.com")
	fmt.Println("URL地址", CertInfo.URL)
	fmt.Println("证书生效时间", CertInfo.CertEffectiveTime)
	fmt.Println("证书生效时间戳", CertInfo.CertEffectiveTimestamp)
	fmt.Println("证书到期时间", CertInfo.CertExpirationTime)
	fmt.Println("证书到期时间戳", CertInfo.CertExpirationTimestamp)
	fmt.Println("证书到期天数", CertInfo.CertExpirationDay)
	fmt.Println("颁发组织", CertInfo.ApplyOrganization)
	fmt.Println("颁发组成名称", CertInfo.ApplyOrganizationCommonName)
	fmt.Println("颁发对象", CertInfo.IssueOrganization)
	fmt.Println("颁发对象名称", CertInfo.IssueOrganizationCommonName)
}
