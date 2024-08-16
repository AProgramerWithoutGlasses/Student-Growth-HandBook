package jrx_model

type QueryTeacherParamStruct struct {
	Gender        string `json:"gender"`
	Ban           *bool  `json:"ban"`
	SearchSelect  string `json:"search_select"`
	SearchMessage string `json:"search_message"`
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
}

type QueryTeacherResStruct struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	Ban         *bool  `json:"ban"`
	IsManager   bool   `json:"is_manager"`
	WhatManager string `json:"what_manager"`
}
