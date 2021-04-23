package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `json:"author_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("Title Required")
	}
	if p.Content == "" {
		return errors.New("Content Required")
	}
	if p.AuthorID < 1 {
		return errors.New("Author Required")
	}
	return nil
}
