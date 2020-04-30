// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package session

import (
	fsession "github.com/fasthttp/session/v2"
	"github.com/gofiber/fiber"
)

// Store is a wrapper arround the session.Store
type Store struct {
	ctx  *fiber.Ctx
	sess *Session
	core *fsession.Store
}

// ID returns session id
func (s *Store) ID() string {
	return string(s.core.GetSessionID())
}

// Save storage before response
func (s *Store) Save() {
	s.sess.core.Save(s.ctx.Fasthttp, s.core)
}

// Get get data by key
func (s *Store) Get(key string) interface{} {
	return s.core.Get(key)
}

// Set get data by key
func (s *Store) Set(key string, value interface{}) {
	s.core.Set(key, value)
}

// Delete delete data by key
func (s *Store) Delete(key string) {
	s.core.Delete(key)
}

// Destroy session and cookies
func (s *Store) Destroy() {
	s.sess.core.Destroy(s.ctx.Fasthttp)
}

// Regenerate session id
func (s *Store) Regenerate() {
	// https://github.com/fasthttp/session/blob/master/session.go#L205
	s.sess.core.Regenerate(s.ctx.Fasthttp)
}
