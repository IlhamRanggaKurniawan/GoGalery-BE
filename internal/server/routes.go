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
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
	)

	userRepository := user.NewUserRepository(s.DB)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)

	contentRepository := content.NewContentRepository(s.DB)
	contentService := content.NewContentService(contentRepository)
	contentHandler := content.NewHandler(contentService)

	commentRepository := comment.NewContentRepository(s.DB)
	commentService := comment.NewContentService(commentRepository)
	commentHandler := comment.NewHandler(commentService)

	likeRepository := like.NewLikeRepository(s.DB)
	likeService := like.NewLikeService(likeRepository)
	likeHandler := like.NewHandler(likeService)

	saveRepository := save.NewSaveRepository(s.DB)
	saveService := save.NewSaveService(saveRepository)
	saveHandler := save.NewHandler(saveService)

	feedbackRepository := feedback.NewFeedbackRepository(s.DB)
	feedbackService := feedback.NewFeedbackService(feedbackRepository)
	feedbackHandler := feedback.NewHandler(feedbackService)

	messageRepository := message.NewMessageRepository(s.DB)
	messageService := message.NewMessageService(messageRepository)
	messageHandler := message.NewHandler(messageService)

	directMessageRepository := directmessage.NewDirectMessageRepository(s.DB)
	directMessageService := directmessage.NewDirectMessageService(directMessageRepository)
	directMessageHandler := directmessage.NewHandler(directMessageService)

	groupChatRepository := groupchat.NewGroupChatRepository(s.DB)
	groupChatService := groupchat.NewGroupChatService(groupChatRepository)
	groupChatHandler := groupchat.NewHandler(groupChatService)

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

	mux.HandleFunc("POST /user/register", userHandler.Register)
	mux.HandleFunc("POST /user/login", userHandler.Login)
	mux.HandleFunc("GET /user/findall", userHandler.FindAllUsers)
	mux.HandleFunc("GET /user/findone", userHandler.FindUser)
	mux.HandleFunc("PUT /user/update", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /user/logout", userHandler.Logout)
	mux.HandleFunc("DELETE /user/delete", userHandler.DeleteUser)

	mux.HandleFunc("GET /token", userHandler.GetToken)

	mux.HandleFunc("POST /content/upload", contentHandler.UploadContent)
	mux.HandleFunc("GET /content/findall", contentHandler.GetAllContent)
	mux.HandleFunc("GET /content/findone", contentHandler.GetOneContent)
	mux.HandleFunc("PUT /content/update", contentHandler.UpdateContent)
	mux.HandleFunc("DELETE /content/delete", contentHandler.DeleteContent)

	mux.HandleFunc("POST /comment/send", commentHandler.SendComment)
	mux.HandleFunc("GET /comment/findall", commentHandler.GetAllComments)
	mux.HandleFunc("PUT /comment/update", commentHandler.UpdateComment)
	mux.HandleFunc("DELETE /comment/delete", commentHandler.DeleteComment)

	mux.HandleFunc("POST /like/create", likeHandler.LikeContent)
	mux.HandleFunc("GET /like/findall", likeHandler.GetAllLikes)
	mux.HandleFunc("GET /like/findone", likeHandler.GetOneLike)
	mux.HandleFunc("DELETE /like/delete", likeHandler.UnlikeContent)

	mux.HandleFunc("POST /save/create", saveHandler.SaveContent)
	mux.HandleFunc("GET /save/findall", saveHandler.GetAllSaves)
	mux.HandleFunc("GET /save/findone", saveHandler.GetOneSave)
	mux.HandleFunc("DELETE /save/delete", saveHandler.UnsaveContent)

	mux.HandleFunc("POST /feedback/create", feedbackHandler.SendFeedback)
	mux.HandleFunc("GET /feedback/findall", feedbackHandler.GetAllFeedbacks)

	mux.HandleFunc("POST /message/create", messageHandler.SendMessage)
	mux.HandleFunc("GET /message/findall", messageHandler.GetAllMessage)
	mux.HandleFunc("PUT /message/update", messageHandler.UpdateMessage)
	mux.HandleFunc("DELETE /message/delete", messageHandler.DeleteMessage)

	mux.HandleFunc("POST /dm/create", directMessageHandler.CreateDirectMessage)
	mux.HandleFunc("GET /dm/findall", directMessageHandler.GetAllDirectMessages)
	mux.HandleFunc("GET /dm/findone", directMessageHandler.GetOneDirectMessage)
	mux.HandleFunc("DELETE /dm/delete", directMessageHandler.DeleteDirectMessage)

	mux.HandleFunc("POST /gc/create", groupChatHandler.CreateGroupChat)
	mux.HandleFunc("GET /gc/findall", groupChatHandler.GetAllGroupChats)
	mux.HandleFunc("GET /gc/findone", groupChatHandler.GetOneGroupChat)
	mux.HandleFunc("PUT /gc/update", groupChatHandler.UpdateGroupChat)
	mux.HandleFunc("DELETE /gc/delete", groupChatHandler.DeleteGroupChat)

	mux.HandleFunc("POST /aiconv/create", aIConversationHandler.CreateConversation)
	mux.HandleFunc("GET /aiconv/findone", aIConversationHandler.GetConversation)
	mux.HandleFunc("DELETE /aiconv/delete", aIConversationHandler.DeleteConversation)

	mux.HandleFunc("POST /aimessage/create", aiMessageHandler.SendMessage)
	mux.HandleFunc("GET /aimessage/findall", aiMessageHandler.GetAllMessages)
	mux.HandleFunc("PUT /aimessage/update", aiMessageHandler.UpdateMessage)
	mux.HandleFunc("DELETE /aimessage/delete", aiMessageHandler.DeleteMessage)

	mux.HandleFunc("POST /follow/create", followHandler.FollowUser)
	mux.HandleFunc("GET /follow/findall", followHandler.GetAllFollows)
	mux.HandleFunc("GET /follow/findone", followHandler.CheckFollowing)
	mux.HandleFunc("DELETE /follow/delete", followHandler.UnfollowUser)

	mux.HandleFunc("POST /notification/create", notificationHandler.CreateNotification)
	mux.HandleFunc("GET /notification/findall", notificationHandler.GetAllNotifications)
	mux.HandleFunc("PUT /notification/update", notificationHandler.UpdateNotifications)
	mux.HandleFunc("DELETE /notification/delete", notificationHandler.DeleteNotifications)

	return stack(mux)
}
