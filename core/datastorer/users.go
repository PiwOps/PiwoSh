package datastorer

func GetUserByID(id uint64) (User, error) {
    var user User 
    result := DB.First(&user, id)
    if result.Error != nil {
        return User{}, result.Error
    }

    return user, nil
}

func CreateUser(id uint64) error {
    user := User{ID: id}
    result := DB.Create(&user)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func UpdateUser(user User) error {
    result := DB.Save(&user)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func DeleteUser(id uint64) error {
    result := DB.Delete(&User{}, id)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func GetAllUsers() ([]User, error) {
    var users []User
    result := DB.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }

    return users, nil
}

func MigrateAllUsers() error {
    return nil
}