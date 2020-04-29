package graphql

import (
	"database/sql"	//Prepare, Exec, Query functions
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"
)

var dbSchema graphql.Schema

func init() {
	//Setup Schema/Server of the fields
	rootQueryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery", 
		Fields: graphql.Fields{
			//given 'tutorial' field and its ID, return tutorial 
			"tutorial": &graphql.Field{
					Type:        tutorialType,	//graphQL returned by this endoint
					Description: "Get Tutorial By ID",
					//Retrieve ID argument of type int
					Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
									Type: graphql.Int,
							},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(int)
						if ok {
							db, err := sql.Open("sqlite3", "graphql/local.db")
							if err != nil {
								log.Fatal(err)
							}
							defer db.Close()
							var tutorial Tutorial
							err = db.QueryRow("SELECT ID, Title FROM tutorials where ID = ?", id).Scan(&tutorial.ID, &tutorial.Title)
							if err != nil {
								fmt.Println(err)
							}
							return tutorial, nil
						}
						return nil, nil
					},
			},
			//given 'list' field, return all tutorials
			"list": &graphql.Field{
					Type:        graphql.NewList(tutorialType),
					Description: "Get Tutorial List",
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, err := sql.Open("sqlite3", "graphql/local.db")
						if err != nil {
							log.Fatal(err)
						}
						defer db.Close()
						var tutorials []Tutorial
						results, err := db.Query("SELECT * FROM tutorials")
						if err != nil {
							fmt.Println(err)
						}
						for results.Next() {	//while the next row of data (results)
							var tutorial Tutorial
							err = results.Scan(&tutorial.ID, &tutorial.Title)	//scan fills the interface with 'results'
							if err != nil {
								fmt.Println(err)
							}
							log.Println(tutorial)
							tutorials = append(tutorials, tutorial)
						}
						return tutorials, nil
					},
			},
		},
	})
	schemaConfig := graphql.SchemaConfig{Query: rootQueryType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	dbSchema = schema
}

func populateDB(){
	database, err := sql.Open("sqlite3", "graphql/local.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	// statement, _ := database.Prepare("CREATE TABLE tutorials (id int, title string);")
	// statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO tutorials VALUES (?, ?)")
	// statement.Exec(1, "First Tutrial")
	// statement.Exec(2, "Second Tutrial")
	// statement.Exec(3, "Third Tutrial")
}

func Test(){
	//populateDB()

	// Query
	query := `
	{
		list {
			id
			title
		}
	}
	`
	query2 := `
    {
        tutorial(id: 1) {
            id
            title
        }
    }
	`
	params := graphql.Params{Schema: dbSchema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	params = graphql.Params{Schema: dbSchema, RequestString: query2}
	r = graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ = json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}