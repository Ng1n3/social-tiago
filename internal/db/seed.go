package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Ng1n3/social/internal/store"
)

var usernames = []string{
	"Aisha",
	"David",
	"Fatima",
	"John",
	"Halima",
	"Samuel",
	"Amina",
	"Michael",
	"Ngozi",
	"Joseph",
	"Adanna",
	"Daniel",
	"Zainab",
	"Emmanuel",
	"Hadiza",
	"Victor",
	"Chioma",
	"Joshua",
	"Aishat",
	"Peter",
	"Blessing",
	"Isaac",
	"Khadija",
	"Gabriel",
	"Ruth",
	"Caleb",
	"Mariam",
	"Elijah",
	"Hannah",
	"Abraham",
	"Deborah",
	"Solomon",
	"Esther",
	"Moses",
	"Sarah",
	"Ezekiel",
	"Mercy",
	"Isaiah",
	"Peace",
	"Jeremiah",
	"Joy",
	"Grace",
	"Abigail",
	"Precious",
	"Miracle",
	"Paul",
	"Faith",
}

var titles = []string{
	"Go for Beginners",
	"Web Dev with Go",
	"Building APIs in Go",
	"Go Concurrency",
	"Testing Go Code",
	"Go Modules Explained",
	"Working with Databases",
	"Go and Microservices",
	"Effective Go",
	"Go Performance Tips",
	"Go Best Practices",
	"Go Design Patterns",
	"Go and DevOps",
	"Go for Data Science",
	"Go for Machine Learning",
	"Go and Cloud Computing",
	"Go and IoT",
	"Go and Security",
	"Go Community",
	"Go Resources",
}

var contents = []string{
	"Go is a powerful and versatile programming language.",
	"It is known for its simplicity and efficiency.",
	"Go is excellent for building web applications.",
	"Concurrency in Go is easy to manage.",
	"Testing is an important part of the development process.",
	"Modules help organize Go code.",
	"Go can be used to interact with various databases.",
	"Microservices are a popular architectural pattern.",
	"Writing effective Go code requires practice.",
	"Performance optimization is crucial for production systems.",
	"Following best practices improves code quality.",
	"Design patterns can make code more maintainable.",
	"Go is often used in DevOps environments.",
	"Go is gaining traction in the data science community.",
	"Machine learning libraries are available for Go.",
	"Go is well-suited for cloud computing.",
	"IoT devices can be programmed with Go.",
	"Security is a critical aspect of software development.",
	"The Go community is very supportive.",
	"Many helpful resources are available for Go developers.",
}

var tags = []string{
	"go",
	"golang",
	"programming",
	"webdev",
	"backend",
	"api",
	"rest",
	"microservices",
	"cloud",
	"aws",
	"gcp",
	"azure",
	"database",
	"sql",
	"nosql",
	"concurrency",
	"testing",
	"devops",
	"security",
	"performance",
	"bestpractices",
	"designpatterns",
	"modules",
	"generics",
	"goroutines",
	"channels",
	"structs",
	"interfaces",
	"http",
	"json",
	"encoding",
	"algorithms",
	"datastructures",
	"debugging",
	"cli",
	"grpc",
	"protobuf",
	"kubernetes",
	"docker",
	"ci/cd",
}

var comments = []string{
	"Great post!",
	"This was really helpful, thanks!",
	"I learned something new today.",
	"Excellent article, keep up the good work!",
	"Thanks for sharing your insights.",
	"This is exactly what I was looking for.",
	"I have a question about this...",
	"Could you elaborate on that?",
	"Interesting perspective.",
	"I disagree with this point...",
	"This is a bit confusing.",
	"Needs more examples.",
	"Overall, a good read.",
	"Looking forward to more content like this.",
	"Keep writing!",
}

func Seed(store store.Storage) {
	ctx := context.Background()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	log.Println("Cleaning existing data...")
	if err := cleanExistingData(ctx, store); err != nil {
		log.Println("Error cleaning existing data: ", err)
		return
	}

	users := generateUsers(100)
	log.Println("Creating users...")
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePost(200, users)
	log.Println("Creating posts...")
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	comments := generateComments(20, users, posts)
	log.Println("Creating comments...")
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		randomSuffix := fmt.Sprintf("%d_%d", i, rand.Intn(100))
		username := usernames[i%len(usernames)] + randomSuffix
		users[i] = &store.User{
			// Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", 1),
			// Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", 1) + "@example.com",
			Username: username,
			Email:    username + "@example.com",
			Password: "12345",
		}
	}

	return users
}

func generatePost(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {

	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}

func cleanExistingData(ctx context.Context, store store.Storage) error {
	if err := store.Comments.DeleteSeedAll(ctx); err != nil {
		return fmt.Errorf("cleaning comments: %w", err)
	}
	if err := store.Posts.DeleteSeedAll(ctx); err != nil {
		return fmt.Errorf("cleaning posts: %w", err)
	}
	if err := store.Users.DeleteSeedAll(ctx); err != nil {
		return fmt.Errorf("cleaning users: %w", err)
	}
	return nil
}
