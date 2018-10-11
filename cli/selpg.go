package main
import (
  "fmt"
  flag "github.com/spf13/pflag"
  "os"
  "os/exec"
  "bufio"
  "strings"
  "io"
)

type selpgArgs struct {
	start, end int
	in_filename string
	page_len int //the number of lines in a page
	page_type string //seperate pages with fixed number of lines or specified symbol
	print_dest string //print to screen or file
}



func main() {

  sa := new(selpgArgs)

  process_args(sa)
  fmt.Println(*sa)
  process_input(*sa)

  /*

  var port int
  flag.IntVarP(&port, "p","p", 8000, "defaults 8000")
  flag.Parse()

  fmt.Printf("port = %d\n", port)
  fmt.Printf("other args %+v\n", flag.Args())
  */
}


func process_args(sa *selpgArgs){
	flag.IntVarP(&sa.start,"s","s",-1,"the start page")
	flag.IntVarP(&sa.end,"e","e",-1,"the end page")
	flag.IntVarP(&sa.page_len,"l","l",72,"default page number")
	flag.StringVarP(&sa.print_dest,"d","d","","printer")

	has_f := flag.BoolP("f","f",false,"")
	flag.Parse()

	if *has_f{
		sa.page_type = "f"
		sa.page_len = -1
	} else{
		sa.page_type = "l"
	}

	if flag.NArg() ==1 {
		sa.in_filename = flag.Arg(0)
	} else {
		sa.in_filename = ""
	}

	page_ok :=sa.start >= 1 && sa.start<=sa.end
	param_num_ok := flag.NArg() ==1 || flag.NArg() ==0
	page_type_ok := !(sa.page_type == "f" && sa.page_len != -1)
	if(!page_ok || !param_num_ok || !page_type_ok){
		usage()
		os.Exit(1)
	}
}

func process_input(sa selpgArgs){
	line_counter := 0
	page_counter := 1

	fin := os.Stdin
  fout := os.Stdout

  var inpipe io.WriteCloser
  var err error

  if sa.in_filename != "" {
    fin, err := os.Open(sa.in_filename)
    if err != nil {
      fmt.Fprintf(os.Stderr, "selpg: could not open input file \"%s\" \n", sa.in_filename)
      fmt.Println(err)
      usage()
      os.Exit(1)
    }
    defer fin.Close()
  }

  if sa.print_dest != "" {
    cmd := exec.Command("grep", "-nf", "keyword")
    inpipe, err = cmd.StdinPipe()
    if err !=nil {
      fmt.Println(err)
      os.Exit(1)
    }
    defer inpipe.Close()
    cmd.Stdout = fout
    cmd.Start()
  }

  if sa.page_type == "l" {
    line := bufio.NewScanner(fin)
    for line.Scan() {
      if page_counter >= sa.start && page_counter <= sa.end {
        fout.Write([]byte(line.Text() + "\n"))
        if sa.print_dest != "" {
          inpipe.Write([]byte(line.Text() + "\n"))
        }
      }
      line_counter++
      if line_counter%sa.page_len ==0 {
        page_counter++
        line_counter =0
      }
    }
  } else {
    rd := bufio.NewReader(fin)
    for {
      page, ferr := rd.ReadString('\f')
      if ferr !=nil || ferr == io.EOF {
        if ferr == io.EOF {
          if page_counter >= sa.start && page_counter <= sa.end {
            fmt.Fprintf(fout, "%s", page)
          }
        }
        break
      }
      page = strings.Replace(page, "\f", "", -1)
      page_counter++
      if page_counter >= sa.start && page_counter <= sa.end {
        fmt.Fprintf(fout, "%s", page)
      }
    }
  
  if page_counter < sa.end {
    fmt.Fprintf(os.Stderr, "./selpg: end_page (%d) greater than total pages (%d), less output than expected\n", sa.end, page_counter)
    }
  }
}

func usage(){
fmt.Fprintf(os.Stderr, "\nUSAGE: ./selpg [-s start_page] [-e end_page] [ -l lines_per_page | -f ] [ -d dest ] [ in_filename ]\n")
}
