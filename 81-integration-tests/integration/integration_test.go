package integration_test

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contacts Integration Suite")
}

// var _ = BeforeSuite(func() {
// 	// create two docker instances
// 	// teardown two docker instances
// 	//docker run -d --name=mysql1 -p 33306:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 docker.io/library/mysql
// 	cmd := exec.Command("docker", "run", "-d", "--name=mysql3", "-p", "34406:3306", "-e", "MYSQL_ROOT_PASSWORD=admin123", "-e", "MYSQL_USER=admin", "-e", "MYSQL_PASSWORD=admin123", "docker.io/library/mysql")
// 	_, err := cmd.Output()
// 	Fail(err.Error())
// })

// var _ = AfterSuite(func() {

// })

var _ = Describe("Integration", func() {
	//fmt.Println("Hello")
	BeforeEach(func() {
		cmd := exec.Command("docker", "run", "-d", "--name=mysql3", "-p", "34406:3306", "-e", "MYSQL_ROOT_PASSWORD=admin123", "-e", "MYSQL_USER=admin", "-e", "MYSQL_PASSWORD=admin123", "docker.io/library/mysql")
		id, err := cmd.Output()
		fmt.Println(id)
		if err != nil {
			Fail(err.Error())
		}

		cmd = exec.Command("docker", "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "mysql3") // create a command
		ip, err := cmd.Output()
		if err != nil {
			Fail(err.Error())
		}

		cmd = exec.Command("docker", "run", "-d", "--name=app1", "-e", "PORT=58089", "-e", `DSN=admin:admin123@tcp(`+string(ip)+`:34406)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0`, "jpalaparthi/app:v0.0.1")

		_, err = cmd.Output()
		if err != nil {
			Fail(err.Error())
		}
	})
	AfterEach(func() {
		cmd := exec.Command("docker", "rm", "-f", "mysql3 app1")
		_, err := cmd.Output()
		if err != nil {
			Fail(err.Error())
		}
	})

	It("testing docker mysql instance", func() {
		cmd := exec.Command("docker", "ps")
		ps, err := cmd.Output()
		if !strings.Contains(string(ps), "mysql3") {
			Fail("mysql docker instance does not exist")
		}
		if err != nil {
			Fail(err.Error())
		}
	})

	It("testing docker app instance", func() {
		cmd := exec.Command("docker", "ps")
		ps, err := cmd.Output()
		if !strings.Contains(string(ps), "app1") {
			Fail("app docker instance does not exist")
		}
		if err != nil {
			Fail(err.Error())
		}
	})

	// insert contact
	It("create contact", func() {

		url := "localhost:58089/contact"
		method := "POST"

		payload := strings.NewReader(`{
		  "name": "Abhinav",
		  "email": "Abhi@victoria.com",
		  "mobile": "32432423"
	  }`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			Fail(err.Error())
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			Fail(err.Error())
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			Fail(err.Error())
		}
		fmt.Println(string(body))

		if !strings.Contains(string(body), "1") {
			Fail("unable to get the id")
		}

	})
})
