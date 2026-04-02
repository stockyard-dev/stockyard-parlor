package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Conversation struct {
	ID string `json:"id"`
	VisitorName string `json:"visitor_name"`
	VisitorEmail string `json:"visitor_email"`
	Subject string `json:"subject"`
	Messages string `json:"messages"`
	Status string `json:"status"`
	Assignee string `json:"assignee"`
	PageURL string `json:"page_url"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"parlor.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS conversations(id TEXT PRIMARY KEY,visitor_name TEXT DEFAULT '',visitor_email TEXT DEFAULT '',subject TEXT DEFAULT '',messages TEXT DEFAULT '[]',status TEXT DEFAULT 'open',assignee TEXT DEFAULT '',page_url TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Conversation)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO conversations(id,visitor_name,visitor_email,subject,messages,status,assignee,page_url,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.VisitorName,e.VisitorEmail,e.Subject,e.Messages,e.Status,e.Assignee,e.PageURL,e.CreatedAt);return err}
func(d *DB)Get(id string)*Conversation{var e Conversation;if d.db.QueryRow(`SELECT id,visitor_name,visitor_email,subject,messages,status,assignee,page_url,created_at FROM conversations WHERE id=?`,id).Scan(&e.ID,&e.VisitorName,&e.VisitorEmail,&e.Subject,&e.Messages,&e.Status,&e.Assignee,&e.PageURL,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Conversation{rows,_:=d.db.Query(`SELECT id,visitor_name,visitor_email,subject,messages,status,assignee,page_url,created_at FROM conversations ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Conversation;for rows.Next(){var e Conversation;rows.Scan(&e.ID,&e.VisitorName,&e.VisitorEmail,&e.Subject,&e.Messages,&e.Status,&e.Assignee,&e.PageURL,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM conversations WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM conversations`).Scan(&n);return n}
