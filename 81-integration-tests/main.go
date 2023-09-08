package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "run", "-d", "--name=mysql3", "-p", "34406:3306", "-e", "MYSQL_ROOT_PASSWORD=admin123", "-e", "MYSQL_USER=admin", "-e", "MYSQL_PASSWORD=admin123", "docker.io/library/mysql")
	_, err := cmd.Output()
	fmt.Println(err)

	cmd = exec.Command("docker", "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "mysql2") // create a command
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} // execute command
	fmt.Println(string(output)) //print command output
	//docker inspect mysql2 | grep -w "IPAddress" | awk '{ print $2 }' | head -n 1 | cut -d "," -f1
	cmd = exec.Command("docker", "inspect", "-f", `{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}`, "mysql3") // create a command
	ip, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	cmd = exec.Command("docker", "run", "-d", "--name=app1", "-e", "PORT=58089", "-e", "DSN=admin:admin123@tcp("+string(ip)+":34406)/contactsdb?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0", "jpalaparthi/app:v0.0.1")

	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

}
