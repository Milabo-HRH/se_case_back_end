package routers

import (
	"github.com/gin-gonic/gin"
	"se_case_back_end/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	//r.Use(middleware.CorsMiddleware())
	//
	//r.POST("/user/register", controller.Register)
	//r.POST("/user/login", controller.Login)
	//r.GET("/user/me", middleware.AuthMiddleware, controller.Info)
	//
	//r.POST("/pics", middleware.AuthMiddleware, middleware.AuthPrivilege("normal"), controller.FileUpload)
	//r.GET("/pics/:filename", middleware.AuthPrivilege("guest"), controller.GetPicture)
	//
	//r.GET("/tasks", controller.GetTasks)
	//r.GET("/tasks/unsolved", controller.GetUnsolvedTasks)
	//r.POST("/tasks/publish", middleware.AuthMiddleware, middleware.AuthPrivilege("normal"), controller.PublishTask)
	//
	//r.GET("/annotations", middleware.AuthMiddleware, middleware.AuthPrivilege("important"), controller.GetAnnotations)
	//r.GET("/annotations/unsolved", middleware.AuthMiddleware, middleware.AuthPrivilege("important"), controller.GetUnsolvedAnnotations)
	//r.POST("/annotations/publish", middleware.AuthMiddleware, middleware.AuthPrivilege("normal"), controller.PublishAnnotation)
	//r.POST("/annotations/pass", middleware.AuthMiddleware, middleware.AuthPrivilege("important"), controller.PassAnnotation)
	////r.POST("/annotations/pass", middleware.AuthMiddleware, controller.PassAnnotation)
	//r.POST("/annotations/reject", middleware.AuthMiddleware, middleware.AuthPrivilege("important"), controller.DeleteAnnotation)
	////r.POST("/annotations/reject", middleware.AuthMiddleware, controller.DeleteAnnotation)
	r.GET("/api/record/:patient_id", controller.GetRecordByPID)
	r.GET("/api/view/:id", controller.GetRecord)
	r.POST("/api/newcase", controller.Register)
	return r
}
