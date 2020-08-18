package whois

import (
	"fmt"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/duoflow/whois-email-resolver/loggers"
)

const (
	// DefaultWhoisPort is default whois port
	DefaultWhoisPort = "43"
)

// GetEmailByASN - method to get email by ASN
func GetEmailByASN(asn string, whoisserver string) (string, error) {
	loggers.Info.Printf("whois.GetEmailByASN starts")
	//result, err := whois.Whois(asn, whoisserver)
	result, err := query(asn, whoisserver, DefaultWhoisPort)
	// error handling
	if err != nil {
		return "", fmt.Errorf("GetEmailByASN Error: read from whois server failed: %v", err)
	}
	// parse admin-c ID
	loggers.Info.Printf("whois.GetEmailByASN start finding admin-c")
	re := regexp.MustCompile(`(?m)(admin-c:)( *)(.*)`)
	res := re.FindAllStringSubmatch(result, -1)
	// if error in parsing admin-c: detected
	if res != nil {
		loggers.Info.Printf("whois.GetEmailByASN Success of finding admin-c: %s", res[0][3])
	} else {
		return "", nil
	}
	// Step 2 parse email
	// query new object - admin-c
	loggers.Info.Printf("whois.GetEmailByASN start finding email")
	result, err = query(strings.TrimSpace(res[0][3]), whoisserver, DefaultWhoisPort)
	// parse e-mail:
	re = regexp.MustCompile(`(?m)(e-mail:)( *)(.*)`)
	res = re.FindAllStringSubmatch(result, -1)
	// if error in parsing admin-c: detected
	if res != nil {
		loggers.Info.Printf("whois.GetEmailByASN Success of finding email: %s", res[0][3])
		return strings.TrimSpace(res[0][3]), nil
	}
	// return null values
	loggers.Info.Printf("whois.GetEmailByASN Success of finding admin-c but NOT email:")
	return "", nil
}

// query - send query to server
func query(object string, server string, tcpport string) (string, error) {
	// open connnection
	loggers.Info.Printf("whois.query() setup connection")
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, tcpport), time.Second*30)
	if err != nil {
		return "", fmt.Errorf("whois: connect to whois server failed: %v", err)
	}
	defer conn.Close()
	// set connection write timeout
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
	_, err = conn.Write([]byte(object + "\r\n"))
	if err != nil {
		return "", fmt.Errorf("whois: send to whois server failed: %v", err)
	}
	// set connection read timeout
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", fmt.Errorf("whois: read from whois server failed: %v", err)
	}
	// return result
	return string(buffer), nil
}
