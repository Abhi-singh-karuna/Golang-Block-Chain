package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func InitHandler(c *gin.Context) {
	A := c.Query("A")
	Aval := c.Query("Aval")
	B := c.Query("B")
	Bval := c.Query("Bval")

	cmd := exec.Command("docker", "exec", "cli", "peer", "chaincode", "invoke",
		"--tls", "--cafile", "/opt/home/managedblockchain-tls-chain.pem",
		"--channelID", "mychannel",
		"--name", "mycc",
		"-c", fmt.Sprintf("{\"Args\":[\"Init\", \"%s\", \"%s\", \"%s\", \"%s\"]}", A, Aval, B, Bval))

	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": string(out)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": string(out)})
}

func QueryHandler(c *gin.Context) {
	A := c.Query("A")

	cmd := exec.Command("docker", "exec", "cli", "peer", "chaincode", "query",
		"--tls", "--cafile", "/opt/home/managedblockchain-tls-chain.pem",
		"--channelID", "mychannel",
		"--name", "mycc",
		"-c", fmt.Sprintf("{\"Args\":[\"Query\", \"%s\"]}", A))

	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": string(out)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": string(out)})
}

func main() {
	r := gin.Default()

	r.POST("/init", InitHandler)
	r.GET("/query", QueryHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
