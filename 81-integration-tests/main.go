package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "run", "-d", "--name=mysql3", "-p", "34406:3306", "-e", "MYSQL_ROOT_PASSWORD=admin123", "-e", "MYSQL_USER=admin", "-e", "MYSQL_PASSWORD=admin123", "docker.io/library/mysql")
	_, err := cmd.Output()
	fmt.Println(err)

	//docker inspect mysql2 | grep -w "IPAddress" | awk '{ print $2 }' | head -n 1 | cut -d "," -f1
	cmd = exec.Command("docker", "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "mysql3") // create a command
	ip, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	dns := fmt.Sprintf("admin:admin123@tcp(%v:3606)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0", string(ip))
	cmd = exec.Command("docker", "run", "-d", "-p", "58090:58989", "--name=app1", "-e", "PORT=58089", "-e", "DSN="+dns, "jpalaparthi/app:v0.0.2")

	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("docker", "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "app1") // create a command
	ip1, err := cmd.Output()
	fmt.Println(string(ip1))
	if err != nil {
		fmt.Println(err)
	}

}
