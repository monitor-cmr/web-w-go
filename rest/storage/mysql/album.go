package mysql

import (
    "database/sql"
    "errors"
    "fmt"

    "github.com/monitor-cmr/web-w-go/domain"
)

// AlbumMySQL ...
type AlbumMySQL struct {
    db *sql.DB
}

// NewAlbumMySQL ...
func NewAlbumMySQL(db *sql.DB) *AlbumMySQL {
    return &AlbumMySQL{
        db: db,
    }
}

// SelectAll ...
func (a *AlbumMySQL) SelectAll() ([]domain.Album, error) {
    // SQL statement
    sqlstmt := "SELECT id, title, artist, price FROM albums"

    // run query against sql server
    rows, err := a.db.Query(sqlstmt)

    // Check SQL connection err
    if err != nil {
        fmt.Println("Cannot Connect to SQL Server", err.Error())
        return nil, errors.New("Server Internal Error")
    }

    var (
        // Create empty Album
        record = domain.Album{}
        // Create empty slice Album
        albums = []domain.Album{}
    )

    for rows.Next() {
        err := rows.Scan(&record.ID, &record.Title, &record.Artist, &record.Price)
        // Check scan err
        if err != nil {
            fmt.Println("Error while scanning record", err.Error())
            return nil, errors.New("Server Internal Error")
        }

        // append record to albums slice
        albums = append(albums, record)
    }

    return albums, nil
}

// Select ...
func (a *AlbumMySQL) Select(id int) (*domain.Album, error) {
    var record = domain.Album{}

    // SQL statement
    sqlstmt := "SELECT id, title, artist, price FROM albums where id=?"

    // Single-Row Queries
    err := a.db.QueryRow(sqlstmt, id).Scan(&record.ID, &record.Title, &record.Artist, &record.Price)

    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Record Not Found", err.Error())
            return nil, err
        }

        fmt.Println("Server Scanning Row Error", err.Error())
        return nil, err
    }

    return &record, nil
}

// Save ...
func (a *AlbumMySQL) Save(album domain.Album) (*int, error) {
    // SQL statement
    sqlstmt := "INSERT INTO albums(title, artist, price) VALUES(?, ?, ?)"

    // Execute SQL Statements
    result, err := a.db.Exec(sqlstmt, album.Title, album.Artist, album.Price)

    // Check SQL connection err
    if err != nil {
        fmt.Println("Cannot Connect to SQL Server", err.Error())
        return nil, errors.New("Server Internal Error")
    }

    id, err := result.LastInsertId()

    // error check for LastInsertId function
    if err != nil {
        fmt.Println("Server Get RowID Error", err.Error())
        return nil, errors.New("Server Internal Error")
    }

    // Convert int64 to int
    resultID := int(id)

    return &resultID, nil
}

// Update ...
func (a *AlbumMySQL) Update(album domain.Album) error {
    // SQL statement
    sqlstmt := "UPDATE albums SET title=?, artist=?, price=? where id=?"

    // Execute SQL Statements
    result, err := a.db.Exec(sqlstmt, album.Title, album.Artist, album.Price, album.ID)

    // Check SQL connection err
    if err != nil {
        fmt.Println("Cannot Connect to SQL Server", err.Error())
        return err
    }

    rowCount, err := result.RowsAffected()

    // error check for RowsAffected function
    if err != nil {
        fmt.Println("Server Get RowsAffected Error", err.Error())
        return err
    }

    // there is no row found
    if rowCount == 0 {
        fmt.Println("Record Not Found")
        return sql.ErrNoRows
    }

    return nil

}

// Delete ...
func (a *AlbumMySQL) Delete(id int) error {
    // SQL statement
    sqlstmt := "DELETE FROM albums where id=?"

    // Execute SQL Statements
    result, err := a.db.Exec(sqlstmt, id)

    // Check SQL connection err
    if err != nil {
        fmt.Println("Cannot Connect to SQL Server", err.Error())
        return err
    }

    rowCount, err := result.RowsAffected()

    // error check for RowsAffected function
    if err != nil {
        fmt.Println("Server Get RowsAffected Error", err.Error())
        return err
    }

    // there is no row found
    if rowCount == 0 {
        fmt.Println("Record Not Found")
        return sql.ErrNoRows
    }

    return nil
}

