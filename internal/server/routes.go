package server

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/middleware"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/aIConversation"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/aIMessage"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/comment"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/content"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/directMessage"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/feedback"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/follow"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/groupChat"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/like"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/message"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/notification"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/save"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/user"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.CorsMiddleware,
		middleware.AuthMiddleware,
	)

	userRepository := user.NewUserRepository(s.DB)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService, s.S3Client, s.BucketName)

	commentRepository := comment.NewContentRepository(s.DB)
	commentService := comment.NewContentService(commentRepository)
	commentHandler := comment.NewHandler(commentService)

	likeRepository := like.NewLikeRepository(s.DB)
	likeService := like.NewLikeService(likeRepository)
	likeHandler := like.NewHandler(likeService)

	saveRepository := save.NewSaveRepository(s.DB)
	saveService := save.NewSaveService(saveRepository)
	saveHandler := save.NewHandler(saveService)

	contentRepository := content.NewContentRepository(s.DB)
	contentService := content.NewContentService(contentRepository)
	contentHandler := content.NewHandler(contentService, s.S3Client, s.BucketName, likeService, saveService)

	feedbackRepository := feedback.NewFeedbackRepository(s.DB)
	feedbackService := feedback.NewFeedbackService(feedbackRepository)
	feedbackHandler := feedback.NewHandler(feedbackService)

	messageRepository := message.NewMessageRepository(s.DB)
	messageService := message.NewMessageService(messageRepository)
	messageHandler := message.NewHandler(messageService)

	directMessageRepository := directmessage.NewDirectMessageRepository(s.DB)
	directMessageService := directmessage.NewDirectMessageService(directMessageRepository)
	directMessageHandler := directmessage.NewHandler(directMessageService, messageRepository)

	groupChatRepository := groupchat.NewGroupChatRepository(s.DB)
	groupChatService := groupchat.NewGroupChatService(groupChatRepository)
	groupChatHandler := groupchat.NewHandler(groupChatService, messageRepository)

	aIConversationRepository := aIconversation.NewAIConversationRepository(s.DB)
	aIConversationService := aIconversation.NewAIConversationService(aIConversationRepository)
	aIConversationHandler := aIconversation.NewHandler(aIConversationService)

	aiMessageRepository := aImessage.NewAIMessageRepository(s.DB)
	aiMessageService := aImessage.NewAIMessageService(aiMessageRepository)
	aiMessageHandler := aImessage.NewHandler(aiMessageService)

	followRepository := follow.NewFollowRepository(s.DB)
	followService := follow.NewFollowService(followRepository)
	followHandler := follow.NewHandler(followService)

	notificationRepository := notification.NewNotificationRepository(s.DB)
	notificationService := notification.NewNotificationService(notificationRepository)
	notificationHandler := notification.NewHandler(notificationService)

	mux.HandleFunc("POST /v1/user", userHandler.Register)
	mux.HandleFunc("POST /v1/user/login", userHandler.Login)
	mux.HandleFunc("GET /v1/users/{username}", userHandler.FindAllUsers)
	mux.HandleFunc("GET /v1/users/{userId}/mutual", userHandler.FindAllMutualUsers)
	mux.HandleFunc("GET /v1/user/{username}", userHandler.FindOneUser)
	mux.HandleFunc("PATCH /v1/user/{userId}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /v1/user/{userId}/logout", userHandler.Logout)
	mux.HandleFunc("DELETE /v1/user/{userId}", userHandler.DeleteUser)
	mux.HandleFunc("GET /v1/token", userHandler.GetToken)

	mux.HandleFunc("POST /v1/content", contentHandler.UploadContent)
	mux.HandleFunc("GET /v1/contents/{userId}", contentHandler.GetAllContent)
	mux.HandleFunc("GET /v1/contents/{userId}/following", contentHandler.GetAllContentByFollowing)
	mux.HandleFunc("GET /v1/content/{contentId}", contentHandler.GetOneContent)
	mux.HandleFunc("PATCH /v1/content/{contentId}", contentHandler.UpdateContent)
	mux.HandleFunc("DELETE /v1/content/{contentId}", contentHandler.DeleteContent)

	mux.HandleFunc("POST /v1/comment/{contentId}", commentHandler.SendComment)
	mux.HandleFunc("GET /v1/comments/{contentId}", commentHandler.GetAllComments)
	mux.HandleFunc("PATCH /v1/comment/{commentId}", commentHandler.UpdateComment)
	mux.HandleFunc("DELETE /v1/comment/{commentId}", commentHandler.DeleteComment)

	mux.HandleFunc("POST /v1/like/{contentId}", likeHandler.LikeContent)
	mux.HandleFunc("DELETE /v1/like/{likeId}", likeHandler.UnlikeContent)

	mux.HandleFunc("POST /v1/save/{contentId}", saveHandler.SaveContent)
	mux.HandleFunc("GET /v1/saves/{userId}", saveHandler.GetAllSaves)
	mux.HandleFunc("DELETE /v1/save/{saveId}", saveHandler.UnsaveContent)

	mux.HandleFunc("POST /v1/feedback/{userId}", feedbackHandler.SendFeedback)
	mux.HandleFunc("GET /v1/feedback", feedbackHandler.GetAllFeedbacks)

	mux.HandleFunc("POST /v1/message/dm/{conversationId}", messageHandler.SendPrivateMessage)
	mux.HandleFunc("POST /v1/message/group/{conversationId}", messageHandler.SendGroupMessage)
	mux.HandleFunc("PATCH /v1/message/{messageId}", messageHandler.UpdateMessage)
	mux.HandleFunc("DELETE /v1/message/{messageId}", messageHandler.DeleteMessage)

	mux.HandleFunc("POST /v1/direct", directMessageHandler.CreateDirectMessage)
	mux.HandleFunc("/v1/ws/direct", directMessageHandler.HandleWebSocket)
	mux.HandleFunc("GET /v1/directs/{userId}", directMessageHandler.GetAllDirectMessages)
	mux.HandleFunc("GET /v1/direct", directMessageHandler.GetOneDirectMessageByParticipants)
	mux.HandleFunc("GET /v1/direct/{dmId}", directMessageHandler.GetOneDirectMessage)
	mux.HandleFunc("DELETE /v1/direct/{dmId}", directMessageHandler.DeleteDirectMessage)

	mux.HandleFunc("POST /v1/group", groupChatHandler.CreateGroupChat)
	mux.HandleFunc("POST /v1/group/members/{groupId}", groupChatHandler.AddMembers)
	mux.HandleFunc("/v1/ws/group", groupChatHandler.HandleWebSocket)
	mux.HandleFunc("GET /v1/groups/{userId}", groupChatHandler.GetAllGroupChats)
	mux.HandleFunc("GET /v1/group/{groupId}", groupChatHandler.GetOneGroupChat)
	mux.HandleFunc("PATCH /v1/group/{groupId}", groupChatHandler.UpdateGroupChat)
	mux.HandleFunc("DELETE /v1/group/{groupId}/members/{userId}", groupChatHandler.LeaveGroupChat)
	mux.HandleFunc("DELETE /v1/group/{groupId}", groupChatHandler.DeleteGroupChat)

	mux.HandleFunc("POST /v1/ai/conv/{userId}", aIConversationHandler.CreateConversation)
	mux.HandleFunc("GET /v1/ai/conv/{userId}", aIConversationHandler.GetConversation)
	mux.HandleFunc("DELETE /v1/ai/conv/{conversationId}", aIConversationHandler.DeleteConversation)

	mux.HandleFunc("POST /v1/ai/message/{conversationId}", aiMessageHandler.SendMessage)
	mux.HandleFunc("DELETE /v1/ai/message/{messageId}", aiMessageHandler.DeleteMessage)

	mux.HandleFunc("POST /v1/follow", followHandler.FollowUser)
	mux.HandleFunc("GET /v1/follows/{userId}", followHandler.CountFollow)
	mux.HandleFunc("GET /v1/follow", followHandler.CheckFollowing)
	mux.HandleFunc("DELETE /v1/follow/{followId}", followHandler.UnfollowUser)

	mux.HandleFunc("GET /v1/notifications/{receiverId}", notificationHandler.GetAllNotifications)
	mux.HandleFunc("DELETE /v1/notification/{receiverId}", notificationHandler.DeleteNotifications)

	return stack(mux)
}
