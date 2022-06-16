package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

type verifyUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type verifyUserResponse struct {
	Valid bool `json:"valid"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	verifyReq := verifyUserRequest{
		Username: req.Username,
		Password: req.Password,
	}
	jsonVerifyReq, _ := json.Marshal(verifyReq)

	resp, err := http.Post(server.config.VerifyServerAddr, "application/json", bytes.NewBuffer(jsonVerifyReq))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusUnauthorized, errors.New("provided data is invalid"))
		return
	}

	var verifyResult verifyUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&verifyResult); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !verifyResult.Valid {
		ctx.JSON(http.StatusUnauthorized, errors.New("provided data is invalid"))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.db.Set(accessToken, req.Username, 0)

	rsp := loginUserResponse{
		Username:    req.Username,
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type verifyTokenRequest struct {
	Username    string `json:"username" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
}

func (server *Server) verify(ctx *gin.Context) {
	var req verifyTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username, err := server.db.Get(req.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if username != req.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "valid")
}
