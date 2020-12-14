package mysql

import (
    "database/sql"
    "errors"

    "github.com/castevet6/snippetbox/pkg/models"
)

// SnippetModel = wrapper for SQL connection pool
type SnippetModel struct {
    DB *sql.DB
}

// ** CRUD OPERATIONS **

// insert new snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
    // SQL statement variable
    stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

    // use Exec on connection pool to execute the statement
    result, err := m.DB.Exec(stmt, title, content, expires);
    if err != nil {
        return 0, err
    }

    // get id of newly created record with LastInsertId()
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    // convert id (default int64) to int
    return int(id), nil
}

    
// return snippet by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
    // prepared SQL statement
    stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`
    // get single row result with DB.QueryRow(), returns pointer to sql.Row object holding the result
    row := m.DB.QueryRow(stmt, id)

    // initialize empty snippet object
    s := &models.Snippet{}

    // get snippet object fields with row.Scan()
    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
        // if query returns no rows sql.ErrorNoRows, check err status with errors.Is()
        if errors.Is(err, sql.ErrNoRows) {
            return nil, models.ErrNoRecord
        } else {
            return nil, err
        }
    }

    return s, nil
}
    
// get snippet(s) with latest creation date
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
    stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

    // use Query method
    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    // empty slice to hold models.Snippets objects
    snippets := []*models.Snippet{}

    // rows.Next() to iterate through resultset
    for rows.Next() {
        // create pointer to new Snippet struct
        s := &models.Snippet{}
        // rows.Scan() to get field values from query resultset
        err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
        if err != nil {
            return nil, err
        }
        snippets = append(snippets, s)
    }

    // use rows.Err() to get any errors encountered during iteration
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return snippets, nil
}
