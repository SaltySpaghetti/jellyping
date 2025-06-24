package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/jackc/pgx/v5"
	repository "saltyspaghetti.dev/jellyping/repositories"
)

type UserService struct {
	ctx  context.Context
	conn *pgx.Conn
}

func NewUserService(ctx context.Context, conn *pgx.Conn) *UserService {
	return &UserService{ctx: ctx, conn: conn}
}

func (s *UserService) SetChatId(username string, chatId int64) (repository.User, error) {
	repo := repository.New(s.conn)

	user, err := repo.GetUserByUsername(s.ctx, username)
	if err != nil {
		return repository.User{}, errors.New("user not found")
	}

	params := repository.UpdateChatIdParams{
		Username: username,
		ChatID: pgtype.Int8{
			Int64: chatId,
			Valid: true,
		},
	}
	user, err = repo.UpdateChatId(s.ctx, params)
	if err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]repository.User, error) {
	repo := repository.New(s.conn)

	users, err := repo.ListUsers(s.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetByUsername(username string) (repository.User, error) {
	repo := repository.New(s.conn)
	user, err := repo.GetUserByUsername(s.ctx, username)
	if err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func (s *UserService) GetByChatId(chatId int64) (repository.User, error) {
	repo := repository.New(s.conn)
	user, err := repo.GetUserByChatId(s.ctx, pgtype.Int8{Int64: chatId})
	if err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func (s *UserService) CreateUser(username string, chatId int64) (repository.User, error) {
	repo := repository.New(s.conn)

	params := repository.CreateUserParams{
		Username: username,
		ChatID:   pgtype.Int8{Int64: chatId},
	}
	user, err := repo.CreateUser(s.ctx, params)
	if err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(chatId int64, username string) (repository.User, error) {
	repo := repository.New(s.conn)

	params := repository.UpdateChatIdParams{
		ChatID:   pgtype.Int8{Int64: chatId},
		Username: username,
	}
	user, err := repo.UpdateChatId(s.ctx, params)
	if err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(chatId int64) error {
	repo := repository.New(s.conn)
	err := repo.DeleteUser(s.ctx, pgtype.Int8{Int64: chatId, Valid: true})
	if err != nil {
		return err
	}

	return nil
}
