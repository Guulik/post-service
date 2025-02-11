package constants

const (
	WrongPageError          = "wrong page number"
	WrongPageSizeError      = "wrong page size"
	WrongIdError            = "wrong id"
	TooLongContentError     = "content is too long"
	CommentsNotAllowedError = "comments not allowed for this post"

	CreatingPostError = "failed to create post"
	GetPostError      = "failed to get post"
	PostNotFoundError = "post not found"

	CreatingCommentError = "failed to create new comment"
	GettingCommentError  = "failed to get comments"
	GettingRepliesError  = "failed to get replies"
)

const (
	InternalErrorType = "500 Internal Server Error"
	BadRequestType    = "400 Bad Request"
	NotFoundType      = "404 Not Found Error"
)

const (
	ThereIsNoObserversError = "there is no connection to the observer"
)
