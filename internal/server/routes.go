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

	contentRepository := content.NewContentRepository(s.DB)
	contentService := content.NewContentService(contentRepository)
	contentHandler := content.NewHandler(contentService, s.S3Client, s.BucketName)

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

	mux.HandleFunc("POST /user/register", userHandler.Register)
	mux.HandleFunc("OPTIONS /user/register", userHandler.Register)
	mux.HandleFunc("POST /user/login", userHandler.Login)
	mux.HandleFunc("OPTIONS /user/login", userHandler.Login)
	mux.HandleFunc("GET /user/findall/{username}", userHandler.FindAllUsers)
	mux.HandleFunc("OPTIONS /user/findall/{username}", userHandler.FindAllUsers)
	mux.HandleFunc("GET /user/mutual/{userId}", userHandler.FindAllMutualUsers)
	mux.HandleFunc("OPTIONS /user/mutual/{userId}", userHandler.FindAllMutualUsers)
	mux.HandleFunc("GET /user/findone/{username}", userHandler.FindUser)
	mux.HandleFunc("OPTIONS /user/findone/{username}", userHandler.FindUser)
	mux.HandleFunc("PATCH /user/update/{id}", userHandler.UpdateUser)
	mux.HandleFunc("OPTIONS /user/update/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /user/logout/{id}", userHandler.Logout)
	mux.HandleFunc("OPTIONS /user/logout/{id}", userHandler.Logout)
	mux.HandleFunc("DELETE /user/delete/{userId}", userHandler.DeleteUser)
	mux.HandleFunc("OPTIONS /user/delete/{userId}", userHandler.DeleteUser)

	mux.HandleFunc("GET /token", userHandler.GetToken)
	mux.HandleFunc("OPTIONS /token", userHandler.GetToken)

	mux.HandleFunc("POST /content/upload", contentHandler.UploadContent)
	mux.HandleFunc("OPTIONS /content/upload", contentHandler.UploadContent)
	mux.HandleFunc("GET /content/findall", contentHandler.GetAllContent)
	mux.HandleFunc("OPTIONS /content/findall", contentHandler.GetAllContent)
	mux.HandleFunc("GET /content/findall/following", contentHandler.GetAllContentByFollowing)
	mux.HandleFunc("OPTIONS /content/findall/following", contentHandler.GetAllContentByFollowing)
	mux.HandleFunc("GET /content/findone/{id}", contentHandler.GetOneContent)
	mux.HandleFunc("OPTIONS /content/findone/{id}", contentHandler.GetOneContent)
	mux.HandleFunc("PATCH /content/update", contentHandler.UpdateContent)
	mux.HandleFunc("OPTIONS /content/update", contentHandler.UpdateContent)
	mux.HandleFunc("DELETE /content/delete", contentHandler.DeleteContent)
	mux.HandleFunc("OPTIONS /content/delete", contentHandler.DeleteContent)

	mux.HandleFunc("POST /comment/create/{userId}", commentHandler.SendComment)
	mux.HandleFunc("OPTIONS /comment/create/{userId}", commentHandler.SendComment)
	mux.HandleFunc("GET /comment/findall/{contentId}", commentHandler.GetAllComments)
	mux.HandleFunc("OPTIONS /comment/findall/{contentId}", commentHandler.GetAllComments)
	mux.HandleFunc("PATCH /comment/update", commentHandler.UpdateComment)
	mux.HandleFunc("OPTIONS /comment/update", commentHandler.UpdateComment)
	mux.HandleFunc("DELETE /comment/delete/{id}", commentHandler.DeleteComment)
	mux.HandleFunc("OPTIONS /comment/delete/{id}", commentHandler.DeleteComment)

	mux.HandleFunc("POST /like/create", likeHandler.LikeContent)
	mux.HandleFunc("OPTIONS /like/create", likeHandler.LikeContent)
	mux.HandleFunc("GET /like/findall", likeHandler.GetAllLikes)
	mux.HandleFunc("OPTIONS /like/findall", likeHandler.GetAllLikes)
	mux.HandleFunc("GET /like/findone", likeHandler.GetOneLike)
	mux.HandleFunc("OPTIONS /like/findone", likeHandler.GetOneLike)
	mux.HandleFunc("DELETE /like/delete/{id}", likeHandler.UnlikeContent)
	mux.HandleFunc("OPTIONS /like/delete/{id}", likeHandler.UnlikeContent)

	mux.HandleFunc("POST /saved/create", saveHandler.SaveContent)
	mux.HandleFunc("OPTIONS /saved/create", saveHandler.SaveContent)
	mux.HandleFunc("GET /saved/findall/{userId}", saveHandler.GetAllSaves)
	mux.HandleFunc("OPTIONS /saved/findall/{userId}", saveHandler.GetAllSaves)
	mux.HandleFunc("GET /saved/findone", saveHandler.GetOneSave)
	mux.HandleFunc("OPTIONS /saved/findone", saveHandler.GetOneSave)
	mux.HandleFunc("DELETE /saved/delete/{id}", saveHandler.UnsaveContent)
	mux.HandleFunc("OPTIONS /saved/delete/{id}", saveHandler.UnsaveContent)

	mux.HandleFunc("POST /feedback/create/{id}", feedbackHandler.SendFeedback)
	mux.HandleFunc("OPTIONS /feedback/create/{id}", feedbackHandler.SendFeedback)
	mux.HandleFunc("GET /feedback/findall", feedbackHandler.GetAllFeedbacks)
	mux.HandleFunc("OPTIONS /feedback/findall", feedbackHandler.GetAllFeedbacks)

	mux.HandleFunc("POST /message/dm/create/{id}", messageHandler.SendPrivateMessage)
	mux.HandleFunc("OPTIONS /message/dm/create/{id}", messageHandler.SendPrivateMessage)
	mux.HandleFunc("POST /message/group/create/{id}", messageHandler.SendGroupMessage)
	mux.HandleFunc("OPTIONS /message/group/create/{id}", messageHandler.SendGroupMessage)
	mux.HandleFunc("GET /message/findall", messageHandler.GetAllMessage)
	mux.HandleFunc("OPTIONS /message/findall", messageHandler.GetAllMessage)
	mux.HandleFunc("PATCH /message/update", messageHandler.UpdateMessage)
	mux.HandleFunc("OPTIONS /message/update", messageHandler.UpdateMessage)
	mux.HandleFunc("DELETE /message/delete", messageHandler.DeleteMessage)
	mux.HandleFunc("OPTIONS /message/delete", messageHandler.DeleteMessage)

	mux.HandleFunc("POST /dm/create", directMessageHandler.CreateDirectMessage)
	mux.HandleFunc("OPTIONS /dm/create", directMessageHandler.CreateDirectMessage)
	mux.HandleFunc("/ws/dm", directMessageHandler.HandleWebSocket)
	mux.HandleFunc("GET /dm/findall/{userId}", directMessageHandler.GetAllDirectMessages)
	mux.HandleFunc("OPTIONS /dm/findall/{userId}", directMessageHandler.GetAllDirectMessages)
	mux.HandleFunc("GET /dm/findone", directMessageHandler.GetOneDirectMessageByParticipants)
	mux.HandleFunc("OPTIONS /dm/findone", directMessageHandler.GetOneDirectMessageByParticipants)
	mux.HandleFunc("GET /dm/findone/{id}", directMessageHandler.GetOneDirectMessage)
	mux.HandleFunc("OPTIONS /dm/findone/{id}", directMessageHandler.GetOneDirectMessage)
	mux.HandleFunc("DELETE /dm/delete/{id}", directMessageHandler.DeleteDirectMessage)
	mux.HandleFunc("OPTIONS /dm/delete/{id}", directMessageHandler.DeleteDirectMessage)

	mux.HandleFunc("POST /gc/create", groupChatHandler.CreateGroupChat)
	mux.HandleFunc("OPTIONS /gc/create", groupChatHandler.CreateGroupChat)
	mux.HandleFunc("POST /gc/members/{groupId}", groupChatHandler.AddMembers)
	mux.HandleFunc("OPTIONS /gc/members/{groupId}", groupChatHandler.AddMembers)
	mux.HandleFunc("/ws/gc", groupChatHandler.HandleWebSocket)
	mux.HandleFunc("GET /gc/findall/{userId}", groupChatHandler.GetAllGroupChats)
	mux.HandleFunc("OPTIONS /gc/findall/{userId}", groupChatHandler.GetAllGroupChats)
	mux.HandleFunc("GET /gc/findone/{id}", groupChatHandler.GetOneGroupChat)
	mux.HandleFunc("OPTIONS /gc/findone/{id}", groupChatHandler.GetOneGroupChat)
	mux.HandleFunc("PATCH /gc/update", groupChatHandler.UpdateGroupChat)
	mux.HandleFunc("OPTIONS /gc/update", groupChatHandler.UpdateGroupChat)
	mux.HandleFunc("DELETE /gc/{groupId}/members/{userId}", groupChatHandler.LeaveGroupChat)
	mux.HandleFunc("OPTIONS /gc/{groupId}/members/{userId}", groupChatHandler.LeaveGroupChat)
	mux.HandleFunc("DELETE /gc/delete", groupChatHandler.DeleteGroupChat)
	mux.HandleFunc("OPTIONS /gc/delete", groupChatHandler.DeleteGroupChat)

	mux.HandleFunc("POST /ai/conv/create", aIConversationHandler.CreateConversation)
	mux.HandleFunc("OPTIONS /ai/conv/create", aIConversationHandler.CreateConversation)
	mux.HandleFunc("GET /ai/conv/findone/{userId}", aIConversationHandler.GetConversation)
	mux.HandleFunc("OPTIONS /ai/conv/findone/{userId}", aIConversationHandler.GetConversation)
	mux.HandleFunc("DELETE /ai/conv/delete", aIConversationHandler.DeleteConversation)
	mux.HandleFunc("OPTIONS /ai/conv/delete", aIConversationHandler.DeleteConversation)

	mux.HandleFunc("POST /ai/message/create", aiMessageHandler.SendMessage)
	mux.HandleFunc("OPTIONS /ai/message/create", aiMessageHandler.SendMessage)
	mux.HandleFunc("GET /ai/message/findall", aiMessageHandler.GetAllMessages)
	mux.HandleFunc("OPTIONS /ai/message/findall", aiMessageHandler.GetAllMessages)
	// mux.HandleFunc("PATCH /ai/message/update", aiMessageHandler.UpdateMessage)
	// mux.HandleFunc("OPTIONS /ai/message/update", aiMessageHandler.UpdateMessage)
	mux.HandleFunc("DELETE /ai/message/delete", aiMessageHandler.DeleteMessage)
	mux.HandleFunc("OPTIONS /ai/message/delete", aiMessageHandler.DeleteMessage)

	mux.HandleFunc("POST /follow/create", followHandler.FollowUser)
	mux.HandleFunc("OPTIONS /follow/create", followHandler.FollowUser)
	mux.HandleFunc("GET /follow/findall", followHandler.GetAllFollows)
	mux.HandleFunc("OPTIONS /follow/findall", followHandler.GetAllFollows)
	mux.HandleFunc("GET /follow/findone", followHandler.CheckFollowing)
	mux.HandleFunc("OPTIONS /follow/findone", followHandler.CheckFollowing)
	mux.HandleFunc("DELETE /follow/delete/{id}", followHandler.UnfollowUser)
	mux.HandleFunc("OPTIONS /follow/delete/{id}", followHandler.UnfollowUser)

	mux.HandleFunc("POST /notification/create", notificationHandler.CreateNotification)
	mux.HandleFunc("OPTIONS /notification/create", notificationHandler.CreateNotification)
	mux.HandleFunc("GET /notification/findall/{userId}", notificationHandler.GetAllNotifications)
	mux.HandleFunc("OPTIONS /notification/findall/{userId}", notificationHandler.GetAllNotifications)
	mux.HandleFunc("PATCH /notification/update", notificationHandler.UpdateNotifications)
	mux.HandleFunc("OPTIONS /notification/update", notificationHandler.UpdateNotifications)
	mux.HandleFunc("DELETE /notification/delete", notificationHandler.DeleteNotifications)
	mux.HandleFunc("OPTIONS /notification/delete", notificationHandler.DeleteNotifications)

	return stack(mux)
}
