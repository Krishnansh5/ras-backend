package company

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func get100CompaniesHandler(ctx *gin.Context) {
	var companies []Company

	lastCid,err := strconv.ParseUint(ctx.Param("lastCid"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	err = getNext100Companies(ctx, &companies, uint(lastCid))
	timeElapsed := time.Since(start)

	log.Printf("The `db` call took %s", timeElapsed)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func getAllCompaniesHandler(ctx *gin.Context) {
	var companies []Company

	start := time.Now()

	err := getAllCompanies(ctx, &companies)
	timeElapsed := time.Since(start)
	log.Printf("The `for` loop took %s", timeElapsed)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, companies)
}

func getCompanyHandler(ctx *gin.Context) {
	var company Company

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getCompany(ctx, &company, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}
