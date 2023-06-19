// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package postgresql

import (
	"fmt"
	"testing"

	"database/sql"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/dbtest"
	_ "github.com/lib/pq"
)

func TestPG(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			db, err := sql.Open("postgres", "host=127.0.0.1 port=5433 user=postgres password=073019 sslmode=disable")
			db.Exec("\\c ethereum")
			db.Exec("drop table kvs")
			db.Exec("create table kvs (key char(20), value char(20))")
			// db.Exec("create index key on kvs (key)")

			if err != nil {
				t.Fatal(err)
				fmt.Println("Open ERROR1")
			}
			if db == nil {
				fmt.Println("Open ERROR2")
			}
			return &Database{
				db: db,
			}
		})
	})
}

func BenchmarkPG(b *testing.B) {
	dbtest.BenchDatabaseSuite(b, func() ethdb.KeyValueStore {
		db, err := sql.Open("postgres", "host=127.0.0.1 port=5433 user=postgres password=073019 sslmode=disable")
		db.Exec("\\c ethereum")
		db.Exec("drop table kvs")
		db.Exec("create table kvs (key char(20), value char(20))")
		db.Exec("create index key on kvs (key)")
		if err != nil {
			b.Fatal(err)
		}
		return &Database{
			db: db,
		}
	})
}
