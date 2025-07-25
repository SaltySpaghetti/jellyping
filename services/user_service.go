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

func (s *UserService) UserExists(username string) (bool, error) {
	repo := repository.New(s.conn)

	exists, err := repo.UserExists(s.ctx, username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *UserService) SetChatId(username string, chatId int64) (repository.User, error) {
	repo := repository.New(s.conn)

	_, err := repo.GetUserByUsername(s.ctx, username)
	if err != nil {
		return repository.User{}, errors.New("user not found")
	}

	existingUser, err := repo.GetUserByChatId(s.ctx, pgtype.Int8{Int64: chatId, Valid: true})
	if err == nil && existingUser.Username != username {
		// Set chatId to null for the existing user
		nullParams := repository.UpdateChatIdParams{
			Username: existingUser.Username,
			ChatID:   pgtype.Int8{Valid: false},
		}
		_, err = repo.UpdateChatId(s.ctx, nullParams)
		if err != nil {
			return repository.User{}, err
		}
	}

	params := repository.UpdateChatIdParams{
		Username: username,
		ChatID: pgtype.Int8{
			Int64: chatId,
			Valid: true,
		},
	}
	user, err := repo.UpdateChatId(s.ctx, params)
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

func (s *UserService) CreateUser(username string) (repository.User, error) {
	repo := repository.New(s.conn)

	user, err := repo.CreateUser(s.ctx, username)
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
