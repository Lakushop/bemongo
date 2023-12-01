package bemongo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllUser(mconn, collectionname)
	return GCFReturnStruct(datagedung)
}

func GCFFindUserByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	user := FindUser(mconn, collectionname, datauser)
	return GCFReturnStruct(user)
}

func GCFFindUserByName(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}

	// Jika username kosong, maka respon "false" dan data tidak ada
	if datauser.Username == "" {
		return "false"
	}

	// Jika ada username, mencari data pengguna
	user := FindUserUser(mconn, collectionname, datauser)

	// Jika data pengguna ditemukan, mengembalikan data pengguna dalam format yang sesuai
	if user != (User{}) {
		return GCFReturnStruct(user)
	}

	// Jika tidak ada data pengguna yang ditemukan, mengembalikan "false" dan data tidak ada
	return "false"
}

func GCFDeleteHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	DeleteUser(mconn, collectionname, datauser)
	return GCFReturnStruct(datauser)
}

func GCFUpdateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	ReplaceOneDoc(mconn, collectionname, bson.M{"username": datauser.Username}, datauser)
	return GCFReturnStruct(datauser)
}

func GCFCreateHandlerTokenPaseto(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword
	CreateNewUserRole(mconn, collectionname, datauser)
	tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return err.Error()
	}
	datauser.Token = tokenstring
	return GCFReturnStruct(datauser)
}

func GCFCreateAccountAndToken(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword
	CreateUserAndAddedToeken(PASETOPRIVATEKEYENV, mconn, collectionname, datauser)
	return GCFReturnStruct(datauser)
}
func GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}

	// Hash the password before storing it
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword

	createErr := CreateNewUserRole(mconn, collectionname, datauser)
	fmt.Println(createErr)

	return GCFReturnStruct(datauser)
}

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
		var Response Credential
		Response.Status = false
		mconn := SetConn(MONGOCONNSTRINGENV, dbname)
		var datauser User
		err := json.NewDecoder(r.Body).Decode(&datauser)
		if err != nil {
			Response.Message = "error parsing application/json: " + err.Error()
		} else {
			if IsPasswordValid(mconn, collectionname, datauser) {
				Response.Status = true
				tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
				if err != nil {
					Response.Message = "Gagal Encode Token : " + err.Error()
				} else {
					Response.Message = "Selamat Datang"
					Response.Token = tokenstring
				}
			} else {
				Response.Message = "Password Salah"
			}
		}

		return GCFReturnStruct(Response)
	}

	func GCFReturnStruct(DataStuct interface{}) string {
		jsondata, _ := json.Marshal(DataStuct)
		return string(jsondata)
	}

	// product
	func GCFGetAllProduct(MONGOCONNSTRINGENV, dbname, collectionname string) string {
		mconn := SetConn(MONGOCONNSTRINGENV, dbname)
		datagedung := GetAllProduct(mconn, collectionname)
		return GCFReturnStruct(datagedung)
	}

	func GCFGetAllContentBy(MONGOCONNSTRINGENV, dbname, collectionname string) string {
		mconn := SetConn(MONGOCONNSTRINGENV, dbname)
		datacontent := GetAllContent(mconn, collectionname)
		return GCFReturnStruct(datacontent)
	}

	func GCFCreateProduct(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) Credential {
		var Response Credential
		// ...
		return Response
	}

	Response.Status = false

	// Retrieve the "PUBLICKEY" from the request headers
	publicKey := r.Header.Get("PUBLICKEY")
	if publicKey == "" {
		Response.Message = "Missing PUBLICKEY in headers"
	} else {
		// Process the request with the "PUBLICKEY"
		mconn := SetConn(MONGOCONNSTRINGENV, dbname)
		var dataproduct Product
		err := json.NewDecoder(r.Body).Decode(&dataproduct)
		if err != nil {
			Response.Message = "Error parsing application/json: " + err.Error()
		} else {
			CreateNewProduct(mconn, dbname, Product{
				NomorID:     dataproduct.NomorID,
				Name:        dataproduct.Name,
				Description: dataproduct.Description,
				Price:       dataproduct.Price,
				Stock:       dataproduct.Stock,
				Size:        dataproduct.Size,
				Image:       dataproduct.Image,
			})
			Response.Status = true
			Response.Message = "Berhasil"
			// No token generation here
		}
	}
	return Response
}

