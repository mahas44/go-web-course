package cookie

import (
	"fmt"
	"net/http"
)

func Cookie() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookies := req.Cookies()

	for i, c := range cookies {
		fmt.Fprintln(w, "YOUR COOKIE ", i, ":", c)
	}

	// c1, err := req.Cookie("my-cookie")
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	// }

	// c2, err := req.Cookie("general")
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Fprintln(w, "YOUR COOKIE #1:", c2)
	// }

	// c3, err := req.Cookie("specific")
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Fprintln(w, "YOUR COOKIE #1:", c3)
	// }

}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:   "specific",
		Value:  "some other value about specific things",
		MaxAge: 0,
		Secure: true,
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}
