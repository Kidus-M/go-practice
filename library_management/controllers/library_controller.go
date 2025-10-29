package controllers
bStr = strings.TrimSpace(bStr)
bid, _ := strconv.Atoi(bStr)
fmt.Print("Member ID: ")
mStr, _ := r.ReadString('\n')
mStr = strings.TrimSpace(mStr)
mid, _ := strconv.Atoi(mStr)
err := c.lib.BorrowBook(bid, mid)
if err != nil {
fmt.Println("Error:", err)
return
}
fmt.Println("Book borrowed")
}


func (c *Controller) returnBook(r *bufio.Reader) {
fmt.Print("Book ID to return: ")
bStr, _ := r.ReadString('\n')
bStr = strings.TrimSpace(bStr)
bid, _ := strconv.Atoi(bStr)
fmt.Print("Member ID: ")
mStr, _ := r.ReadString('\n')
mStr = strings.TrimSpace(mStr)
mid, _ := strconv.Atoi(mStr)
err := c.lib.ReturnBook(bid, mid)
if err != nil {
fmt.Println("Error:", err)
return
}
fmt.Println("Book returned")
}


func (c *Controller) listAvailable() {
books := c.lib.ListAvailableBooks()
if len(books) == 0 {
fmt.Println("No available books")
return
}
for _, b := range books {
fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
}
}


func (c *Controller) listBorrowed(r *bufio.Reader) {
fmt.Print("Member ID: ")
mStr, _ := r.ReadString('\n')
mStr = strings.TrimSpace(mStr)
mid, _ := strconv.Atoi(mStr)
books := c.lib.ListBorrowedBooks(mid)
if len(books) == 0 {
fmt.Println("No borrowed books for this member or member not found")
return
}
for _, b := range books {
fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
}
}