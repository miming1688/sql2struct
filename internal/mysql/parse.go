package mysqlparser

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gangming/sql2struct/config"
	"github.com/gangming/sql2struct/utils"
)

func GenerateFile(ddl string) error {
	c, _ := ParseMysqlDDL(ddl)
	dir := config.Cnf.OutputDir
	os.MkdirAll(dir, 0755)
	fileName := filepath.Join(dir, strings.ToLower(c.Name)+".go")
	fd, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer fd.Close()
	fd.Write([]byte(c.GenerateCode()))

	_, err = exec.Command("goimports", "-l", "-w", dir).Output()
	if err != nil {
		utils.PrintRed(err.Error())
	}
	_, err = exec.Command("gofmt", "-l", "-w", dir).Output()
	if err != nil {
		utils.PrintRed(err.Error())
	}
	utils.PrintGreen(fileName + " generate success")
	return nil
}
