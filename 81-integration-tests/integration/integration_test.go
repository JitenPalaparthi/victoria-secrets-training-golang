package integration_test

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
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

		cmd := exec.Command(container, "run", "-d", "--name=mysql3", "--network=host", "-e", "MYSQL_ROOT_PASSWORD=admin123", "-e", "MYSQL_USER=admin", "-e", "MYSQL_PASSWORD=admin123", "-e", "MYSQL_DATABASE=contactsdb", "docker.io/library/mysql")
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
		dns = fmt.Sprintf("admin:admin123@tcp(%v:3306)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0", "127.0.0.1")

		cmd = exec.Command(container, "run", "-d", "--name=app1", "--network=host", "-e", "PORT=50091", "-e", "DSN="+dns, "jpalaparthi/app:v0.0.3")

		_, err = cmd.Output()
		if err != nil {
			fmt.Println(string(err.Error()))
			Fail(err.Error())
		}
	})
	AfterEach(func() {
		// cmd := exec.Command(container, "rm", "-f", "mysql3", "app1")
		// _, err := cmd.Output()
		// if err != nil {
		// 	Fail(err.Error())
		// }
	})

	// insert contact
	It("create contact", func() {

		resp, err := http.Get("http://localhost:50091/ping")
		if err != nil {
			Fail("1------>>" + err.Error())
		}

		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)

		if err != nil {
			Fail("2------>>" + err.Error())
		}
		if string(bytes) != "pong" {
			Fail("----3" + string(bytes))
		}
		// 	cmd := exec.Command(container, "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "app1") // create a command
		// 	ip, err := cmd.Output()
		// 	if err != nil {
		// 		Fail(err.Error())
		// 	}
		// 	ip = []byte(strings.ReplaceAll(string(ip), "\n", ""))

		// 	//url := "http://" + string(ip) + ":58089/contact"
		// 	url := "http://localhost:50091/contact"
		// 	method := "POST"

		// 	payload := strings.NewReader(`{
		// 	  "name": "Abhinav",
		// 	  "email": "Abhi@victoria.com",
		// 	  "mobile": "32432423"
		//   }`)

		// 	client := &http.Client{}
		// 	req, err := http.NewRequest(method, url, payload)

		// 	if err != nil {
		// 		Fail(err.Error())
		// 	}
		// 	req.Header.Add("Content-Type", "application/json")

		// 	res, err := client.Do(req)
		// 	if err != nil {
		// 		Fail(err.Error())
		// 	}
		// 	defer res.Body.Close()

		// 	contact := new(models.Contact)

		// 	err = json.NewDecoder(res.Body).Decode(&contact)

		// 	if err != nil {
		// 		Fail(err.Error())
		// 	}

		// 	if contact == nil {
		// 		Fail("no data found")
		// 	}

		// 	if contact.Email != "Abhi@victoria.com" || contact.Name != "Abhinav" || contact.Mobile != "32432423" {
		// 		Fail("data does not match")
		// 	}

	})
})
