package movies

import (
	"testing"
	"os/exec"
)

func TestInitial(t *testing.T){

    got := ""
    want := `{"type":"success","data":null,"message":""}` + "\n"

	cmd := `curl -sS -X GET "http://localhost:8000/movies"`
	bashcmd := exec.Command("bash", "-c", cmd)
    out, err := bashcmd.Output()
    if err != nil {
        t.Errorf("error  : %q" , err)
        return
    }
	got = string(out)

    if got != want {
        t.Errorf("got    : %q\nwanted : %q", got, want)
    }
}

func TestAddMovies(t *testing.T){

    got := ""
    want := `{"type":"success","data":[{"movieid":"4","moviename":"movie4"},{"movieid":"1","moviename":"movie3"},{"movieid":"2","moviename":"movie1"},{"movieid":"3","moviename":"movie2"}],"message":""}` + "\n"

	movieList := []string{
        `movieid=4&moviename=movie4`,
        `movieid=1&moviename=movie3`,
        `movieid=2&moviename=movie1`,
        `movieid=3&moviename=movie2`,
    }
    for _, entry := range movieList {
        cmd := `curl -sS -X POST "http://localhost:8000/movies?` + entry + `"`
        bashcmd := exec.Command("bash", "-c", cmd)
        _, err := bashcmd.Output()
        if err != nil {
            t.Errorf("error  : %q" , err)
            return
        }
    }

    cmd := `curl -sS -X GET "http://localhost:8000/movies"`
	bashcmd := exec.Command("bash", "-c", cmd)
    out, err := bashcmd.Output()
    if err != nil {
        t.Errorf("error  : %q" , err)
        return
    }
	got = string(out)

    if got != want {
        t.Errorf("got    : %q\nwanted : %q", got, want)
    }
}

func TestDelMovieByID(t *testing.T){

    got := ""
    want := `{"type":"success","data":[{"movieid":"4","moviename":"movie4"},{"movieid":"2","moviename":"movie1"}],"message":""}` + "\n"

	movieIdList := []string{
        `1`,
        `3`,
    }
    for _, entry := range movieIdList {
        cmd := `curl -sS -X DELETE "http://localhost:8000/movies/` + entry + `"`
        bashcmd := exec.Command("bash", "-c", cmd)
        _, err := bashcmd.Output()
        if err != nil {
            t.Errorf("error  : %q" , err)
            return
        }
    }

    cmd := `curl -sS -X GET "http://localhost:8000/movies"`
	bashcmd := exec.Command("bash", "-c", cmd)
    out, err := bashcmd.Output()
    if err != nil {
        t.Errorf("error  : %q" , err)
        return
    }
	got = string(out)

    if got != want {
        t.Errorf("got    : %q\nwanted : %q", got, want)
    }
}

func TestDelMovieAll(t *testing.T){

    got := ""
    want := `{"type":"success","data":null,"message":""}` + "\n"

    cmd := `curl -sS -X DELETE "http://localhost:8000/movies"`
	bashcmd := exec.Command("bash", "-c", cmd)
    out, err := bashcmd.Output()
    if err != nil {
        t.Errorf("error  : %q" , err)
        return
    }

    cmd = `curl -sS -X GET "http://localhost:8000/movies"`
	bashcmd = exec.Command("bash", "-c", cmd)
    out, err = bashcmd.Output()
    if err != nil {
        t.Errorf("error  : %q" , err)
        return
    }
	got = string(out)

    if got != want {
        t.Errorf("got    : %q\nwanted : %q", got, want)
    }
}
