package models

type Student struct {
    ID       int    `json:"id_students"`
    NIS      int    `json:"nis"`
    Name     string `json:"name"`
    Gender   string `json:"gender"`
} 

