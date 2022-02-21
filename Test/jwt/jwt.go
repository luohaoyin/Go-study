package jwt

/*import (
	"errors"
	"time"
)

type MyClaims struct {
	Username  string   `json:"username"`
	jwt.StandardClaims
}

var mySecret = []byte("巴拉拉小能量~")

func keyFunc(_ *jwt.Token)(i interface{},err  error){

	return mySecret,nil
}
//自定义过期时间
const  TokenExpireDuration = time.Hour *1

func GenToken(username string)(aToken,rToken string,err error){
	c := MyClaims{
		 Username:"username",
		 StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:  "bluebell",
		},
	}
	//aToken,err = jwt.NewWithClaims(jwt.SigningMewhodJS256.c).SignedString(mySecret)
	//rToken,err = jwt.NewWithClaims(jwt.SigninggMethodHS256,jwt.StandrdClaims{
	//	ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
	//	Issuer:  "bluebell",
	//}).SignedString(mySecret)
	return
}

func ParseToken(tokenString string)(Claims *McClaims,err error){
	var token *jwt.Token
	Claims = new(MyClaims)
	token,err = jwt.ParseWithClaim(tokenString,Claims,keyFunc){

		if err != nil{
			return
		}
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}
    return
}

func RefreshToken(atoken,rtoken string)(newAToken,nweEToken string,err error){

	if _,err = jwt.Parse(rToken,keyFunc);err !=nil{
         return
	}
    var claims MyClaims
	_,err = jwt.ParseWithClaims(aToken,&claims,keyFunc)
	v,_ := err.(*jwt.ValidationError)
	return
}
*/

