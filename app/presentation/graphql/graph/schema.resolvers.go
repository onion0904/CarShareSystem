package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"log"

	"github.com/onion0904/app/config"
	errDomain "github.com/onion0904/app/domain/error"
	domain_event "github.com/onion0904/app/domain/event"
	mail_Service "github.com/onion0904/app/infrastructure/mail"
	repo "github.com/onion0904/app/infrastructure/repository"
	"github.com/onion0904/app/presentation/graphql/graph/model"
	usecase_event "github.com/onion0904/app/usecase/event"
	usecase_group "github.com/onion0904/app/usecase/group"
	usecase_mail "github.com/onion0904/app/usecase/mail"
	usecase_user "github.com/onion0904/app/usecase/user"
	"github.com/onion0904/go-pkg/jwt"
	VerifiedCode "github.com/onion0904/go-pkg/verified_code"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	userRepo := repo.NewUserRepository(r.DB)
	create := usecase_user.NewSaveUserUseCase(userRepo)

	DTO := usecase_user.SaveUseCaseDto{
		LastName:  input.LastName,
		FirstName: input.FirstName,
		Email:     input.Email,
		Password:  input.Password,
		Icon:      input.Icon,
	}

	user, err := create.Run(ctx, DTO)
	if err != nil {
		return nil, err
	}
	nuser := model.User{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		Icon:      user.Icon,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		GroupIDs:  user.GroupIDs,
		EventIDs:  user.EventIDs,
	}
	return &nuser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	userRepo := repo.NewUserRepository(r.DB)
	update := usecase_user.NewUpdateUserUseCase(userRepo)
	DTO := usecase_user.UpdateUseCaseDto{
		LastName:  *input.LastName,
		FirstName: *input.FirstName,
		Email:     *input.Email,
		Icon:      *input.Icon,
	}
	user, err := update.Run(ctx, id, DTO)
	if err != nil {
		return nil, err
	}
	nuser := model.User{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		Icon:      user.Icon,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		GroupIDs:  user.GroupIDs,
		EventIDs:  user.EventIDs,
	}
	return &nuser, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	userRepo := repo.NewUserRepository(r.DB)
	delete := usecase_user.NewDeleteUseCase(userRepo)
	err := delete.Run(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	create := usecase_group.NewSaveUseCase(groupRepo)

	DTO := usecase_group.SaveUseCaseDto{
		Name: input.Name,
		Icon: input.Icon,
	}

	group, err := create.Run(ctx, DTO)
	if err != nil {
		return nil, err
	}
	ngroup := model.Group{
		ID:        group.ID(),
		Name:      group.Name(),
		Icon:      group.Icon(),
		CreatedAt: group.CreatedAt(),
		UpdatedAt: group.UpdatedAt(),
		UserIDs:   group.UserIDs(),
		EventIDs:  group.EventIDs(),
	}
	return &ngroup, nil
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id string, input model.UpdateGroupInput) (*model.Group, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	update := usecase_group.NewUpdateUseCase(groupRepo)
	DTO := usecase_group.UpdateUseCaseDto{
		Name: *input.Name,
		Icon: *input.Icon,
	}
	group, err := update.Run(ctx, id, DTO)
	if err != nil {
		return nil, err
	}
	ngroup := model.Group{
		ID:        group.ID(),
		Name:      group.Name(),
		Icon:      group.Icon(),
		CreatedAt: group.CreatedAt(),
		UpdatedAt: group.UpdatedAt(),
		UserIDs:   group.UserIDs(),
		EventIDs:  group.EventIDs(),
	}
	return &ngroup, nil
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (bool, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	delete := usecase_group.NewDeleteUseCase(groupRepo)
	err := delete.Run(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddUserToGroup is the resolver for the addUserToGroup field.
func (r *mutationResolver) AddUserToGroup(ctx context.Context, groupID string, userID string) (*model.Group, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	addUser := usecase_group.NewAddUserToGroupUseCase(groupRepo)
	DTO := usecase_group.AddUserToGroupUseCaseDto{
		UserID:  userID,
		GroupID: groupID,
	}
	group, err := addUser.Run(ctx, DTO)
	if err != nil {
		return nil, err
	}
	ngroup := model.Group{
		ID:        group.ID(),
		Name:      group.Name(),
		Icon:      group.Icon(),
		CreatedAt: group.CreatedAt(),
		UpdatedAt: group.UpdatedAt(),
		UserIDs:   group.UserIDs(),
		EventIDs:  group.EventIDs(),
	}
	return &ngroup, nil
}

// RemoveUserFromGroup is the resolver for the removeUserFromGroup field.
func (r *mutationResolver) RemoveUserFromGroup(ctx context.Context, groupID string, userID string) (*model.Group, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	removeUser := usecase_group.NewRemoveUserToGroupUseCase(groupRepo)
	DTO := usecase_group.RemoveUserFromGroupUseCaseDto{
		UserID:  userID,
		GroupID: groupID,
	}
	group, err := removeUser.Run(ctx, DTO)
	if err != nil {
		return nil, err
	}
	ngroup := model.Group{
		ID:        group.ID(),
		Name:      group.Name(),
		Icon:      group.Icon(),
		CreatedAt: group.CreatedAt(),
		UpdatedAt: group.UpdatedAt(),
		UserIDs:   group.UserIDs(),
		EventIDs:  group.EventIDs(),
	}
	return &ngroup, nil
}

// AddEventToGroup is the resolver for the addEventToGroup field.
func (r *mutationResolver) AddEventToGroup(ctx context.Context, groupID string, eventID string) (*model.Group, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	addEvent := usecase_group.NewAddEventToGroupUseCase(groupRepo)
	DTO := usecase_group.AddEventToGroupUseCaseDto{
		EventID: eventID,
		GroupID: groupID,
	}
	group, err := addEvent.Run(ctx, DTO)
	if err != nil {
		return nil, err
	}
	ngroup := model.Group{
		ID:        group.ID(),
		Name:      group.Name(),
		Icon:      group.Icon(),
		CreatedAt: group.CreatedAt(),
		UpdatedAt: group.UpdatedAt(),
		UserIDs:   group.UserIDs(),
		EventIDs:  group.EventIDs(),
	}
	return &ngroup, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEventInput) (*model.Event, error) {
	eventRepo := repo.NewEventRepository(r.DB)
	create := usecase_event.NewEventUseCase(domain_event.NewEventDomainService(eventRepo))
	DTO := usecase_event.AddEventUseCaseDTO{
		UsersID:     input.UserID,
		Together:    input.Together,
		Description: input.Description,
		Important:   input.Important,
	}
	event, err := create.Run(ctx, DTO)
	if err != nil {
		return nil, err
	}
	nevent := model.Event{
		ID:          event.ID(),
		UserID:      event.UserID(),
		Together:    event.Together(),
		Description: event.Description(),
		Year:        event.Year(),
		Month:       event.Month(),
		Day:         event.Day(),
		Date:        event.Date(),
		CreatedAt:   event.CreatedAt(),
		UpdatedAt:   event.UpdatedAt(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
		Important:   event.Important(),
	}
	return &nevent, nil
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	eventRepo := repo.NewEventRepository(r.DB)
	delete := usecase_event.NewDeleteUseCase(eventRepo)
	err := delete.Run(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SendVerificationCode is the resolver for the sendVerificationCode field.
func (r *mutationResolver) SendVerificationCode(ctx context.Context, email string) (bool, error) {
	mailService := mail_Service.NewMailRepository()
	send := usecase_mail.NewSendEmailUseCase(mailService)
	if email == "" {
		return false, errDomain.NewError("メールアドレスが必要です")
	}

	// 認証コードの生成
	vcode, err := VerifiedCode.GenerateVerificationCode()
	if err != nil {
		log.Printf("Error generating verification code: %v", err)
		return false, errDomain.NewError("認証コードの生成に失敗しました")
	}

	// 認証コードを保存
	send.CodeMutex.Lock()
	send.VerificationCodes[email] = vcode
	send.CodeMutex.Unlock()

	// メール送信
	DTO := usecase_mail.SendEmailUseCaseDto{
		Email: email,
		Code:  vcode,
	}
	send.Run(ctx, DTO)
	return true, nil
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.CreateUserInput, vcode string) (*model.AuthUserResponse, error) {
	mailService := mail_Service.NewMailRepository()
	send := usecase_mail.NewSendEmailUseCase(mailService)
	c := config.GetConfig()
	secret := c.JWT.Secret
	if secret == "" {
		log.Printf("JWT secret key is not set")
		return nil, errDomain.NewError("JWT secret key is not set")
	}
	jwtSecret := []byte(secret)

	if input.Email == "" || input.Password == "" || vcode == "" {
		return nil, errDomain.NewError("Email or Password or verified code is not set")
	}
	userRepo := repo.NewUserRepository(r.DB)
	exist, err := userRepo.ExistUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errDomain.NewError("User is already registered")
	}

	send.CodeMutex.Lock()
	expectedCode, exists := send.VerificationCodes[input.Email]
	send.CodeMutex.Unlock()

	if !exists {
		return nil, errDomain.NewError("verified code is not found")
	}

	if expectedCode != vcode {
		return nil, errDomain.NewError("Invalid verified code")
	}

	send.CodeMutex.Lock()
	delete(send.VerificationCodes, input.Email)
	send.CodeMutex.Unlock()

	user, err := r.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	customClaim := jwt.NewCustomClaims(user.Email, user.ID)
	token := jwt.CreateToken(customClaim)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	return &model.AuthUserResponse{
		Token: tokenString,
		User:  user,
	}, nil
}

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, email string, password string) (*model.AuthUserResponse, error) {
	c := config.GetConfig()
	secret := c.JWT.Secret
	if secret == "" {
		log.Printf("JWT secret key is not set")
		return nil, errDomain.NewError("JWT secret key is not set")
	}
	jwtSecret := []byte(secret)

	if email == "" || password == "" {
		return nil, errDomain.NewError("Email or Password or verified code is not set")
	}
	userRepo := repo.NewUserRepository(r.DB)
	exist, err := userRepo.ExistUser(ctx, email, password)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errDomain.NewError("User is already registered")
	}

	find := usecase_user.NewFindUserByEmailPasswordUseCase(userRepo)
	user, err := find.Run(ctx, email, password)
	if err != nil {
		return nil, err
	}
	nuser := model.User{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		Icon:      user.Icon,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		GroupIDs:  user.GroupIDs,
		EventIDs:  user.EventIDs,
	}

	customClaim := jwt.NewCustomClaims(email, user.ID)
	token := jwt.CreateToken(customClaim)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	return &model.AuthUserResponse{
		Token: tokenString,
		User:  &nuser,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userRepo := repo.NewUserRepository(r.DB)
	find := usecase_user.NewFindUserUseCase(userRepo)
	user, err := find.Run(ctx, id)
	if err != nil {
		return nil, err
	}
	nuser := model.User{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		Icon:      user.Icon,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		GroupIDs:  user.GroupIDs,
		EventIDs:  user.EventIDs,
	}
	return &nuser, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, id string) (*model.Group, error) {
	groupRepo := repo.NewGroupRepository(r.DB)
	find := usecase_group.NewFindGroupUseCase(groupRepo)
	group, err := find.Run(ctx, id)
	if err != nil {
		return nil, err
	}
	ngroup := model.Group{
		ID:        group.ID,
		Name:      group.Name,
		Icon:      group.Icon,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
		UserIDs:   group.UserIDs,
		EventIDs:  group.EventIDs,
	}
	return &ngroup, nil
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	eventRepo := repo.NewEventRepository(r.DB)
	find := usecase_event.NewFindEventUseCase(eventRepo)
	event, err := find.Run(ctx, id)
	if err != nil {
		return nil, err
	}
	nevent := model.Event{
		ID:          event.ID,
		UserID:      event.UserID,
		Together:    event.Together,
		Description: event.Description,
		Year:        event.Year,
		Month:       event.Month,
		Day:         event.Day,
		Date:        event.Date,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Important:   event.Important,
	}
	return &nevent, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented: Events - events"))
}

// EventsByMonth is the resolver for the eventsByMonth field.
func (r *queryResolver) EventsByMonth(ctx context.Context, input model.MonthlyEventInput) ([]string, error) {
	eventRepo := repo.NewEventRepository(r.DB)
	find := usecase_event.NewFindMonthEventUseCase(eventRepo)
	events, err := find.Run(ctx, input.Year, input.Month)
	if err != nil {
		return nil, err
	}
	return events.EventIDs, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
