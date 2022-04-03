package main

import (
	"os/exec"
	"testing"

	"github.com/ipatser/utils"
)

var (
	TEST_IP = utils.GetCfgVar("TEST_IP")
)

/*
TEST_IP
- local deployment  : "localhost"
- docker deployment : "ipatser_app_deploy"
*/

func TestMoviesInitial(t *testing.T) {

	got := ""
	want := `{"type":"success","data":null,"message":""}` + "\n"

	cmd := `curl -sS -X GET "http://` + TEST_IP + `:8000/movies"`
	shcmd := exec.Command("sh", "-c", cmd)
	out, err := shcmd.Output()
	if err != nil {
		t.Errorf("error  : %q", err)
		return
	}
	got = string(out)

	if got != want {
		t.Errorf("got    : %q\nwanted : %q", got, want)
	}
}

func TestMoviesAdd(t *testing.T) {

	got := ""
	want := `{"type":"success","data":[{"movieid":"4","moviename":"movie4"},{"movieid":"1","moviename":"movie3"},{"movieid":"2","moviename":"movie1"},{"movieid":"3","moviename":"movie2"}],"message":""}` + "\n"

	movieList := []string{
		`movieid=4&moviename=movie4`,
		`movieid=1&moviename=movie3`,
		`movieid=2&moviename=movie1`,
		`movieid=3&moviename=movie2`,
	}
	for _, entry := range movieList {
		cmd := `curl -sS -X POST "http://` + TEST_IP + `:8000/movies?` + entry + `"`
		shcmd := exec.Command("sh", "-c", cmd)
		_, err := shcmd.Output()
		if err != nil {
			t.Errorf("error  : %q", err)
			return
		}
	}

	cmd := `curl -sS -X GET "http://` + TEST_IP + `:8000/movies"`
	shcmd := exec.Command("sh", "-c", cmd)
	out, err := shcmd.Output()
	if err != nil {
		t.Errorf("error  : %q", err)
		return
	}
	got = string(out)

	if got != want {
		t.Errorf("\ngot    : %q\nwanted : %q", got, want)
	}
}

func TestMoviesDelByID(t *testing.T) {

	got := ""
	want := `{"type":"success","data":[{"movieid":"4","moviename":"movie4"},{"movieid":"2","moviename":"movie1"}],"message":""}` + "\n"

	movieIdList := []string{
		`1`,
		`3`,
	}
	for _, entry := range movieIdList {
		cmd := `curl -sS -X DELETE "http://` + TEST_IP + `:8000/movies/` + entry + `"`
		shcmd := exec.Command("sh", "-c", cmd)
		_, err := shcmd.Output()
		if err != nil {
			t.Errorf("error  : %q", err)
			return
		}
	}

	cmd := `curl -sS -X GET "http://` + TEST_IP + `:8000/movies"`
	shcmd := exec.Command("sh", "-c", cmd)
	out, err := shcmd.Output()
	if err != nil {
		t.Errorf("error  : %q", err)
		return
	}
	got = string(out)

	if got != want {
		t.Errorf("\ngot    : %q\nwanted : %q", got, want)
	}
}

func TestMoviesDelAll(t *testing.T) {

	got := ""
	want := `{"type":"success","data":null,"message":""}` + "\n"

	cmd := `curl -sS -X DELETE "http://` + TEST_IP + `:8000/movies"`
	shcmd := exec.Command("sh", "-c", cmd)
	out, err := shcmd.Output()
	if err != nil {
		t.Errorf("error  : %q", err)
		return
	}

	cmd = `curl -sS -X GET "http://` + TEST_IP + `:8000/movies"`
	shcmd = exec.Command("sh", "-c", cmd)
	out, err = shcmd.Output()
	if err != nil {
		t.Errorf("error  : %q", err)
		return
	}
	got = string(out)

	if got != want {
		t.Errorf("got    : %q\nwanted : %q", got, want)
	}
}

/*
##########
# Benchmark Functions
##########
*/

func BenchmarkMoviesAdd(b *testing.B) {
	/*
	   Issue:
	   pq: remaining connection slots are reserved for non-replication superuser connections
	*/
	movieList := []string{
		`movieid=4&moviename=movie4`,
		`movieid=1&moviename=movie3`,
		`movieid=2&moviename=movie1`,
		`movieid=3&moviename=movie2`,
	}
	for _, entry := range movieList {
		cmd := `curl -sS -X POST "http://` + TEST_IP + `:8000/movies?` + entry + `"`
		shcmd := exec.Command("sh", "-c", cmd)

		_, err := shcmd.Output()
		if err != nil {
			b.Errorf("error  : %q", err)
			return
		}
	}
}

func BenchmarkMoviesDelAll(b *testing.B) {
	cmd := `curl -sS -X DELETE "http://` + TEST_IP + `:8000/movies"`
	shcmd := exec.Command("sh", "-c", cmd)
	_, err := shcmd.Output()
	if err != nil {
		b.Errorf("error  : %q", err)
		return
	}
}
