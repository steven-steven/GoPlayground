package graphql

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func SampleServer1() {
	// Create GraphQL server -> schema
	fields := graphql.Fields{
			"hello": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return "world", nil
					},
			},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
			log.Fatalf("failed to create new schema, error: %v", err)
	}

	//-----------------------
	// Query the schema with a query asking for "hello"
	query := `
			{
					hello
			}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}


//-------------Second Example (return object, not just strings)------------------

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

//populate and return the data (list of tutorials with 1 value)
func populate() []Tutorial {
	author := &Author{Name: "Elliot Forbes", Tutorials: []int{1}}
	tutorial := Tutorial{
			ID:     1,
			Title:  "Go GraphQL Tutorial",
			Author: *author,
			Comments: []Comment{
					Comment{Body: "First Comment"},
			},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)

	return tutorials
}

//---------------------------Define Comment, Author, Tutorial GraphQL Object Types ------------
var commentType = graphql.NewObject(
	graphql.ObjectConfig{	//pass ObjectConfig struct with name and fields of our object
			Name: "Comment",
			Fields: graphql.Fields{
					"body": &graphql.Field{
							Type: graphql.String,
					},
			},
	},
)
var authorType = graphql.NewObject(
	graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{	//object contains name and []int
					"Name": &graphql.Field{
							Type: graphql.String,
					},
					"Tutorials": &graphql.Field{
							Type: graphql.NewList(graphql.Int),	//array type in grapql is NewList
					},
			},
	},
)
var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
					"id": &graphql.Field{
							Type: graphql.Int,
					},
					"title": &graphql.Field{
							Type: graphql.String,
					},
					"author": &graphql.Field{
							Type: authorType,	//our defined graphql object type
					},
					"comments": &graphql.Field{
							Type: graphql.NewList(commentType),
					},
			},
	},
)

//---------------Update Schema ----------
func SampleServer2(){

	tutorials:= populate()

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
							// ResolveParams object contains the argument
							id, ok := p.Args["id"].(int)
							if ok {
									// Query for the data and return it
									for _, tutorial := range tutorials {
											if int(tutorial.ID) == id {
													return tutorial, nil
											}
									}
							}
							return nil, nil
					},
			},
			//given 'list' field, return all tutorials
			"list": &graphql.Field{
					Type:        graphql.NewList(tutorialType),
					Description: "Get Tutorial List",
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
							return tutorials, nil
					},
			},
		},
	})
	//Mutation Type/Schema
	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",		//name the schema object 'Mutation' (unlike RootQuery)				
		Fields: graphql.Fields{	
				"create": &graphql.Field{		//define a field called 'create' (like all the query fields)
						Type:        tutorialType,
						Description: "Create a new Tutorial",
						Args: graphql.FieldConfigArgument{	//recieves non-null string
								"title": &graphql.ArgumentConfig{
										Type: graphql.NewNonNull(graphql.String),
								},
						},
						Resolve: func(params graphql.ResolveParams) (interface{}, error) {
								tutorial := Tutorial{		//create tutorial and append	
										Title: params.Args["title"].(string),
								}
								tutorials = append(tutorials, tutorial)
								return tutorial, nil
						},
				},
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: rootQueryType, Mutation: mutationType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
			log.Fatalf("failed to create new schema, error: %v", err)
	}

	//-----------------------
	// Query the schema with a query.
	//Consider a 'rootQuery' object, we want list field of the object (in the field specify the attributes of the returned schema)
	queryAllTutorial := `
		{
			list {
				id
				title
				comments {
						body
				}
				author {
						Name
						Tutorials
				}
			}
		}
	`
	queryTutorialById := `
		{
			tutorial(id:1) {
				title
				author {
						Name
						Tutorials
				}
			}
		}
	`
	queryToMutate := `
		mutation {
			create(title: "Hello World") {
				title
			}
		}
	`
	//call query 1
	params := graphql.Params{Schema: schema, RequestString: queryAllTutorial}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	//call query 2
	params2 := graphql.Params{Schema: schema, RequestString: queryTutorialById}
	r2 := graphql.Do(params2)
	if len(r2.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r2.Errors)
	}
	rJSON2, _ := json.Marshal(r2)
	fmt.Printf("%s \n", rJSON2)

	//call query 3
	params3 := graphql.Params{Schema: schema, RequestString: queryToMutate}
	r3 := graphql.Do(params3)
	if len(r3.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r3.Errors)
	}
	rJSON3, _ := json.Marshal(r3)
	fmt.Printf("%s \n", rJSON3)
}
