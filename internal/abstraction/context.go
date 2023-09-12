package abstraction

import (
    "net/http"
    // "github.com/gorilla/mux"
    "gorm.io/gorm"
)

type Context struct {
    ResponseWriter http.ResponseWriter
    Request        *http.Request
    Auth           *AuthContext
    Trx            *TrxContext
}

type AuthContext struct {
    ID        int
    Name      string
    Username  string
    Email     string
    CompanyID int
    RoleID    int
}

type TrxContext struct {
    Db *gorm.DB
}
