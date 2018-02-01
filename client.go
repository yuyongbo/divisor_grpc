package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"google.golang.org/grpc"
	"gongyueshu/pb"
	"log"
)

func main(){
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	gcdClient := pb.NewGCDServiceClient(conn)

	r := gin.Default()
	r.GET("/gcd/:a/:b", func(c *gin.Context) { // Parse parameters
		a, err := strconv.ParseUint(c.Param("a"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		b, err := strconv.ParseUint(c.Param("b"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
			return
		} // Call GCD service
		req := &pb.GCDRequest{A: a, B: b}
		if res, err := gcdClient.Compute(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(res.Result)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	if err := r.Run(":3001"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}