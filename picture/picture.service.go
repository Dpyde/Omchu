package picture

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type PictureService interface {
	// UploadPicToR2(fileHeader multipart.FileHeader, bucketName string) (entity.Picture, error)
	UploadPicsToR2(fileHeader []*multipart.FileHeader, bucketName string) ([]entity.Picture, error)
	SavePicturesSer(id uint, picture []entity.Picture) error
	GetPicsByUserId(id uint) ([]entity.Picture, error)
}

type pictureServiceImpl struct {
	repo PictureRepository
}

func NewPictureService(repo PictureRepository) PictureService {
	return &pictureServiceImpl{repo: repo}
}

// Global R2 client
var R2Client *minio.Client

// Initialize Cloudflare R2
func InitR2() {
	// Replace with your actual Cloudflare R2 credentials
	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY")

	// Avoid printing sensitive credentials
	fmt.Println("Initializing Cloudflare R2 with endpoint:", endpoint)

	if strings.HasPrefix(endpoint, "https://") {
		endpoint = endpoint[len("https://"):]
	}
	// Initialize Minio client (Cloudflare R2 uses S3-compatible API)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true, // Use HTTPS
	})
	if err != nil {
		log.Fatalf("Failed to initialize R2 client: %v", err)
	}

	R2Client = client
	log.Println("✅ Cloudflare R2 Client Initialized")
}

func (s *pictureServiceImpl) UploadPicsToR2(fileHeaders []*multipart.FileHeader, bucketName string) ([]entity.Picture, error) {
	var pictures []entity.Picture

	for _, fileHeader := range fileHeaders {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("Failed to open file: %v", err)
			return nil, err
		}
		defer file.Close()

		// Generate a unique file key
		fileKey := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

		// Upload the file to Cloudflare R2
		_, err = R2Client.PutObject(
			context.Background(),
			bucketName,
			fileKey,
			file,
			fileHeader.Size,
			minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")},
		)
		if err != nil {
			log.Printf("Failed to upload file: %v", err)
			return nil, err
		}

		log.Println("File uploaded successfully:", fileKey)

		// Construct the file URL
		fileURL := fmt.Sprintf("https://%s.r2.cloudflarestorage.com/%s/%s", os.Getenv("ACCOUNT_ID"), bucketName, fileKey)

		// Append the uploaded file info to the slice
		pictures = append(pictures, entity.Picture{
			Url: fileURL,
			Key: fileKey,
		})
	}

	return pictures, nil
}

func (s *pictureServiceImpl) SavePicturesSer(userID uint, pictures []entity.Picture) error {
	for i := range pictures {
		pictures[i].UserID = userID
	}
	return s.repo.SavePicturesToDB(pictures)
}

func (s *pictureServiceImpl) GetPicsByUserId(id uint) ([]entity.Picture, error) {
	pictures, err := s.repo.GetPictureFromDB(id)
	if err != nil {
		return nil, err
	}

	return pictures, nil
}

// func (s *pictureServiceImpl) UploadPicToR2(fileHeader multipart.FileHeader, bucketName string) (entity.Picture, error) {
// 	// Open the file
// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		log.Printf("Failed to open file: %v", err)
// 		return entity.Picture{}, err
// 	}
// 	defer file.Close()

// 	fileKey := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)
// 	contentType := "Picture/png" // Modify based on the actual file type
// 	_, err = R2Client.PutObject(
// 		context.Background(),
// 		bucketName,
// 		fileKey,
// 		file,
// 		fileHeader.Size,
// 		minio.PutObjectOptions{ContentType: contentType},
// 	)
// 	if err != nil {
// 		log.Printf("Failed to upload file: %v", err)
// 		return entity.Picture{}, err
// 	}

// 	// Construct the file URL
// 	fileURL := fmt.Sprintf("https://%s.r2.cloudflarestorage.com/%s/%s", os.Getenv("ACCOUNT_ID"), bucketName, fileKey)
// 	picture := entity.Picture{
// 		Url: fileURL,
// 		Key: fileKey,
// 	}

// 	return picture, nil
// }

//	func (s *pictureServiceImpl) SavePictureToDB(id uint, picture *entity.Picture) error {
//		picture.UserID = id
//		if err := s.repo.SavePictureToDB(id, picture); err != nil {
//			return err
//		}
//		return nil
//	}
