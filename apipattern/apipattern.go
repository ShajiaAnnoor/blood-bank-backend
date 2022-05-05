package apipattern

//BloodRequestCreate holds the api string for creating a blood request
const BloodRequestCreate string = "/api/v1/bloodrequest/create"

//BloodRequestUpdate holds the api string for updating a blood request
const BloodRequestUpdate string = "/api/v1/bloodrequest/update"

//BloodRequestRead holds the api string for getting a blood request
const BloodRequestRead string = "/api/v1/bloodrequest/get/{id}"

//BloodRequestDelete holds the api string for deleting a blood request
const BloodRequestDelete string = "/api/v1/bloodrequest/delete"

//DonorCreate holds the api string for creating a donor
const DonorCreate string = "/api/v1/donor/create"

//DonorRead holds the api string for reading comments
const DonorRead string = "/api/v1/donor/get/{id}"

//DonorUpdate holds the api string for updating comment
const DonorUpdate string = "/api/v1/donor/update"

//DonorDelete holds the api string for getting a blood request
const DonorDelete string = "/api/v1/donor/delete"

//NoticeCreate holds the api string for creating a notice
const NoticeCreate string = "/api/v1/notice/create"

//NoticeRead holds the api string for reading notices
const NoticeRead string = "/api/v1/notice/get/{id}"

//NoticeUpdate holds the api string for updating notice
const NoticeUpdate string = "/api/v1/notice/update"

//NoticeDelete holds the api string for getting a notice
const NoticeDelete string = "/api/v1/notice/delete"

//OrganizationCreate holds the api string for creating a organization
const OrganizationCreate string = "/api/v1/organization/create"

//OrganizationRead holds the api string for reading organizations
const OrganizationRead string = "/api/v1/organization/get/{id}"

//OrganizationUpdate holds the api string for updating organization
const OrganizationUpdate string = "/api/v1/organization/update"

//OrganizationDelete holds the api string for getting a organization
const OrganizationDelete string = "/api/v1/organization/delete"

//PatientCreate holds the api string for creating a patient
const PatientCreate string = "/api/v1/patient/create"

//PatientRead holds the api string for reading comments
const PatientRead string = "/api/v1/patient/get/{id}"

//PatientUpdate holds the api string for updating comment
const PatientUpdate string = "/api/v1/patient/update"

//PatientDelete holds the api string for getting a blood request
const PatientDelete string = "/api/v1/patient/delete"

// StaticContentCreate holds the api string for creating a staticcontent
const StaticContentCreate string = "/api/v1/staticcontent/create"

// StaticContentRead holds the api string for reading staticcontents
const StaticContentRead string = "/api/v1/staticcontent/get/{id}"

// StaticContentUpdate holds the api string for updating staticcontent
const StaticContentUpdate string = "/api/v1/staticcontent/update"

// StaticContentDelete holds the api string for getting a staticcontent
const StaticContentDelete string = "/api/v1/staticcontent/delete"

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
