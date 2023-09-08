package integration_test

import (
	"demo/models"
	"encoding/json"
	"fmt"
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

var _ = Describe("Integration", func() {

	container := "docker"

	//fmt.Println("Hello")
	BeforeEach(func() {

		cmd := exec.Command(container, "run", "-d", "--name=mysql3", "-p", "34406:3306", "-e", "MYSQL_ROOT_PASSWORD=admin123", "-e", "MYSQL_USER=admin", "-e", "MYSQL_PASSWORD=admin123", "-e", "MYSQL_DATABASE=contactsdb", "docker.io/library/mysql")
		id, err := cmd.Output()
		fmt.Println(id)
		if err != nil {
			Fail(err.Error())
		}

		cmd = exec.Command(container, "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "mysql3") // create a command
		ip, err := cmd.Output()
		if err != nil {
			Fail(err.Error())
		}
		dns := fmt.Sprintf("admin:admin123@tcp(%v:3306)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0", string(ip))

		cmd = exec.Command(container, "run", "-d", "--name=app1", "-p", "58089:58089", "-e", "DSN="+dns, "jpalaparthi/app:v0.0.2")

		_, err = cmd.Output()
		if err != nil {
			fmt.Println(string(err.Error()))
			Fail(err.Error())
		}
	})
	AfterEach(func() {
		cmd := exec.Command(container, "rm", "-f", "mysql3", "app1")
		_, err := cmd.Output()
		if err != nil {
			Fail(err.Error())
		}
	})

	// insert contact
	It("create contact", func() {

		cmd := exec.Command(container, "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "app1") // create a command
		ip, err := cmd.Output()
		if err != nil {
			Fail(err.Error())
		}
		ip = []byte(strings.ReplaceAll(string(ip), "\n", ""))

		//url := "http://" + string(ip) + ":58089/contact"
		url := "http://127.0.0.1:58089/contact"
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

		contact := new(models.Contact)

		err = json.NewDecoder(res.Body).Decode(&contact)

		if err != nil {
			Fail(err.Error())
		}

		if contact == nil {
			Fail("no data found")
		}

		if contact.Email != "Abhi@victoria.com" || contact.Name != "Abhinav" || contact.Mobile != "32432423" {
			Fail("data does not match")
		}

	})
})
