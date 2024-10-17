package usecase

import "github.com/ThalesMonteir0/entregue-ja/internal/entity"

type UserDTO struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	TypeUserID     int    `json:"type_user_id"`
	OrganizationID int    `json:"organization_id"`
}

type UserDTOResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	TypeUserID     int    `json:"type_user_id"`
	OrganizationID int    `json:"organization_id"`
}

type CreateUserUseCase struct {
	Repository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(repository entity.UserRepositoryInterface) CreateUserUseCase {
	return CreateUserUseCase{
		repository,
	}
}

func (c *CreateUserUseCase) Execute(userDTO UserDTO) (UserDTOResponse, error) {
	user, err := entity.NewUser(
		userDTO.OrganizationID,
		userDTO.TypeUserID,
		userDTO.Name,
		userDTO.Password,
		userDTO.Role,
		userDTO.Email,
	)
	if err != nil {
		return UserDTOResponse{}, err
	}

	if err = user.EncryptPassword(); err != nil {
		return UserDTOResponse{}, err
	}

	if err = c.Repository.Save(user); err != nil {
		return UserDTOResponse{}, err
	}

	UserDTOR := UserDTOResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Role:           user.Role,
		TypeUserID:     user.TypeUserID,
		OrganizationID: user.OrganizationID,
	}

	return UserDTOR, nil
}
