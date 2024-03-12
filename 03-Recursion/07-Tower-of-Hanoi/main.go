package main

import (
	"fmt"
)

const numDisks = 3

// Add a disk to the beginning of the post.
func push(post []int, disk int) []int {
	post = append([]int{disk}, post...)
	return post
}

// Remove the first disk from the post.
// Return that disk and the revised post.
func pop(post []int) (int, []int) {
	return post[0], post[1:]
}

// Move one disk from fromPost to toPost.
func moveDisk(posts [][]int, fromPost, toPost int) {
	disk, post := pop(posts[fromPost])
	fmt.Printf("Move disk %d: from %d to %d\n", disk, fromPost, toPost)
	posts[fromPost] = post
	post = push(posts[toPost], disk)
	posts[toPost] = post
	
	drawPosts(posts)
}

// Draw the posts by showing the size of the disk at each level.
func drawPosts(posts [][]int) {
	for row := 0; row < numDisks; row++ {
		for i := range posts {
			if row < numDisks - len(posts[i]) {
				fmt.Printf("0 ")
			} else {
				fmt.Printf("%d ", posts[i][row - (numDisks - len(posts[i]))])
			}
		}
		fmt.Println()
	}
}

// Move the disks from fromPost to toPost
// using tempPost as temporary storage.
func moveDisks(posts [][]int, numToMove, fromPost, toPost, tempPost int) {
	if numToMove > 1 {
		moveDisks(posts, numToMove - 1, fromPost, tempPost, toPost)
	}
	moveDisk(posts, fromPost, toPost)
	if numToMove > 1 {
		moveDisks(posts, numToMove - 1, tempPost, toPost, fromPost)
	}
}

func main() {
    // Make three posts.
    posts := [][]int { }

    // Push the disks onto post 0 biggest first.
    posts = append(posts, []int{})
    for disk := numDisks; disk > 0; disk-- {
        posts[0] = push(posts[0], disk)
    }

    // Make the other posts empty.
    for p := 1; p < 3; p++ {
        posts = append(posts, []int{})
    }

    // Draw the initial setup.
    drawPosts(posts)

    // Move the disks.
    moveDisks(posts, numDisks, 0, 2, 1)
}

