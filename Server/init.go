package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Server_config struct {
	listen_address  string
	auth_method     byte
	timeout         int
	log_level       string
	max_connections int
	user_tls        bool
	dns_resolver    string
	free_ports      [3]uint16
	paid_ports      [3]uint16
}

func config_from(configFile string) *Server_config {

	readFile, err := os.Open(configFile)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	conf_scan := bufio.NewScanner(readFile)
	conf_scan.Split(bufio.ScanLines)

	var params map[string]interface{}

	for conf_scan.Scan() {

		micro_lines := strings.Split(conf_scan.Text(), ":")

		params[micro_lines[0]] = micro_lines[1]
	}

	return &Server_config{
		listen_address:  params["addr"].(string),
		auth_method:     params["auth"].(byte),
		timeout:         params["timeout"].(int),
		log_level:       params["logs"].(string),
		max_connections: params["m_con"].(int),
		user_tls:        params["tls"].(bool),
		dns_resolver:    params["dns_res"].(string),
		free_ports:      params["f_p"].([3]uint16),
		paid_ports:      params["p_p"].([3]uint16),
	}
}
