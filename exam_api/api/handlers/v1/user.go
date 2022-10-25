package v1

import (
	"context"
	"examLast/exam_api/api/handlers/models"
	"fmt"
	"net/http"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	pbs "examLast/exam_api/genproto/custumer_proto"
	pb "examLast/exam_api/genproto/post_proto"
	pr "examLast/exam_api/genproto/reating_proto"

	//pr "exam_api/genproto/reating_proto"
	l "examLast/exam_api/pkg/logger"
)

type AllthingPost struct {
	Postinfo   pb.Posts
	Posterinfo pbs.CustumerInfo
}

// GetReating get Reating
// @Summary      Get  posts reating api
// @Description this api get posts reating by id
// @Tags Post
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post ID"
// @Success 200 {object} reating.Reatings
// @Router /v1/post/get/reatings/{id}  [get]
func (h *handlerV1) GetPostReating(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.ReatingService().GetPostReating(ctx, &pr.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetCustumers get Custumers
// @Summary      Get only custumers api
// @Description this api get custumers
// @Tags Custumer
// @Accept json
// @Produce json
// @Success 200 {object} custumer.CustumerAll
// @Router /v1/custumer/getList [get]
func (h *handlerV1) GetListCustumers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.CustumerService().ListAllCustum(ctx, &pbs.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// CreatePost create post
// @Summary      create post api
// @Description this api create post
// @Tags Post
// @Accept json
// @Produce json
// @Param request body post.PostForCreate true "Custumer"
// @Success 200 {object} post.PostInfo
// @Router /v1/post/create [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        pb.PostForCreate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore, err := h.serviceManager.PostService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// DeleteCustumer delete Post
// @Summary      delete Post api
// @Description this api posts by id
// @Tags Post
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post id"
// @Success 200 {object} post.EmptyPost
// @Router /v1/post/delet/{id}  [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore, err := h.serviceManager.PostService().Delet(ctx, &pb.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resStore)
}

// CreateStore create store
// @Summary      create store api
// @Description this api create store
// @Tags Custumer
// @Accept json
// @Produce json
// @Param request body custumer.CustumerForCreate true "Custumer"
// @Success 200 {object} custumer.CustumerInfo
// @Router /v1/custumer/create [post]
func (h *handlerV1) CreateCustumer(c *gin.Context) {
	var (
		body        pbs.CustumerForCreate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pbs.CustumerInfo{}
	resStore, err = h.serviceManager.CustumerService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// CreateStore create store
// @Summary      create reating api
// @Description this api create reating
// @Tags Reating
// @Accept json
// @Produce json
// @Param request body reating.ReatingForCreate true "Custumer"
// @Success 200 {object} reating.ReatingInfo
// @Router /v1/reating/create [post]
func (h *handlerV1) CreateReating(c *gin.Context) {
	var (
		body        pr.ReatingForCreate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	if body.Reating > 5 {
		c.JSON(http.StatusBadRequest, "reating false")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pr.ReatingInfo{}
	resStore, err = h.serviceManager.ReatingService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// GetReating get Reating
// @Summary      Get reating api
// @Description this api get reating by id
// @Tags Reating
// @Accept json
// @Produce json
// @Param   id   path  int  true  "reating id"
// @Success 200 {object} reating.ReatingInfo
// @Router /v1/reating/get/{id}  [get]
func (h *handlerV1) GetReating(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.ReatingService().GetReating(ctx, &pr.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteCustumer delete Custumer
// @Summary      delete Reating api
// @Description this api delet reating by id
// @Tags Reating
// @Accept json
// @Produce json
// @Param   id   path  int  true  "reating id"
// @Success 200 {object} reating.EmptyReating
// @Router /v1/reating/delete/{id}  [delete]
func (h *handlerV1) DeleteReating(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore, err := h.serviceManager.ReatingService().Delet(ctx, &pr.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resStore)
}

// DeleteCustumer delete Custumer
// @Summary      delete Custumer api
// @Description this api delet custumer with posts by id
// @Tags Custumer
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Custumer id"
// @Success 200 {object} custumer.Empty
// @Router /v1/custumer/delete/{id}  [delete]
func (h *handlerV1) DeleteCustumer(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resStore, err := h.serviceManager.CustumerService().DeletCustum(ctx, &pbs.GetId{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resStore)
}

// GetCustumer get Custumer
// @Summary      get Custumer api
// @Description this api get custumer with posts by id
// @Tags Custumer
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Custumer id"
// @Success 200 {object} custumer.CustumerInfo
// @Router /v1/custumer/get/{id}  [get]
func (h *handlerV1) GetCustumer(c *gin.Context) {
	reul := models.CustumerAllInfo{}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustumerService().GetByCustumId(ctx, &pbs.GetId{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	reul.Custumer = *response
	response2, err := h.serviceManager.PostService().GetByOwnerID(ctx, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	azaz := reul.Posts
	aza := models.Postc{}
	//reatings := []*pr.ReatingInfo{}
	for _, post := range response2.Posts {
		aza.Post = *post
		response3, err := h.serviceManager.ReatingService().GetPostReating(ctx, &pr.Id{Id: post.Id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		aza.Reatings = *response3
	}
	azaz = append(azaz, aza)

	reul.Posts = azaz

	c.JSON(http.StatusCreated, reul)
}

// CreateStore create store
// @Summary      update reating api
// @Description this api update reating
// @Tags Reating
// @Accept json
// @Produce json
// @Param request body reating.ReatingInfo true "reating"
// @Success 200 {object} reating.ReatingInfo
// @Router /v1/reating/update [put]
func (h *handlerV1) UpdateReating(c *gin.Context) {
	var (
		body        pr.ReatingInfo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pr.ReatingInfo{}
	resStore, err = h.serviceManager.ReatingService().Update(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// GetOnlyPost get Post
// @Summary      Get  post api
// @Description this api get post by id
// @Tags Post
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post ID"
// @Success 200 {object} post.PostInfo
// @Router /v1/post/get/{id}  [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.PostService().GetPost(ctx,&pb.Id{Id: id} )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}


// CreateStore create store
// @Summary      update custumer api
// @Description this api update custumer
// @Tags Custumer
// @Accept json
// @Produce json
// @Param request body custumer.CustumerInfo true "Custumer"
// @Success 200 {object} custumer.CustumerInfo
// @Router /v1/custumer/update [put]
func (h *handlerV1) UpdateCustumer(c *gin.Context) {
	var (
		body        pbs.CustumerInfo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pbs.CustumerInfo{}
	resStore, err = h.serviceManager.CustumerService().Update(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// CreateStore create store
// @Summary      update post api
// @Description this api update Post
// @Tags Post
// @Accept json
// @Produce json
// @Param request body post.PostInfo true "Post"
// @Success 200 {object} post.PostInfo
// @Router /v1/post/update [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pb.PostInfo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pb.PostInfo{}
	resStore, err = h.serviceManager.PostService().Update(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// GetPost get Post
// @Summary      get Post api
// @Description this api get Post by id
// @Tags Post
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Poster id"
// @Success 200 "object"
// @Router /v1/post/allInfo/{id}  [get]
func (h *handlerV1) GetPostAllInfo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetByOwnerID(ctx, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return

	}
	var azaz int64
	for _, poster := range response.Posts {
		azaz = poster.PosterId
	}

	fmt.Println(azaz)
	response2, err := h.serviceManager.CustumerService().GetByCustumId(ctx, &pbs.GetId{Id: azaz})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}
	natija := AllthingPost{}
	natija.Posterinfo = *response2
	natija.Postinfo = *response
	c.JSON(http.StatusOK, natija)
}
