
package models

type Manager struct {
    ID int `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func CheckManager(username, password string) bool {
    var manager Manager
    db.Select("id").Where(Manager{Username : username, Password : password}).First(&manager)
    if manager.ID > 0 {
        return true
    }

    return false
}
