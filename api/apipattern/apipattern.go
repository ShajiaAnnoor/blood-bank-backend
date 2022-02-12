package apipattern

//BloodReqCreate holds the api string for creating a blood request
const BloodReqCreate string = "/api/v1/bloodrequest/create"

//BloodReqUpdate holds the api string for updating a blood request
const BloodReqUpdate string = "/api/v1/bloodrequest/update"

//BloodReqGet holds the api string for getting a blood request
const BloodReqGet string = "/api/v1/bloodrequest/get"

//DonorCreate holds the api string for creating a donor
const DonorCreate string = "/api/v1/comment/create"

//DonorRead holds the api string for reading comments
const DonorRead string = "/api/v1/comment/get/{id}"

//DonorUpdate holds the api string for updating comment
const DonorUpdate string = "/api/v1/comment/update"

//PublicUserList holds the api for giving public user list
const PublicUserList string = "/api/v1/public/alluser"

//UserList holds the api for giving user list
const UserList string = "/api/v1/alluser"

//LoginUser holds the api for logging user
const LoginUser string = "/api/v1/users/login"

//RegistrationToken holds holds the api for giving registration token
const RegistrationToken string = "/api/v1/registration/token"

//UserRegistration holds the api for registering a user
const UserRegistration string = "/api/v1/users/registration"

//UserSearch holds the api for searching user
const UserSearch string = "/api/v1/users/search"

//OtherUser holds the api for giving other people's profile information
const OtherUser string = "/api/v1/users/me2/{id}"

//ProfileUpdate holds the api for updating profile information
const ProfileUpdate string = "/api/v1/users/me"

//ProfileUser holds the api for giving information about a user's profile
const ProfileUser string = "/api/v1/users/me"

//StatusCreate holds the api for creating a status by a user
const StatusCreate string = "/api/v1/status/create"

//StatusByUser holds the api for giving status by user id
const StatusByUser string = "/api/v1/status/getbyuser/{id}"

//NewsFeed holds the api for giving news feed given user id
const NewsFeed string = "/api/v1/newsfeed/user/{id}"

//StatusRead holds the api for reading status by id
const StatusRead string = "/api/v1/status/get/{id}"

//StatusUpdate holds the api for updating status
const StatusUpdate string = "/api/v1/status/update"

//PictureDownload holds the api for downloading picture
const PictureDownload string = "/api/v1/pictures/{id}"

//PictureList holds the api for listing pictures
const PictureList string = "/api/v1/pictures/list"

//ProfilePictureSet holds the api for setting profile picture
const ProfilePictureSet string = "/api/v1/profile_pic"

//ProfilePictureUpload holds the api for uploading profile picture
const ProfilePictureUpload string = "/api/v1/profilePics"

//PictureUpload holds the api for uploading a picture
const PictureUpload string = "/api/v1/pictures"
