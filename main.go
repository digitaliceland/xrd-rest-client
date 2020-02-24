package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL   *url.URL
	XRoadClient string
	XRoadService string
	httpClient *http.Client
}


func (c *Client) do(path string) (string, error) {
	service_path := fmt.Sprintf("/r1/%s/%s", c.XRoadService, path)
	rel := &url.URL{Path: service_path}
	u := c.BaseURL.ResolveReference(rel)
	fmt.Println("Accessing: ", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Road-Client", c.XRoadClient)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := string(body)
	return res, err
}

// Here are the api calls wrapped in
func (c *Client) ping() (string, error) {
	return c.do("ping")
}

func (c *Client) time() (string, error) {
	return c.do("time")
}


func main() {
	xclient := flag.String("client", "CS/ORG/1111/TestClient", "Your X Road Client ID")
	xservice := flag.String("service", "CS/ORG/1111/TestService/TEST123", "Your X Road Service ID")
	ss := flag.String("ss", "http://localhost:80", "Your X Road Security Server URL without trailing /")
	cmd := flag.String("cmd", "", "method to call on your REST server")
	loop := flag.Bool("loop", false, "repeatedly call the function every second")
	flag.Parse()

	u, err := url.Parse(*ss)
	if err != nil {
		fmt.Println("Error parsing service URL: ", ss)
		log.Fatal(err)
	}
	// Now u is a valid parsed url for the security server

	c := &Client{
		BaseURL:    u,
		XRoadClient:  *xclient,
		XRoadService:  *xservice,
		httpClient: http.DefaultClient,
	}
	// We have created the client object and filled in all neccesary data to query the API

	switch *cmd {
	case "ping":
		if *loop {
			for  {
				reply, _ := c.ping()
				fmt.Println("Ping returned: ", reply)
				time.Sleep(1 * time.Second)
			}
		} else {
			reply, _ := c.ping()
			fmt.Println("Ping returned: ", reply)
		}
	case "time":
		if *loop {
			for  {
				reply, _ := c.time()
				fmt.Println("Time returned: ", reply)
				time.Sleep(1 * time.Second)
			}
		} else {
			reply, _ := c.time()
			fmt.Println("Time returned: ", reply)
		}
	default:
		fmt.Println("No command specified")
		flag.Usage()
	}
}