func GCFLoginTest(username, password, MONGOCONNSTRINGENV, dbname, collectionname string) bool {
	// Membuat koneksi ke MongoDB
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)

	// Mencari data pengguna berdasarkan username
	filter := bson.M{"username": username}
	collection := collectionname
	res := atdb.GetOneDoc[User](mconn, collection, filter)

	// Memeriksa apakah pengguna ditemukan dalam database
	if res == (User{}) {
		return false
	}

	// Memeriksa apakah kata sandi cocok
	return CheckPasswordHash(password, res.Password)
}

func InsertDataUserGCF(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	userdata := new(User)
	resp.Status = false
	conn := SetConn(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := HashPassword(userdata.Password)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(conn, userdata.Username, userdata.Role, hash)
		resp.Message = "Berhasil Input data"
	}
	return GCFReturnStruct(resp)
}

// Content

func GCFCreateContent(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var datacontent Content
	err := json.NewDecoder(r.Body).Decode(&datacontent)
	if err != nil {
		return err.Error()
	}

	CreateNewContent(mconn, collectionname, datacontent)
	// setelah create content munculkan response berhasil dan 200

	if CreateResponse(true, "Berhasil", datacontent) != (Response{}) {
		return GCFReturnStruct(CreateResponse(true, "success Create Data Content", datacontent))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Create Data Content", datacontent))
	}
}

func GCFDeleteHandlerContent(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var contentdata Content
	err := json.NewDecoder(r.Body).Decode(&contentdata)
	if err != nil {
		return err.Error()
	}
	DeleteContent(mconn, collectionname, contentdata)
	return GCFReturnStruct(contentdata)
}

func GCFUpdatedContent(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var contentdata Content
	err := json.NewDecoder(r.Body).Decode(&contentdata)
	if err != nil {
		return err.Error()
	}
	ReplaceContent(mconn, collectionname, bson.M{"id": contentdata.ID}, contentdata)
	return GCFReturnStruct(contentdata)
}

func GCFCreateNewBlog(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var blogdata Blog
	err := json.NewDecoder(r.Body).Decode(&blogdata)
	if err != nil {
		return err.Error()
	}
	CreateNewBlog(mconn, collectionname, blogdata)
	return GCFReturnStruct(blogdata)
}

func GCFFindContentAllID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)

	// Inisialisasi variabel datacontent
	var datacontent Content

	// Membaca data JSON dari permintaan HTTP ke dalam datacontent
	err := json.NewDecoder(r.Body).Decode(&datacontent)
	if err != nil {
		return err.Error()
	}

	// Memanggil fungsi FindContentAllId
	content := FindContentAllId(mconn, collectionname, datacontent)

	// Mengembalikan hasil dalam bentuk JSON
	return GCFReturnStruct(content)
}

func GCFFindBlogAllID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)

	// Inisialisasi variabel datacontent
	var datablog Blog

	// Membaca data JSON dari permintaan HTTP ke dalam datacontent
	err := json.NewDecoder(r.Body).Decode(&datablog)
	if err != nil {
		return err.Error()
	}

	// Memanggil fungsi FindContentAllId
	blog := GetIDBlog(mconn, collectionname, datablog)

	// Mengembalikan hasil dalam bentuk JSON
	return GCFReturnStruct(blog)
}

func GCFGetAllBlog(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	datablog := GetAllBlogAll(mconn, collectionname)
	return GCFReturnStruct(datablog)
}

