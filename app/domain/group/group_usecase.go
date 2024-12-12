package group

type GroupUsecase interface {
	Invite() (link string,err error)
}