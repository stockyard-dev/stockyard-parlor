package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-parlor/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.ListConversations();if list==nil{list=[]store.Conversation{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var req struct{Visitor string `json:"visitor"`;Page string `json:"page"`};json.NewDecoder(r.Body).Decode(&req);if req.Visitor==""{req.Visitor="anonymous"};if req.Page==""{req.Page="/"};conv,_:=s.db.CreateConversation(req.Visitor,req.Page);writeJSON(w,201,conv)}
func(s *Server)handleGetMessages(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.GetMessages(id);if list==nil{list=[]store.Message{}};writeJSON(w,200,list)}
func(s *Server)handleAddMessage(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var m store.Message;json.NewDecoder(r.Body).Decode(&m);m.ConversationID=id;if m.Body==""{writeError(w,400,"body required");return};if m.Sender==""{m.Sender="visitor"};s.db.AddMessage(&m);writeJSON(w,201,m)}
func(s *Server)handleClose(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.CloseConversation(id);writeJSON(w,200,map[string]string{"status":"closed"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