func GCFCreateTokenAndSaveToDB(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) (string, error) {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)

	// Inisialisasi variabel datauser
	var datauser User

	// Membaca data JSON dari permintaan HTTP ke dalam datauser
	if err := json.NewDecoder(r.Body).Decode(&datauser); err != nil {
		return "", err // Mengembalikan kesalahan langsung
	}

	// Generate a token for the user
	tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return "", err // Mengembalikan kesalahan langsung
	}
	datauser.Token = tokenstring

	// Simpan pengguna ke dalam basis data
	if err := atdb.InsertOneDoc(mconn, collectionname, datauser); err != nil {
		return tokenstring, nil // Mengembalikan kesalahan langsung
	}

	return tokenstring, nil // Mengembalikan token dan nil untuk kesalahan jika sukses
}
func GCFCreteRegister(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	CreateUser(mconn, collectionname, userdata)
	return GCFReturnStruct(userdata)
}

func GCFLoginAfterCreate(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		tokenstring, err := watoken.Encode(userdata.Username, os.Getenv("PASETOPRIVATEKEYENV"))
		if err != nil {
			return err.Error()
		}
		userdata.Token = tokenstring
		return GCFReturnStruct(userdata)
	} else {
		return "Password Salah"
	}
}

func GCFLoginAfterCreater(MONGOCONNSTRINGENV, dbname, collectionname, privateKeyEnv string, r *http.Request) (string, error) {
	// Ambil data pengguna dari request, misalnya dari body JSON atau form data.
	var userdata User
	// Implement the logic to extract user data from the request (r) here.

	mconn := SetConn(MONGOCONNSTRINGENV, dbname)

	// Lakukan otentikasi pengguna yang baru saja dibuat.
	token, err := AuthenticateUserAndGenerateToken(privateKeyEnv, mconn, collectionname, userdata)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GCFLoginAfterCreatee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		// Password is valid, return a success message or some other response.
		return "Login successful"

	} else {
		// Password is not valid, return an error message.
		return "Password Salah"
	}
}

func GCFLoginAfterCreateee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		// Password is valid, construct and return the GCFReturnStruct.
		response := CreateResponse(true, "Berhasil Login", userdata)
		return GCFReturnStruct(response) // Return GCFReturnStruct directly
	} else {
		// Password is not valid, return an error message.
		return "Password Salah"
	}
}
func GCFLoginAfterCreateeee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		// Password is valid, return a success message or some other response.
		return GCFReturnStruct(userdata)
	} else {
		// Password is not valid, return an error message.
		return "Password Salah"
	}
}

func GCFCreteCommnet(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var commentdata Comment
	err := json.NewDecoder(r.Body).Decode(&commentdata)
	if err != nil {
		return err.Error()
	}

	if err := CreateComment(mconn, collectionname, commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Succes Create Comment", commentdata))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Create Comment", commentdata))
	}
}

func GCFGetAllComment(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	datacomment := GetAllComment(mconn, collectionname)
	if datacomment != nil {
		return GCFReturnStruct(CreateResponse(true, "success Get All Comment", datacomment))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Get All Comment", datacomment))
	}
}
func GFCUpadatedCommnet(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var commentdata Comment
	err := json.NewDecoder(r.Body).Decode(&commentdata)
	if err != nil {
		return err.Error()
	}

	if err := UpdatedComment(mconn, collectionname, commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Updated Comment", commentdata))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Updated Comment", commentdata))
	}
}

func GCFDeletedCommnet(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConn(MONGOCONNSTRINGENV, dbname)
	var commentdata Comment
	if err := json.NewDecoder(r.Body).Decode(&commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(false, "Failed to process request", commentdata))
	}

	if err := DeleteComment(mconn, collectionname, commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Successfully deleted comment", commentdata))
	}

	return GCFReturnStruct(CreateResponse(false, "Failed to delete comment", commentdata))
}
