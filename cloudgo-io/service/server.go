package service 
import(
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

func NewServer () *negroni.Negroni{
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n:= negroni.Classic()

	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)

	return n;
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root,err := os.Getwd(); err !=nil{
			panic("could not retrieve working directory")
		} else {
			webRoot = root 
		}

	}
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.HandleFunc("/js", jsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/",indexHandler()).Methods("GET")
	mx.HandleFunc("/table",tableHandler()).Methods("POST")
	mx.NotFoundHandler = NotImplementedHandler()
}